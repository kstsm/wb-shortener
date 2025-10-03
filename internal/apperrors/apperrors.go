package apperrors

import "errors"

var (
	ErrAliasAlreadyExists = errors.New("alias already exists")
	ErrNotFound           = errors.New("not found")
)
