package application

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/mt1976/mwt-go-dev/adaptor"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

//SvcDataMapPage is cheese
type SvcDataMapPage struct {
	UserMenu        []dm.AppMenuItem
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
	InstanceList    []SystemStoreItem
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
	DataMapID           string
	DataMapName         string
	DataMapFileID       string
	DataMapDescription  string
	DataMapXMLFile      string
	DataMapLastrun      string
	DataMapType         string
	DataMapInstance     string
	DataMapExtension    string
	DataMapInstanceList []SystemStoreItem
}

func DataLoader_Publish_Impl(mux http.ServeMux) {
	mux.HandleFunc("/listSvcDataMap/", DataLoader_HandlerList_Impl)
	mux.HandleFunc("/viewSvcDataMap/", DataLoader_HandlerView_Impl)
	mux.HandleFunc("/editSvcDataMap/", DataLoader_HandlerEdit_Impl)
	mux.HandleFunc("/saveSvcDataMap/", DataLoader_HandlerSave_Impl)
	mux.HandleFunc("/viewSvcDataMapXML/", DataLoader_HandlerViewXML_Impl)
	mux.HandleFunc("/editSvcDataMapXML/", DataLoader_HandlerEditXML_Impl)
	mux.HandleFunc("/saveSvcDataMapXML/", DataLoader_HandlerSaveXML_Impl)
	mux.HandleFunc("/newSvcDataMap/", DataLoader_HandlerNew_Impl)
	mux.HandleFunc("/genSvcDataMap/", DataLoader_HandlerGenerate_Impl)
	mux.HandleFunc("/deleteSvcDataMap/", DataLoader_HandlerDelete_Impl)
	mux.HandleFunc("/editDataMapColumns/", ListLoaderMapStoreHandler)
	mux.HandleFunc("/editLoaderMapStore/", EditLoaderMapStoreHandler)
	mux.HandleFunc("/saveLoaderMapStore/", SaveLoaderMapStoreHandler)
	mux.HandleFunc("/newLoaderMapStore/", NewLoaderMapStoreHandler)
	mux.HandleFunc("/runLoader/", DataLoader_HandlerRun_Impl)
	logs.Publish("Application", dm.DataLoader_Title+" Impl")
	core.Catalog_Add(dm.DataLoaderData_Loader, dm.DataLoaderData_PathList, "", dm.DataLoader_QueryString, core.DataSource_Siena)
}

func DataLoader_HandlerList_Impl(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderList"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	//	requestID := uuid.New()

	core.ServiceMessage(inUTL)

	noRows, loaderList, _ := GetLoaderStoreList()
	_, instanceList, _ := GetSystemStoreList()

	var dataMapItemsList []SvcDataMapItem

	for _, listItem := range loaderList {
		var newDataMapItem SvcDataMapItem
		newDataMapItem.DataMapID = listItem.Id
		newDataMapItem.DataMapName = listItem.Name
		newDataMapItem.DataMapFileID = listItem.Filename
		newDataMapItem.DataMapDescription = listItem.Description
		newDataMapItem.DataMapXMLFile = listItem.Filename
		newDataMapItem.DataMapLastrun = listItem.Lastrun
		newDataMapItem.DataMapType = listItem.Type
		newDataMapItem.DataMapInstance = listItem.Instance
		newDataMapItem.DataMapExtension = listItem.Extension
		newDataMapItem.DataMapInstanceList = instanceList
		//fmt.Println("newDataMapItem", newDataMapItem)
		dataMapItemsList = append(dataMapItemsList, newDataMapItem)
	}

	pageSrvEvironment := SvcDataMapPage{
		UserMenu:        UserMenu_Get(r),
		UserRole:        Session_GetUserRole(r),
		UserNavi:        "NOT USED",
		Title:           core.ApplicationProperties["appname"],
		PageTitle:       PageTitle("Data Loader", core.Action_List),
		NoDataMapIDs:    noRows,
		SvcDataMapItems: dataMapItemsList,
	}

	//fmt.Println("Page Data", pageSrvEvironment)

	ExecuteTemplate(tmpl, w, r, pageSrvEvironment)

}

func DataLoader_HandlerView_Impl(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	buildGridPage("LoaderView", w, r)

}

func buildGridPage(tmpl string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	inUTL := r.URL.Path
	thisID := core.GetURLparam(r, "loaderID")
	core.ServiceMessage(inUTL)
	//title := core.ApplicationProperties["appname"]
	var wrkDataMapCols []DataHdr
	noColumns, wrkLoaderHeadersList, _ := GetLoaderMapStoreListByLoader(thisID)
	for _, colData := range wrkLoaderHeadersList {
		var tmpVal DataHdr
		tmpVal.ColID, _ = strconv.Atoi(colData.Position)
		tmpVal.DataHdrItem = colData.Name
		wrkDataMapCols = append(wrkDataMapCols, tmpVal)
	}

	var wrkDataMapRows []DataRow
	var wrk DataCol
	var dr DataRow

	noRows := GetLoaderDataStoreRowsByLoader(thisID)

	for thisRow := 0; thisRow < noRows; thisRow++ {
		var wrkDataMapColItems []DataCol

		noItems, wrkLoaderRowsList, _ := GetLoaderDataStoreRowList(thisID, strconv.Itoa(thisRow+1))
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
		UserMenu:       UserMenu_Get(r),
		UserRole:       Session_GetUserRole(r),
		UserNavi:       "NOT USED",
		Title:          CardTitle(dm.DataLoader_Title, tmpl),
		PageTitle:      PageTitle(dm.DataLoader_Title, tmpl),
		NoDataMapIDs:   0,
		SvcDataMapCols: wrkDataMapCols,
		DataMapPageID:  thisID,
		DataRows:       wrkDataMapRows,
		JSCols:         noColumns,
		JSRows:         noRows,
	}

	ExecuteTemplate(tmpl, w, r, pageSrvEvironment)
}

func getDataListFile(fileID string) string {
	content, err := ioutil.ReadFile(fileID)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	return string(content)
}

func DataLoader_HandlerEdit_Impl(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	buildGridPage("LoaderEdit", w, r)

}

func DataLoader_HandlerViewXML_Impl(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderTemplateView"

	inUTL := r.URL.Path

	w.Header().Set("Content-Type", "text/html")

	thisID := core.GetURLparam(r, "loaderID")

	core.ServiceMessage(inUTL)

	text, _ := getXMLtemplateBody(thisID)

	pageSrvEvironment := SvcDataMapPage{
		UserMenu:      UserMenu_Get(r),
		UserRole:      Session_GetUserRole(r),
		UserNavi:      "NOT USED",
		Title:         CardTitle("Data Loader XML", core.Action_View),
		PageTitle:     PageTitle("Data Loader XML", core.Action_View),
		DataMapPageID: thisID,
		JSRows:        35,
		FullRecord:    html.UnescapeString(text),
	}

	ExecuteTemplate(tmpl, w, r, pageSrvEvironment)

}

func getXMLtemplateBody(thisID string) (string, error) {
	path := core.ApplicationProperties["datamaptemplatepath"]
	_, loaderItem, _ := GetLoaderStoreByID(thisID)
	fileName := loaderItem.Filename + ".template"
	content, err := core.ReadDataFile(fileName, path)
	if err != nil {
		log.Fatal(err)
	}
	return content, err
}

func putXMLtemplateBody(thisID string, content string) int {
	path := core.ApplicationProperties["datamaptemplatepath"]
	_, loaderItem, _ := GetLoaderStoreByID(thisID)
	fileName := loaderItem.Filename + ".template"
	status := core.WriteDataFile(fileName, path, content)
	return status
}

func DataLoader_HandlerEditXML_Impl(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderTemplateEdit"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")

	thisID := core.GetURLparam(r, "loaderID")

	core.ServiceMessage(inUTL)

	//title := core.ApplicationProperties["appname"]

	// Get Data Here
	fullRec, _ := getXMLtemplateBody(thisID)

	pageEditSvcDataMapXML := SvcDataMapPage{
		UserMenu:      UserMenu_Get(r),
		UserRole:      Session_GetUserRole(r),
		UserNavi:      "NOT USED",
		Title:         CardTitle("Data Loader XML", core.Action_Edit),
		PageTitle:     PageTitle("Data Loader XML", core.Action_Edit),
		DataMapPageID: thisID,
		JSRows:        35,
		FullRecord:    html.UnescapeString(fullRec),
	}

	ExecuteTemplate(tmpl, w, r, pageEditSvcDataMapXML)

}

func DataLoader_HandlerSave_Impl(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	//	tmpl := "editSrvConfiguration"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")

	core.ServiceMessage(inUTL)

	tableColumns := r.FormValue("tableColumns")
	tableRows := r.FormValue("tableRows")
	loaderID := r.FormValue("loaderID")

	cols, _ := strconv.Atoi(tableColumns)
	rows, _ := strconv.Atoi(tableRows)

	cols = cols - 1
	rows = rows - 1

	DeleteLoaderDataStoreByLoader(loaderID)

	for thisRow := 0; thisRow <= rows; thisRow++ {
		for thisCol := 0; thisCol <= cols; thisCol++ {
			var record LoaderDataStoreItem
			findField := fmt.Sprintf("R%dC%d", thisRow, thisCol)
			data := r.FormValue(findField)
			//log.Println("find name", findField, data)
			record.Row = strconv.Itoa(thisRow + 1)
			record.Position = strconv.Itoa(thisCol + 1)
			record.Value = data
			record.Loader = loaderID
			record.Map = record.Position
			UpdateLoaderDataStore(record)
		}
	}

	//fmt.Println("table", tableRows, "x", tableColumns, "loader", loaderID)

	http.Redirect(w, r, "/listSvcDataMap", http.StatusFound)
}

func DataLoader_HandlerSaveXML_Impl(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")

	core.ServiceMessage(inUTL)

	body := r.FormValue("pgContent")
	thisID := r.FormValue("pgid")

	status := putXMLtemplateBody(thisID, body)
	if status != 1 {
		// do nothing
	}

	http.Redirect(w, r, "/listSvcDataMap", http.StatusFound)

}

func DataLoader_HandlerNew_Impl(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	tmpl := "LoaderNew"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")

	core.ServiceMessage(inUTL)

	//title := core.ApplicationProperties["appname"]

	_, instanceList, _ := GetSystemStoreList()

	pageDM := SvcDataMapPage{
		UserMenu:     UserMenu_Get(r),
		UserRole:     Session_GetUserRole(r),
		UserNavi:     "NOT USED",
		Title:        CardTitle("Data Loader", core.Action_New),
		PageTitle:    PageTitle("Data Loader", core.Action_New),
		InstanceList: instanceList,
	}
	//fmt.Println("WCT : Page :", pageDM)

	ExecuteTemplate(tmpl, w, r, pageDM)

}

func DataLoader_HandlerGenerate_Impl(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")

	core.ServiceMessage(inUTL)

	body := r.FormValue("descr")
	id := r.FormValue("name")
	txntype := r.FormValue("Type")
	instanceID := r.FormValue("Instance")
	extensionID := r.FormValue("Extension")

	var s LoaderStoreItem

	s.Name = id
	s.Description = body
	s.Type = txntype
	s.Instance = instanceID
	s.Extension = extensionID

	NewLoaderStore(s, r)

	http.Redirect(w, r, "/listSvcDataMap", http.StatusFound)

}

func DataLoader_HandlerDelete_Impl(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)
	id := core.GetURLparam(r, "loaderID")
	path := core.ApplicationProperties["datamaptemplatepath"]
	status := core.DeleteDataFile(id+".template", path)
	if status != 1 {
		//do nothing
	}

	DeleteLoaderStore(id)
	DeleteLoaderDataStoreByLoader(id)
	DeleteLoaderMapStoreByLoader(id)

	http.Redirect(w, r, "/listSvcDataMap", http.StatusFound)
}

func DataLoader_HandlerRun_Impl(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	core.ServiceMessage(inUTL)
	id := core.GetURLparam(r, "loaderID")

	_, loader, _ := GetLoaderStoreByID(id)

	// Get template
	//instanceID := loader.Instance
	//extensionID := loader.Extension
	//log.Printf("instance id %s %s", instanceID, extensionID)
	importtemplate, err := core.ReadDataFile(loader.Filename+".template", core.ApplicationProperties["datamaptemplatepath"])
	if err != nil {
		log.Println(err.Error())
	}

	origTemplate := importtemplate

	//log.Println(importtemplate)

	noRows, _, _ := GetLoaderDataStoreListByLoaderCols(id, "1")
	log.Println("Processing    : " + strconv.Itoa(noRows) + " of data")
	// get columns
	noColumns, listColumns, _ := GetLoaderMapStoreListByLoader(id)
	//	log.Println("No Cols=", noColumns, "List", listColumns)

	path := core.SienaProperties["transactional_in"]
	if loader.Type == "static" {
		path = core.SienaProperties["static_in"]
	}

	log.Println("Delivery      : " + path)
	extensionID := loader.Extension
	if len(extensionID) == 0 {
		extensionID = ".xml"
	}

	for thisRow := 1; thisRow <= noRows; thisRow++ {
		// Reset the template data
		importtemplate = origTemplate
		for thisColumn := 1; thisColumn <= noColumns; thisColumn++ {
			//	log.Println(thisRow, thisColumn)
			_, dataItem, _ := GetLoaderDataStoreListByLoaderItem(id, strconv.Itoa(thisRow), strconv.Itoa(thisColumn))
			activeColumn := thisColumn - 1
			importtemplate = replaceWildcard(importtemplate, listColumns[activeColumn].Name, dataItem.Value)
		}
		// Special Replacement Section
		newID := uuid.New().String()
		importtemplate = replaceWildcard(importtemplate, "!MSG.ID", newID)
		today := time.Now()
		dayOfYear := fmt.Sprintf("%03d", today.YearDay())
		importtemplate = replaceWildcard(importtemplate, "!TODAY", today.Format(core.DATEFORMATSIENA))
		importtemplate = replaceWildcard(importtemplate, "!DDD", dayOfYear)
		importtemplate = replaceWildcard(importtemplate, "!DD", today.Format(core.DFDD))
		importtemplate = replaceWildcard(importtemplate, "!MM", today.Format(core.DFMM))
		importtemplate = replaceWildcard(importtemplate, "!YY", today.Format(core.DFYY))
		importtemplate = replaceWildcard(importtemplate, "!YYYY", today.Format(core.DFYYYY))
		importtemplate = replaceWildcard(importtemplate, "!hh", today.Format(core.DFhh))
		importtemplate = replaceWildcard(importtemplate, "!mm", today.Format(core.DFmm))
		importtemplate = replaceWildcard(importtemplate, "!ss", today.Format(core.DFss))
		spot := core.CalculateSpotDate(today)

		importtemplate = replaceWildcard(importtemplate, "!SPOT", spot.Format(core.DATEFORMATSIENA))

		importtemplate = replaceWildcard(importtemplate, "!1M", core.CalculateTenorDate(today, "1").Format(core.DATEFORMATSIENA))
		importtemplate = replaceWildcard(importtemplate, "!2M", core.CalculateTenorDate(today, "2").Format(core.DATEFORMATSIENA))
		importtemplate = replaceWildcard(importtemplate, "!3M", core.CalculateTenorDate(today, "3").Format(core.DATEFORMATSIENA))
		importtemplate = replaceWildcard(importtemplate, "!4M", core.CalculateTenorDate(today, "4").Format(core.DATEFORMATSIENA))
		importtemplate = replaceWildcard(importtemplate, "!5M", core.CalculateTenorDate(today, "5").Format(core.DATEFORMATSIENA))
		importtemplate = replaceWildcard(importtemplate, "!6M", core.CalculateTenorDate(today, "6").Format(core.DATEFORMATSIENA))
		importtemplate = replaceWildcard(importtemplate, "!7M", core.CalculateTenorDate(today, "7").Format(core.DATEFORMATSIENA))
		importtemplate = replaceWildcard(importtemplate, "!8M", core.CalculateTenorDate(today, "8").Format(core.DATEFORMATSIENA))
		importtemplate = replaceWildcard(importtemplate, "!9M", core.CalculateTenorDate(today, "9").Format(core.DATEFORMATSIENA))
		importtemplate = replaceWildcard(importtemplate, "!10M", core.CalculateTenorDate(today, "10").Format(core.DATEFORMATSIENA))
		importtemplate = replaceWildcard(importtemplate, "!11M", core.CalculateTenorDate(today, "11").Format(core.DATEFORMATSIENA))
		importtemplate = replaceWildcard(importtemplate, "!12M", core.CalculateTenorDate(today, "12").Format(core.DATEFORMATSIENA))
		importtemplate = replaceWildcard(importtemplate, "!1Y", core.CalculateTenorDate(today, "12").Format(core.DATEFORMATSIENA))
		importtemplate = replaceWildcard(importtemplate, "!2Y", core.CalculateTenorDate(today, "23").Format(core.DATEFORMATSIENA))
		importtemplate = replaceWildcard(importtemplate, "!3Y", core.CalculateTenorDate(today, "36").Format(core.DATEFORMATSIENA))
		importtemplate = replaceWildcard(importtemplate, "!5Y", core.CalculateTenorDate(today, "60").Format(core.DATEFORMATSIENA))
		importtemplate = replaceWildcard(importtemplate, "!FDY", core.CalculateFirstDateOfYear(today).Format(core.DATEFORMATSIENA))
		importtemplate = replaceWildcard(importtemplate, "!SEQ", strconv.Itoa(thisRow))
		importtemplate = replaceWildcard(importtemplate, "!LEI", "213800APCD7UDNQHOI68")

		// Deliver Data

		filename := newID + extensionID

		val := core.WriteDataFileAbsolute(filename, path, importtemplate)
		if val != 0 {
			//do nothing
		}
		//log.Println(importtemplate)

		//logs.Information("Send Message", fileName)
		_ = adaptor.ExternalMessage_Sent(newID, adaptor.Message_FormatXML, "", path, []byte(importtemplate), filename, 10, adaptor.Message_TimeoutAction_Notify)
		//logs.Information("Return from Send Message", fileName)

		//logs.Information("StaticImport_DispatchXML:", err.Error())

	}

	loader.Lastrun = time.Now().Format(core.DATETIMEFORMATUSER)
	PutLoaderStore(loader, r)

	http.Redirect(w, r, "/listSvcDataMap", http.StatusFound)

}

func replaceWildcard(orig string, replaceThis string, withThis string) string {
	return core.ReplaceWildcard(orig, replaceThis, withThis)
}
