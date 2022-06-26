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

//accounttransaction_PageList provides the information for the template for a list of AccountTransactions
type AccountTransaction_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.AccountTransaction
}
//AccountTransaction_Redirect provides a page to return to aftern an action
const (
	
	AccountTransaction_Redirect = dm.AccountTransaction_PathList
	
)

//accounttransaction_Page provides the information for the template for an individual AccountTransaction
type AccountTransaction_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SienaReference         string
	SienaReference_props     dm.FieldProperties
	LegNo         string
	LegNo_props     dm.FieldProperties
	MMLegNo         string
	MMLegNo_props     dm.FieldProperties
	Narrative         string
	Narrative_props     dm.FieldProperties
	Amount         string
	Amount_props     dm.FieldProperties
	StartInterestDate         string
	StartInterestDate_props     dm.FieldProperties
	EndInterestDate         string
	EndInterestDate_props     dm.FieldProperties
	Amortisation         string
	Amortisation_props     dm.FieldProperties
	InterestAmount         string
	InterestAmount_props     dm.FieldProperties
	InterestAction         string
	InterestAction_props     dm.FieldProperties
	FixingDate         string
	FixingDate_props     dm.FieldProperties
	InterestCalculationDate         string
	InterestCalculationDate_props     dm.FieldProperties
	AmendmentAmount         string
	AmendmentAmount_props     dm.FieldProperties
	DealtCcy         string
	DealtCcy_props     dm.FieldProperties
	AmountDp         string
	AmountDp_props     dm.FieldProperties
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//AccountTransaction_Publish annouces the endpoints available for this object
func AccountTransaction_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.AccountTransaction_PathList, AccountTransaction_HandlerList)
	mux.HandleFunc(dm.AccountTransaction_PathView, AccountTransaction_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.AccountTransaction_Title)
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

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = accounttransaction_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.AccountTransaction_TemplateView, w, r, pageDetail)

}






// Builds/Popuplates the AccountTransaction Page 
func accounttransaction_PopulatePage(rD dm.AccountTransaction, pageDetail AccountTransaction_Page) AccountTransaction_Page {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
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
	
	
	//
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SienaReference_props = rD.SienaReference_props
	pageDetail.LegNo_props = rD.LegNo_props
	pageDetail.MMLegNo_props = rD.MMLegNo_props
	pageDetail.Narrative_props = rD.Narrative_props
	pageDetail.Amount_props = rD.Amount_props
	pageDetail.StartInterestDate_props = rD.StartInterestDate_props
	pageDetail.EndInterestDate_props = rD.EndInterestDate_props
	pageDetail.Amortisation_props = rD.Amortisation_props
	pageDetail.InterestAmount_props = rD.InterestAmount_props
	pageDetail.InterestAction_props = rD.InterestAction_props
	pageDetail.FixingDate_props = rD.FixingDate_props
	pageDetail.InterestCalculationDate_props = rD.InterestCalculationDate_props
	pageDetail.AmendmentAmount_props = rD.AmendmentAmount_props
	pageDetail.DealtCcy_props = rD.DealtCcy_props
	pageDetail.AmountDp_props = rD.AmountDp_props
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	