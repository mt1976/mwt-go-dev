package application
// ----------------------------------------------------------------
// Automatically generated  "/application/salesdesk.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : SalesDesk (salesdesk)
// Endpoint 	        : SalesDesk (Desk)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:32
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//salesdesk_PageList provides the information for the template for a list of SalesDesks
type SalesDesk_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.SalesDesk
}
//SalesDesk_Redirect provides a page to return to aftern an action
const (
	
	SalesDesk_Redirect = dm.SalesDesk_PathList
	
)

//salesdesk_Page provides the information for the template for an individual SalesDesk
type SalesDesk_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	Name         string
	Name_props     dm.FieldProperties
	ReportDealsOver         string
	ReportDealsOver_props     dm.FieldProperties
	ReportDealsOverCCY         string
	ReportDealsOverCCY_props     dm.FieldProperties
	AccountTransferCutOffTime         string
	AccountTransferCutOffTime_props     dm.FieldProperties
	AccountTransferCutOffTimeTimeZone         string
	AccountTransferCutOffTimeTimeZone_props     dm.FieldProperties
	AccountTransferCutOffTimeCutOffPeriod         string
	AccountTransferCutOffTimeCutOffPeriod_props     dm.FieldProperties
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//SalesDesk_Publish annouces the endpoints available for this object
func SalesDesk_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.SalesDesk_Path, SalesDesk_Handler)
	mux.HandleFunc(dm.SalesDesk_PathList, SalesDesk_HandlerList)
	mux.HandleFunc(dm.SalesDesk_PathView, SalesDesk_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.SalesDesk_Title)
    core.Catalog_Add(dm.SalesDesk_Title, dm.SalesDesk_Path, "", dm.SalesDesk_QueryString, "Application")
}


//SalesDesk_HandlerList is the handler for the list page
func SalesDesk_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.SalesDesk
	noItems, returnList, _ := dao.SalesDesk_GetList()

	pageDetail := SalesDesk_PageList{
		Title:            CardTitle(dm.SalesDesk_Title, core.Action_List),
		PageTitle:        PageTitle(dm.SalesDesk_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.SalesDesk_TemplateList, w, r, pageDetail)

}


//SalesDesk_HandlerView is the handler used to View a page
func SalesDesk_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.SalesDesk_QueryString)
	_, rD, _ := dao.SalesDesk_GetByID(searchID)

	pageDetail := SalesDesk_Page{
		Title:       CardTitle(dm.SalesDesk_Title, core.Action_View),
		PageTitle:   PageTitle(dm.SalesDesk_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = salesdesk_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.SalesDesk_TemplateView, w, r, pageDetail)

}






// Builds/Popuplates the SalesDesk Page 
func salesdesk_PopulatePage(rD dm.SalesDesk, pageDetail SalesDesk_Page) SalesDesk_Page {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.Name = rD.Name
	pageDetail.ReportDealsOver = rD.ReportDealsOver
	pageDetail.ReportDealsOverCCY = rD.ReportDealsOverCCY
	pageDetail.AccountTransferCutOffTime = rD.AccountTransferCutOffTime
	pageDetail.AccountTransferCutOffTimeTimeZone = rD.AccountTransferCutOffTimeTimeZone
	pageDetail.AccountTransferCutOffTimeCutOffPeriod = rD.AccountTransferCutOffTimeCutOffPeriod
	
	
	//
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.Name_props = rD.Name_props
	pageDetail.ReportDealsOver_props = rD.ReportDealsOver_props
	pageDetail.ReportDealsOverCCY_props = rD.ReportDealsOverCCY_props
	pageDetail.AccountTransferCutOffTime_props = rD.AccountTransferCutOffTime_props
	pageDetail.AccountTransferCutOffTimeTimeZone_props = rD.AccountTransferCutOffTimeTimeZone_props
	pageDetail.AccountTransferCutOffTimeCutOffPeriod_props = rD.AccountTransferCutOffTimeCutOffPeriod_props
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	