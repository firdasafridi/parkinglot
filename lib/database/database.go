package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/firdasafridi/parkinglot/lib/common/log"
)

type Config struct {
	IsDebug bool
}

type Connection struct {
	DB *gorm.DB
}

func New(dbGroup, dsn string, cfg Config) (*Connection, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Infoln(dbGroup, "Success connection to database")

	if cfg.IsDebug {
		db = db.Debug()
	}

	return &Connection{
		DB: db,
	}, nil
}
