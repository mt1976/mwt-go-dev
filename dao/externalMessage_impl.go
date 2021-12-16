package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/externalmessage.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : ExternalMessage (externalmessage)
// Endpoint 	        : ExternalMessage (MessageID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/12/2021 at 09:31:50
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	core "github.com/mt1976/mwt-go-dev/core"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// ExternalMessage_GetByID() returns a single ExternalMessage record
func ExternalMessage_GetActive() (int, []dm.ExternalMessage, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.ExternalMessage_SQLTable)
	tsql = tsql + " WHERE " + dm.ExternalMessage_ResponseDate + "= ''"
	tsql = tsql + " AND " + dm.ExternalMessage_MessageACKNAK + "= ''"
	noItems, externalmessageList, _, _ := externalmessage_Fetch(tsql)

	return noItems, externalmessageList, nil
}
