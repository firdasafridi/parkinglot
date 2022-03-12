package parking

import "time"

type TrxParking struct {
	PlatNo           string    `json:"plat_no" gorm:"column:plat_no"`
	SlotNumber       int64     `json:"slot_number" gorm:"column:slot_number"`
	RegistrationDate time.Time `json:"registration_date" gorm:"column:reg_date"`
}

type HstParking struct {
	HstID            int64     `json:"history_id" gorm:"column:hst_id"`
	PlatNo           string    `json:"plat_no" gorm:"column:plat_no"`
	SlotNumber       int64     `json:"slot_number" gorm:"column:slot_number"`
	RegistrationDate time.Time `json:"registration_date" gorm:"column:reg_date"`
}
