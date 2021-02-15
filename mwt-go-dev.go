package main

import (
	"fmt"
	"net/http"

	"github.com/mbndr/figlet4go"

	"github.com/google/uuid"
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

	var propertiesFileName = "config/wct_Properties.cfg"
	//var responseFormat = "json"
	//	wctProperties := make(map[string]string)
	wctProperties := getProperties(propertiesFileName)
	//fmt.Println(wctProperties)

	// The underscore would be an error
	renderStr, _ := ascii.Render("MWT GO PROTO")
	//	asciiOptions.FontName = "standard"
	fmt.Print(renderStr)

	fmt.Println("Delivery Path :", wctProperties["deliverpath"])
	fmt.Println("Reponse Path :", wctProperties["receivepath"])

	//fmt.Println(id.String())
	//generateMessage(wctProperties, responseFormat) //Calling generateMessage
	//	generateMessage(wctProperties, responseFormat) //Calling generateMessage

	//home := wctProperties["receivepath"]

	//listResponsescli(wctProperties, responseFormat) //Call listResponses

	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/previewRequest/", previewRequestHandler)
	http.HandleFunc("/executeRequest/", executeRequestHandler)
	http.HandleFunc("/viewResponse/", viewResponseHandler)

	fmt.Println("URL", "http://localhost:"+wctProperties["port"])
	httpPort := ":" + wctProperties["port"]
	http.ListenAndServe(httpPort, nil)

}

func generateMessage(wctProperties map[string]string, responseFormat string) {

	id := uuid.New()

	resp := WctMessage{
		WctPayload{
			ApplicationToken:      wctProperties["applicationtoken"],
			RequestID:             id.String(),
			RequestAction:         "SERVICES",
			UniqueUID:             wctProperties["appid"],
			RequestResponseFormat: responseFormat,
		},
	}

	deliverRequest(resp, wctProperties["deliverpath"], id.String(), responseFormat)

}

func homePageHandler(w http.ResponseWriter, r *http.Request) {

	inUTL := r.URL.Path
	if !(inUTL == "/favicon.ico") {

		fmt.Println("WCT : Serving :", inUTL)
		var propertiesFileName = "config/wct_Properties.cfg"
		wctProperties := getProperties(propertiesFileName)

		w.Header().Set("Content-Type", "text/html")

		title := wctProperties["appname"]

		//p, _ := loadHomePage(title)

		noResp, _, respList := listResponseswebNew(wctProperties, "json", w)

		noServices, servicesList, serviceCatalog := getServices(wctProperties, "json")

		p := HomePage{Title: title, Body: "", RequestPath: wctProperties["deliverpath"], ResponsePath: wctProperties["receivepath"], ProcessedPath: wctProperties["processedpath"], NoResponses: noResp, Responses: respList, NoServices: noServices, Services: servicesList, ServiceCatalog: serviceCatalog, Description: "A description of the homepage.", ResponseType: wctProperties["responseformat"]}

		//fmt.Println("serviceCatalog", serviceCatalog)
		//fmt.Println("HomePage=", p.ServiceCatalog)
		renderTemplate(w, "home", p)
	}
}
