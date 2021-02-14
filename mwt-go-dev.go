package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/mbndr/figlet4go"

	"github.com/google/uuid"
	"github.com/jimlawless/cfg"
)

//WctMessage is cheese
type WctMessage struct {
	WctPayload WctPayload `json:"wctRequest"`
}

//WctPayload ...
type WctPayload struct {
	ApplicationToken      string `json:"appToken"`
	SessionToken          string `json:"sessionToken"`
	RequestID             string `json:"requestID"`
	RequestAction         string `json:"requestAction"`
	RequestItem           string `json:"requestItem"`
	RequestParameters     string `json:"requestParameters"`
	UniqueUID             string `json:"uniqueID"`
	RequestResponseFormat string `json:"requestResponseFormat"`
}

//ServiceCatalogItem ...
type ServiceCatalogItem struct {
	Text       string
	Action     string
	Item       string
	Parameters string
	Helptext   string
	isTitle    bool
	CatalogID  int
}

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

//WctResponseMessage is cheese
type WctResponseMessage struct {
	WctReponsePayload WctResponsePayload `json:"wctResponse"`
}

//WctResponsePayload is cheese
type WctResponsePayload struct {
	RequestID            string `json:"requestID"`
	RequestFileName      string `json:"requestFileName"`
	RequestFilePath      string `json:"requestFilePath"`
	RequestFileType      string `json:"requestFileType"`
	RequestConsumed      string `json:"requestConsumed"`
	RequestAction        string `json:"requestAction"`
	RequestItem          string `json:"requestItem"`
	RequestParameters    string `json:"requestParameters"`
	ResponseID           string `json:"responseID"`
	ResponseFileName     string `json:"responseFileName"`
	ResponseFilePath     string `json:"responseFilePath"`
	ResponseFileType     string `json:"responseFileType"`
	ResponseEjected      string `json:"responseEjected"`
	ResponseContentCount string `json:"responseContentCount"`
	ResponseStatus       string `json:"responseStatus"`
	ResponseContent      struct {
		ResponseContentRow []string `json:"responseContentRow"`
	} `json:"responseContent"`
	RequestPayloadCount string `json:"requestPayloadCount"`
	RequestPayload      struct {
		RequestPayloadRow []string `json:"requestPayloadRow"`
	} `json:"requestPayload"`
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
	fmt.Println("URL", "http://localhost:"+wctProperties["port"])
	httpPort := ":" + wctProperties["port"]
	http.ListenAndServe(httpPort, nil)

}

func listResponsescli(wctProperties map[string]string, responseFormat string) {
	files := getResponses(wctProperties["receivepath"], responseFormat)
	fmt.Println("Responses Found :", len(files))
	for _, f := range files {
		fmt.Println("ID :", f)
	}
}

func listResponsesweb(wctProperties map[string]string, responseFormat string, w http.ResponseWriter) {

	files := getResponses(wctProperties["receivepath"], responseFormat)
	var noFound = "<em>Responses Found : </em><code>" + strconv.Itoa(len(files)) + "</code>"
	fmt.Fprintf(w, "%s", "<hr>")
	fmt.Fprintf(w, "%s", noFound)
	fmt.Fprintf(w, "%s", "<hr>")
	for _, f := range files {
		var outString = "<em>ID : </em><code>" + f + "</code><br>"
		fmt.Fprintf(w, "%s", outString)
	}
}

func getResponses(path string, responseFormat string) []string {
	files, err := filepath.Glob(path + "/*." + responseFormat)
	if err != nil {
		log.Fatal(err)
	}
	return files
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

func getProperties(propFile string) map[string]string {
	wctProperties := make(map[string]string)
	err := cfg.Load(propFile, wctProperties)
	if err != nil {
		log.Fatal(err)
	}
	return wctProperties
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
		fmt.Println("HomePage=", p.ServiceCatalog)
		renderTemplate(w, "home", p)
	}
}

func listResponseswebNew(wctProperties map[string]string, responseFormat string, w http.ResponseWriter) (int, string) {

	files := getResponses(wctProperties["receivepath"], responseFormat)
	//	var noFound = "<em>Responses Found : </em><code>" + strconv.Itoa(len(files)) + "</code>"
	noResponses := len(files)
	var responseText bytes.Buffer
	for _, f := range files {
		responseText.WriteString(f + "\n")
	}
	return noResponses, responseText.String()
}

func getServices(wctProperties map[string]string, responseFormat string) (int, string, []ServiceCatalogItem) {

	// serviceCatalog is an array of Service Catalog Items
	var serviceCatalog []ServiceCatalogItem

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

	// Now we get the services array
	var responseFileName = wctProperties["receivepath"] + "/" + id.String() + "." + responseFormat
	var processedFileName = wctProperties["processedpath"] + "/" + id.String() + "." + responseFormat
	fmt.Println("Response Filename :", responseFileName)
	fmt.Println("Processed Filename :", processedFileName)

	//content := ""
	condition := false
	//	text := ""
	noServices := 999
	servicesList := ""
	var wibble WctResponseMessage
	//	var nibble WctResponsePayload
	for !condition {
		fmt.Println("Polling file", responseFileName)
		content, _ := ioutil.ReadFile(responseFileName)
		text := string(content)
		//	fmt.Println("text file", text)
		if text != "" {
			condition = true
			if text != "" {
				json.Unmarshal(content, &wibble)
				//			fmt.Println("responseContent", wibble.WctReponsePayload.ResponseContent)
				//			fmt.Println("responseContentCount", wibble.WctReponsePayload.ResponseContentCount)
				//servicesList = wibble["reponsepayloadcount"]
				var x = wibble.WctReponsePayload.ResponseContentCount
				noServices, _ = strconv.Atoi(x)
				for ii := 0; ii < noServices; ii++ {
					//servicesList += wibble.WctReponsePayload.ResponseContent.ResponseContentRow[ii] + "\n"
					serviceContent := strings.Split(wibble.WctReponsePayload.ResponseContent.ResponseContentRow[ii], "|")

					var item ServiceCatalogItem
					item.Text = serviceContent[0]
					item.Action = serviceContent[1]
					item.Item = serviceContent[2]
					item.Parameters = serviceContent[3]
					item.Helptext = serviceContent[4]
					if item.Action != "" {
						item.isTitle = false
						servicesList += "* " + item.Text + "\n"
					} else {
						item.isTitle = true
						servicesList += item.Text + "\n"
					}
					item.CatalogID = ii

					//fmt.Println("CatalogItem", item)

					serviceCatalog = append(serviceCatalog, item)
				}
				//servicesList = wibble.WctReponsePayload.ResponseContent.ResponseContentRow[1]
				//fmt.Println(servicesList)
			}
			fmt.Println("Delete file", responseFileName)

			var err = os.Remove(responseFileName)
			if err != nil {
				fmt.Println(err.Error())
			}
			err = ioutil.WriteFile(processedFileName, content, 0644)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {

			pollingInterval, _ := strconv.Atoi(wctProperties["pollinginterval"])
			fmt.Println("Snoooze... Zzzzzz.... ", pollingInterval)
			time.Sleep(time.Duration(pollingInterval) * time.Second)
		}
		//fmt.Println(text)
	}

	return noServices, servicesList, serviceCatalog
}
