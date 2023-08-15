package recipe

import (
	"context"
	"errors"

	"github.com/oklog/ulid/v2"
	"github.com/rockavoldy/recipe-api/category"
	"gorm.io/gorm"
)

var (
	ErrNotFound       = errors.New("recipe not found")
	ErrAlreadyDeleted = errors.New("recipe have been deleted")
)

func List(ctx context.Context) ([]Recipe, error) {
	tx := db.WithContext(ctx)
	materials, err := listRecipes(tx)
	if err != nil {
		return nil, err
	}

	return materials, nil
}

func Create(ctx context.Context, data recipeJsonReq) (ulid.ULID, error) {
	recipe, err := NewRecipe(data.Name)
	if err != nil {
		return ulid.ULID{}, err
	}

	// filling recipe's category
	category, err := category.Find(ctx, data.CategoryID)
	if err != nil {
		return ulid.ULID{}, err
	}
	recipe.CategoryID = data.CategoryID
	recipe.Category = category

	tx := db.WithContext(ctx)
	err = tx.Transaction(func(txx *gorm.DB) error {
		// Put it inside the transaction, make sure to rollback when it fails to append materials
		if err := createOrUpdateRecipe(txx, recipe); err != nil {
			return err
		}

		if err := appendMaterials(txx, recipe, data.Materials); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return ulid.ULID{}, err
	}

	return recipe.ID, nil
}

func Find(ctx context.Context, id ulid.ULID) (Recipe, error) {
	tx := db.WithContext(ctx)
	recipe, err := findRecipeById(tx, id)
	if err != nil {
		return Recipe{}, ErrNotFound
	}

	return recipe, nil
}

func Update(ctx context.Context, id ulid.ULID, name string) (Recipe, error) {
	recipe, err := Find(ctx, id)
	if err != nil {
		return Recipe{}, err
	}

	tx := db.WithContext(ctx)
	recipe.Name = name
	if err := createOrUpdateRecipe(tx, recipe); err != nil {
		return Recipe{}, err
	}

	return recipe, nil
}

func Delete(ctx context.Context, id ulid.ULID) error {
	recipe, err := Find(ctx, id)
	if err != nil {
		return err
	}

	tx := db.WithContext(ctx)
	err = deleteRecipe(tx, recipe)
	if err != nil {
		return err
	}

	return nil
}
