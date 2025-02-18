package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/schedule.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Schedule (schedule)
// Endpoint 	        : Schedule (Schedule)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 21/11/2021 at 12:29:51
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

func schedule_NewIDImpl(r dm.Schedule) string {

	//schedule_NewIDImpl should be specified in dao/Schedule_Impl.go
	id := r.Name + core.IDSep + r.Type

	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------
