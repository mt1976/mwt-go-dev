package globals

import (
	"database/sql"
	"time"
)

var SessionToken = ""
var UUID = "authorAdjust"
var SecurityViolation = ""
var DB *sql.DB
var UserRole = "/default"
var UserName = ""
var UserKnowAs = ""
var UserNavi = ""
var SienaSystemDate DateItem

const (
	DATEFORMATPICK     = "20060102T150405"
	DATEFORMATSIENA    = "2006-01-02"
	DATEFORMATYMD      = "20060102"
	DATEFORMATUSER     = "Monday, 02 Jan 2006"
	SIENACPTYSEP       = "\u22EE"
	APPCONFIG          = "properties.cfg"
	SQLCONFIG          = "mssql.cfg"
	SIENACONFIG        = "siena.cfg"
	DATASTORECONFIG    = "datastore.cfg"
	DATETIMEFORMATUSER = DATEFORMATUSER + " @ 15:04:05"
	TIMEHMS            = "15:04:05"
)

//SienaBusinessDateItem is cheese
type DateItem struct {
	Today     string
	Internal  time.Time
	Siena     string
	YYYYMMDD  string
	PICKEpoch string
}
