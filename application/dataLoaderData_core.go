package application
// ----------------------------------------------------------------
// Automatically generated  "/application/dataloaderdata.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DataLoaderData (dataloaderdata)
// Endpoint 	        : DataLoaderData (Id)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:49
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)





//DataLoaderData_Publish annouces the endpoints available for this object
func DataLoaderData_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.DataLoaderData_PathList, DataLoaderData_HandlerList)
	mux.HandleFunc(dm.DataLoaderData_PathView, DataLoaderData_HandlerView)
	mux.HandleFunc(dm.DataLoaderData_PathEdit, DataLoaderData_HandlerEdit)
	//Cannot Create via GUI
	mux.HandleFunc(dm.DataLoaderData_PathSave, DataLoaderData_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.DataLoaderData_Title)
    //No API
}


//DataLoaderData_HandlerList is the handler for the list page
func DataLoaderData_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.DataLoaderData
	noItems, returnList, _ := dao.DataLoaderData_GetList()

	pageDetail := dm.DataLoaderData_PageList{
		Title:            CardTitle(dm.DataLoaderData_Title, core.Action_List),
		PageTitle:        PageTitle(dm.DataLoaderData_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.DataLoaderData_TemplateList, w, r, pageDetail)

}


//DataLoaderData_HandlerView is the handler used to View a page
func DataLoaderData_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DataLoaderData_QueryString)
	_, rD, _ := dao.DataLoaderData_GetByID(searchID)

	pageDetail := dm.DataLoaderData_Page{
		Title:       CardTitle(dm.DataLoaderData_Title, core.Action_View),
		PageTitle:   PageTitle(dm.DataLoaderData_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = dataloaderdata_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.DataLoaderData_TemplateView, w, r, pageDetail)

}


//DataLoaderData_HandlerEdit is the handler used generate the Edit page
func DataLoaderData_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DataLoaderData_QueryString)
	_, rD, _ := dao.DataLoaderData_GetByID(searchID)
	
	pageDetail := dm.DataLoaderData_Page{
		Title:       CardTitle(dm.DataLoaderData_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.DataLoaderData_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = dataloaderdata_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.DataLoaderData_TemplateEdit, w, r, pageDetail)
}


//DataLoaderData_HandlerSave is the handler used process the saving of an DataLoaderData
func DataLoaderData_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	var item dm.DataLoaderData
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.DataLoaderData_SYSId_scrn)
		item.Id = r.FormValue(dm.DataLoaderData_Id_scrn)
		item.Row = r.FormValue(dm.DataLoaderData_Row_scrn)
		item.Position = r.FormValue(dm.DataLoaderData_Position_scrn)
		item.Value = r.FormValue(dm.DataLoaderData_Value_scrn)
		item.Loader = r.FormValue(dm.DataLoaderData_Loader_scrn)
		item.SYSCreated = r.FormValue(dm.DataLoaderData_SYSCreated_scrn)
		item.SYSWho = r.FormValue(dm.DataLoaderData_SYSWho_scrn)
		item.SYSHost = r.FormValue(dm.DataLoaderData_SYSHost_scrn)
		item.SYSUpdated = r.FormValue(dm.DataLoaderData_SYSUpdated_scrn)
		item.Map = r.FormValue(dm.DataLoaderData_Map_scrn)
		item.SYSCreatedBy = r.FormValue(dm.DataLoaderData_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.DataLoaderData_SYSCreatedHost_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.DataLoaderData_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.DataLoaderData_SYSUpdatedHost_scrn)
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.DataLoaderData_Store(item,r)	
	http.Redirect(w, r, dm.DataLoaderData_Redirect, http.StatusFound)
}




// Builds/Popuplates the DataLoaderData Page 
func dataloaderdata_PopulatePage(rD dm.DataLoaderData, pageDetail dm.DataLoaderData_Page) dm.DataLoaderData_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.Id = rD.Id
	pageDetail.Row = rD.Row
	pageDetail.Position = rD.Position
	pageDetail.Value = rD.Value
	pageDetail.Loader = rD.Loader
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSWho = rD.SYSWho
	pageDetail.SYSHost = rD.SYSHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.Map = rD.Map
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.Loader_lookup = dao.DataLoader_GetLookup()
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.Id_props = rD.Id_props
	pageDetail.Row_props = rD.Row_props
	pageDetail.Position_props = rD.Position_props
	pageDetail.Value_props = rD.Value_props
	pageDetail.Loader_props = rD.Loader_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSWho_props = rD.SYSWho_props
	pageDetail.SYSHost_props = rD.SYSHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.Map_props = rD.Map_props
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