package recipematerial

import (
	"encoding/json"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/rockavoldy/recipe-api/common"
	"gopkg.in/guregu/null.v4"
)

type Material struct {
	ID        ulid.ULID
	Name      string
	Recipes   []RecipeMaterial
	CreatedAt time.Time
	UpdatedAt null.Time
}

func NewMaterial(name string) (Material, error) {
	if err := common.ValidateName(name); err != nil {
		return Material{}, err
	}

	material := Material{
		ID:        ulid.Make(),
		Name:      name,
		CreatedAt: time.Now(),
	}

	return material, nil
}

func (m Material) MarshalJSON() ([]byte, error) {
	var j struct {
		ID        ulid.ULID  `json:"id"`
		Name      string     `json:"name"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
	}

	j.ID = m.ID
	j.Name = m.Name
	j.CreatedAt = m.CreatedAt
	j.UpdatedAt = m.UpdatedAt.Ptr()

	return json.Marshal(j)
}

func (m *Material) UnmarshalJSON(data []byte) error {
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

	m = &Material{
		ID:        j.ID,
		Name:      j.Name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	return nil
}
