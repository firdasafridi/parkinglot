package countries

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	domaincounties "github.com/firdasafridi/parkinglot/internal/entity/countries"
	"github.com/firdasafridi/parkinglot/internal/handler/middleware"
	newrelic "github.com/newrelic/go-agent"
	"github.com/pkg/errors"
)

const (
	url     = "https://restcountries.com"
	version = "v3.1"
	path    = "name"

	cityURL = "https://countriesnow.space/api/v0.1/countries/population/cities"
)

var (
	getCountriesURL = fmt.Sprintf("%s/%s/%s", url, version, path)
)

func RequestCountry(ctx context.Context, country string) (resp domaincounties.Response, err error) {
	url := fmt.Sprintf("%s/%s", getCountriesURL, country)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return resp, errors.Wrap(err, "RequestCountry.NewRequestWithContext")
	}

	txn := middleware.GeTrxKey(ctx)

	txSegment := newrelic.StartExternalSegment(txn, req)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return resp, errors.Wrap(err, "RequestCountry.Client.Do")
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return resp, errors.Wrap(err, "RequestCountry.ioutil.ReadAll")
	}

	countries := make([]domaincounties.ResponseCountry, 0)
	if err = json.Unmarshal(body, &countries); err != nil {
		return resp, errors.Wrap(err, "RequestCountry.Unmarshal")
	}

	txSegment.Response = res
	txSegment.End()

	resp.Countries = countries

	cityBody, _ := json.Marshal(&domaincounties.City{
		City: countries[0].Capital[0],
	})

	req2, err := http.NewRequestWithContext(ctx, http.MethodPost, cityURL, bytes.NewBuffer(cityBody))
	if err != nil {
		return resp, errors.Wrap(err, "RequestCountry.NewRequestWithContext")
	}
	req2.Header.Set("Content-Type", "application/json")

	txSegment2 := newrelic.StartExternalSegment(txn, req2)

	res2, err := client.Do(req2)
	if err != nil {
		return resp, errors.Wrap(err, "RequestCountry.Client.Do")
	}

	defer res2.Body.Close()

	body2, err := ioutil.ReadAll(res2.Body)
	if err != nil {
		return resp, errors.Wrap(err, "RequestCountry.ioutil.ReadAll")
	}

	city := domaincounties.CityResponse{}
	if err = json.Unmarshal(body2, &city); err != nil {
		return resp, errors.Wrap(err, "RequestCountry.Unmarshal")
	}

	resp.City = city

	txSegment2.Response = res2
	txSegment2.End()

	return resp, nil
}
