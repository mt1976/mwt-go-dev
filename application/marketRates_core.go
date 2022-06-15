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
// Date & Time		    : 14/06/2022 at 21:32:06
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//marketrates_PageList provides the information for the template for a list of MarketRatess
type MarketRates_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.MarketRates
}
//MarketRates_Redirect provides a page to return to aftern an action
const (
	MarketRates_Redirect = dm.MarketRates_PathList
)

//marketrates_Page provides the information for the template for an individual MarketRates
type MarketRates_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	Id         string
	Bid         string
	Mid         string
	Offer         string
	Market         string
	Tenor         string
	Series         string
	Name         string
	Source         string
	Destination         string
	Class         string
	SYSCreated         string
	SYSWho         string
	SYSHost         string
	Date         string
	SYSUpdated         string
	SYSCreatedBy         string
	SYSCreatedHost         string
	SYSUpdatedBy         string
	SYSUpdatedHost         string
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



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

	pageDetail := MarketRates_PageList{
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

	pageDetail := MarketRates_Page{
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
func marketrates_PopulatePage(rD dm.MarketRates, pageDetail MarketRates_Page) MarketRates_Page {
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	