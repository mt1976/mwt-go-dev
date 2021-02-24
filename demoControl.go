package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/mbndr/figlet4go"
)

//CONST_CONFIG_FILE is cheese
const (
	CONST_CONFIG_FILE = "config/properties.cfg"
)

//HomePage ...
type HomePage struct {
	Title          string
	Body           string
	RequestPath    string
	ResponsePath   string
	ProcessedPath  string
	NoResponses    int
	Responses      []WctResponsePayload
	NoServices     int
	Services       string
	ServiceCatalog []ServiceCatalogItem
	Description    string
	ResponseType   string
}

func main() {

	ascii := figlet4go.NewAsciiRender()

	wctProperties := getProperties(CONST_CONFIG_FILE)

	// The underscore would be an error
	renderStr, _ := ascii.Render("MWT GO PROTO")

	fmt.Print(renderStr)

	fmt.Println("Delivery Path :", wctProperties["deliverpath"])
	fmt.Println("Reponse Path :", wctProperties["receivepath"])

	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/favicon-32x32.png", favicon32Handler)
	http.HandleFunc("/site.webmanifest", faviconManifestHandler)
	http.HandleFunc("/favicon-16x16.png", favicon16Handler)
	http.HandleFunc("/browserconfig.xml", faviconBrowserConfigHandler)
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
	http.HandleFunc("/shutdown/", shutdownHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Println("URL", "http://localhost:"+wctProperties["port"])

	httpPort := ":" + wctProperties["port"]
	http.ListenAndServe(httpPort, nil)

}

func homePageHandler(w http.ResponseWriter, r *http.Request) {

	inUTL := r.URL.Path
	if !(inUTL == "/favicon.ico") {
		tmpl := "home"
		fmt.Println("WCT : Serving :", inUTL)
		//		var propertiesFileName = "config/properties.cfg"
		wctProperties := getProperties(CONST_CONFIG_FILE)

		w.Header().Set("Content-Type", "text/html")

		title := wctProperties["appname"]

		//p, _ := loadHomePage(title)

		noResp, _, respList := listResponseswebNew(wctProperties, "json", w)

		noServices, servicesList, serviceCatalog := getServices(wctProperties, "json")

		p := HomePage{Title: title, Body: "", RequestPath: wctProperties["deliverpath"], ResponsePath: wctProperties["receivepath"], ProcessedPath: wctProperties["processedpath"], NoResponses: noResp, Responses: respList, NoServices: noServices, Services: servicesList, ServiceCatalog: serviceCatalog, Description: "A description of the homepage.", ResponseType: wctProperties["responseformat"]}

		//fmt.Println("serviceCatalog", serviceCatalog)
		//fmt.Println("HomePage=", p.ServiceCatalog)

		t, _ := template.ParseFiles(getTemplateID(tmpl))
		t.Execute(w, p)
	}
}
