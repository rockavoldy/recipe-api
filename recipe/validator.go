package recipe

import (
	"errors"
	"strings"
)

var (
	ErrNameMustAlphanumeric = errors.New("name must be alphanumeric")
	ErrNameTooShort         = errors.New("name must be above 3 character")
)

func validateName(name string) error {
	if len(name) < 3 {
		return ErrNameTooShort
	}
	if ok := strings.ContainsAny(name, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ ,.-_()[]{};:'\""); !ok {
		return ErrNameMustAlphanumeric
	}

	return nil
}
