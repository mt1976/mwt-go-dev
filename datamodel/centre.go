package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/centre.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Centre (centre)
// Endpoint 	        : Centre (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 04/12/2021 at 17:36:43
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Centre struct {

Code        string
Name        string
Country        string
Country_Impl        string

}

const (
	Centre_Title       = "Centre"
	Centre_SQLTable    = "sienaCentre"
	Centre_SQLSearchID = "Code"
	Centre_QueryString = "Code"
	///
	/// Handler Defintions
	///
	Centre_TemplateList = "Centre_List"
	Centre_TemplateView = "Centre_View"
	Centre_TemplateEdit = "Centre_Edit"
	Centre_TemplateNew  = "Centre_New"
	///
	/// Handler Monitor Paths
	///
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
	Centre_Country_Impl   = "Country_Impl" // Country_Impl is a String

	/// Definitions End
)
