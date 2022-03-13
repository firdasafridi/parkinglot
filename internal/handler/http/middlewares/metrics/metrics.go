package metrics

import (
	"log"
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type NewRelic struct {
	AppName    string
	LicenseKey string
}

type Middleware struct {
	NewRelic
}

var nra *newrelic.Application

func New(mw *Middleware) *Middleware {
	var err error

	nra, err = newrelic.NewApplication(
		newrelic.ConfigAppName(mw.AppName),
		newrelic.ConfigLicense(mw.LicenseKey),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		log.Fatalln("[METRIC_MIDDLEWARE]", err)
	}

	return mw
}

func (m *Middleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("[METRIC_MIDDLEWARE]", r.URL.Path)

		next.ServeHTTP(w, r.WithContext(r.Context()))

		txn := nra.StartTransaction(r.URL.Path)
		defer txn.End()

		txn.SetWebRequestHTTP(r)

		// TODO:
		// req := reqinfo.GetRequestInfo(r.Context())
		// req.GetLatency()
	})
}
