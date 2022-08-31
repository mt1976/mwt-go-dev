package adaptor

import (
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/contactdealrefno.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : ContactDealRefNo (contactdealrefno)
// Endpoint 	        : ContactDealRefNo (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 24/06/2022 at 16:06:32
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in contactdealrefno_impl.go

func ContactDealRefNo_NewID_impl(rec dm.ContactDealRefNo) string { return rec.NoteId }
func ContactDealRefNo_Delete_impl(id string) error               { return nil }

func ContactDealRefNo_Update_impl(id string, rec dm.ContactDealRefNo, usr string) error { return nil }

// If there are fields below, create the methods in adaptor\contactdealrefno_impl.go
