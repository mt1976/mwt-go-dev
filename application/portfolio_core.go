package application
// ----------------------------------------------------------------
// Automatically generated  "/application/portfolio.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Portfolio (portfolio)
// Endpoint 	        : Portfolio (Code)
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

//portfolio_PageList provides the information for the template for a list of Portfolios
type Portfolio_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Portfolio
}
//Portfolio_Redirect provides a page to return to aftern an action
const (
	Portfolio_Redirect = dm.Portfolio_PathList
)

//portfolio_Page provides the information for the template for an individual Portfolio
type Portfolio_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	Code         string
	Description1         string
	Description2         string
	IsDefault         string
	InternalId         string
	InternalDeleted         string
	UpdatedTransactionId         string
	UpdatedUserId         string
	UpdatedDateTime         string
	DeletedTransactionId         string
	DeletedUserId         string
	ChangeType         string
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//Portfolio_Publish annouces the endpoints available for this object
func Portfolio_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Portfolio_Path, Portfolio_Handler)
	mux.HandleFunc(dm.Portfolio_PathList, Portfolio_HandlerList)
	mux.HandleFunc(dm.Portfolio_PathView, Portfolio_HandlerView)
	mux.HandleFunc(dm.Portfolio_PathEdit, Portfolio_HandlerEdit)
	mux.HandleFunc(dm.Portfolio_PathNew, Portfolio_HandlerNew)
	mux.HandleFunc(dm.Portfolio_PathSave, Portfolio_HandlerSave)
	mux.HandleFunc(dm.Portfolio_PathDelete, Portfolio_HandlerDelete)
	logs.Publish("Application", dm.Portfolio_Title)
    core.Catalog_Add(dm.Portfolio_Title, dm.Portfolio_Path, "", dm.Portfolio_QueryString, "Application")
}


//Portfolio_HandlerList is the handler for the list page
func Portfolio_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Portfolio
	noItems, returnList, _ := dao.Portfolio_GetList()

	pageDetail := Portfolio_PageList{
		Title:            CardTitle(dm.Portfolio_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Portfolio_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Portfolio_TemplateList, w, r, pageDetail)

}


//Portfolio_HandlerView is the handler used to View a page
func Portfolio_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Portfolio_QueryString)
	_, rD, _ := dao.Portfolio_GetByID(searchID)

	pageDetail := Portfolio_Page{
		Title:       CardTitle(dm.Portfolio_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Portfolio_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = portfolio_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Portfolio_TemplateView, w, r, pageDetail)

}


//Portfolio_HandlerEdit is the handler used generate the Edit page
func Portfolio_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Portfolio_QueryString)
	_, rD, _ := dao.Portfolio_GetByID(searchID)
	
	pageDetail := Portfolio_Page{
		Title:       CardTitle(dm.Portfolio_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Portfolio_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = portfolio_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Portfolio_TemplateEdit, w, r, pageDetail)
}


//Portfolio_HandlerSave is the handler used process the saving of an Portfolio
func Portfolio_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Code"))

	var item dm.Portfolio
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.Code = r.FormValue(dm.Portfolio_Code)
		item.Description1 = r.FormValue(dm.Portfolio_Description1)
		item.Description2 = r.FormValue(dm.Portfolio_Description2)
		item.IsDefault = r.FormValue(dm.Portfolio_IsDefault)
		item.InternalId = r.FormValue(dm.Portfolio_InternalId)
		item.InternalDeleted = r.FormValue(dm.Portfolio_InternalDeleted)
		item.UpdatedTransactionId = r.FormValue(dm.Portfolio_UpdatedTransactionId)
		item.UpdatedUserId = r.FormValue(dm.Portfolio_UpdatedUserId)
		item.UpdatedDateTime = r.FormValue(dm.Portfolio_UpdatedDateTime)
		item.DeletedTransactionId = r.FormValue(dm.Portfolio_DeletedTransactionId)
		item.DeletedUserId = r.FormValue(dm.Portfolio_DeletedUserId)
		item.ChangeType = r.FormValue(dm.Portfolio_ChangeType)
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Portfolio_Store(item,r)	
	http.Redirect(w, r, Portfolio_Redirect, http.StatusFound)
}


//Portfolio_HandlerNew is the handler used process the creation of an Portfolio
func Portfolio_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Portfolio_Page{
		Title:       CardTitle(dm.Portfolio_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Portfolio_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = portfolio_PopulatePage(dm.Portfolio{} , pageDetail) 

	ExecuteTemplate(dm.Portfolio_TemplateNew, w, r, pageDetail)

}	


//Portfolio_HandlerDelete is the handler used process the deletion of an Portfolio
func Portfolio_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Portfolio_QueryString)

	dao.Portfolio_Delete(searchID)	

	http.Redirect(w, r, Portfolio_Redirect, http.StatusFound)
}


// Builds/Popuplates the Portfolio Page 
func portfolio_PopulatePage(rD dm.Portfolio, pageDetail Portfolio_Page) Portfolio_Page {
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.Code = rD.Code
	pageDetail.Description1 = rD.Description1
	pageDetail.Description2 = rD.Description2
	pageDetail.IsDefault = rD.IsDefault
	pageDetail.InternalId = rD.InternalId
	pageDetail.InternalDeleted = rD.InternalDeleted
	pageDetail.UpdatedTransactionId = rD.UpdatedTransactionId
	pageDetail.UpdatedUserId = rD.UpdatedUserId
	pageDetail.UpdatedDateTime = rD.UpdatedDateTime
	pageDetail.DeletedTransactionId = rD.DeletedTransactionId
	pageDetail.DeletedUserId = rD.DeletedUserId
	pageDetail.ChangeType = rD.ChangeType
	
	
	//
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	