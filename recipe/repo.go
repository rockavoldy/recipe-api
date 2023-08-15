package recipe

import (
	"github.com/oklog/ulid/v2"
	"github.com/rockavoldy/recipe-api/recipematerial"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func listRecipes(tx *gorm.DB) ([]recipematerial.Recipe, error) {
	var recipes []recipematerial.Recipe
	res := tx.Preload("Materials.Material").Preload("Materials.Unit").Preload(clause.Associations).Find(&recipes)

	if err := res.Error; err != nil {
		return nil, err
	}

	return recipes, nil
}

func findRecipeById(tx *gorm.DB, id ulid.ULID) (recipematerial.Recipe, error) {
	var recipe recipematerial.Recipe
	res := tx.Preload("Materials.Material").Preload("Materials.Unit").First(&recipe, "id = ?", id)
	if err := res.Error; err != nil {
		return recipematerial.Recipe{}, err
	}

	return recipe, nil
}

func createOrUpdateRecipe(tx *gorm.DB, recipe recipematerial.Recipe) error {
	res := tx.Omit("Materials").Save(&recipe)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func appendMaterials(tx *gorm.DB, recipe recipematerial.Recipe, materialsJson []materialJsonReq) error {
	var recipeMaterialAppend []recipematerial.RecipeMaterial
	tx.Where("recipe_id = ?", recipe.ID).Delete(&recipematerial.RecipeMaterial{})

	for _, rmat := range materialsJson {
		recipeMaterialAppend = append(recipeMaterialAppend, recipematerial.RecipeMaterial{
			RecipeID:   recipe.ID,
			MaterialID: rmat.MaterialID,
			Quantity:   rmat.Quantity,
			UnitID:     rmat.UnitID,
		})
	}

	res := tx.Create(&recipeMaterialAppend)
	if err := res.Error; err != nil {
		return err
	}

	return nil
}

func deleteRecipe(tx *gorm.DB, recipe recipematerial.Recipe) error {
	err := tx.Transaction(func(txx *gorm.DB) error {
		var recipeMaterials []recipematerial.RecipeMaterial
		res := tx.Delete(&recipeMaterials, "recipe_id = ?", recipe.ID)
		if err := res.Error; err != nil {
			return err
		}

		res = tx.Delete(&recipe, "id = ?", recipe.ID)

		if err := res.Error; err != nil {
			return err
		}

		return nil
	})

	return err
}
