package application
// ----------------------------------------------------------------
// Automatically generated  "/application/counterpartycreditrating.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CounterpartyCreditRating (counterpartycreditrating)
// Endpoint 	        : CounterpartyCreditRating (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:07
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//counterpartycreditrating_PageList provides the information for the template for a list of CounterpartyCreditRatings
type CounterpartyCreditRating_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.CounterpartyCreditRating
}
//CounterpartyCreditRating_Redirect provides a page to return to aftern an action
const (
	CounterpartyCreditRating_Redirect = dm.CounterpartyCreditRating_PathList
)

//counterpartycreditrating_Page provides the information for the template for an individual CounterpartyCreditRating
type CounterpartyCreditRating_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	NameFirm         string
	NameCentre         string
	CreditRatingUsage         string
	CreditRatingAgency         string
	CreditRatingName         string
	CompID         string
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//CounterpartyCreditRating_Publish annouces the endpoints available for this object
func CounterpartyCreditRating_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.CounterpartyCreditRating_PathList, CounterpartyCreditRating_HandlerList)
	mux.HandleFunc(dm.CounterpartyCreditRating_PathView, CounterpartyCreditRating_HandlerView)
	mux.HandleFunc(dm.CounterpartyCreditRating_PathEdit, CounterpartyCreditRating_HandlerEdit)
	mux.HandleFunc(dm.CounterpartyCreditRating_PathNew, CounterpartyCreditRating_HandlerNew)
	mux.HandleFunc(dm.CounterpartyCreditRating_PathSave, CounterpartyCreditRating_HandlerSave)
	mux.HandleFunc(dm.CounterpartyCreditRating_PathDelete, CounterpartyCreditRating_HandlerDelete)
	logs.Publish("Application", dm.CounterpartyCreditRating_Title)
    //No API
}


//CounterpartyCreditRating_HandlerList is the handler for the list page
func CounterpartyCreditRating_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.CounterpartyCreditRating
	noItems, returnList, _ := dao.CounterpartyCreditRating_GetList()

	pageDetail := CounterpartyCreditRating_PageList{
		Title:            CardTitle(dm.CounterpartyCreditRating_Title, core.Action_List),
		PageTitle:        PageTitle(dm.CounterpartyCreditRating_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.CounterpartyCreditRating_TemplateList, w, r, pageDetail)

}


//CounterpartyCreditRating_HandlerView is the handler used to View a page
func CounterpartyCreditRating_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CounterpartyCreditRating_QueryString)
	_, rD, _ := dao.CounterpartyCreditRating_GetByID(searchID)

	pageDetail := CounterpartyCreditRating_Page{
		Title:       CardTitle(dm.CounterpartyCreditRating_Title, core.Action_View),
		PageTitle:   PageTitle(dm.CounterpartyCreditRating_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = counterpartycreditrating_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.CounterpartyCreditRating_TemplateView, w, r, pageDetail)

}


//CounterpartyCreditRating_HandlerEdit is the handler used generate the Edit page
func CounterpartyCreditRating_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CounterpartyCreditRating_QueryString)
	_, rD, _ := dao.CounterpartyCreditRating_GetByID(searchID)
	
	pageDetail := CounterpartyCreditRating_Page{
		Title:       CardTitle(dm.CounterpartyCreditRating_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.CounterpartyCreditRating_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = counterpartycreditrating_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.CounterpartyCreditRating_TemplateEdit, w, r, pageDetail)
}


//CounterpartyCreditRating_HandlerSave is the handler used process the saving of an CounterpartyCreditRating
func CounterpartyCreditRating_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("CompID"))

	var item dm.CounterpartyCreditRating
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.NameFirm = r.FormValue(dm.CounterpartyCreditRating_NameFirm_scrn)
		item.NameCentre = r.FormValue(dm.CounterpartyCreditRating_NameCentre_scrn)
		item.CreditRatingUsage = r.FormValue(dm.CounterpartyCreditRating_CreditRatingUsage_scrn)
		item.CreditRatingAgency = r.FormValue(dm.CounterpartyCreditRating_CreditRatingAgency_scrn)
		item.CreditRatingName = r.FormValue(dm.CounterpartyCreditRating_CreditRatingName_scrn)
		item.CompID = r.FormValue(dm.CounterpartyCreditRating_CompID_scrn)
	
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.CounterpartyCreditRating_Store(item,r)	
	http.Redirect(w, r, CounterpartyCreditRating_Redirect, http.StatusFound)
}


//CounterpartyCreditRating_HandlerNew is the handler used process the creation of an CounterpartyCreditRating
func CounterpartyCreditRating_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := CounterpartyCreditRating_Page{
		Title:       CardTitle(dm.CounterpartyCreditRating_Title, core.Action_New),
		PageTitle:   PageTitle(dm.CounterpartyCreditRating_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = counterpartycreditrating_PopulatePage(dm.CounterpartyCreditRating{} , pageDetail) 

	ExecuteTemplate(dm.CounterpartyCreditRating_TemplateNew, w, r, pageDetail)

}	


//CounterpartyCreditRating_HandlerDelete is the handler used process the deletion of an CounterpartyCreditRating
func CounterpartyCreditRating_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.CounterpartyCreditRating_QueryString)

	dao.CounterpartyCreditRating_Delete(searchID)	

	http.Redirect(w, r, CounterpartyCreditRating_Redirect, http.StatusFound)
}


// Builds/Popuplates the CounterpartyCreditRating Page 
func counterpartycreditrating_PopulatePage(rD dm.CounterpartyCreditRating, pageDetail CounterpartyCreditRating_Page) CounterpartyCreditRating_Page {
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.NameFirm = rD.NameFirm
	pageDetail.NameCentre = rD.NameCentre
	pageDetail.CreditRatingUsage = rD.CreditRatingUsage
	pageDetail.CreditRatingAgency = rD.CreditRatingAgency
	pageDetail.CreditRatingName = rD.CreditRatingName
	pageDetail.CompID = rD.CompID
	
	
	//
	// Automatically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	