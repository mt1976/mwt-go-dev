package application
// ----------------------------------------------------------------
// Automatically generated  "/application/schedule.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Schedule (schedule)
// Endpoint 	        : Schedule (Schedule)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:32:08
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//schedule_PageList provides the information for the template for a list of Schedules
type Schedule_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Schedule
}
//Schedule_Redirect provides a page to return to aftern an action
const (
	Schedule_Redirect = dm.Schedule_PathList
)

//schedule_Page provides the information for the template for an individual Schedule
type Schedule_Page struct {
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
	Schedule         string
	Started         string
	Lastrun         string
	Message         string
	SYSCreated         string
	SYSWho         string
	SYSHost         string
	SYSUpdated         string
	Type         string
	SYSCreatedBy         string
	SYSCreatedHost         string
	SYSUpdatedBy         string
	SYSUpdatedHost         string
	Human         string
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//Schedule_Publish annouces the endpoints available for this object
func Schedule_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Schedule_Path, Schedule_Handler)
	mux.HandleFunc(dm.Schedule_PathList, Schedule_HandlerList)
	mux.HandleFunc(dm.Schedule_PathView, Schedule_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	mux.HandleFunc(dm.Schedule_PathSave, Schedule_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Schedule_Title)
    core.Catalog_Add(dm.Schedule_Title, dm.Schedule_Path, "", dm.Schedule_QueryString, "Application")
}


//Schedule_HandlerList is the handler for the list page
func Schedule_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Schedule
	noItems, returnList, _ := dao.Schedule_GetList()

	pageDetail := Schedule_PageList{
		Title:            CardTitle(dm.Schedule_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Schedule_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Schedule_TemplateList, w, r, pageDetail)

}


//Schedule_HandlerView is the handler used to View a page
func Schedule_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Schedule_QueryString)
	_, rD, _ := dao.Schedule_GetByID(searchID)

	pageDetail := Schedule_Page{
		Title:       CardTitle(dm.Schedule_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Schedule_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = schedule_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Schedule_TemplateView, w, r, pageDetail)

}



//Schedule_HandlerSave is the handler used process the saving of an Schedule
func Schedule_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	var item dm.Schedule
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.Schedule_SYSId)
		item.Id = r.FormValue(dm.Schedule_Id)
		item.Name = r.FormValue(dm.Schedule_Name)
		item.Description = r.FormValue(dm.Schedule_Description)
		item.Schedule = r.FormValue(dm.Schedule_Schedule)
		item.Started = r.FormValue(dm.Schedule_Started)
		item.Lastrun = r.FormValue(dm.Schedule_Lastrun)
		item.Message = r.FormValue(dm.Schedule_Message)
		item.SYSCreated = r.FormValue(dm.Schedule_SYSCreated)
		item.SYSWho = r.FormValue(dm.Schedule_SYSWho)
		item.SYSHost = r.FormValue(dm.Schedule_SYSHost)
		item.SYSUpdated = r.FormValue(dm.Schedule_SYSUpdated)
		item.Type = r.FormValue(dm.Schedule_Type)
		item.SYSCreatedBy = r.FormValue(dm.Schedule_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.Schedule_SYSCreatedHost)
		item.SYSUpdatedBy = r.FormValue(dm.Schedule_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.Schedule_SYSUpdatedHost)
		item.Human = r.FormValue(dm.Schedule_Human)
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Schedule_Store(item,r)	
	http.Redirect(w, r, Schedule_Redirect, http.StatusFound)
}




// Builds/Popuplates the Schedule Page 
func schedule_PopulatePage(rD dm.Schedule, pageDetail Schedule_Page) Schedule_Page {
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.Id = rD.Id
	pageDetail.Name = rD.Name
	pageDetail.Description = rD.Description
	pageDetail.Schedule = rD.Schedule
	pageDetail.Started = rD.Started
	pageDetail.Lastrun = rD.Lastrun
	pageDetail.Message = rD.Message
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSWho = rD.SYSWho
	pageDetail.SYSHost = rD.SYSHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.Type = rD.Type
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.Human = rD.Human
	
	
	//
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	