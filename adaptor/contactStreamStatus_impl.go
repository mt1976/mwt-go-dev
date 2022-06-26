package adaptor

import (
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/contactstreamstatus.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : ContactStreamStatus (contactstreamstatus)
// Endpoint 	        : ContactStreamStatus (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 24/06/2022 at 15:41:21
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in contactstreamstatus_impl.go

func ContactStreamStatus_NewID_impl(rec dm.ContactStreamStatus) string { return rec.StatusId }
func ContactStreamStatus_Delete_impl(id string) error                  { return nil }

func ContactStreamStatus_Update_impl(id string, rec dm.ContactStreamStatus, usr string) error {
	return nil
}

// If there are fields below, create the methods in adaptor\contactstreamstatus_impl.go
