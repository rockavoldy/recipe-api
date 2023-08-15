package recipe

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/rockavoldy/recipe-api/category"
	"github.com/rockavoldy/recipe-api/common"
	"github.com/rockavoldy/recipe-api/material"
	"github.com/rockavoldy/recipe-api/unit"
	"gopkg.in/guregu/null.v4"
)

type Recipe struct {
	ID         ulid.ULID
	Name       string
	CategoryID ulid.ULID
	Category   category.Category
	Materials  []material.Material `gorm:"many2many:recipe_materials;"`
	CreatedAt  time.Time
	UpdatedAt  null.Time
}

// will be use by gorm
type RecipeMaterial struct {
	RecipeID   ulid.ULID `gorm:"primaryKey"`
	MaterialID ulid.ULID `gorm:"primaryKey"`
	Quantity   int
	UnitID     ulid.ULID `gorm:"primaryKey"`
	Unit       unit.Unit
}

func NewRecipe(name string) (Recipe, error) {
	if err := validateName(name); err != nil {
		return Recipe{}, err
	}

	recipe := Recipe{
		ID:        ulid.Make(),
		Name:      name,
		CreatedAt: time.Now(),
	}

	return recipe, nil
}

type materialJsonReq struct {
	MaterialID ulid.ULID `json:"material_id"`
	Quantity   int       `json:"quantity"`
	UnitID     ulid.ULID `json:"unit_id"`
}

type recipeJsonReq struct {
	Name       string            `json:"name"`
	CategoryID ulid.ULID         `json:"category_id"`
	Materials  []materialJsonReq `json:"materials"`
}

func (r Recipe) MarshalJSON() ([]byte, error) {
	var j struct {
		ID         ulid.ULID           `json:"id"`
		Name       string              `json:"name"`
		CategoryID ulid.ULID           `json:"category_id"`
		Materials  []material.Material `json:"materials"`
		CreatedAt  time.Time           `json:"created_at"`
		UpdatedAt  *time.Time          `json:"updated_at,omitempty"`
	}

	j.ID = r.ID
	j.Name = r.Name
	j.CategoryID = r.CategoryID
	j.Materials = r.Materials
	j.CreatedAt = r.CreatedAt
	j.UpdatedAt = r.UpdatedAt.Ptr()

	return json.Marshal(j)
}

func (r *Recipe) UnmarshalJSON(data []byte) error {
	var j struct {
		ID         ulid.ULID           `json:"id"`
		Name       string              `json:"name"`
		CategoryID ulid.ULID           `json:"category_id"`
		Materials  []material.Material `json:"materials"`
		CreatedAt  string              `json:"created_at,omitempty"`
		UpdatedAt  null.String         `json:"updated_at,omitempty"`
	}

	err := json.Unmarshal(data, &j)
	if err != nil {
		log.Println("unmarshal recipe")
		return errors.New("unmarshal recipe")
	}

	createdAt, err := time.Parse(time.RFC3339, j.CreatedAt)
	if err != nil {
		return err
	}
	updatedAt := common.ParseNullStringToTime(j.UpdatedAt)
	r = &Recipe{
		ID:         j.ID,
		Name:       j.Name,
		CategoryID: j.CategoryID,
		Materials:  j.Materials,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
	}

	return nil
}
