package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/mbndr/figlet4go"

	"github.com/google/uuid"
	"github.com/jimlawless/cfg"
)

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

//Page ...
type Page struct {
	Title        string
	Body         string
	RequestPath  string
	ResponsePath string
	NoResponses  int
	Responses    string
}

//WctMessage is cheese
type WctMessage struct {
	WctPayload WctPayload `json:"wctRequest"`
}

func main() {

	ascii := figlet4go.NewAsciiRender()
	//	asciiOptions := figlet4go.NewRenderOptions()

	var propertiesFileName = "wct_Properties.cfg"
	var responseFormat = "json"
	//	wctProperties := make(map[string]string)
	wctProperties := getProperties(propertiesFileName)
	//fmt.Println(wctProperties)

	// The underscore would be an error
	renderStr, _ := ascii.Render("MWT GO PROTO")
	//	asciiOptions.FontName = "standard"
	fmt.Print(renderStr)

	fmt.Println("Delivery Path :", wctProperties["deliverpath"])
	fmt.Println("Responce Path :", wctProperties["receivepath"])

	//fmt.Println(id.String())
	generateMessage(wctProperties, responseFormat) //Calling generateMessage
	generateMessage(wctProperties, responseFormat) //Calling generateMessage

	//home := wctProperties["receivepath"]

	listResponsescli(wctProperties, responseFormat) //Call listResponses

	fmt.Printf("\n")
	fmt.Println("DONE!")
	fmt.Printf("\n")

	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":8080", nil)

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

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	inUTL := r.URL.Path
	if !(inUTL == "/favicon.ico") {

		fmt.Println("WCT : Serving :", inUTL)

		var propertiesFileName = "wct_Properties.cfg"
		wctProperties := getProperties(propertiesFileName)

		w.Header().Set("Content-Type", "text/html")

		title := wctProperties["appname"]
		//p, _ := loadPage(title)

		noResp, respText := listResponseswebNew(wctProperties, "json", w)
		p := Page{Title: title, Body: "", RequestPath: wctProperties["deliverpath"], ResponsePath: wctProperties["receivepath"], NoResponses: noResp, Responses: respText}

		renderTemplate(w, "page", p)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, p Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
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
