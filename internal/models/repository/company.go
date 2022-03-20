package repository

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Company struct {
	ID         string    `gorm:"primaryKey;<-:false"`
	Name       string    `gorm:"name"`
	Address    Address   `gorm:"address"`
	IndustryID string    `gorm:"industry_id"`
	Industry   Industry  `gorm:"industries;foreignKey:IndustryID;references:id"`
	Created    time.Time `gorm:"<-:false"`
}

type Address struct {
	Street   string `gorm:"street" json:"street"`
	Postcode string `gorm:"postcode" json:"postcode"`
	City     string `gorm:"city" json:"city"`
}

func (a *Address) Value() (v driver.Value, err error) {
	return json.Marshal(a)
}

func (a *Address) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("invalid data")
	}

	return json.Unmarshal(b, &a)
}

type Industry struct {
	ID           string  `gorm:"<-:false"`
	Name         string  `gorm:"name"`
	MarketValue  float64 `gorm:"market_value"`
	Co2Footprint string  `gorm:"co2_footprint"`
}

type CompanyIndustries struct {
	CompanyID  string `gorm:"company_id"`
	IndustryID string `gorm:"industry_id"`
}

type SearchParams struct {
	ID string
}
