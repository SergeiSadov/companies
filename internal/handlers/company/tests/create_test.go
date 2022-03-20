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

func Test_CreateSuccess(t *testing.T) {
	companyRepoMock := mock_company.NewMockIRepository(gomock.NewController(t))
	companyRepoMock.EXPECT().Create(gomock.Any(), &CreateCompanyRepoReq).Times(1).Return(&CreateCompanyRepoResp, nil)

	handler := prepareHandler(&params{
		companyRepo: companyRepoMock,
	})

	ctx, err := prepareContext(&contextParams{
		request: CreateCompanyRequestOk,
	})
	assert.NoError(t, err)

	handler.Create(ctx)
	resp, err := extractCreateResponse(ctx)
	assert.NoError(t, err)
	assert.Equal(t, CreateCompanyExpectedResponseOk, resp)
}

func Test_CreateBadRequest(t *testing.T) {
	companyRepoMock := mock_company.NewMockIRepository(gomock.NewController(t))
	companyRepoMock.EXPECT().Create(gomock.Any(), &CreateCompanyRepoReq).Times(1).Return(nil, company.ErrAlreadyExist)

	handler := prepareHandler(&params{
		companyRepo: companyRepoMock,
	})

	ctx, err := prepareContext(&contextParams{
		request: CreateCompanyRequestOk,
	})
	assert.NoError(t, err)

	handler.Create(ctx)
	assert.NoError(t, err)
	assert.Equal(t, fasthttp.StatusBadRequest, ctx.Response.StatusCode())
}

func extractCreateResponse(ctx *fasthttp.RequestCtx) (res *api.CreateCompanyResponse, err error) {
	return res, json.Unmarshal(ctx.Response.Body(), &res)
}
