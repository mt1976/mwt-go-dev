package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/centre.go"
// ----------------------------------------------------------------
// Package            : datamodel
// Object 			  : Centre
// Endpoint Root 	  : Centre
// Search QueryString : Code
// From   			  : 
// ----------------------------------------------------------------
// Template Generator : RussetAlbatross [r0-21.11.01]
// ----------------------------------------------------------------
// Date & Time		  : 19/11/2021 at 17:16:03
// Who & Where		  : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Centre struct {
	AppInternalID string  // Special field for internal use only
Code        string
Name        string
Country        string
Country_Enri        string

}

const (
	Centre_Title       = "Centres"
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
	Centre_Country_Enri   = "Country_Enri" // Country_Enri is a String

	/// Definitions End
)
