package unit

import (
	"encoding/json"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/rockavoldy/recipe-api/common"
	"gopkg.in/guregu/null.v4"
)

type Unit struct {
	ID        ulid.ULID
	Name      string
	CreatedAt time.Time
	UpdatedAt null.Time
}

func NewUnit(name string) (Unit, error) {
	if err := validateName(name); err != nil {
		return Unit{}, err
	}

	unit := Unit{
		ID:        ulid.Make(),
		Name:      name,
		CreatedAt: time.Now(),
	}

	return unit, nil
}

func (u Unit) MarshalJSON() ([]byte, error) {
	var j struct {
		ID        ulid.ULID  `json:"id"`
		Name      string     `json:"name"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
	}

	j.ID = u.ID
	j.Name = u.Name
	j.CreatedAt = u.CreatedAt
	j.UpdatedAt = u.UpdatedAt.Ptr()

	return json.Marshal(j)
}

func (u *Unit) UnmarshalJSON(data []byte) error {
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

	u = &Unit{
		ID:        j.ID,
		Name:      j.Name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	return nil
}
