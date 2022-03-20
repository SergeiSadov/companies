package definitions

import (
	"github.com/sarulabs/di"

	"companies/internal/configs"
	"companies/internal/db"
	"companies/internal/repositories/company"
	"companies/pkg/error_adapters/sql_adapter"
)

const (
	ConnectionDef  = "connection"
	CompanyRepoDef = "company_repo"
)

func GetConnectionDef() di.Def {
	return di.Def{
		Name:  ConnectionDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get(CfgDef).(configs.Config)

			conn, err := db.New(cfg.Database.Dialect, cfg.Database.PrepareDSN())
			if err != nil {
				return nil, err
			}

			return conn, nil
		},
	}
}

func GetCompanyRepoDef() di.Def {
	return di.Def{
		Name:  CompanyRepoDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			conn := ctn.Get(ConnectionDef).(db.IDatabase)
			gormDB, err := conn.GetConn()
			if err != nil {
				return nil, err
			}

			return company.NewRepository(gormDB, sql_adapter.New(company.PreparedErrorsMap)), nil
		},
	}
}
