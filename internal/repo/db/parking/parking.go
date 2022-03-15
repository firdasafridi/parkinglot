package parking

import (
	"context"

	parkingdomain "github.com/firdasafridi/parkinglot/internal/entity/parking"
	"github.com/firdasafridi/parkinglot/lib/database"
)

type ParkingLotDB interface {
	GetList(ctx context.Context) (listTrxParking []*parkingdomain.TrxParking, err error)
	ParkVehicle(ctx context.Context, trxParking *parkingdomain.TrxParking) error
	LeaveParkingLot(ctx context.Context, platNo string) error
	GetEmptyParkingLot(ctx context.Context) (parkinglotID int64, err error)
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
		return nil, tx.Error
	}

	return listTrxParking, nil
}

func (db *ParkingDB) GetEmptyParkingLot(ctx context.Context) (parkingLotID int64, err error) {
	data := parkingdomain.MapParking{}
	tx := db.Conn.DB.Table(TblMapParkingLot).Where("plat_no = ''").First(&data)

	if tx.Error != nil {
		return data.ID, tx.Error
	}

	return data.ID, nil
}

func (db *ParkingDB) ParkVehicle(ctx context.Context, trxParking *parkingdomain.TrxParking) error {
	tx := db.Conn.DB.Table(TblTrxParking).Save(trxParking)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (db *ParkingDB) LeaveParkingLot(ctx context.Context, platNo string) error {
	tx := db.Conn.DB.Table(TblTrxParking).Delete(&parkingdomain.TrxParking{}, "plat_no = ?", platNo)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
