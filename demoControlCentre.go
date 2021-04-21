package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/google/uuid"

	"github.com/mbndr/figlet4go"
)

//CONST_CONFIG_FILE is cheese
const (
	APPCONFIG       = "properties.cfg"
	SQLCONFIG       = "mssql.cfg"
	SIENACONFIG     = "siena.cfg"
	DATEFORMATPICK  = "20060102T150405"
	DATEFORMATSIENA = "2006-01-02"
	DATEFORMATYMD   = "20060102"
	DATEFORMATUSER  = "Monday, 02 Jan 2006"
	SIENACPTYSEP    = "\u22EE"
)

var gSessionToken = ""
var gUUID = "authorAdjust"
var gSecurityViolation = ""
var gDB *sql.DB
var gUserRole = "/default"
var gUserName = ""
var gUserKnowAs = ""
var gUserNavi = ""
var gSienaSystemDate sienaBusinessDateItem

func main() {

	ascii := figlet4go.NewAsciiRender()

	wctProperties := getProperties(APPCONFIG)

	// The underscore would be an error
	renderStr, _ := ascii.Render(wctProperties["appname"])
	tmpHostname, _ := os.Hostname()

	fmt.Print(renderStr)
	log.Println("INSTANCE")
	log.Println("Machine    :", tmpHostname)
	log.Println("PATHS")
	log.Println("Delivery   :", wctProperties["deliverpath"])
	log.Println("Response   :", wctProperties["receivepath"])
	log.Println("Processed  :", wctProperties["processedpath"])
	log.Println("")
	log.Println("APPLICATION")
	log.Println("Name       :", wctProperties["appname"])
	log.Println("Msg Format :", wctProperties["responseformat"])
	log.Println("Licence    :", wctProperties["licname"])
	log.Println("Lic URL    :", wctProperties["liclink"])
	log.Println("")
	log.Println("RELEASE")
	log.Println("Release ID :", wctProperties["releaseid"])
	log.Println("Level      :", wctProperties["releaselevel"])
	log.Println("Number     :", wctProperties["releasenumber"])
	log.Println("")

	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/login", valLoginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/home", homePageHandler)
	http.HandleFunc("/srvServiceCatalog", srvServiceCatalogHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/favicon-32x32.png", favicon32Handler)
	http.HandleFunc("/site.webmanifest", faviconManifestHandler)
	http.HandleFunc("/favicon-16x16.png", favicon16Handler)
	http.HandleFunc("/browserconfig.xml", faviconBrowserConfigHandler)
	http.HandleFunc("/listResponses/", listResponsesHandler)
	http.HandleFunc("/previewRequest/", previewRequestHandler)
	http.HandleFunc("/executeRequest/", executeRequestHandler)
	http.HandleFunc("/viewResponse/", viewResponseHandler)
	http.HandleFunc("/deleteResponse/", deleteResponseHandler)
	http.HandleFunc("/clearQueues/", clearQueuesHandler)
	http.HandleFunc("/viewSrvEnvironment/", viewSrvEnvironmentHandler)
	http.HandleFunc("/listSrvConfiguration/", listSrvConfigurationHandler)
	http.HandleFunc("/viewSrvConfiguration/", viewSrvConfigurationHandler)
	http.HandleFunc("/editSrvConfiguration/", editSrvConfigurationHandler)
	http.HandleFunc("/saveSrvConfiguration/", saveSrvConfigurationHandler)
	http.HandleFunc("/viewAppConfiguration/", viewAppConfigurationHandler)
	http.HandleFunc("/listSvcDataMap/", listSvcDataMapHandler)
	http.HandleFunc("/viewSvcDataMap/", viewSvcDataMapHandler)
	http.HandleFunc("/editSvcDataMap/", editSvcDataMapHandler)
	http.HandleFunc("/saveSvcDataMap/", saveSvcDataMapHandler)
	http.HandleFunc("/viewSvcDataMapXML/", viewSvcDataMapXMLHandler)
	http.HandleFunc("/editSvcDataMapXML/", editSvcDataMapXMLHandler)
	http.HandleFunc("/saveSvcDataMapXML/", saveSvcDataMapXMLHandler)
	http.HandleFunc("/newSvcDataMap/", newSvcDataMapHandler)
	http.HandleFunc("/genSvcDataMap/", genSvcDataMapHandler)
	http.HandleFunc("/deleteSvcDataMap/", deleteSvcDataMapHandler)
	http.HandleFunc("/listSienaCountry/", listSienaCountryHandler)
	http.HandleFunc("/viewSienaCountry/", viewSienaCountryHandler)
	http.HandleFunc("/editSienaCountry/", editSienaCountryHandler)
	http.HandleFunc("/saveSienaCountry/", saveSienaCountryHandler)
	http.HandleFunc("/newSienaCountry/", newSienaCountryHandler)

	http.HandleFunc("/listSienaSector/", listSienaSectorHandler)
	http.HandleFunc("/viewSienaSector/", viewSienaSectorHandler)
	http.HandleFunc("/editSienaSector/", editSienaSectorHandler)
	http.HandleFunc("/saveSienaSector/", saveSienaSectorHandler)
	http.HandleFunc("/newSienaSector/", newSienaSectorHandler)

	http.HandleFunc("/listSienaFirm/", listSienaFirmHandler)
	http.HandleFunc("/viewSienaFirm/", viewSienaFirmHandler)
	http.HandleFunc("/editSienaFirm/", editSienaFirmHandler)
	http.HandleFunc("/saveSienaFirm/", saveSienaFirmHandler)
	http.HandleFunc("/newSienaFirm/", newSienaFirmHandler)

	http.HandleFunc("/listSienaPortfolio/", listSienaPortfolioHandler)
	http.HandleFunc("/viewSienaPortfolio/", viewSienaPortfolioHandler)
	http.HandleFunc("/editSienaPortfolio/", editSienaPortfolioHandler)
	http.HandleFunc("/saveSienaPortfolio/", saveSienaPortfolioHandler)
	http.HandleFunc("/newSienaPortfolio/", newSienaPortfolioHandler)

	http.HandleFunc("/listSienaCentre/", listSienaCentreHandler)
	http.HandleFunc("/viewSienaCentre/", viewSienaCentreHandler)
	http.HandleFunc("/editSienaCentre/", editSienaCentreHandler)
	http.HandleFunc("/saveSienaCentre/", saveSienaCentreHandler)
	http.HandleFunc("/newSienaCentre/", newSienaCentreHandler)

	http.HandleFunc("/listSienaBook/", listSienaBookHandler)
	http.HandleFunc("/viewSienaBook/", viewSienaBookHandler)
	http.HandleFunc("/editSienaBook/", editSienaBookHandler)
	http.HandleFunc("/saveSienaBook/", saveSienaBookHandler)
	http.HandleFunc("/newSienaBook/", newSienaBookHandler)

	http.HandleFunc("/listSienaBroker/", listSienaBrokerHandler)
	http.HandleFunc("/viewSienaBroker/", viewSienaBrokerHandler)
	http.HandleFunc("/editSienaBroker/", editSienaBrokerHandler)
	http.HandleFunc("/saveSienaBroker/", saveSienaBrokerHandler)
	http.HandleFunc("/newSienaBroker/", newSienaBrokerHandler)

	http.HandleFunc("/listSienaAccount/", listSienaAccountHandler)
	http.HandleFunc("/viewSienaAccount/", viewSienaAccountHandler)

	http.HandleFunc("/listSienaCurrency/", listSienaCurrencyHandler)
	http.HandleFunc("/listSienaCurrencyPair/", listSienaCurrencyPairHandler)

	http.HandleFunc("/listSienaMandatedUser/", listSienaMandatedUserHandler)
	http.HandleFunc("/viewSienaMandatedUser/", viewSienaMandatedUserHandler)
	http.HandleFunc("/editSienaMandatedUser/", editSienaMandatedUserHandler)
	http.HandleFunc("/saveSienaMandatedUser/", saveSienaMandatedUserHandler)
	http.HandleFunc("/newSienaMandatedUser/", newSienaMandatedUserHandler)

	http.HandleFunc("/listSienaCounterparty/", listSienaCounterpartyHandler)
	http.HandleFunc("/viewSienaCounterparty/", viewSienaCounterpartyHandler)
	http.HandleFunc("/editSienaCounterparty/", editSienaCounterpartyHandler)
	http.HandleFunc("/saveSienaCounterparty/", saveSienaCounterpartyHandler)
	http.HandleFunc("/newSienaCounterparty/", newSienaCounterpartyHandler)

	http.HandleFunc("/editSienaCounterpartyAddress/", editSienaCounterpartyAddressHandler)
	http.HandleFunc("/saveSienaCounterpartyAddress/", saveSienaCounterpartyAddressHandler)
	http.HandleFunc("/editSienaCounterpartyExtensions/", editSienaCounterpartyExtensionsHandler)
	http.HandleFunc("/saveSienaCounterpartyExtensions/", saveSienaCounterpartyExtensionsHandler)

	http.HandleFunc("/listSienaCounterpartyPayee/", listSienaCounterpartyPayeeHandler)
	http.HandleFunc("/viewSienaCounterpartyPayee/", viewSienaCounterpartyPayeeHandler)
	http.HandleFunc("/editSienaCounterpartyPayee/", editSienaCounterpartyPayeeHandler)
	http.HandleFunc("/saveSienaCounterpartyPayee/", saveSienaCounterpartyPayeeHandler)
	http.HandleFunc("/newSienaCounterpartyPayee/", newSienaCounterpartyPayeeHandler)

	http.HandleFunc("/listSienaCounterpartyImportID/", listSienaCounterpartyImportIDHandler)
	http.HandleFunc("/viewSienaCounterpartyImportID/", viewSienaCounterpartyImportIDHandler)
	http.HandleFunc("/editSienaCounterpartyImportID/", editSienaCounterpartyImportIDHandler)
	http.HandleFunc("/saveSienaCounterpartyImportID/", saveSienaCounterpartyImportIDHandler)
	http.HandleFunc("/newSienaCounterpartyImportID/", newSienaCounterpartyImportIDHandler)

	http.HandleFunc("/listSienaDealList/", listSienaDealListHandler)
	http.HandleFunc("/viewSienaDealList/", viewSienaDealListHandler)
	http.HandleFunc("/listSienaAccountLadder/", listSienaAccountLadderHandler)
	http.HandleFunc("/listSienaAccountTransactions/", listSienaAccountTransactionsHandler)
	//http.HandleFunc("/saveSienaDealList/", saveSienaDealListHandler)
	//http.HandleFunc("/newSienaDealList/", newSienaDealListHandler)

	http.HandleFunc("/listSienaCounterpartyGroup/", listSienaCounterpartyGroupHandler)
	http.HandleFunc("/viewSienaCounterpartyGroup/", viewSienaCounterpartyGroupHandler)
	http.HandleFunc("/editSienaCounterpartyGroup/", editSienaCounterpartyGroupHandler)
	http.HandleFunc("/saveSienaCounterpartyGroup/", saveSienaCounterpartyGroupHandler)
	http.HandleFunc("/newSienaCounterpartyGroup/", newSienaCounterpartyGroupHandler)

	http.HandleFunc("/dashboard/", sienaDashboardHandler)

	http.HandleFunc("/shutdown/", shutdownHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	db, _ := sienaConnect()
	_, gSienaSystemDate, _ = getSienaBusinessDate(db)
	log.Println("Siena System Date:", gSienaSystemDate.Internal.Format(DATEFORMATUSER))
	// Get menu
	menuCount, _ := fetchappMenuData("")
	log.Println("No. Menu Items   :", menuCount)

	log.Println("")
	log.Println("READY STEADY GO!!!")

	log.Println("URL        :", "http://localhost:"+wctProperties["port"])

	httpPort := ":" + wctProperties["port"]
	http.ListenAndServe(httpPort, nil)

}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/favicon.ico")
}

func favicon16Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/favicon-16x16.png")
}

func favicon32Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/favicon-32x32.png")
}

func faviconManifestHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "site.webmanifest")
}

func faviconBrowserConfigHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "browserconfig.xml")
}

func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	wctProperties := getProperties(APPCONFIG)

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	requestID := uuid.New()

	log.Println("Servicing :", inUTL)

	requestMessage := buildRequestMessage(requestID.String(), "SHUTDOWN", "", "", "", wctProperties)

	fmt.Println("requestMessage", requestMessage)
	fmt.Println("SEND MESSAGE")
	sendRequest(requestMessage, requestID.String(), wctProperties)
	m := http.NewServeMux()

	s := http.Server{Addr: wctProperties["port"], Handler: m}
	s.Shutdown(context.Background())
	//	r.URL.Path = "/viewResponse?uuid=" + requestID.String()
	//	viewResponseHandler(w, r)

}
func clearQueuesHandler(w http.ResponseWriter, r *http.Request) {

	//var propertiesFileName = "config/properties.cfg"
	wctProperties := getProperties(APPCONFIG)
	//	tmpl := "viewResponse"
	inUTL := r.URL.Path
	//requestID := uuid.New()

	log.Println("Servicing :", inUTL)

	//fmt.Println("delivPath", wctProperties["deliverpath"])
	err1 := RemoveContents(wctProperties["deliverpath"])
	if err1 != nil {
		fmt.Println(err1)
	}

	//fmt.Println("recPath", wctProperties["receivepath"])

	err2 := RemoveContents(wctProperties["receivepath"])
	if err2 != nil {
		fmt.Println(err2)
	}
	//fmt.Println("procPath", wctProperties["processedpath"])

	err3 := RemoveContents(wctProperties["processedpath"])
	if err3 != nil {
		fmt.Println(err3)
	}

	homePageHandler(w, r)

}
