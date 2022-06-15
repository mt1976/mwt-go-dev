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

func SimulatorSWIFT_GetList_impl() (int, []dm.SimulatorSWIFT, error) {
	var dummy dm.SimulatorSWIFT
	var dummys []dm.SimulatorSWIFT
	dummy.ID = "1"
	dummy.FileName = "1.xml"
	dummy.MessageRaw = "{1:F01RIKOEE2XAXXX0000000000}{2:O3001130191210RIKOEE22XXXX11112222221912191135N}{3:{108:469000-CPTYSPOT}}{4::15A::20:469000-CPTYSPOT:22A:NEWT:94A:BILA:22C:RIKO223804RIKO22:82A:RIKOEE22XXX:87A:RIKOEE22XXX:15B::30T:20200402:30V:20200406:36:1,3804:32B:EUR36221,39:57A:/97822879:33B:USD50000,:53A:TESTUS00XXX:56A:INBKGB00:57A:TESTUS00XXXTESTUS00-}"
	dummy.MessageFmt = "1MessageRaw"
	dummy.Action = "1Action"

	dummys = append(dummys, dummy)

	return len(dummys), dummys, nil
}
func SimulatorSWIFT_GetByID_impl(id string) (int, dm.SimulatorSWIFT, error) {
	var dummy dm.SimulatorSWIFT

	dummy.ID = "1"
	dummy.FileName = "1.xml"
	dummy.MessageFmt = "1MessageFmt"
	dummy.MessageRaw = "{1:F01RIKOEE2XAXXX0000000000}{2:O3001130191210RIKOEE22XXXX11112222221912191135N}{3:{108:469000-CPTYSPOT}}{4::15A::20:469000-CPTYSPOT:22A:NEWT:94A:BILA:22C:RIKO223804RIKO22:82A:RIKOEE22XXX:87A:RIKOEE22XXX:15B::30T:20200402:30V:20200406:36:1,3804:32B:EUR36221,39:57A:/97822879:33B:USD50000,:53A:TESTUS00XXX:56A:INBKGB00:57A:TESTUS00XXXTESTUS00-}"
	dummy.Action = "1Action"
	return 1, dummy, nil
}

func SimulatorSWIFT_NewID_impl(dm.SimulatorSWIFT) string { return "" }
func SimulatorSWIFT_Delete_impl(id string) error         { return nil }

func SimulatorSWIFT_Update_impl(id string, rec dm.SimulatorSWIFT, usr string) error { return nil }

// If there are fields below, create the methods in adaptor\simulatorswift_impl.go

func SimulatorSWIFT_MessageFmt_Override_Store_impl(rec dm.SimulatorSWIFT) error { return nil }

func SimulatorSWIFT_MessageFmt_Override_Fetch_impl(rec dm.SimulatorSWIFT) string { return "" }

func SimulatorSWIFT_Simulator_ProcessResponse_impl(filename string) error {
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
