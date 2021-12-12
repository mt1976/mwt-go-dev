package application
// ----------------------------------------------------------------
// Automatically generated  "/application/dealtype.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DealType (dealtype)
// Endpoint 	        : DealType (DealTypeKey)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:14
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//dealtype_PageList provides the information for the template for a list of DealTypes
type DealType_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.DealType
}

//dealtype_Page provides the information for the template for an individual DealType
type DealType_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		DealTypeKey string
		DealTypeShortName string
		HostKey string
		IsActive string
		Interbook string
		BackOfficeLink string
		HasTicket string
		CurrencyOverride string
		CurrencyHolderCurrency string
		AllBooks string
		FundamentalDealTypeKey string
		RelatedDealType string
		BookName string
		ExportMethod string
		DefaultUserLayoffBooks string
		RFQ string
		OBS string
		KID string
		InternalId string
		InternalDeleted string
		UpdatedTransactionId string
		UpdatedUserId string
		UpdatedDateTime string
		DeletedTransactionId string
		DeletedUserId string
		ChangeType string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	DealType_Redirect = dm.DealType_PathList
)

//DealType_Publish annouces the endpoints available for this object
func DealType_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.DealType_PathList, DealType_HandlerList)
	mux.HandleFunc(dm.DealType_PathView, DealType_HandlerView)
	mux.HandleFunc(dm.DealType_PathEdit, DealType_HandlerEdit)
	mux.HandleFunc(dm.DealType_PathNew, DealType_HandlerNew)
	mux.HandleFunc(dm.DealType_PathSave, DealType_HandlerSave)
	mux.HandleFunc(dm.DealType_PathDelete, DealType_HandlerDelete)
	logs.Publish("Siena", dm.DealType_Title)
    //No API
}

//DealType_HandlerList is the handler for the list page
func DealType_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.DealType
	noItems, returnList, _ := dao.DealType_GetList()

	pageDetail := DealType_PageList{
		Title:            CardTitle(dm.DealType_Title, core.Action_List),
		PageTitle:        PageTitle(dm.DealType_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.DealType_TemplateList, w, r, pageDetail)

}

//DealType_HandlerView is the handler used to View a page
func DealType_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DealType_QueryString)
	_, rD, _ := dao.DealType_GetByID(searchID)

	pageDetail := DealType_Page{
		Title:       CardTitle(dm.DealType_Title, core.Action_View),
		PageTitle:   PageTitle(dm.DealType_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.DealTypeKey = rD.DealTypeKey
pageDetail.DealTypeShortName = rD.DealTypeShortName
pageDetail.HostKey = rD.HostKey
pageDetail.IsActive = rD.IsActive
pageDetail.Interbook = rD.Interbook
pageDetail.BackOfficeLink = rD.BackOfficeLink
pageDetail.HasTicket = rD.HasTicket
pageDetail.CurrencyOverride = rD.CurrencyOverride
pageDetail.CurrencyHolderCurrency = rD.CurrencyHolderCurrency
pageDetail.AllBooks = rD.AllBooks
pageDetail.FundamentalDealTypeKey = rD.FundamentalDealTypeKey
pageDetail.RelatedDealType = rD.RelatedDealType
pageDetail.BookName = rD.BookName
pageDetail.ExportMethod = rD.ExportMethod
pageDetail.DefaultUserLayoffBooks = rD.DefaultUserLayoffBooks
pageDetail.RFQ = rD.RFQ
pageDetail.OBS = rD.OBS
pageDetail.KID = rD.KID
pageDetail.InternalId = rD.InternalId
pageDetail.InternalDeleted = rD.InternalDeleted
pageDetail.UpdatedTransactionId = rD.UpdatedTransactionId
pageDetail.UpdatedUserId = rD.UpdatedUserId
pageDetail.UpdatedDateTime = rD.UpdatedDateTime
pageDetail.DeletedTransactionId = rD.DeletedTransactionId
pageDetail.DeletedUserId = rD.DeletedUserId
pageDetail.ChangeType = rD.ChangeType


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.DealType_TemplateView, w, r, pageDetail)

}

//DealType_HandlerEdit is the handler used generate the Edit page
func DealType_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DealType_QueryString)
	_, rD, _ := dao.DealType_GetByID(searchID)
	
	pageDetail := DealType_Page{
		Title:       CardTitle(dm.DealType_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.DealType_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.DealTypeKey = rD.DealTypeKey
pageDetail.DealTypeShortName = rD.DealTypeShortName
pageDetail.HostKey = rD.HostKey
pageDetail.IsActive = rD.IsActive
pageDetail.Interbook = rD.Interbook
pageDetail.BackOfficeLink = rD.BackOfficeLink
pageDetail.HasTicket = rD.HasTicket
pageDetail.CurrencyOverride = rD.CurrencyOverride
pageDetail.CurrencyHolderCurrency = rD.CurrencyHolderCurrency
pageDetail.AllBooks = rD.AllBooks
pageDetail.FundamentalDealTypeKey = rD.FundamentalDealTypeKey
pageDetail.RelatedDealType = rD.RelatedDealType
pageDetail.BookName = rD.BookName
pageDetail.ExportMethod = rD.ExportMethod
pageDetail.DefaultUserLayoffBooks = rD.DefaultUserLayoffBooks
pageDetail.RFQ = rD.RFQ
pageDetail.OBS = rD.OBS
pageDetail.KID = rD.KID
pageDetail.InternalId = rD.InternalId
pageDetail.InternalDeleted = rD.InternalDeleted
pageDetail.UpdatedTransactionId = rD.UpdatedTransactionId
pageDetail.UpdatedUserId = rD.UpdatedUserId
pageDetail.UpdatedDateTime = rD.UpdatedDateTime
pageDetail.DeletedTransactionId = rD.DeletedTransactionId
pageDetail.DeletedUserId = rD.DeletedUserId
pageDetail.ChangeType = rD.ChangeType


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.DealType_TemplateEdit, w, r, pageDetail)


}

//DealType_HandlerSave is the handler used process the saving of an DealType
func DealType_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("DealTypeKey"))

	var item dm.DealType
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.DealTypeKey = r.FormValue(dm.DealType_DealTypeKey)
		item.DealTypeShortName = r.FormValue(dm.DealType_DealTypeShortName)
		item.HostKey = r.FormValue(dm.DealType_HostKey)
		item.IsActive = r.FormValue(dm.DealType_IsActive)
		item.Interbook = r.FormValue(dm.DealType_Interbook)
		item.BackOfficeLink = r.FormValue(dm.DealType_BackOfficeLink)
		item.HasTicket = r.FormValue(dm.DealType_HasTicket)
		item.CurrencyOverride = r.FormValue(dm.DealType_CurrencyOverride)
		item.CurrencyHolderCurrency = r.FormValue(dm.DealType_CurrencyHolderCurrency)
		item.AllBooks = r.FormValue(dm.DealType_AllBooks)
		item.FundamentalDealTypeKey = r.FormValue(dm.DealType_FundamentalDealTypeKey)
		item.RelatedDealType = r.FormValue(dm.DealType_RelatedDealType)
		item.BookName = r.FormValue(dm.DealType_BookName)
		item.ExportMethod = r.FormValue(dm.DealType_ExportMethod)
		item.DefaultUserLayoffBooks = r.FormValue(dm.DealType_DefaultUserLayoffBooks)
		item.RFQ = r.FormValue(dm.DealType_RFQ)
		item.OBS = r.FormValue(dm.DealType_OBS)
		item.KID = r.FormValue(dm.DealType_KID)
		item.InternalId = r.FormValue(dm.DealType_InternalId)
		item.InternalDeleted = r.FormValue(dm.DealType_InternalDeleted)
		item.UpdatedTransactionId = r.FormValue(dm.DealType_UpdatedTransactionId)
		item.UpdatedUserId = r.FormValue(dm.DealType_UpdatedUserId)
		item.UpdatedDateTime = r.FormValue(dm.DealType_UpdatedDateTime)
		item.DeletedTransactionId = r.FormValue(dm.DealType_DeletedTransactionId)
		item.DeletedUserId = r.FormValue(dm.DealType_DeletedUserId)
		item.ChangeType = r.FormValue(dm.DealType_ChangeType)
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.DealType_Store(item,r)	

	http.Redirect(w, r, DealType_Redirect, http.StatusFound)
}

//DealType_HandlerNew is the handler used process the creation of an DealType
func DealType_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := DealType_Page{
		Title:       CardTitle(dm.DealType_Title, core.Action_New),
		PageTitle:   PageTitle(dm.DealType_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.DealTypeKey = ""
pageDetail.DealTypeShortName = ""
pageDetail.HostKey = ""
pageDetail.IsActive = ""
pageDetail.Interbook = ""
pageDetail.BackOfficeLink = ""
pageDetail.HasTicket = ""
pageDetail.CurrencyOverride = ""
pageDetail.CurrencyHolderCurrency = ""
pageDetail.AllBooks = ""
pageDetail.FundamentalDealTypeKey = ""
pageDetail.RelatedDealType = ""
pageDetail.BookName = ""
pageDetail.ExportMethod = ""
pageDetail.DefaultUserLayoffBooks = ""
pageDetail.RFQ = ""
pageDetail.OBS = ""
pageDetail.KID = ""
pageDetail.InternalId = ""
pageDetail.InternalDeleted = ""
pageDetail.UpdatedTransactionId = ""
pageDetail.UpdatedUserId = ""
pageDetail.UpdatedDateTime = ""
pageDetail.DeletedTransactionId = ""
pageDetail.DeletedUserId = ""
pageDetail.ChangeType = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.DealType_TemplateNew, w, r, pageDetail)

}

//DealType_HandlerDelete is the handler used process the deletion of an DealType
func DealType_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.DealType_QueryString)

	dao.DealType_Delete(searchID)	

	http.Redirect(w, r, DealType_Redirect, http.StatusFound)
}