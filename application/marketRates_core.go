package application
// ----------------------------------------------------------------
// Automatically generated  "/application/marketrates.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : MarketRates (marketrates)
// Endpoint 	        : MarketRates (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:54
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)





//MarketRates_Publish annouces the endpoints available for this object
func MarketRates_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.MarketRates_Path, MarketRates_Handler)
	mux.HandleFunc(dm.MarketRates_PathList, MarketRates_HandlerList)
	mux.HandleFunc(dm.MarketRates_PathView, MarketRates_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.MarketRates_Title)
    core.Catalog_Add(dm.MarketRates_Title, dm.MarketRates_Path, "", dm.MarketRates_QueryString, "Application")
}


//MarketRates_HandlerList is the handler for the list page
func MarketRates_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.MarketRates
	noItems, returnList, _ := dao.MarketRates_GetList()

	pageDetail := dm.MarketRates_PageList{
		Title:            CardTitle(dm.MarketRates_Title, core.Action_List),
		PageTitle:        PageTitle(dm.MarketRates_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.MarketRates_TemplateList, w, r, pageDetail)

}


//MarketRates_HandlerView is the handler used to View a page
func MarketRates_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.MarketRates_QueryString)
	_, rD, _ := dao.MarketRates_GetByID(searchID)

	pageDetail := dm.MarketRates_Page{
		Title:       CardTitle(dm.MarketRates_Title, core.Action_View),
		PageTitle:   PageTitle(dm.MarketRates_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = marketrates_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.MarketRates_TemplateView, w, r, pageDetail)

}






// Builds/Popuplates the MarketRates Page 
func marketrates_PopulatePage(rD dm.MarketRates, pageDetail dm.MarketRates_Page) dm.MarketRates_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.Id = rD.Id
	pageDetail.Bid = rD.Bid
	pageDetail.Mid = rD.Mid
	pageDetail.Offer = rD.Offer
	pageDetail.Market = rD.Market
	pageDetail.Tenor = rD.Tenor
	pageDetail.Series = rD.Series
	pageDetail.Name = rD.Name
	pageDetail.Source = rD.Source
	pageDetail.Destination = rD.Destination
	pageDetail.Class = rD.Class
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSWho = rD.SYSWho
	pageDetail.SYSHost = rD.SYSHost
	pageDetail.Date = rD.Date
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.Id_props = rD.Id_props
	pageDetail.Bid_props = rD.Bid_props
	pageDetail.Mid_props = rD.Mid_props
	pageDetail.Offer_props = rD.Offer_props
	pageDetail.Market_props = rD.Market_props
	pageDetail.Tenor_props = rD.Tenor_props
	pageDetail.Series_props = rD.Series_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.Source_props = rD.Source_props
	pageDetail.Destination_props = rD.Destination_props
	pageDetail.Class_props = rD.Class_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSWho_props = rD.SYSWho_props
	pageDetail.SYSHost_props = rD.SYSHost_props
	pageDetail.Date_props = rD.Date_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	