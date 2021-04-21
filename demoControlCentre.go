package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/google/uuid"

	"github.com/mbndr/figlet4go"
	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
	siena "github.com/mt1976/mwt-go-dev/siena"
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

func main() {

	ascii := figlet4go.NewAsciiRender()

	wctProperties := application.GetProperties(APPCONFIG)

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

	http.HandleFunc("/", application.LoginHandler)
	http.HandleFunc("/login", application.ValidateLoginHandler)
	http.HandleFunc("/logout", application.LogoutHandler)
	http.HandleFunc("/home", application.HomePageHandler)
	http.HandleFunc("/srvServiceCatalog", application.ServiceCatalogHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/favicon-32x32.png", favicon32Handler)
	http.HandleFunc("/site.webmanifest", faviconManifestHandler)
	http.HandleFunc("/favicon-16x16.png", favicon16Handler)
	http.HandleFunc("/browserconfig.xml", faviconBrowserConfigHandler)
	http.HandleFunc("/listResponses/", application.ListResponsesHandler)
	http.HandleFunc("/previewRequest/", application.PreviewRequestHandler)
	http.HandleFunc("/executeRequest/", application.ExecuteRequestHandler)
	http.HandleFunc("/viewResponse/", application.ViewResponseHandler)
	http.HandleFunc("/deleteResponse/", application.DeleteResponseHandler)
	http.HandleFunc("/clearQueues/", clearQueuesHandler)
	http.HandleFunc("/viewSrvEnvironment/", application.ViewSrvEnvironmentHandler)
	http.HandleFunc("/listSrvConfiguration/", application.ListSrvConfigurationHandler)
	http.HandleFunc("/viewSrvConfiguration/", application.ViewSrvConfigurationHandler)
	http.HandleFunc("/editSrvConfiguration/", application.EditSrvConfigurationHandler)
	http.HandleFunc("/saveSrvConfiguration/", application.SaveSrvConfigurationHandler)
	http.HandleFunc("/viewAppConfiguration/", application.ViewAppConfigurationHandler)
	http.HandleFunc("/listSvcDataMap/", siena.ListSvcDataMapHandler)
	http.HandleFunc("/viewSvcDataMap/", siena.ViewSvcDataMapHandler)
	http.HandleFunc("/editSvcDataMap/", siena.EditSvcDataMapHandler)
	http.HandleFunc("/saveSvcDataMap/", siena.SaveSvcDataMapHandler)
	http.HandleFunc("/viewSvcDataMapXML/", siena.ViewSvcDataMapXMLHandler)
	http.HandleFunc("/editSvcDataMapXML/", siena.EditSvcDataMapXMLHandler)
	http.HandleFunc("/saveSvcDataMapXML/", siena.SaveSvcDataMapXMLHandler)
	http.HandleFunc("/newSvcDataMap/", siena.NewSvcDataMapHandler)
	http.HandleFunc("/genSvcDataMap/", siena.GenSvcDataMapHandler)
	http.HandleFunc("/deleteSvcDataMap/", siena.DeleteSvcDataMapHandler)
	http.HandleFunc("/listSienaCountry/", siena.ListSienaCountryHandler)
	http.HandleFunc("/viewSienaCountry/", siena.ViewSienaCountryHandler)
	http.HandleFunc("/editSienaCountry/", siena.EditSienaCountryHandler)
	http.HandleFunc("/saveSienaCountry/", siena.SaveSienaCountryHandler)
	http.HandleFunc("/newSienaCountry/", siena.NewSienaCountryHandler)

	http.HandleFunc("/listSienaSector/", siena.ListSienaSectorHandler)
	http.HandleFunc("/viewSienaSector/", siena.ViewSienaSectorHandler)
	http.HandleFunc("/editSienaSector/", siena.EditSienaSectorHandler)
	http.HandleFunc("/saveSienaSector/", siena.SaveSienaSectorHandler)
	http.HandleFunc("/newSienaSector/", siena.NewSienaSectorHandler)

	http.HandleFunc("/listSienaFirm/", siena.ListSienaFirmHandler)
	http.HandleFunc("/viewSienaFirm/", siena.ViewSienaFirmHandler)
	http.HandleFunc("/editSienaFirm/", siena.EditSienaFirmHandler)
	http.HandleFunc("/saveSienaFirm/", siena.SaveSienaFirmHandler)
	http.HandleFunc("/newSienaFirm/", siena.NewSienaFirmHandler)

	http.HandleFunc("/listSienaPortfolio/", siena.ListSienaPortfolioHandler)
	http.HandleFunc("/viewSienaPortfolio/", siena.ViewSienaPortfolioHandler)
	http.HandleFunc("/editSienaPortfolio/", siena.EditSienaPortfolioHandler)
	http.HandleFunc("/saveSienaPortfolio/", siena.SaveSienaPortfolioHandler)
	http.HandleFunc("/newSienaPortfolio/", siena.NewSienaPortfolioHandler)

	http.HandleFunc("/listSienaCentre/", siena.ListSienaCentreHandler)
	http.HandleFunc("/viewSienaCentre/", siena.ViewSienaCentreHandler)
	http.HandleFunc("/editSienaCentre/", siena.EditSienaCentreHandler)
	http.HandleFunc("/saveSienaCentre/", siena.SaveSienaCentreHandler)
	http.HandleFunc("/newSienaCentre/", siena.NewSienaCentreHandler)

	http.HandleFunc("/listSienaBook/", siena.ListSienaBookHandler)
	http.HandleFunc("/viewSienaBook/", siena.ViewSienaBookHandler)
	http.HandleFunc("/editSienaBook/", siena.EditSienaBookHandler)
	http.HandleFunc("/saveSienaBook/", siena.SaveSienaBookHandler)
	http.HandleFunc("/newSienaBook/", siena.NewSienaBookHandler)

	http.HandleFunc("/listSienaBroker/", siena.ListSienaBrokerHandler)
	http.HandleFunc("/viewSienaBroker/", siena.ViewSienaBrokerHandler)
	http.HandleFunc("/editSienaBroker/", siena.EditSienaBrokerHandler)
	http.HandleFunc("/saveSienaBroker/", siena.SaveSienaBrokerHandler)
	http.HandleFunc("/newSienaBroker/", siena.NewSienaBrokerHandler)

	http.HandleFunc("/listSienaAccount/", siena.ListSienaAccountHandler)
	http.HandleFunc("/viewSienaAccount/", siena.ViewSienaAccountHandler)

	http.HandleFunc("/listSienaCurrency/", siena.ListSienaCurrencyHandler)
	http.HandleFunc("/listSienaCurrencyPair/", siena.ListSienaCurrencyPairHandler)

	http.HandleFunc("/listSienaMandatedUser/", siena.ListSienaMandatedUserHandler)
	http.HandleFunc("/viewSienaMandatedUser/", siena.ViewSienaMandatedUserHandler)
	http.HandleFunc("/editSienaMandatedUser/", siena.EditSienaMandatedUserHandler)
	http.HandleFunc("/saveSienaMandatedUser/", siena.SaveSienaMandatedUserHandler)
	http.HandleFunc("/newSienaMandatedUser/", siena.NewSienaMandatedUserHandler)

	http.HandleFunc("/listSienaCounterparty/", siena.ListSienaCounterpartyHandler)
	http.HandleFunc("/viewSienaCounterparty/", siena.ViewSienaCounterpartyHandler)
	http.HandleFunc("/editSienaCounterparty/", siena.EditSienaCounterpartyHandler)
	http.HandleFunc("/saveSienaCounterparty/", siena.SaveSienaCounterpartyHandler)
	http.HandleFunc("/newSienaCounterparty/", siena.NewSienaCounterpartyHandler)

	http.HandleFunc("/editSienaCounterpartyAddress/", siena.EditSienaCounterpartyAddressHandler)
	http.HandleFunc("/saveSienaCounterpartyAddress/", siena.SaveSienaCounterpartyAddressHandler)
	http.HandleFunc("/editSienaCounterpartyExtensions/", siena.EditSienaCounterpartyExtensionsHandler)
	http.HandleFunc("/saveSienaCounterpartyExtensions/", siena.SaveSienaCounterpartyExtensionsHandler)

	http.HandleFunc("/listSienaCounterpartyPayee/", siena.ListSienaCounterpartyPayeeHandler)
	http.HandleFunc("/viewSienaCounterpartyPayee/", siena.ViewSienaCounterpartyPayeeHandler)
	http.HandleFunc("/editSienaCounterpartyPayee/", siena.EditSienaCounterpartyPayeeHandler)
	http.HandleFunc("/saveSienaCounterpartyPayee/", siena.SaveSienaCounterpartyPayeeHandler)
	http.HandleFunc("/newSienaCounterpartyPayee/", siena.NewSienaCounterpartyPayeeHandler)

	http.HandleFunc("/listSienaCounterpartyImportID/", siena.ListSienaCounterpartyImportIDHandler)
	http.HandleFunc("/viewSienaCounterpartyImportID/", siena.ViewSienaCounterpartyImportIDHandler)
	http.HandleFunc("/editSienaCounterpartyImportID/", siena.EditSienaCounterpartyImportIDHandler)
	http.HandleFunc("/saveSienaCounterpartyImportID/", siena.SaveSienaCounterpartyImportIDHandler)
	http.HandleFunc("/newSienaCounterpartyImportID/", siena.NewSienaCounterpartyImportIDHandler)

	http.HandleFunc("/listSienaDealList/", siena.ListSienaDealListHandler)
	http.HandleFunc("/viewSienaDealList/", siena.ViewSienaDealListHandler)
	http.HandleFunc("/listSienaAccountLadder/", siena.ListSienaAccountLadderHandler)
	http.HandleFunc("/listSienaAccountTransactions/", siena.ListSienaAccountTransactionsHandler)
	//http.HandleFunc("/saveSienaDealList/", siena.SaveSienaDealListHandler)
	//http.HandleFunc("/newSienaDealList/", siena.NewSienaDealListHandler)

	http.HandleFunc("/listSienaCounterpartyGroup/", siena.ListSienaCounterpartyGroupHandler)
	http.HandleFunc("/viewSienaCounterpartyGroup/", siena.ViewSienaCounterpartyGroupHandler)
	http.HandleFunc("/editSienaCounterpartyGroup/", siena.EditSienaCounterpartyGroupHandler)
	http.HandleFunc("/saveSienaCounterpartyGroup/", siena.SaveSienaCounterpartyGroupHandler)
	http.HandleFunc("/newSienaCounterpartyGroup/", siena.NewSienaCounterpartyGroupHandler)

	http.HandleFunc("/dashboard/", siena.SienaDashboardHandler)

	http.HandleFunc("/shutdown/", shutdownHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	db, _ := siena.Connect()
	_, tempDate, _ := siena.GetBusinessDate(db)
	globals.SienaSystemDate = tempDate
	log.Println("Siena System Date:", tempDate.Internal.Format(globals.DATEFORMATUSER))
	// Get menu
	menuCount, _ := application.FetchappMenuData("")
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

//// TODO: migrage the following three functions to appsupport
func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	wctProperties := application.GetProperties(APPCONFIG)

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	requestID := uuid.New()

	log.Println("Servicing :", inUTL)

	requestMessage := application.BuildRequestMessage(requestID.String(), "SHUTDOWN", "", "", "", wctProperties)

	fmt.Println("requestMessage", requestMessage)
	fmt.Println("SEND MESSAGE")
	application.SendRequest(requestMessage, requestID.String(), wctProperties)
	m := http.NewServeMux()

	s := http.Server{Addr: wctProperties["port"], Handler: m}
	s.Shutdown(context.Background())
	//	r.URL.Path = "/viewResponse?uuid=" + requestID.String()
	//	viewResponseHandler(w, r)

}
func clearQueuesHandler(w http.ResponseWriter, r *http.Request) {

	//var propertiesFileName = "config/properties.cfg"
	wctProperties := application.GetProperties(APPCONFIG)
	//	tmpl := "viewResponse"
	inUTL := r.URL.Path
	//requestID := uuid.New()
	log.Println("Servicing :", inUTL)
	//fmt.Println("delivPath", wctProperties["deliverpath"])
	err1 := application.RemoveContents(wctProperties["deliverpath"])
	if err1 != nil {
		fmt.Println(err1)
	}
	//fmt.Println("recPath", wctProperties["receivepath"])
	err2 := application.RemoveContents(wctProperties["receivepath"])
	if err2 != nil {
		fmt.Println(err2)
	}
	//fmt.Println("procPath", wctProperties["processedpath"])
	err3 := application.RemoveContents(wctProperties["processedpath"])
	if err3 != nil {
		fmt.Println(err3)
	}
	application.HomePageHandler(w, r)
}
func clearResponsesHandler(w http.ResponseWriter, r *http.Request) {
	//var propertiesFileName = "config/properties.cfg"
	wctProperties := application.GetProperties(globals.APPCONFIG)
	//	tmpl := "viewResponse"
	inUTL := r.URL.Path
	//requestID := uuid.New()
	log.Println("Servicing :", inUTL)
	application.RemoveContents(wctProperties["receivepath"])
	application.HomePageHandler(w, r)
}
