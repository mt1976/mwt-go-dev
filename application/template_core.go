package application
// ----------------------------------------------------------------
// Automatically generated  "/application/template.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Template (template)
// Endpoint 	        : Template (TemplateID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:20
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//template_PageList provides the information for the template for a list of Templates
type Template_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Template
}

//template_Page provides the information for the template for an individual Template
type Template_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		SYSId string
		FIELD1 string
		FIELD2 string
		SYSCreated string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdated string
		SYSUpdatedHost string
		SYSUpdatedBy string
		ID string
		ExtraField_Extra string
		ExtraField2_Extra string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	Template_Redirect = dm.Template_PathList
)

//Template_Publish annouces the endpoints available for this object
func Template_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.Template_PathList, Template_HandlerList)
	mux.HandleFunc(dm.Template_PathView, Template_HandlerView)
	mux.HandleFunc(dm.Template_PathEdit, Template_HandlerEdit)
	mux.HandleFunc(dm.Template_PathNew, Template_HandlerNew)
	mux.HandleFunc(dm.Template_PathSave, Template_HandlerSave)
	mux.HandleFunc(dm.Template_PathDelete, Template_HandlerDelete)
	logs.Publish("Application", dm.Template_Title)
    //No API
}

//Template_HandlerList is the handler for the list page
func Template_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Template
	noItems, returnList, _ := dao.Template_GetList()

	pageDetail := Template_PageList{
		Title:            CardTitle(dm.Template_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Template_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Template_TemplateList, w, r, pageDetail)

}

//Template_HandlerView is the handler used to View a page
func Template_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Template_QueryString)
	_, rD, _ := dao.Template_GetByID(searchID)

	pageDetail := Template_Page{
		Title:       CardTitle(dm.Template_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Template_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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

pageDetail.ExtraField_Extra = rD.ExtraField_Extra
pageDetail.ExtraField2_Extra = rD.ExtraField2_Extra

// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Template_TemplateView, w, r, pageDetail)

}

//Template_HandlerEdit is the handler used generate the Edit page
func Template_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Template_QueryString)
	_, rD, _ := dao.Template_GetByID(searchID)
	
	pageDetail := Template_Page{
		Title:       CardTitle(dm.Template_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Template_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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

pageDetail.ExtraField_Extra = rD.ExtraField_Extra
pageDetail.ExtraField2_Extra = rD.ExtraField2_Extra

// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Template_TemplateEdit, w, r, pageDetail)


}

//Template_HandlerSave is the handler used process the saving of an Template
func Template_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ID"))

	var item dm.Template
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.Template_SYSId)
		item.FIELD1 = r.FormValue(dm.Template_FIELD1)
		item.FIELD2 = r.FormValue(dm.Template_FIELD2)
		item.SYSCreated = r.FormValue(dm.Template_SYSCreated)
		item.SYSCreatedBy = r.FormValue(dm.Template_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.Template_SYSCreatedHost)
		item.SYSUpdated = r.FormValue(dm.Template_SYSUpdated)
		item.SYSUpdatedHost = r.FormValue(dm.Template_SYSUpdatedHost)
		item.SYSUpdatedBy = r.FormValue(dm.Template_SYSUpdatedBy)
		item.ID = r.FormValue(dm.Template_ID)
		item.ExtraField_Extra = r.FormValue(dm.Template_ExtraField_Extra)
		item.ExtraField2_Extra = r.FormValue(dm.Template_ExtraField2_Extra)
		
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.Template_Store(item,r)	

	http.Redirect(w, r, Template_Redirect, http.StatusFound)
}

//Template_HandlerNew is the handler used process the creation of an Template
func Template_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Template_Page{
		Title:       CardTitle(dm.Template_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Template_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.FIELD1 = ""
pageDetail.FIELD2 = ""
pageDetail.SYSCreated = ""
pageDetail.SYSCreatedBy = ""
pageDetail.SYSCreatedHost = ""
pageDetail.SYSUpdated = ""
pageDetail.SYSUpdatedHost = ""
pageDetail.SYSUpdatedBy = ""
pageDetail.ID = ""

pageDetail.ExtraField_Extra = ""
pageDetail.ExtraField2_Extra = ""

// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Template_TemplateNew, w, r, pageDetail)

}

//Template_HandlerDelete is the handler used process the deletion of an Template
func Template_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Template_QueryString)

	dao.Template_Delete(searchID)	

	http.Redirect(w, r, Template_Redirect, http.StatusFound)
}