package tests

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
	"go.uber.org/zap"

	"companies/internal/handlers"
	"companies/internal/handlers/company"
	"companies/internal/handlers/company/validators"
	company_repository "companies/internal/repositories/company"
	company_use_case "companies/internal/usecases/company"
	"companies/pkg/error_adapters/http_adapter"
)

type params struct {
	companyRepo company_repository.IRepository
}

type contextParams struct {
	request    interface{}
	userValues []userValue
}

type userValue struct {
	key   string
	value string
}

func prepareHandler(p *params) (handler company.IHandler) {
	logger := zap.L()
	return company.NewHandler(company.Config{
		Internal: handlers.NewHttpHandler(logger),
		UseCase:  company_use_case.New(p.companyRepo, company_use_case.NewAdapter()),
		Logger:   logger,
		ErrorAdapter: http_adapter.New(fasthttp.StatusInternalServerError,
			http_adapter.AdaptNotFoundError,
			http_adapter.AdaptBadRequestError,
		),
		Validator: validators.New(validators.PreparedValidatorParams),
	})
}

func prepareContext(p *contextParams) (ctx *fasthttp.RequestCtx, err error) {
	ctx = new(fasthttp.RequestCtx)

	for i := range p.userValues {
		ctx.SetUserValue(p.userValues[i].key, p.userValues[i].value)
	}

	if p.request != nil {
		body, err := json.Marshal(p.request)
		if err != nil {
			return nil, err
		}
		ctx.Request.SetBody(body)
	}

	return ctx, nil
}
