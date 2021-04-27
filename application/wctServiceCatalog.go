package application

import (
	"html"
	"net/http"
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

func GetServices(unused map[string]string, responseFormat string, r *http.Request) (int, string, []ServiceCatalogItem) {

	// serviceCatalog is an array of Service Catalog Items
	var serviceCatalog []ServiceCatalogItem

	id := uuid.New()

	resp := BuildRequestMessage(id.String(), "SERVICES", "", "", "", GetUserSessionToken(r))

	deliverRequest(resp, id.String())

	var servicesList string

	catalogResponse := GetResponseAsync(id.String(), r)

	noServices, _ := strconv.Atoi(catalogResponse.ResponseContentCount)
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
