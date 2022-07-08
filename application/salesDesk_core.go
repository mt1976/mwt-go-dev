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
// Date & Time		    : 28/06/2022 at 16:10:56
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)





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

	pageDetail := dm.SalesDesk_PageList{
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

	pageDetail := dm.SalesDesk_Page{
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
func salesdesk_PopulatePage(rD dm.SalesDesk, pageDetail dm.SalesDesk_Page) dm.SalesDesk_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.Name = rD.Name
	pageDetail.ReportDealsOver = rD.ReportDealsOver
	pageDetail.ReportDealsOverCCY = rD.ReportDealsOverCCY
	pageDetail.AccountTransferCutOffTime = rD.AccountTransferCutOffTime
	pageDetail.AccountTransferCutOffTimeTimeZone = rD.AccountTransferCutOffTimeTimeZone
	pageDetail.AccountTransferCutOffTimeCutOffPeriod = rD.AccountTransferCutOffTimeCutOffPeriod
	
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.Name_props = rD.Name_props
	pageDetail.ReportDealsOver_props = rD.ReportDealsOver_props
	pageDetail.ReportDealsOverCCY_props = rD.ReportDealsOverCCY_props
	pageDetail.AccountTransferCutOffTime_props = rD.AccountTransferCutOffTime_props
	pageDetail.AccountTransferCutOffTimeTimeZone_props = rD.AccountTransferCutOffTimeTimeZone_props
	pageDetail.AccountTransferCutOffTimeCutOffPeriod_props = rD.AccountTransferCutOffTimeCutOffPeriod_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	