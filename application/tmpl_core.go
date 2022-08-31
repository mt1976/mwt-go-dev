package application
// ----------------------------------------------------------------
// Automatically generated  "/application/tmpl.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Tmpl (tmpl)
// Endpoint 	        : Tmpl (TemplateID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:58
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)





//Tmpl_Publish annouces the endpoints available for this object
func Tmpl_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Tmpl_Path, Tmpl_Handler)
	mux.HandleFunc(dm.Tmpl_PathList, Tmpl_HandlerList)
	mux.HandleFunc(dm.Tmpl_PathView, Tmpl_HandlerView)
	mux.HandleFunc(dm.Tmpl_PathEdit, Tmpl_HandlerEdit)
	mux.HandleFunc(dm.Tmpl_PathNew, Tmpl_HandlerNew)
	mux.HandleFunc(dm.Tmpl_PathSave, Tmpl_HandlerSave)
	mux.HandleFunc(dm.Tmpl_PathDelete, Tmpl_HandlerDelete)
	logs.Publish("Application", dm.Tmpl_Title)
    core.Catalog_Add(dm.Tmpl_Title, dm.Tmpl_Path, "", dm.Tmpl_QueryString, "Application")
}


//Tmpl_HandlerList is the handler for the list page
func Tmpl_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Tmpl
	noItems, returnList, _ := dao.Tmpl_GetList()

	pageDetail := dm.Tmpl_PageList{
		Title:            CardTitle(dm.Tmpl_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Tmpl_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Tmpl_TemplateList, w, r, pageDetail)

}


//Tmpl_HandlerView is the handler used to View a page
func Tmpl_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Tmpl_QueryString)
	_, rD, _ := dao.Tmpl_GetByID(searchID)

	pageDetail := dm.Tmpl_Page{
		Title:       CardTitle(dm.Tmpl_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Tmpl_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = tmpl_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Tmpl_TemplateView, w, r, pageDetail)

}


//Tmpl_HandlerEdit is the handler used generate the Edit page
func Tmpl_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Tmpl_QueryString)
	_, rD, _ := dao.Tmpl_GetByID(searchID)
	
	pageDetail := dm.Tmpl_Page{
		Title:       CardTitle(dm.Tmpl_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Tmpl_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = tmpl_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Tmpl_TemplateEdit, w, r, pageDetail)
}


//Tmpl_HandlerSave is the handler used process the saving of an Tmpl
func Tmpl_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ID"))

	var item dm.Tmpl
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.Tmpl_SYSId_scrn)
		item.FIELD1 = r.FormValue(dm.Tmpl_FIELD1_scrn)
		item.FIELD2 = r.FormValue(dm.Tmpl_FIELD2_scrn)
		item.SYSCreated = r.FormValue(dm.Tmpl_SYSCreated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.Tmpl_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.Tmpl_SYSCreatedHost_scrn)
		item.SYSUpdated = r.FormValue(dm.Tmpl_SYSUpdated_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.Tmpl_SYSUpdatedHost_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.Tmpl_SYSUpdatedBy_scrn)
		item.ID = r.FormValue(dm.Tmpl_ID_scrn)
		item.ExtraField = r.FormValue(dm.Tmpl_ExtraField_scrn)
		item.ExtraField2 = r.FormValue(dm.Tmpl_ExtraField2_scrn)
		item.ExtraField3 = r.FormValue(dm.Tmpl_ExtraField3_scrn)
		item.TDate = r.FormValue(dm.Tmpl_TDate_scrn)
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Tmpl_Store(item,r)	
	http.Redirect(w, r, dm.Tmpl_Redirect, http.StatusFound)
}


//Tmpl_HandlerNew is the handler used process the creation of an Tmpl
func Tmpl_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.Tmpl_New()

	pageDetail := dm.Tmpl_Page{
		Title:       CardTitle(dm.Tmpl_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Tmpl_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = tmpl_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Tmpl_TemplateNew, w, r, pageDetail)

}	


//Tmpl_HandlerDelete is the handler used process the deletion of an Tmpl
func Tmpl_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Tmpl_QueryString)

	dao.Tmpl_Delete(searchID)	

	http.Redirect(w, r, dm.Tmpl_Redirect, http.StatusFound)
}


// Builds/Popuplates the Tmpl Page 
func tmpl_PopulatePage(rD dm.Tmpl, pageDetail dm.Tmpl_Page) dm.Tmpl_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.FIELD1 = rD.FIELD1
	pageDetail.FIELD2 = rD.FIELD2
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.ID = rD.ID
	
	pageDetail.ExtraField = rD.ExtraField
	pageDetail.ExtraField2 = rD.ExtraField2
	pageDetail.ExtraField3 = rD.ExtraField3
	pageDetail.TDate = rD.TDate
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	pageDetail.FIELD1_lookup = dao.StubLists_Get("YN")
	
	
	pageDetail.FIELD2_lookup = dao.Firm_GetLookup()
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.FIELD1_props = rD.FIELD1_props
	pageDetail.FIELD2_props = rD.FIELD2_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.ID_props = rD.ID_props
	pageDetail.ExtraField_props = rD.ExtraField_props
	pageDetail.ExtraField2_props = rD.ExtraField2_props
	pageDetail.ExtraField3_props = rD.ExtraField3_props
	pageDetail.TDate_props = rD.TDate_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	