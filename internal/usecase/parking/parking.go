package parking

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	parkingdomain "github.com/firdasafridi/parkinglot/internal/entity/parking"
	parkingdb "github.com/firdasafridi/parkinglot/internal/repo/db/parking"
	"github.com/firdasafridi/parkinglot/lib/common"
	"github.com/firdasafridi/parkinglot/lib/common/commonerr"
)

type ParkingUC interface {
	GetAllParkingData(ctx context.Context) (listTrxParking []*parkingdomain.TrxParking, err error)
	ParkVehicle(ctx context.Context, platNo string) error
	LeaveParkingLot(ctx context.Context, platNo string) error
	GetParkingLotByPlatNumber(ctx context.Context, platNo string) (parkinglotID int64, err error)
	GetEmptyParkingLot(ctx context.Context) (parkinglotID int64, err error)
	GetParkingHistoryByDate(ctx context.Context, date parkingdomain.ParkingDate) ([]*parkingdomain.HstParking, error)
	GetParkingHistoryDailyReport(ctx context.Context) (*parkingdomain.ParkingReport, error)

	countriesUC
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

	return listTrxParking, nil
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

func (uc *Parking) GetParkingHistoryByDate(ctx context.Context, date parkingdomain.ParkingDate) ([]*parkingdomain.HstParking, error) {
	startDate, err := time.Parse(common.YYYYMMDDDash, date.StartDate)
	if err != nil {
		return nil, commonerr.SetNewBadRequest("start date", "invalid start date. use YYYY-MM-DD format")
	}
	endDate, err := time.Parse(common.YYYYMMDDDash, date.EndDate)
	if err != nil {
		return nil, commonerr.SetNewBadRequest("end date", "invalid end date. use YYYY-MM-DD format")
	}
	hstParking, err := uc.ParkingDB.GetParkingHistoryByDate(ctx, startDate, endDate)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrap(err, "Database.GetList")
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, commonerr.SetNewNotFound("data", "parking data not found")
	}
	return hstParking, nil

}

func (uc *Parking) GetParkingHistoryDailyReport(ctx context.Context) (*parkingdomain.ParkingReport, error) {
	result, err := uc.ParkingDB.GetParkingHistoryDailyReport(ctx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrap(err, "Database.GetList")
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, commonerr.SetNewNotFound("data", "parking data not found")
	}
	return result, nil
}

func (uc *Parking) ParkVehicle(ctx context.Context, platNo string) error {
	// get empty parking lot
	parkingSLotNo, err := uc.ParkingDB.GetEmptyParkingLot(ctx)
	if err != nil {
		return errors.Wrap(err, "Database.GetEmptyParkingLot")
	}

	err = uc.ParkingDB.ParkVehicle(ctx, &parkingdomain.TrxParking{
		PlatNo:     platNo,
		SlotNumber: parkingSLotNo.ID,
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
