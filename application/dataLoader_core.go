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

//dataloader_PageList provides the information for the template for a list of DataLoaders
type DataLoader_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.DataLoader
}

//dataloader_Page provides the information for the template for an individual DataLoader
type DataLoader_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		SYSId string
		Id string
		Name string
		Description string
		Filename string
		Lastrun string
		SYSCreated string
		SYSWho string
		SYSHost string
		SYSUpdated string
		Type string
		Instance string
		Extension string
		SYSCreatedBy string
		SYSUpdatedHost string
		SYSUpdatedBy string
		SYSCreatedHost string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	DataLoader_Redirect = dm.DataLoader_PathList
)

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

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

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

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

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
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.DataLoader_Store(item,r)	

	http.Redirect(w, r, DataLoader_Redirect, http.StatusFound)
}

//DataLoader_HandlerNew is the handler used process the creation of an DataLoader
func DataLoader_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := DataLoader_Page{
		Title:       CardTitle(dm.DataLoader_Title, core.Action_New),
		PageTitle:   PageTitle(dm.DataLoader_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.Id = ""
pageDetail.Name = ""
pageDetail.Description = ""
pageDetail.Filename = ""
pageDetail.Lastrun = ""
pageDetail.SYSCreated = ""
pageDetail.SYSWho = ""
pageDetail.SYSHost = ""
pageDetail.SYSUpdated = ""
pageDetail.Type = ""
pageDetail.Instance = ""
pageDetail.Extension = ""
pageDetail.SYSCreatedBy = ""
pageDetail.SYSUpdatedHost = ""
pageDetail.SYSUpdatedBy = ""
pageDetail.SYSCreatedHost = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.DataLoader_TemplateNew, w, r, pageDetail)

}

//DataLoader_HandlerDelete is the handler used process the deletion of an DataLoader
func DataLoader_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.DataLoader_QueryString)

	dao.DataLoader_Delete(searchID)	

	http.Redirect(w, r, DataLoader_Redirect, http.StatusFound)
}