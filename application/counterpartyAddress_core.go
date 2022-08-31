package application
// ----------------------------------------------------------------
// Automatically generated  "/application/counterpartyaddress.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CounterpartyAddress (counterpartyaddress)
// Endpoint 	        : CounterpartyAddress (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:47
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)





//CounterpartyAddress_Publish annouces the endpoints available for this object
func CounterpartyAddress_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.CounterpartyAddress_PathList, CounterpartyAddress_HandlerList)
	mux.HandleFunc(dm.CounterpartyAddress_PathView, CounterpartyAddress_HandlerView)
	mux.HandleFunc(dm.CounterpartyAddress_PathEdit, CounterpartyAddress_HandlerEdit)
	mux.HandleFunc(dm.CounterpartyAddress_PathNew, CounterpartyAddress_HandlerNew)
	mux.HandleFunc(dm.CounterpartyAddress_PathSave, CounterpartyAddress_HandlerSave)
	mux.HandleFunc(dm.CounterpartyAddress_PathDelete, CounterpartyAddress_HandlerDelete)
	logs.Publish("Application", dm.CounterpartyAddress_Title)
    //No API
}


//CounterpartyAddress_HandlerList is the handler for the list page
func CounterpartyAddress_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.CounterpartyAddress
	noItems, returnList, _ := dao.CounterpartyAddress_GetList()

	pageDetail := dm.CounterpartyAddress_PageList{
		Title:            CardTitle(dm.CounterpartyAddress_Title, core.Action_List),
		PageTitle:        PageTitle(dm.CounterpartyAddress_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.CounterpartyAddress_TemplateList, w, r, pageDetail)

}


//CounterpartyAddress_HandlerView is the handler used to View a page
func CounterpartyAddress_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CounterpartyAddress_QueryString)
	_, rD, _ := dao.CounterpartyAddress_GetByID(searchID)

	pageDetail := dm.CounterpartyAddress_Page{
		Title:       CardTitle(dm.CounterpartyAddress_Title, core.Action_View),
		PageTitle:   PageTitle(dm.CounterpartyAddress_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = counterpartyaddress_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.CounterpartyAddress_TemplateView, w, r, pageDetail)

}


//CounterpartyAddress_HandlerEdit is the handler used generate the Edit page
func CounterpartyAddress_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CounterpartyAddress_QueryString)
	_, rD, _ := dao.CounterpartyAddress_GetByID(searchID)
	
	pageDetail := dm.CounterpartyAddress_Page{
		Title:       CardTitle(dm.CounterpartyAddress_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.CounterpartyAddress_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = counterpartyaddress_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.CounterpartyAddress_TemplateEdit, w, r, pageDetail)
}


//CounterpartyAddress_HandlerSave is the handler used process the saving of an CounterpartyAddress
func CounterpartyAddress_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("CompID"))

	var item dm.CounterpartyAddress
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.NameFirm = r.FormValue(dm.CounterpartyAddress_NameFirm_scrn)
		item.NameCentre = r.FormValue(dm.CounterpartyAddress_NameCentre_scrn)
		item.Address1 = r.FormValue(dm.CounterpartyAddress_Address1_scrn)
		item.Address2 = r.FormValue(dm.CounterpartyAddress_Address2_scrn)
		item.CityTown = r.FormValue(dm.CounterpartyAddress_CityTown_scrn)
		item.County = r.FormValue(dm.CounterpartyAddress_County_scrn)
		item.Postcode = r.FormValue(dm.CounterpartyAddress_Postcode_scrn)
		item.CompID = r.FormValue(dm.CounterpartyAddress_CompID_scrn)
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.CounterpartyAddress_Store(item,r)	
	http.Redirect(w, r, dm.CounterpartyAddress_Redirect, http.StatusFound)
}


//CounterpartyAddress_HandlerNew is the handler used process the creation of an CounterpartyAddress
func CounterpartyAddress_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.CounterpartyAddress_New()

	pageDetail := dm.CounterpartyAddress_Page{
		Title:       CardTitle(dm.CounterpartyAddress_Title, core.Action_New),
		PageTitle:   PageTitle(dm.CounterpartyAddress_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = counterpartyaddress_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.CounterpartyAddress_TemplateNew, w, r, pageDetail)

}	


//CounterpartyAddress_HandlerDelete is the handler used process the deletion of an CounterpartyAddress
func CounterpartyAddress_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.CounterpartyAddress_QueryString)

	dao.CounterpartyAddress_Delete(searchID)	

	http.Redirect(w, r, dm.CounterpartyAddress_Redirect, http.StatusFound)
}


// Builds/Popuplates the CounterpartyAddress Page 
func counterpartyaddress_PopulatePage(rD dm.CounterpartyAddress, pageDetail dm.CounterpartyAddress_Page) dm.CounterpartyAddress_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.NameFirm = rD.NameFirm
	pageDetail.NameCentre = rD.NameCentre
	pageDetail.Address1 = rD.Address1
	pageDetail.Address2 = rD.Address2
	pageDetail.CityTown = rD.CityTown
	pageDetail.County = rD.County
	pageDetail.Postcode = rD.Postcode
	pageDetail.CompID = rD.CompID
	
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.NameFirm_props = rD.NameFirm_props
	pageDetail.NameCentre_props = rD.NameCentre_props
	pageDetail.Address1_props = rD.Address1_props
	pageDetail.Address2_props = rD.Address2_props
	pageDetail.CityTown_props = rD.CityTown_props
	pageDetail.County_props = rD.County_props
	pageDetail.Postcode_props = rD.Postcode_props
	pageDetail.CompID_props = rD.CompID_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	