package parking

import (
	"context"

	domaincounties "github.com/firdasafridi/parkinglot/internal/entity/countries"
	"github.com/firdasafridi/parkinglot/internal/repo/api/countries"
)

type countriesUC interface {
	GetCounties(ctx context.Context, country string) (detailCountries []domaincounties.ResponseCountry, err error)
}

func (uc *Parking) GetCounties(ctx context.Context, country string) (detailCountries []domaincounties.ResponseCountry, err error) {
	return countries.RequestCountry(ctx, country)
}
