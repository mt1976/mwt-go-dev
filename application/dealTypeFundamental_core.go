package application
// ----------------------------------------------------------------
// Automatically generated  "/application/dealtypefundamental.go"
// ----------------------------------------------------------------
// Package            : application
// Object 			  : DealTypeFundamental
// Endpoint Root 	  : DealTypeFundamental
// Search QueryString : DealTypeKey
// From   			  : 
// ----------------------------------------------------------------
// Template Generator : RussetAlbatross [r0-21.11.01]
// ----------------------------------------------------------------
// Date & Time		  : 18/11/2021 at 21:34:19
// Who & Where		  : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"html/template"
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//dealtypefundamental_PageList provides the information for the template for a list of DealTypeFundamentals
type dealtypefundamental_PageList struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.DealTypeFundamental
}

//dealtypefundamental_Page provides the information for the template for an individual DealTypeFundamental
type dealtypefundamental_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	AppInternalID  string
	// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
		DealTypeKey string
		Amendment string
		DefaultFrequency string
		Directions string
		SettledTermDealType string
		Optn string
		AllowPledge string
		Takeup string
		MismatchDealType string
		SettledHedgeTermDealType string
		SettlementCode string
		TermSubType string
		FundingDealType string
		TransferType string
		TermDealType string
		NegotiableInstrumentType string
		Mismatch string
		ComplexTransferSubType string
		LayOffDealType string
		NIAccount string
		SimpleMMsubtype string
		SwapDealType string
		Positions string
		OptionOutright string
		SettledHedgeSpotDealType string
		StraightThroughInterestMethod string
		SubType string
		Rollover string
		DefaultIssuer string
		DefaultStartDate string
		Fee string
		NDF string
		FXFX string
		ONIA string
		MarginSubType string
		TransferDealType string
		IsFX string
		Ordr string
		OptionStyle string
		SpotDealType string
		CanIssue string
		CanShort string
		FXMarginTradingType string
		Internal string
		TicketBasename string
		InterestRateFutureType string
		TradingLimitProductCode string
		Forward string
		MaturityNotificationPeriod string
		NotificationEvents string
		SwapSubType string
		ProductCode string
		Funding string
		AllocationPricing string
		CancelPeriod string
		MMMarginTradingType string
		OptionSpot string
		Transfer string
		NotificationPeriod string
		Paymentdateshift string
		CloseOut string
		FXOptionPricing string
		SettledHedgeOutrightDealType string
		RepoBond string
		RepoTerm string
		RepoType string
		DateRule string
		CorpTransferDealType string
		GenerateFXImage string
		Variant string
		HedgeTermDealType string
		PricingModel string
		HedgeOutrightDealType string
		Fixing string
		Payment string
		MT string
		SettlementInstructionStyle string
		QuoteHistoryRequired string
		Brokerage string
		ExposureDisabled string
		ChildInheritsTradingEntity string
		InternalId string
		InternalDeleted string
		UpdatedTransactionId string
		UpdatedUserId string
		UpdatedDateTime string
		DeletedTransactionId string
		DeletedUserId string
		ChangeType string
	
	// Automatically generated 18/11/2021 by matttownsend on silicon.local - END
}

const (
	DealTypeFundamental_Redirect = dm.DealTypeFundamental_PathList
)

//DealTypeFundamental_Publish annouces the endpoints available for this object
func DealTypeFundamental_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.DealTypeFundamental_PathList, DealTypeFundamental_HandlerList)
	mux.HandleFunc(dm.DealTypeFundamental_PathView, DealTypeFundamental_HandlerView)
	mux.HandleFunc(dm.DealTypeFundamental_PathEdit, DealTypeFundamental_HandlerEdit)
	mux.HandleFunc(dm.DealTypeFundamental_PathNew, DealTypeFundamental_HandlerNew)
	mux.HandleFunc(dm.DealTypeFundamental_PathSave, DealTypeFundamental_HandlerSave)
	mux.HandleFunc(dm.DealTypeFundamental_PathDelete, DealTypeFundamental_HandlerDelete)
	logs.Publish("Siena", dm.DealTypeFundamental_Title)
}

//DealTypeFundamental_HandlerList is the handler for the list page
func DealTypeFundamental_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.DealTypeFundamental
	noItems, returnList, _ := dao.DealTypeFundamental_GetList()


	pageDetail := dealtypefundamental_PageList{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.DealTypeFundamental_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.DealTypeFundamental_TemplateList, core.GetUserRole(r)))
	t.Execute(w, pageDetail)
}

//DealTypeFundamental_HandlerView is the handler used to View a page
func DealTypeFundamental_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DealTypeFundamental_QueryString)
	_, rD, _ := dao.DealTypeFundamental_GetByID(searchID)

	pageDetail := dealtypefundamental_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.DealTypeFundamental_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:		     rD.AppInternalID,
		// 
		// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
		DealTypeKey: rD.DealTypeKey,
		Amendment: rD.Amendment,
		DefaultFrequency: rD.DefaultFrequency,
		Directions: rD.Directions,
		SettledTermDealType: rD.SettledTermDealType,
		Optn: rD.Optn,
		AllowPledge: rD.AllowPledge,
		Takeup: rD.Takeup,
		MismatchDealType: rD.MismatchDealType,
		SettledHedgeTermDealType: rD.SettledHedgeTermDealType,
		SettlementCode: rD.SettlementCode,
		TermSubType: rD.TermSubType,
		FundingDealType: rD.FundingDealType,
		TransferType: rD.TransferType,
		TermDealType: rD.TermDealType,
		NegotiableInstrumentType: rD.NegotiableInstrumentType,
		Mismatch: rD.Mismatch,
		ComplexTransferSubType: rD.ComplexTransferSubType,
		LayOffDealType: rD.LayOffDealType,
		NIAccount: rD.NIAccount,
		SimpleMMsubtype: rD.SimpleMMsubtype,
		SwapDealType: rD.SwapDealType,
		Positions: rD.Positions,
		OptionOutright: rD.OptionOutright,
		SettledHedgeSpotDealType: rD.SettledHedgeSpotDealType,
		StraightThroughInterestMethod: rD.StraightThroughInterestMethod,
		SubType: rD.SubType,
		Rollover: rD.Rollover,
		DefaultIssuer: rD.DefaultIssuer,
		DefaultStartDate: rD.DefaultStartDate,
		Fee: rD.Fee,
		NDF: rD.NDF,
		FXFX: rD.FXFX,
		ONIA: rD.ONIA,
		MarginSubType: rD.MarginSubType,
		TransferDealType: rD.TransferDealType,
		IsFX: rD.IsFX,
		Ordr: rD.Ordr,
		OptionStyle: rD.OptionStyle,
		SpotDealType: rD.SpotDealType,
		CanIssue: rD.CanIssue,
		CanShort: rD.CanShort,
		FXMarginTradingType: rD.FXMarginTradingType,
		Internal: rD.Internal,
		TicketBasename: rD.TicketBasename,
		InterestRateFutureType: rD.InterestRateFutureType,
		TradingLimitProductCode: rD.TradingLimitProductCode,
		Forward: rD.Forward,
		MaturityNotificationPeriod: rD.MaturityNotificationPeriod,
		NotificationEvents: rD.NotificationEvents,
		SwapSubType: rD.SwapSubType,
		ProductCode: rD.ProductCode,
		Funding: rD.Funding,
		AllocationPricing: rD.AllocationPricing,
		CancelPeriod: rD.CancelPeriod,
		MMMarginTradingType: rD.MMMarginTradingType,
		OptionSpot: rD.OptionSpot,
		Transfer: rD.Transfer,
		NotificationPeriod: rD.NotificationPeriod,
		Paymentdateshift: rD.Paymentdateshift,
		CloseOut: rD.CloseOut,
		FXOptionPricing: rD.FXOptionPricing,
		SettledHedgeOutrightDealType: rD.SettledHedgeOutrightDealType,
		RepoBond: rD.RepoBond,
		RepoTerm: rD.RepoTerm,
		RepoType: rD.RepoType,
		DateRule: rD.DateRule,
		CorpTransferDealType: rD.CorpTransferDealType,
		GenerateFXImage: rD.GenerateFXImage,
		Variant: rD.Variant,
		HedgeTermDealType: rD.HedgeTermDealType,
		PricingModel: rD.PricingModel,
		HedgeOutrightDealType: rD.HedgeOutrightDealType,
		Fixing: rD.Fixing,
		Payment: rD.Payment,
		MT: rD.MT,
		SettlementInstructionStyle: rD.SettlementInstructionStyle,
		QuoteHistoryRequired: rD.QuoteHistoryRequired,
		Brokerage: rD.Brokerage,
		ExposureDisabled: rD.ExposureDisabled,
		ChildInheritsTradingEntity: rD.ChildInheritsTradingEntity,
		InternalId: rD.InternalId,
		InternalDeleted: rD.InternalDeleted,
		UpdatedTransactionId: rD.UpdatedTransactionId,
		UpdatedUserId: rD.UpdatedUserId,
		UpdatedDateTime: rD.UpdatedDateTime,
		DeletedTransactionId: rD.DeletedTransactionId,
		DeletedUserId: rD.DeletedUserId,
		ChangeType: rD.ChangeType,
		
		// Automatically generated 18/11/2021 by matttownsend on silicon.local - END
		//
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.DealTypeFundamental_TemplateView, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//DealTypeFundamental_HandlerEdit is the handler used generate the Edit page
func DealTypeFundamental_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DealTypeFundamental_QueryString)
	_, rD, _ := dao.DealTypeFundamental_GetByID(searchID)
	
	pageDetail := dealtypefundamental_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.DealTypeFundamental_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:          rD.AppInternalID,
		// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
			DealTypeKey: rD.DealTypeKey,
			Amendment: rD.Amendment,
			DefaultFrequency: rD.DefaultFrequency,
			Directions: rD.Directions,
			SettledTermDealType: rD.SettledTermDealType,
			Optn: rD.Optn,
			AllowPledge: rD.AllowPledge,
			Takeup: rD.Takeup,
			MismatchDealType: rD.MismatchDealType,
			SettledHedgeTermDealType: rD.SettledHedgeTermDealType,
			SettlementCode: rD.SettlementCode,
			TermSubType: rD.TermSubType,
			FundingDealType: rD.FundingDealType,
			TransferType: rD.TransferType,
			TermDealType: rD.TermDealType,
			NegotiableInstrumentType: rD.NegotiableInstrumentType,
			Mismatch: rD.Mismatch,
			ComplexTransferSubType: rD.ComplexTransferSubType,
			LayOffDealType: rD.LayOffDealType,
			NIAccount: rD.NIAccount,
			SimpleMMsubtype: rD.SimpleMMsubtype,
			SwapDealType: rD.SwapDealType,
			Positions: rD.Positions,
			OptionOutright: rD.OptionOutright,
			SettledHedgeSpotDealType: rD.SettledHedgeSpotDealType,
			StraightThroughInterestMethod: rD.StraightThroughInterestMethod,
			SubType: rD.SubType,
			Rollover: rD.Rollover,
			DefaultIssuer: rD.DefaultIssuer,
			DefaultStartDate: rD.DefaultStartDate,
			Fee: rD.Fee,
			NDF: rD.NDF,
			FXFX: rD.FXFX,
			ONIA: rD.ONIA,
			MarginSubType: rD.MarginSubType,
			TransferDealType: rD.TransferDealType,
			IsFX: rD.IsFX,
			Ordr: rD.Ordr,
			OptionStyle: rD.OptionStyle,
			SpotDealType: rD.SpotDealType,
			CanIssue: rD.CanIssue,
			CanShort: rD.CanShort,
			FXMarginTradingType: rD.FXMarginTradingType,
			Internal: rD.Internal,
			TicketBasename: rD.TicketBasename,
			InterestRateFutureType: rD.InterestRateFutureType,
			TradingLimitProductCode: rD.TradingLimitProductCode,
			Forward: rD.Forward,
			MaturityNotificationPeriod: rD.MaturityNotificationPeriod,
			NotificationEvents: rD.NotificationEvents,
			SwapSubType: rD.SwapSubType,
			ProductCode: rD.ProductCode,
			Funding: rD.Funding,
			AllocationPricing: rD.AllocationPricing,
			CancelPeriod: rD.CancelPeriod,
			MMMarginTradingType: rD.MMMarginTradingType,
			OptionSpot: rD.OptionSpot,
			Transfer: rD.Transfer,
			NotificationPeriod: rD.NotificationPeriod,
			Paymentdateshift: rD.Paymentdateshift,
			CloseOut: rD.CloseOut,
			FXOptionPricing: rD.FXOptionPricing,
			SettledHedgeOutrightDealType: rD.SettledHedgeOutrightDealType,
			RepoBond: rD.RepoBond,
			RepoTerm: rD.RepoTerm,
			RepoType: rD.RepoType,
			DateRule: rD.DateRule,
			CorpTransferDealType: rD.CorpTransferDealType,
			GenerateFXImage: rD.GenerateFXImage,
			Variant: rD.Variant,
			HedgeTermDealType: rD.HedgeTermDealType,
			PricingModel: rD.PricingModel,
			HedgeOutrightDealType: rD.HedgeOutrightDealType,
			Fixing: rD.Fixing,
			Payment: rD.Payment,
			MT: rD.MT,
			SettlementInstructionStyle: rD.SettlementInstructionStyle,
			QuoteHistoryRequired: rD.QuoteHistoryRequired,
			Brokerage: rD.Brokerage,
			ExposureDisabled: rD.ExposureDisabled,
			ChildInheritsTradingEntity: rD.ChildInheritsTradingEntity,
			InternalId: rD.InternalId,
			InternalDeleted: rD.InternalDeleted,
			UpdatedTransactionId: rD.UpdatedTransactionId,
			UpdatedUserId: rD.UpdatedUserId,
			UpdatedDateTime: rD.UpdatedDateTime,
			DeletedTransactionId: rD.DeletedTransactionId,
			DeletedUserId: rD.DeletedUserId,
			ChangeType: rD.ChangeType,
		
	// Automatically generated 18/11/2021 by matttownsend on silicon.local - END
		//Post Import Actions - START

		// Post Import Actions - END
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.DealTypeFundamental_TemplateEdit, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//DealTypeFundamental_HandlerSave is the handler used process the saving of an DealTypeFundamental
func DealTypeFundamental_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ID"))

	var item dm.DealTypeFundamental

	//item.AppInternalID = r.FormValue("AppInternalID")
	// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
		item.DealTypeKey = r.FormValue(dm.DealTypeFundamental_DealTypeKey)
		item.Amendment = r.FormValue(dm.DealTypeFundamental_Amendment)
		item.DefaultFrequency = r.FormValue(dm.DealTypeFundamental_DefaultFrequency)
		item.Directions = r.FormValue(dm.DealTypeFundamental_Directions)
		item.SettledTermDealType = r.FormValue(dm.DealTypeFundamental_SettledTermDealType)
		item.Optn = r.FormValue(dm.DealTypeFundamental_Optn)
		item.AllowPledge = r.FormValue(dm.DealTypeFundamental_AllowPledge)
		item.Takeup = r.FormValue(dm.DealTypeFundamental_Takeup)
		item.MismatchDealType = r.FormValue(dm.DealTypeFundamental_MismatchDealType)
		item.SettledHedgeTermDealType = r.FormValue(dm.DealTypeFundamental_SettledHedgeTermDealType)
		item.SettlementCode = r.FormValue(dm.DealTypeFundamental_SettlementCode)
		item.TermSubType = r.FormValue(dm.DealTypeFundamental_TermSubType)
		item.FundingDealType = r.FormValue(dm.DealTypeFundamental_FundingDealType)
		item.TransferType = r.FormValue(dm.DealTypeFundamental_TransferType)
		item.TermDealType = r.FormValue(dm.DealTypeFundamental_TermDealType)
		item.NegotiableInstrumentType = r.FormValue(dm.DealTypeFundamental_NegotiableInstrumentType)
		item.Mismatch = r.FormValue(dm.DealTypeFundamental_Mismatch)
		item.ComplexTransferSubType = r.FormValue(dm.DealTypeFundamental_ComplexTransferSubType)
		item.LayOffDealType = r.FormValue(dm.DealTypeFundamental_LayOffDealType)
		item.NIAccount = r.FormValue(dm.DealTypeFundamental_NIAccount)
		item.SimpleMMsubtype = r.FormValue(dm.DealTypeFundamental_SimpleMMsubtype)
		item.SwapDealType = r.FormValue(dm.DealTypeFundamental_SwapDealType)
		item.Positions = r.FormValue(dm.DealTypeFundamental_Positions)
		item.OptionOutright = r.FormValue(dm.DealTypeFundamental_OptionOutright)
		item.SettledHedgeSpotDealType = r.FormValue(dm.DealTypeFundamental_SettledHedgeSpotDealType)
		item.StraightThroughInterestMethod = r.FormValue(dm.DealTypeFundamental_StraightThroughInterestMethod)
		item.SubType = r.FormValue(dm.DealTypeFundamental_SubType)
		item.Rollover = r.FormValue(dm.DealTypeFundamental_Rollover)
		item.DefaultIssuer = r.FormValue(dm.DealTypeFundamental_DefaultIssuer)
		item.DefaultStartDate = r.FormValue(dm.DealTypeFundamental_DefaultStartDate)
		item.Fee = r.FormValue(dm.DealTypeFundamental_Fee)
		item.NDF = r.FormValue(dm.DealTypeFundamental_NDF)
		item.FXFX = r.FormValue(dm.DealTypeFundamental_FXFX)
		item.ONIA = r.FormValue(dm.DealTypeFundamental_ONIA)
		item.MarginSubType = r.FormValue(dm.DealTypeFundamental_MarginSubType)
		item.TransferDealType = r.FormValue(dm.DealTypeFundamental_TransferDealType)
		item.IsFX = r.FormValue(dm.DealTypeFundamental_IsFX)
		item.Ordr = r.FormValue(dm.DealTypeFundamental_Ordr)
		item.OptionStyle = r.FormValue(dm.DealTypeFundamental_OptionStyle)
		item.SpotDealType = r.FormValue(dm.DealTypeFundamental_SpotDealType)
		item.CanIssue = r.FormValue(dm.DealTypeFundamental_CanIssue)
		item.CanShort = r.FormValue(dm.DealTypeFundamental_CanShort)
		item.FXMarginTradingType = r.FormValue(dm.DealTypeFundamental_FXMarginTradingType)
		item.Internal = r.FormValue(dm.DealTypeFundamental_Internal)
		item.TicketBasename = r.FormValue(dm.DealTypeFundamental_TicketBasename)
		item.InterestRateFutureType = r.FormValue(dm.DealTypeFundamental_InterestRateFutureType)
		item.TradingLimitProductCode = r.FormValue(dm.DealTypeFundamental_TradingLimitProductCode)
		item.Forward = r.FormValue(dm.DealTypeFundamental_Forward)
		item.MaturityNotificationPeriod = r.FormValue(dm.DealTypeFundamental_MaturityNotificationPeriod)
		item.NotificationEvents = r.FormValue(dm.DealTypeFundamental_NotificationEvents)
		item.SwapSubType = r.FormValue(dm.DealTypeFundamental_SwapSubType)
		item.ProductCode = r.FormValue(dm.DealTypeFundamental_ProductCode)
		item.Funding = r.FormValue(dm.DealTypeFundamental_Funding)
		item.AllocationPricing = r.FormValue(dm.DealTypeFundamental_AllocationPricing)
		item.CancelPeriod = r.FormValue(dm.DealTypeFundamental_CancelPeriod)
		item.MMMarginTradingType = r.FormValue(dm.DealTypeFundamental_MMMarginTradingType)
		item.OptionSpot = r.FormValue(dm.DealTypeFundamental_OptionSpot)
		item.Transfer = r.FormValue(dm.DealTypeFundamental_Transfer)
		item.NotificationPeriod = r.FormValue(dm.DealTypeFundamental_NotificationPeriod)
		item.Paymentdateshift = r.FormValue(dm.DealTypeFundamental_Paymentdateshift)
		item.CloseOut = r.FormValue(dm.DealTypeFundamental_CloseOut)
		item.FXOptionPricing = r.FormValue(dm.DealTypeFundamental_FXOptionPricing)
		item.SettledHedgeOutrightDealType = r.FormValue(dm.DealTypeFundamental_SettledHedgeOutrightDealType)
		item.RepoBond = r.FormValue(dm.DealTypeFundamental_RepoBond)
		item.RepoTerm = r.FormValue(dm.DealTypeFundamental_RepoTerm)
		item.RepoType = r.FormValue(dm.DealTypeFundamental_RepoType)
		item.DateRule = r.FormValue(dm.DealTypeFundamental_DateRule)
		item.CorpTransferDealType = r.FormValue(dm.DealTypeFundamental_CorpTransferDealType)
		item.GenerateFXImage = r.FormValue(dm.DealTypeFundamental_GenerateFXImage)
		item.Variant = r.FormValue(dm.DealTypeFundamental_Variant)
		item.HedgeTermDealType = r.FormValue(dm.DealTypeFundamental_HedgeTermDealType)
		item.PricingModel = r.FormValue(dm.DealTypeFundamental_PricingModel)
		item.HedgeOutrightDealType = r.FormValue(dm.DealTypeFundamental_HedgeOutrightDealType)
		item.Fixing = r.FormValue(dm.DealTypeFundamental_Fixing)
		item.Payment = r.FormValue(dm.DealTypeFundamental_Payment)
		item.MT = r.FormValue(dm.DealTypeFundamental_MT)
		item.SettlementInstructionStyle = r.FormValue(dm.DealTypeFundamental_SettlementInstructionStyle)
		item.QuoteHistoryRequired = r.FormValue(dm.DealTypeFundamental_QuoteHistoryRequired)
		item.Brokerage = r.FormValue(dm.DealTypeFundamental_Brokerage)
		item.ExposureDisabled = r.FormValue(dm.DealTypeFundamental_ExposureDisabled)
		item.ChildInheritsTradingEntity = r.FormValue(dm.DealTypeFundamental_ChildInheritsTradingEntity)
		item.InternalId = r.FormValue(dm.DealTypeFundamental_InternalId)
		item.InternalDeleted = r.FormValue(dm.DealTypeFundamental_InternalDeleted)
		item.UpdatedTransactionId = r.FormValue(dm.DealTypeFundamental_UpdatedTransactionId)
		item.UpdatedUserId = r.FormValue(dm.DealTypeFundamental_UpdatedUserId)
		item.UpdatedDateTime = r.FormValue(dm.DealTypeFundamental_UpdatedDateTime)
		item.DeletedTransactionId = r.FormValue(dm.DealTypeFundamental_DeletedTransactionId)
		item.DeletedUserId = r.FormValue(dm.DealTypeFundamental_DeletedUserId)
		item.ChangeType = r.FormValue(dm.DealTypeFundamental_ChangeType)
	
	// Automatically generated 18/11/2021 by matttownsend on silicon.local - END

	dao.DealTypeFundamental_Store(item)	

	http.Redirect(w, r, DealTypeFundamental_Redirect, http.StatusFound)
}

//DealTypeFundamental_HandlerNew is the handler used process the creation of an DealTypeFundamental
func DealTypeFundamental_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := dealtypefundamental_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.DealTypeFundamental_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:   "NEW",
		// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
			DealTypeKey: "",
			Amendment: "True",
			DefaultFrequency: "0",
			Directions: "",
			SettledTermDealType: "",
			Optn: "True",
			AllowPledge: "True",
			Takeup: "True",
			MismatchDealType: "",
			SettledHedgeTermDealType: "",
			SettlementCode: "",
			TermSubType: "",
			FundingDealType: "",
			TransferType: "",
			TermDealType: "",
			NegotiableInstrumentType: "",
			Mismatch: "True",
			ComplexTransferSubType: "",
			LayOffDealType: "",
			NIAccount: "0",
			SimpleMMsubtype: "0",
			SwapDealType: "",
			Positions: "",
			OptionOutright: "",
			SettledHedgeSpotDealType: "",
			StraightThroughInterestMethod: "True",
			SubType: "",
			Rollover: "True",
			DefaultIssuer: "",
			DefaultStartDate: "0",
			Fee: "",
			NDF: "True",
			FXFX: "True",
			ONIA: "True",
			MarginSubType: "0",
			TransferDealType: "",
			IsFX: "True",
			Ordr: "",
			OptionStyle: "",
			SpotDealType: "",
			CanIssue: "True",
			CanShort: "True",
			FXMarginTradingType: "0",
			Internal: "True",
			TicketBasename: "",
			InterestRateFutureType: "",
			TradingLimitProductCode: "",
			Forward: "True",
			MaturityNotificationPeriod: "",
			NotificationEvents: "",
			SwapSubType: "",
			ProductCode: "",
			Funding: "True",
			AllocationPricing: "",
			CancelPeriod: "",
			MMMarginTradingType: "0",
			OptionSpot: "",
			Transfer: "True",
			NotificationPeriod: "",
			Paymentdateshift: "0",
			CloseOut: "True",
			FXOptionPricing: "",
			SettledHedgeOutrightDealType: "",
			RepoBond: "",
			RepoTerm: "",
			RepoType: "0",
			DateRule: "",
			CorpTransferDealType: "",
			GenerateFXImage: "True",
			Variant: "",
			HedgeTermDealType: "",
			PricingModel: "",
			HedgeOutrightDealType: "",
			Fixing: "True",
			Payment: "True",
			MT: "True",
			SettlementInstructionStyle: "",
			QuoteHistoryRequired: "True",
			Brokerage: "True",
			ExposureDisabled: "True",
			ChildInheritsTradingEntity: "True",
			InternalId: "0",
			InternalDeleted: "",
			UpdatedTransactionId: "",
			UpdatedUserId: "",
			UpdatedDateTime: "",
			DeletedTransactionId: "",
			DeletedUserId: "",
			ChangeType: "",
		
		// Automatically generated 18/11/2021 by matttownsend on silicon.local - END
		//
		// Post Import Actions - START

		// Post Import Actions - END
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.DealTypeFundamental_TemplateNew, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//DealTypeFundamental_HandlerDelete is the handler used process the deletion of an DealTypeFundamental
func DealTypeFundamental_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.DealTypeFundamental_QueryString)

	dao.DealTypeFundamental_Delete(searchID)	

	http.Redirect(w, r, DealTypeFundamental_Redirect, http.StatusFound)
}