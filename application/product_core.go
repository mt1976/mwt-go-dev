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

//product_PageList provides the information for the template for a list of Products
type Product_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Product
}
//Product_Redirect provides a page to return to aftern an action
const (
	
	Product_Redirect = dm.Product_PathList
	
)

//product_Page provides the information for the template for an individual Product
type Product_Page struct {
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
	Name         string
	Name_props     dm.FieldProperties
	Factor         string
	Factor_props     dm.FieldProperties
	MaxTermPrecedence         string
	MaxTermPrecedence_props     dm.FieldProperties
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



//Product_Publish annouces the endpoints available for this object
func Product_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Product_Path, Product_Handler)
	mux.HandleFunc(dm.Product_PathList, Product_HandlerList)
	mux.HandleFunc(dm.Product_PathView, Product_HandlerView)
	mux.HandleFunc(dm.Product_PathEdit, Product_HandlerEdit)
	mux.HandleFunc(dm.Product_PathNew, Product_HandlerNew)
	mux.HandleFunc(dm.Product_PathSave, Product_HandlerSave)
	mux.HandleFunc(dm.Product_PathDelete, Product_HandlerDelete)
	logs.Publish("Application", dm.Product_Title)
    core.Catalog_Add(dm.Product_Title, dm.Product_Path, "", dm.Product_QueryString, "Application")
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

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = product_PopulatePage(rD , pageDetail) 

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

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = product_PopulatePage(rD , pageDetail) 

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
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.Code = r.FormValue(dm.Product_Code_scrn)
		item.Name = r.FormValue(dm.Product_Name_scrn)
		item.Factor = r.FormValue(dm.Product_Factor_scrn)
		item.MaxTermPrecedence = r.FormValue(dm.Product_MaxTermPrecedence_scrn)
		item.InternalId = r.FormValue(dm.Product_InternalId_scrn)
		item.InternalDeleted = r.FormValue(dm.Product_InternalDeleted_scrn)
		item.UpdatedTransactionId = r.FormValue(dm.Product_UpdatedTransactionId_scrn)
		item.UpdatedUserId = r.FormValue(dm.Product_UpdatedUserId_scrn)
		item.UpdatedDateTime = r.FormValue(dm.Product_UpdatedDateTime_scrn)
		item.DeletedTransactionId = r.FormValue(dm.Product_DeletedTransactionId_scrn)
		item.DeletedUserId = r.FormValue(dm.Product_DeletedUserId_scrn)
		item.ChangeType = r.FormValue(dm.Product_ChangeType_scrn)
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
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

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.Product_New()

	pageDetail := Product_Page{
		Title:       CardTitle(dm.Product_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Product_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = product_PopulatePage(rD , pageDetail) 

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


// Builds/Popuplates the Product Page 
func product_PopulatePage(rD dm.Product, pageDetail Product_Page) Product_Page {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
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
	
	
	//
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.Code_props = rD.Code_props
	pageDetail.Name_props = rD.Name_props
	pageDetail.Factor_props = rD.Factor_props
	pageDetail.MaxTermPrecedence_props = rD.MaxTermPrecedence_props
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