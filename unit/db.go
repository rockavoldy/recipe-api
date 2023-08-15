package unit

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db           *gorm.DB
	ErrMissingDB = errors.New("cannot assign nil db")
)

func SetDB(gormDb *gorm.DB) error {
	if gormDb == nil {
		return ErrMissingDB
	}

	db = gormDb
	db.AutoMigrate(&Unit{})
	return nil
}
