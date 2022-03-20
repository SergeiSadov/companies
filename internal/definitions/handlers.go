package definitions

import (
	"github.com/sarulabs/di"
	"go.uber.org/zap"

	"companies/internal/handlers"
	"companies/internal/handlers/company"
	"companies/internal/handlers/company/validators"
	company_use_case "companies/internal/usecases/company"
	"companies/pkg/error_adapters/http_adapter"
)

const (
	InternalHandlerDef = "internal_handler"
	CompanyHandlerDef  = "company_handler"
)

func GetInternalHandlerDef() di.Def {
	return di.Def{
		Name:  InternalHandlerDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			logger := ctn.Get(LoggerDef).(*zap.Logger)

			return handlers.NewHttpHandler(logger), nil
		},
	}
}

func GetCompanyHandlerDef() di.Def {
	return di.Def{
		Name:  CompanyHandlerDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			handler := ctn.Get(InternalHandlerDef).(handlers.IHandler)
			companyUseCase := ctn.Get(CompanyUseCase).(company_use_case.IUseCase)
			logger := ctn.Get(LoggerDef).(*zap.Logger)
			httpAdapter := ctn.Get(CompanyHandlerHttpAdapter).(http_adapter.IErrorAdapter)

			return company.NewHandler(company.Config{
				Internal:     handler,
				UseCase:      companyUseCase,
				Logger:       logger,
				ErrorAdapter: httpAdapter,
				Validator:    validators.New(validators.PreparedValidatorParams),
			}), nil
		},
	}
}
