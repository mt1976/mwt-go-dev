package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	_ "github.com/denisenkom/go-mssqldb"

	application "github.com/mt1976/mwt-go-dev/application"
	core "github.com/mt1976/mwt-go-dev/core"
	"github.com/mt1976/mwt-go-dev/dao"
	scheduler "github.com/mt1976/mwt-go-dev/jobs"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func main() {

	//ascii := figlet4go.NewAsciiRender()

	//wctProperties := application.GetProperties(APPCONFIG)

	// The underscore would be an error
	//	renderStr, _ := ascii.Render(wctProperties["appname"])
	tmpHostname, _ := os.Hostname()
	logs.Break()
	logs.Header("MISSION CONTROL")
	logs.Break()

	logs.Information("Initialising...", "")

	core.Initialise()

	logs.Success("Initialised")
	logs.Break()
	logs.Header("Scheduling Jobs")
	logs.Break()

	//log.Println("TEST>")
	//scheduler.RunJobCob("TEST")
	//log.Println("<TEST")

	scheduler.Start()

	logs.Success("Jobs Scheduled")
	logs.Break()
	logs.Header("Starting Handlers")
	logs.Break()

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

	application.Book_Publish(*mux)
	application.Product_Publish(*mux)
	application.DealType_Publish(*mux)
	application.DealTypeFundamental_Publish(*mux)

	application.Broker_Publish(*mux)

	application.Account_Publish(*mux)
	application.AccountTransaction_Publish(*mux)
	application.AccountLadder_Publish(*mux)
	application.Payee_Publish(*mux)
	application.Payee_PublishImpl(*mux)

	application.Currency_Publish(*mux)
	application.CurrencyPair_Publish(*mux)

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

	mux.HandleFunc("/listSienaCounterpartyImportID/", application.ListSienaCounterpartyImportIDHandler)
	mux.HandleFunc("/viewSienaCounterpartyImportID/", application.ViewSienaCounterpartyImportIDHandler)
	mux.HandleFunc("/editSienaCounterpartyImportID/", application.EditSienaCounterpartyImportIDHandler)
	mux.HandleFunc("/saveSienaCounterpartyImportID/", application.SaveSienaCounterpartyImportIDHandler)
	mux.HandleFunc("/newSienaCounterpartyImportID/", application.NewSienaCounterpartyImportIDHandler)

	mux.HandleFunc("/listSienaDealList/", application.ListSienaDealListHandler)
	mux.HandleFunc("/viewSienaDealList/", application.ViewSienaDealListHandler)

	//mux.HandleFunc("/saveSienaDealList/", application.SaveSienaDealListHandler)
	//mux.HandleFunc("/newSienaDealList/", application.NewSienaDealListHandler)

	application.CounterpartyGroup_Publish(*mux)

	mux.HandleFunc("/dashboard/", application.SienaDashboardHandler)
	application.DealingInterface_Publish(*mux)

	application.Credentials_Publish(*mux)

	application.Message_Publish(*mux)

	application.Translation_Publish(*mux)

	application.Schedule_Publish(*mux)

	application.Session_Publish(*mux)

	// mux.HandleFunc("/listSystemStore/", application.ListSystemStoreHandler)
	// mux.HandleFunc("/viewSystemStore/", application.ViewSystemStoreHandler)
	// mux.HandleFunc("/editSystemStore/", application.EditSystemStoreHandler)
	// mux.HandleFunc("/deleteSystemStore/", application.DeleteSystemStoreHandler)
	// mux.HandleFunc("/saveSystemStore/", application.SaveSystemStoreHandler)
	// mux.HandleFunc("/newSystemStore/", application.NewSystemStoreHandler)

	application.Systems_Publish(*mux)

	application.Simulator_FundsChecker_PublishImpl(*mux)

	application.Template_Publish(*mux)
	application.MarketRates_Publish(*mux)
	application.Cache_Publish(*mux)
	application.DealConversation_Publish(*mux)

	//mux.HandleFunc("/injectSQLViews/", application.SQLInjection_HandlerRun)
	application.SQLInjection_Publish(*mux)

	//	mux.HandleFunc("/refreshCache/", application.RefreshCacheHandler)
	application.DataCache_Publish(*mux)
	application.NegotiableInstrument_Publish(*mux)
	application.NegotiableInstrument_PublishImpl(*mux)

	application.CMNotes_Publish(*mux)
	dao.Onboard_Test()
	application.CounterpartyOnboarding_Publish(*mux)

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	logs.Success("Handlers Started")
	logs.Break()
	logs.Header("Application Information")
	logs.Break()

	logs.Header("Application")
	logs.Information("Name", core.ApplicationProperties["appname"])
	logs.Information("Host Name", tmpHostname)
	logs.Information("Server Release", fmt.Sprintf("%s [r%s-%s]", core.ApplicationProperties["releaseid"], core.ApplicationProperties["releaselevel"], core.ApplicationProperties["releasenumber"]))
	logs.Information("Server Date", time.Now().Format(core.DATEFORMATUSER))
	if core.IsChildInstance {
		logs.Information("Server Mode", "Primary System")
	} else {
		logs.Information("Server Mode", "Secondary System")
	}
	logs.Information("Licence", core.ApplicationProperties["licname"])
	logs.Information("Lic URL", core.ApplicationProperties["liclink"])
	logs.Header("Runtime")
	logs.Information("GO Version", runtime.Version())
	logs.Information("Operating System", runtime.GOOS+" ("+runtime.GOARCH+")")
	logs.Header("Application Database (MSSQL)")
	logs.Information("Server", core.ApplicationPropertiesDB["server"])
	logs.Information("Database", core.ApplicationPropertiesDB["database"])
	logs.Information("Schema", core.ApplicationPropertiesDB["schema"])
	logs.Information("Parent Schema", core.ApplicationPropertiesDB["parentschema"])

	logs.Header("Siena")
	_, tempDate, _ := application.GetBusinessDate(core.SienaDB)
	core.SienaSystemDate = tempDate
	logs.Information("System", core.SienaProperties["name"])
	logs.Information("System Date", core.SienaSystemDate.Internal.Format(core.DATEFORMATUSER))

	logs.Header("Siena Database (MSSQL)")
	logs.Information("Server", core.SienaPropertiesDB["server"])
	logs.Information("Database", core.SienaPropertiesDB["database"])
	logs.Information("Schema", core.SienaPropertiesDB["schema"])
	logs.Information("Parent Schema", core.SienaPropertiesDB["parentschema"])

	logs.Header("Siena Connectivity")
	logs.Information("TXNs Delivery", core.SienaProperties["transactional_in"])
	logs.Information("TXNs Response", core.SienaProperties["transactional_out"])
	logs.Information("Static Delivery", core.SienaProperties["static_in"])
	logs.Information("Static Response", core.SienaProperties["static_out"])
	logs.Information("Funds Check Request", core.SienaProperties["funds_out"])
	logs.Information("Funds Check Response", core.SienaProperties["funds_in"])
	logs.Information("Rates & Prices Delivery", core.SienaProperties["rates_in"])

	logs.Header("Sessions")
	logs.Information("Session Life", core.ApplicationProperties["sessionlife"])

	//scheduler.RunJobLSE("")
	//scheduler.RunJobFII("")
	//jobs.RatesFXSpot_Run()

	logs.Header("READY STEADY GO!!!")
	logs.Information("Initialisation", "Vrooom, Vrooooom, Vroooooooo..."+logs.Character_Bike+logs.Character_Bike+logs.Character_Bike+logs.Character_Bike)
	logs.Break()
	logs.URI("http://localhost:" + core.ApplicationProperties["port"])
	logs.Break()

	httpPort := ":" + core.ApplicationProperties["port"]

	// Wrap your handlers with the LoadAndSave() middleware.
	//spew.Dump(mux)
	http.ListenAndServe(httpPort, core.SessionManager.LoadAndSave(mux))
	logs.Break()
	core.Log_uptime()

}

func Main_Publish(mux http.ServeMux) {

	mux.HandleFunc("/shutdown/", shutdownHandler)
	mux.HandleFunc("/clearQueues/", clearQueuesHandler)
	mux.HandleFunc("/put", putHandler)
	mux.HandleFunc("/get", getHandler)
	logs.Publish("Main", "Application")
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
