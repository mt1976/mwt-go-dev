package adaptor

import (
	"strings"

	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/test1.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Test1 (test1)
// Endpoint 	        : Test1 (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 12:41:06
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in test1_impl.go

func Test1_GetList_impl() (int, []dm.Test1, error)        { return 0, nil, nil }
func Test1_GetByID_impl(id string) (int, dm.Test1, error) { return 0, dm.Test1{}, nil }

func Test1_GetByReverseLookup_impl(id string, lookupField string) (int, dm.Test1, error) {
	return 0, dm.Test1{}, nil
}

func Test1_NewID_impl(rec dm.Test1) string { return rec.ID }
func Test1_Delete_impl(id string) error    { return nil }

func Test1_Update_impl(id string, rec dm.Test1, usr string) error { return nil }

// If there are fields below, create the methods in adaptor\test1_impl.go

func Test1_Cheese_OnStore_impl(fieldval string, rec dm.Test1, usr string) (string, error) {
	return fieldval, nil
}
func Test1_Onion_OnStore_impl(fieldval string, rec dm.Test1, usr string) (string, error) {
	return fieldval, nil
}

func Test1_Cheese_OnFetch_impl(rec dm.Test1) string { return "" }
func Test1_Onion_OnFetch_impl(rec dm.Test1) string  { return "" }

func Test1_Simulator_ProcessResponse_impl(filename string) error {
	logs.Success("Test1_Simulator_ProcessResponse:" + filename)
	// tokenise the filename, get last element
	tokens := strings.Split(filename, "/")
	last := tokens[len(tokens)-1]
	core.Notification_Normal("New Test1Template Message Detected", last)
	return test1_Simulator_ProcessResponse_impl(tokens, last)
}

func test1_Simulator_ProcessResponse_impl(tokens []string, latestToken string) error {
	return nil
}

// If there are fields below, create the methods in adaptor\test1_impl.go
// START - GET API/Callout
// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
//

// START - Validation API/Callout
// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
//
// Test1_Cheese_Impl provides validation/actions for Cheese
func Test1_Cheese_Impl(iAction string, iId string, iValue string, iRec dm.Test1, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Test1", "Cheese", VAL+"-"+iAction, iId)
	return "", fP
}

// Test1_Onion_Impl provides validation/actions for Onion
func Test1_Onion_Impl(iAction string, iId string, iValue string, iRec dm.Test1, fP dm.FieldProperties) (string, dm.FieldProperties) {
	logs.Callout("Test1", "Onion", VAL+"-"+iAction, iId)
	return "", fP
}

//
// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
// END - Validation API/Callout
