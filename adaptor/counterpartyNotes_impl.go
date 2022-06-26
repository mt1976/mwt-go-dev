package adaptor

import (
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/counterpartynotes.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : CounterpartyNotes (counterpartynotes)
// Endpoint 	        : CounterpartyNotes (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 24/06/2022 at 15:04:35
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in counterpartynotes_impl.go

func CounterpartyNotes_NewID_impl(rec dm.CounterpartyNotes) string { return rec.NoteId }
func CounterpartyNotes_Delete_impl(id string) error                { return nil }

func CounterpartyNotes_Update_impl(id string, rec dm.CounterpartyNotes, usr string) error { return nil }

// If there are fields below, create the methods in adaptor\counterpartynotes_impl.go
