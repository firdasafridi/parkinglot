/**
** TODO: #2 Implementation middleware HTTP
**/
package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/firdasafridi/parkinglot/internal/config"
	"github.com/firdasafridi/parkinglot/lib/common/log"
	"github.com/firdasafridi/parkinglot/lib/util/nr"
	"github.com/go-chi/chi"
	newrelic "github.com/newrelic/go-agent"
)

type Config struct {
	Server config.Server
}

// HandlerCustomMetricsTimeout will help handling custom metrics timeout and put the data based on new tag that define
func (cfg *Config) HandlerCustomMetricsTimeout(next http.Handler) http.Handler {
	app := nr.GetApp()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx        = r.Context()
			tNow       = time.Now()
			method     = r.Method
			errMessage = ""
			isError    bool
		)

		// initial context with code 200
		ctx = SetHttpCode(ctx, 200, nil)

		// run serving http
		next.ServeHTTP(w, r.WithContext(ctx))

		// path add router path pattern
		path := chi.RouteContext(ctx).RoutePattern()

		// calculate response time
		then := time.Since(tNow).Milliseconds()

		// get code and response time
		code := GetHttpCode(ctx)

		// handle error message
		if code.Error != nil {
			errMessage = code.Error.Error()
			isError = true
		}

		// send custom http tracing
		if err := app.RecordCustomEvent("HTTPTracing", map[string]interface{}{
			"method":           method,
			"path":             path,
			"http_code":        code.Code,
			"error":            errMessage,
			"is_error":         isError,
			"response_time_ms": then,
		}); err != nil {
			log.Errorln("app.RecordCustomEvent", err)
		}

	})
}

func (cfg *Config) HandlerNR(next http.Handler) http.Handler {
	app := nr.GetApp()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		pattern := chi.RouteContext(r.Context()).RoutePattern()

		// get transaction newrelic
		txn := app.StartTransaction(pattern, w, r)

		ctx := SetTrxKey(r.Context(), txn)
		defer txn.End()

		r = newrelic.RequestWithTransactionContext(r.WithContext(ctx), txn)

		next.ServeHTTP(w, r)

	})
}

type httpCode struct {
	Error error
	Code  int
}

type nrTrx struct {
	Error error
	Code  int
}

var (
	httpCodeKey = httpCode{}
	nrTrxKey    = nrTrx{}
)

// GetHttpCode get structure http code and handling error
func GetHttpCode(ctx context.Context) (code *httpCode) {
	if ctx == nil {
		return
	}

	codePnt, ok := ctx.Value(httpCodeKey).(*httpCode)
	if !ok {
		return
	}
	return codePnt
}

// SetHttpCode set structure http code and handling error
func SetHttpCode(ctx context.Context, code int, err error) context.Context {
	if ctx == nil {
		return nil
	}

	codePnt, ok := ctx.Value(httpCodeKey).(*httpCode)
	if !ok {
		return context.WithValue(ctx, httpCodeKey, &httpCode{
			Error: err,
			Code:  code,
		})
	}

	codePnt.Code = code
	codePnt.Error = err
	return context.WithValue(ctx, httpCodeKey, codePnt)
}

// GeTrxKey get transaction newrelic apps
func GeTrxKey(ctx context.Context) (nrTrxValue newrelic.Transaction) {
	if ctx == nil {
		return
	}

	nrTrxValue, ok := ctx.Value(nrTrxKey).(newrelic.Transaction)
	if !ok {
		return
	}
	return nrTrxValue
}

// SetTrxKey set transaction newrelic apps
func SetTrxKey(ctx context.Context, tx newrelic.Transaction) context.Context {
	if ctx == nil {
		return nil
	}

	return context.WithValue(ctx, nrTrxKey, tx)
}
