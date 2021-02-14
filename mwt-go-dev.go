package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	Responses      string
	NoServices     int
	Services       string
	ServiceCatalog []ServiceCatalogItem
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
	js, _ := json.Marshal(resp)

	//	fmt.Printf("\n")
	//	fmt.Printf("%s", js)
	//	fmt.Printf("\n")

	var fileName = wctProperties["deliverpath"] + "/" + id.String() + "." + responseFormat
	fmt.Println("Request Filename :", fileName)
	//	fmt.Printf("\n")
	_ = ioutil.WriteFile(fileName, js, 0644)
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

		noResp, respText := listResponseswebNew(wctProperties, "json", w)

		noServices, servicesList, serviceCatalog := getServices(wctProperties, "json")

		p := HomePage{Title: title, Body: "", RequestPath: wctProperties["deliverpath"], ResponsePath: wctProperties["receivepath"], ProcessedPath: wctProperties["processedpath"], NoResponses: noResp, Responses: respText, NoServices: noServices, Services: servicesList, ServiceCatalog: serviceCatalog}

		//fmt.Println("serviceCatalog", serviceCatalog)
		//fmt.Println("HomePage=", p.ServiceCatalog)
		renderTemplate(w, "home", p)
	}
}
