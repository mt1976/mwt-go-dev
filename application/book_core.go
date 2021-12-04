package application
// ----------------------------------------------------------------
// Automatically generated  "/application/book.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Book (book)
// Endpoint 	        : Book (Book)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 04/12/2021 at 17:36:42
// Who & Where		    : matttownsend on silicon.local
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
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Book
}

//book_Page provides the information for the template for an individual Book
type Book_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	// Automatically generated 04/12/2021 by matttownsend on silicon.local - START
		BookName string
		FullName string
		PLManage string
		PLTransfer string
		DerivePL string
		CostOfCarry string
		CostOfFunding string
		LotAllocationMethod string
		InternalId string
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 04/12/2021 by matttownsend on silicon.local - END
}

const (
	Book_Redirect = dm.Book_PathList
)

//Book_Publish annouces the endpoints available for this object
func Book_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Book_PathList, Book_HandlerList)
	mux.HandleFunc(dm.Book_PathView, Book_HandlerView)
	mux.HandleFunc(dm.Book_PathEdit, Book_HandlerEdit)
	
	mux.HandleFunc(dm.Book_PathSave, Book_HandlerSave)
	mux.HandleFunc(dm.Book_PathDelete, Book_HandlerDelete)
	logs.Publish("Siena", dm.Book_Title)
	
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
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.Book_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}

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
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Book_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 04/12/2021 by matttownsend on silicon.local - START
pageDetail.BookName = rD.BookName
pageDetail.FullName = rD.FullName
pageDetail.PLManage = rD.PLManage
pageDetail.PLTransfer = rD.PLTransfer
pageDetail.DerivePL = rD.DerivePL
pageDetail.CostOfCarry = rD.CostOfCarry
pageDetail.CostOfFunding = rD.CostOfFunding
pageDetail.LotAllocationMethod = rD.LotAllocationMethod
pageDetail.InternalId = rD.InternalId
// Automatically generated 04/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 04/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 04/12/2021 by matttownsend on silicon.local - END


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
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Book_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 04/12/2021 by matttownsend on silicon.local - START
pageDetail.BookName = rD.BookName
pageDetail.FullName = rD.FullName
pageDetail.PLManage = rD.PLManage
pageDetail.PLTransfer = rD.PLTransfer
pageDetail.DerivePL = rD.DerivePL
pageDetail.CostOfCarry = rD.CostOfCarry
pageDetail.CostOfFunding = rD.CostOfFunding
pageDetail.LotAllocationMethod = rD.LotAllocationMethod
pageDetail.InternalId = rD.InternalId
// Automatically generated 04/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 04/12/2021 by matttownsend on silicon.local - END
		//

	// Automatically generated 04/12/2021 by matttownsend on silicon.local - END

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
	// Automatically generated 04/12/2021 by matttownsend on silicon.local - START
		item.BookName = r.FormValue(dm.Book_BookName)
		item.FullName = r.FormValue(dm.Book_FullName)
		item.PLManage = r.FormValue(dm.Book_PLManage)
		item.PLTransfer = r.FormValue(dm.Book_PLTransfer)
		item.DerivePL = r.FormValue(dm.Book_DerivePL)
		item.CostOfCarry = r.FormValue(dm.Book_CostOfCarry)
		item.CostOfFunding = r.FormValue(dm.Book_CostOfFunding)
		item.LotAllocationMethod = r.FormValue(dm.Book_LotAllocationMethod)
		item.InternalId = r.FormValue(dm.Book_InternalId)
	
	// Automatically generated 04/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 04/12/2021 by matttownsend on silicon.local - END

	dao.Book_Store(item)	

	http.Redirect(w, r, Book_Redirect, http.StatusFound)
}

//Book_HandlerNew is the handler used process the creation of an Book
func Book_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Book_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Book_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 04/12/2021 by matttownsend on silicon.local - START
pageDetail.BookName = ""
pageDetail.FullName = ""
pageDetail.PLManage = ""
pageDetail.PLTransfer = ""
pageDetail.DerivePL = ""
pageDetail.CostOfCarry = ""
pageDetail.CostOfFunding = ""
pageDetail.LotAllocationMethod = ""
pageDetail.InternalId = ""
// Automatically generated 04/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 04/12/2021 by matttownsend on silicon.local - END
		//

		ExecuteTemplate(dm.Book_TemplateNew, w, r, pageDetail)

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