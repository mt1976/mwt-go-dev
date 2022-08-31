package adaptor

import (
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/contactstreamtype.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : ContactStreamType (contactstreamtype)
// Endpoint 	        : ContactStreamType (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 24/06/2022 at 15:25:05
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in contactstreamtype_impl.go

func ContactStreamType_NewID_impl(rec dm.ContactStreamType) string { return rec.TypeId }
func ContactStreamType_Delete_impl(id string) error                { return nil }

func ContactStreamType_Update_impl(id string, rec dm.ContactStreamType, usr string) error { return nil }

// If there are fields below, create the methods in adaptor\contactstreamtype_impl.go
