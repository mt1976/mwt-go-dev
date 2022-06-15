package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/centre.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Centre (centre)
// Endpoint 	        : Centre (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:31:51
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//Centre defines the datamolde for the Centre object
type Centre struct {

Code       string
Name       string
Country       string
Country_lookup []Lookup_Item

}

const (
	Centre_Title       = "Centre"
	Centre_SQLTable    = "sienaCentre"
	Centre_SQLSearchID = "Code"
	Centre_QueryString = "Code"
	///
	/// Handler Defintions
	///
	Centre_Template     = "Centre"
	Centre_TemplateList = "Centre_List"
	Centre_TemplateView = "Centre_View"
	Centre_TemplateEdit = "Centre_Edit"
	Centre_TemplateNew  = "Centre_New"
	///
	/// Handler Monitor Paths
	///
	Centre_Path       = "/API/Centre/"
	Centre_PathList   = "/CentreList/"
	Centre_PathView   = "/CentreView/"
	Centre_PathEdit   = "/CentreEdit/"
	Centre_PathNew    = "/CentreNew/"
	Centre_PathSave   = "/CentreSave/"
	Centre_PathDelete = "/CentreDelete/"
	///
	/// SQL Field Definitions
	///
Centre_Code   = "Code" // Code is a String
Centre_Name   = "Name" // Name is a String
Centre_Country   = "Country" // Country is a String

	/// Definitions End
)
