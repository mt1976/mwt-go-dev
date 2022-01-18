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

	monitors "github.com/mt1976/mwt-go-dev/adaptor/monitors"
	application "github.com/mt1976/mwt-go-dev/application"
	core "github.com/mt1976/mwt-go-dev/core"
	scheduler "github.com/mt1976/mwt-go-dev/jobs"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func main() {

	logs.Break()
	logs.Header("MISSION CONTROL")
	logs.Break()

	logs.Information("Initialising...", "")

	core.Initialise()

	logs.Success("Initialised")
	logs.Break()
	logs.Header("Scheduling Jobs")
	logs.Break()

	scheduler.Start()

	logs.Success("Jobs Scheduled")
	logs.Break()
	logs.Header("Publish Endpoints")
	logs.Break()
	// Setup Endpoints
	mux := http.NewServeMux()
	// At least one "mux" handler is required - Dont remove this
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	// At least one "mux" handler is required - Dont remove this

	Main_Publish(*mux)
	core.LoginLogout_Publish_Impl(*mux)
	application.Resources_Publish_Impl(*mux)
	application.Home_Publish_Impl(*mux)
	application.Session_Publish_Impl(*mux)

	application.Configuration_Publish_Impl(*mux)

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
	application.AccountTransaction_Publish_Impl(*mux)
	application.AccountLadder_Publish_Impl(*mux)
	application.Payee_Publish(*mux)
	application.Payee_PublishImpl(*mux)

	application.Currency_Publish(*mux)
	application.CurrencyPair_Publish(*mux)

	application.Mandate_Publish(*mux)
	application.Mandate_PublishImpl(*mux)

	application.CounterpartyGroup_Publish(*mux)
	application.Dashboard_Publish_Impl(*mux)
	application.DealingInterface_Publish(*mux)

	application.Credentials_Publish(*mux)

	application.Message_Publish(*mux)

	application.Translation_Publish(*mux)

	application.Schedule_Publish(*mux)

	application.Session_Publish(*mux)

	application.Systems_Publish(*mux)

	application.Simulator_SienaFundsChecker_Publish_Impl(*mux)

	application.Template_Publish(*mux)
	application.MarketRates_Publish(*mux)
	application.Cache_Publish(*mux)
	application.DealConversation_Publish(*mux)

	application.SQLInjection_Publish(*mux)

	application.NegotiableInstrument_Publish(*mux)
	application.NegotiableInstrument_Publish_Impl(*mux)

	application.CMNotes_Publish(*mux)
	application.CounterpartyOnboarding_Publish(*mux)
	application.CounterpartyImport_Publish(*mux)
	application.CounterpartyName_Publish(*mux)
	application.CounterpartyAddress_Publish(*mux)
	application.CounterpartyExtensions_Publish(*mux)
	application.CounterpartyCreditRating_Publish(*mux)
	application.Counterparty_Publish(*mux)
	application.Transaction_Publish(*mux)
	application.Counterparty_Publish_Impl(*mux)

	application.DataLoader_Publish(*mux)
	application.DataLoaderData_Publish(*mux)
	application.DataLoaderMap_Publish(*mux)
	application.DataLoader_Publish_Impl(*mux)

	application.ExternalMessage_Publish(*mux)
	application.Catalog_Publish(*mux)
	// End of Endpoints

	logs.Header("Publish API")

	core.Catalog_List()

	logs.Success("Endpoints Published")
	logs.Break()
	logs.Header("Start Watchers")
	logs.Break()
	//go monitors.StaticDataImporter_Watch()
	go monitors.Simulator_SienaFundsChecker_Watch()
	go monitors.Simulator_SienaDealImporter_Watch()
	go monitors.Simulator_SienaStaticDataImporter_Watch()
	logs.Success("Watchers Started")
	Application_Info()
	//scheduler.RunJobLSE("")
	//scheduler.RunJobFII("")
	//jobs.RatesFXSpot_Run()
	//spew.Dump(mux)

	//core.Notification_Test()

	logs.Header("READY STEADY GO!!!")
	logs.Information("Initialisation", "Vrooom, Vrooooom, Vroooooooo..."+logs.Character_Bike+logs.Character_Bike+logs.Character_Bike+logs.Character_Bike)
	logs.Break()

	httpProtocol := core.ApplicationProperties["protocol"]
	logs.URI(httpProtocol + "://localhost:" + core.ApplicationProperties["port"])
	logs.Break()

	httpPort := ":" + core.ApplicationProperties["port"]

	if core.ApplicationProperties["environment"] == "development" {

		http.ListenAndServe(httpPort, core.SessionManager.LoadAndSave(mux))

	} else {

		certPath := core.ApplicationProperties["certpath"]
		certName := core.ApplicationProperties["certname"]

		pwd, _ := os.Getwd()
		certLocation := pwd + certPath + certName
		certKey := certLocation + ".key"
		certCrt := certLocation + ".crt"

		logs.Information("Certificate Path", certPath)
		logs.Information("Certificate Name", certName)
		logs.Information("Certificate Location", certLocation)
		logs.Information("Certificate Key", certKey)
		logs.Information("Certificate Crt", certCrt)

		log.Fatal(http.ListenAndServeTLS(httpPort, certCrt, certKey, core.SessionManager.LoadAndSave(mux)))
	}
	logs.Break()
	core.Log_uptime()
	core.Log_uptime()

}

func Main_Publish(mux http.ServeMux) {

	mux.HandleFunc("/shutdown/", application_HandlerShutdown)
	mux.HandleFunc("/clearQueues/", application_HandlerClearQueues)
	mux.HandleFunc("/put", application_HandlerPUT)
	mux.HandleFunc("/get", application_HandlerGET)
	logs.Publish("Main", "Application")
}

//// TODO: migrage the following three functions to appsupport
func application_HandlerShutdown(w http.ResponseWriter, r *http.Request) {
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
func application_HandlerClearQueues(w http.ResponseWriter, r *http.Request) {

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
	application.Home_HandlerView(w, r)
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

func application_HandlerPUT(w http.ResponseWriter, r *http.Request) {
	// Store a new key and value in the session data.
	core.SessionManager.Put(r.Context(), "message", "Hello from a session!")
}

func application_HandlerGET(w http.ResponseWriter, r *http.Request) {
	// Use the GetString helper to retrieve the string value associated with a
	// key. The zero value is returned if the key does not exist.
	msg := core.SessionManager.GetString(r.Context(), "message")
	io.WriteString(w, msg)
}

func Application_Info() {
	tmpHostname, _ := os.Hostname()
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
}
