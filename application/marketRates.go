package application
// ----------------------------------------------------------------
// Automatically generated  "/application/marketrates.go"
// ----------------------------------------------------------------
// Package            : application
// Object 			  : MarketRates
// Endpoint Root 	  : MarketRates
// Search QueryString : ID
// From   			  : 
// ----------------------------------------------------------------
// Template Generator : RussetAlbatross [r0-21.11.01]
// ----------------------------------------------------------------
// Date & Time		  : 15/11/2021 at 23:50:08
// Who & Where		  : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"html/template"
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//marketrates_PageList provides the information for the template for a list of MarketRatess
type marketrates_PageList struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.MarketRates
}

//marketrates_Page provides the information for the template for an individual MarketRates
type marketrates_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	AppInternalID  string
	// Automatically generated 15/11/2021 by matttownsend on silicon.local - START
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
	
	// Automatically generated 15/11/2021 by matttownsend on silicon.local - END
}

const (
	MarketRates_Redirect = dm.MarketRates_PathList
)

//MarketRates_Publish annouces the endpoints available for this object
func MarketRates_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.MarketRates_PathList, MarketRates_HandlerList)
	mux.HandleFunc(dm.MarketRates_PathView, MarketRates_HandlerView)
	
	
	
	
	logs.Publish("Application", dm.MarketRates_Title)
}

//MarketRates_HandlerList is the handler for the list page
func MarketRates_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.MarketRates
	noItems, returnList, _ := dao.MarketRates_GetList()


	pageDetail := marketrates_PageList{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.MarketRates_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.MarketRates_TemplateList, core.GetUserRole(r)))
	t.Execute(w, pageDetail)
}

//MarketRates_HandlerView is the handler used to View a page
func MarketRates_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.MarketRates_QueryString)
	_, rD, _ := dao.MarketRates_GetByID(searchID)

	pageDetail := marketrates_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.MarketRates_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:		     rD.AppInternalID,
		// 
		// Automatically generated 15/11/2021 by matttownsend on silicon.local - START
		SYSId: rD.SYSId,
		Id: rD.Id,
		Bid: rD.Bid,
		Mid: rD.Mid,
		Offer: rD.Offer,
		Market: rD.Market,
		Tenor: rD.Tenor,
		Series: rD.Series,
		Name: rD.Name,
		Source: rD.Source,
		Destination: rD.Destination,
		Class: rD.Class,
		SYSCreated: rD.SYSCreated,
		SYSWho: rD.SYSWho,
		SYSHost: rD.SYSHost,
		Date: rD.Date,
		SYSUpdated: rD.SYSUpdated,
		SYSCreatedBy: rD.SYSCreatedBy,
		SYSCreatedHost: rD.SYSCreatedHost,
		SYSUpdatedBy: rD.SYSUpdatedBy,
		SYSUpdatedHost: rD.SYSUpdatedHost,
		
		// Automatically generated 15/11/2021 by matttownsend on silicon.local - END
		//
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.MarketRates_TemplateView, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//MarketRates_HandlerEdit is the handler used generate the Edit page
func MarketRates_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.MarketRates_QueryString)
	_, rD, _ := dao.MarketRates_GetByID(searchID)
	
	pageDetail := marketrates_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.MarketRates_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:          rD.AppInternalID,
		// Automatically generated 15/11/2021 by matttownsend on silicon.local - START
			SYSId: rD.SYSId,
			Id: rD.Id,
			Bid: rD.Bid,
			Mid: rD.Mid,
			Offer: rD.Offer,
			Market: rD.Market,
			Tenor: rD.Tenor,
			Series: rD.Series,
			Name: rD.Name,
			Source: rD.Source,
			Destination: rD.Destination,
			Class: rD.Class,
			SYSCreated: rD.SYSCreated,
			SYSWho: rD.SYSWho,
			SYSHost: rD.SYSHost,
			Date: rD.Date,
			SYSUpdated: rD.SYSUpdated,
			SYSCreatedBy: rD.SYSCreatedBy,
			SYSCreatedHost: rD.SYSCreatedHost,
			SYSUpdatedBy: rD.SYSUpdatedBy,
			SYSUpdatedHost: rD.SYSUpdatedHost,
		
	// Automatically generated 15/11/2021 by matttownsend on silicon.local - END
		//Post Import Actions - START

		// Post Import Actions - END
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.MarketRates_TemplateEdit, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//MarketRates_HandlerSave is the handler used process the saving of an MarketRates
func MarketRates_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ID"))

	var item dm.MarketRates

	//item.AppInternalID = r.FormValue("AppInternalID")
	// Automatically generated 15/11/2021 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue("SYSId")
		item.Id = r.FormValue("Id")
		item.Bid = r.FormValue("Bid")
		item.Mid = r.FormValue("Mid")
		item.Offer = r.FormValue("Offer")
		item.Market = r.FormValue("Market")
		item.Tenor = r.FormValue("Tenor")
		item.Series = r.FormValue("Series")
		item.Name = r.FormValue("Name")
		item.Source = r.FormValue("Source")
		item.Destination = r.FormValue("Destination")
		item.Class = r.FormValue("Class")
		item.SYSCreated = r.FormValue("SYSCreated")
		item.SYSWho = r.FormValue("SYSWho")
		item.SYSHost = r.FormValue("SYSHost")
		item.Date = r.FormValue("Date")
		item.SYSUpdated = r.FormValue("SYSUpdated")
		item.SYSCreatedBy = r.FormValue("SYSCreatedBy")
		item.SYSCreatedHost = r.FormValue("SYSCreatedHost")
		item.SYSUpdatedBy = r.FormValue("SYSUpdatedBy")
		item.SYSUpdatedHost = r.FormValue("SYSUpdatedHost")
	
	// Automatically generated 15/11/2021 by matttownsend on silicon.local - END

	dao.MarketRates_Store(item)	

	http.Redirect(w, r, MarketRates_Redirect, http.StatusFound)
}

//MarketRates_HandlerNew is the handler used process the creation of an MarketRates
func MarketRates_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := marketrates_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.MarketRates_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:   "NEW",
		// Automatically generated 15/11/2021 by matttownsend on silicon.local - START
			SYSId: "0",
			Id: "",
			Bid: "",
			Mid: "",
			Offer: "",
			Market: "",
			Tenor: "",
			Series: "",
			Name: "",
			Source: "",
			Destination: "",
			Class: "",
			SYSCreated: "",
			SYSWho: "",
			SYSHost: "",
			Date: "",
			SYSUpdated: "",
			SYSCreatedBy: "",
			SYSCreatedHost: "",
			SYSUpdatedBy: "",
			SYSUpdatedHost: "",
		
		// Automatically generated 15/11/2021 by matttownsend on silicon.local - END
		//
		// Post Import Actions - START

		// Post Import Actions - END
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.MarketRates_TemplateNew, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//MarketRates_HandlerDelete is the handler used process the deletion of an MarketRates
func MarketRates_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.MarketRates_QueryString)

	dao.MarketRates_Delete(searchID)	

	http.Redirect(w, r, MarketRates_Redirect, http.StatusFound)
}