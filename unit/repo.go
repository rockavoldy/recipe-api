package unit

import (
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

func listCategories(tx *gorm.DB) ([]Unit, error) {
	var units []Unit
	res := tx.Find(&units)

	if err := res.Error; err != nil {
		return nil, err
	}

	return units, nil
}

func findCategoryById(tx *gorm.DB, id ulid.ULID) (Unit, error) {
	var unit Unit
	res := tx.First(&unit, "id = ?", id)

	if err := res.Error; err != nil {
		return Unit{}, err
	}

	return unit, nil
}

func createOrUpdateCategory(tx *gorm.DB, unit Unit) error {
	res := tx.Save(&unit)

	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func deleteCategory(tx *gorm.DB, unit Unit) error {
	res := tx.Delete(&unit, "id = ?", unit.ID)

	if err := res.Error; err != nil {
		return err
	}

	return nil
}
