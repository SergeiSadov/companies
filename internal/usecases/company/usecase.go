package company

import (
	"context"

	"companies/internal/models/api"
	"companies/internal/repositories/company"
)

type IUseCase interface {
	Create(ctx context.Context, req *api.CreateCompanyRequest) (response *api.CreateCompanyResponse, err error)
	Get(ctx context.Context, req *api.GetCompanyRequest) (response *api.GetCompanyResponse, err error)
	Update(ctx context.Context, req *api.UpdateCompanyRequest) (response *api.UpdateCompanyResponse, err error)
	Delete(ctx context.Context, req *api.DeleteCompanyRequest) (err error)
}

type UseCase struct {
	companyRepository company.IRepository
	adapter           IAdapter
}

func New(
	companyRepository company.IRepository,
	adapter IAdapter,
) *UseCase {
	return &UseCase{
		companyRepository: companyRepository,
		adapter:           adapter,
	}
}

func (u *UseCase) Create(ctx context.Context, req *api.CreateCompanyRequest) (response *api.CreateCompanyResponse, err error) {
	created, err := u.companyRepository.Create(ctx, u.adapter.AdaptCreateReqToRepo(req))
	if err != nil {
		return response, err
	}

	return u.adapter.AdaptRepoToCreateResp(created), nil
}

func (u *UseCase) Get(ctx context.Context, req *api.GetCompanyRequest) (response *api.GetCompanyResponse, err error) {
	res, err := u.companyRepository.Get(ctx, u.adapter.AdaptGetReqToRepo(req))
	if err != nil {
		return response, err
	}

	return u.adapter.AdaptRepoToGetResp(res), nil
}

func (u *UseCase) Update(ctx context.Context, req *api.UpdateCompanyRequest) (response *api.UpdateCompanyResponse, err error) {
	created, err := u.companyRepository.Update(ctx, u.adapter.AdaptUpdateReqToRepo(req))
	if err != nil {
		return response, err
	}

	return u.adapter.AdaptRepoToUpdateResp(created), nil
}

func (u *UseCase) Delete(ctx context.Context, req *api.DeleteCompanyRequest) (err error) {
	return u.companyRepository.Delete(ctx, u.adapter.AdaptDeleteReqToRepo(req))
}
