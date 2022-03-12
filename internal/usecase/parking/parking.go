package parking

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	parkingdomain "github.com/firdasafridi/parkinglot/internal/entity/parking"
	parkingdb "github.com/firdasafridi/parkinglot/internal/repo/db/parking"
	"github.com/firdasafridi/parkinglot/lib/common/commonerr"
)

type ParkingUC interface {
	GetAllParkingData(ctx context.Context) (listTrxParking []*parkingdomain.TrxParking, err error)
}

type Parking struct {
	ParkingDB parkingdb.ParkingLotDB
}

func New(parking *Parking) *Parking {
	return parking
}

func (uc *Parking) GetAllParkingData(ctx context.Context) (listTrxParking []*parkingdomain.TrxParking, err error) {
	listTrxParking, err = uc.ParkingDB.GetList(ctx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrap(err, "Database.GetList")
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, commonerr.SetNewNotFound("data", "parking data not found")
	}

	return listTrxParking, err
}
