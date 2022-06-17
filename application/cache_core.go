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
// Date & Time		    : 17/06/2022 at 18:38:06
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
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
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Cache
}
//Cache_Redirect provides a page to return to aftern an action
const (
	Cache_Redirect = dm.Cache_PathList
)

//cache_Page provides the information for the template for an individual Cache
type Cache_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	Id         string
	Object         string
	Field         string
	Value         string
	Expiry         string
	SYSCreated         string
	SYSWho         string
	SYSHost         string
	SYSUpdated         string
	Source         string
	SYSCreatedBy         string
	SYSCreatedHost         string
	SYSUpdatedBy         string
	SYSUpdatedHost         string
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



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
	
	pageDetail := Cache_Page{
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
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Cache_Store(item,r)	
	http.Redirect(w, r, Cache_Redirect, http.StatusFound)
}




// Builds/Popuplates the Cache Page 
func cache_PopulatePage(rD dm.Cache, pageDetail Cache_Page) Cache_Page {
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	// Automatically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	