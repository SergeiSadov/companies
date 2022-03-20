package http_adapter

import (
	"github.com/valyala/fasthttp"

	"companies/internal/repositories/company"
	"companies/pkg/errors"
)

type ErrorToHttpCodeAdapter func(err error) (code int)

func AdaptNotFoundError(err error) (code int) {
	if err == company.ErrNotFound {
		return fasthttp.StatusNotFound
	}

	return code
}

func AdaptBadRequestError(err error) (code int) {
	switch err {
	case company.ErrAlreadyExist,
		errors.ErrEmptyParamID,
		errors.ErrInvalidJSON,
		errors.ErrInvalidCompanyName,
		errors.ErrInvalidIndustryName,
		errors.ErrInvalidCo2Footprint,
		errors.ErrInvalidUUID:
		return fasthttp.StatusBadRequest
	default:
		return code
	}
}
