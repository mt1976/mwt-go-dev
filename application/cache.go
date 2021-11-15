package application
// ----------------------------------------------------------------
// Automatically generated  "/application/cache.go"
// ----------------------------------------------------------------
// Package            : application
// Object 			  : Cache
// Endpoint Root 	  : Cache
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

//cache_PageList provides the information for the template for a list of Caches
type cache_PageList struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Cache
}

//cache_Page provides the information for the template for an individual Cache
type cache_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	AppInternalID  string
	// Automatically generated 15/11/2021 by matttownsend on silicon.local - START
		SYSId string
		Id string
		Object string
		Field string
		Value string
		Expiry string
		SYSCreated string
		SYSWho string
		SYSHost string
		SYSUpdated string
		Source string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdatedBy string
		SYSUpdatedHost string
	
	// Automatically generated 15/11/2021 by matttownsend on silicon.local - END
}

const (
	Cache_Redirect = dm.Cache_PathList
)

//Cache_Publish annouces the endpoints available for this object
func Cache_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Cache_PathList, Cache_HandlerList)
	mux.HandleFunc(dm.Cache_PathView, Cache_HandlerView)
	mux.HandleFunc(dm.Cache_PathEdit, Cache_HandlerEdit)
	
	mux.HandleFunc(dm.Cache_PathSave, Cache_HandlerSave)
	
	logs.Publish("Application", dm.Cache_Title)
}

//Cache_HandlerList is the handler for the list page
func Cache_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Cache
	noItems, returnList, _ := dao.Cache_GetList()


	pageDetail := cache_PageList{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.Cache_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Cache_TemplateList, core.GetUserRole(r)))
	t.Execute(w, pageDetail)
}

//Cache_HandlerView is the handler used to View a page
func Cache_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Cache_QueryString)
	_, rD, _ := dao.Cache_GetByID(searchID)

	pageDetail := cache_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Cache_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:		     rD.AppInternalID,
		// 
		// Automatically generated 15/11/2021 by matttownsend on silicon.local - START
		SYSId: rD.SYSId,
		Id: rD.Id,
		Object: rD.Object,
		Field: rD.Field,
		Value: rD.Value,
		Expiry: rD.Expiry,
		SYSCreated: rD.SYSCreated,
		SYSWho: rD.SYSWho,
		SYSHost: rD.SYSHost,
		SYSUpdated: rD.SYSUpdated,
		Source: rD.Source,
		SYSCreatedBy: rD.SYSCreatedBy,
		SYSCreatedHost: rD.SYSCreatedHost,
		SYSUpdatedBy: rD.SYSUpdatedBy,
		SYSUpdatedHost: rD.SYSUpdatedHost,
		
		// Automatically generated 15/11/2021 by matttownsend on silicon.local - END
		//
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Cache_TemplateView, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Cache_HandlerEdit is the handler used generate the Edit page
func Cache_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Cache_QueryString)
	_, rD, _ := dao.Cache_GetByID(searchID)
	
	pageDetail := cache_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Cache_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:          rD.AppInternalID,
		// Automatically generated 15/11/2021 by matttownsend on silicon.local - START
			SYSId: rD.SYSId,
			Id: rD.Id,
			Object: rD.Object,
			Field: rD.Field,
			Value: rD.Value,
			Expiry: rD.Expiry,
			SYSCreated: rD.SYSCreated,
			SYSWho: rD.SYSWho,
			SYSHost: rD.SYSHost,
			SYSUpdated: rD.SYSUpdated,
			Source: rD.Source,
			SYSCreatedBy: rD.SYSCreatedBy,
			SYSCreatedHost: rD.SYSCreatedHost,
			SYSUpdatedBy: rD.SYSUpdatedBy,
			SYSUpdatedHost: rD.SYSUpdatedHost,
		
	// Automatically generated 15/11/2021 by matttownsend on silicon.local - END
		//Post Import Actions - START

		// Post Import Actions - END
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Cache_TemplateEdit, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Cache_HandlerSave is the handler used process the saving of an Cache
func Cache_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ID"))

	var item dm.Cache

	//item.AppInternalID = r.FormValue("AppInternalID")
	// Automatically generated 15/11/2021 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue("SYSId")
		item.Id = r.FormValue("Id")
		item.Object = r.FormValue("Object")
		item.Field = r.FormValue("Field")
		item.Value = r.FormValue("Value")
		item.Expiry = r.FormValue("Expiry")
		item.SYSCreated = r.FormValue("SYSCreated")
		item.SYSWho = r.FormValue("SYSWho")
		item.SYSHost = r.FormValue("SYSHost")
		item.SYSUpdated = r.FormValue("SYSUpdated")
		item.Source = r.FormValue("Source")
		item.SYSCreatedBy = r.FormValue("SYSCreatedBy")
		item.SYSCreatedHost = r.FormValue("SYSCreatedHost")
		item.SYSUpdatedBy = r.FormValue("SYSUpdatedBy")
		item.SYSUpdatedHost = r.FormValue("SYSUpdatedHost")
	
	// Automatically generated 15/11/2021 by matttownsend on silicon.local - END

	dao.Cache_Store(item)	

	http.Redirect(w, r, Cache_Redirect, http.StatusFound)
}

//Cache_HandlerNew is the handler used process the creation of an Cache
func Cache_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := cache_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Cache_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:   "NEW",
		// Automatically generated 15/11/2021 by matttownsend on silicon.local - START
			SYSId: "0",
			Id: "",
			Object: "",
			Field: "",
			Value: "",
			Expiry: "",
			SYSCreated: "",
			SYSWho: "",
			SYSHost: "",
			SYSUpdated: "",
			Source: "",
			SYSCreatedBy: "",
			SYSCreatedHost: "",
			SYSUpdatedBy: "",
			SYSUpdatedHost: "",
		
		// Automatically generated 15/11/2021 by matttownsend on silicon.local - END
		//
		// Post Import Actions - START

		// Post Import Actions - END
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Cache_TemplateNew, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Cache_HandlerDelete is the handler used process the deletion of an Cache
func Cache_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Cache_QueryString)

	dao.Cache_Delete(searchID)	

	http.Redirect(w, r, Cache_Redirect, http.StatusFound)
}