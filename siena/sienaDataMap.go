package siena

import (
	"fmt"
	"html"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

//SvcDataMapPage is cheese
type SvcDataMapPage struct {
	UserMenu        []application.AppMenuItem
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
	DataMapID          string
	DataMapName        string
	DataMapFileID      string
	DataMapDescription string
	DataMapXMLFile     string
}

func ListSvcDataMapHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderList"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	//	requestID := uuid.New()

	application.ServiceMessage(inUTL)

	noRows, loaderList, _ := application.GetLoaderStoreList()

	var dataMapItemsList []SvcDataMapItem

	for _, listItem := range loaderList {
		var newDataMapItem SvcDataMapItem
		newDataMapItem.DataMapID = listItem.Id
		newDataMapItem.DataMapName = listItem.Name
		newDataMapItem.DataMapFileID = listItem.Filename
		newDataMapItem.DataMapDescription = listItem.Description
		newDataMapItem.DataMapXMLFile = listItem.Filename
		//fmt.Println("newDataMapItem", newDataMapItem)
		dataMapItemsList = append(dataMapItemsList, newDataMapItem)
	}

	pageSrvEvironment := SvcDataMapPage{
		UserMenu:        application.GetUserMenu(r),
		UserRole:        application.GetUserRole(r),
		UserNavi:        "NOT USED",
		Title:           globals.ApplicationProperties["appname"],
		PageTitle:       "List Data Maps",
		NoDataMapIDs:    noRows,
		SvcDataMapItems: dataMapItemsList,
	}

	//fmt.Println("Page Data", pageSrvEvironment)

	//thisTemplate:= application.GetTemplateID(tmpl,application.GetUserRole(r))
	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSrvEvironment)

}

func ViewSvcDataMapHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	buildGridPage("LoaderView", w, r)

}

func buildGridPage(tmpl string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	inUTL := r.URL.Path
	thisID := application.GetURLparam(r, "loaderID")
	application.ServiceMessage(inUTL)
	title := globals.ApplicationProperties["appname"]
	var wrkDataMapCols []DataHdr
	noColumns, wrkLoaderHeadersList, _ := application.GetLoaderMapStoreListByLoader(thisID)
	for _, colData := range wrkLoaderHeadersList {
		var tmpVal DataHdr
		tmpVal.ColID, _ = strconv.Atoi(colData.Position)
		tmpVal.DataHdrItem = colData.Name
		wrkDataMapCols = append(wrkDataMapCols, tmpVal)
	}

	var wrkDataMapRows []DataRow
	var wrk DataCol
	var dr DataRow

	noRows := application.GetLoaderDataStoreRowsByLoader(thisID)

	for thisRow := 0; thisRow < noRows; thisRow++ {
		var wrkDataMapColItems []DataCol

		noItems, wrkLoaderRowsList, _ := application.GetLoaderDataStoreRowList(thisID, strconv.Itoa(thisRow+1))
		for thisCol := 0; thisCol < noColumns; thisCol++ {
			//		log.Println("R=", thisRow, "C=", thisCol)
			wrk.DataItem = ""
			//	log.Println(thisCol, thisRow, noItems)
			if thisCol < noItems {
				var thisColData = wrkLoaderRowsList[thisCol]
				wrk.DataItem = thisColData.Value
			}
			wrk.DIcol = thisCol
			wrk.DIrow = thisRow
			wrkDataMapColItems = append(wrkDataMapColItems, wrk)
		}
		//	log.Println("WDMCOL=", wrkDataMapColItems)
		dr.RowID = thisRow
		dr.DataRowItem = wrkDataMapColItems
		wrkDataMapRows = append(wrkDataMapRows, dr)
	}

	pageSrvEvironment := SvcDataMapPage{
		UserMenu:       application.GetUserMenu(r),
		UserRole:       application.GetUserRole(r),
		UserNavi:       "NOT USED",
		Title:          title,
		PageTitle:      "View Data Maps",
		NoDataMapIDs:   0,
		SvcDataMapCols: wrkDataMapCols,
		DataMapPageID:  thisID,
		DataRows:       wrkDataMapRows,
		JSCols:         noColumns,
		JSRows:         noRows,
	}

	application.GetTemplateID(tmpl, application.GetUserRole(r))
	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSrvEvironment)
}

func getDataListFile(fileID string) string {
	content, err := ioutil.ReadFile(fileID)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	return string(content)
}

func EditSvcDataMapHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	buildGridPage("LoaderEdit", w, r)

}

func ViewSvcDataMapXMLHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderTemplateView"

	inUTL := r.URL.Path

	w.Header().Set("Content-Type", "text/html")

	thisID := application.GetURLparam(r, "loaderID")

	application.ServiceMessage(inUTL)

	text, _ := getXMLtemplateBody(thisID)

	pageSrvEvironment := SvcDataMapPage{
		UserMenu:      application.GetUserMenu(r),
		UserRole:      application.GetUserRole(r),
		UserNavi:      "NOT USED",
		Title:         "Title",
		PageTitle:     "View XML Template",
		DataMapPageID: thisID,
		JSRows:        35,
		FullRecord:    html.UnescapeString(text),
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSrvEvironment)

}

func getXMLtemplateBody(thisID string) (string, error) {
	path := globals.ApplicationProperties["datamaptemplatepath"]
	_, loaderItem, _ := application.GetLoaderStoreByID(thisID)
	fileName := loaderItem.Filename + ".xml"
	content, err := application.ReadDataFile(fileName, path)
	if err != nil {
		log.Fatal(err)
	}
	return content, err
}

func putXMLtemplateBody(thisID string, content string) int {
	path := globals.ApplicationProperties["datamaptemplatepath"]
	_, loaderItem, _ := application.GetLoaderStoreByID(thisID)
	fileName := loaderItem.Filename + ".xml"
	status := application.WriteDataFile(fileName, path, content)
	return status
}

func EditSvcDataMapXMLHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderTemplateEdit"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")

	thisID := application.GetURLparam(r, "loaderID")

	application.ServiceMessage(inUTL)

	title := globals.ApplicationProperties["appname"]

	// Get Data Here
	fullRec, _ := getXMLtemplateBody(thisID)

	pageEditSvcDataMapXML := SvcDataMapPage{
		UserMenu:      application.GetUserMenu(r),
		UserRole:      application.GetUserRole(r),
		UserNavi:      "NOT USED",
		Title:         title,
		PageTitle:     "View XML Template",
		DataMapPageID: thisID,
		JSRows:        35,
		FullRecord:    html.UnescapeString(fullRec),
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))

	//fmt.Println("error", err)
	t.Execute(w, pageEditSvcDataMapXML)

}

func SaveSvcDataMapHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	//	tmpl := "editSrvConfiguration"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")

	application.ServiceMessage(inUTL)

	tableColumns := r.FormValue("tableColumns")
	tableRows := r.FormValue("tableRows")
	loaderID := r.FormValue("loaderID")

	cols, _ := strconv.Atoi(tableColumns)
	rows, _ := strconv.Atoi(tableRows)

	cols = cols - 1
	rows = rows - 1

	application.DeleteLoaderDataStoreByLoader(loaderID)

	for thisRow := 0; thisRow <= rows; thisRow++ {
		for thisCol := 0; thisCol <= cols; thisCol++ {
			var record application.LoaderDataStoreItem
			findField := fmt.Sprintf("R%dC%d", thisRow, thisCol)
			data := r.FormValue(findField)
			//log.Println("find name", findField, data)
			record.Row = strconv.Itoa(thisRow + 1)
			record.Position = strconv.Itoa(thisCol + 1)
			record.Value = data
			record.Loader = loaderID
			record.Map = record.Position
			application.UpdateLoaderDataStore(record)
		}
	}

	//fmt.Println("table", tableRows, "x", tableColumns, "loader", loaderID)

	ListSvcDataMapHandler(w, r)
}

func SaveSvcDataMapXMLHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")

	application.ServiceMessage(inUTL)

	body := r.FormValue("pgContent")
	thisID := r.FormValue("pgid")

	status := putXMLtemplateBody(thisID, body)
	if status != 1 {
		// do nothing
	}

	ListSvcDataMapHandler(w, r)
}

func NewSvcDataMapHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderNew"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")

	application.ServiceMessage(inUTL)

	title := globals.ApplicationProperties["appname"]

	pageDM := SvcDataMapPage{
		UserMenu:  application.GetUserMenu(r),
		UserRole:  application.GetUserRole(r),
		UserNavi:  "NOT USED",
		Title:     title,
		PageTitle: "New Data Loader Template",
	}
	//fmt.Println("WCT : Page :", pageDM)

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageDM)

}

func GenSvcDataMapHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")

	application.ServiceMessage(inUTL)

	body := r.FormValue("descr")
	id := r.FormValue("name")

	var s application.LoaderStoreItem

	s.Name = id
	s.Description = body

	application.NewLoaderStore(s)

	ListSvcDataMapHandler(w, r)

}

func DeleteSvcDataMapHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessage(inUTL)
	id := application.GetURLparam(r, "loaderID")
	path := globals.ApplicationProperties["datamaptemplatepath"]
	status := application.DeleteDataFile(id+".xml", path)
	if status != 1 {
		//do nothing
	}
	application.DeleteLoaderStore(id)

	ListSvcDataMapHandler(w, r)
}
