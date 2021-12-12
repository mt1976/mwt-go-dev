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
// Date & Time		    : 12/12/2021 at 16:13:17
// Who & Where		    : matttownsend on silicon.local
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
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Message
}

//message_Page provides the information for the template for an individual Message
type Message_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		SYSId string
		Id string
		Message string
		SYSCreated string
		SYSWho string
		SYSHost string
		SYSUpdated string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdatedBy string
		SYSUpdatedHost string
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	Message_Redirect = dm.Message_PathList
)

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

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

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

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

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
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.Message_SYSId)
		item.Id = r.FormValue(dm.Message_Id)
		item.Message = r.FormValue(dm.Message_Message)
		item.SYSCreated = r.FormValue(dm.Message_SYSCreated)
		item.SYSWho = r.FormValue(dm.Message_SYSWho)
		item.SYSHost = r.FormValue(dm.Message_SYSHost)
		item.SYSUpdated = r.FormValue(dm.Message_SYSUpdated)
		item.SYSCreatedBy = r.FormValue(dm.Message_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.Message_SYSCreatedHost)
		item.SYSUpdatedBy = r.FormValue(dm.Message_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.Message_SYSUpdatedHost)
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.Message_Store(item,r)	

	http.Redirect(w, r, Message_Redirect, http.StatusFound)
}

//Message_HandlerNew is the handler used process the creation of an Message
func Message_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Message_Page{
		Title:       CardTitle(dm.Message_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Message_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.Id = ""
pageDetail.Message = ""
pageDetail.SYSCreated = ""
pageDetail.SYSWho = ""
pageDetail.SYSHost = ""
pageDetail.SYSUpdated = ""
pageDetail.SYSCreatedBy = ""
pageDetail.SYSCreatedHost = ""
pageDetail.SYSUpdatedBy = ""
pageDetail.SYSUpdatedHost = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Message_TemplateNew, w, r, pageDetail)

}

//Message_HandlerDelete is the handler used process the deletion of an Message
func Message_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Message_QueryString)

	dao.Message_Delete(searchID)	

	http.Redirect(w, r, Message_Redirect, http.StatusFound)
}