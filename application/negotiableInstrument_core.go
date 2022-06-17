package application
// ----------------------------------------------------------------
// Automatically generated  "/application/negotiableinstrument.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : NegotiableInstrument (negotiableinstrument)
// Endpoint 	        : NegotiableInstrument (Id)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:12
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//negotiableinstrument_PageList provides the information for the template for a list of NegotiableInstruments
type NegotiableInstrument_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.NegotiableInstrument
}
//NegotiableInstrument_Redirect provides a page to return to aftern an action
const (
	NegotiableInstrument_Redirect = dm.NegotiableInstrument_PathList
)

//negotiableinstrument_Page provides the information for the template for an individual NegotiableInstrument
type NegotiableInstrument_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	Id         string
	LongName         string
	Isin         string
	Tidm         string
	Sedol         string
	IssueDate         string
	MaturityDate         string
	CouponValue         string
	CouponType         string
	Segment         string
	Sector         string
	CodeConventionCalculateAccrual         string
	MinimumDenomination         string
	DenominationCurrency         string
	TradingCurrency         string
	Type         string
	FlatYield         string
	PaymentCouponDate         string
	PeriodOfCoupon         string
	ExCouponDate         string
	DateOfIndexInflation         string
	UnitOfQuotation         string
	SYSCreated         string
	SYSWho         string
	SYSHost         string
	SYSUpdated         string
	Issuer         string
	IssueAmount         string
	RunningYield         string
	LEI         string
	CUSIP         string
	SYSUpdatedHost         string
	SYSCreatedBy         string
	SYSCreatedHost         string
	SYSUpdatedBy         string
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//NegotiableInstrument_Publish annouces the endpoints available for this object
func NegotiableInstrument_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.NegotiableInstrument_Path, NegotiableInstrument_Handler)
	mux.HandleFunc(dm.NegotiableInstrument_PathList, NegotiableInstrument_HandlerList)
	mux.HandleFunc(dm.NegotiableInstrument_PathView, NegotiableInstrument_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	mux.HandleFunc(dm.NegotiableInstrument_PathSave, NegotiableInstrument_HandlerSave)
	//Cannot Delete via GUI
	logs.Publish("Application", dm.NegotiableInstrument_Title)
    core.Catalog_Add(dm.NegotiableInstrument_Title, dm.NegotiableInstrument_Path, "", dm.NegotiableInstrument_QueryString, "Application")
}


//NegotiableInstrument_HandlerList is the handler for the list page
func NegotiableInstrument_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.NegotiableInstrument
	noItems, returnList, _ := dao.NegotiableInstrument_GetList()

	pageDetail := NegotiableInstrument_PageList{
		Title:            CardTitle(dm.NegotiableInstrument_Title, core.Action_List),
		PageTitle:        PageTitle(dm.NegotiableInstrument_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.NegotiableInstrument_TemplateList, w, r, pageDetail)

}


//NegotiableInstrument_HandlerView is the handler used to View a page
func NegotiableInstrument_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.NegotiableInstrument_QueryString)
	_, rD, _ := dao.NegotiableInstrument_GetByID(searchID)

	pageDetail := NegotiableInstrument_Page{
		Title:       CardTitle(dm.NegotiableInstrument_Title, core.Action_View),
		PageTitle:   PageTitle(dm.NegotiableInstrument_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = negotiableinstrument_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.NegotiableInstrument_TemplateView, w, r, pageDetail)

}



//NegotiableInstrument_HandlerSave is the handler used process the saving of an NegotiableInstrument
func NegotiableInstrument_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	var item dm.NegotiableInstrument
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.SYSId = r.FormValue(dm.NegotiableInstrument_SYSId_scrn)
		item.Id = r.FormValue(dm.NegotiableInstrument_Id_scrn)
		item.LongName = r.FormValue(dm.NegotiableInstrument_LongName_scrn)
		item.Isin = r.FormValue(dm.NegotiableInstrument_Isin_scrn)
		item.Tidm = r.FormValue(dm.NegotiableInstrument_Tidm_scrn)
		item.Sedol = r.FormValue(dm.NegotiableInstrument_Sedol_scrn)
		item.IssueDate = r.FormValue(dm.NegotiableInstrument_IssueDate_scrn)
		item.MaturityDate = r.FormValue(dm.NegotiableInstrument_MaturityDate_scrn)
		item.CouponValue = r.FormValue(dm.NegotiableInstrument_CouponValue_scrn)
		item.CouponType = r.FormValue(dm.NegotiableInstrument_CouponType_scrn)
		item.Segment = r.FormValue(dm.NegotiableInstrument_Segment_scrn)
		item.Sector = r.FormValue(dm.NegotiableInstrument_Sector_scrn)
		item.CodeConventionCalculateAccrual = r.FormValue(dm.NegotiableInstrument_CodeConventionCalculateAccrual_scrn)
		item.MinimumDenomination = r.FormValue(dm.NegotiableInstrument_MinimumDenomination_scrn)
		item.DenominationCurrency = r.FormValue(dm.NegotiableInstrument_DenominationCurrency_scrn)
		item.TradingCurrency = r.FormValue(dm.NegotiableInstrument_TradingCurrency_scrn)
		item.Type = r.FormValue(dm.NegotiableInstrument_Type_scrn)
		item.FlatYield = r.FormValue(dm.NegotiableInstrument_FlatYield_scrn)
		item.PaymentCouponDate = r.FormValue(dm.NegotiableInstrument_PaymentCouponDate_scrn)
		item.PeriodOfCoupon = r.FormValue(dm.NegotiableInstrument_PeriodOfCoupon_scrn)
		item.ExCouponDate = r.FormValue(dm.NegotiableInstrument_ExCouponDate_scrn)
		item.DateOfIndexInflation = r.FormValue(dm.NegotiableInstrument_DateOfIndexInflation_scrn)
		item.UnitOfQuotation = r.FormValue(dm.NegotiableInstrument_UnitOfQuotation_scrn)
		item.SYSCreated = r.FormValue(dm.NegotiableInstrument_SYSCreated_scrn)
		item.SYSWho = r.FormValue(dm.NegotiableInstrument_SYSWho_scrn)
		item.SYSHost = r.FormValue(dm.NegotiableInstrument_SYSHost_scrn)
		item.SYSUpdated = r.FormValue(dm.NegotiableInstrument_SYSUpdated_scrn)
		item.Issuer = r.FormValue(dm.NegotiableInstrument_Issuer_scrn)
		item.IssueAmount = r.FormValue(dm.NegotiableInstrument_IssueAmount_scrn)
		item.RunningYield = r.FormValue(dm.NegotiableInstrument_RunningYield_scrn)
		item.LEI = r.FormValue(dm.NegotiableInstrument_LEI_scrn)
		item.CUSIP = r.FormValue(dm.NegotiableInstrument_CUSIP_scrn)
		item.SYSUpdatedHost = r.FormValue(dm.NegotiableInstrument_SYSUpdatedHost_scrn)
		item.SYSCreatedBy = r.FormValue(dm.NegotiableInstrument_SYSCreatedBy_scrn)
		item.SYSCreatedHost = r.FormValue(dm.NegotiableInstrument_SYSCreatedHost_scrn)
		item.SYSUpdatedBy = r.FormValue(dm.NegotiableInstrument_SYSUpdatedBy_scrn)
	
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.NegotiableInstrument_Store(item,r)	
	http.Redirect(w, r, NegotiableInstrument_Redirect, http.StatusFound)
}




// Builds/Popuplates the NegotiableInstrument Page 
func negotiableinstrument_PopulatePage(rD dm.NegotiableInstrument, pageDetail NegotiableInstrument_Page) NegotiableInstrument_Page {
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SYSId = rD.SYSId
	pageDetail.Id = rD.Id
	pageDetail.LongName = rD.LongName
	pageDetail.Isin = rD.Isin
	pageDetail.Tidm = rD.Tidm
	pageDetail.Sedol = rD.Sedol
	pageDetail.IssueDate = rD.IssueDate
	pageDetail.MaturityDate = rD.MaturityDate
	pageDetail.CouponValue = rD.CouponValue
	pageDetail.CouponType = rD.CouponType
	pageDetail.Segment = rD.Segment
	pageDetail.Sector = rD.Sector
	pageDetail.CodeConventionCalculateAccrual = rD.CodeConventionCalculateAccrual
	pageDetail.MinimumDenomination = rD.MinimumDenomination
	pageDetail.DenominationCurrency = rD.DenominationCurrency
	pageDetail.TradingCurrency = rD.TradingCurrency
	pageDetail.Type = rD.Type
	pageDetail.FlatYield = rD.FlatYield
	pageDetail.PaymentCouponDate = rD.PaymentCouponDate
	pageDetail.PeriodOfCoupon = rD.PeriodOfCoupon
	pageDetail.ExCouponDate = rD.ExCouponDate
	pageDetail.DateOfIndexInflation = rD.DateOfIndexInflation
	pageDetail.UnitOfQuotation = rD.UnitOfQuotation
	pageDetail.SYSCreated = rD.SYSCreated
	pageDetail.SYSWho = rD.SYSWho
	pageDetail.SYSHost = rD.SYSHost
	pageDetail.SYSUpdated = rD.SYSUpdated
	pageDetail.Issuer = rD.Issuer
	pageDetail.IssueAmount = rD.IssueAmount
	pageDetail.RunningYield = rD.RunningYield
	pageDetail.LEI = rD.LEI
	pageDetail.CUSIP = rD.CUSIP
	pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
	pageDetail.SYSCreatedBy = rD.SYSCreatedBy
	pageDetail.SYSCreatedHost = rD.SYSCreatedHost
	pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
	
	
	//
	// Automatically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	