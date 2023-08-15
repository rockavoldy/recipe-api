package recipematerial

import (
	"encoding/json"

	"github.com/oklog/ulid/v2"
	"github.com/rockavoldy/recipe-api/unit"
)

type RecipeMaterial struct {
	RecipeID   ulid.ULID `gorm:"primaryKey"`
	Recipe     Recipe    `gorm:"foreignKey:RecipeID;references:ID"`
	MaterialID ulid.ULID `gorm:"primaryKey"`
	Material   Material  `gorm:"foreignKey:MaterialID;references:ID"`
	Quantity   int
	UnitID     ulid.ULID `gorm:"primaryKey"`
	Unit       unit.Unit
}

func (rm RecipeMaterial) MarshalJSON() ([]byte, error) {
	var j struct {
		RecipeID   ulid.ULID `json:"-"`
		MaterialID ulid.ULID `json:"-"`
		Material   string    `json:"material"`
		Quantity   int       `json:"quantity"`
		UnitID     ulid.ULID `json:"-"`
		Unit       string    `json:"unit"`
	}

	j.Material = rm.Material.Name
	j.Quantity = rm.Quantity
	j.Unit = rm.Unit.Name

	return json.Marshal(&j)
}

func (rm *RecipeMaterial) UnmarshalJSON(data []byte) error {
	var j struct {
		RecipeID   ulid.ULID `json:"recipe_id"`
		Recipe     Recipe    `json:"-"`
		MaterialID ulid.ULID `json:"material_id"`
		Material   Material  `json:"-"`
		Quantity   int       `json:"quantity"`
		UnitID     ulid.ULID `json:"unit_id"`
		Unit       unit.Unit `json:"-"`
	}

	err := json.Unmarshal(data, &j)
	if err != nil {
		return err
	}

	rm = &RecipeMaterial{
		RecipeID:   j.RecipeID,
		Recipe:     j.Recipe,
		MaterialID: j.MaterialID,
		Material:   j.Material,
		Quantity:   j.Quantity,
		UnitID:     j.UnitID,
		Unit:       j.Unit,
	}

	return nil
}
