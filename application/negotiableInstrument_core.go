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
// Date & Time		    : 28/06/2022 at 16:10:53
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)





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

	pageDetail := dm.NegotiableInstrument_PageList{
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

	pageDetail := dm.NegotiableInstrument_Page{
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
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.NegotiableInstrument_Store(item,r)	
	http.Redirect(w, r, dm.NegotiableInstrument_Redirect, http.StatusFound)
}




// Builds/Popuplates the NegotiableInstrument Page 
func negotiableinstrument_PopulatePage(rD dm.NegotiableInstrument, pageDetail dm.NegotiableInstrument_Page) dm.NegotiableInstrument_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
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
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SYSId_props = rD.SYSId_props
	pageDetail.Id_props = rD.Id_props
	pageDetail.LongName_props = rD.LongName_props
	pageDetail.Isin_props = rD.Isin_props
	pageDetail.Tidm_props = rD.Tidm_props
	pageDetail.Sedol_props = rD.Sedol_props
	pageDetail.IssueDate_props = rD.IssueDate_props
	pageDetail.MaturityDate_props = rD.MaturityDate_props
	pageDetail.CouponValue_props = rD.CouponValue_props
	pageDetail.CouponType_props = rD.CouponType_props
	pageDetail.Segment_props = rD.Segment_props
	pageDetail.Sector_props = rD.Sector_props
	pageDetail.CodeConventionCalculateAccrual_props = rD.CodeConventionCalculateAccrual_props
	pageDetail.MinimumDenomination_props = rD.MinimumDenomination_props
	pageDetail.DenominationCurrency_props = rD.DenominationCurrency_props
	pageDetail.TradingCurrency_props = rD.TradingCurrency_props
	pageDetail.Type_props = rD.Type_props
	pageDetail.FlatYield_props = rD.FlatYield_props
	pageDetail.PaymentCouponDate_props = rD.PaymentCouponDate_props
	pageDetail.PeriodOfCoupon_props = rD.PeriodOfCoupon_props
	pageDetail.ExCouponDate_props = rD.ExCouponDate_props
	pageDetail.DateOfIndexInflation_props = rD.DateOfIndexInflation_props
	pageDetail.UnitOfQuotation_props = rD.UnitOfQuotation_props
	pageDetail.SYSCreated_props = rD.SYSCreated_props
	pageDetail.SYSWho_props = rD.SYSWho_props
	pageDetail.SYSHost_props = rD.SYSHost_props
	pageDetail.SYSUpdated_props = rD.SYSUpdated_props
	pageDetail.Issuer_props = rD.Issuer_props
	pageDetail.IssueAmount_props = rD.IssueAmount_props
	pageDetail.RunningYield_props = rD.RunningYield_props
	pageDetail.LEI_props = rD.LEI_props
	pageDetail.CUSIP_props = rD.CUSIP_props
	pageDetail.SYSUpdatedHost_props = rD.SYSUpdatedHost_props
	pageDetail.SYSCreatedBy_props = rD.SYSCreatedBy_props
	pageDetail.SYSCreatedHost_props = rD.SYSCreatedHost_props
	pageDetail.SYSUpdatedBy_props = rD.SYSUpdatedBy_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	