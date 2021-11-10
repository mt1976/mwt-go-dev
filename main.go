package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"

	application "github.com/mt1976/mwt-go-dev/application"
	core "github.com/mt1976/mwt-go-dev/core"
	scheduler "github.com/mt1976/mwt-go-dev/jobs"
)

func main() {

	//ascii := figlet4go.NewAsciiRender()

	//wctProperties := application.GetProperties(APPCONFIG)

	// The underscore would be an error
	//	renderStr, _ := ascii.Render(wctProperties["appname"])
	tmpHostname, _ := os.Hostname()
	line := strings.Repeat("-", 100)
	log.Println(line)

	core.LOG_header("Initialising ...")

	core.Initialise()

	core.LOG_success("Initialised")

	core.LOG_header("Scheduling Jobs")
	//log.Println("TEST>")
	//scheduler.RunJobCob("TEST")
	//log.Println("<TEST")
	scheduler.Start()
	core.LOG_success("Jobs Scheduled")

	core.LOG_header("Starting Handlers")
	mux := http.NewServeMux()

	Main_Publish(*mux)
	mux.HandleFunc("/", core.LoginHandler)
	mux.HandleFunc("/login", core.ValidateLoginHandler)
	mux.HandleFunc("/logout", core.LogoutHandler)
	mux.HandleFunc("/home", application.HomePageHandler)
	//mux.HandleFunc("/srvServiceCatalog", application.ServiceCatalogHandler)
	mux.HandleFunc("/favicon.ico", application.FaviconHandler)
	mux.HandleFunc("/favicon-32x32.png", application.Favicon32Handler)
	mux.HandleFunc("/site.webmanifest", application.FaviconManifestHandler)
	mux.HandleFunc("/favicon-16x16.png", application.Favicon16Handler)
	mux.HandleFunc("/browserconfig.xml", application.FaviconBrowserConfigHandler)
	// mux.HandleFunc("/listResponses/", application.ListResponsesHandler)
	// mux.HandleFunc("/previewRequest/", application.PreviewRequestHandler)
	// mux.HandleFunc("/executeRequest/", application.ExecuteRequestHandler)
	// mux.HandleFunc("/viewResponse/", application.ViewResponseHandler)
	// mux.HandleFunc("/deleteResponse/", application.DeleteResponseHandler)

	// mux.HandleFunc("/viewSrvEnvironment/", application.ViewSrvEnvironmentHandler)
	// mux.HandleFunc("/listSrvConfiguration/", application.ListSrvConfigurationHandler)
	// mux.HandleFunc("/viewSrvConfiguration/", application.ViewSrvConfigurationHandler)
	// mux.HandleFunc("/editSrvConfiguration/", application.EditSrvConfigurationHandler)
	// mux.HandleFunc("/saveSrvConfiguration/", application.SaveSrvConfigurationHandler)
	mux.HandleFunc("/viewAppConfiguration/", application.ViewAppConfigurationHandler)

	mux.HandleFunc("/listSvcDataMap/", application.ListSvcDataMapHandler)
	mux.HandleFunc("/viewSvcDataMap/", application.ViewSvcDataMapHandler)
	mux.HandleFunc("/editSvcDataMap/", application.EditSvcDataMapHandler)
	mux.HandleFunc("/saveSvcDataMap/", application.SaveSvcDataMapHandler)
	mux.HandleFunc("/viewSvcDataMapXML/", application.ViewSvcDataMapXMLHandler)
	mux.HandleFunc("/editSvcDataMapXML/", application.EditSvcDataMapXMLHandler)
	mux.HandleFunc("/saveSvcDataMapXML/", application.SaveSvcDataMapXMLHandler)
	mux.HandleFunc("/newSvcDataMap/", application.NewSvcDataMapHandler)
	mux.HandleFunc("/genSvcDataMap/", application.GenSvcDataMapHandler)
	mux.HandleFunc("/deleteSvcDataMap/", application.DeleteSvcDataMapHandler)
	mux.HandleFunc("/editDataMapColumns/", application.ListLoaderMapStoreHandler)
	mux.HandleFunc("/editLoaderMapStore/", application.EditLoaderMapStoreHandler)
	mux.HandleFunc("/saveLoaderMapStore/", application.SaveLoaderMapStoreHandler)
	mux.HandleFunc("/newLoaderMapStore/", application.NewLoaderMapStoreHandler)
	mux.HandleFunc("/runLoader/", application.RunDataLoaderHandler)

	application.Country_Publish(*mux)
	application.Sector_Publish(*mux)
	application.Firm_Publish(*mux)
	application.Portfolio_Publish(*mux)

	application.Centre_Publish(*mux)

	mux.HandleFunc("/listSienaBook/", application.ListSienaBookHandler)
	mux.HandleFunc("/viewSienaBook/", application.ViewSienaBookHandler)
	mux.HandleFunc("/editSienaBook/", application.EditSienaBookHandler)
	mux.HandleFunc("/saveSienaBook/", application.SaveSienaBookHandler)
	mux.HandleFunc("/newSienaBook/", application.NewSienaBookHandler)

	mux.HandleFunc("/listSienaBroker/", application.ListSienaBrokerHandler)
	mux.HandleFunc("/viewSienaBroker/", application.ViewSienaBrokerHandler)
	mux.HandleFunc("/editSienaBroker/", application.EditSienaBrokerHandler)
	mux.HandleFunc("/saveSienaBroker/", application.SaveSienaBrokerHandler)
	mux.HandleFunc("/newSienaBroker/", application.NewSienaBrokerHandler)

	application.Account_Publish(*mux)

	mux.HandleFunc("/listSienaCurrency/", application.ListSienaCurrencyHandler)
	mux.HandleFunc("/listSienaCurrencyPair/", application.ListSienaCurrencyPairHandler)

	mux.HandleFunc("/listSienaMandatedUser/", application.ListSienaMandatedUserHandler)
	mux.HandleFunc("/viewSienaMandatedUser/", application.ViewSienaMandatedUserHandler)
	mux.HandleFunc("/editSienaMandatedUser/", application.EditSienaMandatedUserHandler)
	mux.HandleFunc("/saveSienaMandatedUser/", application.SaveSienaMandatedUserHandler)
	mux.HandleFunc("/newSienaMandatedUser/", application.NewSienaMandatedUserHandler)

	mux.HandleFunc("/listSienaCounterparty/", application.ListSienaCounterpartyHandler)
	mux.HandleFunc("/viewSienaCounterparty/", application.ViewSienaCounterpartyHandler)
	mux.HandleFunc("/editSienaCounterparty/", application.EditSienaCounterpartyHandler)
	mux.HandleFunc("/saveSienaCounterparty/", application.SaveSienaCounterpartyHandler)
	mux.HandleFunc("/newSienaCounterparty/", application.NewSienaCounterpartyHandler)

	mux.HandleFunc("/editSienaCounterpartyAddress/", application.EditSienaCounterpartyAddressHandler)
	mux.HandleFunc("/saveSienaCounterpartyAddress/", application.SaveSienaCounterpartyAddressHandler)
	mux.HandleFunc("/editSienaCounterpartyExtensions/", application.EditSienaCounterpartyExtensionsHandler)
	mux.HandleFunc("/saveSienaCounterpartyExtensions/", application.SaveSienaCounterpartyExtensionsHandler)

	mux.HandleFunc("/listSienaCounterpartyPayee/", application.ListSienaCounterpartyPayeeHandler)
	mux.HandleFunc("/viewSienaCounterpartyPayee/", application.ViewSienaCounterpartyPayeeHandler)
	mux.HandleFunc("/editSienaCounterpartyPayee/", application.EditSienaCounterpartyPayeeHandler)
	mux.HandleFunc("/saveSienaCounterpartyPayee/", application.SaveSienaCounterpartyPayeeHandler)
	mux.HandleFunc("/newSienaCounterpartyPayee/", application.NewSienaCounterpartyPayeeHandler)

	mux.HandleFunc("/listSienaCounterpartyImportID/", application.ListSienaCounterpartyImportIDHandler)
	mux.HandleFunc("/viewSienaCounterpartyImportID/", application.ViewSienaCounterpartyImportIDHandler)
	mux.HandleFunc("/editSienaCounterpartyImportID/", application.EditSienaCounterpartyImportIDHandler)
	mux.HandleFunc("/saveSienaCounterpartyImportID/", application.SaveSienaCounterpartyImportIDHandler)
	mux.HandleFunc("/newSienaCounterpartyImportID/", application.NewSienaCounterpartyImportIDHandler)

	mux.HandleFunc("/listSienaDealList/", application.ListSienaDealListHandler)
	mux.HandleFunc("/viewSienaDealList/", application.ViewSienaDealListHandler)
	mux.HandleFunc("/listSienaAccountLadder/", application.ListSienaAccountLadderHandler)
	mux.HandleFunc("/listSienaAccountTransactions/", application.ListSienaAccountTransactionsHandler)
	//mux.HandleFunc("/saveSienaDealList/", application.SaveSienaDealListHandler)
	//mux.HandleFunc("/newSienaDealList/", application.NewSienaDealListHandler)

	mux.HandleFunc("/listSienaCounterpartyGroup/", application.ListSienaCounterpartyGroupHandler)
	mux.HandleFunc("/viewSienaCounterpartyGroup/", application.ViewSienaCounterpartyGroupHandler)
	mux.HandleFunc("/editSienaCounterpartyGroup/", application.EditSienaCounterpartyGroupHandler)
	mux.HandleFunc("/saveSienaCounterpartyGroup/", application.SaveSienaCounterpartyGroupHandler)
	mux.HandleFunc("/newSienaCounterpartyGroup/", application.NewSienaCounterpartyGroupHandler)

	mux.HandleFunc("/dashboard/", application.SienaDashboardHandler)

	application.Credentials_Publish(*mux)

	mux.HandleFunc("/listMessageStore/", application.ListMessageStoreHandler)
	mux.HandleFunc("/viewMessageStore/", application.ViewMessageStoreHandler)
	mux.HandleFunc("/editMessageStore/", application.EditMessageStoreHandler)
	mux.HandleFunc("/saveMessageStore/", application.SaveMessageStoreHandler)

	// mux.HandleFunc("/listTranslationStore/", application.ListTranslationStoreHandler)
	// mux.HandleFunc("/viewTranslationStore/", application.ViewTranslationStoreHandler)
	// mux.HandleFunc("/editTranslationStore/", application.EditTranslationStoreHandler)
	// mux.HandleFunc("/saveTranslationStore/", application.SaveTranslationStoreHandler)
	application.Translation_Publish(*mux)

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

	//mux.HandleFunc("/injectSQLViews/", application.SQLInjection_HandlerRun)
	application.SQLInjection_Publish(*mux)

	//	mux.HandleFunc("/refreshCache/", application.RefreshCacheHandler)
	application.DataCache_Publish(*mux)

	mux.HandleFunc("/listGiltsDataStore/", application.ListLSEGiltsDataStoreHandler)
	mux.HandleFunc("/viewLSEGiltsDataStore/", application.ViewLSEGiltsDataStoreHandler)
	mux.HandleFunc("/selectLSEGiltsDataStore/", application.SelectLSEGiltsDataStoreHandler)
	mux.HandleFunc("/deselectLSEGiltsDataStore/", application.DeselectLSEGiltsDataStoreHandler)

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	core.LOG_success("Handlers Started")
	log.Println(line)
	core.LOG_header("Application Information")
	core.LOG_header("Application")
	core.LOG_info("Name", core.ApplicationProperties["appname"])
	core.LOG_info("Host Name", tmpHostname)
	core.LOG_info("Server Release", fmt.Sprintf("%s [r%s-%s]", core.ApplicationProperties["releaseid"], core.ApplicationProperties["releaselevel"], core.ApplicationProperties["releasenumber"]))
	core.LOG_info("Server Date", time.Now().Format(core.DATEFORMATUSER))
	if core.IsChildInstance {
		core.LOG_info("Server Mode", "Primary System")
	} else {
		core.LOG_info("Server Mode", "Secondary System")
	}
	core.LOG_info("Licence", core.ApplicationProperties["licname"])
	core.LOG_info("Lic URL", core.ApplicationProperties["liclink"])
	core.LOG_header("Runtime")
	core.LOG_info("GO Version", runtime.Version())
	core.LOG_info("Operating System", runtime.GOOS)
	core.LOG_header("Application Database (MSSQL)")
	core.LOG_info("Server", core.ApplicationPropertiesDB["server"])
	core.LOG_info("Database", core.ApplicationPropertiesDB["database"])
	core.LOG_info("Schema", core.ApplicationPropertiesDB["schema"])
	core.LOG_info("Parent Schema", core.ApplicationPropertiesDB["parentschema"])

	core.LOG_header("Siena")
	_, tempDate, _ := application.GetBusinessDate(core.SienaDB)
	core.SienaSystemDate = tempDate
	core.LOG_info("System", core.SienaProperties["name"])
	core.LOG_info("System Date", core.SienaSystemDate.Internal.Format(core.DATEFORMATUSER))

	core.LOG_header("Siena Database (MSSQL)")
	core.LOG_info("Server", core.SienaPropertiesDB["server"])
	core.LOG_info("Database", core.SienaPropertiesDB["database"])
	core.LOG_info("Schema", core.SienaPropertiesDB["schema"])
	core.LOG_info("Parent Schema", core.SienaPropertiesDB["parentschema"])

	core.LOG_header("Siena Connectivity")
	core.LOG_info("TXNs Delivery", core.SienaProperties["transactional_in"])
	core.LOG_info("TXNs Response", core.SienaProperties["transactional_out"])
	core.LOG_info("Static Delivery", core.SienaProperties["static_in"])
	core.LOG_info("Static Response", core.SienaProperties["static_out"])
	core.LOG_info("Funds Check Request", core.SienaProperties["funds_out"])
	core.LOG_info("Funds Check Response", core.SienaProperties["funds_in"])
	core.LOG_info("Rates & Prices Delivery", core.SienaProperties["rates_in"])

	core.LOG_header("Sessions")
	core.LOG_info("Session Life", core.ApplicationProperties["sessionlife"])

	//scheduler.RunJobLSE("")
	//scheduler.RunJobFII("")

	core.LOG_header("READY STEADY GO!!!")
	log.Println("URI            :", core.ColorPurple+"http://localhost:"+core.ApplicationProperties["port"]+core.ColorReset)
	log.Println(line)

	//rec, err := das.Query(core.ApplicationDB, "SELECT * FROM dbo.credentialsStore")
	log.Println(line)

	// log.Println("rec:", rec, err, len(rec))
	// for i, _ := range rec {
	// 	log.Println("row:", i, rec[i])
	// 	//	log.Println("row:", row[])
	// }
	//spew.Dump(rec)

	// jobs.RatesFXSpot_Run()
	// jobs.SessionHouseKeeping_Run()
	// jobs.RatesFXSpot_Run()
	// jobs.SessionHouseKeeping_Run()
	// jobs.RatesFXSpot_Run()
	// jobs.SessionHouseKeeping_Run()

	httpPort := ":" + core.ApplicationProperties["port"]
	//http.ListenAndServe(httpPort, nil)

	//scheduler.RunJobFII("")
	//s, _ := application.GLIEF_leiLookup("GB00BL68HJ26")
	//info("LEI", s)
	// Wrap your handlers with the LoadAndSave() middleware.
	http.ListenAndServe(httpPort, core.SessionManager.LoadAndSave(mux))

	core.Log_uptime()

}

func Main_Publish(mux http.ServeMux) {

	mux.HandleFunc("/shutdown/", shutdownHandler)
	mux.HandleFunc("/clearQueues/", clearQueuesHandler)
	mux.HandleFunc("/put", putHandler)
	mux.HandleFunc("/get", getHandler)
	core.LOG_mux("Main", "Application")
}

//// TODO: migrage the following three functions to appsupport
func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	//	wctProperties := application.GetProperties(APPCONFIG)

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	//requestID := uuid.New()http.ListenAndServe(":8080", nil)

	log.Println("Servicing :", inUTL)
	//requestMessage := application.BuildRequestMessage(requestID.String(), "SHUTDOWN", line, line, line, wctProperties)

	//fmt.Println("requestMessage", requestMessage)
	//fmt.Println("SEND MESSAGE")
	//application.SendRequest(requestMessage, requestID.String(), core.ApplicationProperties)
	m := http.NewServeMux()

	s := http.Server{Addr: core.ApplicationProperties["port"], Handler: m}
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
	err1 := core.RemoveContents(core.ApplicationProperties["deliverpath"])
	if err1 != nil {
		fmt.Println(err1)
	}
	//fmt.Println("recPath", wctProperties["receivepath"])
	err2 := core.RemoveContents(core.ApplicationProperties["receivepath"])
	if err2 != nil {
		fmt.Println(err2)
	}
	//fmt.Println("procPath", wctProperties["processedpath"])
	err3 := core.RemoveContents(core.ApplicationProperties["processedpath"])
	if err3 != nil {
		fmt.Println(err3)
	}
	application.HomePageHandler(w, r)
}

// func clearResponsesHandler(w http.ResponseWriter, r *http.Request) {
// 	//var propertiesFileName = "config/properties.cfg"
// 	//	wctProperties := application.GetProperties(core.APPCONFIG)
// 	//	tmpl := "viewResponse"
// 	inUTL := r.URL.Path
// 	//requestID := uuid.New()
// 	log.Println("Servicing :", inUTL)
// 	application.RemoveContents(core.ApplicationProperties["receivepath"])
// 	application.HomePageHandler(w, r)
// }

func putHandler(w http.ResponseWriter, r *http.Request) {
	// Store a new key and value in the session data.
	core.SessionManager.Put(r.Context(), "message", "Hello from a session!")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	// Use the GetString helper to retrieve the string value associated with a
	// key. The zero value is returned if the key does not exist.
	msg := core.SessionManager.GetString(r.Context(), "message")
	io.WriteString(w, msg)
}
