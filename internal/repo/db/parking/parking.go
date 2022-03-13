package parking

import (
	"context"

	parkingdomain "github.com/firdasafridi/parkinglot/internal/entity/parking"
	"github.com/firdasafridi/parkinglot/lib/database"
)

type ParkingLotDB interface {
	GetList(ctx context.Context) (listTrxParking []*parkingdomain.TrxParking, err error)
	GetParkingLotByPlatNumber(ctx context.Context, platNo string) (parkingdomain.MapParking, error)
	GetEmptyParkingLot(ctx context.Context) (parkingdomain.MapParking, error)
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

func (db *ParkingDB) GetParkingLotByPlatNumber(ctx context.Context, platNo string) (data parkingdomain.MapParking, err error) {
	tx := db.Conn.DB.Table(TblMapParkingLot).Where("plat_no = ?", platNo).First(&data)
	
	if tx.Error != nil {
		return data, tx.Error
	}

	return
}

func (db *ParkingDB) GetEmptyParkingLot(ctx context.Context) (data parkingdomain.MapParking, err error) {
	tx := db.Conn.DB.Table(TblMapParkingLot).Where("plat_no = ''").First(&data)
	
	if tx.Error != nil {
		return data, tx.Error
	}

	return
}
