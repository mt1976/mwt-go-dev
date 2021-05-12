package globals

import (
	"database/sql"
	"fmt"
	"io/ioutil"
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

var ApplicationProperties map[string]string
var ApplicationPropertiesDB map[string]string
var ApplicationDB *sql.DB
var InstanceProperties map[string]string

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
	DFNANO                   = "2006-01-02T15:04:05.999999"
	DFDD                     = "02"
	DFMM                     = "01"
	DFYY                     = "06"
	DFYYYY                   = "2006"
	DFmm                     = "04"
	DFss                     = "05"
	DFhh                     = "15"
	DATETIMEFORMATSQLSERVER  = "2006-01-02 15:04:05"
	SIENACPTYSEP             = "\u22EE"
	APPCONFIG                = "application.cfg"
	SQLCONFIG                = "sienaDB.cfg"
	SIENACONFIG              = "siena.cfg"
	DATASTORECONFIG          = "applicationDB.cfg"
	INSTANCECONFIG           = "instance.cfg"
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
	log.Println("Vrooom....")
	SessionToken = ""
	UUID = "authorAdjust"
	SecurityViolation = ""

	PreInitialise()

	//SienaSystemDate DateItem
	ApplicationProperties = getProperties(APPCONFIG)
	SienaProperties = getProperties(SIENACONFIG)
	SienaPropertiesDB = getProperties(SQLCONFIG)
	ApplicationPropertiesDB = getProperties(DATASTORECONFIG)
	InstanceProperties = getProperties(INSTANCECONFIG)
	//
	log.Println("Connecting to DB's")
	ApplicationDB, _ = GlobalsDatabaseConnect(ApplicationPropertiesDB)
	SienaDB, _ = GlobalsDatabaseConnect(SienaPropertiesDB)
	//

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
func GlobalsDatabaseConnect(mssqlConfig map[string]string) (*sql.DB, error) {
	// Connect to SQL Server DB
	//mssqlConfig := getProperties(config)

	server := mssqlConfig["server"]
	user := mssqlConfig["user"]
	password := mssqlConfig["password"]
	port := mssqlConfig["port"]
	database := mssqlConfig["database"]
	instance := mssqlConfig["instance"]
	log.Println("Connecting to " + database)

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		server, user, password, port, database)
	dbInstance, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	keepalive, _ := time.ParseDuration("-1h")
	dbInstance.SetConnMaxLifetime(keepalive)
	stmt, err := dbInstance.Prepare("select @@version")
	row := stmt.QueryRow()
	var result string

	err = row.Scan(&result)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}

	dbName := database
	if len(instance) != 0 {
		dbName = database + "-" + instance
	}

	checkDBstmt := "SELECT create_date FROM sys.databases WHERE name = '" + dbName + "'"

	stmt2, err2 := dbInstance.Prepare(checkDBstmt)
	log.Println(checkDBstmt)
	//spew.Dump(stmt2)
	dbCheck := stmt2.QueryRow()
	var result2 string

	err2 = dbCheck.Scan(&result2)
	if err2 != nil {
		log.Println("Database does not exist", "Generating", dbName)
		CreateDatabase(dbInstance, mssqlConfig, dbName)
	} else {
		log.Println("Database Exists " + result2)
	}

	//fmt.Printf("%s\n", result)
	return dbInstance, err
}

func CreateDatabase(dbInstance *sql.DB, mssqlConfig map[string]string, dbName string) {
	log.Println("poo")

	createDBSQL := "CREATE DATABASE [" + dbName + "]"

	_, errCreateDBSQL := dbInstance.Exec(createDBSQL)
	if errCreateDBSQL != nil {
		log.Panic(errCreateDBSQL)
	}

	log.Println("Database Created" + dbName)

}

func GlobalsDatabasePoke(dbInstance *sql.DB, mssqlConfig map[string]string) *sql.DB {
	log.Printf("Fingering     : Server '%s' Database '%s' Schema '%s'", mssqlConfig["server"], mssqlConfig["database"], mssqlConfig["schema"])
	err := dbInstance.Ping()
	if err != nil {
		log.Printf("Reconnecting  : Server '%s' Database '%s' Schema '%s' (%s)", mssqlConfig["server"], mssqlConfig["database"], mssqlConfig["schema"], err.Error())
		// Try to reconnect
		dbInstance, err = GlobalsDatabaseConnect(mssqlConfig)
		if err != nil {
			log.Println(err.Error())
		}
	}
	return dbInstance
}

// Load a Properties File
func getProperties(inPropertiesFile string) map[string]string {
	wctProperties := make(map[string]string)
	//machineName, _ := os.Hostname()
	// TODO: Dockerise
	// For docker - if can't find properties file (create one from the template properties file)
	propertiesFileName := "config/" + inPropertiesFile
	if fileExists(propertiesFileName) {
		// Do nothign this is ok
	} else {

		ok := copyDataFile(inPropertiesFile, "config/", "config/fileSystem/config")
		if !ok {
			log.Println("Issue in Copy Function")
		}
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

func PreInitialise() {

}

func readDataFile(fileName string, path string) (string, error) {
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName
	if len(path) != 0 {
		filePath = pwd + path + "/" + fileName
	}

	// Check it exists - If not create it
	if !(fileExists(filePath)) {
		writeDataFile(fileName, path, "")
	}

	//log.Println("Read          :", filePath)
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// Convert []byte to string and print to screen
	return string(content), err
}

func writeDataFile(fileName string, path string, content string) (bool, error) {
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName
	if len(path) != 0 {
		filePath = pwd + path + "/" + fileName
	}
	//log.Println("Write         :", filePath)

	message := []byte(content)
	err := ioutil.WriteFile(filePath, message, 0644)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	return false, nil
}

func deleteDataFile(fileName string, path string) int {
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName
	if len(path) != 0 {
		filePath = pwd + path + "/" + fileName
	}
	//log.Println("Delete        :", filePath)

	// delete file

	if fileExists(filePath) {
		var err = os.Remove(filePath)
		if err != nil {
			log.Fatal(err.Error())
			return -1
		}
	}
	fmt.Println("File Deleted - " + fileName + " - " + path)
	return 1
}

func copyDataFile(fileName string, fromPath string, toPath string) bool {

	log.Panicln("Copying " + fileName + " from " + fromPath + " to " + toPath)

	content, err := readDataFile(fileName, fromPath)
	if err != nil {
		log.Panicln(err.Error())
	}

	ok, err2 := writeDataFile(fileName, toPath, content)
	if err2 != nil {
		log.Panicln(err2.Error())
	}

	if !ok {
		log.Panicln("Unable to Copy " + fileName + " from " + fromPath + " to " + toPath)
	}

	return true
}
