package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

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

	listResponses(wctProperties, responseFormat) //Call listResponses

	fmt.Printf("\n")
	fmt.Println("DONE!")
	fmt.Printf("\n")

}

func listResponses(wctProperties map[string]string, responseFormat string) {
	files, err := filepath.Glob(wctProperties["receivepath"] + "/*." + responseFormat)

	if err != nil {

		log.Fatal(err)
	}

	fmt.Println("Responses Found :", len(files))

	for _, f := range files {

		fmt.Println("ID :", f)

	}
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
