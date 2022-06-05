package adaptor

import (
	"strings"

	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/simulatorswift.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : SimulatorSWIFT (simulatorswift)
// Endpoint 	        : SimulatorSWIFT (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 05/06/2022 at 12:40:29
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

// The following functions should be created in simulatorswift_impl.go

func SimulatorSWIFT_GetList_impl() (int, []dm.SimulatorSWIFT, error) { return 0, nil, nil }
func SimulatorSWIFT_GetByID_impl(id string) (int, dm.SimulatorSWIFT, error) {
	return 0, dm.SimulatorSWIFT{}, nil
}

func SimulatorSWIFT_NewID_impl(dm.SimulatorSWIFT) string { return "" }
func SimulatorSWIFT_Delete_impl(id string) error         { return nil }

func SimulatorSWIFT_Update_impl(id string, rec dm.SimulatorSWIFT, usr string) error { return nil }

// If there are fields below, create the methods in adaptor\simulatorswift_impl.go

func SimulatorSWIFT_MessageFmt_Override_Store_impl(rec dm.SimulatorSWIFT) error { return nil }

func SimulatorSWIFT_MessageFmt_Override_Fetch_impl(rec dm.SimulatorSWIFT) string { return "" }

func Simulator_SimulatorSWIFT_ProcessResponse_impl(filename string) error {
	logs.Success("SimulatorSWIFT_ProcessResponse:" + filename)
	// tokenise the filename, get last element
	tokens := strings.Split(filename, "/")
	last := tokens[len(tokens)-1]
	core.Notification_Normal("New SWIFT Simulator Message Detected", last)
	return simulatorswift_ProcessResponse_impl(tokens, last)
}

func simulatorswift_ProcessResponse_impl(tokens []string, latestToken string) error {
	return nil
}
