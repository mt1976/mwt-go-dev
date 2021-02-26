package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

//ServiceCatalogItem ...
type ServiceCatalogItem struct {
	Text       string
	Action     string
	Item       string
	Parameters string
	Helptext   string
	isTitle    bool
	CatalogID  int
	UUID       string
}

func getServices(wctProperties map[string]string, responseFormat string) (int, string, []ServiceCatalogItem) {

	// serviceCatalog is an array of Service Catalog Items
	var serviceCatalog []ServiceCatalogItem

	id := uuid.New()

	resp := buildRequestMessage(id.String(), "SERVICES", "", "", "", wctProperties)

	deliverRequest(resp, wctProperties["deliverpath"], id.String(), wctProperties["responseformat"])

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

				var x = wibble.WctReponsePayload.ResponseContentCount
				noServices, _ = strconv.Atoi(x)
				for ii := 0; ii < noServices; ii++ {

					serviceContent := strings.Split(wibble.WctReponsePayload.ResponseContent.ResponseContentRow[ii], "|")

					var item ServiceCatalogItem
					item.Text = serviceContent[0]
					item.Action = serviceContent[1]
					item.Item = serviceContent[2]
					item.Parameters = serviceContent[3]
					item.Helptext = html.UnescapeString(serviceContent[4])
					if item.Action != "" {
						item.isTitle = false
						servicesList += "* " + item.Text + "\n"
					} else {
						item.isTitle = true
						servicesList += item.Text + "\n"
					}
					item.CatalogID = ii
					requestID := uuid.New()
					item.UUID = requestID.String()
					//fmt.Println("CatalogItem", item)

					serviceCatalog = append(serviceCatalog, item)
				}

			}

			err := deleteResponse(id.String(), wctProperties)
			if err != nil {
				fmt.Println(err.Error())
			}

			err = ioutil.WriteFile(processedFileName, content, 0644)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {

			doSnooze(wctProperties["pollinginterval"])

		}
		//fmt.Println(text)
	}

	return noServices, servicesList, serviceCatalog
}
