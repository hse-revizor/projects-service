package project

import (
	"errors"
)

var (
	ErrProjectExists      = errors.New("project account already exists")
	ErrProjectNotFound    = errors.New("project account not found")
	ErrForeignKeyError = errors.New("foreign key error")
	ErrInvalidField    = errors.New("error invalid field")
)
