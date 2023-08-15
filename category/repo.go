package category

import (
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

func listCategories(tx *gorm.DB) ([]Category, error) {
	var categories []Category
	res := tx.Find(&categories)

	if err := res.Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func findCategoryById(tx *gorm.DB, id ulid.ULID) (Category, error) {
	var category Category
	res := tx.First(&category, "id = ?", id)

	if err := res.Error; err != nil {
		return Category{}, err
	}

	return category, nil
}

func createOrUpdateCategory(tx *gorm.DB, category Category) error {
	res := tx.Save(&category)

	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func deleteCategory(tx *gorm.DB, category Category) error {
	res := tx.Delete(&category, "id = ?", category.ID)

	if err := res.Error; err != nil {
		return err
	}

	return nil
}
