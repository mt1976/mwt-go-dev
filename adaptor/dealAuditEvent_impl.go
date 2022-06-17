package adaptor

import (
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/dealauditevent.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : DealAuditEvent (dealauditevent)
// Endpoint 	        : DealAuditEvent (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/06/2022 at 16:50:10
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in dealauditevent_impl.go

func DealAuditEvent_NewID_impl(rec dm.DealAuditEvent) string { return rec.InternalId }
func DealAuditEvent_Delete_impl(id string) error             { return nil }

func DealAuditEvent_Update_impl(id string, rec dm.DealAuditEvent, usr string) error { return nil }

// If there are fields below, create the methods in adaptor\dealauditevent_impl.go
