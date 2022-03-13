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
	GetParkingHistoryByDate(ctx context.Context, date parkingdomain.ParkingDate) ([]*parkingdomain.HstParking, error)
	GetParkingHistoryDailyReport(ctx context.Context) (*parkingdomain.ParkingReport, error)
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
