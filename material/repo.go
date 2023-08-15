package material

import (
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

func listMaterials(tx *gorm.DB) ([]Material, error) {
	var materials []Material
	res := tx.Find(&materials)

	if err := res.Error; err != nil {
		return nil, err
	}

	return materials, nil
}

func findMaterialById(tx *gorm.DB, id ulid.ULID) (Material, error) {
	var material Material
	res := tx.First(&material, "id = ?", id)

	if err := res.Error; err != nil {
		return Material{}, err
	}

	return material, nil
}

func createOrUpdateMaterial(tx *gorm.DB, material Material) error {
	res := tx.Save(&material)

	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func deleteMaterial(tx *gorm.DB, material Material) error {
	res := tx.Delete(&material, "id = ?", material.ID)

	if err := res.Error; err != nil {
		return err
	}

	return nil
}
