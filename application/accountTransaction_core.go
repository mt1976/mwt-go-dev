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
// Date & Time		    : 28/06/2022 at 16:10:42
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
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

	pageDetail := dm.AccountTransaction_PageList{
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

	pageDetail := dm.AccountTransaction_Page{
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
func accounttransaction_PopulatePage(rD dm.AccountTransaction, pageDetail dm.AccountTransaction_Page) dm.AccountTransaction_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
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
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	