package parking

import (
	"context"

	parkingdomain "github.com/firdasafridi/parkinglot/internal/entity/parking"
	"github.com/firdasafridi/parkinglot/lib/database"
)

type ParkingLotDB interface {
	GetList(ctx context.Context) (listTrxParking []*parkingdomain.TrxParking, err error)
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
