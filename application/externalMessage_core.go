package application
// ----------------------------------------------------------------
// Automatically generated  "/application/externalmessage.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : ExternalMessage (externalmessage)
// Endpoint 	        : ExternalMessage (MessageID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/12/2021 at 09:31:50
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//externalmessage_PageList provides the information for the template for a list of ExternalMessages
type ExternalMessage_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.ExternalMessage
}

//externalmessage_Page provides the information for the template for an individual ExternalMessage
type ExternalMessage_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 14/12/2021 by matttownsend on silicon.local - START
		SYSId string
		MessageID string
		MessageFormat string
		MessageDeliveredTo string
		MessageBody string
		MessageFilename string
		MessageLife string
		MessageDate string
		MessageTime string
		MessageTimeoutAction string
		MessageACKNAK string
		ResponseID string
		ResponseFilename string
		ResponseBody string
		ResponseDate string
		ResponseTime string
		ResponseAdditionalInfo string
		SYSCreated string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdated string
		SYSUpdatedBy string
		SYSUpdatedHost string
		MessageTimeout string
		MessageEmitted string
		ResponseRecieved string
		MessageClass string
		AppID string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 14/12/2021 by matttownsend on silicon.local - END
}

const (
	ExternalMessage_Redirect = dm.ExternalMessage_PathList
)

//ExternalMessage_Publish annouces the endpoints available for this object
func ExternalMessage_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.ExternalMessage_Path, ExternalMessage_Handler)
	mux.HandleFunc(dm.ExternalMessage_PathList, ExternalMessage_HandlerList)
	mux.HandleFunc(dm.ExternalMessage_PathView, ExternalMessage_HandlerView)
	mux.HandleFunc(dm.ExternalMessage_PathEdit, ExternalMessage_HandlerEdit)
	//Cannot Create via GUI
	mux.HandleFunc(dm.ExternalMessage_PathSave, ExternalMessage_HandlerSave)
	mux.HandleFunc(dm.ExternalMessage_PathDelete, ExternalMessage_HandlerDelete)
	logs.Publish("Application", dm.ExternalMessage_Title)
    core.Catalog_Add(dm.ExternalMessage_Title, dm.ExternalMessage_Path, "", dm.ExternalMessage_QueryString, "APP")
}

//ExternalMessage_HandlerList is the handler for the list page
func ExternalMessage_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.ExternalMessage
	noItems, returnList, _ := dao.ExternalMessage_GetList()

	pageDetail := ExternalMessage_PageList{
		Title:            CardTitle(dm.ExternalMessage_Title, core.Action_List),
		PageTitle:        PageTitle(dm.ExternalMessage_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.ExternalMessage_TemplateList, w, r, pageDetail)

}

//ExternalMessage_HandlerView is the handler used to View a page
func ExternalMessage_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.ExternalMessage_QueryString)
	_, rD, _ := dao.ExternalMessage_GetByID(searchID)

	pageDetail := ExternalMessage_Page{
		Title:       CardTitle(dm.ExternalMessage_Title, core.Action_View),
		PageTitle:   PageTitle(dm.ExternalMessage_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 14/12/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.MessageID = rD.MessageID
pageDetail.MessageFormat = rD.MessageFormat
pageDetail.MessageDeliveredTo = rD.MessageDeliveredTo
pageDetail.MessageBody = rD.MessageBody
pageDetail.MessageFilename = rD.MessageFilename
pageDetail.MessageLife = rD.MessageLife
pageDetail.MessageDate = rD.MessageDate
pageDetail.MessageTime = rD.MessageTime
pageDetail.MessageTimeoutAction = rD.MessageTimeoutAction
pageDetail.MessageACKNAK = rD.MessageACKNAK
pageDetail.ResponseID = rD.ResponseID
pageDetail.ResponseFilename = rD.ResponseFilename
pageDetail.ResponseBody = rD.ResponseBody
pageDetail.ResponseDate = rD.ResponseDate
pageDetail.ResponseTime = rD.ResponseTime
pageDetail.ResponseAdditionalInfo = rD.ResponseAdditionalInfo
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
pageDetail.MessageTimeout = rD.MessageTimeout
pageDetail.MessageEmitted = rD.MessageEmitted
pageDetail.ResponseRecieved = rD.ResponseRecieved
pageDetail.MessageClass = rD.MessageClass
pageDetail.AppID = rD.AppID


// Automatically generated 14/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 14/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 14/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.ExternalMessage_TemplateView, w, r, pageDetail)

}

//ExternalMessage_HandlerEdit is the handler used generate the Edit page
func ExternalMessage_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.ExternalMessage_QueryString)
	_, rD, _ := dao.ExternalMessage_GetByID(searchID)
	
	pageDetail := ExternalMessage_Page{
		Title:       CardTitle(dm.ExternalMessage_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.ExternalMessage_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 14/12/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.MessageID = rD.MessageID
pageDetail.MessageFormat = rD.MessageFormat
pageDetail.MessageDeliveredTo = rD.MessageDeliveredTo
pageDetail.MessageBody = rD.MessageBody
pageDetail.MessageFilename = rD.MessageFilename
pageDetail.MessageLife = rD.MessageLife
pageDetail.MessageDate = rD.MessageDate
pageDetail.MessageTime = rD.MessageTime
pageDetail.MessageTimeoutAction = rD.MessageTimeoutAction
pageDetail.MessageACKNAK = rD.MessageACKNAK
pageDetail.ResponseID = rD.ResponseID
pageDetail.ResponseFilename = rD.ResponseFilename
pageDetail.ResponseBody = rD.ResponseBody
pageDetail.ResponseDate = rD.ResponseDate
pageDetail.ResponseTime = rD.ResponseTime
pageDetail.ResponseAdditionalInfo = rD.ResponseAdditionalInfo
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
pageDetail.MessageTimeout = rD.MessageTimeout
pageDetail.MessageEmitted = rD.MessageEmitted
pageDetail.ResponseRecieved = rD.ResponseRecieved
pageDetail.MessageClass = rD.MessageClass
pageDetail.AppID = rD.AppID


// Automatically generated 14/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 14/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 14/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.ExternalMessage_TemplateEdit, w, r, pageDetail)


}

//ExternalMessage_HandlerSave is the handler used process the saving of an ExternalMessage
func ExternalMessage_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("MessageID"))

	var item dm.ExternalMessage
	// Automatically generated 14/12/2021 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.ExternalMessage_SYSId)
		item.MessageID = r.FormValue(dm.ExternalMessage_MessageID)
		item.MessageFormat = r.FormValue(dm.ExternalMessage_MessageFormat)
		item.MessageDeliveredTo = r.FormValue(dm.ExternalMessage_MessageDeliveredTo)
		item.MessageBody = r.FormValue(dm.ExternalMessage_MessageBody)
		item.MessageFilename = r.FormValue(dm.ExternalMessage_MessageFilename)
		item.MessageLife = r.FormValue(dm.ExternalMessage_MessageLife)
		item.MessageDate = r.FormValue(dm.ExternalMessage_MessageDate)
		item.MessageTime = r.FormValue(dm.ExternalMessage_MessageTime)
		item.MessageTimeoutAction = r.FormValue(dm.ExternalMessage_MessageTimeoutAction)
		item.MessageACKNAK = r.FormValue(dm.ExternalMessage_MessageACKNAK)
		item.ResponseID = r.FormValue(dm.ExternalMessage_ResponseID)
		item.ResponseFilename = r.FormValue(dm.ExternalMessage_ResponseFilename)
		item.ResponseBody = r.FormValue(dm.ExternalMessage_ResponseBody)
		item.ResponseDate = r.FormValue(dm.ExternalMessage_ResponseDate)
		item.ResponseTime = r.FormValue(dm.ExternalMessage_ResponseTime)
		item.ResponseAdditionalInfo = r.FormValue(dm.ExternalMessage_ResponseAdditionalInfo)
		item.SYSCreated = r.FormValue(dm.ExternalMessage_SYSCreated)
		item.SYSCreatedBy = r.FormValue(dm.ExternalMessage_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.ExternalMessage_SYSCreatedHost)
		item.SYSUpdated = r.FormValue(dm.ExternalMessage_SYSUpdated)
		item.SYSUpdatedBy = r.FormValue(dm.ExternalMessage_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.ExternalMessage_SYSUpdatedHost)
		item.MessageTimeout = r.FormValue(dm.ExternalMessage_MessageTimeout)
		item.MessageEmitted = r.FormValue(dm.ExternalMessage_MessageEmitted)
		item.ResponseRecieved = r.FormValue(dm.ExternalMessage_ResponseRecieved)
		item.MessageClass = r.FormValue(dm.ExternalMessage_MessageClass)
		item.AppID = r.FormValue(dm.ExternalMessage_AppID)
	

	// Automatically generated 14/12/2021 by matttownsend on silicon.local - END

	dao.ExternalMessage_Store(item,r)	

	http.Redirect(w, r, ExternalMessage_Redirect, http.StatusFound)
}

//ExternalMessage_HandlerNew is the handler used process the creation of an ExternalMessage
func ExternalMessage_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := ExternalMessage_Page{
		Title:       CardTitle(dm.ExternalMessage_Title, core.Action_New),
		PageTitle:   PageTitle(dm.ExternalMessage_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 14/12/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.MessageID = ""
pageDetail.MessageFormat = ""
pageDetail.MessageDeliveredTo = ""
pageDetail.MessageBody = ""
pageDetail.MessageFilename = ""
pageDetail.MessageLife = ""
pageDetail.MessageDate = ""
pageDetail.MessageTime = ""
pageDetail.MessageTimeoutAction = ""
pageDetail.MessageACKNAK = ""
pageDetail.ResponseID = ""
pageDetail.ResponseFilename = ""
pageDetail.ResponseBody = ""
pageDetail.ResponseDate = ""
pageDetail.ResponseTime = ""
pageDetail.ResponseAdditionalInfo = ""
pageDetail.SYSCreated = ""
pageDetail.SYSCreatedBy = ""
pageDetail.SYSCreatedHost = ""
pageDetail.SYSUpdated = ""
pageDetail.SYSUpdatedBy = ""
pageDetail.SYSUpdatedHost = ""
pageDetail.MessageTimeout = ""
pageDetail.MessageEmitted = ""
pageDetail.ResponseRecieved = ""
pageDetail.MessageClass = ""
pageDetail.AppID = ""


// Automatically generated 14/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 14/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.ExternalMessage_TemplateNew, w, r, pageDetail)

}

//ExternalMessage_HandlerDelete is the handler used process the deletion of an ExternalMessage
func ExternalMessage_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.ExternalMessage_QueryString)

	dao.ExternalMessage_Delete(searchID)	

	http.Redirect(w, r, ExternalMessage_Redirect, http.StatusFound)
}