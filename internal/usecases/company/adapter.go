package company

import (
	"companies/internal/models/api"
	"companies/internal/models/repository"
)

type IAdapter interface {
	AdaptCreateReqToRepo(req *api.CreateCompanyRequest) (adapted *repository.Company)
	AdaptRepoToCreateResp(req *repository.Company) (adapted *api.CreateCompanyResponse)

	AdaptUpdateReqToRepo(req *api.UpdateCompanyRequest) (adapted *repository.Company)
	AdaptRepoToUpdateResp(req *repository.Company) (adapted *api.UpdateCompanyResponse)

	AdaptGetReqToRepo(req *api.GetCompanyRequest) (adapted *repository.SearchParams)
	AdaptRepoToGetResp(req *repository.Company) (adapted *api.GetCompanyResponse)

	AdaptDeleteReqToRepo(req *api.DeleteCompanyRequest) (adapted *repository.SearchParams)
}

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a *Adapter) AdaptCreateReqToRepo(req *api.CreateCompanyRequest) (adapted *repository.Company) {
	return &repository.Company{
		Name: req.Name,
		Address: repository.Address{
			Street:   req.Address.Street,
			Postcode: req.Address.Postcode,
			City:     req.Address.City,
		},
		Industry: repository.Industry{
			ID:           req.Industry.ID,
			Name:         req.Industry.Name,
			MarketValue:  req.Industry.MarketValue,
			Co2Footprint: req.Industry.Co2Footprint,
		},
	}
}

func (a *Adapter) AdaptRepoToCreateResp(req *repository.Company) (adapted *api.CreateCompanyResponse) {
	return &api.CreateCompanyResponse{
		Company: api.Company{
			ID:   req.ID,
			Name: req.Name,
			Address: api.Address{
				Street:   req.Address.Street,
				Postcode: req.Address.Postcode,
				City:     req.Address.City,
			},
			Industry: api.Industry{
				ID:           req.Industry.ID,
				Name:         req.Industry.Name,
				MarketValue:  req.Industry.MarketValue,
				Co2Footprint: req.Industry.Co2Footprint,
			},
			Created: float64(req.Created.Unix()),
		},
	}
}

func (a *Adapter) AdaptUpdateReqToRepo(req *api.UpdateCompanyRequest) (adapted *repository.Company) {
	return &repository.Company{
		ID:   req.ID,
		Name: req.Name,
		Address: repository.Address{
			Street:   req.Address.Street,
			Postcode: req.Address.Postcode,
			City:     req.Address.City,
		},
		Industry: repository.Industry{
			ID:           req.Industry.ID,
			Name:         req.Industry.Name,
			MarketValue:  req.Industry.MarketValue,
			Co2Footprint: req.Industry.Co2Footprint,
		},
	}
}

func (a *Adapter) AdaptRepoToUpdateResp(req *repository.Company) (adapted *api.UpdateCompanyResponse) {
	return &api.UpdateCompanyResponse{
		Company: api.Company{
			ID:   req.ID,
			Name: req.Name,
			Address: api.Address{
				Street:   req.Address.Street,
				Postcode: req.Address.Postcode,
				City:     req.Address.City,
			},
			Industry: api.Industry{
				ID:           req.Industry.ID,
				Name:         req.Industry.Name,
				MarketValue:  req.Industry.MarketValue,
				Co2Footprint: req.Industry.Co2Footprint,
			},
			Created: float64(req.Created.Unix()),
		},
	}
}

func (a *Adapter) AdaptGetReqToRepo(req *api.GetCompanyRequest) (adapted *repository.SearchParams) {
	return &repository.SearchParams{ID: req.ID}
}

func (a *Adapter) AdaptRepoToGetResp(req *repository.Company) (adapted *api.GetCompanyResponse) {
	return &api.GetCompanyResponse{
		Company: api.Company{
			ID:   req.ID,
			Name: req.Name,
			Address: api.Address{
				Street:   req.Address.Street,
				Postcode: req.Address.Postcode,
				City:     req.Address.City,
			},
			Industry: api.Industry{
				ID:           req.Industry.ID,
				Name:         req.Industry.Name,
				MarketValue:  req.Industry.MarketValue,
				Co2Footprint: req.Industry.Co2Footprint,
			},
			Created: float64(req.Created.Unix()),
		},
	}
}

func (a *Adapter) AdaptDeleteReqToRepo(req *api.DeleteCompanyRequest) (adapted *repository.SearchParams) {
	return &repository.SearchParams{ID: req.ID}
}
