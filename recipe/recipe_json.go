package recipe

import "github.com/oklog/ulid/v2"

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
