package material

import (
	"encoding/json"
	"log"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/rockavoldy/recipe-api/common"
	"github.com/rockavoldy/recipe-api/unit"
	"gopkg.in/guregu/null.v4"
)

type Material struct {
	ID        ulid.ULID
	Name      string
	CreatedAt time.Time
	UpdatedAt null.Time
	Quantity  int       `gorm:"-:migration;->"`
	UnitID    ulid.ULID `gorm:"-:migration;->"`
	Unit      unit.Unit `gorm:"-:migration;->"`
}

func NewMaterial(name string) (Material, error) {
	if err := validateName(name); err != nil {
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
		CreatedAt time.Time  `json:"created_at,omitempty"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
		Quantity  int        `json:"quantity,omitempty"`
		Unit      string     `json:"unit,omitempty"`
	}

	j.ID = m.ID
	j.Name = m.Name
	j.CreatedAt = m.CreatedAt
	j.UpdatedAt = m.UpdatedAt.Ptr()
	j.Quantity = m.Quantity
	j.Unit = m.Unit.Name

	return json.Marshal(j)
}

func (m *Material) UnmarshalJSON(data []byte) error {
	var j struct {
		ID        ulid.ULID   `json:"id"`
		Name      string      `json:"name"`
		CreatedAt string      `json:"created_at,omitempty"`
		UpdatedAt null.String `json:"updated_at,omitempty"`
		Quantity  int         `json:"quantity,omitempty"`
		UnitID    ulid.ULID   `json:"unit_id"`
	}

	err := json.Unmarshal(data, &j)
	if err != nil {
		log.Println("unmarshal material")
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
		Quantity:  j.Quantity,
		UnitID:    j.UnitID,
	}

	return nil
}
