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
// Date & Time		    : 26/06/2022 at 18:48:31
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
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	Code         string
	Code_props     dm.FieldProperties
	Description1         string
	Description1_props     dm.FieldProperties
	Description2         string
	Description2_props     dm.FieldProperties
	IsDefault         string
	IsDefault_props     dm.FieldProperties
	InternalId         string
	InternalId_props     dm.FieldProperties
	InternalDeleted         string
	InternalDeleted_props     dm.FieldProperties
	UpdatedTransactionId         string
	UpdatedTransactionId_props     dm.FieldProperties
	UpdatedUserId         string
	UpdatedUserId_props     dm.FieldProperties
	UpdatedDateTime         string
	UpdatedDateTime_props     dm.FieldProperties
	DeletedTransactionId         string
	DeletedTransactionId_props     dm.FieldProperties
	DeletedUserId         string
	DeletedUserId_props     dm.FieldProperties
	ChangeType         string
	ChangeType_props     dm.FieldProperties
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.Code = r.FormValue(dm.Portfolio_Code_scrn)
		item.Description1 = r.FormValue(dm.Portfolio_Description1_scrn)
		item.Description2 = r.FormValue(dm.Portfolio_Description2_scrn)
		item.IsDefault = r.FormValue(dm.Portfolio_IsDefault_scrn)
		item.InternalId = r.FormValue(dm.Portfolio_InternalId_scrn)
		item.InternalDeleted = r.FormValue(dm.Portfolio_InternalDeleted_scrn)
		item.UpdatedTransactionId = r.FormValue(dm.Portfolio_UpdatedTransactionId_scrn)
		item.UpdatedUserId = r.FormValue(dm.Portfolio_UpdatedUserId_scrn)
		item.UpdatedDateTime = r.FormValue(dm.Portfolio_UpdatedDateTime_scrn)
		item.DeletedTransactionId = r.FormValue(dm.Portfolio_DeletedTransactionId_scrn)
		item.DeletedUserId = r.FormValue(dm.Portfolio_DeletedUserId_scrn)
		item.ChangeType = r.FormValue(dm.Portfolio_ChangeType_scrn)
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.Portfolio_New()

	pageDetail := Portfolio_Page{
		Title:       CardTitle(dm.Portfolio_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Portfolio_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = portfolio_PopulatePage(rD , pageDetail) 

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
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.Code_props = rD.Code_props
	pageDetail.Description1_props = rD.Description1_props
	pageDetail.Description2_props = rD.Description2_props
	pageDetail.IsDefault_props = rD.IsDefault_props
	pageDetail.InternalId_props = rD.InternalId_props
	pageDetail.InternalDeleted_props = rD.InternalDeleted_props
	pageDetail.UpdatedTransactionId_props = rD.UpdatedTransactionId_props
	pageDetail.UpdatedUserId_props = rD.UpdatedUserId_props
	pageDetail.UpdatedDateTime_props = rD.UpdatedDateTime_props
	pageDetail.DeletedTransactionId_props = rD.DeletedTransactionId_props
	pageDetail.DeletedUserId_props = rD.DeletedUserId_props
	pageDetail.ChangeType_props = rD.ChangeType_props
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	