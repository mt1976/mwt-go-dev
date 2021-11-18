package application
// ----------------------------------------------------------------
// Automatically generated  "/application/currency.go"
// ----------------------------------------------------------------
// Package            : application
// Object 			  : Currency
// Endpoint Root 	  : Currency
// Search QueryString : Code
// From   			  : 
// ----------------------------------------------------------------
// Template Generator : RussetAlbatross [r0-21.11.01]
// ----------------------------------------------------------------
// Date & Time		  : 18/11/2021 at 21:34:18
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

//currency_PageList provides the information for the template for a list of Currencys
type currency_PageList struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Currency
}

//currency_Page provides the information for the template for an individual Currency
type currency_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	AppInternalID  string
	// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
		Code string
		Name string
		AmountDp string
		Country string
		CountryName string
		IntBase string
		KeydateBase string
		InterestRateTolerance string
		CheckPayTo string
		LatinAmericanSettlement string
		DefaultLayOffBookKey string
		CutOffTimeCutOffTime string
		CutOffTimeTimeZone string
		CutOffTimeDerivedDataUTCOffset string
		CutOffTimeDerivedDataHasDaylightSaving string
		CutOffTimeDerivedDataDaylightStart string
		CutOffTimeDerivedDataDaylightEnd string
		DealerInterventionQuoteTimeout string
		CutOffTimeCutOffPeriod string
		StripRateFutureExchangeCode string
		StripRateFutureCurrencyContractCurrencyIsoCode string
		StripRateFutureCurrencyContractFutureContractCode string
		OvernightFundingSpreadBid string
		OvernightFundingSpreadOffer string
	
	// Automatically generated 18/11/2021 by matttownsend on silicon.local - END
}

const (
	Currency_Redirect = dm.Currency_PathList
)

//Currency_Publish annouces the endpoints available for this object
func Currency_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Currency_PathList, Currency_HandlerList)
	mux.HandleFunc(dm.Currency_PathView, Currency_HandlerView)
	
	
	
	
	logs.Publish("Siena", dm.Currency_Title)
}

//Currency_HandlerList is the handler for the list page
func Currency_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Currency
	noItems, returnList, _ := dao.Currency_GetList()


	pageDetail := currency_PageList{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.Currency_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Currency_TemplateList, core.GetUserRole(r)))
	t.Execute(w, pageDetail)
}

//Currency_HandlerView is the handler used to View a page
func Currency_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Currency_QueryString)
	_, rD, _ := dao.Currency_GetByID(searchID)

	pageDetail := currency_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Currency_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:		     rD.AppInternalID,
		// 
		// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
		Code: rD.Code,
		Name: rD.Name,
		AmountDp: rD.AmountDp,
		Country: rD.Country,
		CountryName: rD.CountryName,
		IntBase: rD.IntBase,
		KeydateBase: rD.KeydateBase,
		InterestRateTolerance: rD.InterestRateTolerance,
		CheckPayTo: rD.CheckPayTo,
		LatinAmericanSettlement: rD.LatinAmericanSettlement,
		DefaultLayOffBookKey: rD.DefaultLayOffBookKey,
		CutOffTimeCutOffTime: rD.CutOffTimeCutOffTime,
		CutOffTimeTimeZone: rD.CutOffTimeTimeZone,
		CutOffTimeDerivedDataUTCOffset: rD.CutOffTimeDerivedDataUTCOffset,
		CutOffTimeDerivedDataHasDaylightSaving: rD.CutOffTimeDerivedDataHasDaylightSaving,
		CutOffTimeDerivedDataDaylightStart: rD.CutOffTimeDerivedDataDaylightStart,
		CutOffTimeDerivedDataDaylightEnd: rD.CutOffTimeDerivedDataDaylightEnd,
		DealerInterventionQuoteTimeout: rD.DealerInterventionQuoteTimeout,
		CutOffTimeCutOffPeriod: rD.CutOffTimeCutOffPeriod,
		StripRateFutureExchangeCode: rD.StripRateFutureExchangeCode,
		StripRateFutureCurrencyContractCurrencyIsoCode: rD.StripRateFutureCurrencyContractCurrencyIsoCode,
		StripRateFutureCurrencyContractFutureContractCode: rD.StripRateFutureCurrencyContractFutureContractCode,
		OvernightFundingSpreadBid: rD.OvernightFundingSpreadBid,
		OvernightFundingSpreadOffer: rD.OvernightFundingSpreadOffer,
		
		// Automatically generated 18/11/2021 by matttownsend on silicon.local - END
		//
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Currency_TemplateView, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Currency_HandlerEdit is the handler used generate the Edit page
func Currency_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Currency_QueryString)
	_, rD, _ := dao.Currency_GetByID(searchID)
	
	pageDetail := currency_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Currency_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:          rD.AppInternalID,
		// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
			Code: rD.Code,
			Name: rD.Name,
			AmountDp: rD.AmountDp,
			Country: rD.Country,
			CountryName: rD.CountryName,
			IntBase: rD.IntBase,
			KeydateBase: rD.KeydateBase,
			InterestRateTolerance: rD.InterestRateTolerance,
			CheckPayTo: rD.CheckPayTo,
			LatinAmericanSettlement: rD.LatinAmericanSettlement,
			DefaultLayOffBookKey: rD.DefaultLayOffBookKey,
			CutOffTimeCutOffTime: rD.CutOffTimeCutOffTime,
			CutOffTimeTimeZone: rD.CutOffTimeTimeZone,
			CutOffTimeDerivedDataUTCOffset: rD.CutOffTimeDerivedDataUTCOffset,
			CutOffTimeDerivedDataHasDaylightSaving: rD.CutOffTimeDerivedDataHasDaylightSaving,
			CutOffTimeDerivedDataDaylightStart: rD.CutOffTimeDerivedDataDaylightStart,
			CutOffTimeDerivedDataDaylightEnd: rD.CutOffTimeDerivedDataDaylightEnd,
			DealerInterventionQuoteTimeout: rD.DealerInterventionQuoteTimeout,
			CutOffTimeCutOffPeriod: rD.CutOffTimeCutOffPeriod,
			StripRateFutureExchangeCode: rD.StripRateFutureExchangeCode,
			StripRateFutureCurrencyContractCurrencyIsoCode: rD.StripRateFutureCurrencyContractCurrencyIsoCode,
			StripRateFutureCurrencyContractFutureContractCode: rD.StripRateFutureCurrencyContractFutureContractCode,
			OvernightFundingSpreadBid: rD.OvernightFundingSpreadBid,
			OvernightFundingSpreadOffer: rD.OvernightFundingSpreadOffer,
		
	// Automatically generated 18/11/2021 by matttownsend on silicon.local - END
		//Post Import Actions - START

		// Post Import Actions - END
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Currency_TemplateEdit, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Currency_HandlerSave is the handler used process the saving of an Currency
func Currency_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ID"))

	var item dm.Currency

	//item.AppInternalID = r.FormValue("AppInternalID")
	// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
		item.Code = r.FormValue(dm.Currency_Code)
		item.Name = r.FormValue(dm.Currency_Name)
		item.AmountDp = r.FormValue(dm.Currency_AmountDp)
		item.Country = r.FormValue(dm.Currency_Country)
		item.CountryName = r.FormValue(dm.Currency_CountryName)
		item.IntBase = r.FormValue(dm.Currency_IntBase)
		item.KeydateBase = r.FormValue(dm.Currency_KeydateBase)
		item.InterestRateTolerance = r.FormValue(dm.Currency_InterestRateTolerance)
		item.CheckPayTo = r.FormValue(dm.Currency_CheckPayTo)
		item.LatinAmericanSettlement = r.FormValue(dm.Currency_LatinAmericanSettlement)
		item.DefaultLayOffBookKey = r.FormValue(dm.Currency_DefaultLayOffBookKey)
		item.CutOffTimeCutOffTime = r.FormValue(dm.Currency_CutOffTimeCutOffTime)
		item.CutOffTimeTimeZone = r.FormValue(dm.Currency_CutOffTimeTimeZone)
		item.CutOffTimeDerivedDataUTCOffset = r.FormValue(dm.Currency_CutOffTimeDerivedDataUTCOffset)
		item.CutOffTimeDerivedDataHasDaylightSaving = r.FormValue(dm.Currency_CutOffTimeDerivedDataHasDaylightSaving)
		item.CutOffTimeDerivedDataDaylightStart = r.FormValue(dm.Currency_CutOffTimeDerivedDataDaylightStart)
		item.CutOffTimeDerivedDataDaylightEnd = r.FormValue(dm.Currency_CutOffTimeDerivedDataDaylightEnd)
		item.DealerInterventionQuoteTimeout = r.FormValue(dm.Currency_DealerInterventionQuoteTimeout)
		item.CutOffTimeCutOffPeriod = r.FormValue(dm.Currency_CutOffTimeCutOffPeriod)
		item.StripRateFutureExchangeCode = r.FormValue(dm.Currency_StripRateFutureExchangeCode)
		item.StripRateFutureCurrencyContractCurrencyIsoCode = r.FormValue(dm.Currency_StripRateFutureCurrencyContractCurrencyIsoCode)
		item.StripRateFutureCurrencyContractFutureContractCode = r.FormValue(dm.Currency_StripRateFutureCurrencyContractFutureContractCode)
		item.OvernightFundingSpreadBid = r.FormValue(dm.Currency_OvernightFundingSpreadBid)
		item.OvernightFundingSpreadOffer = r.FormValue(dm.Currency_OvernightFundingSpreadOffer)
	
	// Automatically generated 18/11/2021 by matttownsend on silicon.local - END

	dao.Currency_Store(item)	

	http.Redirect(w, r, Currency_Redirect, http.StatusFound)
}

//Currency_HandlerNew is the handler used process the creation of an Currency
func Currency_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := currency_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Currency_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:   "NEW",
		// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
			Code: "",
			Name: "",
			AmountDp: "0",
			Country: "",
			CountryName: "",
			IntBase: "",
			KeydateBase: "",
			InterestRateTolerance: "0.00",
			CheckPayTo: "True",
			LatinAmericanSettlement: "True",
			DefaultLayOffBookKey: "",
			CutOffTimeCutOffTime: "",
			CutOffTimeTimeZone: "",
			CutOffTimeDerivedDataUTCOffset: "",
			CutOffTimeDerivedDataHasDaylightSaving: "True",
			CutOffTimeDerivedDataDaylightStart: "",
			CutOffTimeDerivedDataDaylightEnd: "",
			DealerInterventionQuoteTimeout: "0",
			CutOffTimeCutOffPeriod: "",
			StripRateFutureExchangeCode: "",
			StripRateFutureCurrencyContractCurrencyIsoCode: "",
			StripRateFutureCurrencyContractFutureContractCode: "",
			OvernightFundingSpreadBid: "0.00",
			OvernightFundingSpreadOffer: "0.00",
		
		// Automatically generated 18/11/2021 by matttownsend on silicon.local - END
		//
		// Post Import Actions - START

		// Post Import Actions - END
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Currency_TemplateNew, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Currency_HandlerDelete is the handler used process the deletion of an Currency
func Currency_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Currency_QueryString)

	dao.Currency_Delete(searchID)	

	http.Redirect(w, r, Currency_Redirect, http.StatusFound)
}