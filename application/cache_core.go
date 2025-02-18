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
// Date & Time		    : 28/06/2022 at 16:10:44
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
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

	pageDetail := dm.Cache_PageList{
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

	pageDetail := dm.Cache_Page{
		Title:       CardTitle(dm.Cache_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Cache_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = cache_PopulatePage(rD , pageDetail) 

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
	
	pageDetail := dm.Cache_Page{
		Title:       CardTitle(dm.Cache_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Cache_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = cache_PopulatePage(rD , pageDetail) 

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
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.Cache_SYSId_scrn)
		item.Id = r.FormValue(dm.Cache_Id_scrn)
		item.Object = r.FormValue(dm.Cache_Object_scrn)
		item.Field = r.FormValue(dm.Cache_Field_scrn)
		item.Value = r.FormValue(dm.Cache_Value_scrn)
		item.Expiry = r.FormValue(dm.Cache_Expiry_scrn)
		item.SYSCreated = r.FormValue(dm.Cache_SYSCreated_scrn)
		item.SYSWho = r.FormValue(dm.Cache_SYSWho_scrn)
		item.SYSHost = r.FormValue(dm.Cache_SYSHost_scrn)
		item.SYSUpdated = r.FormValue(dm.Cache_SYSUpdated_scrn)
		item.Source = r.FormValue(dm.Cache_Source_scrn)
		item.SYSCreatedBy = r.FormValue(dm.Cache_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.Cache_SYSCreatedHost_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.Cache_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.Cache_SYSUpdatedHost_scrn)
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Cache_Store(item,r)	
	http.Redirect(w, r, dm.Cache_Redirect, http.StatusFound)
}




// Builds/Popuplates the Cache Page 
func cache_PopulatePage(rD dm.Cache, pageDetail dm.Cache_Page) dm.Cache_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
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
	
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.Id_props = rD.Id_props
	pageDetail.Object_props = rD.Object_props
	pageDetail.Field_props = rD.Field_props
	pageDetail.Value_props = rD.Value_props
	pageDetail.Expiry_props = rD.Expiry_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSWho_props = rD.SYSWho_props
	pageDetail.SYSHost_props = rD.SYSHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.Source_props = rD.Source_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	