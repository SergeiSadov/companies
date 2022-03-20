package errors

import "errors"

var (
	ErrEmptyParamID = errors.New("empty id param")
	ErrInvalidJSON  = errors.New("invalid json")
)
