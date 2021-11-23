package application
// ----------------------------------------------------------------
// Automatically generated  "/application/credentials.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Credentials (credentials)
// Endpoint 	        : Credentials (Id)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 22/11/2021 at 21:11:40
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"html/template"
	"net/http"

	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
	logs    "github.com/mt1976/mwt-go-dev/logs"
)

//credentials_PageList provides the information for the template for a list of Credentialss
type Credentials_PageList struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []dm.Credentials
}

//credentials_Page provides the information for the template for an individual Credentials
type Credentials_Page struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	Title       string
	PageTitle   string
	AppInternalID  string
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
		SYSId string
		Id string
		Username string
		Password string
		Firstname string
		Lastname string
		Knownas string
		Email string
		Issued string
		Expiry string
		RoleType string
		Brand string
		SYSCreated string
		SYSWho string
		SYSHost string
		SYSUpdated string
		SYSCreatedBy string
		SYSCreatedHost string
		SYSUpdatedBy string
		SYSUpdatedHost string
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
}

const (
	Credentials_Redirect = dm.Credentials_PathList
)

//Credentials_Publish annouces the endpoints available for this object
func Credentials_Publish(mux http.ServeMux) {
	mux.HandleFunc(dm.Credentials_PathList, Credentials_HandlerList)
	mux.HandleFunc(dm.Credentials_PathView, Credentials_HandlerView)
	mux.HandleFunc(dm.Credentials_PathEdit, Credentials_HandlerEdit)
	mux.HandleFunc(dm.Credentials_PathNew, Credentials_HandlerNew)
	mux.HandleFunc(dm.Credentials_PathSave, Credentials_HandlerSave)
	mux.HandleFunc(dm.Credentials_PathDelete, Credentials_HandlerDelete)
	logs.Publish("Application", dm.Credentials_Title)
	
	// credentials_PublishImpl should be specified in application/credentials_Impl.go
	// to provide the implementation for the special case.
	// override function should be defined as
	// credentials_PublishImpl(mux http.ServeMux) {...}
	// TODO - this is a temporary hack to get the special case working
	// Add to main.go >>> credentials_PublishImpl(mux)
	
	
}

//Credentials_HandlerList is the handler for the list page
func Credentials_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []dm.Credentials
	noItems, returnList, _ := dao.Credentials_GetList()


	pageDetail := Credentials_PageList{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.Credentials_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Credentials_TemplateList, core.GetUserRole(r)))
	t.Execute(w, pageDetail)
}

//Credentials_HandlerView is the handler used to View a page
func Credentials_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Credentials_QueryString)
	_, rD, _ := dao.Credentials_GetByID(searchID)

	pageDetail := Credentials_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Credentials_Title, core.Action_View),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:		     rD.AppInternalID,
	}

		// 
		// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Id = rD.Id
pageDetail.Username = rD.Username
pageDetail.Password = rD.Password
pageDetail.Firstname = rD.Firstname
pageDetail.Lastname = rD.Lastname
pageDetail.Knownas = rD.Knownas
pageDetail.Email = rD.Email
pageDetail.Issued = rD.Issued
pageDetail.Expiry = rD.Expiry
pageDetail.RoleType = rD.RoleType
pageDetail.Brand = rD.Brand
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSWho = rD.SYSWho
pageDetail.SYSHost = rD.SYSHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
// Automatically generated 22/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
		//


	// credentials_HandlerViewImpl should be specified in application/credentials_Impl.go
	// to provide the implementation for the special case.
	// override should return mux - override function should be defined as
	// func credentials_HandlerViewImpl(pageDetail Credentials_Page) Credentials_Page {return pageDetail}
	pageDetail = credentials_HandlerViewImpl(pageDetail)

	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END


	t, _ := template.ParseFiles(core.GetTemplateID(dm.Credentials_TemplateView, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Credentials_HandlerEdit is the handler used generate the Edit page
func Credentials_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	searchID := core.GetURLparam(r, dm.Credentials_QueryString)
	_, rD, _ := dao.Credentials_GetByID(searchID)
	
	pageDetail := Credentials_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Credentials_Title, core.Action_Edit),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:          rD.AppInternalID,
	}

		// 
		// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = rD.SYSId
pageDetail.Id = rD.Id
pageDetail.Username = rD.Username
pageDetail.Password = rD.Password
pageDetail.Firstname = rD.Firstname
pageDetail.Lastname = rD.Lastname
pageDetail.Knownas = rD.Knownas
pageDetail.Email = rD.Email
pageDetail.Issued = rD.Issued
pageDetail.Expiry = rD.Expiry
pageDetail.RoleType = rD.RoleType
pageDetail.Brand = rD.Brand
pageDetail.SYSCreated = rD.SYSCreated
pageDetail.SYSWho = rD.SYSWho
pageDetail.SYSHost = rD.SYSHost
pageDetail.SYSUpdated = rD.SYSUpdated
pageDetail.SYSCreatedBy = rD.SYSCreatedBy
pageDetail.SYSCreatedHost = rD.SYSCreatedHost
pageDetail.SYSUpdatedBy = rD.SYSUpdatedBy
pageDetail.SYSUpdatedHost = rD.SYSUpdatedHost
// Automatically generated 22/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
		//

	// credentials_HandlerEditImpl should be specified in application/credentials_Impl.go
	// to provide the implementation for the special case.
	// override should return mux - override function should be defined as
	// func credentials_HandlerEditImpl(pageDetail Credentials_Page) Credentials_Page {return pageDetail}
	pageDetail = credentials_HandlerEditImpl(pageDetail)

	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Credentials_TemplateEdit, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Credentials_HandlerSave is the handler used process the saving of an Credentials
func Credentials_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path+r.FormValue("Id"))

	var item dm.Credentials

	//item.AppInternalID = r.FormValue("AppInternalID")
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
		item.SYSId = r.FormValue(dm.Credentials_SYSId)
		item.Id = r.FormValue(dm.Credentials_Id)
		item.Username = r.FormValue(dm.Credentials_Username)
		item.Password = r.FormValue(dm.Credentials_Password)
		item.Firstname = r.FormValue(dm.Credentials_Firstname)
		item.Lastname = r.FormValue(dm.Credentials_Lastname)
		item.Knownas = r.FormValue(dm.Credentials_Knownas)
		item.Email = r.FormValue(dm.Credentials_Email)
		item.Issued = r.FormValue(dm.Credentials_Issued)
		item.Expiry = r.FormValue(dm.Credentials_Expiry)
		item.RoleType = r.FormValue(dm.Credentials_RoleType)
		item.Brand = r.FormValue(dm.Credentials_Brand)
		item.SYSCreated = r.FormValue(dm.Credentials_SYSCreated)
		item.SYSWho = r.FormValue(dm.Credentials_SYSWho)
		item.SYSHost = r.FormValue(dm.Credentials_SYSHost)
		item.SYSUpdated = r.FormValue(dm.Credentials_SYSUpdated)
		item.SYSCreatedBy = r.FormValue(dm.Credentials_SYSCreatedBy)
		item.SYSCreatedHost = r.FormValue(dm.Credentials_SYSCreatedHost)
		item.SYSUpdatedBy = r.FormValue(dm.Credentials_SYSUpdatedBy)
		item.SYSUpdatedHost = r.FormValue(dm.Credentials_SYSUpdatedHost)
	
	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END

	// credentials_HandlerSaveImpl should be specified in application/credentials_Impl.go
	// to provide the implementation for the special case.
	// override should return mux - override function should be defined as
	// func credentials_HandlerSaveImpl(item dm.Credentials) dm.Credentials {return item}
	item = credentials_HandlerSaveImpl(item)

	// Automatically generated 22/11/2021 by matttownsend on silicon.local - END

	dao.Credentials_Store(item)	

	http.Redirect(w, r, Credentials_Redirect, http.StatusFound)
}

//Credentials_HandlerNew is the handler used process the creation of an Credentials
func Credentials_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageDetail := Credentials_Page{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Credentials_Title, core.Action_New),
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		AppInternalID:   "NEW",	
	}

		// 
		// Automatically generated 22/11/2021 by matttownsend on silicon.local - START
pageDetail.SYSId = ""
pageDetail.Id = ""
pageDetail.Username = ""
pageDetail.Password = ""
pageDetail.Firstname = ""
pageDetail.Lastname = ""
pageDetail.Knownas = ""
pageDetail.Email = ""
pageDetail.Issued = ""
pageDetail.Expiry = ""
pageDetail.RoleType = ""
pageDetail.Brand = ""
pageDetail.SYSCreated = ""
pageDetail.SYSWho = ""
pageDetail.SYSHost = ""
pageDetail.SYSUpdated = ""
pageDetail.SYSCreatedBy = ""
pageDetail.SYSCreatedHost = ""
pageDetail.SYSUpdatedBy = ""
pageDetail.SYSUpdatedHost = ""
// Automatically generated 22/11/2021 by matttownsend on silicon.local - Enrichment Fields Below
// Automatically generated 22/11/2021 by matttownsend on silicon.local - END
		//

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Credentials_TemplateNew, core.GetUserRole(r)))
	t.Execute(w, pageDetail)

}

//Credentials_HandlerDelete is the handler used process the deletion of an Credentials
func Credentials_HandlerDelete(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	logs.Servicing(r.URL.Path)
	searchID := core.GetURLparam(r, dm.Credentials_QueryString)

	dao.Credentials_Delete(searchID)	

	http.Redirect(w, r, Credentials_Redirect, http.StatusFound)
}