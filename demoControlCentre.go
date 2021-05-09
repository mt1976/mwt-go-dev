package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"

	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
	scheduler "github.com/mt1976/mwt-go-dev/jobs"
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

	//ascii := figlet4go.NewAsciiRender()

	//wctProperties := application.GetProperties(APPCONFIG)

	// The underscore would be an error
	//	renderStr, _ := ascii.Render(wctProperties["appname"])
	tmpHostname, _ := os.Hostname()
	line := strings.Repeat("-", 100)
	log.Println(line)

	header("Initialising ...")
	globals.Initialise()
	done("Initialised")
	//	log.Println(line)
	header("Caching ...")
	//application.InitialiseCache()
	done("Cache Refreshed")
	//log.Println(line)

	header("Scheduling Jobs")
	scheduler.Start()
	done("Jobs Scheduled")

	header("Starting Handlers")
	mux := http.NewServeMux()
	mux.HandleFunc("/put", putHandler)
	mux.HandleFunc("/get", getHandler)

	mux.HandleFunc("/", application.LoginHandler)
	mux.HandleFunc("/login", application.ValidateLoginHandler)
	mux.HandleFunc("/logout", application.LogoutHandler)
	mux.HandleFunc("/home", application.HomePageHandler)
	mux.HandleFunc("/srvServiceCatalog", application.ServiceCatalogHandler)
	mux.HandleFunc("/favicon.ico", application.FaviconHandler)
	mux.HandleFunc("/favicon-32x32.png", application.Favicon32Handler)
	mux.HandleFunc("/site.webmanifest", application.FaviconManifestHandler)
	mux.HandleFunc("/favicon-16x16.png", application.Favicon16Handler)
	mux.HandleFunc("/browserconfig.xml", application.FaviconBrowserConfigHandler)
	mux.HandleFunc("/listResponses/", application.ListResponsesHandler)
	mux.HandleFunc("/previewRequest/", application.PreviewRequestHandler)
	mux.HandleFunc("/executeRequest/", application.ExecuteRequestHandler)
	mux.HandleFunc("/viewResponse/", application.ViewResponseHandler)
	mux.HandleFunc("/deleteResponse/", application.DeleteResponseHandler)
	mux.HandleFunc("/clearQueues/", clearQueuesHandler)
	mux.HandleFunc("/viewSrvEnvironment/", application.ViewSrvEnvironmentHandler)
	mux.HandleFunc("/listSrvConfiguration/", application.ListSrvConfigurationHandler)
	mux.HandleFunc("/viewSrvConfiguration/", application.ViewSrvConfigurationHandler)
	mux.HandleFunc("/editSrvConfiguration/", application.EditSrvConfigurationHandler)
	mux.HandleFunc("/saveSrvConfiguration/", application.SaveSrvConfigurationHandler)
	mux.HandleFunc("/viewAppConfiguration/", application.ViewAppConfigurationHandler)

	mux.HandleFunc("/listSvcDataMap/", siena.ListSvcDataMapHandler)
	mux.HandleFunc("/viewSvcDataMap/", siena.ViewSvcDataMapHandler)
	mux.HandleFunc("/editSvcDataMap/", siena.EditSvcDataMapHandler)
	mux.HandleFunc("/saveSvcDataMap/", siena.SaveSvcDataMapHandler)
	mux.HandleFunc("/viewSvcDataMapXML/", siena.ViewSvcDataMapXMLHandler)
	mux.HandleFunc("/editSvcDataMapXML/", siena.EditSvcDataMapXMLHandler)
	mux.HandleFunc("/saveSvcDataMapXML/", siena.SaveSvcDataMapXMLHandler)
	mux.HandleFunc("/newSvcDataMap/", siena.NewSvcDataMapHandler)
	mux.HandleFunc("/genSvcDataMap/", siena.GenSvcDataMapHandler)
	mux.HandleFunc("/deleteSvcDataMap/", siena.DeleteSvcDataMapHandler)
	mux.HandleFunc("/editDataMapColumns/", application.ListLoaderMapStoreHandler)
	mux.HandleFunc("/editLoaderMapStore/", application.EditLoaderMapStoreHandler)
	mux.HandleFunc("/saveLoaderMapStore/", application.SaveLoaderMapStoreHandler)
	mux.HandleFunc("/newLoaderMapStore/", application.NewLoaderMapStoreHandler)
	mux.HandleFunc("/runLoader/", siena.RunDataLoaderHandler)

	mux.HandleFunc("/listSienaCountry/", siena.ListSienaCountryHandler)
	mux.HandleFunc("/viewSienaCountry/", siena.ViewSienaCountryHandler)
	mux.HandleFunc("/editSienaCountry/", siena.EditSienaCountryHandler)
	mux.HandleFunc("/saveSienaCountry/", siena.SaveSienaCountryHandler)
	mux.HandleFunc("/newSienaCountry/", siena.NewSienaCountryHandler)

	mux.HandleFunc("/listSienaSector/", siena.ListSienaSectorHandler)
	mux.HandleFunc("/viewSienaSector/", siena.ViewSienaSectorHandler)
	mux.HandleFunc("/editSienaSector/", siena.EditSienaSectorHandler)
	mux.HandleFunc("/saveSienaSector/", siena.SaveSienaSectorHandler)
	mux.HandleFunc("/newSienaSector/", siena.NewSienaSectorHandler)

	mux.HandleFunc("/listSienaFirm/", siena.ListSienaFirmHandler)
	mux.HandleFunc("/viewSienaFirm/", siena.ViewSienaFirmHandler)
	mux.HandleFunc("/editSienaFirm/", siena.EditSienaFirmHandler)
	mux.HandleFunc("/saveSienaFirm/", siena.SaveSienaFirmHandler)
	mux.HandleFunc("/newSienaFirm/", siena.NewSienaFirmHandler)

	mux.HandleFunc("/listSienaPortfolio/", siena.ListSienaPortfolioHandler)
	mux.HandleFunc("/viewSienaPortfolio/", siena.ViewSienaPortfolioHandler)
	mux.HandleFunc("/editSienaPortfolio/", siena.EditSienaPortfolioHandler)
	mux.HandleFunc("/saveSienaPortfolio/", siena.SaveSienaPortfolioHandler)
	mux.HandleFunc("/newSienaPortfolio/", siena.NewSienaPortfolioHandler)

	mux.HandleFunc("/listSienaCentre/", siena.ListSienaCentreHandler)
	mux.HandleFunc("/viewSienaCentre/", siena.ViewSienaCentreHandler)
	mux.HandleFunc("/editSienaCentre/", siena.EditSienaCentreHandler)
	mux.HandleFunc("/saveSienaCentre/", siena.SaveSienaCentreHandler)
	mux.HandleFunc("/newSienaCentre/", siena.NewSienaCentreHandler)

	mux.HandleFunc("/listSienaBook/", siena.ListSienaBookHandler)
	mux.HandleFunc("/viewSienaBook/", siena.ViewSienaBookHandler)
	mux.HandleFunc("/editSienaBook/", siena.EditSienaBookHandler)
	mux.HandleFunc("/saveSienaBook/", siena.SaveSienaBookHandler)
	mux.HandleFunc("/newSienaBook/", siena.NewSienaBookHandler)

	mux.HandleFunc("/listSienaBroker/", siena.ListSienaBrokerHandler)
	mux.HandleFunc("/viewSienaBroker/", siena.ViewSienaBrokerHandler)
	mux.HandleFunc("/editSienaBroker/", siena.EditSienaBrokerHandler)
	mux.HandleFunc("/saveSienaBroker/", siena.SaveSienaBrokerHandler)
	mux.HandleFunc("/newSienaBroker/", siena.NewSienaBrokerHandler)

	mux.HandleFunc("/listSienaAccount/", siena.ListSienaAccountHandler)
	mux.HandleFunc("/viewSienaAccount/", siena.ViewSienaAccountHandler)

	mux.HandleFunc("/listSienaCurrency/", siena.ListSienaCurrencyHandler)
	mux.HandleFunc("/listSienaCurrencyPair/", siena.ListSienaCurrencyPairHandler)

	mux.HandleFunc("/listSienaMandatedUser/", siena.ListSienaMandatedUserHandler)
	mux.HandleFunc("/viewSienaMandatedUser/", siena.ViewSienaMandatedUserHandler)
	mux.HandleFunc("/editSienaMandatedUser/", siena.EditSienaMandatedUserHandler)
	mux.HandleFunc("/saveSienaMandatedUser/", siena.SaveSienaMandatedUserHandler)
	mux.HandleFunc("/newSienaMandatedUser/", siena.NewSienaMandatedUserHandler)

	mux.HandleFunc("/listSienaCounterparty/", siena.ListSienaCounterpartyHandler)
	mux.HandleFunc("/viewSienaCounterparty/", siena.ViewSienaCounterpartyHandler)
	mux.HandleFunc("/editSienaCounterparty/", siena.EditSienaCounterpartyHandler)
	mux.HandleFunc("/saveSienaCounterparty/", siena.SaveSienaCounterpartyHandler)
	mux.HandleFunc("/newSienaCounterparty/", siena.NewSienaCounterpartyHandler)

	mux.HandleFunc("/editSienaCounterpartyAddress/", siena.EditSienaCounterpartyAddressHandler)
	mux.HandleFunc("/saveSienaCounterpartyAddress/", siena.SaveSienaCounterpartyAddressHandler)
	mux.HandleFunc("/editSienaCounterpartyExtensions/", siena.EditSienaCounterpartyExtensionsHandler)
	mux.HandleFunc("/saveSienaCounterpartyExtensions/", siena.SaveSienaCounterpartyExtensionsHandler)

	mux.HandleFunc("/listSienaCounterpartyPayee/", siena.ListSienaCounterpartyPayeeHandler)
	mux.HandleFunc("/viewSienaCounterpartyPayee/", siena.ViewSienaCounterpartyPayeeHandler)
	mux.HandleFunc("/editSienaCounterpartyPayee/", siena.EditSienaCounterpartyPayeeHandler)
	mux.HandleFunc("/saveSienaCounterpartyPayee/", siena.SaveSienaCounterpartyPayeeHandler)
	mux.HandleFunc("/newSienaCounterpartyPayee/", siena.NewSienaCounterpartyPayeeHandler)

	mux.HandleFunc("/listSienaCounterpartyImportID/", siena.ListSienaCounterpartyImportIDHandler)
	mux.HandleFunc("/viewSienaCounterpartyImportID/", siena.ViewSienaCounterpartyImportIDHandler)
	mux.HandleFunc("/editSienaCounterpartyImportID/", siena.EditSienaCounterpartyImportIDHandler)
	mux.HandleFunc("/saveSienaCounterpartyImportID/", siena.SaveSienaCounterpartyImportIDHandler)
	mux.HandleFunc("/newSienaCounterpartyImportID/", siena.NewSienaCounterpartyImportIDHandler)

	mux.HandleFunc("/listSienaDealList/", siena.ListSienaDealListHandler)
	mux.HandleFunc("/viewSienaDealList/", siena.ViewSienaDealListHandler)
	mux.HandleFunc("/listSienaAccountLadder/", siena.ListSienaAccountLadderHandler)
	mux.HandleFunc("/listSienaAccountTransactions/", siena.ListSienaAccountTransactionsHandler)
	//mux.HandleFunc("/saveSienaDealList/", siena.SaveSienaDealListHandler)
	//mux.HandleFunc("/newSienaDealList/", siena.NewSienaDealListHandler)

	mux.HandleFunc("/listSienaCounterpartyGroup/", siena.ListSienaCounterpartyGroupHandler)
	mux.HandleFunc("/viewSienaCounterpartyGroup/", siena.ViewSienaCounterpartyGroupHandler)
	mux.HandleFunc("/editSienaCounterpartyGroup/", siena.EditSienaCounterpartyGroupHandler)
	mux.HandleFunc("/saveSienaCounterpartyGroup/", siena.SaveSienaCounterpartyGroupHandler)
	mux.HandleFunc("/newSienaCounterpartyGroup/", siena.NewSienaCounterpartyGroupHandler)

	mux.HandleFunc("/dashboard/", siena.SienaDashboardHandler)

	mux.HandleFunc("/listCredentialsStore/", application.ListCredentialsStoreHandler)
	mux.HandleFunc("/viewCredentialsStore/", application.ViewCredentialStoreHandler)
	mux.HandleFunc("/editCredentialsStore/", application.EditCredentialStoreHandler)
	mux.HandleFunc("/deleteCredentialsStore/", application.DeleteCredentialStoreHandler)
	mux.HandleFunc("/saveCredentialsStore/", application.SaveCredentialStoreHandler)
	mux.HandleFunc("/newCredentialsStore/", application.NewCredentialStoreHandler)
	mux.HandleFunc("/banCredentialsStore/", application.BanCredentialStoreHandler)
	mux.HandleFunc("/activateCredentialsStore/", application.ActivateCredentialStoreHandler)

	mux.HandleFunc("/listMessageStore/", application.ListMessageStoreHandler)
	mux.HandleFunc("/viewMessageStore/", application.ViewMessageStoreHandler)
	mux.HandleFunc("/editMessageStore/", application.EditMessageStoreHandler)
	mux.HandleFunc("/saveMessageStore/", application.SaveMessageStoreHandler)

	mux.HandleFunc("/listScheduleStore/", application.ListScheduleStoreHandler)
	mux.HandleFunc("/viewScheduleStore/", application.ViewScheduleStoreHandler)

	mux.HandleFunc("/listSessionStore/", application.ListSessionStoreHandler)
	mux.HandleFunc("/deleteSessionStore/", application.DeleteSessionStoreHandler)

	mux.HandleFunc("/listSystemStore/", application.ListSystemStoreHandler)
	mux.HandleFunc("/viewSystemStore/", application.ViewSystemStoreHandler)
	mux.HandleFunc("/editSystemStore/", application.EditSystemStoreHandler)
	mux.HandleFunc("/deleteSystemStore/", application.DeleteSystemStoreHandler)
	mux.HandleFunc("/saveSystemStore/", application.SaveSystemStoreHandler)
	mux.HandleFunc("/newSystemStore/", application.NewSystemStoreHandler)

	mux.HandleFunc("/listFundsCheck/", application.ListFundsCheckHandler)
	mux.HandleFunc("/viewFundsCheck/", application.ViewFundsCheckHandler)
	mux.HandleFunc("/actionFundsCheck/", application.ActionFundsCheckHandler)
	mux.HandleFunc("/submitFundsCheck/", application.SubmitFundsCheckHandler)

	mux.HandleFunc("/injectSQLViews/", application.SQLInjectionHandler)

	mux.HandleFunc("/refreshCache/", application.RefreshCacheHandler)

	mux.HandleFunc("/shutdown/", shutdownHandler)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	done("Handlers Started")
	log.Println(line)
	header("Application Information")
	header("Application")
	log.Println("Host          :", tmpHostname)
	log.Println("Name          :", globals.ApplicationProperties["appname"])
	log.Println("Server Date   :", time.Now().Format(globals.DATEFORMATUSER))
	log.Println("Licence       :", globals.ApplicationProperties["licname"])
	log.Println("Lic URL       :", globals.ApplicationProperties["liclink"])
	log.Println("Session Life  :", globals.ApplicationProperties["sessionlife"])
	header("Paths")
	log.Println("Delivery      :", globals.ApplicationProperties["deliverpath"])
	log.Println("Response      :", globals.ApplicationProperties["receivepath"])
	log.Println("Processed     :", globals.ApplicationProperties["processedpath"])
	header("Software Release")
	log.Printf("Release       : %s [r%s-%s]", globals.ApplicationProperties["releaseid"], globals.ApplicationProperties["releaselevel"], globals.ApplicationProperties["releasenumber"])
	log.Println("Name          :", globals.ApplicationProperties["releaseid"])
	log.Println("Level         :", globals.ApplicationProperties["releaselevel"])
	log.Println("Number        :", globals.ApplicationProperties["releasenumber"])
	header("Siena")
	_, tempDate, _ := siena.GetBusinessDate(globals.SienaDB)
	globals.SienaSystemDate = tempDate
	log.Println("System Date   :", globals.SienaSystemDate.Internal.Format(globals.DATEFORMATUSER))

	// Get menu
	//menuCount, _ := application.FetchappMenuData("")
	//log.Println("No. Menu Items   :", menuCount)

	log.Println(line)
	header("READY STEADY GO!!!")
	log.Println("URI           :", globals.ColorPurple+"http://localhost:"+globals.ApplicationProperties["port"]+globals.ColorReset)
	log.Println(line)

	httpPort := ":" + globals.ApplicationProperties["port"]
	//http.ListenAndServe(httpPort, nil)

	// Wrap your handlers with the LoadAndSave() middleware.
	http.ListenAndServe(httpPort, globals.SessionManager.LoadAndSave(mux))

}

//// TODO: migrage the following three functions to appsupport
func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	//	wctProperties := application.GetProperties(APPCONFIG)

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	//requestID := uuid.New()

	log.Println("Servicing :", inUTL)
	//requestMessage := application.BuildRequestMessage(requestID.String(), "SHUTDOWN", line, line, line, wctProperties)

	//fmt.Println("requestMessage", requestMessage)
	//fmt.Println("SEND MESSAGE")
	//application.SendRequest(requestMessage, requestID.String(), globals.ApplicationProperties)
	m := http.NewServeMux()

	s := http.Server{Addr: globals.ApplicationProperties["port"], Handler: m}
	s.Shutdown(context.Background())
	//	r.URL.Path = "/viewResponse?uuid=" + requestID.String()
	//	viewResponseHandler(w, r)

}
func clearQueuesHandler(w http.ResponseWriter, r *http.Request) {

	//var propertiesFileName = "config/properties.cfg"
	//wctProperties := application.GetProperties(APPCONFIG)
	//	tmpl := "viewResponse"
	inUTL := r.URL.Path
	//requestID := uuid.New()
	log.Println("Servicing :", inUTL)
	//fmt.Println("delivPath", wctProperties["deliverpath"])
	err1 := application.RemoveContents(globals.ApplicationProperties["deliverpath"])
	if err1 != nil {
		fmt.Println(err1)
	}
	//fmt.Println("recPath", wctProperties["receivepath"])
	err2 := application.RemoveContents(globals.ApplicationProperties["receivepath"])
	if err2 != nil {
		fmt.Println(err2)
	}
	//fmt.Println("procPath", wctProperties["processedpath"])
	err3 := application.RemoveContents(globals.ApplicationProperties["processedpath"])
	if err3 != nil {
		fmt.Println(err3)
	}
	application.HomePageHandler(w, r)
}
func clearResponsesHandler(w http.ResponseWriter, r *http.Request) {
	//var propertiesFileName = "config/properties.cfg"
	//	wctProperties := application.GetProperties(globals.APPCONFIG)
	//	tmpl := "viewResponse"
	inUTL := r.URL.Path
	//requestID := uuid.New()
	log.Println("Servicing :", inUTL)
	application.RemoveContents(globals.ApplicationProperties["receivepath"])
	application.HomePageHandler(w, r)
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	// Store a new key and value in the session data.
	globals.SessionManager.Put(r.Context(), "message", "Hello from a session!")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	// Use the GetString helper to retrieve the string value associated with a
	// key. The zero value is returned if the key does not exist.
	msg := globals.SessionManager.GetString(r.Context(), "message")
	io.WriteString(w, msg)
}

func header(s string) {
	log.Println(globals.ColorYellow + s + globals.ColorReset)
	//log.Println(strings.Repeat("-", len(s)))
}
func done(s string) {
	log.Println(globals.ColorYellow + globals.Tick + globals.ColorReset + " " + s)
}
