package definitions

import (
	"github.com/sarulabs/di"
	"go.uber.org/zap"

	"companies/internal/configs"
	"companies/internal/handlers"
	"companies/internal/handlers/company"
	"companies/internal/router"
)

const (
	RouterDef = "router"
)

func GetRouterDef() di.Def {
	return di.Def{
		Name:  RouterDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get(CfgDef).(configs.Config)
			logger := ctn.Get(LoggerDef).(*zap.Logger)
			companyHandler := ctn.Get(CompanyHandlerDef).(company.IHandler)
			internalHandler := ctn.Get(InternalHandlerDef).(handlers.IHandler)

			return router.New(router.Config{
				InternalHandler: internalHandler,
				CompanyHandler:  companyHandler,
				App:             cfg.App,
				Logger:          logger,
			}), nil
		},
	}
}
