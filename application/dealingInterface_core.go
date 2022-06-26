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
// Date & Time		    : 26/06/2022 at 18:48:27
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
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	Name         string
	Name_props     dm.FieldProperties
	AcceptReducedAmount         string
	AcceptReducedAmount_props     dm.FieldProperties
	QuoteAsIndicative         string
	QuoteAsIndicative_props     dm.FieldProperties
	RateTimeOut         string
	RateTimeOut_props     dm.FieldProperties
	PropagationDelay         string
	PropagationDelay_props     dm.FieldProperties
	CheckLiquidity         string
	CheckLiquidity_props     dm.FieldProperties
	ChangeQuoteDirection         string
	ChangeQuoteDirection_props     dm.FieldProperties
	GenerateRejectedDeals         string
	GenerateRejectedDeals_props     dm.FieldProperties
	SpotUpdatesForForwardQuotes         string
	SpotUpdatesForForwardQuotes_props     dm.FieldProperties
	SettlementInstructionStyle         string
	SettlementInstructionStyle_props     dm.FieldProperties
	CanRetractQuotes         string
	CanRetractQuotes_props     dm.FieldProperties
	CancelESPifNotPriced         string
	CancelESPifNotPriced_props     dm.FieldProperties
	CancelRFQSifNotPriced         string
	CancelRFQSifNotPriced_props     dm.FieldProperties
	CancelonDealingSuspended         string
	CancelonDealingSuspended_props     dm.FieldProperties
	CreditCheckedatDI         string
	CreditCheckedatDI_props     dm.FieldProperties
	DuplicateCheckonExternalRef         string
	DuplicateCheckonExternalRef_props     dm.FieldProperties
	LimitCheckQuote         string
	LimitCheckQuote_props     dm.FieldProperties
	LimitCheckonRFQDealSubmission         string
	LimitCheckonRFQDealSubmission_props     dm.FieldProperties
	ListenonLimits         string
	ListenonLimits_props     dm.FieldProperties
	MarginStyle         string
	MarginStyle_props     dm.FieldProperties
	UseRerouteDefinitionOnly         string
	UseRerouteDefinitionOnly_props     dm.FieldProperties
	BypassConfirmation         string
	BypassConfirmation_props     dm.FieldProperties
	DIOnAcceptance         string
	DIOnAcceptance_props     dm.FieldProperties
	IgnoreESPAmountRules         string
	IgnoreESPAmountRules_props     dm.FieldProperties
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.Name = r.FormValue(dm.DealingInterface_Name_scrn)
		item.AcceptReducedAmount = r.FormValue(dm.DealingInterface_AcceptReducedAmount_scrn)
		item.QuoteAsIndicative = r.FormValue(dm.DealingInterface_QuoteAsIndicative_scrn)
		item.RateTimeOut = r.FormValue(dm.DealingInterface_RateTimeOut_scrn)
		item.PropagationDelay = r.FormValue(dm.DealingInterface_PropagationDelay_scrn)
		item.CheckLiquidity = r.FormValue(dm.DealingInterface_CheckLiquidity_scrn)
		item.ChangeQuoteDirection = r.FormValue(dm.DealingInterface_ChangeQuoteDirection_scrn)
		item.GenerateRejectedDeals = r.FormValue(dm.DealingInterface_GenerateRejectedDeals_scrn)
		item.SpotUpdatesForForwardQuotes = r.FormValue(dm.DealingInterface_SpotUpdatesForForwardQuotes_scrn)
		item.SettlementInstructionStyle = r.FormValue(dm.DealingInterface_SettlementInstructionStyle_scrn)
		item.CanRetractQuotes = r.FormValue(dm.DealingInterface_CanRetractQuotes_scrn)
		item.CancelESPifNotPriced = r.FormValue(dm.DealingInterface_CancelESPifNotPriced_scrn)
		item.CancelRFQSifNotPriced = r.FormValue(dm.DealingInterface_CancelRFQSifNotPriced_scrn)
		item.CancelonDealingSuspended = r.FormValue(dm.DealingInterface_CancelonDealingSuspended_scrn)
		item.CreditCheckedatDI = r.FormValue(dm.DealingInterface_CreditCheckedatDI_scrn)
		item.DuplicateCheckonExternalRef = r.FormValue(dm.DealingInterface_DuplicateCheckonExternalRef_scrn)
		item.LimitCheckQuote = r.FormValue(dm.DealingInterface_LimitCheckQuote_scrn)
		item.LimitCheckonRFQDealSubmission = r.FormValue(dm.DealingInterface_LimitCheckonRFQDealSubmission_scrn)
		item.ListenonLimits = r.FormValue(dm.DealingInterface_ListenonLimits_scrn)
		item.MarginStyle = r.FormValue(dm.DealingInterface_MarginStyle_scrn)
		item.UseRerouteDefinitionOnly = r.FormValue(dm.DealingInterface_UseRerouteDefinitionOnly_scrn)
		item.BypassConfirmation = r.FormValue(dm.DealingInterface_BypassConfirmation_scrn)
		item.DIOnAcceptance = r.FormValue(dm.DealingInterface_DIOnAcceptance_scrn)
		item.IgnoreESPAmountRules = r.FormValue(dm.DealingInterface_IgnoreESPAmountRules_scrn)
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.DealingInterface_New()

	pageDetail := DealingInterface_Page{
		Title:       CardTitle(dm.DealingInterface_Title, core.Action_New),
		PageTitle:   PageTitle(dm.DealingInterface_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = dealinginterface_PopulatePage(rD , pageDetail) 

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
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.Name_props = rD.Name_props
	pageDetail.AcceptReducedAmount_props = rD.AcceptReducedAmount_props
	pageDetail.QuoteAsIndicative_props = rD.QuoteAsIndicative_props
	pageDetail.RateTimeOut_props = rD.RateTimeOut_props
	pageDetail.PropagationDelay_props = rD.PropagationDelay_props
	pageDetail.CheckLiquidity_props = rD.CheckLiquidity_props
	pageDetail.ChangeQuoteDirection_props = rD.ChangeQuoteDirection_props
	pageDetail.GenerateRejectedDeals_props = rD.GenerateRejectedDeals_props
	pageDetail.SpotUpdatesForForwardQuotes_props = rD.SpotUpdatesForForwardQuotes_props
	pageDetail.SettlementInstructionStyle_props = rD.SettlementInstructionStyle_props
	pageDetail.CanRetractQuotes_props = rD.CanRetractQuotes_props
	pageDetail.CancelESPifNotPriced_props = rD.CancelESPifNotPriced_props
	pageDetail.CancelRFQSifNotPriced_props = rD.CancelRFQSifNotPriced_props
	pageDetail.CancelonDealingSuspended_props = rD.CancelonDealingSuspended_props
	pageDetail.CreditCheckedatDI_props = rD.CreditCheckedatDI_props
	pageDetail.DuplicateCheckonExternalRef_props = rD.DuplicateCheckonExternalRef_props
	pageDetail.LimitCheckQuote_props = rD.LimitCheckQuote_props
	pageDetail.LimitCheckonRFQDealSubmission_props = rD.LimitCheckonRFQDealSubmission_props
	pageDetail.ListenonLimits_props = rD.ListenonLimits_props
	pageDetail.MarginStyle_props = rD.MarginStyle_props
	pageDetail.UseRerouteDefinitionOnly_props = rD.UseRerouteDefinitionOnly_props
	pageDetail.BypassConfirmation_props = rD.BypassConfirmation_props
	pageDetail.DIOnAcceptance_props = rD.DIOnAcceptance_props
	pageDetail.IgnoreESPAmountRules_props = rD.IgnoreESPAmountRules_props
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	