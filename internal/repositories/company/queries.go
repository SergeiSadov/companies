package company

const (
	InsertIntoCompanies = "INSERT INTO " +
		TableCompanies + " (name, address, industry_id) VALUES (?, ?, ?) RETURNING id, created"
	InsertIntoIndustries = "INSERT INTO " + TableIndustries +
		" (name, market_value, co2_footprint) VALUES (?, ?, ?) RETURNING id"
	DeleteFromIndustries = "DELETE FROM industries WHERE id NOT IN (SELECT industry_id FROM companies)"
)

const (
	ParamID = "id = ?"
)
