package application
// ----------------------------------------------------------------
// Automatically generated  "/application/systems.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Systems (systems)
// Endpoint 	        : Systems (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:33
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//systems_PageList provides the information for the template for a list of Systemss
type Systems_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Systems
}
//Systems_Redirect provides a page to return to aftern an action
const (
	
	Systems_Redirect = dm.Systems_PathList
	
)

//systems_Page provides the information for the template for an individual Systems
type Systems_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	SYSId_props     dm.FieldProperties
	Id         string
	Id_props     dm.FieldProperties
	Name         string
	Name_props     dm.FieldProperties
	Staticin         string
	Staticin_props     dm.FieldProperties
	Staticout         string
	Staticout_props     dm.FieldProperties
	Txnin         string
	Txnin_props     dm.FieldProperties
	Txnout         string
	Txnout_props     dm.FieldProperties
	Fundscheckin         string
	Fundscheckin_props     dm.FieldProperties
	Fundscheckout         string
	Fundscheckout_props     dm.FieldProperties
	SYSCreated         string
	SYSCreated_props     dm.FieldProperties
	SYSWho         string
	SYSWho_props     dm.FieldProperties
	SYSHost         string
	SYSHost_props     dm.FieldProperties
	SYSUpdated         string
	SYSUpdated_props     dm.FieldProperties
	SYSCreatedBy         string
	SYSCreatedBy_props     dm.FieldProperties
	SYSCreatedHost         string
	SYSCreatedHost_props     dm.FieldProperties
	SYSUpdatedBy         string
	SYSUpdatedBy_props     dm.FieldProperties
	SYSUpdatedHost         string
	SYSUpdatedHost_props     dm.FieldProperties
	SWIFTin         string
	SWIFTin_props     dm.FieldProperties
	SWIFTout         string
	SWIFTout_props     dm.FieldProperties
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//Systems_Publish annouces the endpoints available for this object
func Systems_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.Systems_PathList, Systems_HandlerList)
	mux.HandleFunc(dm.Systems_PathView, Systems_HandlerView)
	mux.HandleFunc(dm.Systems_PathEdit, Systems_HandlerEdit)
	mux.HandleFunc(dm.Systems_PathNew, Systems_HandlerNew)
	mux.HandleFunc(dm.Systems_PathSave, Systems_HandlerSave)
	mux.HandleFunc(dm.Systems_PathDelete, Systems_HandlerDelete)
	logs.Publish("Application", dm.Systems_Title)
    //No API
}


//Systems_HandlerList is the handler for the list page
func Systems_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Systems
	noItems, returnList, _ := dao.Systems_GetList()

	pageDetail := Systems_PageList{
		Title:            CardTitle(dm.Systems_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Systems_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Systems_TemplateList, w, r, pageDetail)

}


//Systems_HandlerView is the handler used to View a page
func Systems_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Systems_QueryString)
	_, rD, _ := dao.Systems_GetByID(searchID)

	pageDetail := Systems_Page{
		Title:       CardTitle(dm.Systems_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Systems_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = systems_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Systems_TemplateView, w, r, pageDetail)

}


//Systems_HandlerEdit is the handler used generate the Edit page
func Systems_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Systems_QueryString)
	_, rD, _ := dao.Systems_GetByID(searchID)
	
	pageDetail := Systems_Page{
		Title:       CardTitle(dm.Systems_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Systems_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = systems_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Systems_TemplateEdit, w, r, pageDetail)
}


//Systems_HandlerSave is the handler used process the saving of an Systems
func Systems_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	var item dm.Systems
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.Systems_SYSId_scrn)
		item.Id = r.FormValue(dm.Systems_Id_scrn)
		item.Name = r.FormValue(dm.Systems_Name_scrn)
		item.Staticin = r.FormValue(dm.Systems_Staticin_scrn)
		item.Staticout = r.FormValue(dm.Systems_Staticout_scrn)
		item.Txnin = r.FormValue(dm.Systems_Txnin_scrn)
		item.Txnout = r.FormValue(dm.Systems_Txnout_scrn)
		item.Fundscheckin = r.FormValue(dm.Systems_Fundscheckin_scrn)
		item.Fundscheckout = r.FormValue(dm.Systems_Fundscheckout_scrn)
		item.SYSCreated = r.FormValue(dm.Systems_SYSCreated_scrn)
		item.SYSWho = r.FormValue(dm.Systems_SYSWho_scrn)
		item.SYSHost = r.FormValue(dm.Systems_SYSHost_scrn)
		item.SYSUpdated = r.FormValue(dm.Systems_SYSUpdated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.Systems_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.Systems_SYSCreatedHost_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.Systems_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.Systems_SYSUpdatedHost_scrn)
		item.SWIFTin = r.FormValue(dm.Systems_SWIFTin_scrn)
		item.SWIFTout = r.FormValue(dm.Systems_SWIFTout_scrn)
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Systems_Store(item,r)	
	http.Redirect(w, r, Systems_Redirect, http.StatusFound)
}


//Systems_HandlerNew is the handler used process the creation of an Systems
func Systems_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.Systems_New()

	pageDetail := Systems_Page{
		Title:       CardTitle(dm.Systems_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Systems_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = systems_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Systems_TemplateNew, w, r, pageDetail)

}	


//Systems_HandlerDelete is the handler used process the deletion of an Systems
func Systems_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Systems_QueryString)

	dao.Systems_Delete(searchID)	

	http.Redirect(w, r, Systems_Redirect, http.StatusFound)
}


// Builds/Popuplates the Systems Page 
func systems_PopulatePage(rD dm.Systems, pageDetail Systems_Page) Systems_Page {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.Id = rD.Id
	pageDetail.Name = rD.Name
	pageDetail.Staticin = rD.Staticin
	pageDetail.Staticout = rD.Staticout
	pageDetail.Txnin = rD.Txnin
	pageDetail.Txnout = rD.Txnout
	pageDetail.Fundscheckin = rD.Fundscheckin
	pageDetail.Fundscheckout = rD.Fundscheckout
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSWho = rD.SYSWho
	pageDetail.SYSHost = rD.SYSHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.SWIFTin = rD.SWIFTin
	pageDetail.SWIFTout = rD.SWIFTout
	
	
	//
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.Id_props = rD.Id_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.Staticin_props = rD.Staticin_props
	pageDetail.Staticout_props = rD.Staticout_props
	pageDetail.Txnin_props = rD.Txnin_props
	pageDetail.Txnout_props = rD.Txnout_props
	pageDetail.Fundscheckin_props = rD.Fundscheckin_props
	pageDetail.Fundscheckout_props = rD.Fundscheckout_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSWho_props = rD.SYSWho_props
	pageDetail.SYSHost_props = rD.SYSHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.SWIFTin_props = rD.SWIFTin_props
	pageDetail.SWIFTout_props = rD.SWIFTout_props
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	