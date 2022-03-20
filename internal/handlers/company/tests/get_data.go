package tests

import (
	"companies/internal/models/api"
	"companies/internal/models/repository"
)

var (
	GetRequestOk = []userValue{{
		key:   "companyId",
		value: defaultUUID,
	}}

	GetCompanyExpectedResponseOk = &api.GetCompanyResponse{
		Company: api.Company{
			ID:   defaultUUID,
			Name: "company name",
			Address: api.Address{
				Street:   "Main street, 1",
				Postcode: "12345",
				City:     "Default city",
			},
			Industry: api.Industry{
				ID:           "abc",
				Name:         "retail",
				MarketValue:  1000,
				Co2Footprint: "large",
			},
			Created: float64(DefaultTime.Unix()),
		},
	}

	GetCompanyRepoReq = repository.SearchParams{
		ID: defaultUUID,
	}

	GetCompanyRepoResp = repository.Company{
		ID:   defaultUUID,
		Name: CreateCompanyRequestOk.Name,
		Address: repository.Address{
			Street:   CreateCompanyRequestOk.Address.Street,
			Postcode: CreateCompanyRequestOk.Address.Postcode,
			City:     CreateCompanyRequestOk.Address.City,
		},
		Industry: repository.Industry{
			ID:           CreateCompanyRequestOk.Industry.ID,
			Name:         CreateCompanyRequestOk.Industry.Name,
			MarketValue:  CreateCompanyRequestOk.Industry.MarketValue,
			Co2Footprint: CreateCompanyRequestOk.Industry.Co2Footprint,
		},
		Created: DefaultTime,
	}
)
