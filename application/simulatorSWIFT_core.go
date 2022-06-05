package application
// ----------------------------------------------------------------
// Automatically generated  "/application/simulatorswift.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : SimulatorSWIFT (simulatorswift)
// Endpoint 	        : SimulatorSWIFT (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 05/06/2022 at 12:56:56
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//simulatorswift_PageList provides the information for the template for a list of SimulatorSWIFTs
type SimulatorSWIFT_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.SimulatorSWIFT
}

//simulatorswift_Page provides the information for the template for an individual SimulatorSWIFT
type SimulatorSWIFT_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 05/06/2022 by matttownsend on silicon.local - START
		ID string
		FileName string
		MessageRaw string
		MessageFmt string
		Action string
	
	
	
	
	
	
	
	
	
	// Automatically generated 05/06/2022 by matttownsend on silicon.local - END
}

const (
	SimulatorSWIFT_Redirect = dm.SimulatorSWIFT_PathList
)

//SimulatorSWIFT_Publish annouces the endpoints available for this object
func SimulatorSWIFT_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.SimulatorSWIFT_Path, SimulatorSWIFT_Handler)
	mux.HandleFunc(dm.SimulatorSWIFT_PathList, SimulatorSWIFT_HandlerList)
	mux.HandleFunc(dm.SimulatorSWIFT_PathView, SimulatorSWIFT_HandlerView)
	mux.HandleFunc(dm.SimulatorSWIFT_PathEdit, SimulatorSWIFT_HandlerEdit)
	mux.HandleFunc(dm.SimulatorSWIFT_PathNew, SimulatorSWIFT_HandlerNew)
	mux.HandleFunc(dm.SimulatorSWIFT_PathSave, SimulatorSWIFT_HandlerSave)
	mux.HandleFunc(dm.SimulatorSWIFT_PathDelete, SimulatorSWIFT_HandlerDelete)
	logs.Publish("Siena", dm.SimulatorSWIFT_Title)
    core.Catalog_Add(dm.SimulatorSWIFT_Title, dm.SimulatorSWIFT_Path, "", dm.SimulatorSWIFT_QueryString, "SIENA")
}

//SimulatorSWIFT_HandlerList is the handler for the list page
func SimulatorSWIFT_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.SimulatorSWIFT
	noItems, returnList, _ := dao.SimulatorSWIFT_GetList()

	pageDetail := SimulatorSWIFT_PageList{
		Title:            CardTitle(dm.SimulatorSWIFT_Title, core.Action_List),
		PageTitle:        PageTitle(dm.SimulatorSWIFT_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.SimulatorSWIFT_TemplateList, w, r, pageDetail)

}

//SimulatorSWIFT_HandlerView is the handler used to View a page
func SimulatorSWIFT_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.SimulatorSWIFT_QueryString)
	_, rD, _ := dao.SimulatorSWIFT_GetByID(searchID)

	pageDetail := SimulatorSWIFT_Page{
		Title:       CardTitle(dm.SimulatorSWIFT_Title, core.Action_View),
		PageTitle:   PageTitle(dm.SimulatorSWIFT_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 05/06/2022 by matttownsend on silicon.local - START
pageDetail.ID = rD.ID
pageDetail.FileName = rD.FileName
pageDetail.MessageRaw = rD.MessageRaw
pageDetail.MessageFmt = rD.MessageFmt
pageDetail.Action = rD.Action


// Automatically generated 05/06/2022 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 05/06/2022 by matttownsend on silicon.local - END
		//


	// Automatically generated 05/06/2022 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.SimulatorSWIFT_TemplateView, w, r, pageDetail)

}

//SimulatorSWIFT_HandlerEdit is the handler used generate the Edit page
func SimulatorSWIFT_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.SimulatorSWIFT_QueryString)
	_, rD, _ := dao.SimulatorSWIFT_GetByID(searchID)
	
	pageDetail := SimulatorSWIFT_Page{
		Title:       CardTitle(dm.SimulatorSWIFT_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.SimulatorSWIFT_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 05/06/2022 by matttownsend on silicon.local - START
pageDetail.ID = rD.ID
pageDetail.FileName = rD.FileName
pageDetail.MessageRaw = rD.MessageRaw
pageDetail.MessageFmt = rD.MessageFmt
pageDetail.Action = rD.Action


// Automatically generated 05/06/2022 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 05/06/2022 by matttownsend on silicon.local - END

	// Automatically generated 05/06/2022 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.SimulatorSWIFT_TemplateEdit, w, r, pageDetail)


}

//SimulatorSWIFT_HandlerSave is the handler used process the saving of an SimulatorSWIFT
func SimulatorSWIFT_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ID"))

	var item dm.SimulatorSWIFT
	// Automatically generated 05/06/2022 by matttownsend on silicon.local - START
		item.ID = r.FormValue(dm.SimulatorSWIFT_ID)
		item.FileName = r.FormValue(dm.SimulatorSWIFT_FileName)
		item.MessageRaw = r.FormValue(dm.SimulatorSWIFT_MessageRaw)
		item.MessageFmt = r.FormValue(dm.SimulatorSWIFT_MessageFmt)
		item.Action = r.FormValue(dm.SimulatorSWIFT_Action)
		
	

	// Automatically generated 05/06/2022 by matttownsend on silicon.local - END

	dao.SimulatorSWIFT_Store(item,r)	

	http.Redirect(w, r, SimulatorSWIFT_Redirect, http.StatusFound)
}

//SimulatorSWIFT_HandlerNew is the handler used process the creation of an SimulatorSWIFT
func SimulatorSWIFT_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := SimulatorSWIFT_Page{
		Title:       CardTitle(dm.SimulatorSWIFT_Title, core.Action_New),
		PageTitle:   PageTitle(dm.SimulatorSWIFT_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 05/06/2022 by matttownsend on silicon.local - START
pageDetail.ID = ""
pageDetail.FileName = ""
pageDetail.MessageRaw = ""
pageDetail.MessageFmt = ""
pageDetail.Action = ""


// Automatically generated 05/06/2022 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 05/06/2022 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.SimulatorSWIFT_TemplateNew, w, r, pageDetail)

}

//SimulatorSWIFT_HandlerDelete is the handler used process the deletion of an SimulatorSWIFT
func SimulatorSWIFT_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.SimulatorSWIFT_QueryString)

	dao.SimulatorSWIFT_Delete(searchID)	

	http.Redirect(w, r, SimulatorSWIFT_Redirect, http.StatusFound)
}