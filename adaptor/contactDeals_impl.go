package adaptor

import (
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/contactdeals.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : ContactDeals (contactdeals)
// Endpoint 	        : ContactDeals (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 24/06/2022 at 16:01:55
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in contactdeals_impl.go

func ContactDeals_NewID_impl(rec dm.ContactDeals) string { return rec.NoteId }
func ContactDeals_Delete_impl(id string) error           { return nil }

func ContactDeals_Update_impl(id string, rec dm.ContactDeals, usr string) error { return nil }

// If there are fields below, create the methods in adaptor\contactdeals_impl.go
