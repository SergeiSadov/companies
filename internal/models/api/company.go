package api

type Company struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Address  Address  `json:"address"`
	Industry Industry `json:"industry"`
	Created  float64  `json:"created"`
}

type Address struct {
	Street   string `json:"street"`
	Postcode string `json:"postcode"`
	City     string `json:"city"`
}

type Industry struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	MarketValue  float64 `json:"marketValue"`
	Co2Footprint string  `json:"co2Footprint"`
}

type CreateCompanyRequest struct {
	Name     string   `json:"name"`
	Address  Address  `json:"address"`
	Industry Industry `json:"industry"`
}

type CreateCompanyResponse struct {
	Company
}

type UpdateCompanyRequest struct {
	Company
}

type UpdateCompanyResponse struct {
	Company
}

type GetCompanyRequest struct {
	ID string
}

type GetCompanyResponse struct {
	Company
}

type DeleteCompanyRequest struct {
	ID string
}
