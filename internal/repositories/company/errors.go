package company

import (
	"errors"
)

const (
	UniqueViolationErr = "23505"
	NotFoundErr        = "20000"
)

var (
	ErrAlreadyExist = errors.New("already exists")
	ErrNotFound     = errors.New("not found")
)

var PreparedErrorsMap = map[string]error{
	UniqueViolationErr: ErrAlreadyExist,
	NotFoundErr:        ErrNotFound,
}
