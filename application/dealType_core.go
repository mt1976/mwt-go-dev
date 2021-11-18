package application
// ----------------------------------------------------------------
// Automatically generated  "/application/dealtype.go"
// ----------------------------------------------------------------
// Package            : application
// Object 			  : DealType
// Endpoint Root 	  : DealType
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

//dealtype_PageList provides the information for the template for a list of DealTypes
type dealtype_PageList struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.DealType
}

//dealtype_Page provides the information for the template for an individual DealType
type dealtype_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	AppInternalID  string
	// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
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
	
	// Automatically generated 18/11/2021 by matttownsend on silicon.local - END
}

const (
	DealType_Redirect = dm.DealType_PathList
)

//DealType_Publish annouces the endpoints available for this object
func DealType_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.DealType_PathList, DealType_HandlerList)
	mux.HandleFunc(dm.DealType_PathView, DealType_HandlerView)
	mux.HandleFunc(dm.DealType_PathEdit, DealType_HandlerEdit)
	mux.HandleFunc(dm.DealType_PathNew, DealType_HandlerNew)
	mux.HandleFunc(dm.DealType_PathSave, DealType_HandlerSave)
	mux.HandleFunc(dm.DealType_PathDelete, DealType_HandlerDelete)
	logs.Publish("Siena", dm.DealType_Title)
}

//DealType_HandlerList is the handler for the list page
func DealType_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.DealType
	noItems, returnList, _ := dao.DealType_GetList()


	pageDetail := dealtype_PageList{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.DealType_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.DealType_TemplateList, core.GetUserRole(r)))
	t.Execute(w, pageDetail)
}

//DealType_HandlerView is the handler used to View a page
func DealType_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DealType_QueryString)
	_, rD, _ := dao.DealType_GetByID(searchID)

	pageDetail := dealtype_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.DealType_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:		     rD.AppInternalID,
		// 
		// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
		DealTypeKey: rD.DealTypeKey,
		DealTypeShortName: rD.DealTypeShortName,
		HostKey: rD.HostKey,
		IsActive: rD.IsActive,
		Interbook: rD.Interbook,
		BackOfficeLink: rD.BackOfficeLink,
		HasTicket: rD.HasTicket,
		CurrencyOverride: rD.CurrencyOverride,
		CurrencyHolderCurrency: rD.CurrencyHolderCurrency,
		AllBooks: rD.AllBooks,
		FundamentalDealTypeKey: rD.FundamentalDealTypeKey,
		RelatedDealType: rD.RelatedDealType,
		BookName: rD.BookName,
		ExportMethod: rD.ExportMethod,
		DefaultUserLayoffBooks: rD.DefaultUserLayoffBooks,
		RFQ: rD.RFQ,
		OBS: rD.OBS,
		KID: rD.KID,
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

	t, _ := template.ParseFiles(core.GetTemplateID(dm.DealType_TemplateView, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//DealType_HandlerEdit is the handler used generate the Edit page
func DealType_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DealType_QueryString)
	_, rD, _ := dao.DealType_GetByID(searchID)
	
	pageDetail := dealtype_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.DealType_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:          rD.AppInternalID,
		// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
			DealTypeKey: rD.DealTypeKey,
			DealTypeShortName: rD.DealTypeShortName,
			HostKey: rD.HostKey,
			IsActive: rD.IsActive,
			Interbook: rD.Interbook,
			BackOfficeLink: rD.BackOfficeLink,
			HasTicket: rD.HasTicket,
			CurrencyOverride: rD.CurrencyOverride,
			CurrencyHolderCurrency: rD.CurrencyHolderCurrency,
			AllBooks: rD.AllBooks,
			FundamentalDealTypeKey: rD.FundamentalDealTypeKey,
			RelatedDealType: rD.RelatedDealType,
			BookName: rD.BookName,
			ExportMethod: rD.ExportMethod,
			DefaultUserLayoffBooks: rD.DefaultUserLayoffBooks,
			RFQ: rD.RFQ,
			OBS: rD.OBS,
			KID: rD.KID,
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

	t, _ := template.ParseFiles(core.GetTemplateID(dm.DealType_TemplateEdit, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//DealType_HandlerSave is the handler used process the saving of an DealType
func DealType_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ID"))

	var item dm.DealType

	//item.AppInternalID = r.FormValue("AppInternalID")
	// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
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
	
	// Automatically generated 18/11/2021 by matttownsend on silicon.local - END

	dao.DealType_Store(item)	

	http.Redirect(w, r, DealType_Redirect, http.StatusFound)
}

//DealType_HandlerNew is the handler used process the creation of an DealType
func DealType_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := dealtype_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.DealType_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:   "NEW",
		// Automatically generated 18/11/2021 by matttownsend on silicon.local - START
			DealTypeKey: "",
			DealTypeShortName: "",
			HostKey: "",
			IsActive: "True",
			Interbook: "True",
			BackOfficeLink: "True",
			HasTicket: "True",
			CurrencyOverride: "True",
			CurrencyHolderCurrency: "",
			AllBooks: "True",
			FundamentalDealTypeKey: "",
			RelatedDealType: "",
			BookName: "",
			ExportMethod: "",
			DefaultUserLayoffBooks: "True",
			RFQ: "True",
			OBS: "True",
			KID: "True",
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

	t, _ := template.ParseFiles(core.GetTemplateID(dm.DealType_TemplateNew, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//DealType_HandlerDelete is the handler used process the deletion of an DealType
func DealType_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.DealType_QueryString)

	dao.DealType_Delete(searchID)	

	http.Redirect(w, r, DealType_Redirect, http.StatusFound)
}