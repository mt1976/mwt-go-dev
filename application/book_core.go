package application
// ----------------------------------------------------------------
// Automatically generated  "/application/book.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Book (book)
// Endpoint 	        : Book (Book)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:18
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//book_PageList provides the information for the template for a list of Books
type Book_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Book
}
//Book_Redirect provides a page to return to aftern an action
const (
	
	Book_Redirect = dm.Book_PathList
	
)

//book_Page provides the information for the template for an individual Book
type Book_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	BookName         string
	BookName_props     dm.FieldProperties
	FullName         string
	FullName_props     dm.FieldProperties
	PLManage         string
	PLManage_props     dm.FieldProperties
	PLTransfer         string
	PLTransfer_props     dm.FieldProperties
	DerivePL         string
	DerivePL_props     dm.FieldProperties
	CostOfCarry         string
	CostOfCarry_props     dm.FieldProperties
	CostOfFunding         string
	CostOfFunding_props     dm.FieldProperties
	LotAllocationMethod         string
	LotAllocationMethod_props     dm.FieldProperties
	InternalId         string
	InternalId_props     dm.FieldProperties
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//Book_Publish annouces the endpoints available for this object
func Book_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Book_Path, Book_Handler)
	mux.HandleFunc(dm.Book_PathList, Book_HandlerList)
	mux.HandleFunc(dm.Book_PathView, Book_HandlerView)
	mux.HandleFunc(dm.Book_PathEdit, Book_HandlerEdit)
	//Cannot Create via GUI
	mux.HandleFunc(dm.Book_PathSave, Book_HandlerSave)
	mux.HandleFunc(dm.Book_PathDelete, Book_HandlerDelete)
	logs.Publish("Application", dm.Book_Title)
    core.Catalog_Add(dm.Book_Title, dm.Book_Path, "", dm.Book_QueryString, "Application")
}


//Book_HandlerList is the handler for the list page
func Book_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Book
	noItems, returnList, _ := dao.Book_GetList()

	pageDetail := Book_PageList{
		Title:            CardTitle(dm.Book_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Book_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Book_TemplateList, w, r, pageDetail)

}


//Book_HandlerView is the handler used to View a page
func Book_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Book_QueryString)
	_, rD, _ := dao.Book_GetByID(searchID)

	pageDetail := Book_Page{
		Title:       CardTitle(dm.Book_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Book_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = book_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Book_TemplateView, w, r, pageDetail)

}


//Book_HandlerEdit is the handler used generate the Edit page
func Book_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Book_QueryString)
	_, rD, _ := dao.Book_GetByID(searchID)
	
	pageDetail := Book_Page{
		Title:       CardTitle(dm.Book_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Book_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = book_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Book_TemplateEdit, w, r, pageDetail)
}


//Book_HandlerSave is the handler used process the saving of an Book
func Book_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("BookName"))

	var item dm.Book
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.BookName = r.FormValue(dm.Book_BookName_scrn)
		item.FullName = r.FormValue(dm.Book_FullName_scrn)
		item.PLManage = r.FormValue(dm.Book_PLManage_scrn)
		item.PLTransfer = r.FormValue(dm.Book_PLTransfer_scrn)
		item.DerivePL = r.FormValue(dm.Book_DerivePL_scrn)
		item.CostOfCarry = r.FormValue(dm.Book_CostOfCarry_scrn)
		item.CostOfFunding = r.FormValue(dm.Book_CostOfFunding_scrn)
		item.LotAllocationMethod = r.FormValue(dm.Book_LotAllocationMethod_scrn)
		item.InternalId = r.FormValue(dm.Book_InternalId_scrn)
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.Book_Store(item,r)	
	http.Redirect(w, r, Book_Redirect, http.StatusFound)
}



//Book_HandlerDelete is the handler used process the deletion of an Book
func Book_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Book_QueryString)

	dao.Book_Delete(searchID)	

	http.Redirect(w, r, Book_Redirect, http.StatusFound)
}


// Builds/Popuplates the Book Page 
func book_PopulatePage(rD dm.Book, pageDetail Book_Page) Book_Page {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.BookName = rD.BookName
	pageDetail.FullName = rD.FullName
	pageDetail.PLManage = rD.PLManage
	pageDetail.PLTransfer = rD.PLTransfer
	pageDetail.DerivePL = rD.DerivePL
	pageDetail.CostOfCarry = rD.CostOfCarry
	pageDetail.CostOfFunding = rD.CostOfFunding
	pageDetail.LotAllocationMethod = rD.LotAllocationMethod
	pageDetail.InternalId = rD.InternalId
	
	
	//
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.BookName_props = rD.BookName_props
	pageDetail.FullName_props = rD.FullName_props
	pageDetail.PLManage_props = rD.PLManage_props
	pageDetail.PLTransfer_props = rD.PLTransfer_props
	pageDetail.DerivePL_props = rD.DerivePL_props
	pageDetail.CostOfCarry_props = rD.CostOfCarry_props
	pageDetail.CostOfFunding_props = rD.CostOfFunding_props
	pageDetail.LotAllocationMethod_props = rD.LotAllocationMethod_props
	pageDetail.InternalId_props = rD.InternalId_props
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	