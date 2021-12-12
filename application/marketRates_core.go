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
// Date & Time		    : 12/12/2021 at 16:13:18
// Who & Where		    : matttownsend on silicon.local
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
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.MarketRates
}

//marketrates_Page provides the information for the template for an individual MarketRates
type MarketRates_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		SYSId string
		Id string
		Bid string
		Mid string
		Offer string
		Market string
		Tenor string
		Series string
		Name string
		Source string
		Destination string
		Class string
		SYSCreated string
		SYSWho string
		SYSHost string
		Date string
		SYSUpdated string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdatedBy string
		SYSUpdatedHost string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	MarketRates_Redirect = dm.MarketRates_PathList
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
    core.Catalog_Add(dm.MarketRates_Title, dm.MarketRates_Path, "", dm.MarketRates_QueryString, "APP")
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

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.MarketRates_TemplateView, w, r, pageDetail)

}

//MarketRates_HandlerEdit is the handler used generate the Edit page
func MarketRates_HandlerEdit(w http.ResponseWriter, r *http.Request) {
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
		Title:       CardTitle(dm.MarketRates_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.MarketRates_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.MarketRates_TemplateEdit, w, r, pageDetail)


}

//MarketRates_HandlerSave is the handler used process the saving of an MarketRates
func MarketRates_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	var item dm.MarketRates
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.MarketRates_SYSId)
		item.Id = r.FormValue(dm.MarketRates_Id)
		item.Bid = r.FormValue(dm.MarketRates_Bid)
		item.Mid = r.FormValue(dm.MarketRates_Mid)
		item.Offer = r.FormValue(dm.MarketRates_Offer)
		item.Market = r.FormValue(dm.MarketRates_Market)
		item.Tenor = r.FormValue(dm.MarketRates_Tenor)
		item.Series = r.FormValue(dm.MarketRates_Series)
		item.Name = r.FormValue(dm.MarketRates_Name)
		item.Source = r.FormValue(dm.MarketRates_Source)
		item.Destination = r.FormValue(dm.MarketRates_Destination)
		item.Class = r.FormValue(dm.MarketRates_Class)
		item.SYSCreated = r.FormValue(dm.MarketRates_SYSCreated)
		item.SYSWho = r.FormValue(dm.MarketRates_SYSWho)
		item.SYSHost = r.FormValue(dm.MarketRates_SYSHost)
		item.Date = r.FormValue(dm.MarketRates_Date)
		item.SYSUpdated = r.FormValue(dm.MarketRates_SYSUpdated)
		item.SYSCreatedBy = r.FormValue(dm.MarketRates_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.MarketRates_SYSCreatedHost)
		item.SYSUpdatedBy = r.FormValue(dm.MarketRates_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.MarketRates_SYSUpdatedHost)
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.MarketRates_Store(item,r)	

	http.Redirect(w, r, MarketRates_Redirect, http.StatusFound)
}

//MarketRates_HandlerNew is the handler used process the creation of an MarketRates
func MarketRates_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := MarketRates_Page{
		Title:       CardTitle(dm.MarketRates_Title, core.Action_New),
		PageTitle:   PageTitle(dm.MarketRates_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.Id = ""
pageDetail.Bid = ""
pageDetail.Mid = ""
pageDetail.Offer = ""
pageDetail.Market = ""
pageDetail.Tenor = ""
pageDetail.Series = ""
pageDetail.Name = ""
pageDetail.Source = ""
pageDetail.Destination = ""
pageDetail.Class = ""
pageDetail.SYSCreated = ""
pageDetail.SYSWho = ""
pageDetail.SYSHost = ""
pageDetail.Date = ""
pageDetail.SYSUpdated = ""
pageDetail.SYSCreatedBy = ""
pageDetail.SYSCreatedHost = ""
pageDetail.SYSUpdatedBy = ""
pageDetail.SYSUpdatedHost = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.MarketRates_TemplateNew, w, r, pageDetail)

}

//MarketRates_HandlerDelete is the handler used process the deletion of an MarketRates
func MarketRates_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.MarketRates_QueryString)

	dao.MarketRates_Delete(searchID)	

	http.Redirect(w, r, MarketRates_Redirect, http.StatusFound)
}