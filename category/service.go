package category

import (
	"context"
	"errors"

	"github.com/oklog/ulid/v2"
)

var (
	ErrNotFound       = errors.New("category not found")
	ErrAlreadyDeleted = errors.New("category have been deleted")
)

func List(ctx context.Context) ([]Category, error) {
	tx := db.WithContext(ctx)
	categories, err := listCategories(tx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func Create(ctx context.Context, name string) (ulid.ULID, error) {
	category, err := NewCategory(name)
	if err != nil {
		return ulid.ULID{}, nil
	}

	tx := db.WithContext(ctx)
	if err := createOrUpdateCategory(tx, category); err != nil {
		return ulid.ULID{}, err
	}

	return category.ID, nil
}

func Find(ctx context.Context, id ulid.ULID) (Category, error) {
	tx := db.WithContext(ctx)
	category, err := findCategoryById(tx, id)
	if err != nil {
		return Category{}, ErrNotFound
	}

	return category, nil
}

func Update(ctx context.Context, id ulid.ULID, name string) (Category, error) {
	category, err := Find(ctx, id)
	if err != nil {
		return Category{}, err
	}

	tx := db.WithContext(ctx)
	category.Name = name
	if err := createOrUpdateCategory(tx, category); err != nil {
		return Category{}, err
	}

	return category, nil
}

func Delete(ctx context.Context, id ulid.ULID) error {
	category, err := Find(ctx, id)
	if err != nil {
		return err
	}

	tx := db.WithContext(ctx)
	err = deleteCategory(tx, category)
	if err != nil {
		return err
	}

	return nil
}
