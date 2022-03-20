package tests

import (
	"time"

	"companies/internal/models/api"
	"companies/internal/models/repository"
)

const (
	defaultUUID   = "4f6b2a25-8616-4855-89a8-44cb68b76fbf"
	string256Char = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam maximus elit sed dui interdum, sed dapibus sem condimentum. Nunc consectetur nisi rutrum est efficitur, ac tristique sem convallis. Morbi sit amet tortor erat. Etiam efficitur turpis sit amet nunc rutrum dictum. Cras ultrices ullamcorper velit vitae auctor. Cras eu convallis urna, non rutrum urna. Morbi nec mollis est. Suspendisse nec sodales sem, rhoncus semper sapien. Vivamus consequat quam vel lorem viverra consectetur. Duis est neque, ullamcorper nec nibh et, imperdiet eleifend magna. Praesent vestibulum elementum nulla, cursus hendrerit mauris vehicula et. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. In lacinia, ligula et faucibus faucibus, libero dolor auctor sapien, non accumsan urna tellus eu libero.\n\nCras luctus interdum nibh, eget placerat dui rutrum et. Integer condimentum mauris magna, vel feugiat nunc hendrerit sed. Integer blandit justo non ante auctor lobortis. Quisque sollicitudin sed nisi at malesuada. Nam eu tortor diam. Nullam libero nisl, fermentum et imperdiet in, ullamcorper sit amet neque. Aenean eget urna ex. Cras lacinia gravida pulvinar. Donec egestas dapibus dolor, non tempus mi dignissim et. Sed ullamcorper, tellus in elementum auctor, lacus nisi ullamcorper mi, in venenatis turpis sapien id felis. Vivamus egestas lobortis tortor, vel scelerisque nulla feugiat sed. Mauris velit ipsum, posuere vitae lacus at, feugiat malesuada diam. Nullam eget erat arcu.\n\nCurabitur auctor orci dolor, ut pulvinar mi tempor eget. Sed at erat in neque elementum suscipit in tempus libero. Ut risus turpis, scelerisque posuere ornare sit amet, egestas sit amet purus. Nullam eu quam non."
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
			ID:           defaultUUID,
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
				ID:           defaultUUID,
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
