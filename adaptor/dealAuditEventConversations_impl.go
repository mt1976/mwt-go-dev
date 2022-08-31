package adaptor

import (
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/dealauditeventconversations.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : DealAuditEventConversations (dealauditeventconversations)
// Endpoint 	        : DealAuditEventConversations (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 24/06/2022 at 16:47:42
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in dealauditeventconversations_impl.go

func DealAuditEventConversations_NewID_impl(rec dm.DealAuditEventConversations) string {
	return rec.InternalId
}
func DealAuditEventConversations_Delete_impl(id string) error { return nil }

func DealAuditEventConversations_Update_impl(id string, rec dm.DealAuditEventConversations, usr string) error {
	return nil
}

// If there are fields below, create the methods in adaptor\dealauditeventconversations_impl.go
