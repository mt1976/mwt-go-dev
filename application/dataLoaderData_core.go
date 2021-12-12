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
// Date & Time		    : 12/12/2021 at 16:13:12
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//dataloaderdata_PageList provides the information for the template for a list of DataLoaderDatas
type DataLoaderData_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.DataLoaderData
}

//dataloaderdata_Page provides the information for the template for an individual DataLoaderData
type DataLoaderData_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		SYSId string
		Id string
		Row string
		Position string
		Value string
		Loader string
		SYSCreated string
		SYSWho string
		SYSHost string
		SYSUpdated string
		Map string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdatedBy string
		SYSUpdatedHost string
		Loader_Lookup string
		LoaderDescription_Lookup string
		LoaderType_Lookup string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	Loader_Lookup_List	[]dm.DataLoader
	LoaderDescription_Lookup_List	[]dm.DataLoader
	LoaderType_Lookup_List	[]dm.DataLoader
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	DataLoaderData_Redirect = dm.DataLoaderData_PathList
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

	pageDetail := DataLoaderData_PageList{
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

	pageDetail := DataLoaderData_Page{
		Title:       CardTitle(dm.DataLoaderData_Title, core.Action_View),
		PageTitle:   PageTitle(dm.DataLoaderData_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,Loader_Lookup_Name,_:= dao.DataLoader_GetByID(rD.Loader)
pageDetail.Loader_Lookup = Loader_Lookup_Name.Name
_,Loader_Lookup_Description,_:= dao.DataLoader_GetByID(rD.Loader)
pageDetail.LoaderDescription_Lookup = Loader_Lookup_Description.Description
_,Loader_Lookup_Type,_:= dao.DataLoader_GetByID(rD.Loader)
pageDetail.LoaderType_Lookup = Loader_Lookup_Type.Type
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

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
	
	pageDetail := DataLoaderData_Page{
		Title:       CardTitle(dm.DataLoaderData_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.DataLoaderData_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,Loader_Lookup_Name,_:= dao.DataLoader_GetByID(rD.Loader)
pageDetail.Loader_Lookup = Loader_Lookup_Name.Name
_,pageDetail.Loader_Lookup_List,_ = dao.DataLoader_GetList()
_,Loader_Lookup_Description,_:= dao.DataLoader_GetByID(rD.Loader)
pageDetail.LoaderDescription_Lookup = Loader_Lookup_Description.Description
_,pageDetail.LoaderDescription_Lookup_List,_ = dao.DataLoader_GetList()
_,Loader_Lookup_Type,_:= dao.DataLoader_GetByID(rD.Loader)
pageDetail.LoaderType_Lookup = Loader_Lookup_Type.Type
_,pageDetail.LoaderType_Lookup_List,_ = dao.DataLoader_GetList()
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

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
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.DataLoaderData_SYSId)
		item.Id = r.FormValue(dm.DataLoaderData_Id)
		item.Row = r.FormValue(dm.DataLoaderData_Row)
		item.Position = r.FormValue(dm.DataLoaderData_Position)
		item.Value = r.FormValue(dm.DataLoaderData_Value)
		item.Loader = r.FormValue(dm.DataLoaderData_Loader)
		item.SYSCreated = r.FormValue(dm.DataLoaderData_SYSCreated)
		item.SYSWho = r.FormValue(dm.DataLoaderData_SYSWho)
		item.SYSHost = r.FormValue(dm.DataLoaderData_SYSHost)
		item.SYSUpdated = r.FormValue(dm.DataLoaderData_SYSUpdated)
		item.Map = r.FormValue(dm.DataLoaderData_Map)
		item.SYSCreatedBy = r.FormValue(dm.DataLoaderData_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.DataLoaderData_SYSCreatedHost)
		item.SYSUpdatedBy = r.FormValue(dm.DataLoaderData_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.DataLoaderData_SYSUpdatedHost)
		item.Loader_Lookup = r.FormValue(dm.DataLoaderData_Loader_Lookup)
		item.LoaderDescription_Lookup = r.FormValue(dm.DataLoaderData_LoaderDescription_Lookup)
		item.LoaderType_Lookup = r.FormValue(dm.DataLoaderData_LoaderType_Lookup)
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.DataLoaderData_Store(item,r)	

	http.Redirect(w, r, DataLoaderData_Redirect, http.StatusFound)
}

//DataLoaderData_HandlerNew is the handler used process the creation of an DataLoaderData
func DataLoaderData_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := DataLoaderData_Page{
		Title:       CardTitle(dm.DataLoaderData_Title, core.Action_New),
		PageTitle:   PageTitle(dm.DataLoaderData_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.Id = ""
pageDetail.Row = ""
pageDetail.Position = ""
pageDetail.Value = ""
pageDetail.Loader = ""
pageDetail.SYSCreated = ""
pageDetail.SYSWho = ""
pageDetail.SYSHost = ""
pageDetail.SYSUpdated = ""
pageDetail.Map = ""
pageDetail.SYSCreatedBy = ""
pageDetail.SYSCreatedHost = ""
pageDetail.SYSUpdatedBy = ""
pageDetail.SYSUpdatedHost = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
pageDetail.Loader_Lookup = ""
_,pageDetail.Loader_Lookup_List,_ = dao.DataLoader_GetList()
pageDetail.LoaderDescription_Lookup = ""
_,pageDetail.LoaderDescription_Lookup_List,_ = dao.DataLoader_GetList()
pageDetail.LoaderType_Lookup = ""
_,pageDetail.LoaderType_Lookup_List,_ = dao.DataLoader_GetList()
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.DataLoaderData_TemplateNew, w, r, pageDetail)

}

//DataLoaderData_HandlerDelete is the handler used process the deletion of an DataLoaderData
func DataLoaderData_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.DataLoaderData_QueryString)

	dao.DataLoaderData_Delete(searchID)	

	http.Redirect(w, r, DataLoaderData_Redirect, http.StatusFound)
}