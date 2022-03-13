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
	GetParkingLotByPlatNumber(ctx context.Context, platNo string) (parkinglotID int64, err error)
	GetEmptyParkingLot(ctx context.Context) (parkinglotID int64, err error)
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

func (uc *Parking) GetParkingLotByPlatNumber(ctx context.Context, platNo string) (parkingLotID int64, err error) {

	if platNo == "" {
		return 0, commonerr.SetNewBadRequest("Plat Number", "Plat Number is required")
	}

	parkingLot, err := uc.ParkingDB.GetParkingLotByPlatNumber(ctx, platNo)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, errors.Wrap(err, "Database.GetList")
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, commonerr.SetNewNotFound("data", "parking data not found")
	}

	return parkingLot.ID, nil
}

func (uc *Parking) GetEmptyParkingLot(ctx context.Context) (parkingLotID int64, err error) {
	parkingLot, err := uc.ParkingDB.GetEmptyParkingLot(ctx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, errors.Wrap(err, "Database.GetList")
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, commonerr.SetNewNotFound("data", "parking data not found")
	}

	return parkingLot.ID, nil
}
