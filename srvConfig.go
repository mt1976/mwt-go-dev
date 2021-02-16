package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

//ConfigViewPage is cheese
type ConfigViewPage struct {
	Title       string
	PageTitle   string
	ConfigItems []ConfigItem
}

//ConfigItem is cheese
type ConfigItem struct {
	ItemID    int
	ItemLabel string
	ItemValue string
}

func viewSrvConfigHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "viewSrvConfig"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	requestID := uuid.New()

	fmt.Println("WCT : Serving :", inUTL)

	title := wctProperties["appname"]

	// Get Data Here
	//_, _, serviceCatalog := getServices(wctProperties, wctProperties["responseformat"])
	requestMessage := buildRequestMessage(requestID.String(), "@CONFIG", "", "", wctProperties)

	fmt.Println("requestMessage", requestMessage)
	fmt.Println("SEND MESSAGE")
	sendRequest(requestMessage, requestID.String(), wctProperties)

	responseMessge := waitForResponse(requestID.String(), wctProperties)
	fmt.Println("responseMessge", responseMessge)

	//outString := ""
	noRows := len(responseMessge.ResponseContent.ResponseContentRow)

	var configsList []ConfigItem

	for ii := 0; ii < noRows; ii++ {
		var fred ConfigItem
		serviceContent := strings.Split(responseMessge.ResponseContent.ResponseContentRow[ii], "|")
		fred.ItemID = ii
		fred.ItemLabel = serviceContent[0]
		fred.ItemValue = serviceContent[1]
		configsList = append(configsList, fred)
	}

	pageSrvConfigView := ConfigViewPage{
		Title:       title,
		PageTitle:   "View Server Config",
		ConfigItems: configsList,
	}

	fmt.Println("Page Data", pageSrvConfigView)

	//thisTemplate:= getTemplateID(tmpl)
	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSrvConfigView)

}
