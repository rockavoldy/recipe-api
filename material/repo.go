package material

import (
	"github.com/oklog/ulid/v2"
	"github.com/rockavoldy/recipe-api/recipematerial"
	"gorm.io/gorm"
)

func listMaterials(tx *gorm.DB) ([]recipematerial.Material, error) {
	var materials []recipematerial.Material
	res := tx.Find(&materials)

	if err := res.Error; err != nil {
		return nil, err
	}

	return materials, nil
}

func findMaterialById(tx *gorm.DB, id ulid.ULID) (recipematerial.Material, error) {
	var material recipematerial.Material
	res := tx.First(&material, "id = ?", id)

	if err := res.Error; err != nil {
		return recipematerial.Material{}, err
	}

	return material, nil
}

func createOrUpdateMaterial(tx *gorm.DB, material recipematerial.Material) error {
	res := tx.Save(&material)

	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func deleteMaterial(tx *gorm.DB, material recipematerial.Material) error {
	res := tx.Delete(&material, "id = ?", material.ID)

	if err := res.Error; err != nil {
		return err
	}

	return nil
}
