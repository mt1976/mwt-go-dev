package application
// ----------------------------------------------------------------
// Automatically generated  "/application/account.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Account (account)
// Endpoint 	        : Account (AccountNo)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 01/12/2021 at 20:36:36
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//account_PageList provides the information for the template for a list of Accounts
type Account_PageList struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Account
}

//account_Page provides the information for the template for an individual Account
type Account_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	// Automatically generated 01/12/2021 by matttownsend on silicon.local - START
		SienaReference string
		CustomerSienaView string
		SienaCommonRef string
		Status string
		StartDate string
		MaturityDate string
		ContractNumber string
		ExternalReference string
		CCY string
		Book string
		MandatedUser string
		BackOfficeNotes string
		CashBalance string
		AccountNumber string
		AccountName string
		LedgerBalance string
		Portfolio string
		AgreementId string
		BackOfficeRefNo string
		ISIN string
		UTI string
		CCYName string
		BookName string
		PortfolioName string
		Centre string
		Firm string
		CCYDp string
		CCY_Impl string
		Book_Impl string
		Portfolio_Impl string
		Centre_Impl string
		Firm_Impl string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	CCY_Impl_List	[]dm.Currency
	Book_Impl_List	[]dm.Book
	Portfolio_Impl_List	[]dm.Portfolio
	Centre_Impl_List	[]dm.Centre
	Firm_Impl_List	[]dm.Firm
	
	// Automatically generated 01/12/2021 by matttownsend on silicon.local - END
}

const (
	Account_Redirect = dm.Account_PathList
)

//Account_Publish annouces the endpoints available for this object
func Account_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Account_PathList, Account_HandlerList)
	mux.HandleFunc(dm.Account_PathView, Account_HandlerView)
	
	
	
	
	logs.Publish("Siena", dm.Account_Title)
	
}

//Account_HandlerList is the handler for the list page
func Account_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Account
	noItems, returnList, _ := dao.Account_GetList()

	pageDetail := Account_PageList{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.Account_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
	}

		ExecuteTemplate(dm.Account_TemplateList, w, r, pageDetail)

}

//Account_HandlerView is the handler used to View a page
func Account_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Account_QueryString)
	_, rD, _ := dao.Account_GetByID(searchID)

	pageDetail := Account_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Account_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
	}

		// 
		// Automatically generated 01/12/2021 by matttownsend on silicon.local - START
pageDetail.SienaReference = rD.SienaReference
pageDetail.CustomerSienaView = rD.CustomerSienaView
pageDetail.SienaCommonRef = rD.SienaCommonRef
pageDetail.Status = rD.Status
pageDetail.StartDate = rD.StartDate
pageDetail.MaturityDate = rD.MaturityDate
pageDetail.ContractNumber = rD.ContractNumber
pageDetail.ExternalReference = rD.ExternalReference
pageDetail.CCY = rD.CCY
pageDetail.Book = rD.Book
pageDetail.MandatedUser = rD.MandatedUser
pageDetail.BackOfficeNotes = rD.BackOfficeNotes
pageDetail.CashBalance = rD.CashBalance
pageDetail.AccountNumber = rD.AccountNumber
pageDetail.AccountName = rD.AccountName
pageDetail.LedgerBalance = rD.LedgerBalance
pageDetail.Portfolio = rD.Portfolio
pageDetail.AgreementId = rD.AgreementId
pageDetail.BackOfficeRefNo = rD.BackOfficeRefNo
pageDetail.ISIN = rD.ISIN
pageDetail.UTI = rD.UTI
pageDetail.CCYName = rD.CCYName
pageDetail.BookName = rD.BookName
pageDetail.PortfolioName = rD.PortfolioName
pageDetail.Centre = rD.Centre
pageDetail.Firm = rD.Firm
pageDetail.CCYDp = rD.CCYDp
// Automatically generated 01/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,CCY_Lookup,_:= dao.Currency_GetByID(rD.CCY)
pageDetail.CCY_Impl = CCY_Lookup.Name
_,Book_Lookup,_:= dao.Book_GetByID(rD.Book)
pageDetail.Book_Impl = Book_Lookup.FullName
_,Portfolio_Lookup,_:= dao.Portfolio_GetByID(rD.Portfolio)
pageDetail.Portfolio_Impl = Portfolio_Lookup.Description1
_,Centre_Lookup,_:= dao.Centre_GetByID(rD.Centre)
pageDetail.Centre_Impl = Centre_Lookup.Name
_,Firm_Lookup,_:= dao.Firm_GetByID(rD.Firm)
pageDetail.Firm_Impl = Firm_Lookup.FullName
// Automatically generated 01/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 01/12/2021 by matttownsend on silicon.local - END


		ExecuteTemplate(dm.Account_TemplateView, w, r, pageDetail)


}

//Account_HandlerEdit is the handler used generate the Edit page
func Account_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Account_QueryString)
	_, rD, _ := dao.Account_GetByID(searchID)
	
	pageDetail := Account_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Account_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
	}

		// 
		// Automatically generated 01/12/2021 by matttownsend on silicon.local - START
pageDetail.SienaReference = rD.SienaReference
pageDetail.CustomerSienaView = rD.CustomerSienaView
pageDetail.SienaCommonRef = rD.SienaCommonRef
pageDetail.Status = rD.Status
pageDetail.StartDate = rD.StartDate
pageDetail.MaturityDate = rD.MaturityDate
pageDetail.ContractNumber = rD.ContractNumber
pageDetail.ExternalReference = rD.ExternalReference
pageDetail.CCY = rD.CCY
pageDetail.Book = rD.Book
pageDetail.MandatedUser = rD.MandatedUser
pageDetail.BackOfficeNotes = rD.BackOfficeNotes
pageDetail.CashBalance = rD.CashBalance
pageDetail.AccountNumber = rD.AccountNumber
pageDetail.AccountName = rD.AccountName
pageDetail.LedgerBalance = rD.LedgerBalance
pageDetail.Portfolio = rD.Portfolio
pageDetail.AgreementId = rD.AgreementId
pageDetail.BackOfficeRefNo = rD.BackOfficeRefNo
pageDetail.ISIN = rD.ISIN
pageDetail.UTI = rD.UTI
pageDetail.CCYName = rD.CCYName
pageDetail.BookName = rD.BookName
pageDetail.PortfolioName = rD.PortfolioName
pageDetail.Centre = rD.Centre
pageDetail.Firm = rD.Firm
pageDetail.CCYDp = rD.CCYDp
// Automatically generated 01/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
_,CCY_Lookup,_:= dao.Currency_GetByID(rD.CCY)
pageDetail.CCY_Impl = CCY_Lookup.Name
_,pageDetail.CCY_Impl_List,_ = dao.Currency_GetList()
_,Book_Lookup,_:= dao.Book_GetByID(rD.Book)
pageDetail.Book_Impl = Book_Lookup.FullName
_,pageDetail.Book_Impl_List,_ = dao.Book_GetList()
_,Portfolio_Lookup,_:= dao.Portfolio_GetByID(rD.Portfolio)
pageDetail.Portfolio_Impl = Portfolio_Lookup.Description1
_,pageDetail.Portfolio_Impl_List,_ = dao.Portfolio_GetList()
_,Centre_Lookup,_:= dao.Centre_GetByID(rD.Centre)
pageDetail.Centre_Impl = Centre_Lookup.Name
_,pageDetail.Centre_Impl_List,_ = dao.Centre_GetList()
_,Firm_Lookup,_:= dao.Firm_GetByID(rD.Firm)
pageDetail.Firm_Impl = Firm_Lookup.FullName
_,pageDetail.Firm_Impl_List,_ = dao.Firm_GetList()
// Automatically generated 01/12/2021 by matttownsend on silicon.local - END
		//

	// Automatically generated 01/12/2021 by matttownsend on silicon.local - END

		ExecuteTemplate(dm.Account_TemplateEdit, w, r, pageDetail)


}

//Account_HandlerSave is the handler used process the saving of an Account
func Account_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("SienaReference"))

	var item dm.Account
	// Automatically generated 01/12/2021 by matttownsend on silicon.local - START
		item.SienaReference = r.FormValue(dm.Account_SienaReference)
		item.CustomerSienaView = r.FormValue(dm.Account_CustomerSienaView)
		item.SienaCommonRef = r.FormValue(dm.Account_SienaCommonRef)
		item.Status = r.FormValue(dm.Account_Status)
		item.StartDate = r.FormValue(dm.Account_StartDate)
		item.MaturityDate = r.FormValue(dm.Account_MaturityDate)
		item.ContractNumber = r.FormValue(dm.Account_ContractNumber)
		item.ExternalReference = r.FormValue(dm.Account_ExternalReference)
		item.CCY = r.FormValue(dm.Account_CCY)
		item.Book = r.FormValue(dm.Account_Book)
		item.MandatedUser = r.FormValue(dm.Account_MandatedUser)
		item.BackOfficeNotes = r.FormValue(dm.Account_BackOfficeNotes)
		item.CashBalance = r.FormValue(dm.Account_CashBalance)
		item.AccountNumber = r.FormValue(dm.Account_AccountNumber)
		item.AccountName = r.FormValue(dm.Account_AccountName)
		item.LedgerBalance = r.FormValue(dm.Account_LedgerBalance)
		item.Portfolio = r.FormValue(dm.Account_Portfolio)
		item.AgreementId = r.FormValue(dm.Account_AgreementId)
		item.BackOfficeRefNo = r.FormValue(dm.Account_BackOfficeRefNo)
		item.ISIN = r.FormValue(dm.Account_ISIN)
		item.UTI = r.FormValue(dm.Account_UTI)
		item.CCYName = r.FormValue(dm.Account_CCYName)
		item.BookName = r.FormValue(dm.Account_BookName)
		item.PortfolioName = r.FormValue(dm.Account_PortfolioName)
		item.Centre = r.FormValue(dm.Account_Centre)
		item.Firm = r.FormValue(dm.Account_Firm)
		item.CCYDp = r.FormValue(dm.Account_CCYDp)
		item.CCY_Impl = r.FormValue(dm.Account_CCY_Impl)
		item.Book_Impl = r.FormValue(dm.Account_Book_Impl)
		item.Portfolio_Impl = r.FormValue(dm.Account_Portfolio_Impl)
		item.Centre_Impl = r.FormValue(dm.Account_Centre_Impl)
		item.Firm_Impl = r.FormValue(dm.Account_Firm_Impl)
	
	// Automatically generated 01/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 01/12/2021 by matttownsend on silicon.local - END

	dao.Account_Store(item)	

	http.Redirect(w, r, Account_Redirect, http.StatusFound)
}

//Account_HandlerNew is the handler used process the creation of an Account
func Account_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Account_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Account_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
	}

		// 
		// Automatically generated 01/12/2021 by matttownsend on silicon.local - START
pageDetail.SienaReference = ""
pageDetail.CustomerSienaView = ""
pageDetail.SienaCommonRef = ""
pageDetail.Status = ""
pageDetail.StartDate = ""
pageDetail.MaturityDate = ""
pageDetail.ContractNumber = ""
pageDetail.ExternalReference = ""
pageDetail.CCY = ""
pageDetail.Book = ""
pageDetail.MandatedUser = ""
pageDetail.BackOfficeNotes = ""
pageDetail.CashBalance = ""
pageDetail.AccountNumber = ""
pageDetail.AccountName = ""
pageDetail.LedgerBalance = ""
pageDetail.Portfolio = ""
pageDetail.AgreementId = ""
pageDetail.BackOfficeRefNo = ""
pageDetail.ISIN = ""
pageDetail.UTI = ""
pageDetail.CCYName = ""
pageDetail.BookName = ""
pageDetail.PortfolioName = ""
pageDetail.Centre = ""
pageDetail.Firm = ""
pageDetail.CCYDp = ""
// Automatically generated 01/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
pageDetail.CCY_Impl = ""
_,pageDetail.CCY_Impl_List,_ = dao.Currency_GetList()
pageDetail.Book_Impl = ""
_,pageDetail.Book_Impl_List,_ = dao.Book_GetList()
pageDetail.Portfolio_Impl = ""
_,pageDetail.Portfolio_Impl_List,_ = dao.Portfolio_GetList()
pageDetail.Centre_Impl = ""
_,pageDetail.Centre_Impl_List,_ = dao.Centre_GetList()
pageDetail.Firm_Impl = ""
_,pageDetail.Firm_Impl_List,_ = dao.Firm_GetList()
// Automatically generated 01/12/2021 by matttownsend on silicon.local - END
		//

		ExecuteTemplate(dm.Account_TemplateNew, w, r, pageDetail)

}

//Account_HandlerDelete is the handler used process the deletion of an Account
func Account_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Account_QueryString)

	dao.Account_Delete(searchID)	

	http.Redirect(w, r, Account_Redirect, http.StatusFound)
}