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
	ParkVehicle(ctx context.Context, platNo string) error
	LeaveParkingLot(ctx context.Context, platNo string) error
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

func (uc *Parking) ParkVehicle(ctx context.Context, platNo string) error {
	// get empty parking lot
	parkingSLotNo, err := uc.ParkingDB.GetEmptyParkingLot(ctx)
	if err != nil {
		return errors.Wrap(err, "Database.GetEmptyParkingLot")
	}

	err = uc.ParkingDB.ParkVehicle(ctx, &parkingdomain.TrxParking{
		PlatNo:     platNo,
		SlotNumber: parkingSLotNo,
	})
	if err != nil {
		return errors.Wrap(err, "Database.ParkVehicle")
	}

	return nil
}

func (uc *Parking) LeaveParkingLot(ctx context.Context, platNo string) error {
	err := uc.ParkingDB.LeaveParkingLot(ctx, platNo)
	if err != nil {
		return errors.Wrap(err, "Database.LeaveParkingLot")
	}

	return nil
}
