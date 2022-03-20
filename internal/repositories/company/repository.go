package company

import (
	"context"

	"gorm.io/gorm"

	"companies/internal/models/repository"
	"companies/pkg/error_adapters/sql_adapter"
)

const (
	TableCompanies  = "companies"
	TableIndustries = "industries"

	FieldIndustry = "Industry"
)

type IRepository interface {
	Create(ctx context.Context, company *repository.Company) (created *repository.Company, err error)
	Get(ctx context.Context, params *repository.SearchParams) (company *repository.Company, err error)
	Update(ctx context.Context, company *repository.Company) (updated *repository.Company, err error)
	Delete(ctx context.Context, params *repository.SearchParams) (err error)
}

type Repository struct {
	db           *gorm.DB
	errorAdapter sql_adapter.IErrorAdapter
}

func NewRepository(db *gorm.DB, errorAdapter sql_adapter.IErrorAdapter) *Repository {
	return &Repository{
		db:           db,
		errorAdapter: errorAdapter,
	}
}

func (r *Repository) Create(ctx context.Context, company *repository.Company) (created *repository.Company, err error) {
	if err = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) (err error) {
		if err = r.upsertIndustry(tx, &company.Industry); err != nil {
			return
		}
		company.IndustryID = company.Industry.ID

		if err = r.createCompany(tx, company); err != nil {
			return
		}

		return
	}); err != nil {
		return created, r.errorAdapter.AdaptSqlErr(err)
	}

	return company, nil
}

func (r *Repository) Get(ctx context.Context, params *repository.SearchParams) (company *repository.Company, err error) {
	if err = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) (err error) {
		company = new(repository.Company)

		if err = tx.Table(TableCompanies).Where(ParamID, params.ID).Find(&company).Error; err != nil {
			return
		}

		if company.ID == "" {
			return ErrNotFound
		}

		if err = tx.Model(&company).Association(FieldIndustry).Find(&company.Industry); err != nil {
			return
		}

		return
	}); err != nil {
		return company, r.errorAdapter.AdaptSqlErr(err)
	}

	return company, nil
}

func (r *Repository) Update(ctx context.Context, company *repository.Company) (updated *repository.Company, err error) {
	if err = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) (err error) {
		if err = r.upsertIndustry(tx, &company.Industry); err != nil {
			return
		}

		if err = r.updateCompany(tx, company); err != nil {
			return
		}
		return
	}); err != nil {
		return updated, r.errorAdapter.AdaptSqlErr(err)
	}

	return company, nil
}

func (r *Repository) Delete(ctx context.Context, params *repository.SearchParams) (err error) {
	if err = r.db.Debug().WithContext(ctx).Transaction(func(tx *gorm.DB) (err error) {
		res := tx.Table(TableCompanies).Delete(&repository.Company{ID: params.ID})
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return ErrNotFound
		}

		//delete industries with no relations left
		if err = tx.Exec(DeleteFromIndustries).Error; err != nil {
			return err
		}

		return
	}); err != nil {
		return r.errorAdapter.AdaptSqlErr(err)
	}

	return nil
}

func (r *Repository) createCompany(tx *gorm.DB, company *repository.Company) (err error) {
	if err = tx.Raw(InsertIntoCompanies, company.Name, company.Address, company.IndustryID).
		Scan(&company).Error; err != nil {
		return
	}

	return
}

func (r *Repository) upsertIndustry(tx *gorm.DB, industry *repository.Industry) (err error) {
	if industry.ID != "" && tx.Table(TableIndustries).Where(ParamID, industry.ID).Updates(&industry).RowsAffected != 0 {
		return
	}

	return tx.Raw(InsertIntoIndustries, industry.Name, industry.MarketValue, industry.Co2Footprint).Scan(&industry.ID).Error
}

func (r *Repository) updateCompany(tx *gorm.DB, company *repository.Company) (err error) {
	res := tx.Table(TableCompanies).Where(ParamID, company.ID).Updates(&company).Find(&company)
	if err = res.Error; err != nil {
		return err
	}

	if res.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
