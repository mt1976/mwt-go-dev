package application
// ----------------------------------------------------------------
// Automatically generated  "/application/schedule.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Schedule (schedule)
// Endpoint 	        : Schedule (Schedule)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 21/11/2021 at 15:44:05
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"html/template"
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//schedule_PageList provides the information for the template for a list of Schedules
type Schedule_PageList struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Schedule
}

//schedule_Page provides the information for the template for an individual Schedule
type Schedule_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	AppInternalID  string
	// Automatically generated 21/11/2021 by matttownsend on silicon.local - START
		SYSId string
		Id string
		Name string
		Description string
		Schedule string
		Started string
		Lastrun string
		Message string
		SYSCreated string
		SYSWho string
		SYSHost string
		SYSUpdated string
		Type string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdatedBy string
		SYSUpdatedHost string
		Human string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 21/11/2021 by matttownsend on silicon.local - END
}

const (
	Schedule_Redirect = dm.Schedule_PathList
)

//Schedule_Publish annouces the endpoints available for this object
func Schedule_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Schedule_PathList, Schedule_HandlerList)
	mux.HandleFunc(dm.Schedule_PathView, Schedule_HandlerView)
	
	
	
	
	logs.Publish("Application", dm.Schedule_Title)
	
	// schedule_PublishImpl should be specified in application/schedule_Impl.go
	// to provide the implementation for the special case.
	// override should return mux - override function should be defined as
	// schedule_PublishImpl(mux http.ServeMux) http.ServeMux {...}
	mux = schedule_PublishImpl(mux)
	logs.Publish("Application", dm.Schedule_Title + " Special")
	
}

//Schedule_HandlerList is the handler for the list page
func Schedule_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Schedule
	noItems, returnList, _ := dao.Schedule_GetList()


	pageDetail := Schedule_PageList{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.Schedule_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Schedule_TemplateList, core.GetUserRole(r)))
	t.Execute(w, pageDetail)
}

//Schedule_HandlerView is the handler used to View a page
func Schedule_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Schedule_QueryString)
	_, rD, _ := dao.Schedule_GetByID(searchID)

	pageDetail := Schedule_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Schedule_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:		     rD.AppInternalID,
	}

		// 
		// Automatically generated 21/11/2021 by matttownsend on silicon.local - START
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
// Automatically generated 21/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 21/11/2021 by matttownsend on silicon.local - END
		//


	// schedule_HandlerViewImpl should be specified in application/schedule_Impl.go
	// to provide the implementation for the special case.
	// override should return mux - override function should be defined as
	// func schedule_HandlerViewImpl(pageDetail Schedule_Page) Schedule_Page {return pageDetail}
	pageDetail = schedule_HandlerViewImpl(pageDetail)

	// Automatically generated 21/11/2021 by matttownsend on silicon.local - END


	t, _ := template.ParseFiles(core.GetTemplateID(dm.Schedule_TemplateView, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Schedule_HandlerEdit is the handler used generate the Edit page
func Schedule_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Schedule_QueryString)
	_, rD, _ := dao.Schedule_GetByID(searchID)
	
	pageDetail := Schedule_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Schedule_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:          rD.AppInternalID,
	}

		// 
		// Automatically generated 21/11/2021 by matttownsend on silicon.local - START
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
// Automatically generated 21/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 21/11/2021 by matttownsend on silicon.local - END
		//

	// schedule_HandlerEditImpl should be specified in application/schedule_Impl.go
	// to provide the implementation for the special case.
	// override should return mux - override function should be defined as
	// func schedule_HandlerEditImpl(pageDetail Schedule_Page) Schedule_Page {return pageDetail}
	pageDetail = schedule_HandlerEditImpl(pageDetail)

	// Automatically generated 21/11/2021 by matttownsend on silicon.local - END

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Schedule_TemplateEdit, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Schedule_HandlerSave is the handler used process the saving of an Schedule
func Schedule_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	var item dm.Schedule

	//item.AppInternalID = r.FormValue("AppInternalID")
	// Automatically generated 21/11/2021 by matttownsend on silicon.local - START
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
	
	// Automatically generated 21/11/2021 by matttownsend on silicon.local - END

	// schedule_HandlerSaveImpl should be specified in application/schedule_Impl.go
	// to provide the implementation for the special case.
	// override should return mux - override function should be defined as
	// func schedule_HandlerSaveImpl(item dm.Schedule) dm.Schedule {return item}
	item = schedule_HandlerSaveImpl(item)

	// Automatically generated 21/11/2021 by matttownsend on silicon.local - END

	dao.Schedule_Store(item)	

	http.Redirect(w, r, Schedule_Redirect, http.StatusFound)
}

//Schedule_HandlerNew is the handler used process the creation of an Schedule
func Schedule_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Schedule_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Schedule_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:   "NEW",	
	}

		// 
		// Automatically generated 21/11/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.Id = ""
pageDetail.Name = ""
pageDetail.Description = ""
pageDetail.Schedule = ""
pageDetail.Started = ""
pageDetail.Lastrun = ""
pageDetail.Message = ""
pageDetail.SYSCreated = ""
pageDetail.SYSWho = ""
pageDetail.SYSHost = ""
pageDetail.SYSUpdated = ""
pageDetail.Type = ""
pageDetail.SYSCreatedBy = ""
pageDetail.SYSCreatedHost = ""
pageDetail.SYSUpdatedBy = ""
pageDetail.SYSUpdatedHost = ""
pageDetail.Human = ""
// Automatically generated 21/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 21/11/2021 by matttownsend on silicon.local - END
		//

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Schedule_TemplateNew, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Schedule_HandlerDelete is the handler used process the deletion of an Schedule
func Schedule_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Schedule_QueryString)

	dao.Schedule_Delete(searchID)	

	http.Redirect(w, r, Schedule_Redirect, http.StatusFound)
}