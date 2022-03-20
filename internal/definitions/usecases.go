package definitions

import (
	"github.com/sarulabs/di"

	company_repo "companies/internal/repositories/company"
	"companies/internal/usecases/company"
)

const (
	CompanyUseCase = "company_use_case"
)

func GetCompanyUseCaseDef() di.Def {
	return di.Def{
		Name:  CompanyUseCase,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			repo := ctn.Get(CompanyRepoDef).(company_repo.IRepository)

			return company.New(repo, company.NewAdapter()), nil
		},
	}
}
