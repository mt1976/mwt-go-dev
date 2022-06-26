package application
// ----------------------------------------------------------------
// Automatically generated  "/application/owner.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Owner (owner)
// Endpoint 	        : Owner (Owner)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:30
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//owner_PageList provides the information for the template for a list of Owners
type Owner_PageList struct {
	SessionInfo      dm.SessionInfo
	UserMenu         dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Owner
}
//Owner_Redirect provides a page to return to aftern an action
const (
	
	Owner_Redirect = dm.Owner_PathList
	
)

//owner_Page provides the information for the template for an individual Owner
type Owner_Page struct {
	SessionInfo      dm.SessionInfo
	UserMenu    	 dm.AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	UserName         string
	UserName_props     dm.FieldProperties
	FullName         string
	FullName_props     dm.FieldProperties
	Type         string
	Type_props     dm.FieldProperties
	TradingEntity         string
	TradingEntity_props     dm.FieldProperties
	DefaultEnterBook         string
	DefaultEnterBook_props     dm.FieldProperties
	EmailAddress         string
	EmailAddress_props     dm.FieldProperties
	Enabled         string
	Enabled_props     dm.FieldProperties
	ExternalUserIds         string
	ExternalUserIds_props     dm.FieldProperties
	Language         string
	Language_props     dm.FieldProperties
	LocalCurrency         string
	LocalCurrency_props     dm.FieldProperties
	Role         string
	Role_props     dm.FieldProperties
	TelephoneNumber         string
	TelephoneNumber_props     dm.FieldProperties
	TokenId         string
	TokenId_props     dm.FieldProperties
	Entity         string
	Entity_props     dm.FieldProperties
	UserCode         string
	UserCode_props     dm.FieldProperties
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
}



//Owner_Publish annouces the endpoints available for this object
func Owner_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Owner_Path, Owner_Handler)
	mux.HandleFunc(dm.Owner_PathList, Owner_HandlerList)
	mux.HandleFunc(dm.Owner_PathView, Owner_HandlerView)
	//Cannot Edit via GUI
	//Cannot Create via GUI
	//Cannot Save via GUI
	//Cannot Delete via GUI
	logs.Publish("Application", dm.Owner_Title)
    core.Catalog_Add(dm.Owner_Title, dm.Owner_Path, "", dm.Owner_QueryString, "Application")
}


//Owner_HandlerList is the handler for the list page
func Owner_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Owner
	noItems, returnList, _ := dao.Owner_GetList()

	pageDetail := Owner_PageList{
		Title:            CardTitle(dm.Owner_Title, core.Action_List),
		PageTitle:        PageTitle(dm.Owner_Title, core.Action_List),
		ItemsOnPage: 	  noItems,
		ItemList:         returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         Session_GetUserRole(r),
	}
	
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	
	ExecuteTemplate(dm.Owner_TemplateList, w, r, pageDetail)

}


//Owner_HandlerView is the handler used to View a page
func Owner_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Owner_QueryString)
	_, rD, _ := dao.Owner_GetByID(searchID)

	pageDetail := Owner_Page{
		Title:       CardTitle(dm.Owner_Title, core.Action_View),
		PageTitle:   PageTitle(dm.Owner_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	pageDetail = owner_PopulatePage(rD , pageDetail) 

	ExecuteTemplate(dm.Owner_TemplateView, w, r, pageDetail)

}






// Builds/Popuplates the Owner Page 
func owner_PopulatePage(rD dm.Owner, pageDetail Owner_Page) Owner_Page {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	pageDetail.UserName = rD.UserName
	pageDetail.FullName = rD.FullName
	pageDetail.Type = rD.Type
	pageDetail.TradingEntity = rD.TradingEntity
	pageDetail.DefaultEnterBook = rD.DefaultEnterBook
	pageDetail.EmailAddress = rD.EmailAddress
	pageDetail.Enabled = rD.Enabled
	pageDetail.ExternalUserIds = rD.ExternalUserIds
	pageDetail.Language = rD.Language
	pageDetail.LocalCurrency = rD.LocalCurrency
	pageDetail.Role = rD.Role
	pageDetail.TelephoneNumber = rD.TelephoneNumber
	pageDetail.TokenId = rD.TokenId
	pageDetail.Entity = rD.Entity
	pageDetail.UserCode = rD.UserCode
	
	
	//
	// Automatically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local - Enrichment Fields Below
	//
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	pageDetail.UserName_props = rD.UserName_props
	pageDetail.FullName_props = rD.FullName_props
	pageDetail.Type_props = rD.Type_props
	pageDetail.TradingEntity_props = rD.TradingEntity_props
	pageDetail.DefaultEnterBook_props = rD.DefaultEnterBook_props
	pageDetail.EmailAddress_props = rD.EmailAddress_props
	pageDetail.Enabled_props = rD.Enabled_props
	pageDetail.ExternalUserIds_props = rD.ExternalUserIds_props
	pageDetail.Language_props = rD.Language_props
	pageDetail.LocalCurrency_props = rD.LocalCurrency_props
	pageDetail.Role_props = rD.Role_props
	pageDetail.TelephoneNumber_props = rD.TelephoneNumber_props
	pageDetail.TokenId_props = rD.TokenId_props
	pageDetail.Entity_props = rD.Entity_props
	pageDetail.UserCode_props = rD.UserCode_props
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//spew.Dump(pageDetail)
return pageDetail
}	