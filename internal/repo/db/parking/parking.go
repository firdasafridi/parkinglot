package parking

import (
	"context"
	"time"

	parkingdomain "github.com/firdasafridi/parkinglot/internal/entity/parking"
	"github.com/firdasafridi/parkinglot/lib/database"
)

type ParkingLotDB interface {
	GetList(ctx context.Context) (listTrxParking []*parkingdomain.TrxParking, err error)
	GetParkingHistoryByDate(ctx context.Context, startDate, endDate time.Time) ([]*parkingdomain.HstParking, error)
	GetParkingHistoryDailyReport(ctx context.Context) (*parkingdomain.ParkingReport, error)
}

type ParkingDB struct {
	Conn *database.Connection
}

// New create new repository parking db
func New(parkingDB *ParkingDB) *ParkingDB {
	return parkingDB
}

func (db *ParkingDB) GetList(ctx context.Context) (listTrxParking []*parkingdomain.TrxParking, err error) {
	listTrxParking = make([]*parkingdomain.TrxParking, 0)

	tx := db.Conn.DB.Table(TblTrxParking).Find(&listTrxParking)

	if tx.Error != nil {
		return nil, err
	}

	return listTrxParking, nil
}

func (db *ParkingDB) GetParkingHistoryByDate(ctx context.Context, startDate, endDate time.Time) ([]*parkingdomain.HstParking, error) {
	hstParking := []*parkingdomain.HstParking{}
	result := db.Conn.DB.Table(TblHstParking).
		Where("reg_date > ?", startDate).
		Where("reg_date < ?", endDate).
		Find(&hstParking)

	if result.Error != nil {
		return nil, result.Error
	}
	return hstParking, nil
}

func (db *ParkingDB) GetParkingHistoryDailyReport(ctx context.Context) (*parkingdomain.ParkingReport, error) {
	dailyReport := []*parkingdomain.DailyReport{}
	result := db.Conn.DB.Table(TblHstParking).
		Select("count(hst_id) as total_vehicle, DATE_FORMAT(reg_date, '%Y-%m-%d') as date").
		Group("date").
		Order("date").
		Find(&dailyReport)

	if result.Error != nil {
		return nil, result.Error
	}
	reportResult := &parkingdomain.ParkingReport{
		TotalDays: len(dailyReport),
		Reports:   dailyReport,
	}
	return reportResult, nil
}
