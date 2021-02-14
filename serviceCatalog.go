package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

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
