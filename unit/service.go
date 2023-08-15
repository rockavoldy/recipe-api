package unit

import (
	"context"
	"errors"

	"github.com/oklog/ulid/v2"
)

var (
	ErrNotFound       = errors.New("unit not found")
	ErrAlreadyDeleted = errors.New("unit have been deleted")
)

func List(ctx context.Context) ([]Unit, error) {
	tx := db.WithContext(ctx)
	units, err := listCategories(tx)
	if err != nil {
		return nil, err
	}

	return units, nil
}

func Create(ctx context.Context, name string) (ulid.ULID, error) {
	unit, err := NewUnit(name)
	if err != nil {
		return ulid.ULID{}, nil
	}

	tx := db.WithContext(ctx)
	if err := createOrUpdateCategory(tx, unit); err != nil {
		return ulid.ULID{}, err
	}

	return unit.ID, nil
}

func Find(ctx context.Context, id ulid.ULID) (Unit, error) {
	tx := db.WithContext(ctx)
	unit, err := findCategoryById(tx, id)
	if err != nil {
		return Unit{}, ErrNotFound
	}

	return unit, nil
}

func Update(ctx context.Context, id ulid.ULID, name string) (Unit, error) {
	unit, err := Find(ctx, id)
	if err != nil {
		return Unit{}, err
	}

	tx := db.WithContext(ctx)
	unit.Name = name
	if err := createOrUpdateCategory(tx, unit); err != nil {
		return Unit{}, err
	}

	return unit, nil
}

func Delete(ctx context.Context, id ulid.ULID) error {
	unit, err := Find(ctx, id)
	if err != nil {
		return err
	}

	tx := db.WithContext(ctx)
	err = deleteCategory(tx, unit)
	if err != nil {
		return err
	}

	return nil
}
