package application
// ----------------------------------------------------------------
// Automatically generated  "/application/dealtypefundamental.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DealTypeFundamental (dealtypefundamental)
// Endpoint 	        : DealTypeFundamental (DealTypeKey)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:10
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//dealtypefundamental_PageList provides the information for the template for a list of DealTypeFundamentals
type DealTypeFundamental_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.DealTypeFundamental
}
//DealTypeFundamental_Redirect provides a page to return to aftern an action
const (
	DealTypeFundamental_Redirect = dm.DealTypeFundamental_PathList
)

//dealtypefundamental_Page provides the information for the template for an individual DealTypeFundamental
type DealTypeFundamental_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	DealTypeKey         string
	Amendment         string
	DefaultFrequency         string
	Directions         string
	SettledTermDealType         string
	Optn         string
	AllowPledge         string
	Takeup         string
	MismatchDealType         string
	SettledHedgeTermDealType         string
	SettlementCode         string
	TermSubType         string
	FundingDealType         string
	TransferType         string
	TermDealType         string
	NegotiableInstrumentType         string
	Mismatch         string
	ComplexTransferSubType         string
	LayOffDealType         string
	NIAccount         string
	SimpleMMsubtype         string
	SwapDealType         string
	Positions         string
	OptionOutright         string
	SettledHedgeSpotDealType         string
	StraightThroughInterestMethod         string
	SubType         string
	Rollover         string
	DefaultIssuer         string
	DefaultStartDate         string
	Fee         string
	NDF         string
	FXFX         string
	ONIA         string
	MarginSubType         string
	TransferDealType         string
	IsFX         string
	Ordr         string
	OptionStyle         string
	SpotDealType         string
	CanIssue         string
	CanShort         string
	FXMarginTradingType         string
	Internal         string
	TicketBasename         string
	InterestRateFutureType         string
	TradingLimitProductCode         string
	Forward         string
	MaturityNotificationPeriod         string
	NotificationEvents         string
	SwapSubType         string
	ProductCode         string
	Funding         string
	AllocationPricing         string
	CancelPeriod         string
	MMMarginTradingType         string
	OptionSpot         string
	Transfer         string
	NotificationPeriod         string
	Paymentdateshift         string
	CloseOut         string
	FXOptionPricing         string
	SettledHedgeOutrightDealType         string
	RepoBond         string
	RepoTerm         string
	RepoType         string
	DateRule         string
	CorpTransferDealType         string
	GenerateFXImage         string
	Variant         string
	HedgeTermDealType         string
	PricingModel         string
	HedgeOutrightDealType         string
	Fixing         string
	Payment         string
	MT         string
	SettlementInstructionStyle         string
	QuoteHistoryRequired         string
	Brokerage         string
	ExposureDisabled         string
	CreditLine         string
	Encumbered         string
	InternalId         string
	InternalDeleted         string
	UpdatedTransactionId         string
	UpdatedUserId         string
	UpdatedDateTime         string
	DeletedTransactionId         string
	DeletedUserId         string
	ChangeType         string
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//DealTypeFundamental_Publish annouces the endpoints available for this object
func DealTypeFundamental_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.DealTypeFundamental_PathList, DealTypeFundamental_HandlerList)
	mux.HandleFunc(dm.DealTypeFundamental_PathView, DealTypeFundamental_HandlerView)
	mux.HandleFunc(dm.DealTypeFundamental_PathEdit, DealTypeFundamental_HandlerEdit)
	mux.HandleFunc(dm.DealTypeFundamental_PathNew, DealTypeFundamental_HandlerNew)
	mux.HandleFunc(dm.DealTypeFundamental_PathSave, DealTypeFundamental_HandlerSave)
	mux.HandleFunc(dm.DealTypeFundamental_PathDelete, DealTypeFundamental_HandlerDelete)
	logs.Publish("Application", dm.DealTypeFundamental_Title)
    //No API
}


//DealTypeFundamental_HandlerList is the handler for the list page
func DealTypeFundamental_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.DealTypeFundamental
	noItems, returnList, _ := dao.DealTypeFundamental_GetList()

	pageDetail := DealTypeFundamental_PageList{
		Title:            CardTitle(dm.DealTypeFundamental_Title, core.Action_List),
		PageTitle:        PageTitle(dm.DealTypeFundamental_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.DealTypeFundamental_TemplateList, w, r, pageDetail)

}


//DealTypeFundamental_HandlerView is the handler used to View a page
func DealTypeFundamental_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DealTypeFundamental_QueryString)
	_, rD, _ := dao.DealTypeFundamental_GetByID(searchID)

	pageDetail := DealTypeFundamental_Page{
		Title:       CardTitle(dm.DealTypeFundamental_Title, core.Action_View),
		PageTitle:   PageTitle(dm.DealTypeFundamental_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = dealtypefundamental_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.DealTypeFundamental_TemplateView, w, r, pageDetail)

}


//DealTypeFundamental_HandlerEdit is the handler used generate the Edit page
func DealTypeFundamental_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DealTypeFundamental_QueryString)
	_, rD, _ := dao.DealTypeFundamental_GetByID(searchID)
	
	pageDetail := DealTypeFundamental_Page{
		Title:       CardTitle(dm.DealTypeFundamental_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.DealTypeFundamental_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = dealtypefundamental_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.DealTypeFundamental_TemplateEdit, w, r, pageDetail)
}


//DealTypeFundamental_HandlerSave is the handler used process the saving of an DealTypeFundamental
func DealTypeFundamental_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("DealTypeKey"))

	var item dm.DealTypeFundamental
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.DealTypeKey = r.FormValue(dm.DealTypeFundamental_DealTypeKey_scrn)
		item.Amendment = r.FormValue(dm.DealTypeFundamental_Amendment_scrn)
		item.DefaultFrequency = r.FormValue(dm.DealTypeFundamental_DefaultFrequency_scrn)
		item.Directions = r.FormValue(dm.DealTypeFundamental_Directions_scrn)
		item.SettledTermDealType = r.FormValue(dm.DealTypeFundamental_SettledTermDealType_scrn)
		item.Optn = r.FormValue(dm.DealTypeFundamental_Optn_scrn)
		item.AllowPledge = r.FormValue(dm.DealTypeFundamental_AllowPledge_scrn)
		item.Takeup = r.FormValue(dm.DealTypeFundamental_Takeup_scrn)
		item.MismatchDealType = r.FormValue(dm.DealTypeFundamental_MismatchDealType_scrn)
		item.SettledHedgeTermDealType = r.FormValue(dm.DealTypeFundamental_SettledHedgeTermDealType_scrn)
		item.SettlementCode = r.FormValue(dm.DealTypeFundamental_SettlementCode_scrn)
		item.TermSubType = r.FormValue(dm.DealTypeFundamental_TermSubType_scrn)
		item.FundingDealType = r.FormValue(dm.DealTypeFundamental_FundingDealType_scrn)
		item.TransferType = r.FormValue(dm.DealTypeFundamental_TransferType_scrn)
		item.TermDealType = r.FormValue(dm.DealTypeFundamental_TermDealType_scrn)
		item.NegotiableInstrumentType = r.FormValue(dm.DealTypeFundamental_NegotiableInstrumentType_scrn)
		item.Mismatch = r.FormValue(dm.DealTypeFundamental_Mismatch_scrn)
		item.ComplexTransferSubType = r.FormValue(dm.DealTypeFundamental_ComplexTransferSubType_scrn)
		item.LayOffDealType = r.FormValue(dm.DealTypeFundamental_LayOffDealType_scrn)
		item.NIAccount = r.FormValue(dm.DealTypeFundamental_NIAccount_scrn)
		item.SimpleMMsubtype = r.FormValue(dm.DealTypeFundamental_SimpleMMsubtype_scrn)
		item.SwapDealType = r.FormValue(dm.DealTypeFundamental_SwapDealType_scrn)
		item.Positions = r.FormValue(dm.DealTypeFundamental_Positions_scrn)
		item.OptionOutright = r.FormValue(dm.DealTypeFundamental_OptionOutright_scrn)
		item.SettledHedgeSpotDealType = r.FormValue(dm.DealTypeFundamental_SettledHedgeSpotDealType_scrn)
		item.StraightThroughInterestMethod = r.FormValue(dm.DealTypeFundamental_StraightThroughInterestMethod_scrn)
		item.SubType = r.FormValue(dm.DealTypeFundamental_SubType_scrn)
		item.Rollover = r.FormValue(dm.DealTypeFundamental_Rollover_scrn)
		item.DefaultIssuer = r.FormValue(dm.DealTypeFundamental_DefaultIssuer_scrn)
		item.DefaultStartDate = r.FormValue(dm.DealTypeFundamental_DefaultStartDate_scrn)
		item.Fee = r.FormValue(dm.DealTypeFundamental_Fee_scrn)
		item.NDF = r.FormValue(dm.DealTypeFundamental_NDF_scrn)
		item.FXFX = r.FormValue(dm.DealTypeFundamental_FXFX_scrn)
		item.ONIA = r.FormValue(dm.DealTypeFundamental_ONIA_scrn)
		item.MarginSubType = r.FormValue(dm.DealTypeFundamental_MarginSubType_scrn)
		item.TransferDealType = r.FormValue(dm.DealTypeFundamental_TransferDealType_scrn)
		item.IsFX = r.FormValue(dm.DealTypeFundamental_IsFX_scrn)
		item.Ordr = r.FormValue(dm.DealTypeFundamental_Ordr_scrn)
		item.OptionStyle = r.FormValue(dm.DealTypeFundamental_OptionStyle_scrn)
		item.SpotDealType = r.FormValue(dm.DealTypeFundamental_SpotDealType_scrn)
		item.CanIssue = r.FormValue(dm.DealTypeFundamental_CanIssue_scrn)
		item.CanShort = r.FormValue(dm.DealTypeFundamental_CanShort_scrn)
		item.FXMarginTradingType = r.FormValue(dm.DealTypeFundamental_FXMarginTradingType_scrn)
		item.Internal = r.FormValue(dm.DealTypeFundamental_Internal_scrn)
		item.TicketBasename = r.FormValue(dm.DealTypeFundamental_TicketBasename_scrn)
		item.InterestRateFutureType = r.FormValue(dm.DealTypeFundamental_InterestRateFutureType_scrn)
		item.TradingLimitProductCode = r.FormValue(dm.DealTypeFundamental_TradingLimitProductCode_scrn)
		item.Forward = r.FormValue(dm.DealTypeFundamental_Forward_scrn)
		item.MaturityNotificationPeriod = r.FormValue(dm.DealTypeFundamental_MaturityNotificationPeriod_scrn)
		item.NotificationEvents = r.FormValue(dm.DealTypeFundamental_NotificationEvents_scrn)
		item.SwapSubType = r.FormValue(dm.DealTypeFundamental_SwapSubType_scrn)
		item.ProductCode = r.FormValue(dm.DealTypeFundamental_ProductCode_scrn)
		item.Funding = r.FormValue(dm.DealTypeFundamental_Funding_scrn)
		item.AllocationPricing = r.FormValue(dm.DealTypeFundamental_AllocationPricing_scrn)
		item.CancelPeriod = r.FormValue(dm.DealTypeFundamental_CancelPeriod_scrn)
		item.MMMarginTradingType = r.FormValue(dm.DealTypeFundamental_MMMarginTradingType_scrn)
		item.OptionSpot = r.FormValue(dm.DealTypeFundamental_OptionSpot_scrn)
		item.Transfer = r.FormValue(dm.DealTypeFundamental_Transfer_scrn)
		item.NotificationPeriod = r.FormValue(dm.DealTypeFundamental_NotificationPeriod_scrn)
		item.Paymentdateshift = r.FormValue(dm.DealTypeFundamental_Paymentdateshift_scrn)
		item.CloseOut = r.FormValue(dm.DealTypeFundamental_CloseOut_scrn)
		item.FXOptionPricing = r.FormValue(dm.DealTypeFundamental_FXOptionPricing_scrn)
		item.SettledHedgeOutrightDealType = r.FormValue(dm.DealTypeFundamental_SettledHedgeOutrightDealType_scrn)
		item.RepoBond = r.FormValue(dm.DealTypeFundamental_RepoBond_scrn)
		item.RepoTerm = r.FormValue(dm.DealTypeFundamental_RepoTerm_scrn)
		item.RepoType = r.FormValue(dm.DealTypeFundamental_RepoType_scrn)
		item.DateRule = r.FormValue(dm.DealTypeFundamental_DateRule_scrn)
		item.CorpTransferDealType = r.FormValue(dm.DealTypeFundamental_CorpTransferDealType_scrn)
		item.GenerateFXImage = r.FormValue(dm.DealTypeFundamental_GenerateFXImage_scrn)
		item.Variant = r.FormValue(dm.DealTypeFundamental_Variant_scrn)
		item.HedgeTermDealType = r.FormValue(dm.DealTypeFundamental_HedgeTermDealType_scrn)
		item.PricingModel = r.FormValue(dm.DealTypeFundamental_PricingModel_scrn)
		item.HedgeOutrightDealType = r.FormValue(dm.DealTypeFundamental_HedgeOutrightDealType_scrn)
		item.Fixing = r.FormValue(dm.DealTypeFundamental_Fixing_scrn)
		item.Payment = r.FormValue(dm.DealTypeFundamental_Payment_scrn)
		item.MT = r.FormValue(dm.DealTypeFundamental_MT_scrn)
		item.SettlementInstructionStyle = r.FormValue(dm.DealTypeFundamental_SettlementInstructionStyle_scrn)
		item.QuoteHistoryRequired = r.FormValue(dm.DealTypeFundamental_QuoteHistoryRequired_scrn)
		item.Brokerage = r.FormValue(dm.DealTypeFundamental_Brokerage_scrn)
		item.ExposureDisabled = r.FormValue(dm.DealTypeFundamental_ExposureDisabled_scrn)
		item.CreditLine = r.FormValue(dm.DealTypeFundamental_CreditLine_scrn)
		item.Encumbered = r.FormValue(dm.DealTypeFundamental_Encumbered_scrn)
		item.InternalId = r.FormValue(dm.DealTypeFundamental_InternalId_scrn)
		item.InternalDeleted = r.FormValue(dm.DealTypeFundamental_InternalDeleted_scrn)
		item.UpdatedTransactionId = r.FormValue(dm.DealTypeFundamental_UpdatedTransactionId_scrn)
		item.UpdatedUserId = r.FormValue(dm.DealTypeFundamental_UpdatedUserId_scrn)
		item.UpdatedDateTime = r.FormValue(dm.DealTypeFundamental_UpdatedDateTime_scrn)
		item.DeletedTransactionId = r.FormValue(dm.DealTypeFundamental_DeletedTransactionId_scrn)
		item.DeletedUserId = r.FormValue(dm.DealTypeFundamental_DeletedUserId_scrn)
		item.ChangeType = r.FormValue(dm.DealTypeFundamental_ChangeType_scrn)
	
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.DealTypeFundamental_Store(item,r)	
	http.Redirect(w, r, DealTypeFundamental_Redirect, http.StatusFound)
}


//DealTypeFundamental_HandlerNew is the handler used process the creation of an DealTypeFundamental
func DealTypeFundamental_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := DealTypeFundamental_Page{
		Title:       CardTitle(dm.DealTypeFundamental_Title, core.Action_New),
		PageTitle:   PageTitle(dm.DealTypeFundamental_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = dealtypefundamental_PopulatePage(dm.DealTypeFundamental{} , pageDetail) 

	ExecuteTemplate(dm.DealTypeFundamental_TemplateNew, w, r, pageDetail)

}	


//DealTypeFundamental_HandlerDelete is the handler used process the deletion of an DealTypeFundamental
func DealTypeFundamental_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.DealTypeFundamental_QueryString)

	dao.DealTypeFundamental_Delete(searchID)	

	http.Redirect(w, r, DealTypeFundamental_Redirect, http.StatusFound)
}


// Builds/Popuplates the DealTypeFundamental Page 
func dealtypefundamental_PopulatePage(rD dm.DealTypeFundamental, pageDetail DealTypeFundamental_Page) DealTypeFundamental_Page {
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.DealTypeKey = rD.DealTypeKey
	pageDetail.Amendment = rD.Amendment
	pageDetail.DefaultFrequency = rD.DefaultFrequency
	pageDetail.Directions = rD.Directions
	pageDetail.SettledTermDealType = rD.SettledTermDealType
	pageDetail.Optn = rD.Optn
	pageDetail.AllowPledge = rD.AllowPledge
	pageDetail.Takeup = rD.Takeup
	pageDetail.MismatchDealType = rD.MismatchDealType
	pageDetail.SettledHedgeTermDealType = rD.SettledHedgeTermDealType
	pageDetail.SettlementCode = rD.SettlementCode
	pageDetail.TermSubType = rD.TermSubType
	pageDetail.FundingDealType = rD.FundingDealType
	pageDetail.TransferType = rD.TransferType
	pageDetail.TermDealType = rD.TermDealType
	pageDetail.NegotiableInstrumentType = rD.NegotiableInstrumentType
	pageDetail.Mismatch = rD.Mismatch
	pageDetail.ComplexTransferSubType = rD.ComplexTransferSubType
	pageDetail.LayOffDealType = rD.LayOffDealType
	pageDetail.NIAccount = rD.NIAccount
	pageDetail.SimpleMMsubtype = rD.SimpleMMsubtype
	pageDetail.SwapDealType = rD.SwapDealType
	pageDetail.Positions = rD.Positions
	pageDetail.OptionOutright = rD.OptionOutright
	pageDetail.SettledHedgeSpotDealType = rD.SettledHedgeSpotDealType
	pageDetail.StraightThroughInterestMethod = rD.StraightThroughInterestMethod
	pageDetail.SubType = rD.SubType
	pageDetail.Rollover = rD.Rollover
	pageDetail.DefaultIssuer = rD.DefaultIssuer
	pageDetail.DefaultStartDate = rD.DefaultStartDate
	pageDetail.Fee = rD.Fee
	pageDetail.NDF = rD.NDF
	pageDetail.FXFX = rD.FXFX
	pageDetail.ONIA = rD.ONIA
	pageDetail.MarginSubType = rD.MarginSubType
	pageDetail.TransferDealType = rD.TransferDealType
	pageDetail.IsFX = rD.IsFX
	pageDetail.Ordr = rD.Ordr
	pageDetail.OptionStyle = rD.OptionStyle
	pageDetail.SpotDealType = rD.SpotDealType
	pageDetail.CanIssue = rD.CanIssue
	pageDetail.CanShort = rD.CanShort
	pageDetail.FXMarginTradingType = rD.FXMarginTradingType
	pageDetail.Internal = rD.Internal
	pageDetail.TicketBasename = rD.TicketBasename
	pageDetail.InterestRateFutureType = rD.InterestRateFutureType
	pageDetail.TradingLimitProductCode = rD.TradingLimitProductCode
	pageDetail.Forward = rD.Forward
	pageDetail.MaturityNotificationPeriod = rD.MaturityNotificationPeriod
	pageDetail.NotificationEvents = rD.NotificationEvents
	pageDetail.SwapSubType = rD.SwapSubType
	pageDetail.ProductCode = rD.ProductCode
	pageDetail.Funding = rD.Funding
	pageDetail.AllocationPricing = rD.AllocationPricing
	pageDetail.CancelPeriod = rD.CancelPeriod
	pageDetail.MMMarginTradingType = rD.MMMarginTradingType
	pageDetail.OptionSpot = rD.OptionSpot
	pageDetail.Transfer = rD.Transfer
	pageDetail.NotificationPeriod = rD.NotificationPeriod
	pageDetail.Paymentdateshift = rD.Paymentdateshift
	pageDetail.CloseOut = rD.CloseOut
	pageDetail.FXOptionPricing = rD.FXOptionPricing
	pageDetail.SettledHedgeOutrightDealType = rD.SettledHedgeOutrightDealType
	pageDetail.RepoBond = rD.RepoBond
	pageDetail.RepoTerm = rD.RepoTerm
	pageDetail.RepoType = rD.RepoType
	pageDetail.DateRule = rD.DateRule
	pageDetail.CorpTransferDealType = rD.CorpTransferDealType
	pageDetail.GenerateFXImage = rD.GenerateFXImage
	pageDetail.Variant = rD.Variant
	pageDetail.HedgeTermDealType = rD.HedgeTermDealType
	pageDetail.PricingModel = rD.PricingModel
	pageDetail.HedgeOutrightDealType = rD.HedgeOutrightDealType
	pageDetail.Fixing = rD.Fixing
	pageDetail.Payment = rD.Payment
	pageDetail.MT = rD.MT
	pageDetail.SettlementInstructionStyle = rD.SettlementInstructionStyle
	pageDetail.QuoteHistoryRequired = rD.QuoteHistoryRequired
	pageDetail.Brokerage = rD.Brokerage
	pageDetail.ExposureDisabled = rD.ExposureDisabled
	pageDetail.CreditLine = rD.CreditLine
	pageDetail.Encumbered = rD.Encumbered
	pageDetail.InternalId = rD.InternalId
	pageDetail.InternalDeleted = rD.InternalDeleted
	pageDetail.UpdatedTransactionId = rD.UpdatedTransactionId
	pageDetail.UpdatedUserId = rD.UpdatedUserId
	pageDetail.UpdatedDateTime = rD.UpdatedDateTime
	pageDetail.DeletedTransactionId = rD.DeletedTransactionId
	pageDetail.DeletedUserId = rD.DeletedUserId
	pageDetail.ChangeType = rD.ChangeType
	
	
	//
	// Automatically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	