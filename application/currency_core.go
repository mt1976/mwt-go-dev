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
// Date & Time		    : 17/06/2022 at 18:38:08
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
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
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Currency
}
//Currency_Redirect provides a page to return to aftern an action
const (
	Currency_Redirect = dm.Currency_PathList
)

//currency_Page provides the information for the template for an individual Currency
type Currency_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	Code         string
	Name         string
	AmountDp         string
	Country         string
	CountryName         string
	IntBase         string
	KeydateBase         string
	InterestRateTolerance         string
	CheckPayTo         string
	LatinAmericanSettlement         string
	DefaultLayOffBookKey         string
	CutOffTimeCutOffTime         string
	CutOffTimeTimeZone         string
	CutOffTimeDerivedDataUTCOffset         string
	CutOffTimeDerivedDataHasDaylightSaving         string
	CutOffTimeDerivedDataDaylightStart         string
	CutOffTimeDerivedDataDaylightEnd         string
	DealerInterventionQuoteTimeout         string
	CutOffTimeCutOffPeriod         string
	StripRateFutureExchangeCode         string
	StripRateFutureCurrencyContractCurrencyIsoCode         string
	StripRateFutureCurrencyContractFutureContractCode         string
	OvernightFundingSpreadBid         string
	OvernightFundingSpreadOffer         string
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//Currency_Publish annouces the endpoints available for this object
func Currency_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Currency_Path, Currency_Handler)
	mux.HandleFunc(dm.Currency_PathList, Currency_HandlerList)
	mux.HandleFunc(dm.Currency_PathView, Currency_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Currency_Title)
    core.Catalog_Add(dm.Currency_Title, dm.Currency_Path, "", dm.Currency_QueryString, "Application")
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

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = currency_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Currency_TemplateView, w, r, pageDetail)

}






// Builds/Popuplates the Currency Page 
func currency_PopulatePage(rD dm.Currency, pageDetail Currency_Page) Currency_Page {
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
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
	
	
	//
	// Automatically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	