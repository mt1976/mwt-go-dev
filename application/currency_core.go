package application
// ----------------------------------------------------------------
// Automatically generated  "/application/currency.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Currency (currency)
// Endpoint 	        : Currency (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:11
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//currency_PageList provides the information for the template for a list of Currencys
type Currency_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Currency
}

//currency_Page provides the information for the template for an individual Currency
type Currency_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	Currency_Redirect = dm.Currency_PathList
)

//Currency_Publish annouces the endpoints available for this object
func Currency_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Currency_Path, Currency_Handler)
	mux.HandleFunc(dm.Currency_PathList, Currency_HandlerList)
	mux.HandleFunc(dm.Currency_PathView, Currency_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Siena", dm.Currency_Title)
    core.Catalog_Add(dm.Currency_Title, dm.Currency_Path, "", dm.Currency_QueryString, "APP")
}

//Currency_HandlerList is the handler for the list page
func Currency_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Currency
	noItems, returnList, _ := dao.Currency_GetList()

	pageDetail := Currency_PageList{
		Title:            CardTitle(dm.Currency_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Currency_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Currency_TemplateList, w, r, pageDetail)

}

//Currency_HandlerView is the handler used to View a page
func Currency_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Currency_QueryString)
	_, rD, _ := dao.Currency_GetByID(searchID)

	pageDetail := Currency_Page{
		Title:       CardTitle(dm.Currency_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Currency_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
pageDetail.AmountDp = rD.AmountDp
pageDetail.Country = rD.Country
pageDetail.CountryName = rD.CountryName
pageDetail.IntBase = rD.IntBase
pageDetail.KeydateBase = rD.KeydateBase
pageDetail.InterestRateTolerance = rD.InterestRateTolerance
pageDetail.CheckPayTo = rD.CheckPayTo
pageDetail.LatinAmericanSettlement = rD.LatinAmericanSettlement
pageDetail.DefaultLayOffBookKey = rD.DefaultLayOffBookKey
pageDetail.CutOffTimeCutOffTime = rD.CutOffTimeCutOffTime
pageDetail.CutOffTimeTimeZone = rD.CutOffTimeTimeZone
pageDetail.CutOffTimeDerivedDataUTCOffset = rD.CutOffTimeDerivedDataUTCOffset
pageDetail.CutOffTimeDerivedDataHasDaylightSaving = rD.CutOffTimeDerivedDataHasDaylightSaving
pageDetail.CutOffTimeDerivedDataDaylightStart = rD.CutOffTimeDerivedDataDaylightStart
pageDetail.CutOffTimeDerivedDataDaylightEnd = rD.CutOffTimeDerivedDataDaylightEnd
pageDetail.DealerInterventionQuoteTimeout = rD.DealerInterventionQuoteTimeout
pageDetail.CutOffTimeCutOffPeriod = rD.CutOffTimeCutOffPeriod
pageDetail.StripRateFutureExchangeCode = rD.StripRateFutureExchangeCode
pageDetail.StripRateFutureCurrencyContractCurrencyIsoCode = rD.StripRateFutureCurrencyContractCurrencyIsoCode
pageDetail.StripRateFutureCurrencyContractFutureContractCode = rD.StripRateFutureCurrencyContractFutureContractCode
pageDetail.OvernightFundingSpreadBid = rD.OvernightFundingSpreadBid
pageDetail.OvernightFundingSpreadOffer = rD.OvernightFundingSpreadOffer


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Currency_TemplateView, w, r, pageDetail)

}

//Currency_HandlerEdit is the handler used generate the Edit page
func Currency_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Currency_QueryString)
	_, rD, _ := dao.Currency_GetByID(searchID)
	
	pageDetail := Currency_Page{
		Title:       CardTitle(dm.Currency_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.Currency_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.Code = rD.Code
pageDetail.Name = rD.Name
pageDetail.AmountDp = rD.AmountDp
pageDetail.Country = rD.Country
pageDetail.CountryName = rD.CountryName
pageDetail.IntBase = rD.IntBase
pageDetail.KeydateBase = rD.KeydateBase
pageDetail.InterestRateTolerance = rD.InterestRateTolerance
pageDetail.CheckPayTo = rD.CheckPayTo
pageDetail.LatinAmericanSettlement = rD.LatinAmericanSettlement
pageDetail.DefaultLayOffBookKey = rD.DefaultLayOffBookKey
pageDetail.CutOffTimeCutOffTime = rD.CutOffTimeCutOffTime
pageDetail.CutOffTimeTimeZone = rD.CutOffTimeTimeZone
pageDetail.CutOffTimeDerivedDataUTCOffset = rD.CutOffTimeDerivedDataUTCOffset
pageDetail.CutOffTimeDerivedDataHasDaylightSaving = rD.CutOffTimeDerivedDataHasDaylightSaving
pageDetail.CutOffTimeDerivedDataDaylightStart = rD.CutOffTimeDerivedDataDaylightStart
pageDetail.CutOffTimeDerivedDataDaylightEnd = rD.CutOffTimeDerivedDataDaylightEnd
pageDetail.DealerInterventionQuoteTimeout = rD.DealerInterventionQuoteTimeout
pageDetail.CutOffTimeCutOffPeriod = rD.CutOffTimeCutOffPeriod
pageDetail.StripRateFutureExchangeCode = rD.StripRateFutureExchangeCode
pageDetail.StripRateFutureCurrencyContractCurrencyIsoCode = rD.StripRateFutureCurrencyContractCurrencyIsoCode
pageDetail.StripRateFutureCurrencyContractFutureContractCode = rD.StripRateFutureCurrencyContractFutureContractCode
pageDetail.OvernightFundingSpreadBid = rD.OvernightFundingSpreadBid
pageDetail.OvernightFundingSpreadOffer = rD.OvernightFundingSpreadOffer


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Currency_TemplateEdit, w, r, pageDetail)


}

//Currency_HandlerSave is the handler used process the saving of an Currency
func Currency_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Code"))

	var item dm.Currency
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
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
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.Currency_Store(item,r)	

	http.Redirect(w, r, Currency_Redirect, http.StatusFound)
}

//Currency_HandlerNew is the handler used process the creation of an Currency
func Currency_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Currency_Page{
		Title:       CardTitle(dm.Currency_Title, core.Action_New),
		PageTitle:   PageTitle(dm.Currency_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.Code = ""
pageDetail.Name = ""
pageDetail.AmountDp = ""
pageDetail.Country = ""
pageDetail.CountryName = ""
pageDetail.IntBase = ""
pageDetail.KeydateBase = ""
pageDetail.InterestRateTolerance = ""
pageDetail.CheckPayTo = ""
pageDetail.LatinAmericanSettlement = ""
pageDetail.DefaultLayOffBookKey = ""
pageDetail.CutOffTimeCutOffTime = ""
pageDetail.CutOffTimeTimeZone = ""
pageDetail.CutOffTimeDerivedDataUTCOffset = ""
pageDetail.CutOffTimeDerivedDataHasDaylightSaving = ""
pageDetail.CutOffTimeDerivedDataDaylightStart = ""
pageDetail.CutOffTimeDerivedDataDaylightEnd = ""
pageDetail.DealerInterventionQuoteTimeout = ""
pageDetail.CutOffTimeCutOffPeriod = ""
pageDetail.StripRateFutureExchangeCode = ""
pageDetail.StripRateFutureCurrencyContractCurrencyIsoCode = ""
pageDetail.StripRateFutureCurrencyContractFutureContractCode = ""
pageDetail.OvernightFundingSpreadBid = ""
pageDetail.OvernightFundingSpreadOffer = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.Currency_TemplateNew, w, r, pageDetail)

}

//Currency_HandlerDelete is the handler used process the deletion of an Currency
func Currency_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Currency_QueryString)

	dao.Currency_Delete(searchID)	

	http.Redirect(w, r, Currency_Redirect, http.StatusFound)
}