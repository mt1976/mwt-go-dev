package application
// ----------------------------------------------------------------
// Automatically generated  "/application/translation.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Translation (translation)
// Endpoint 	        : Translation (Message)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 21/11/2021 at 15:44:05
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"html/template"
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//translation_PageList provides the information for the template for a list of Translations
type Translation_PageList struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Translation
}

//translation_Page provides the information for the template for an individual Translation
type Translation_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	AppInternalID  string
	// Automatically generated 21/11/2021 by matttownsend on silicon.local - START
		SYSId string
		Id string
		Class string
		Message string
		Translation string
		SYSCreated string
		SYSWho string
		SYSHost string
		SYSUpdated string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdatedBy string
		SYSUpdatedHost string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 21/11/2021 by matttownsend on silicon.local - END
}

const (
	Translation_Redirect = dm.Translation_PathList
)

//Translation_Publish annouces the endpoints available for this object
func Translation_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Translation_PathList, Translation_HandlerList)
	mux.HandleFunc(dm.Translation_PathView, Translation_HandlerView)
	mux.HandleFunc(dm.Translation_PathEdit, Translation_HandlerEdit)
	
	mux.HandleFunc(dm.Translation_PathSave, Translation_HandlerSave)
	
	logs.Publish("Application", dm.Translation_Title)
	
}

//Translation_HandlerList is the handler for the list page
func Translation_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Translation
	noItems, returnList, _ := dao.Translation_GetList()


	pageDetail := Translation_PageList{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.Translation_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Translation_TemplateList, core.GetUserRole(r)))
	t.Execute(w, pageDetail)
}

//Translation_HandlerView is the handler used to View a page
func Translation_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Translation_QueryString)
	_, rD, _ := dao.Translation_GetByID(searchID)

	pageDetail := Translation_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Translation_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:		     rD.AppInternalID,
	}

		// 
		// Automatically generated 21/11/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Id = rD.Id
pageDetail.Class = rD.Class
pageDetail.Message = rD.Message
pageDetail.Translation = rD.Translation
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSWho = rD.SYSWho
pageDetail.SYSHost = rD.SYSHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
// Automatically generated 21/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 21/11/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 21/11/2021 by matttownsend on silicon.local - END


	t, _ := template.ParseFiles(core.GetTemplateID(dm.Translation_TemplateView, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Translation_HandlerEdit is the handler used generate the Edit page
func Translation_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Translation_QueryString)
	_, rD, _ := dao.Translation_GetByID(searchID)
	
	pageDetail := Translation_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Translation_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:          rD.AppInternalID,
	}

		// 
		// Automatically generated 21/11/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Id = rD.Id
pageDetail.Class = rD.Class
pageDetail.Message = rD.Message
pageDetail.Translation = rD.Translation
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSWho = rD.SYSWho
pageDetail.SYSHost = rD.SYSHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
// Automatically generated 21/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 21/11/2021 by matttownsend on silicon.local - END
		//

	// Automatically generated 21/11/2021 by matttownsend on silicon.local - END

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Translation_TemplateEdit, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Translation_HandlerSave is the handler used process the saving of an Translation
func Translation_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	var item dm.Translation

	//item.AppInternalID = r.FormValue("AppInternalID")
	// Automatically generated 21/11/2021 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.Translation_SYSId)
		item.Id = r.FormValue(dm.Translation_Id)
		item.Class = r.FormValue(dm.Translation_Class)
		item.Message = r.FormValue(dm.Translation_Message)
		item.Translation = r.FormValue(dm.Translation_Translation)
		item.SYSCreated = r.FormValue(dm.Translation_SYSCreated)
		item.SYSWho = r.FormValue(dm.Translation_SYSWho)
		item.SYSHost = r.FormValue(dm.Translation_SYSHost)
		item.SYSUpdated = r.FormValue(dm.Translation_SYSUpdated)
		item.SYSCreatedBy = r.FormValue(dm.Translation_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.Translation_SYSCreatedHost)
		item.SYSUpdatedBy = r.FormValue(dm.Translation_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.Translation_SYSUpdatedHost)
	
	// Automatically generated 21/11/2021 by matttownsend on silicon.local - END

	// Automatically generated 21/11/2021 by matttownsend on silicon.local - END

	dao.Translation_Store(item)	

	http.Redirect(w, r, Translation_Redirect, http.StatusFound)
}

//Translation_HandlerNew is the handler used process the creation of an Translation
func Translation_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Translation_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Translation_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:   "NEW",	
	}

		// 
		// Automatically generated 21/11/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.Id = ""
pageDetail.Class = ""
pageDetail.Message = ""
pageDetail.Translation = ""
pageDetail.SYSCreated = ""
pageDetail.SYSWho = ""
pageDetail.SYSHost = ""
pageDetail.SYSUpdated = ""
pageDetail.SYSCreatedBy = ""
pageDetail.SYSCreatedHost = ""
pageDetail.SYSUpdatedBy = ""
pageDetail.SYSUpdatedHost = ""
// Automatically generated 21/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 21/11/2021 by matttownsend on silicon.local - END
		//

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Translation_TemplateNew, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Translation_HandlerDelete is the handler used process the deletion of an Translation
func Translation_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Translation_QueryString)

	dao.Translation_Delete(searchID)	

	http.Redirect(w, r, Translation_Redirect, http.StatusFound)
}