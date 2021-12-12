package application
// ----------------------------------------------------------------
// Automatically generated  "/application/dealconversation.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DealConversation (dealconversation)
// Endpoint 	        : DealConversation (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:13
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//dealconversation_PageList provides the information for the template for a list of DealConversations
type DealConversation_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.DealConversation
}

//dealconversation_Page provides the information for the template for an individual DealConversation
type DealConversation_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 []dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		SienaReference string
		Status string
		MessageType string
		ContractNumber string
		AckReference string
		NewTX string
		LegNo string
		Summary string
		BusinessDate string
		TXNo string
		ExternalSystem string
		MessageLogReference string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
}

const (
	DealConversation_Redirect = dm.DealConversation_PathList
)

//DealConversation_Publish annouces the endpoints available for this object
func DealConversation_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.DealConversation_PathList, DealConversation_HandlerList)
	mux.HandleFunc(dm.DealConversation_PathView, DealConversation_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Siena", dm.DealConversation_Title)
    //No API
}

//DealConversation_HandlerList is the handler for the list page
func DealConversation_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.DealConversation
	noItems, returnList, _ := dao.DealConversation_GetList()

	pageDetail := DealConversation_PageList{
		Title:            CardTitle(dm.DealConversation_Title, core.Action_List),
		PageTitle:        PageTitle(dm.DealConversation_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.DealConversation_TemplateList, w, r, pageDetail)

}

//DealConversation_HandlerView is the handler used to View a page
func DealConversation_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DealConversation_QueryString)
	_, rD, _ := dao.DealConversation_GetByID(searchID)

	pageDetail := DealConversation_Page{
		Title:       CardTitle(dm.DealConversation_Title, core.Action_View),
		PageTitle:   PageTitle(dm.DealConversation_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.SienaReference = rD.SienaReference
pageDetail.Status = rD.Status
pageDetail.MessageType = rD.MessageType
pageDetail.ContractNumber = rD.ContractNumber
pageDetail.AckReference = rD.AckReference
pageDetail.NewTX = rD.NewTX
pageDetail.LegNo = rD.LegNo
pageDetail.Summary = rD.Summary
pageDetail.BusinessDate = rD.BusinessDate
pageDetail.TXNo = rD.TXNo
pageDetail.ExternalSystem = rD.ExternalSystem
pageDetail.MessageLogReference = rD.MessageLogReference


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//


	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.DealConversation_TemplateView, w, r, pageDetail)

}

//DealConversation_HandlerEdit is the handler used generate the Edit page
func DealConversation_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.DealConversation_QueryString)
	_, rD, _ := dao.DealConversation_GetByID(searchID)
	
	pageDetail := DealConversation_Page{
		Title:       CardTitle(dm.DealConversation_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.DealConversation_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.SienaReference = rD.SienaReference
pageDetail.Status = rD.Status
pageDetail.MessageType = rD.MessageType
pageDetail.ContractNumber = rD.ContractNumber
pageDetail.AckReference = rD.AckReference
pageDetail.NewTX = rD.NewTX
pageDetail.LegNo = rD.LegNo
pageDetail.Summary = rD.Summary
pageDetail.BusinessDate = rD.BusinessDate
pageDetail.TXNo = rD.TXNo
pageDetail.ExternalSystem = rD.ExternalSystem
pageDetail.MessageLogReference = rD.MessageLogReference


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.DealConversation_TemplateEdit, w, r, pageDetail)


}

//DealConversation_HandlerSave is the handler used process the saving of an DealConversation
func DealConversation_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("MessageLogReference"))

	var item dm.DealConversation
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
		item.SienaReference = r.FormValue(dm.DealConversation_SienaReference)
		item.Status = r.FormValue(dm.DealConversation_Status)
		item.MessageType = r.FormValue(dm.DealConversation_MessageType)
		item.ContractNumber = r.FormValue(dm.DealConversation_ContractNumber)
		item.AckReference = r.FormValue(dm.DealConversation_AckReference)
		item.NewTX = r.FormValue(dm.DealConversation_NewTX)
		item.LegNo = r.FormValue(dm.DealConversation_LegNo)
		item.Summary = r.FormValue(dm.DealConversation_Summary)
		item.BusinessDate = r.FormValue(dm.DealConversation_BusinessDate)
		item.TXNo = r.FormValue(dm.DealConversation_TXNo)
		item.ExternalSystem = r.FormValue(dm.DealConversation_ExternalSystem)
		item.MessageLogReference = r.FormValue(dm.DealConversation_MessageLogReference)
	

	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END

	dao.DealConversation_Store(item,r)	

	http.Redirect(w, r, DealConversation_Redirect, http.StatusFound)
}

//DealConversation_HandlerNew is the handler used process the creation of an DealConversation
func DealConversation_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := DealConversation_Page{
		Title:       CardTitle(dm.DealConversation_Title, core.Action_New),
		PageTitle:   PageTitle(dm.DealConversation_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

		// 
		// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
pageDetail.SienaReference = ""
pageDetail.Status = ""
pageDetail.MessageType = ""
pageDetail.ContractNumber = ""
pageDetail.AckReference = ""
pageDetail.NewTX = ""
pageDetail.LegNo = ""
pageDetail.Summary = ""
pageDetail.BusinessDate = ""
pageDetail.TXNo = ""
pageDetail.ExternalSystem = ""
pageDetail.MessageLogReference = ""


// Automatically generated 12/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.DealConversation_TemplateNew, w, r, pageDetail)

}

//DealConversation_HandlerDelete is the handler used process the deletion of an DealConversation
func DealConversation_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.DealConversation_QueryString)

	dao.DealConversation_Delete(searchID)	

	http.Redirect(w, r, DealConversation_Redirect, http.StatusFound)
}