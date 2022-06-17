package application
// ----------------------------------------------------------------
// Automatically generated  "/application/customeronboarding.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : CustomerOnboarding (customeronboarding)
// Endpoint 	        : CustomerOnboarding (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 17/06/2022 at 18:38:03
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//customeronboarding_PageList provides the information for the template for a list of CustomerOnboardings
type CustomerOnboarding_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.CustomerOnboarding
}
//CustomerOnboarding_Redirect provides a page to return to aftern an action
const (
	CustomerOnboarding_Redirect = dm.CustomerOnboarding_PathList
)

//customeronboarding_Page provides the information for the template for an individual CustomerOnboarding
type CustomerOnboarding_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	ID         string
	CustomerName         string
	CustomerAddress         string
	CustomerBOLID         string
	CustomerFirmName         string
	CustomerType         string
	CustomerRDC         string
	CustomerSortCode         string
	CustomerGMClientNo         string
	CustomerDefaultBook         string
	CustomerRegion         string
	CustomerCategory         string
	CustomerTelephoneNo         string
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//CustomerOnboarding_Publish annouces the endpoints available for this object
func CustomerOnboarding_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.CustomerOnboarding_Path, CustomerOnboarding_Handler)
	mux.HandleFunc(dm.CustomerOnboarding_PathList, CustomerOnboarding_HandlerList)
	mux.HandleFunc(dm.CustomerOnboarding_PathView, CustomerOnboarding_HandlerView)
	mux.HandleFunc(dm.CustomerOnboarding_PathEdit, CustomerOnboarding_HandlerEdit)
	mux.HandleFunc(dm.CustomerOnboarding_PathNew, CustomerOnboarding_HandlerNew)
	mux.HandleFunc(dm.CustomerOnboarding_PathSave, CustomerOnboarding_HandlerSave)
	mux.HandleFunc(dm.CustomerOnboarding_PathDelete, CustomerOnboarding_HandlerDelete)
	logs.Publish("Application", dm.CustomerOnboarding_Title)
    core.Catalog_Add(dm.CustomerOnboarding_Title, dm.CustomerOnboarding_Path, "", dm.CustomerOnboarding_QueryString, "Application")
}


//CustomerOnboarding_HandlerList is the handler for the list page
func CustomerOnboarding_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.CustomerOnboarding
	noItems, returnList, _ := dao.CustomerOnboarding_GetList()

	pageDetail := CustomerOnboarding_PageList{
		Title:            CardTitle(dm.CustomerOnboarding_Title, core.Action_List),
		PageTitle:        PageTitle(dm.CustomerOnboarding_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.CustomerOnboarding_TemplateList, w, r, pageDetail)

}


//CustomerOnboarding_HandlerView is the handler used to View a page
func CustomerOnboarding_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CustomerOnboarding_QueryString)
	_, rD, _ := dao.CustomerOnboarding_GetByID(searchID)

	pageDetail := CustomerOnboarding_Page{
		Title:       CardTitle(dm.CustomerOnboarding_Title, core.Action_View),
		PageTitle:   PageTitle(dm.CustomerOnboarding_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = customeronboarding_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.CustomerOnboarding_TemplateView, w, r, pageDetail)

}


//CustomerOnboarding_HandlerEdit is the handler used generate the Edit page
func CustomerOnboarding_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.CustomerOnboarding_QueryString)
	_, rD, _ := dao.CustomerOnboarding_GetByID(searchID)
	
	pageDetail := CustomerOnboarding_Page{
		Title:       CardTitle(dm.CustomerOnboarding_Title, core.Action_Edit),
		PageTitle:   PageTitle(dm.CustomerOnboarding_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = customeronboarding_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.CustomerOnboarding_TemplateEdit, w, r, pageDetail)
}


//CustomerOnboarding_HandlerSave is the handler used process the saving of an CustomerOnboarding
func CustomerOnboarding_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("ID"))

	var item dm.CustomerOnboarding
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
		item.ID = r.FormValue(dm.CustomerOnboarding_ID_scrn)
		item.CustomerName = r.FormValue(dm.CustomerOnboarding_CustomerName_scrn)
		item.CustomerAddress = r.FormValue(dm.CustomerOnboarding_CustomerAddress_scrn)
		item.CustomerBOLID = r.FormValue(dm.CustomerOnboarding_CustomerBOLID_scrn)
		item.CustomerFirmName = r.FormValue(dm.CustomerOnboarding_CustomerFirmName_scrn)
		item.CustomerType = r.FormValue(dm.CustomerOnboarding_CustomerType_scrn)
		item.CustomerRDC = r.FormValue(dm.CustomerOnboarding_CustomerRDC_scrn)
		item.CustomerSortCode = r.FormValue(dm.CustomerOnboarding_CustomerSortCode_scrn)
		item.CustomerGMClientNo = r.FormValue(dm.CustomerOnboarding_CustomerGMClientNo_scrn)
		item.CustomerDefaultBook = r.FormValue(dm.CustomerOnboarding_CustomerDefaultBook_scrn)
		item.CustomerRegion = r.FormValue(dm.CustomerOnboarding_CustomerRegion_scrn)
		item.CustomerCategory = r.FormValue(dm.CustomerOnboarding_CustomerCategory_scrn)
		item.CustomerTelephoneNo = r.FormValue(dm.CustomerOnboarding_CustomerTelephoneNo_scrn)
	
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	dao.CustomerOnboarding_Store(item,r)	
	http.Redirect(w, r, CustomerOnboarding_Redirect, http.StatusFound)
}


//CustomerOnboarding_HandlerNew is the handler used process the creation of an CustomerOnboarding
func CustomerOnboarding_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := CustomerOnboarding_Page{
		Title:       CardTitle(dm.CustomerOnboarding_Title, core.Action_New),
		PageTitle:   PageTitle(dm.CustomerOnboarding_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = customeronboarding_PopulatePage(dm.CustomerOnboarding{} , pageDetail) 

	ExecuteTemplate(dm.CustomerOnboarding_TemplateNew, w, r, pageDetail)

}	


//CustomerOnboarding_HandlerDelete is the handler used process the deletion of an CustomerOnboarding
func CustomerOnboarding_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.CustomerOnboarding_QueryString)

	dao.CustomerOnboarding_Delete(searchID)	

	http.Redirect(w, r, CustomerOnboarding_Redirect, http.StatusFound)
}


// Builds/Popuplates the CustomerOnboarding Page 
func customeronboarding_PopulatePage(rD dm.CustomerOnboarding, pageDetail CustomerOnboarding_Page) CustomerOnboarding_Page {
	// START
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.ID = rD.ID
	pageDetail.CustomerName = rD.CustomerName
	pageDetail.CustomerAddress = rD.CustomerAddress
	pageDetail.CustomerBOLID = rD.CustomerBOLID
	pageDetail.CustomerFirmName = rD.CustomerFirmName
	pageDetail.CustomerType = rD.CustomerType
	pageDetail.CustomerRDC = rD.CustomerRDC
	pageDetail.CustomerSortCode = rD.CustomerSortCode
	pageDetail.CustomerGMClientNo = rD.CustomerGMClientNo
	pageDetail.CustomerDefaultBook = rD.CustomerDefaultBook
	pageDetail.CustomerRegion = rD.CustomerRegion
	pageDetail.CustomerCategory = rD.CustomerCategory
	pageDetail.CustomerTelephoneNo = rD.CustomerTelephoneNo
	
	
	//
	// Automatically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 17/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
return pageDetail
}	