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

//dealtype_PageList provides the information for the template for a list of DealTypes
type DealType_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.DealType
}
//DealType_Redirect provides a page to return to aftern an action
const (
	
	DealType_Redirect = dm.DealType_PathList
	
)

//dealtype_Page provides the information for the template for an individual DealType
type DealType_Page struct {
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
	DealTypeShortName         string
	DealTypeShortName_props     dm.FieldProperties
	HostKey         string
	HostKey_props     dm.FieldProperties
	IsActive         string
	IsActive_props     dm.FieldProperties
	Interbook         string
	Interbook_props     dm.FieldProperties
	BackOfficeLink         string
	BackOfficeLink_props     dm.FieldProperties
	HasTicket         string
	HasTicket_props     dm.FieldProperties
	CurrencyOverride         string
	CurrencyOverride_props     dm.FieldProperties
	CurrencyHolderCurrency         string
	CurrencyHolderCurrency_props     dm.FieldProperties
	AllBooks         string
	AllBooks_props     dm.FieldProperties
	FundamentalDealTypeKey         string
	FundamentalDealTypeKey_props     dm.FieldProperties
	RelatedDealType         string
	RelatedDealType_props     dm.FieldProperties
	BookName         string
	BookName_props     dm.FieldProperties
	ExportMethod         string
	ExportMethod_props     dm.FieldProperties
	DefaultUserLayoffBooks         string
	DefaultUserLayoffBooks_props     dm.FieldProperties
	RFQ         string
	RFQ_props     dm.FieldProperties
	OBS         string
	OBS_props     dm.FieldProperties
	KID         string
	KID_props     dm.FieldProperties
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



//DealType_Publish annouces the endpoints available for this object
func DealType_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.DealType_PathList, DealType_HandlerList)
	mux.HandleFunc(dm.DealType_PathView, DealType_HandlerView)
	mux.HandleFunc(dm.DealType_PathEdit, DealType_HandlerEdit)
	mux.HandleFunc(dm.DealType_PathNew, DealType_HandlerNew)
	mux.HandleFunc(dm.DealType_PathSave, DealType_HandlerSave)
	mux.HandleFunc(dm.DealType_PathDelete, DealType_HandlerDelete)
	logs.Publish("Application", dm.DealType_Title)
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

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = dealtype_PopulatePage(rD , pageDetail) 

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

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = dealtype_PopulatePage(rD , pageDetail) 

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
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.DealTypeKey = r.FormValue(dm.DealType_DealTypeKey_scrn)
		item.DealTypeShortName = r.FormValue(dm.DealType_DealTypeShortName_scrn)
		item.HostKey = r.FormValue(dm.DealType_HostKey_scrn)
		item.IsActive = r.FormValue(dm.DealType_IsActive_scrn)
		item.Interbook = r.FormValue(dm.DealType_Interbook_scrn)
		item.BackOfficeLink = r.FormValue(dm.DealType_BackOfficeLink_scrn)
		item.HasTicket = r.FormValue(dm.DealType_HasTicket_scrn)
		item.CurrencyOverride = r.FormValue(dm.DealType_CurrencyOverride_scrn)
		item.CurrencyHolderCurrency = r.FormValue(dm.DealType_CurrencyHolderCurrency_scrn)
		item.AllBooks = r.FormValue(dm.DealType_AllBooks_scrn)
		item.FundamentalDealTypeKey = r.FormValue(dm.DealType_FundamentalDealTypeKey_scrn)
		item.RelatedDealType = r.FormValue(dm.DealType_RelatedDealType_scrn)
		item.BookName = r.FormValue(dm.DealType_BookName_scrn)
		item.ExportMethod = r.FormValue(dm.DealType_ExportMethod_scrn)
		item.DefaultUserLayoffBooks = r.FormValue(dm.DealType_DefaultUserLayoffBooks_scrn)
		item.RFQ = r.FormValue(dm.DealType_RFQ_scrn)
		item.OBS = r.FormValue(dm.DealType_OBS_scrn)
		item.KID = r.FormValue(dm.DealType_KID_scrn)
		item.InternalId = r.FormValue(dm.DealType_InternalId_scrn)
		item.InternalDeleted = r.FormValue(dm.DealType_InternalDeleted_scrn)
		item.UpdatedTransactionId = r.FormValue(dm.DealType_UpdatedTransactionId_scrn)
		item.UpdatedUserId = r.FormValue(dm.DealType_UpdatedUserId_scrn)
		item.UpdatedDateTime = r.FormValue(dm.DealType_UpdatedDateTime_scrn)
		item.DeletedTransactionId = r.FormValue(dm.DealType_DeletedTransactionId_scrn)
		item.DeletedUserId = r.FormValue(dm.DealType_DeletedUserId_scrn)
		item.ChangeType = r.FormValue(dm.DealType_ChangeType_scrn)
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
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

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)
	_, _, rD, _ := dao.DealType_New()

	pageDetail := DealType_Page{
		Title:       CardTitle(dm.DealType_Title, core.Action_New),
		PageTitle:   PageTitle(dm.DealType_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = dealtype_PopulatePage(rD , pageDetail) 

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


// Builds/Popuplates the DealType Page 
func dealtype_PopulatePage(rD dm.DealType, pageDetail DealType_Page) DealType_Page {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
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
	
	
	//
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.DealTypeKey_props = rD.DealTypeKey_props
	pageDetail.DealTypeShortName_props = rD.DealTypeShortName_props
	pageDetail.HostKey_props = rD.HostKey_props
	pageDetail.IsActive_props = rD.IsActive_props
	pageDetail.Interbook_props = rD.Interbook_props
	pageDetail.BackOfficeLink_props = rD.BackOfficeLink_props
	pageDetail.HasTicket_props = rD.HasTicket_props
	pageDetail.CurrencyOverride_props = rD.CurrencyOverride_props
	pageDetail.CurrencyHolderCurrency_props = rD.CurrencyHolderCurrency_props
	pageDetail.AllBooks_props = rD.AllBooks_props
	pageDetail.FundamentalDealTypeKey_props = rD.FundamentalDealTypeKey_props
	pageDetail.RelatedDealType_props = rD.RelatedDealType_props
	pageDetail.BookName_props = rD.BookName_props
	pageDetail.ExportMethod_props = rD.ExportMethod_props
	pageDetail.DefaultUserLayoffBooks_props = rD.DefaultUserLayoffBooks_props
	pageDetail.RFQ_props = rD.RFQ_props
	pageDetail.OBS_props = rD.OBS_props
	pageDetail.KID_props = rD.KID_props
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