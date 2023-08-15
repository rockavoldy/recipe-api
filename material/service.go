package material

import (
	"context"
	"errors"

	"github.com/oklog/ulid/v2"
	"github.com/rockavoldy/recipe-api/recipematerial"
)

var (
	ErrNotFound       = errors.New("material not found")
	ErrAlreadyDeleted = errors.New("material have been deleted")
)

func List(ctx context.Context) ([]recipematerial.Material, error) {
	tx := db.WithContext(ctx)
	materials, err := listMaterials(tx)
	if err != nil {
		return nil, err
	}

	return materials, nil
}

func Create(ctx context.Context, name string) (ulid.ULID, error) {
	material, err := recipematerial.NewMaterial(name)
	if err != nil {
		return ulid.ULID{}, nil
	}

	tx := db.WithContext(ctx)
	if err := createOrUpdateMaterial(tx, material); err != nil {
		return ulid.ULID{}, err
	}

	return material.ID, nil
}

func Find(ctx context.Context, id ulid.ULID) (recipematerial.Material, error) {
	tx := db.WithContext(ctx)
	material, err := findMaterialById(tx, id)
	if err != nil {
		return recipematerial.Material{}, ErrNotFound
	}

	return material, nil
}

func Update(ctx context.Context, id ulid.ULID, name string) (recipematerial.Material, error) {
	material, err := Find(ctx, id)
	if err != nil {
		return recipematerial.Material{}, err
	}

	tx := db.WithContext(ctx)
	material.Name = name
	if err := createOrUpdateMaterial(tx, material); err != nil {
		return recipematerial.Material{}, err
	}

	return material, nil
}

func Delete(ctx context.Context, id ulid.ULID) error {
	material, err := Find(ctx, id)
	if err != nil {
		return err
	}

	tx := db.WithContext(ctx)
	err = deleteMaterial(tx, material)
	if err != nil {
		return err
	}

	return nil
}
