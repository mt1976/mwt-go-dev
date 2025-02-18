package application
// ----------------------------------------------------------------
// Automatically generated  "/application/accountladder.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : AccountLadder (accountladder)
// Endpoint 	        : AccountLadder (AccountNo)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:41
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)





//AccountLadder_Publish annouces the endpoints available for this object
func AccountLadder_Publish(mux http.ServeMux) {
	//No API
	mux.HandleFunc(dm.AccountLadder_PathList, AccountLadder_HandlerList)
	mux.HandleFunc(dm.AccountLadder_PathView, AccountLadder_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.AccountLadder_Title)
    //No API
}


//AccountLadder_HandlerList is the handler for the list page
func AccountLadder_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.AccountLadder
	noItems, returnList, _ := dao.AccountLadder_GetList()

	pageDetail := dm.AccountLadder_PageList{
		Title:            CardTitle(dm.AccountLadder_Title, core.Action_List),
		PageTitle:        PageTitle(dm.AccountLadder_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.AccountLadder_TemplateList, w, r, pageDetail)

}


//AccountLadder_HandlerView is the handler used to View a page
func AccountLadder_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.AccountLadder_QueryString)
	_, rD, _ := dao.AccountLadder_GetByID(searchID)

	pageDetail := dm.AccountLadder_Page{
		Title:       CardTitle(dm.AccountLadder_Title, core.Action_View),
		PageTitle:   PageTitle(dm.AccountLadder_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = accountladder_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.AccountLadder_TemplateView, w, r, pageDetail)

}






// Builds/Popuplates the AccountLadder Page 
func accountladder_PopulatePage(rD dm.AccountLadder, pageDetail dm.AccountLadder_Page) dm.AccountLadder_Page {
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SienaReference = rD.SienaReference
	pageDetail.BusinessDate = rD.BusinessDate
	pageDetail.ContractNumber = rD.ContractNumber
	pageDetail.Balance = rD.Balance
	pageDetail.DealtCcy = rD.DealtCcy
	pageDetail.AmountDp = rD.AmountDp
	
	
	//
	// Automatically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.SienaReference_props = rD.SienaReference_props
	pageDetail.BusinessDate_props = rD.BusinessDate_props
	pageDetail.ContractNumber_props = rD.ContractNumber_props
	pageDetail.Balance_props = rD.Balance_props
	pageDetail.DealtCcy_props = rD.DealtCcy_props
	pageDetail.AmountDp_props = rD.AmountDp_props
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	