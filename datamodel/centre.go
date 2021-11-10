package datamodel

//Centre is cheese
type Centre struct {
	Code        string
	Name        string
	Country     string
	Action      string
	CountryName string
}

const (
	Centre_Title       = "Centre"
	Centre_SQLTable    = "sienaCentre"
	Centre_QueryString = "SienaCentre"
	///
	/// Handler Views
	///
	Centre_TemplateList = "Centre_List"
	Centre_TemplateView = "Centre_View"
	Centre_TemplateEdit = "Centre_Edit"
	Centre_TemplateNew  = "Centre_New"
	///
	/// Field Definitions
	///
	Centre_Code        = "Code"
	Centre_Name        = "Name"
	Centre_Country     = "Country"
	Centre_CountryName = "CountryName"
)
