package application
// ----------------------------------------------------------------
// Automatically generated  "/application/product.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Product (product)
// Endpoint 	        : Product (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:18
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//product_PageList provides the information for the template for a list of Products
type Product_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Product
}

//product_Page provides the information for the template for an individual Product
type Product_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		Code string
		Name string
		Factor string
		MaxTermPrecedence string
		InternalId string
		InternalDeleted string
		UpdatedTransactionId string
		UpdatedUserId string
		UpdatedDateTime string
		DeletedTransactionId string
		DeletedUserId string
		ChangeType string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	Product_Redirect = dm.Product_PathList
)

//Product_Publish annouces the endpoints available for this object
func Product_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Product_Path, Product_Handler)
	mux.HandleFunc(dm.Product_PathList, Product_HandlerList)
	mux.HandleFunc(dm.Product_PathView, Product_HandlerView)
	mux.HandleFunc(dm.Product_PathEdit, Product_HandlerEdit)
	mux.HandleFunc(dm.Product_PathNew, Product_HandlerNew)
	mux.HandleFunc(dm.Product_PathSave, Product_HandlerSave)
	mux.HandleFunc(dm.Product_PathDelete, Product_HandlerDelete)
	logs.Publish("Siena", dm.Product_Title)
    core.Catalog_Add(dm.Product_Title, dm.Product_Path, "", dm.Product_QueryString, "APP")
}

//Product_HandlerList is the handler for the list page
func Product_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Product
	noItems, returnList, _ := dao.Product_GetList()

	pageDetail := Product_PageList{
		Title:            CardTitle(dm.Product_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Product_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Product_TemplateList, w, r, pageDetail)

}

//Product_HandlerView is the handler used to View a page
func Product_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Product_QueryString)
	_, rD, _ := dao.Product_GetByID(searchID)

	pageDetail := Product_Page{
		Title:       CardTitle(dm.Product_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Product_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
pageDetail.Factor = rD.Factor
pageDetail.MaxTermPrecedence = rD.MaxTermPrecedence
pageDetail.InternalId = rD.InternalId
pageDetail.InternalDeleted = rD.InternalDeleted
pageDetail.UpdatedTransactionId = rD.UpdatedTransactionId
pageDetail.UpdatedUserId = rD.UpdatedUserId
pageDetail.UpdatedDateTime = rD.UpdatedDateTime
pageDetail.DeletedTransactionId = rD.DeletedTransactionId
pageDetail.DeletedUserId = rD.DeletedUserId
pageDetail.ChangeType = rD.ChangeType


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Product_TemplateView, w, r, pageDetail)

}

//Product_HandlerEdit is the handler used generate the Edit page
func Product_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Product_QueryString)
	_, rD, _ := dao.Product_GetByID(searchID)
	
	pageDetail := Product_Page{
		Title:       CardTitle(dm.Product_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Product_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
pageDetail.Factor = rD.Factor
pageDetail.MaxTermPrecedence = rD.MaxTermPrecedence
pageDetail.InternalId = rD.InternalId
pageDetail.InternalDeleted = rD.InternalDeleted
pageDetail.UpdatedTransactionId = rD.UpdatedTransactionId
pageDetail.UpdatedUserId = rD.UpdatedUserId
pageDetail.UpdatedDateTime = rD.UpdatedDateTime
pageDetail.DeletedTransactionId = rD.DeletedTransactionId
pageDetail.DeletedUserId = rD.DeletedUserId
pageDetail.ChangeType = rD.ChangeType


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Product_TemplateEdit, w, r, pageDetail)


}

//Product_HandlerSave is the handler used process the saving of an Product
func Product_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Code"))

	var item dm.Product
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.Code = r.FormValue(dm.Product_Code)
		item.Name = r.FormValue(dm.Product_Name)
		item.Factor = r.FormValue(dm.Product_Factor)
		item.MaxTermPrecedence = r.FormValue(dm.Product_MaxTermPrecedence)
		item.InternalId = r.FormValue(dm.Product_InternalId)
		item.InternalDeleted = r.FormValue(dm.Product_InternalDeleted)
		item.UpdatedTransactionId = r.FormValue(dm.Product_UpdatedTransactionId)
		item.UpdatedUserId = r.FormValue(dm.Product_UpdatedUserId)
		item.UpdatedDateTime = r.FormValue(dm.Product_UpdatedDateTime)
		item.DeletedTransactionId = r.FormValue(dm.Product_DeletedTransactionId)
		item.DeletedUserId = r.FormValue(dm.Product_DeletedUserId)
		item.ChangeType = r.FormValue(dm.Product_ChangeType)
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.Product_Store(item,r)	

	http.Redirect(w, r, Product_Redirect, http.StatusFound)
}

//Product_HandlerNew is the handler used process the creation of an Product
func Product_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Product_Page{
		Title:       CardTitle(dm.Product_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Product_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.Code = ""
pageDetail.Name = ""
pageDetail.Factor = ""
pageDetail.MaxTermPrecedence = ""
pageDetail.InternalId = ""
pageDetail.InternalDeleted = ""
pageDetail.UpdatedTransactionId = ""
pageDetail.UpdatedUserId = ""
pageDetail.UpdatedDateTime = ""
pageDetail.DeletedTransactionId = ""
pageDetail.DeletedUserId = ""
pageDetail.ChangeType = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Product_TemplateNew, w, r, pageDetail)

}

//Product_HandlerDelete is the handler used process the deletion of an Product
func Product_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Product_QueryString)

	dao.Product_Delete(searchID)	

	http.Redirect(w, r, Product_Redirect, http.StatusFound)
}