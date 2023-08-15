package category

import (
	"encoding/json"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/rockavoldy/recipe-api/common"
	"gopkg.in/guregu/null.v4"
)

type Category struct {
	ID        ulid.ULID
	Name      string
	CreatedAt time.Time
	UpdatedAt null.Time
}

func NewCategory(name string) (Category, error) {
	if err := common.ValidateName(name); err != nil {
		return Category{}, err
	}

	category := Category{
		ID:        ulid.Make(),
		Name:      name,
		CreatedAt: time.Now(),
	}

	return category, nil
}

func (c Category) MarshalJSON() ([]byte, error) {
	var j struct {
		ID        ulid.ULID  `json:"id"`
		Name      string     `json:"name"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
	}

	j.ID = c.ID
	j.Name = c.Name
	j.CreatedAt = c.CreatedAt
	j.UpdatedAt = c.UpdatedAt.Ptr()

	return json.Marshal(j)
}

func (c *Category) UnmarshalJSON(data []byte) error {
	var j struct {
		ID        ulid.ULID   `json:"id"`
		Name      string      `json:"name"`
		CreatedAt string      `json:"created_at"`
		UpdatedAt null.String `json:"updated_at"`
	}

	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	createdAt, err := time.Parse(time.RFC3339, j.CreatedAt)
	if err != nil {
		return err
	}

	updatedAt := common.ParseNullStringToTime(j.UpdatedAt)

	c = &Category{
		ID:        j.ID,
		Name:      j.Name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	return nil
}
