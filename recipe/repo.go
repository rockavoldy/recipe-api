package recipe

import (
	"log"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

func listRecipes(tx *gorm.DB) ([]Recipe, error) {
	var recipes []Recipe
	res := tx.Preload("Materials").Find(&recipes)

	if err := res.Error; err != nil {
		return nil, err
	}

	return recipes, nil
}

func findRecipeById(tx *gorm.DB, id ulid.ULID) (Recipe, error) {
	var recipe Recipe
	res := tx.Preload("Materials").First(&recipe, "id = ?", id)

	if err := res.Error; err != nil {
		return Recipe{}, err
	}

	return recipe, nil
}

func createOrUpdateRecipe(tx *gorm.DB, recipe Recipe) error {
	res := tx.Omit("Materials").Save(&recipe)
	if err := res.Error; err != nil {
		log.Println("repo recipe")
		return err
	}

	return nil
}

func appendMaterials(tx *gorm.DB, recipe Recipe, materialsJson []materialJsonReq) error {
	var recipeMaterialAppend []RecipeMaterial
	tx.Where("recipe_id = ?", recipe.ID).Delete(&RecipeMaterial{})

	for _, rmat := range materialsJson {
		recipeMaterialAppend = append(recipeMaterialAppend, RecipeMaterial{
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

func deleteRecipe(tx *gorm.DB, recipe Recipe) error {
	res := tx.Preload("Materials").Delete(&recipe, "id = ?", recipe.ID)

	if err := res.Error; err != nil {
		return err
	}

	return nil
}
