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
// Date & Time		    : 14/06/2022 at 21:31:47
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//accountladder_PageList provides the information for the template for a list of AccountLadders
type AccountLadder_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.AccountLadder
}
//AccountLadder_Redirect provides a page to return to aftern an action
const (
	AccountLadder_Redirect = dm.AccountLadder_PathList
)

//accountladder_Page provides the information for the template for an individual AccountLadder
type AccountLadder_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SienaReference         string
	BusinessDate         string
	ContractNumber         string
	Balance         string
	DealtCcy         string
	AmountDp         string
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



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

	pageDetail := AccountLadder_PageList{
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

	pageDetail := AccountLadder_Page{
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
func accountladder_PopulatePage(rD dm.AccountLadder, pageDetail AccountLadder_Page) AccountLadder_Page {
	// START
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.SienaReference = rD.SienaReference
	pageDetail.BusinessDate = rD.BusinessDate
	pageDetail.ContractNumber = rD.ContractNumber
	pageDetail.Balance = rD.Balance
	pageDetail.DealtCcy = rD.DealtCcy
	pageDetail.AmountDp = rD.AmountDp
	
	
	//
	// Automatically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 14/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	