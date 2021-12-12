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
// Date & Time		    : 12/12/2021 at 16:13:19
// Who & Where		    : matttownsend on silicon.local
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
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Systems
}

//systems_Page provides the information for the template for an individual Systems
type Systems_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		SYSId string
		Id string
		Name string
		Staticin string
		Staticout string
		Txnin string
		Txnout string
		Fundscheckin string
		Fundscheckout string
		SYSCreated string
		SYSWho string
		SYSHost string
		SYSUpdated string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdatedBy string
		SYSUpdatedHost string
		SWIFTin string
		SWIFTout string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	Systems_Redirect = dm.Systems_PathList
)

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

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

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

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

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
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.Systems_SYSId)
		item.Id = r.FormValue(dm.Systems_Id)
		item.Name = r.FormValue(dm.Systems_Name)
		item.Staticin = r.FormValue(dm.Systems_Staticin)
		item.Staticout = r.FormValue(dm.Systems_Staticout)
		item.Txnin = r.FormValue(dm.Systems_Txnin)
		item.Txnout = r.FormValue(dm.Systems_Txnout)
		item.Fundscheckin = r.FormValue(dm.Systems_Fundscheckin)
		item.Fundscheckout = r.FormValue(dm.Systems_Fundscheckout)
		item.SYSCreated = r.FormValue(dm.Systems_SYSCreated)
		item.SYSWho = r.FormValue(dm.Systems_SYSWho)
		item.SYSHost = r.FormValue(dm.Systems_SYSHost)
		item.SYSUpdated = r.FormValue(dm.Systems_SYSUpdated)
		item.SYSCreatedBy = r.FormValue(dm.Systems_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.Systems_SYSCreatedHost)
		item.SYSUpdatedBy = r.FormValue(dm.Systems_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.Systems_SYSUpdatedHost)
		item.SWIFTin = r.FormValue(dm.Systems_SWIFTin)
		item.SWIFTout = r.FormValue(dm.Systems_SWIFTout)
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

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

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Systems_Page{
		Title:       CardTitle(dm.Systems_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Systems_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.Id = ""
pageDetail.Name = ""
pageDetail.Staticin = ""
pageDetail.Staticout = ""
pageDetail.Txnin = ""
pageDetail.Txnout = ""
pageDetail.Fundscheckin = ""
pageDetail.Fundscheckout = ""
pageDetail.SYSCreated = ""
pageDetail.SYSWho = ""
pageDetail.SYSHost = ""
pageDetail.SYSUpdated = ""
pageDetail.SYSCreatedBy = ""
pageDetail.SYSCreatedHost = ""
pageDetail.SYSUpdatedBy = ""
pageDetail.SYSUpdatedHost = ""
pageDetail.SWIFTin = ""
pageDetail.SWIFTout = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

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