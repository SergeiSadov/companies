package tests

import (
	"time"

	"companies/internal/models/api"
	"companies/internal/models/repository"
)

const (
	defaultUUID = "4f6b2a25-8616-4855-89a8-44cb68b76fbf"
)

var (
	DefaultTime = time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC)

	CreateCompanyRequestOk = &api.CreateCompanyRequest{
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
	}

	CreateCompanyExpectedResponseOk = &api.CreateCompanyResponse{
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

	CreateCompanyRepoReq = repository.Company{
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
	}

	CreateCompanyRepoResp = repository.Company{
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
