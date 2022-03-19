package common

import (
	"net/url"

	"github.com/firdasafridi/parkinglot/lib/common/commonerr"
	"github.com/gorilla/schema"
)

const (
	YYYYMMDDDash = "2006-01-02"
)

func DecodeSchema(values url.Values, val interface{}) error {
	decoder := schema.NewDecoder()
	if err := decoder.Decode(val, values); err != nil {
		return commonerr.SetNewBadRequest("URL Params", err.Error())
	}
	return nil
}
