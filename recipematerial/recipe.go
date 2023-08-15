package recipematerial

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/rockavoldy/recipe-api/category"
	"github.com/rockavoldy/recipe-api/common"
	"gopkg.in/guregu/null.v4"
)

type Recipe struct {
	ID         ulid.ULID
	Name       string
	CategoryID ulid.ULID
	Category   category.Category
	Materials  []RecipeMaterial
	CreatedAt  time.Time
	UpdatedAt  null.Time
}

func NewRecipe(name string) (Recipe, error) {
	if err := common.ValidateName(name); err != nil {
		return Recipe{}, err
	}

	recipe := Recipe{
		ID:        ulid.Make(),
		Name:      name,
		CreatedAt: time.Now(),
	}

	return recipe, nil
}

func (r Recipe) MarshalJSON() ([]byte, error) {
	var j struct {
		ID         ulid.ULID        `json:"id"`
		Name       string           `json:"name"`
		Materials  []RecipeMaterial `json:"materials"`
		CategoryID ulid.ULID        `json:"category_id"`
		CreatedAt  time.Time        `json:"created_at"`
		UpdatedAt  *time.Time       `json:"updated_at,omitempty"`
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
		ID         ulid.ULID        `json:"id"`
		Name       string           `json:"name"`
		Materials  []RecipeMaterial `json:"materials"`
		CategoryID ulid.ULID        `json:"category_id"`
		CreatedAt  string           `json:"created_at"`
		UpdatedAt  null.String      `json:"updated_at"`
	}

	err := json.Unmarshal(data, &j)
	if err != nil {
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
