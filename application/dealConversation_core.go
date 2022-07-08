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
// Date & Time		    : 28/06/2022 at 16:10:50
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
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
	logs.Publish("Application", dm.DealConversation_Title)
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

	pageDetail := dm.DealConversation_PageList{
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

	pageDetail := dm.DealConversation_Page{
		Title:       CardTitle(dm.DealConversation_Title, core.Action_View),
		PageTitle:   PageTitle(dm.DealConversation_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = dealconversation_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.DealConversation_TemplateView, w, r, pageDetail)

}






// Builds/Popuplates the DealConversation Page 
func dealconversation_PopulatePage(rD dm.DealConversation, pageDetail dm.DealConversation_Page) dm.DealConversation_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
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
	
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SienaReference_props = rD.SienaReference_props
	pageDetail.Status_props = rD.Status_props
	pageDetail.MessageType_props = rD.MessageType_props
	pageDetail.ContractNumber_props = rD.ContractNumber_props
	pageDetail.AckReference_props = rD.AckReference_props
	pageDetail.NewTX_props = rD.NewTX_props
	pageDetail.LegNo_props = rD.LegNo_props
	pageDetail.Summary_props = rD.Summary_props
	pageDetail.BusinessDate_props = rD.BusinessDate_props
	pageDetail.TXNo_props = rD.TXNo_props
	pageDetail.ExternalSystem_props = rD.ExternalSystem_props
	pageDetail.MessageLogReference_props = rD.MessageLogReference_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	