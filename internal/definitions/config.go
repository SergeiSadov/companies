package definitions

import (
	"github.com/sarulabs/di"

	"companies/internal/configs"
)

const (
	CfgDef = "config"
)

func GetConfigDef() di.Def {
	return di.Def{
		Name:  CfgDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return configs.Setup()
		},
	}
}
