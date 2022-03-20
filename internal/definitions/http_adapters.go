package definitions

import (
	"github.com/sarulabs/di"
	"github.com/valyala/fasthttp"

	"companies/pkg/error_adapters/http_adapter"
)

const (
	CompanyHandlerHttpAdapter = "company_handler_http_adapter"
)

func GetCompanyHandlerHttpAdapterDef() di.Def {
	return di.Def{
		Name:  CompanyHandlerHttpAdapter,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return http_adapter.New(fasthttp.StatusInternalServerError,
				http_adapter.AdaptNotFoundError, http_adapter.AdaptBadRequestError), nil
		},
	}
}
