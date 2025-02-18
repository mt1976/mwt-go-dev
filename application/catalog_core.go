package application
// ----------------------------------------------------------------
// Automatically generated  "/application/catalog.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Catalog (catalog)
// Endpoint 	        : Catalog (ID)
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





//Catalog_Publish annouces the endpoints available for this object
func Catalog_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Catalog_Path, Catalog_Handler)
	mux.HandleFunc(dm.Catalog_PathList, Catalog_HandlerList)
	mux.HandleFunc(dm.Catalog_PathView, Catalog_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Catalog_Title)
    core.Catalog_Add(dm.Catalog_Title, dm.Catalog_Path, "", dm.Catalog_QueryString, "Application")
}


//Catalog_HandlerList is the handler for the list page
func Catalog_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Catalog
	noItems, returnList, _ := dao.Catalog_GetList()

	pageDetail := dm.Catalog_PageList{
		Title:            CardTitle(dm.Catalog_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Catalog_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Catalog_TemplateList, w, r, pageDetail)

}


//Catalog_HandlerView is the handler used to View a page
func Catalog_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Catalog_QueryString)
	_, rD, _ := dao.Catalog_GetByID(searchID)

	pageDetail := dm.Catalog_Page{
		Title:       CardTitle(dm.Catalog_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Catalog_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = catalog_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Catalog_TemplateView, w, r, pageDetail)

}






// Builds/Popuplates the Catalog Page 
func catalog_PopulatePage(rD dm.Catalog, pageDetail dm.Catalog_Page) dm.Catalog_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.ID = rD.ID
	pageDetail.Endpoint = rD.Endpoint
	pageDetail.Descr = rD.Descr
	pageDetail.Query = rD.Query
	pageDetail.Source = rD.Source
	
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.ID_props = rD.ID_props
	pageDetail.Endpoint_props = rD.Endpoint_props
	pageDetail.Descr_props = rD.Descr_props
	pageDetail.Query_props = rD.Query_props
	pageDetail.Source_props = rD.Source_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	