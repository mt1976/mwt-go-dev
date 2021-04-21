package siena

import (
	"fmt"
	"html"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	support "github.com/mt1976/mwt-go-dev/appsupport"
)

//SvcDataMapPage is cheese
type SvcDataMapPage struct {
	UserMenu        []AppMenuItem
	UserRole        string
	UserNavi        string
	Title           string
	PageTitle       string
	NoDataMapIDs    int
	SvcDataMapItems []SvcDataMapItem
	SvcDataMapCols  []DataHdr
	DataMapPageID   string
	DataRows        []DataRow
	JSRows          int
	JSCols          int
	FullRecord      string
}

//DataRow is cheese
type DataRow struct {
	RowID       int
	DataRowItem []DataCol
}

//DataHdr is cheese
type DataHdr struct {
	ColID       int
	DataHdrItem string
}

//DataCol is cheese
type DataCol struct {
	DataItem string
	DIrow    int
	DIcol    int
}

//SvcDataMapItem is cheese
type SvcDataMapItem struct {
	DataMapID          int
	DataMapName        string
	DataMapFileID      string
	DataMapDescription string
	DataMapXMLFile     string
}

func listSvcDataMapHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	tmpl := "listSvcDataMap"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	requestID := uuid.New()

	log.Println("Servicing :", inUTL)

	title := wctProperties["appname"]

	// Get Data Here
	//_, _, serviceCatalog := getServices(wctProperties, wctProperties["responseformat"])
	requestMessage := buildRequestMessage(requestID.String(), "@DATAMAP", "LIST", "", "", wctProperties)

	//fmt.Println("requestMessage", requestMessage)
	//fmt.Println("SEND MESSAGE")
	sendRequest(requestMessage, requestID.String(), wctProperties)

	responseMessage := getResponseAsync(requestID.String(), wctProperties, r)
	//fmt.Println("responseMessage", responseMessage)

	//outString := ""
	noRows := len(responseMessage.ResponseContent.ResponseContentRow)

	var dataMapItemsList []SvcDataMapItem

	for ii := 0; ii < noRows; ii++ {
		var newDataMapItem SvcDataMapItem
		dataMapContentRow := strings.Split(responseMessage.ResponseContent.ResponseContentRow[ii], "|")
		newDataMapItem.DataMapID = ii
		newDataMapItem.DataMapName = dataMapContentRow[1]
		newDataMapItem.DataMapFileID = dataMapContentRow[2]
		newDataMapItem.DataMapDescription = html.UnescapeString(dataMapContentRow[3])
		newDataMapItem.DataMapXMLFile = dataMapContentRow[4]
		//fmt.Println("newDataMapItem", newDataMapItem)
		dataMapItemsList = append(dataMapItemsList, newDataMapItem)
	}

	pageSrvEvironment := SvcDataMapPage{
		UserMenu:        getappMenuData(gUserRole),
		UserRole:        gUserRole,
		UserNavi:        gUserNavi,
		Title:           title,
		PageTitle:       "List Data Maps",
		NoDataMapIDs:    noRows,
		SvcDataMapItems: dataMapItemsList,
	}

	//fmt.Println("Page Data", pageSrvEvironment)

	//thisTemplate:= support.GetTemplateID(tmpl,gUserRole)
	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, gUserRole))
	t.Execute(w, pageSrvEvironment)

}

func viewSvcDataMapHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	tmpl := "viewSvcDataMap"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	requestID := uuid.New()

	thisID := support.GetURLparam(r, "dataMapName")
	//fmt.Println(thisID)

	log.Println("Servicing :", inUTL)

	title := wctProperties["appname"]

	// Get Data Here
	//_, _, serviceCatalog := getServices(wctProperties, wctProperties["responseformat"])
	requestMessage := buildRequestMessage(requestID.String(), "@DATAMAP", "VIEW", thisID, "", wctProperties)

	//fmt.Println("requestMessage", requestMessage)
	//fmt.Println("SEND MESSAGE")
	sendRequest(requestMessage, requestID.String(), wctProperties)

	responseMessage := getResponseAsync(requestID.String(), wctProperties, r)
	//fmt.Println("*** GOT RESPONSE ***")
	//fmt.Println("responseMessage", responseMessage)

	//outString := ""
	noRows := len(responseMessage.ResponseContent.ResponseContentRow)

	var wrkDataMapCols []DataHdr
	var wrkDataMapRows []DataRow
	//wrkDataMapCols[1] = "POO"
	//fmt.Println("defined wrkDataMapCols", wrkDataMapCols)
	var firstDataMapContentRow = strings.Split(responseMessage.ResponseContent.ResponseContentRow[0], "|")
	//fmt.Println("Calculating Columns")
	noCols := len(firstDataMapContentRow)
	//fmt.Println("firstDataMapContentRow=", firstDataMapContentRow)
	//fmt.Println("noCols=", noCols)
	for jj := 0; jj < noCols; jj++ {
		headerVal := firstDataMapContentRow[jj]
		var tmpVal DataHdr
		tmpVal.ColID = jj
		tmpVal.DataHdrItem = headerVal
		wrkDataMapCols = append(wrkDataMapCols, tmpVal)
	}
	//fmt.Println("defined wrkDataMapCols", wrkDataMapCols)

	// Deliberatly skip the first row
	for kk := 1; kk < noRows; kk++ {

		var theDataMapContentRow = strings.Split(responseMessage.ResponseContent.ResponseContentRow[kk], "|")

		//fmt.Println("theDataMapContentRow", theDataMapContentRow)

		var wrkDataMapColItems []DataCol
		for jj := 0; jj < noCols; jj++ {
			var wrk DataCol
			//fmt.Println("theDataMapContentRow", kk, jj, theDataMapContentRow[jj])
			wrk.DataItem = theDataMapContentRow[jj]
			wrk.DIcol = jj
			wrk.DIrow = kk
			wrkDataMapColItems = append(wrkDataMapColItems, wrk)
		}
		//fmt.Println("wrkDataMapColItems", kk, wrkDataMapColItems)
		var dr DataRow
		dr.RowID = kk
		dr.DataRowItem = wrkDataMapColItems
		//headerVal := firstDataMapContentRow[jj]
		wrkDataMapRows = append(wrkDataMapRows, dr)
	}
	//fmt.Println("wrkDataMapRows", wrkDataMapRows)

	pageSrvEvironment := SvcDataMapPage{
		UserMenu:       getappMenuData(gUserRole),
		UserRole:       gUserRole,
		UserNavi:       gUserNavi,
		Title:          title,
		PageTitle:      "View Data Maps",
		NoDataMapIDs:   0,
		SvcDataMapCols: wrkDataMapCols,
		DataMapPageID:  thisID,
		DataRows:       wrkDataMapRows,
		JSCols:         noCols - 1,
		JSRows:         noRows - 1,
	}

	//fmt.Println("Page Data", pageSrvEvironment)

	//thisTemplate:= support.GetTemplateID(tmpl,gUserRole)
	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, gUserRole))
	t.Execute(w, pageSrvEvironment)

}

func editSvcDataMapHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	tmpl := "editSvcDataMap"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	requestID := uuid.New()

	thisID := support.GetURLparam(r, "dataMapName")
	//fmt.Println(thisID)

	log.Println("Servicing :", inUTL)

	title := wctProperties["appname"]

	// Get Data Here
	//_, _, serviceCatalog := getServices(wctProperties, wctProperties["responseformat"])
	requestMessage := buildRequestMessage(requestID.String(), "@DATAMAP", "RAW", thisID, "", wctProperties)

	//fmt.Println("requestMessage", requestMessage)
	//fmt.Println("SEND MESSAGE")
	sendRequest(requestMessage, requestID.String(), wctProperties)

	responseMessage := getResponseAsync(requestID.String(), wctProperties, r)
	//fmt.Println("*** GOT RESPONSE ***")
	fmt.Println("responseMessage", responseMessage)

	//outString := ""
	noRows := len(responseMessage.ResponseContent.ResponseContentRow)

	fullRec := strings.Join(responseMessage.ResponseContent.ResponseContentRow, " \n")

	pageSrvEvironment := SvcDataMapPage{
		UserMenu:      getappMenuData(gUserRole),
		UserRole:      gUserRole,
		UserNavi:      gUserNavi,
		Title:         title,
		PageTitle:     "Edit Data Map",
		DataMapPageID: thisID,
		JSRows:        noRows - 1,
		FullRecord:    fullRec,
	}

	//fmt.Println("Page Data", pageSrvEvironment)

	//thisTemplate:= support.GetTemplateID(tmpl,gUserRole)
	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, gUserRole))
	t.Execute(w, pageSrvEvironment)

}

func viewSvcDataMapXMLHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	tmpl := "viewSvcDataMapXML"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	requestID := uuid.New()

	thisID := support.GetURLparam(r, "dataMapName")
	//fmt.Println(thisID)

	log.Println("Servicing :", inUTL)

	title := wctProperties["appname"]

	// Get Data Here
	//_, _, serviceCatalog := getServices(wctProperties, wctProperties["responseformat"])
	requestMessage := buildRequestMessage(requestID.String(), "@DATAXML", "VIEW", thisID, "", wctProperties)

	//fmt.Println("requestMessage", requestMessage)
	//fmt.Println("SEND MESSAGE")
	sendRequest(requestMessage, requestID.String(), wctProperties)

	responseMessage := getResponseAsync(requestID.String(), wctProperties, r)
	//fmt.Println("*** GOT RESPONSE ***")
	fmt.Println("responseMessage", responseMessage)

	//outString := ""
	noRows := len(responseMessage.ResponseContent.ResponseContentRow)

	fullRec := strings.Join(responseMessage.ResponseContent.ResponseContentRow, " \n")

	pageSrvEvironment := SvcDataMapPage{
		UserMenu:      getappMenuData(gUserRole),
		UserRole:      gUserRole,
		UserNavi:      gUserNavi,
		Title:         title,
		PageTitle:     "View XML Template",
		DataMapPageID: thisID,
		JSRows:        noRows - 1,
		FullRecord:    html.UnescapeString(fullRec),
	}

	//fmt.Println("Page Data", pageSrvEvironment)

	//thisTemplate:= support.GetTemplateID(tmpl,gUserRole)
	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, gUserRole))
	t.Execute(w, pageSrvEvironment)

}

func editSvcDataMapXMLHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	tmpl := "editSvcDataMapXML"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	requestID := uuid.New()

	thisID := support.GetURLparam(r, "dataMapName")
	//fmt.Println(thisID)

	log.Println("Servicing :", inUTL)

	title := wctProperties["appname"]

	// Get Data Here
	//_, _, serviceCatalog := getServices(wctProperties, wctProperties["responseformat"])
	requestMessage := buildRequestMessage(requestID.String(), "@DATAXML", "VIEW", thisID, "", wctProperties)

	//fmt.Println("requestMessage", requestMessage)
	//fmt.Println("SEND MESSAGE")
	sendRequest(requestMessage, requestID.String(), wctProperties)

	responseMessage := getResponseAsync(requestID.String(), wctProperties, r)
	//fmt.Println("*** GOT RESPONSE ***")
	fmt.Println("responseMessage", responseMessage)

	//outString := ""
	noRows := len(responseMessage.ResponseContent.ResponseContentRow)

	fullRec := strings.Join(responseMessage.ResponseContent.ResponseContentRow, " \n")
	fullRec = html.UnescapeString(fullRec)
	pageEditSvcDataMapXML := SvcDataMapPage{
		UserMenu:      getappMenuData(gUserRole),
		UserRole:      gUserRole,
		UserNavi:      gUserNavi,
		Title:         title,
		PageTitle:     "View XML Template",
		DataMapPageID: thisID,
		JSRows:        noRows - 1,
		FullRecord:    html.UnescapeString(fullRec),
	}

	fmt.Println("Page Data", pageEditSvcDataMapXML)

	//thisTemplate:= support.GetTemplateID(tmpl,gUserRole)
	t, err := template.ParseFiles(support.GetTemplateID(tmpl, gUserRole))
	fmt.Println("error", err)
	t.Execute(w, pageEditSvcDataMapXML)

}

func saveSvcDataMapHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	//	tmpl := "editSrvConfiguration"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")

	requestID := uuid.New()
	//	maxRows, _ := strconv.Atoi(wctProperties["maxtextboxrows"])
	//recordID := support.GetURLparam(r, "pgid")
	//recordContent := support.GetURLparam(r, "pgContent")

	log.Println("Servicing :", inUTL)

	title := wctProperties["appname"]

	body := r.FormValue("pgContent")
	id := r.FormValue("pgid")

	//p := &Page{Title: title, Body: []byte(body)}
	//	fmt.Println("METHOD",r.Method)
	fmt.Println("TITLE", title)
	//	fmt.Println("ID", recordID)
	fmt.Println("ID", id)
	fmt.Println("BODY", body)
	//fmt.Println("REC", recordContent)
	//	fmt.Println("ARSE", r)
	//	fmt.Println("parse",r.ParseForm())

	requestMessage := buildRequestMessage(requestID.String(), "@DATAMAP", "SAVE", id, body, wctProperties)

	fmt.Println("requestMessage", requestMessage)
	fmt.Println("SEND MESSAGE")
	sendRequest(requestMessage, requestID.String(), wctProperties)

	listSvcDataMapHandler(w, r)

	// Get Data Here
	//_, _, serviceCatalog := getServices(wctProperties, wctProperties["responseformat"])

	//	t, _ := template.ParseFiles(support.GetTemplateID(tmpl,gUserRole))
	//	t.Execute(w, pageSrvConfigView)

}

func saveSvcDataMapXMLHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	//	tmpl := "editSrvConfiguration"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")

	requestID := uuid.New()
	//	maxRows, _ := strconv.Atoi(wctProperties["maxtextboxrows"])
	//recordID := support.GetURLparam(r, "pgid")
	//recordContent := support.GetURLparam(r, "pgContent")

	log.Println("Servicing :", inUTL)

	title := wctProperties["appname"]

	body := r.FormValue("pgContent")
	id := r.FormValue("pgid")

	//p := &Page{Title: title, Body: []byte(body)}
	//	fmt.Println("METHOD",r.Method)
	fmt.Println("TITLE", title)
	//	fmt.Println("ID", recordID)
	fmt.Println("ID", id)
	fmt.Println("BODY", body)
	//fmt.Println("REC", recordContent)
	//	fmt.Println("ARSE", r)
	//	fmt.Println("parse",r.ParseForm())

	requestMessage := buildRequestMessage(requestID.String(), "@DATAXML", "SAVE", id, body, wctProperties)

	fmt.Println("requestMessage", requestMessage)
	fmt.Println("SEND MESSAGE")
	sendRequest(requestMessage, requestID.String(), wctProperties)

	listSvcDataMapHandler(w, r)

	// Get Data Here
	//_, _, serviceCatalog := getServices(wctProperties, wctProperties["responseformat"])

	//	t, _ := template.ParseFiles(support.GetTemplateID(tmpl,gUserRole))
	//	t.Execute(w, pageSrvConfigView)

}

func newSvcDataMapHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	tmpl := "newSvcDataMap"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")

	log.Println("Servicing :", inUTL)

	title := wctProperties["appname"]

	pageDM := SvcDataMapPage{
		UserMenu:  getappMenuData(gUserRole),
		UserRole:  gUserRole,
		UserNavi:  gUserNavi,
		Title:     title,
		PageTitle: "New Data Loader Template",
	}
	fmt.Println("WCT : Page :", pageDM)

	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, gUserRole))
	t.Execute(w, pageDM)

}

func genSvcDataMapHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	//	tmpl := "editSrvConfiguration"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")

	requestID := uuid.New()
	//	maxRows, _ := strconv.Atoi(wctProperties["maxtextboxrows"])
	//recordID := support.GetURLparam(r, "pgid")
	//recordContent := support.GetURLparam(r, "pgContent")

	log.Println("Servicing :", inUTL)

	title := wctProperties["appname"]

	body := r.FormValue("descr")
	id := r.FormValue("name")

	//p := &Page{Title: title, Body: []byte(body)}
	//	fmt.Println("METHOD",r.Method)
	fmt.Println("TITLE", title)
	//	fmt.Println("ID", recordID)
	fmt.Println("ID", id)
	fmt.Println("BODY", body)
	//fmt.Println("REC", recordContent)
	//	fmt.Println("ARSE", r)
	//	fmt.Println("parse",r.ParseForm())

	requestMessage := buildRequestMessage(requestID.String(), "@DATAMAP", "NEW", id, body, wctProperties)

	fmt.Println("requestMessage", requestMessage)
	fmt.Println("SEND MESSAGE")

	sendRequest(requestMessage, requestID.String(), wctProperties)

	listSvcDataMapHandler(w, r)

	// Get Data Here
	//_, _, serviceCatalog := getServices(wctProperties, wctProperties["responseformat"])

	//	t, _ := template.ParseFiles(support.GetTemplateID(tmpl,gUserRole))
	//	t.Execute(w, pageSrvConfigView)

}

func deleteSvcDataMapHandler(w http.ResponseWriter, r *http.Request) {
	wctProperties := support.GetProperties(APPCONFIG)
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	requestID := uuid.New()
	log.Println("Servicing :", inUTL)
	id := support.GetURLparam(r, "dataMapName")
	requestMessage := buildRequestMessage(requestID.String(), "@DATAMAP", "DELETE", id, "", wctProperties)
	fmt.Println("requestMessage", requestMessage)
	fmt.Println("SEND MESSAGE")
	sendRequest(requestMessage, requestID.String(), wctProperties)
	listSvcDataMapHandler(w, r)
}
