package http_adapter

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"

	"companies/internal/repositories/company"
)

var (
	expectedMap = map[error]int{
		company.ErrNotFound:     fasthttp.StatusNotFound,
		company.ErrAlreadyExist: fasthttp.StatusBadRequest,
	}
)

func Test_Adapter(t *testing.T) {
	defaultCode := fasthttp.StatusInternalServerError
	adapter := New(defaultCode, AdaptNotFoundError, AdaptBadRequestError)

	assert.Equal(t, expectedMap[company.ErrNotFound], adapter.AdaptToHttpCode(company.ErrNotFound))
	assert.Equal(t, expectedMap[company.ErrAlreadyExist], adapter.AdaptToHttpCode(company.ErrAlreadyExist))
	assert.Equal(t, defaultCode, adapter.AdaptToHttpCode(errors.New("")))
}
