package validators

import "companies/internal/models/api"

var (
	PreparedValidatorParams = api.CompanyValidationParams{
		CompanyNameLen:  255,
		IndustryNameLen: 255,
		Co2FootprintLen: 500,
	}
)
