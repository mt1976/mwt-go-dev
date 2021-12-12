package application
// ----------------------------------------------------------------
// Automatically generated  "/application/dataloadermap.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DataLoaderMap (dataloadermap)
// Endpoint 	        : DataLoaderMap (Id)
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

//dataloadermap_PageList provides the information for the template for a list of DataLoaderMaps
type DataLoaderMap_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.DataLoaderMap
}

//dataloadermap_Page provides the information for the template for an individual DataLoaderMap
type DataLoaderMap_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		SYSId string
		Id string
		Name string
		Position string
		Loader string
		SYSCreated string
		SYSWho string
		SYSHost string
		SYSUpdated string
		Int_position string
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
	DataLoaderMap_Redirect = dm.DataLoaderMap_PathList
)

//DataLoaderMap_Publish annouces the endpoints available for this object
func DataLoaderMap_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.DataLoaderMap_PathList, DataLoaderMap_HandlerList)
	mux.HandleFunc(dm.DataLoaderMap_PathView, DataLoaderMap_HandlerView)
	mux.HandleFunc(dm.DataLoaderMap_PathEdit, DataLoaderMap_HandlerEdit)
	//Cannot Create via GUI
	mux.HandleFunc(dm.DataLoaderMap_PathSave, DataLoaderMap_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.DataLoaderMap_Title)
    //No API
}

//DataLoaderMap_HandlerList is the handler for the list page
func DataLoaderMap_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.DataLoaderMap
	noItems, returnList, _ := dao.DataLoaderMap_GetList()

	pageDetail := DataLoaderMap_PageList{
		Title:            CardTitle(dm.DataLoaderMap_Title, core.Action_List),
		PageTitle:        PageTitle(dm.DataLoaderMap_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.DataLoaderMap_TemplateList, w, r, pageDetail)

}

//DataLoaderMap_HandlerView is the handler used to View a page
func DataLoaderMap_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DataLoaderMap_QueryString)
	_, rD, _ := dao.DataLoaderMap_GetByID(searchID)

	pageDetail := DataLoaderMap_Page{
		Title:       CardTitle(dm.DataLoaderMap_Title, core.Action_View),
		PageTitle:   PageTitle(dm.DataLoaderMap_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Id = rD.Id
pageDetail.Name = rD.Name
pageDetail.Position = rD.Position
pageDetail.Loader = rD.Loader
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSWho = rD.SYSWho
pageDetail.SYSHost = rD.SYSHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.Int_position = rD.Int_position
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

	ExecuteTemplate(dm.DataLoaderMap_TemplateView, w, r, pageDetail)

}

//DataLoaderMap_HandlerEdit is the handler used generate the Edit page
func DataLoaderMap_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DataLoaderMap_QueryString)
	_, rD, _ := dao.DataLoaderMap_GetByID(searchID)
	
	pageDetail := DataLoaderMap_Page{
		Title:       CardTitle(dm.DataLoaderMap_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.DataLoaderMap_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Id = rD.Id
pageDetail.Name = rD.Name
pageDetail.Position = rD.Position
pageDetail.Loader = rD.Loader
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSWho = rD.SYSWho
pageDetail.SYSHost = rD.SYSHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.Int_position = rD.Int_position
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

	ExecuteTemplate(dm.DataLoaderMap_TemplateEdit, w, r, pageDetail)


}

//DataLoaderMap_HandlerSave is the handler used process the saving of an DataLoaderMap
func DataLoaderMap_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	var item dm.DataLoaderMap
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.DataLoaderMap_SYSId)
		item.Id = r.FormValue(dm.DataLoaderMap_Id)
		item.Name = r.FormValue(dm.DataLoaderMap_Name)
		item.Position = r.FormValue(dm.DataLoaderMap_Position)
		item.Loader = r.FormValue(dm.DataLoaderMap_Loader)
		item.SYSCreated = r.FormValue(dm.DataLoaderMap_SYSCreated)
		item.SYSWho = r.FormValue(dm.DataLoaderMap_SYSWho)
		item.SYSHost = r.FormValue(dm.DataLoaderMap_SYSHost)
		item.SYSUpdated = r.FormValue(dm.DataLoaderMap_SYSUpdated)
		item.Int_position = r.FormValue(dm.DataLoaderMap_Int_position)
		item.SYSCreatedBy = r.FormValue(dm.DataLoaderMap_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.DataLoaderMap_SYSCreatedHost)
		item.SYSUpdatedBy = r.FormValue(dm.DataLoaderMap_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.DataLoaderMap_SYSUpdatedHost)
		item.Loader_Lookup = r.FormValue(dm.DataLoaderMap_Loader_Lookup)
		item.LoaderDescription_Lookup = r.FormValue(dm.DataLoaderMap_LoaderDescription_Lookup)
		item.LoaderType_Lookup = r.FormValue(dm.DataLoaderMap_LoaderType_Lookup)
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.DataLoaderMap_Store(item,r)	

	http.Redirect(w, r, DataLoaderMap_Redirect, http.StatusFound)
}

//DataLoaderMap_HandlerNew is the handler used process the creation of an DataLoaderMap
func DataLoaderMap_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := DataLoaderMap_Page{
		Title:       CardTitle(dm.DataLoaderMap_Title, core.Action_New),
		PageTitle:   PageTitle(dm.DataLoaderMap_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.Id = ""
pageDetail.Name = ""
pageDetail.Position = ""
pageDetail.Loader = ""
pageDetail.SYSCreated = ""
pageDetail.SYSWho = ""
pageDetail.SYSHost = ""
pageDetail.SYSUpdated = ""
pageDetail.Int_position = ""
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

	ExecuteTemplate(dm.DataLoaderMap_TemplateNew, w, r, pageDetail)

}

//DataLoaderMap_HandlerDelete is the handler used process the deletion of an DataLoaderMap
func DataLoaderMap_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.DataLoaderMap_QueryString)

	dao.DataLoaderMap_Delete(searchID)	

	http.Redirect(w, r, DataLoaderMap_Redirect, http.StatusFound)
}