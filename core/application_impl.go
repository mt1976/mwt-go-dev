package core

import (
	"fmt"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/davecgh/go-spew/spew"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func Uptime() time.Duration {
	return time.Since(startTime)
}

func Log_uptime() {
	logs.System("App Uptime : " + Uptime().String())
}

func Initialise() {
	//logs.Message("Initialisation", "Vroom")
	logs.Information("Initialisation", "Vroom "+Bike)

	SessionToken = ""
	UUID = "authorAdjust"
	SecurityViolation = ""

	SystemHostname, _ = os.Hostname()

	PWD, _ = os.Getwd()

	PreInitialise()

	//SienaSystemDate DateItem
	ApplicationProperties = getProperties(APPCONFIG)
	SienaProperties = getProperties(SIENACONFIG)
	SienaPropertiesDB = getProperties(SQLCONFIG)
	ApplicationPropertiesDB = getProperties(DATASTORECONFIG)
	InstanceProperties = getProperties(INSTANCECONFIG)
	ApplicationStubLists = getProperties(APPSTUBLISTS)
	SienaTablesLookup = getProperties("sienaTables.cfg")
	//MasterPropertiesDB = getProperties(DATASTORECONFIG)
	//MasterPropertiesDB["database"] = "master"

	IsChildInstance = false
	if len(ApplicationPropertiesDB["instance"]) != 0 {
		IsChildInstance = true
	}
	//

	logs.Information("Connecting to application databases...", "")

	//MasterDB, _ = GlobalsDatabaseConnect(MasterPropertiesDB)

	ApplicationDB, _ = Database_Connect(ApplicationPropertiesDB)
	//log.Printf("Initialisation", ApplicationDB.Stats())

	SienaDB, _ = Database_Connect(SienaPropertiesDB)
	//log.Printf("Initialisation", SienaDB.Stats())
	logs.Success("Connections established")

	//

	//spew.Dump(ApplicationDB)
	// TODO: get a list of additional DB's to connect to (from the SRS.sienadbStore)
	// TODO: load them into the var sourceAccess []*sql.DB slice

	SessionManager = scs.New()
	fmt.Printf("ApplicationSessionLife(): %v\n", ApplicationSessionLife())
	life, err := time.ParseDuration(ApplicationSessionLife())
	if err != nil {
		logs.Fatal("No Session Life Found", err)
	}
	SessionManager.Lifetime = life
	SessionManager.IdleTimeout = 20 * time.Minute
	SessionManager.Cookie.Name = GetCookieIdentity()
	SessionManager.Cookie.HttpOnly = true
	SessionManager.Cookie.Persist = true
	//SessionManager.Cookie.SameSite = http.SameSiteStrictMode
	SessionManager.Cookie.Secure = false
	spew.Dump(SessionManager)
	fmt.Printf("SessionManager: %v\n", SessionManager)
	logs.Information("Initialisation", "Vroooom Vrooooom! "+Bike+Bike)
}
