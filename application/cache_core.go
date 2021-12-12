package application
// ----------------------------------------------------------------
// Automatically generated  "/application/cache.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Cache (cache)
// Endpoint 	        : Cache (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:08
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//cache_PageList provides the information for the template for a list of Caches
type Cache_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Cache
}

//cache_Page provides the information for the template for an individual Cache
type Cache_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	Cache_Redirect = dm.Cache_PathList
)

//Cache_Publish annouces the endpoints available for this object
func Cache_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.Cache_PathList, Cache_HandlerList)
	mux.HandleFunc(dm.Cache_PathView, Cache_HandlerView)
	mux.HandleFunc(dm.Cache_PathEdit, Cache_HandlerEdit)
	//Cannot Create via GUI
	mux.HandleFunc(dm.Cache_PathSave, Cache_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Cache_Title)
    //No API
}

//Cache_HandlerList is the handler for the list page
func Cache_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Cache
	noItems, returnList, _ := dao.Cache_GetList()

	pageDetail := Cache_PageList{
		Title:            CardTitle(dm.Cache_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Cache_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Cache_TemplateList, w, r, pageDetail)

}

//Cache_HandlerView is the handler used to View a page
func Cache_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Cache_QueryString)
	_, rD, _ := dao.Cache_GetByID(searchID)

	pageDetail := Cache_Page{
		Title:       CardTitle(dm.Cache_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Cache_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Id = rD.Id
pageDetail.Object = rD.Object
pageDetail.Field = rD.Field
pageDetail.Value = rD.Value
pageDetail.Expiry = rD.Expiry
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSWho = rD.SYSWho
pageDetail.SYSHost = rD.SYSHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.Source = rD.Source
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Cache_TemplateView, w, r, pageDetail)

}

//Cache_HandlerEdit is the handler used generate the Edit page
func Cache_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Cache_QueryString)
	_, rD, _ := dao.Cache_GetByID(searchID)
	
	pageDetail := Cache_Page{
		Title:       CardTitle(dm.Cache_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Cache_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Id = rD.Id
pageDetail.Object = rD.Object
pageDetail.Field = rD.Field
pageDetail.Value = rD.Value
pageDetail.Expiry = rD.Expiry
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSWho = rD.SYSWho
pageDetail.SYSHost = rD.SYSHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.Source = rD.Source
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Cache_TemplateEdit, w, r, pageDetail)


}

//Cache_HandlerSave is the handler used process the saving of an Cache
func Cache_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	var item dm.Cache
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.Cache_SYSId)
		item.Id = r.FormValue(dm.Cache_Id)
		item.Object = r.FormValue(dm.Cache_Object)
		item.Field = r.FormValue(dm.Cache_Field)
		item.Value = r.FormValue(dm.Cache_Value)
		item.Expiry = r.FormValue(dm.Cache_Expiry)
		item.SYSCreated = r.FormValue(dm.Cache_SYSCreated)
		item.SYSWho = r.FormValue(dm.Cache_SYSWho)
		item.SYSHost = r.FormValue(dm.Cache_SYSHost)
		item.SYSUpdated = r.FormValue(dm.Cache_SYSUpdated)
		item.Source = r.FormValue(dm.Cache_Source)
		item.SYSCreatedBy = r.FormValue(dm.Cache_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.Cache_SYSCreatedHost)
		item.SYSUpdatedBy = r.FormValue(dm.Cache_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.Cache_SYSUpdatedHost)
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.Cache_Store(item,r)	

	http.Redirect(w, r, Cache_Redirect, http.StatusFound)
}

//Cache_HandlerNew is the handler used process the creation of an Cache
func Cache_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Cache_Page{
		Title:       CardTitle(dm.Cache_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Cache_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.Id = ""
pageDetail.Object = ""
pageDetail.Field = ""
pageDetail.Value = ""
pageDetail.Expiry = ""
pageDetail.SYSCreated = ""
pageDetail.SYSWho = ""
pageDetail.SYSHost = ""
pageDetail.SYSUpdated = ""
pageDetail.Source = ""
pageDetail.SYSCreatedBy = ""
pageDetail.SYSCreatedHost = ""
pageDetail.SYSUpdatedBy = ""
pageDetail.SYSUpdatedHost = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Cache_TemplateNew, w, r, pageDetail)

}

//Cache_HandlerDelete is the handler used process the deletion of an Cache
func Cache_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Cache_QueryString)

	dao.Cache_Delete(searchID)	

	http.Redirect(w, r, Cache_Redirect, http.StatusFound)
}