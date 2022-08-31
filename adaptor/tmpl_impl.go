package adaptor

import (
	"strings"
	"time"

	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/tmpl.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Tmpl (tmpl)
// Endpoint 	        : Tmpl (TemplateID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 25/06/2022 at 13:24:13
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in tmpl_impl.go

// If there are fields below, create the methods in adaptor\tmpl_impl.go
// START - GET API/Callout
// Dynamically generated 25/06/2022 by matttownsend (Matt Townsend) on silicon.local
//
func Tmpl_ExtraField_OnStore_impl(fieldval string, rec dm.Tmpl, usr string) (string, error) {
	logs.Callout("Tmpl", "ExtraField", PUT, rec.ID)
	return fieldval, nil
}
func Tmpl_ExtraField3_OnStore_impl(fieldval string, rec dm.Tmpl, usr string) (string, error) {
	logs.Callout("Tmpl", "ExtraField3", PUT, rec.ID)
	return fieldval, nil
}
func Tmpl_TDate_OnStore_impl(fieldval string, rec dm.Tmpl, usr string) (string, error) {
	logs.Callout("Tmpl", "TDate", PUT, rec.ID)
	return fieldval, nil
}

//
// Dynamically generated 25/06/2022 by matttownsend (Matt Townsend) on silicon.local
// END - GET API/Callout
//
// START - PUT API/Callout
// Dynamically generated 25/06/2022 by matttownsend (Matt Townsend) on silicon.local
//
func Tmpl_ExtraField_OnFetch_impl(rec dm.Tmpl) string {
	logs.Callout("Tmpl", "ExtraField", GET, rec.ID)
	return rec.ExtraField
}
func Tmpl_ExtraField3_OnFetch_impl(rec dm.Tmpl) string {
	logs.Callout("Tmpl", "ExtraField3", GET, rec.ID)
	return rec.ExtraField3
}
func Tmpl_TDate_OnFetch_impl(rec dm.Tmpl) string {
	logs.Callout("Tmpl", "TDate", GET, rec.ID)
	return rec.TDate
}

//
// Dynamically generated 25/06/2022 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 25/06/2022 by matttownsend (Matt Townsend) on silicon.local
//
// Tmpl_ExtraField_impl provides validation/actions for ExtraField
func Tmpl_ExtraField_impl(iAction string, iId string, iValue string, iRec dm.Tmpl, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Tmpl", "ExtraField", VAL+"-"+iAction, iId)

	fP.MsgType = core.FieldMessage_POSITIVE
	fP.MsgMessage = "WOW! MUCH SUCCESS"
	//	fP.ErrorCode = "0"
	fP.MsgLog = true
	fP.MsgGlyph = "far fa-thumbs-up"

	fP = dm.AddFieldMessage(fP, core.FieldMessage_POSITIVE, "WOW! MUCH SUCCESS", false, "far fa-thumbs-up")

	return "Sploot", fP
}

// Tmpl_ExtraField3_impl provides validation/actions for ExtraField3
func Tmpl_ExtraField3_impl(iAction string, iId string, iValue string, iRec dm.Tmpl, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Tmpl", "ExtraField3", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("Tmpl_TDate_impl" + " - Invalid Action")
	}
	fP = dm.AddFieldMessage(fP, core.FieldMessage_ERROR, "WOW! MUCH FAILURE", false, "far fa-thumbs-down")
	return iAction, fP
}

// Tmpl_TDate_impl provides validation/actions for TDate
func Tmpl_TDate_impl(iAction string, iId string, iValue string, iRec dm.Tmpl, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Tmpl", "TDate", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:
		t := time.Now()
		fP.DefaultValue = t.Format("2006-01-02")
		logs.Information("SetDate", fP.DefaultValue)
	case PUT:

	case GET:

	default:
		logs.Warning("Tmpl_TDate_impl" + " - Invalid Action")
	}
	return fP.DefaultValue, fP
}

//
// Dynamically generated 25/06/2022 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout

// If there are fields below, create the methods in adaptor\tmpl_impl.go
// START - GET API/Callout
// Dynamically generated 25/06/2022 by matttownsend (Matt Townsend) on silicon.local
//
func Tmpl_ID_OnStore_impl(fieldval string, rec dm.Tmpl, usr string) (string, error) {
	logs.Callout("Tmpl", "ID", PUT, rec.ID)
	return fieldval, nil
}

//
// Dynamically generated 25/06/2022 by matttownsend (Matt Townsend) on silicon.local
// END - GET API/Callout
//
// START - PUT API/Callout
// Dynamically generated 25/06/2022 by matttownsend (Matt Townsend) on silicon.local
//
func Tmpl_ID_OnFetch_impl(rec dm.Tmpl) string {
	logs.Callout("Tmpl", "ID", GET, rec.ID)
	return rec.ID
}

//
// Dynamically generated 25/06/2022 by matttownsend (Matt Townsend) on silicon.local
// END - PUT API/Callout

// START - Validation API/Callout
// Dynamically generated 25/06/2022 by matttownsend (Matt Townsend) on silicon.local
//
// Tmpl_ID_impl provides validation/actions for ID
func Tmpl_ID_impl(iAction string, iId string, iValue string, iRec dm.Tmpl, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Tmpl", "ID", VAL+"-"+iAction, iId+" - "+iValue)
	iValue = strings.TrimSpace(iValue)
	switch iAction {
	case VAL:

	case NEW:
		iValue = core.GetUUID()
	case PUT:

	case GET:

	default:
		logs.Warning("Tmpl_TDate_impl" + " - Invalid Action")
	}

	fP = dm.AddFieldMessage(fP, core.FieldMessage_NORMAL, "WOW! MUCH NORMAL", false, "far fa-thumbs-down")
	logs.Callout("Tmpl", "ID", VAL+"-"+iAction, iId+" - "+iValue)
	return iValue, fP
}

//
// Dynamically generated 25/06/2022 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout

func Tmpl_ObjectValidation_impl(iAction string, iId string, iRec dm.Tmpl) (dm.Tmpl, error, string) {
	logs.Callout("Tmpl", "ObjectValidation", VAL+"-"+iAction, iId)
	switch iAction {
	case VAL:

	case NEW:

	case PUT:

	case GET:

	default:
		logs.Warning("Tmpl_TDate_impl" + " - Invalid Action")
	}
	return iRec, nil, ""
}
