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
	//	asciiOptions := figlet4go.NewRenderOptions()

	//var propertiesFileName = getAppPropertiesFile
	//var responseFormat = "json"
	//	wctProperties := make(map[string]string)
	wctProperties := getProperties(CONST_CONFIG_FILE)
	//fmt.Println(wctProperties)

	// The underscore would be an error
	renderStr, _ := ascii.Render("MWT GO PROTO")
	//	asciiOptions.FontName = "standard"
	fmt.Print(renderStr)

	fmt.Println("Delivery Path :", wctProperties["deliverpath"])
	fmt.Println("Reponse Path :", wctProperties["receivepath"])

	//fmt.Println(id.String())
	//processSimpleRequestMessage(wctProperties, responseFormat) //Calling processSimpleRequestMessage
	//	processSimpleRequestMessage(wctProperties, responseFormat) //Calling processSimpleRequestMessage

	//home := wctProperties["receivepath"]

	//listResponsescli(wctProperties, responseFormat) //Call listResponses

	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/previewRequest/", previewRequestHandler)
	http.HandleFunc("/executeRequest/", executeRequestHandler)
	http.HandleFunc("/viewResponse/", viewResponseHandler)
	http.HandleFunc("/deleteResponse/", deleteResponseHandler)
	http.HandleFunc("/clearQueues/", clearQueuesHandler)
	http.HandleFunc("/viewSrvConfig", viewSrvConfigHandler)

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
