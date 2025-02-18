package adaptor

import (
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// ----------------------------------------------------------------

// Automatically generated  "/adaptor/contactstream.go"

// ----------------------------------------------------------------

// Package              : adaptor

// Object 			    : ContactStream (contactstream)

// Endpoint 	        : ContactStream (ID)

// For Project          : github.com/mt1976/mwt-go-dev/

// ----------------------------------------------------------------

// Template Generator   : delinquentDysprosium [r4-21.12.31]

// Date & Time		    : 24/06/2022 at 15:25:01

// Who & Where		    : matttownsend (Matt Townsend) on silicon.local

// ----------------------------------------------------------------

// The following functions should be created in contactstream_impl.go

func ContactStream_NewID_impl(rec dm.ContactStream) string { return rec.StreamId }

func ContactStream_Delete_impl(id string) error { return nil }

func ContactStream_Update_impl(id string, rec dm.ContactStream, usr string) error { return nil }

// If there are fields below, create the methods in adaptor\contactstream_impl.go
