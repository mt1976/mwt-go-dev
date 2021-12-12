package application
// ----------------------------------------------------------------
// Automatically generated  "/application/accounttransaction.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : AccountTransaction (accounttransaction)
// Endpoint 	        : AccountTransaction (AccountNo)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:06
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//accounttransaction_PageList provides the information for the template for a list of AccountTransactions
type AccountTransaction_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.AccountTransaction
}

//accounttransaction_Page provides the information for the template for an individual AccountTransaction
type AccountTransaction_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		SienaReference string
		LegNo string
		MMLegNo string
		Narrative string
		Amount string
		StartInterestDate string
		EndInterestDate string
		Amortisation string
		InterestAmount string
		InterestAction string
		FixingDate string
		InterestCalculationDate string
		AmendmentAmount string
		DealtCcy string
		AmountDp string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	AccountTransaction_Redirect = dm.AccountTransaction_PathList
)

//AccountTransaction_Publish annouces the endpoints available for this object
func AccountTransaction_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.AccountTransaction_PathList, AccountTransaction_HandlerList)
	mux.HandleFunc(dm.AccountTransaction_PathView, AccountTransaction_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Siena", dm.AccountTransaction_Title)
    //No API
}

//AccountTransaction_HandlerList is the handler for the list page
func AccountTransaction_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.AccountTransaction
	noItems, returnList, _ := dao.AccountTransaction_GetList()

	pageDetail := AccountTransaction_PageList{
		Title:            CardTitle(dm.AccountTransaction_Title, core.Action_List),
		PageTitle:        PageTitle(dm.AccountTransaction_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.AccountTransaction_TemplateList, w, r, pageDetail)

}

//AccountTransaction_HandlerView is the handler used to View a page
func AccountTransaction_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.AccountTransaction_QueryString)
	_, rD, _ := dao.AccountTransaction_GetByID(searchID)

	pageDetail := AccountTransaction_Page{
		Title:       CardTitle(dm.AccountTransaction_Title, core.Action_View),
		PageTitle:   PageTitle(dm.AccountTransaction_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.SienaReference = rD.SienaReference
pageDetail.LegNo = rD.LegNo
pageDetail.MMLegNo = rD.MMLegNo
pageDetail.Narrative = rD.Narrative
pageDetail.Amount = rD.Amount
pageDetail.StartInterestDate = rD.StartInterestDate
pageDetail.EndInterestDate = rD.EndInterestDate
pageDetail.Amortisation = rD.Amortisation
pageDetail.InterestAmount = rD.InterestAmount
pageDetail.InterestAction = rD.InterestAction
pageDetail.FixingDate = rD.FixingDate
pageDetail.InterestCalculationDate = rD.InterestCalculationDate
pageDetail.AmendmentAmount = rD.AmendmentAmount
pageDetail.DealtCcy = rD.DealtCcy
pageDetail.AmountDp = rD.AmountDp


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.AccountTransaction_TemplateView, w, r, pageDetail)

}

//AccountTransaction_HandlerEdit is the handler used generate the Edit page
func AccountTransaction_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.AccountTransaction_QueryString)
	_, rD, _ := dao.AccountTransaction_GetByID(searchID)
	
	pageDetail := AccountTransaction_Page{
		Title:       CardTitle(dm.AccountTransaction_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.AccountTransaction_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.SienaReference = rD.SienaReference
pageDetail.LegNo = rD.LegNo
pageDetail.MMLegNo = rD.MMLegNo
pageDetail.Narrative = rD.Narrative
pageDetail.Amount = rD.Amount
pageDetail.StartInterestDate = rD.StartInterestDate
pageDetail.EndInterestDate = rD.EndInterestDate
pageDetail.Amortisation = rD.Amortisation
pageDetail.InterestAmount = rD.InterestAmount
pageDetail.InterestAction = rD.InterestAction
pageDetail.FixingDate = rD.FixingDate
pageDetail.InterestCalculationDate = rD.InterestCalculationDate
pageDetail.AmendmentAmount = rD.AmendmentAmount
pageDetail.DealtCcy = rD.DealtCcy
pageDetail.AmountDp = rD.AmountDp


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.AccountTransaction_TemplateEdit, w, r, pageDetail)


}

//AccountTransaction_HandlerSave is the handler used process the saving of an AccountTransaction
func AccountTransaction_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("SienaReference"))

	var item dm.AccountTransaction
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.SienaReference = r.FormValue(dm.AccountTransaction_SienaReference)
		item.LegNo = r.FormValue(dm.AccountTransaction_LegNo)
		item.MMLegNo = r.FormValue(dm.AccountTransaction_MMLegNo)
		item.Narrative = r.FormValue(dm.AccountTransaction_Narrative)
		item.Amount = r.FormValue(dm.AccountTransaction_Amount)
		item.StartInterestDate = r.FormValue(dm.AccountTransaction_StartInterestDate)
		item.EndInterestDate = r.FormValue(dm.AccountTransaction_EndInterestDate)
		item.Amortisation = r.FormValue(dm.AccountTransaction_Amortisation)
		item.InterestAmount = r.FormValue(dm.AccountTransaction_InterestAmount)
		item.InterestAction = r.FormValue(dm.AccountTransaction_InterestAction)
		item.FixingDate = r.FormValue(dm.AccountTransaction_FixingDate)
		item.InterestCalculationDate = r.FormValue(dm.AccountTransaction_InterestCalculationDate)
		item.AmendmentAmount = r.FormValue(dm.AccountTransaction_AmendmentAmount)
		item.DealtCcy = r.FormValue(dm.AccountTransaction_DealtCcy)
		item.AmountDp = r.FormValue(dm.AccountTransaction_AmountDp)
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.AccountTransaction_Store(item,r)	

	http.Redirect(w, r, AccountTransaction_Redirect, http.StatusFound)
}

//AccountTransaction_HandlerNew is the handler used process the creation of an AccountTransaction
func AccountTransaction_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := AccountTransaction_Page{
		Title:       CardTitle(dm.AccountTransaction_Title, core.Action_New),
		PageTitle:   PageTitle(dm.AccountTransaction_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.SienaReference = ""
pageDetail.LegNo = ""
pageDetail.MMLegNo = ""
pageDetail.Narrative = ""
pageDetail.Amount = ""
pageDetail.StartInterestDate = ""
pageDetail.EndInterestDate = ""
pageDetail.Amortisation = ""
pageDetail.InterestAmount = ""
pageDetail.InterestAction = ""
pageDetail.FixingDate = ""
pageDetail.InterestCalculationDate = ""
pageDetail.AmendmentAmount = ""
pageDetail.DealtCcy = ""
pageDetail.AmountDp = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.AccountTransaction_TemplateNew, w, r, pageDetail)

}

//AccountTransaction_HandlerDelete is the handler used process the deletion of an AccountTransaction
func AccountTransaction_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.AccountTransaction_QueryString)

	dao.AccountTransaction_Delete(searchID)	

	http.Redirect(w, r, AccountTransaction_Redirect, http.StatusFound)
}