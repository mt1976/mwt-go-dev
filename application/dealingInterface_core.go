package application
// ----------------------------------------------------------------
// Automatically generated  "/application/dealinginterface.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DealingInterface (dealinginterface)
// Endpoint 	        : DealingInterface (Name)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/06/2022 at 21:32:03
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//dealinginterface_PageList provides the information for the template for a list of DealingInterfaces
type DealingInterface_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.DealingInterface
}
//DealingInterface_Redirect provides a page to return to aftern an action
const (
	DealingInterface_Redirect = dm.DealingInterface_PathList
)

//dealinginterface_Page provides the information for the template for an individual DealingInterface
type DealingInterface_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	Name         string
	AcceptReducedAmount         string
	QuoteAsIndicative         string
	RateTimeOut         string
	PropagationDelay         string
	CheckLiquidity         string
	ChangeQuoteDirection         string
	GenerateRejectedDeals         string
	SpotUpdatesForForwardQuotes         string
	SettlementInstructionStyle         string
	CanRetractQuotes         string
	CancelESPifNotPriced         string
	CancelRFQSifNotPriced         string
	CancelonDealingSuspended         string
	CreditCheckedatDI         string
	DuplicateCheckonExternalRef         string
	LimitCheckQuote         string
	LimitCheckonRFQDealSubmission         string
	ListenonLimits         string
	MarginStyle         string
	UseRerouteDefinitionOnly         string
	BypassConfirmation         string
	DIOnAcceptance         string
	IgnoreESPAmountRules         string
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//DealingInterface_Publish annouces the endpoints available for this object
func DealingInterface_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.DealingInterface_PathList, DealingInterface_HandlerList)
	mux.HandleFunc(dm.DealingInterface_PathView, DealingInterface_HandlerView)
	mux.HandleFunc(dm.DealingInterface_PathEdit, DealingInterface_HandlerEdit)
	mux.HandleFunc(dm.DealingInterface_PathNew, DealingInterface_HandlerNew)
	mux.HandleFunc(dm.DealingInterface_PathSave, DealingInterface_HandlerSave)
	mux.HandleFunc(dm.DealingInterface_PathDelete, DealingInterface_HandlerDelete)
	logs.Publish("Application", dm.DealingInterface_Title)
    //No API
}


//DealingInterface_HandlerList is the handler for the list page
func DealingInterface_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.DealingInterface
	noItems, returnList, _ := dao.DealingInterface_GetList()

	pageDetail := DealingInterface_PageList{
		Title:            CardTitle(dm.DealingInterface_Title, core.Action_List),
		PageTitle:        PageTitle(dm.DealingInterface_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.DealingInterface_TemplateList, w, r, pageDetail)

}


//DealingInterface_HandlerView is the handler used to View a page
func DealingInterface_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DealingInterface_QueryString)
	_, rD, _ := dao.DealingInterface_GetByID(searchID)

	pageDetail := DealingInterface_Page{
		Title:       CardTitle(dm.DealingInterface_Title, core.Action_View),
		PageTitle:   PageTitle(dm.DealingInterface_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = dealinginterface_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.DealingInterface_TemplateView, w, r, pageDetail)

}


//DealingInterface_HandlerEdit is the handler used generate the Edit page
func DealingInterface_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DealingInterface_QueryString)
	_, rD, _ := dao.DealingInterface_GetByID(searchID)
	
	pageDetail := DealingInterface_Page{
		Title:       CardTitle(dm.DealingInterface_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.DealingInterface_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = dealinginterface_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.DealingInterface_TemplateEdit, w, r, pageDetail)
}


//DealingInterface_HandlerSave is the handler used process the saving of an DealingInterface
func DealingInterface_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Name"))

	var item dm.DealingInterface
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.Name = r.FormValue(dm.DealingInterface_Name)
		item.AcceptReducedAmount = r.FormValue(dm.DealingInterface_AcceptReducedAmount)
		item.QuoteAsIndicative = r.FormValue(dm.DealingInterface_QuoteAsIndicative)
		item.RateTimeOut = r.FormValue(dm.DealingInterface_RateTimeOut)
		item.PropagationDelay = r.FormValue(dm.DealingInterface_PropagationDelay)
		item.CheckLiquidity = r.FormValue(dm.DealingInterface_CheckLiquidity)
		item.ChangeQuoteDirection = r.FormValue(dm.DealingInterface_ChangeQuoteDirection)
		item.GenerateRejectedDeals = r.FormValue(dm.DealingInterface_GenerateRejectedDeals)
		item.SpotUpdatesForForwardQuotes = r.FormValue(dm.DealingInterface_SpotUpdatesForForwardQuotes)
		item.SettlementInstructionStyle = r.FormValue(dm.DealingInterface_SettlementInstructionStyle)
		item.CanRetractQuotes = r.FormValue(dm.DealingInterface_CanRetractQuotes)
		item.CancelESPifNotPriced = r.FormValue(dm.DealingInterface_CancelESPifNotPriced)
		item.CancelRFQSifNotPriced = r.FormValue(dm.DealingInterface_CancelRFQSifNotPriced)
		item.CancelonDealingSuspended = r.FormValue(dm.DealingInterface_CancelonDealingSuspended)
		item.CreditCheckedatDI = r.FormValue(dm.DealingInterface_CreditCheckedatDI)
		item.DuplicateCheckonExternalRef = r.FormValue(dm.DealingInterface_DuplicateCheckonExternalRef)
		item.LimitCheckQuote = r.FormValue(dm.DealingInterface_LimitCheckQuote)
		item.LimitCheckonRFQDealSubmission = r.FormValue(dm.DealingInterface_LimitCheckonRFQDealSubmission)
		item.ListenonLimits = r.FormValue(dm.DealingInterface_ListenonLimits)
		item.MarginStyle = r.FormValue(dm.DealingInterface_MarginStyle)
		item.UseRerouteDefinitionOnly = r.FormValue(dm.DealingInterface_UseRerouteDefinitionOnly)
		item.BypassConfirmation = r.FormValue(dm.DealingInterface_BypassConfirmation)
		item.DIOnAcceptance = r.FormValue(dm.DealingInterface_DIOnAcceptance)
		item.IgnoreESPAmountRules = r.FormValue(dm.DealingInterface_IgnoreESPAmountRules)
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.DealingInterface_Store(item,r)	
	http.Redirect(w, r, DealingInterface_Redirect, http.StatusFound)
}


//DealingInterface_HandlerNew is the handler used process the creation of an DealingInterface
func DealingInterface_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := DealingInterface_Page{
		Title:       CardTitle(dm.DealingInterface_Title, core.Action_New),
		PageTitle:   PageTitle(dm.DealingInterface_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = dealinginterface_PopulatePage(dm.DealingInterface{} , pageDetail) 

	ExecuteTemplate(dm.DealingInterface_TemplateNew, w, r, pageDetail)

}	


//DealingInterface_HandlerDelete is the handler used process the deletion of an DealingInterface
func DealingInterface_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.DealingInterface_QueryString)

	dao.DealingInterface_Delete(searchID)	

	http.Redirect(w, r, DealingInterface_Redirect, http.StatusFound)
}


// Builds/Popuplates the DealingInterface Page 
func dealinginterface_PopulatePage(rD dm.DealingInterface, pageDetail DealingInterface_Page) DealingInterface_Page {
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.Name = rD.Name
	pageDetail.AcceptReducedAmount = rD.AcceptReducedAmount
	pageDetail.QuoteAsIndicative = rD.QuoteAsIndicative
	pageDetail.RateTimeOut = rD.RateTimeOut
	pageDetail.PropagationDelay = rD.PropagationDelay
	pageDetail.CheckLiquidity = rD.CheckLiquidity
	pageDetail.ChangeQuoteDirection = rD.ChangeQuoteDirection
	pageDetail.GenerateRejectedDeals = rD.GenerateRejectedDeals
	pageDetail.SpotUpdatesForForwardQuotes = rD.SpotUpdatesForForwardQuotes
	pageDetail.SettlementInstructionStyle = rD.SettlementInstructionStyle
	pageDetail.CanRetractQuotes = rD.CanRetractQuotes
	pageDetail.CancelESPifNotPriced = rD.CancelESPifNotPriced
	pageDetail.CancelRFQSifNotPriced = rD.CancelRFQSifNotPriced
	pageDetail.CancelonDealingSuspended = rD.CancelonDealingSuspended
	pageDetail.CreditCheckedatDI = rD.CreditCheckedatDI
	pageDetail.DuplicateCheckonExternalRef = rD.DuplicateCheckonExternalRef
	pageDetail.LimitCheckQuote = rD.LimitCheckQuote
	pageDetail.LimitCheckonRFQDealSubmission = rD.LimitCheckonRFQDealSubmission
	pageDetail.ListenonLimits = rD.ListenonLimits
	pageDetail.MarginStyle = rD.MarginStyle
	pageDetail.UseRerouteDefinitionOnly = rD.UseRerouteDefinitionOnly
	pageDetail.BypassConfirmation = rD.BypassConfirmation
	pageDetail.DIOnAcceptance = rD.DIOnAcceptance
	pageDetail.IgnoreESPAmountRules = rD.IgnoreESPAmountRules
	
	
	//
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	