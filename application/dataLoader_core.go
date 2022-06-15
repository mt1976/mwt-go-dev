package application
// ----------------------------------------------------------------
// Automatically generated  "/application/dataloader.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DataLoader (dataloader)
// Endpoint 	        : DataLoader (Id)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:32:01
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//dataloader_PageList provides the information for the template for a list of DataLoaders
type DataLoader_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.DataLoader
}
//DataLoader_Redirect provides a page to return to aftern an action
const (
	DataLoader_Redirect = dm.DataLoader_PathList
)

//dataloader_Page provides the information for the template for an individual DataLoader
type DataLoader_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	Id         string
	Name         string
	Description         string
	Filename         string
	Lastrun         string
	SYSCreated         string
	SYSWho         string
	SYSHost         string
	SYSUpdated         string
	Type         string
	Instance         string
	Extension         string
	SYSCreatedBy         string
	SYSUpdatedHost         string
	SYSUpdatedBy         string
	SYSCreatedHost         string
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//DataLoader_Publish annouces the endpoints available for this object
func DataLoader_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.DataLoader_PathList, DataLoader_HandlerList)
	mux.HandleFunc(dm.DataLoader_PathView, DataLoader_HandlerView)
	mux.HandleFunc(dm.DataLoader_PathEdit, DataLoader_HandlerEdit)
	//Cannot Create via GUI
	mux.HandleFunc(dm.DataLoader_PathSave, DataLoader_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.DataLoader_Title)
    //No API
}


//DataLoader_HandlerList is the handler for the list page
func DataLoader_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.DataLoader
	noItems, returnList, _ := dao.DataLoader_GetList()

	pageDetail := DataLoader_PageList{
		Title:            CardTitle(dm.DataLoader_Title, core.Action_List),
		PageTitle:        PageTitle(dm.DataLoader_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.DataLoader_TemplateList, w, r, pageDetail)

}


//DataLoader_HandlerView is the handler used to View a page
func DataLoader_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DataLoader_QueryString)
	_, rD, _ := dao.DataLoader_GetByID(searchID)

	pageDetail := DataLoader_Page{
		Title:       CardTitle(dm.DataLoader_Title, core.Action_View),
		PageTitle:   PageTitle(dm.DataLoader_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = dataloader_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.DataLoader_TemplateView, w, r, pageDetail)

}


//DataLoader_HandlerEdit is the handler used generate the Edit page
func DataLoader_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DataLoader_QueryString)
	_, rD, _ := dao.DataLoader_GetByID(searchID)
	
	pageDetail := DataLoader_Page{
		Title:       CardTitle(dm.DataLoader_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.DataLoader_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = dataloader_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.DataLoader_TemplateEdit, w, r, pageDetail)
}


//DataLoader_HandlerSave is the handler used process the saving of an DataLoader
func DataLoader_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	var item dm.DataLoader
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.DataLoader_SYSId)
		item.Id = r.FormValue(dm.DataLoader_Id)
		item.Name = r.FormValue(dm.DataLoader_Name)
		item.Description = r.FormValue(dm.DataLoader_Description)
		item.Filename = r.FormValue(dm.DataLoader_Filename)
		item.Lastrun = r.FormValue(dm.DataLoader_Lastrun)
		item.SYSCreated = r.FormValue(dm.DataLoader_SYSCreated)
		item.SYSWho = r.FormValue(dm.DataLoader_SYSWho)
		item.SYSHost = r.FormValue(dm.DataLoader_SYSHost)
		item.SYSUpdated = r.FormValue(dm.DataLoader_SYSUpdated)
		item.Type = r.FormValue(dm.DataLoader_Type)
		item.Instance = r.FormValue(dm.DataLoader_Instance)
		item.Extension = r.FormValue(dm.DataLoader_Extension)
		item.SYSCreatedBy = r.FormValue(dm.DataLoader_SYSCreatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.DataLoader_SYSUpdatedHost)
		item.SYSUpdatedBy = r.FormValue(dm.DataLoader_SYSUpdatedBy)
		item.SYSCreatedHost = r.FormValue(dm.DataLoader_SYSCreatedHost)
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.DataLoader_Store(item,r)	
	http.Redirect(w, r, DataLoader_Redirect, http.StatusFound)
}




// Builds/Popuplates the DataLoader Page 
func dataloader_PopulatePage(rD dm.DataLoader, pageDetail DataLoader_Page) DataLoader_Page {
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.Id = rD.Id
	pageDetail.Name = rD.Name
	pageDetail.Description = rD.Description
	pageDetail.Filename = rD.Filename
	pageDetail.Lastrun = rD.Lastrun
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSWho = rD.SYSWho
	pageDetail.SYSHost = rD.SYSHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.Type = rD.Type
	pageDetail.Instance = rD.Instance
	pageDetail.Extension = rD.Extension
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	
	
	//
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	