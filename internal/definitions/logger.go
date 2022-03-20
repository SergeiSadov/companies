package definitions

import (
	"github.com/sarulabs/di"
	"go.uber.org/zap"
)

const (
	LoggerDef = "logger"
)

func GetLoggerDef() di.Def {
	return di.Def{
		Name:  LoggerDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return zap.NewExample(), nil
		},
	}
}
