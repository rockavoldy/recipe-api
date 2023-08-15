package recipe

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
	err := db.SetupJoinTable(&Recipe{}, "Materials", &RecipeMaterial{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&Recipe{})
	return nil
}
