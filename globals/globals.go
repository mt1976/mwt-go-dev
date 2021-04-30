package globals

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jimlawless/cfg"
)

var SessionToken = ""
var UUID = "authorAdjust"
var SecurityViolation = ""
var DB *sql.DB
var sourceAccess []*sql.DB

//var UserRole = "/default"
//var UserName = ""
//var UserKnowAs = ""
//var UserNavi = ""

var ApplicationProperties map[string]string
var ApplicationPropertiesDB map[string]string
var ApplicationDB *sql.DB

var SienaProperties map[string]string
var SienaPropertiesDB map[string]string
var SienaDB *sql.DB
var SienaSystemDate DateItem

var SessionManager *scs.SessionManager

const (
	DATEFORMATPICK           = "20060102T150405"
	DATEFORMATSIENA          = "2006-01-02"
	DATEFORMATYMD            = "20060102"
	DATEFORMATUSER           = "02/01/2006"
	DATETIMEFORMATSQLSERVER  = "2006-01-02 15:04:05"
	SIENACPTYSEP             = "\u22EE"
	APPCONFIG                = "properties.cfg"
	SQLCONFIG                = "mssql.cfg"
	SIENACONFIG              = "siena.cfg"
	DATASTORECONFIG          = "datastore.cfg"
	DATETIMEFORMATUSER       = DATEFORMATUSER + " 15:04:05"
	TIMEHMS                  = "15:04:05"
	ColorReset               = "\033[0m"
	ColorRed                 = "\033[31m"
	ColorGreen               = "\033[32m"
	ColorYellow              = "\033[33m"
	ColorBlue                = "\033[34m"
	ColorPurple              = "\033[35m"
	ColorCyan                = "\033[36m"
	ColorWhite               = "\033[37m"
	SessionRole              = "1891835972"
	SessionNavi              = "6782444386"
	SessionKnowAs            = "0627218437"
	SessionUserName          = "9214790474"
	SessionTokenID           = "1516099114"
	SessionUUID              = "0663644127"
	SessionSecurityViolation = "4097340829"
	SessionAppToken          = "1117429826"
	IDSep                    = "|"
	Tick                     = "\u2713"
)

type Connection struct {
	count int
	Pool  []ConnectionItem
}

type ConnectionItem struct {
	ID               string
	Name             string
	DealimportIn     string
	DealimportOut    string
	StaticIn         string
	StaticOut        string
	Database         *sql.DB
	ConnectionString sienaDBItem
}

type sienaDBItem struct {
	ID         string
	Server     string
	Port       string
	User       string
	Password   string
	Database   string
	Schema     string
	Active     string
	SYSCreated string
	SYSWho     string
	SYSHost    string
	SYSUpdated string
}

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
	//UserRole = "/default"
	//UserName = ""
	//UserKnowAs = ""
	//UserNavi = ""

	//SienaSystemDate DateItem
	ApplicationProperties = getProperties(APPCONFIG)
	SienaProperties = getProperties(SIENACONFIG)
	SienaPropertiesDB = getProperties(SQLCONFIG)
	ApplicationPropertiesDB = getProperties(DATASTORECONFIG)

	ApplicationDB, _ = globalsDatabaseConnect(ApplicationPropertiesDB)
	SienaDB, _ = globalsDatabaseConnect(SienaPropertiesDB)

	// TODO: get a list of additional DB's to connect to (from the SRS.sienadbStore)
	// TODO: load them into the var sourceAccess []*sql.DB slice

	SessionManager = scs.New()
	life, err := time.ParseDuration(ApplicationProperties["sessionlife"])
	if err != nil {
		log.Fatal("No Session Life Found", err, life)
	}
	SessionManager.Lifetime = life
	SessionManager.IdleTimeout = 20 * time.Minute
	SessionManager.Cookie.Name = ApplicationProperties["releaseid"]
	SessionManager.Cookie.HttpOnly = true
	SessionManager.Cookie.Persist = true
	//SessionManager.Cookie.SameSite = http.SameSiteStrictMode
	SessionManager.Cookie.Secure = false

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
