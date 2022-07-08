package application
// ----------------------------------------------------------------
// Automatically generated  "/application/mandate.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Mandate (mandate)
// Endpoint 	        : Mandate (Mandate)
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





//Mandate_Publish annouces the endpoints available for this object
func Mandate_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Mandate_Path, Mandate_Handler)
	mux.HandleFunc(dm.Mandate_PathList, Mandate_HandlerList)
	mux.HandleFunc(dm.Mandate_PathView, Mandate_HandlerView)
	mux.HandleFunc(dm.Mandate_PathEdit, Mandate_HandlerEdit)
	mux.HandleFunc(dm.Mandate_PathNew, Mandate_HandlerNew)
	mux.HandleFunc(dm.Mandate_PathSave, Mandate_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Mandate_Title)
    core.Catalog_Add(dm.Mandate_Title, dm.Mandate_Path, "", dm.Mandate_QueryString, "Application")
}


//Mandate_HandlerList is the handler for the list page
func Mandate_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Mandate
	noItems, returnList, _ := dao.Mandate_GetList()

	pageDetail := dm.Mandate_PageList{
		Title:            CardTitle(dm.Mandate_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Mandate_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Mandate_TemplateList, w, r, pageDetail)

}


//Mandate_HandlerView is the handler used to View a page
func Mandate_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Mandate_QueryString)
	_, rD, _ := dao.Mandate_GetByID(searchID)

	pageDetail := dm.Mandate_Page{
		Title:       CardTitle(dm.Mandate_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Mandate_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = mandate_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Mandate_TemplateView, w, r, pageDetail)

}


//Mandate_HandlerEdit is the handler used generate the Edit page
func Mandate_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Mandate_QueryString)
	_, rD, _ := dao.Mandate_GetByID(searchID)
	
	pageDetail := dm.Mandate_Page{
		Title:       CardTitle(dm.Mandate_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Mandate_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = mandate_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Mandate_TemplateEdit, w, r, pageDetail)
}


//Mandate_HandlerSave is the handler used process the saving of an Mandate
func Mandate_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("CompID"))

	var item dm.Mandate
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.MandatedUserKeyCounterpartyFirm = r.FormValue(dm.Mandate_MandatedUserKeyCounterpartyFirm_scrn)
		item.MandatedUserKeyCounterpartyCentre = r.FormValue(dm.Mandate_MandatedUserKeyCounterpartyCentre_scrn)
		item.MandatedUserKeyUserName = r.FormValue(dm.Mandate_MandatedUserKeyUserName_scrn)
		item.TelephoneNumber = r.FormValue(dm.Mandate_TelephoneNumber_scrn)
		item.EmailAddress = r.FormValue(dm.Mandate_EmailAddress_scrn)
		item.Active = r.FormValue(dm.Mandate_Active_scrn)
		item.FirstName = r.FormValue(dm.Mandate_FirstName_scrn)
		item.Surname = r.FormValue(dm.Mandate_Surname_scrn)
		item.DateOfBirth = r.FormValue(dm.Mandate_DateOfBirth_scrn)
		item.Postcode = r.FormValue(dm.Mandate_Postcode_scrn)
		item.NationalIDNo = r.FormValue(dm.Mandate_NationalIDNo_scrn)
		item.PassportNo = r.FormValue(dm.Mandate_PassportNo_scrn)
		item.Country = r.FormValue(dm.Mandate_Country_scrn)
		item.CountryName = r.FormValue(dm.Mandate_CountryName_scrn)
		item.FirmName = r.FormValue(dm.Mandate_FirmName_scrn)
		item.CentreName = r.FormValue(dm.Mandate_CentreName_scrn)
		item.Notify = r.FormValue(dm.Mandate_Notify_scrn)
		item.SystemUser = r.FormValue(dm.Mandate_SystemUser_scrn)
		item.CompID = r.FormValue(dm.Mandate_CompID_scrn)
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Mandate_Store(item,r)	
	http.Redirect(w, r, dm.Mandate_Redirect, http.StatusFound)
}


//Mandate_HandlerNew is the handler used process the creation of an Mandate
func Mandate_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.Mandate_New()

	pageDetail := dm.Mandate_Page{
		Title:       CardTitle(dm.Mandate_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Mandate_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = mandate_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Mandate_TemplateNew, w, r, pageDetail)

}	



// Builds/Popuplates the Mandate Page 
func mandate_PopulatePage(rD dm.Mandate, pageDetail dm.Mandate_Page) dm.Mandate_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.MandatedUserKeyCounterpartyFirm = rD.MandatedUserKeyCounterpartyFirm
	pageDetail.MandatedUserKeyCounterpartyCentre = rD.MandatedUserKeyCounterpartyCentre
	pageDetail.MandatedUserKeyUserName = rD.MandatedUserKeyUserName
	pageDetail.TelephoneNumber = rD.TelephoneNumber
	pageDetail.EmailAddress = rD.EmailAddress
	pageDetail.Active = rD.Active
	pageDetail.FirstName = rD.FirstName
	pageDetail.Surname = rD.Surname
	pageDetail.DateOfBirth = rD.DateOfBirth
	pageDetail.Postcode = rD.Postcode
	pageDetail.NationalIDNo = rD.NationalIDNo
	pageDetail.PassportNo = rD.PassportNo
	pageDetail.Country = rD.Country
	pageDetail.CountryName = rD.CountryName
	pageDetail.FirmName = rD.FirmName
	pageDetail.CentreName = rD.CentreName
	pageDetail.Notify = rD.Notify
	pageDetail.SystemUser = rD.SystemUser
	pageDetail.CompID = rD.CompID
	
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	pageDetail.MandatedUserKeyCounterpartyFirm_lookup = dao.Firm_GetLookup()
	
	
	
	pageDetail.MandatedUserKeyCounterpartyCentre_lookup = dao.Centre_GetLookup()
	
	
	
	
	
	
	
	
	
	
	pageDetail.Active_lookup = dao.StubLists_Get("tf")
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.Country_lookup = dao.Country_GetLookup()
	
	
	
	
	
	
	
	
	
	
	pageDetail.Notify_lookup = dao.StubLists_Get("tf")
	
	
	
	
	
	
	pageDetail.MandatedUserKeyCounterpartyFirm_props = rD.MandatedUserKeyCounterpartyFirm_props
	pageDetail.MandatedUserKeyCounterpartyCentre_props = rD.MandatedUserKeyCounterpartyCentre_props
	pageDetail.MandatedUserKeyUserName_props = rD.MandatedUserKeyUserName_props
	pageDetail.TelephoneNumber_props = rD.TelephoneNumber_props
	pageDetail.EmailAddress_props = rD.EmailAddress_props
	pageDetail.Active_props = rD.Active_props
	pageDetail.FirstName_props = rD.FirstName_props
	pageDetail.Surname_props = rD.Surname_props
	pageDetail.DateOfBirth_props = rD.DateOfBirth_props
	pageDetail.Postcode_props = rD.Postcode_props
	pageDetail.NationalIDNo_props = rD.NationalIDNo_props
	pageDetail.PassportNo_props = rD.PassportNo_props
	pageDetail.Country_props = rD.Country_props
	pageDetail.CountryName_props = rD.CountryName_props
	pageDetail.FirmName_props = rD.FirmName_props
	pageDetail.CentreName_props = rD.CentreName_props
	pageDetail.Notify_props = rD.Notify_props
	pageDetail.SystemUser_props = rD.SystemUser_props
	pageDetail.CompID_props = rD.CompID_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	