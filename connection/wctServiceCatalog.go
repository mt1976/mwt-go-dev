package connection

import (
	"html"
	"strconv"
	"strings"
	"net/http"
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

func getServices(wctProperties map[string]string, responseFormat string, r *http.Request) (int, string, []ServiceCatalogItem) {

	// serviceCatalog is an array of Service Catalog Items
	var serviceCatalog []ServiceCatalogItem

	id := uuid.New()

	resp := buildRequestMessage(id.String(), "SERVICES", "", "", "", wctProperties)

	deliverRequest(resp, wctProperties["deliverpath"], id.String(), wctProperties["responseformat"])

	// Now we get the services array
	//var responseFileName = wctProperties["receivepath"] + "/" + id.String() + "." + responseFormat
	//var processedFileName = wctProperties["processedpath"] + "/" + id.String() + "." + responseFormat
	//fmt.Println("Response Filename :", responseFileName)
	//fmt.Println("Processed Filename :", processedFileName)

	var servicesList string

	catalogResponse := getResponseAsync(id.String(), wctProperties, r)



				noServices,_ := strconv.Atoi(catalogResponse.ResponseContentCount)
				for ii := 0; ii < noServices; ii++ {

					serviceContent := strings.Split(catalogResponse.ResponseContent.ResponseContentRow[ii], "|")

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



	return noServices, servicesList, serviceCatalog
}
