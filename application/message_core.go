package application
// ----------------------------------------------------------------
// Automatically generated  "/application/message.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Message (message)
// Endpoint 	        : Message (Message)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:30
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//message_PageList provides the information for the template for a list of Messages
type Message_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Message
}
//Message_Redirect provides a page to return to aftern an action
const (
	
	Message_Redirect = dm.Message_PathList
	
)

//message_Page provides the information for the template for an individual Message
type Message_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	SYSId_props     dm.FieldProperties
	Id         string
	Id_props     dm.FieldProperties
	Message         string
	Message_props     dm.FieldProperties
	SYSCreated         string
	SYSCreated_props     dm.FieldProperties
	SYSWho         string
	SYSWho_props     dm.FieldProperties
	SYSHost         string
	SYSHost_props     dm.FieldProperties
	SYSUpdated         string
	SYSUpdated_props     dm.FieldProperties
	SYSCreatedBy         string
	SYSCreatedBy_props     dm.FieldProperties
	SYSCreatedHost         string
	SYSCreatedHost_props     dm.FieldProperties
	SYSUpdatedBy         string
	SYSUpdatedBy_props     dm.FieldProperties
	SYSUpdatedHost         string
	SYSUpdatedHost_props     dm.FieldProperties
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//Message_Publish annouces the endpoints available for this object
func Message_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.Message_PathList, Message_HandlerList)
	mux.HandleFunc(dm.Message_PathView, Message_HandlerView)
	mux.HandleFunc(dm.Message_PathEdit, Message_HandlerEdit)
	//Cannot Create via GUI
	mux.HandleFunc(dm.Message_PathSave, Message_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Message_Title)
    //No API
}


//Message_HandlerList is the handler for the list page
func Message_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Message
	noItems, returnList, _ := dao.Message_GetList()

	pageDetail := Message_PageList{
		Title:            CardTitle(dm.Message_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Message_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Message_TemplateList, w, r, pageDetail)

}


//Message_HandlerView is the handler used to View a page
func Message_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Message_QueryString)
	_, rD, _ := dao.Message_GetByID(searchID)

	pageDetail := Message_Page{
		Title:       CardTitle(dm.Message_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Message_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = message_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Message_TemplateView, w, r, pageDetail)

}


//Message_HandlerEdit is the handler used generate the Edit page
func Message_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Message_QueryString)
	_, rD, _ := dao.Message_GetByID(searchID)
	
	pageDetail := Message_Page{
		Title:       CardTitle(dm.Message_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Message_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = message_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Message_TemplateEdit, w, r, pageDetail)
}


//Message_HandlerSave is the handler used process the saving of an Message
func Message_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	var item dm.Message
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.Message_SYSId_scrn)
		item.Id = r.FormValue(dm.Message_Id_scrn)
		item.Message = r.FormValue(dm.Message_Message_scrn)
		item.SYSCreated = r.FormValue(dm.Message_SYSCreated_scrn)
		item.SYSWho = r.FormValue(dm.Message_SYSWho_scrn)
		item.SYSHost = r.FormValue(dm.Message_SYSHost_scrn)
		item.SYSUpdated = r.FormValue(dm.Message_SYSUpdated_scrn)
		item.SYSCreatedBy = r.FormValue(dm.Message_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.Message_SYSCreatedHost_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.Message_SYSUpdatedBy_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.Message_SYSUpdatedHost_scrn)
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Message_Store(item,r)	
	http.Redirect(w, r, Message_Redirect, http.StatusFound)
}




// Builds/Popuplates the Message Page 
func message_PopulatePage(rD dm.Message, pageDetail Message_Page) Message_Page {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.Id = rD.Id
	pageDetail.Message = rD.Message
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSWho = rD.SYSWho
	pageDetail.SYSHost = rD.SYSHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	
	
	//
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.Id_props = rD.Id_props
	pageDetail.Message_props = rD.Message_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSWho_props = rD.SYSWho_props
	pageDetail.SYSHost_props = rD.SYSHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	