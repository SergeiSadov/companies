package validators

import (
	"unicode/utf8"

	"github.com/google/uuid"

	"companies/internal/models/api"
	"companies/pkg/errors"
)

type IValidator interface {
	ValidateCreateRequest(req *api.CreateCompanyRequest) (err error)
	ValidateUpdateRequest(req *api.UpdateCompanyRequest) (err error)
}

type Validator struct {
	validationParams api.CompanyValidationParams
}

func New(validationParams api.CompanyValidationParams) *Validator {
	return &Validator{
		validationParams: validationParams,
	}
}

func (v *Validator) ValidateCreateRequest(req *api.CreateCompanyRequest) (err error) {
	if utf8.RuneCountInString(req.Name) > v.validationParams.CompanyNameLen || len(req.Name) == 0 {
		return errors.ErrInvalidCompanyName
	}

	if utf8.RuneCountInString(req.Industry.Name) > v.validationParams.IndustryNameLen || len(req.Industry.Name) == 0 {
		return errors.ErrInvalidIndustryName
	}

	if utf8.RuneCountInString(req.Industry.Co2Footprint) > v.validationParams.Co2FootprintLen || len(req.Industry.Co2Footprint) == 0 {
		return errors.ErrInvalidCo2Footprint
	}

	if req.Industry.ID != "" {
		if _, err := uuid.Parse(req.Industry.ID); err != nil {
			return errors.ErrInvalidUUID
		}
	}

	return nil
}

func (v *Validator) ValidateUpdateRequest(req *api.UpdateCompanyRequest) (err error) {
	if !v.validUUID(req.ID) {
		return errors.ErrInvalidUUID
	}

	if utf8.RuneCountInString(req.Name) > v.validationParams.CompanyNameLen || len(req.Name) == 0 {
		return errors.ErrInvalidCompanyName
	}

	if utf8.RuneCountInString(req.Industry.Name) > v.validationParams.IndustryNameLen || len(req.Industry.Name) == 0 {
		return errors.ErrInvalidIndustryName
	}

	if utf8.RuneCountInString(req.Industry.Co2Footprint) > v.validationParams.Co2FootprintLen || len(req.Industry.Co2Footprint) == 0 {
		return errors.ErrInvalidCo2Footprint
	}

	if req.Industry.ID != "" && !v.validUUID(req.Industry.ID) {
		return errors.ErrInvalidUUID
	}

	return nil
}

func (v *Validator) validUUID(input string) (valid bool) {
	if _, err := uuid.Parse(input); err != nil {
		return false
	}

	return true
}
