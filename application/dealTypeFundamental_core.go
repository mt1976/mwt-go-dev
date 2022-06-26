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
// Date & Time		    : 26/06/2022 at 18:48:28
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
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	DealTypeKey         string
	DealTypeKey_props     dm.FieldProperties
	Amendment         string
	Amendment_props     dm.FieldProperties
	DefaultFrequency         string
	DefaultFrequency_props     dm.FieldProperties
	Directions         string
	Directions_props     dm.FieldProperties
	SettledTermDealType         string
	SettledTermDealType_props     dm.FieldProperties
	Optn         string
	Optn_props     dm.FieldProperties
	AllowPledge         string
	AllowPledge_props     dm.FieldProperties
	Takeup         string
	Takeup_props     dm.FieldProperties
	MismatchDealType         string
	MismatchDealType_props     dm.FieldProperties
	SettledHedgeTermDealType         string
	SettledHedgeTermDealType_props     dm.FieldProperties
	SettlementCode         string
	SettlementCode_props     dm.FieldProperties
	TermSubType         string
	TermSubType_props     dm.FieldProperties
	FundingDealType         string
	FundingDealType_props     dm.FieldProperties
	TransferType         string
	TransferType_props     dm.FieldProperties
	TermDealType         string
	TermDealType_props     dm.FieldProperties
	NegotiableInstrumentType         string
	NegotiableInstrumentType_props     dm.FieldProperties
	Mismatch         string
	Mismatch_props     dm.FieldProperties
	ComplexTransferSubType         string
	ComplexTransferSubType_props     dm.FieldProperties
	LayOffDealType         string
	LayOffDealType_props     dm.FieldProperties
	NIAccount         string
	NIAccount_props     dm.FieldProperties
	SimpleMMsubtype         string
	SimpleMMsubtype_props     dm.FieldProperties
	SwapDealType         string
	SwapDealType_props     dm.FieldProperties
	Positions         string
	Positions_props     dm.FieldProperties
	OptionOutright         string
	OptionOutright_props     dm.FieldProperties
	SettledHedgeSpotDealType         string
	SettledHedgeSpotDealType_props     dm.FieldProperties
	StraightThroughInterestMethod         string
	StraightThroughInterestMethod_props     dm.FieldProperties
	SubType         string
	SubType_props     dm.FieldProperties
	Rollover         string
	Rollover_props     dm.FieldProperties
	DefaultIssuer         string
	DefaultIssuer_props     dm.FieldProperties
	DefaultStartDate         string
	DefaultStartDate_props     dm.FieldProperties
	Fee         string
	Fee_props     dm.FieldProperties
	NDF         string
	NDF_props     dm.FieldProperties
	FXFX         string
	FXFX_props     dm.FieldProperties
	ONIA         string
	ONIA_props     dm.FieldProperties
	MarginSubType         string
	MarginSubType_props     dm.FieldProperties
	TransferDealType         string
	TransferDealType_props     dm.FieldProperties
	IsFX         string
	IsFX_props     dm.FieldProperties
	Ordr         string
	Ordr_props     dm.FieldProperties
	OptionStyle         string
	OptionStyle_props     dm.FieldProperties
	SpotDealType         string
	SpotDealType_props     dm.FieldProperties
	CanIssue         string
	CanIssue_props     dm.FieldProperties
	CanShort         string
	CanShort_props     dm.FieldProperties
	FXMarginTradingType         string
	FXMarginTradingType_props     dm.FieldProperties
	Internal         string
	Internal_props     dm.FieldProperties
	TicketBasename         string
	TicketBasename_props     dm.FieldProperties
	InterestRateFutureType         string
	InterestRateFutureType_props     dm.FieldProperties
	TradingLimitProductCode         string
	TradingLimitProductCode_props     dm.FieldProperties
	Forward         string
	Forward_props     dm.FieldProperties
	MaturityNotificationPeriod         string
	MaturityNotificationPeriod_props     dm.FieldProperties
	NotificationEvents         string
	NotificationEvents_props     dm.FieldProperties
	SwapSubType         string
	SwapSubType_props     dm.FieldProperties
	ProductCode         string
	ProductCode_props     dm.FieldProperties
	Funding         string
	Funding_props     dm.FieldProperties
	AllocationPricing         string
	AllocationPricing_props     dm.FieldProperties
	CancelPeriod         string
	CancelPeriod_props     dm.FieldProperties
	MMMarginTradingType         string
	MMMarginTradingType_props     dm.FieldProperties
	OptionSpot         string
	OptionSpot_props     dm.FieldProperties
	Transfer         string
	Transfer_props     dm.FieldProperties
	NotificationPeriod         string
	NotificationPeriod_props     dm.FieldProperties
	Paymentdateshift         string
	Paymentdateshift_props     dm.FieldProperties
	CloseOut         string
	CloseOut_props     dm.FieldProperties
	FXOptionPricing         string
	FXOptionPricing_props     dm.FieldProperties
	SettledHedgeOutrightDealType         string
	SettledHedgeOutrightDealType_props     dm.FieldProperties
	RepoBond         string
	RepoBond_props     dm.FieldProperties
	RepoTerm         string
	RepoTerm_props     dm.FieldProperties
	RepoType         string
	RepoType_props     dm.FieldProperties
	DateRule         string
	DateRule_props     dm.FieldProperties
	CorpTransferDealType         string
	CorpTransferDealType_props     dm.FieldProperties
	GenerateFXImage         string
	GenerateFXImage_props     dm.FieldProperties
	Variant         string
	Variant_props     dm.FieldProperties
	HedgeTermDealType         string
	HedgeTermDealType_props     dm.FieldProperties
	PricingModel         string
	PricingModel_props     dm.FieldProperties
	HedgeOutrightDealType         string
	HedgeOutrightDealType_props     dm.FieldProperties
	Fixing         string
	Fixing_props     dm.FieldProperties
	Payment         string
	Payment_props     dm.FieldProperties
	MT         string
	MT_props     dm.FieldProperties
	SettlementInstructionStyle         string
	SettlementInstructionStyle_props     dm.FieldProperties
	QuoteHistoryRequired         string
	QuoteHistoryRequired_props     dm.FieldProperties
	Brokerage         string
	Brokerage_props     dm.FieldProperties
	ExposureDisabled         string
	ExposureDisabled_props     dm.FieldProperties
	CreditLine         string
	CreditLine_props     dm.FieldProperties
	Encumbered         string
	Encumbered_props     dm.FieldProperties
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
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.DealTypeFundamental_New()

	pageDetail := DealTypeFundamental_Page{
		Title:       CardTitle(dm.DealTypeFundamental_Title, core.Action_New),
		PageTitle:   PageTitle(dm.DealTypeFundamental_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = dealtypefundamental_PopulatePage(rD , pageDetail) 

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
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.DealTypeKey_props = rD.DealTypeKey_props
	pageDetail.Amendment_props = rD.Amendment_props
	pageDetail.DefaultFrequency_props = rD.DefaultFrequency_props
	pageDetail.Directions_props = rD.Directions_props
	pageDetail.SettledTermDealType_props = rD.SettledTermDealType_props
	pageDetail.Optn_props = rD.Optn_props
	pageDetail.AllowPledge_props = rD.AllowPledge_props
	pageDetail.Takeup_props = rD.Takeup_props
	pageDetail.MismatchDealType_props = rD.MismatchDealType_props
	pageDetail.SettledHedgeTermDealType_props = rD.SettledHedgeTermDealType_props
	pageDetail.SettlementCode_props = rD.SettlementCode_props
	pageDetail.TermSubType_props = rD.TermSubType_props
	pageDetail.FundingDealType_props = rD.FundingDealType_props
	pageDetail.TransferType_props = rD.TransferType_props
	pageDetail.TermDealType_props = rD.TermDealType_props
	pageDetail.NegotiableInstrumentType_props = rD.NegotiableInstrumentType_props
	pageDetail.Mismatch_props = rD.Mismatch_props
	pageDetail.ComplexTransferSubType_props = rD.ComplexTransferSubType_props
	pageDetail.LayOffDealType_props = rD.LayOffDealType_props
	pageDetail.NIAccount_props = rD.NIAccount_props
	pageDetail.SimpleMMsubtype_props = rD.SimpleMMsubtype_props
	pageDetail.SwapDealType_props = rD.SwapDealType_props
	pageDetail.Positions_props = rD.Positions_props
	pageDetail.OptionOutright_props = rD.OptionOutright_props
	pageDetail.SettledHedgeSpotDealType_props = rD.SettledHedgeSpotDealType_props
	pageDetail.StraightThroughInterestMethod_props = rD.StraightThroughInterestMethod_props
	pageDetail.SubType_props = rD.SubType_props
	pageDetail.Rollover_props = rD.Rollover_props
	pageDetail.DefaultIssuer_props = rD.DefaultIssuer_props
	pageDetail.DefaultStartDate_props = rD.DefaultStartDate_props
	pageDetail.Fee_props = rD.Fee_props
	pageDetail.NDF_props = rD.NDF_props
	pageDetail.FXFX_props = rD.FXFX_props
	pageDetail.ONIA_props = rD.ONIA_props
	pageDetail.MarginSubType_props = rD.MarginSubType_props
	pageDetail.TransferDealType_props = rD.TransferDealType_props
	pageDetail.IsFX_props = rD.IsFX_props
	pageDetail.Ordr_props = rD.Ordr_props
	pageDetail.OptionStyle_props = rD.OptionStyle_props
	pageDetail.SpotDealType_props = rD.SpotDealType_props
	pageDetail.CanIssue_props = rD.CanIssue_props
	pageDetail.CanShort_props = rD.CanShort_props
	pageDetail.FXMarginTradingType_props = rD.FXMarginTradingType_props
	pageDetail.Internal_props = rD.Internal_props
	pageDetail.TicketBasename_props = rD.TicketBasename_props
	pageDetail.InterestRateFutureType_props = rD.InterestRateFutureType_props
	pageDetail.TradingLimitProductCode_props = rD.TradingLimitProductCode_props
	pageDetail.Forward_props = rD.Forward_props
	pageDetail.MaturityNotificationPeriod_props = rD.MaturityNotificationPeriod_props
	pageDetail.NotificationEvents_props = rD.NotificationEvents_props
	pageDetail.SwapSubType_props = rD.SwapSubType_props
	pageDetail.ProductCode_props = rD.ProductCode_props
	pageDetail.Funding_props = rD.Funding_props
	pageDetail.AllocationPricing_props = rD.AllocationPricing_props
	pageDetail.CancelPeriod_props = rD.CancelPeriod_props
	pageDetail.MMMarginTradingType_props = rD.MMMarginTradingType_props
	pageDetail.OptionSpot_props = rD.OptionSpot_props
	pageDetail.Transfer_props = rD.Transfer_props
	pageDetail.NotificationPeriod_props = rD.NotificationPeriod_props
	pageDetail.Paymentdateshift_props = rD.Paymentdateshift_props
	pageDetail.CloseOut_props = rD.CloseOut_props
	pageDetail.FXOptionPricing_props = rD.FXOptionPricing_props
	pageDetail.SettledHedgeOutrightDealType_props = rD.SettledHedgeOutrightDealType_props
	pageDetail.RepoBond_props = rD.RepoBond_props
	pageDetail.RepoTerm_props = rD.RepoTerm_props
	pageDetail.RepoType_props = rD.RepoType_props
	pageDetail.DateRule_props = rD.DateRule_props
	pageDetail.CorpTransferDealType_props = rD.CorpTransferDealType_props
	pageDetail.GenerateFXImage_props = rD.GenerateFXImage_props
	pageDetail.Variant_props = rD.Variant_props
	pageDetail.HedgeTermDealType_props = rD.HedgeTermDealType_props
	pageDetail.PricingModel_props = rD.PricingModel_props
	pageDetail.HedgeOutrightDealType_props = rD.HedgeOutrightDealType_props
	pageDetail.Fixing_props = rD.Fixing_props
	pageDetail.Payment_props = rD.Payment_props
	pageDetail.MT_props = rD.MT_props
	pageDetail.SettlementInstructionStyle_props = rD.SettlementInstructionStyle_props
	pageDetail.QuoteHistoryRequired_props = rD.QuoteHistoryRequired_props
	pageDetail.Brokerage_props = rD.Brokerage_props
	pageDetail.ExposureDisabled_props = rD.ExposureDisabled_props
	pageDetail.CreditLine_props = rD.CreditLine_props
	pageDetail.Encumbered_props = rD.Encumbered_props
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