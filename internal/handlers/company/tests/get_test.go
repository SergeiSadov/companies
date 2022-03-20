package tests

import (
	"encoding/json"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"

	"companies/internal/models/api"
	"companies/internal/repositories/company"
	mock_company "companies/internal/repositories/company/mock"
)

func Test_GetSuccess(t *testing.T) {
	companyRepoMock := mock_company.NewMockIRepository(gomock.NewController(t))
	companyRepoMock.EXPECT().Get(gomock.Any(), &GetCompanyRepoReq).Times(1).Return(&GetCompanyRepoResp, nil)

	handler := prepareHandler(&params{
		companyRepo: companyRepoMock,
	})

	ctx, err := prepareContext(&contextParams{
		userValues: GetRequestOk,
	})
	assert.NoError(t, err)

	handler.Get(ctx)
	resp, err := extractGetResponse(ctx)
	assert.NoError(t, err)
	assert.Equal(t, GetCompanyExpectedResponseOk, resp)
}

func Test_GetNotFound(t *testing.T) {
	companyRepoMock := mock_company.NewMockIRepository(gomock.NewController(t))
	companyRepoMock.EXPECT().Get(gomock.Any(), &GetCompanyRepoReq).Times(1).Return(nil, company.ErrNotFound)

	handler := prepareHandler(&params{
		companyRepo: companyRepoMock,
	})

	ctx, err := prepareContext(&contextParams{
		userValues: GetRequestOk,
	})
	assert.NoError(t, err)

	handler.Get(ctx)
	assert.NoError(t, err)
	assert.Equal(t, fasthttp.StatusNotFound, ctx.Response.StatusCode())
}

func Test_GetBadRequest(t *testing.T) {
	companyRepoMock := mock_company.NewMockIRepository(gomock.NewController(t))

	handler := prepareHandler(&params{
		companyRepo: companyRepoMock,
	})

	ctx, err := prepareContext(&contextParams{})
	assert.NoError(t, err)

	handler.Get(ctx)
	assert.NoError(t, err)
	assert.Equal(t, fasthttp.StatusBadRequest, ctx.Response.StatusCode())
}

func extractGetResponse(ctx *fasthttp.RequestCtx) (res *api.GetCompanyResponse, err error) {
	return res, json.Unmarshal(ctx.Response.Body(), &res)
}
