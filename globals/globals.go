package globals

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jimlawless/cfg"
)

var SessionToken = ""
var UUID = "authorAdjust"
var SecurityViolation = ""
var DB *sql.DB
var UserRole = "/default"
var UserName = ""
var UserKnowAs = ""
var UserNavi = ""

var ApplicationProperties map[string]string
var ApplicationPropertiesDB map[string]string
var ApplicationDB *sql.DB

var SienaProperties map[string]string
var SienaPropertiesDB map[string]string
var SienaDB *sql.DB
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
	ColorReset         = "\033[0m"
	ColorRed           = "\033[31m"
	ColorGreen         = "\033[32m"
	ColorYellow        = "\033[33m"
	ColorBlue          = "\033[34m"
	ColorPurple        = "\033[35m"
	ColorCyan          = "\033[36m"
	ColorWhite         = "\033[37m"
)

//SienaBusinessDateItem is cheese
type DateItem struct {
	Today     string
	Internal  time.Time
	Siena     string
	YYYYMMDD  string
	PICKEpoch string
}

func Initialise() {
	SessionToken = ""
	UUID = "authorAdjust"
	SecurityViolation = ""
	//DB *sql.DB
	//Datasource_db *sql.DB
	UserRole = "/default"
	UserName = ""
	UserKnowAs = ""
	UserNavi = ""

	//SienaSystemDate DateItem
	ApplicationProperties = getProperties(APPCONFIG)
	SienaProperties = getProperties(SIENACONFIG)
	SienaPropertiesDB = getProperties(SQLCONFIG)
	ApplicationPropertiesDB = getProperties(DATASTORECONFIG)

	//var SienaPropertiesDB map[string]string
	//var ApplicationPropertiesDB map[string]string
	//var SienaProperties map[string]string

	ApplicationDB, _ = globalsDatabaseConnect(ApplicationPropertiesDB)
	SienaDB, _ = globalsDatabaseConnect(SienaPropertiesDB)

}

// DataStoreConnect connects application to its datastore database
func globalsDatabaseConnect(mssqlConfig map[string]string) (*sql.DB, error) {
	// Connect to SQL Server DB
	//mssqlConfig := getProperties(config)

	server := mssqlConfig["server"]
	user := mssqlConfig["user"]
	password := mssqlConfig["password"]
	port := mssqlConfig["port"]
	database := mssqlConfig["database"]

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		server, user, password, port, database)
	dbInstance, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	stmt, err := dbInstance.Prepare("select @@version")
	row := stmt.QueryRow()
	var result string

	err = row.Scan(&result)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	//fmt.Printf("%s\n", result)
	return dbInstance, err
}

//comment
func getProperties(inPropertiesFile string) map[string]string {
	wctProperties := make(map[string]string)
	machineName, _ := os.Hostname()
	propertiesFileName := "config/" + inPropertiesFile
	localisedFileName := "config/" + machineName + "/" + inPropertiesFile
	if fileExists(localisedFileName) {
		propertiesFileName = localisedFileName
	}
	err := cfg.Load(propertiesFileName, wctProperties)

	if err != nil {
		log.Fatal(err)
	}
	return wctProperties
}

// FileExists returns true if the specified file existing on the filesystem
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
