package datamodel

import (
	"github.com/andreyvit/sqlexpr"
)

//Country is cheese
type Country struct {
	Code      string
	Name      string
	ShortCode string
	EU_EEA    string
	Action    string
}

// const (
// 	Country_Code      = "Code"
// 	Country_Name      = "Name"
// 	Country_ShortCode = "ShortCode"
// 	Country_EU_EEA    = "EU_EEA"
// )

//appCacheStoreItem is cheese
const (
	//Action     string
	Country_Title     = "Countries"
	Country_Table     = "sienaCountry"
	Country_Code      = "Code"
	Country_Name      = "Name"
	Country_ShortCode = "ShortCode"
	Country_EU_EEA    = "EU_EEA"
	// Define SQL Components

	Country_SQLCode      = sqlexpr.Column(Country_Code)
	Country_SQLName      = sqlexpr.Column(Country_Name)
	Country_SQLShortCode = sqlexpr.Column(Country_ShortCode)
	Country_SQLEU_EEA    = sqlexpr.Column(Country_EU_EEA)
)
