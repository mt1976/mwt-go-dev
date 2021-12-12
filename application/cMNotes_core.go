package application
// ----------------------------------------------------------------
// Automatically generated  "/application/cmnotes.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CMNotes (cmnotes)
// Endpoint 	        : CMNotes (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:08
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//cmnotes_PageList provides the information for the template for a list of CMNotess
type CMNotes_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.CMNotes
}

//cmnotes_Page provides the information for the template for an individual CMNotes
type CMNotes_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		NoteId string
		StreamId string
		Summary string
		Details string
		RecordState string
		CreatedBy string
		CreatedDateTime string
	
	
	
	
	
	
	
	
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	CMNotes_Redirect = dm.CMNotes_PathList
)

//CMNotes_Publish annouces the endpoints available for this object
func CMNotes_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.CMNotes_PathList, CMNotes_HandlerList)
	mux.HandleFunc(dm.CMNotes_PathView, CMNotes_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Siena", dm.CMNotes_Title)
    //No API
}

//CMNotes_HandlerList is the handler for the list page
func CMNotes_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.CMNotes
	noItems, returnList, _ := dao.CMNotes_GetList()

	pageDetail := CMNotes_PageList{
		Title:            CardTitle(dm.CMNotes_Title, core.Action_List),
		PageTitle:        PageTitle(dm.CMNotes_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.CMNotes_TemplateList, w, r, pageDetail)

}

//CMNotes_HandlerView is the handler used to View a page
func CMNotes_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CMNotes_QueryString)
	_, rD, _ := dao.CMNotes_GetByID(searchID)

	pageDetail := CMNotes_Page{
		Title:       CardTitle(dm.CMNotes_Title, core.Action_View),
		PageTitle:   PageTitle(dm.CMNotes_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.NoteId = rD.NoteId
pageDetail.StreamId = rD.StreamId
pageDetail.Summary = rD.Summary
pageDetail.Details = rD.Details
pageDetail.RecordState = rD.RecordState
pageDetail.CreatedBy = rD.CreatedBy
pageDetail.CreatedDateTime = rD.CreatedDateTime


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.CMNotes_TemplateView, w, r, pageDetail)

}

//CMNotes_HandlerEdit is the handler used generate the Edit page
func CMNotes_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CMNotes_QueryString)
	_, rD, _ := dao.CMNotes_GetByID(searchID)
	
	pageDetail := CMNotes_Page{
		Title:       CardTitle(dm.CMNotes_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.CMNotes_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.NoteId = rD.NoteId
pageDetail.StreamId = rD.StreamId
pageDetail.Summary = rD.Summary
pageDetail.Details = rD.Details
pageDetail.RecordState = rD.RecordState
pageDetail.CreatedBy = rD.CreatedBy
pageDetail.CreatedDateTime = rD.CreatedDateTime


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.CMNotes_TemplateEdit, w, r, pageDetail)


}

//CMNotes_HandlerSave is the handler used process the saving of an CMNotes
func CMNotes_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("NoteId"))

	var item dm.CMNotes
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.NoteId = r.FormValue(dm.CMNotes_NoteId)
		item.StreamId = r.FormValue(dm.CMNotes_StreamId)
		item.Summary = r.FormValue(dm.CMNotes_Summary)
		item.Details = r.FormValue(dm.CMNotes_Details)
		item.RecordState = r.FormValue(dm.CMNotes_RecordState)
		item.CreatedBy = r.FormValue(dm.CMNotes_CreatedBy)
		item.CreatedDateTime = r.FormValue(dm.CMNotes_CreatedDateTime)
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.CMNotes_Store(item,r)	

	http.Redirect(w, r, CMNotes_Redirect, http.StatusFound)
}

//CMNotes_HandlerNew is the handler used process the creation of an CMNotes
func CMNotes_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := CMNotes_Page{
		Title:       CardTitle(dm.CMNotes_Title, core.Action_New),
		PageTitle:   PageTitle(dm.CMNotes_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.NoteId = ""
pageDetail.StreamId = ""
pageDetail.Summary = ""
pageDetail.Details = ""
pageDetail.RecordState = ""
pageDetail.CreatedBy = ""
pageDetail.CreatedDateTime = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.CMNotes_TemplateNew, w, r, pageDetail)

}

//CMNotes_HandlerDelete is the handler used process the deletion of an CMNotes
func CMNotes_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.CMNotes_QueryString)

	dao.CMNotes_Delete(searchID)	

	http.Redirect(w, r, CMNotes_Redirect, http.StatusFound)
}