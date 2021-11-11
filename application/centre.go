package application

import (
	"html/template"
	"net/http"

	adaptor "github.com/mt1976/mwt-go-dev/adaptor"
	core "github.com/mt1976/mwt-go-dev/core"
	dao "github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

//sienaCentrePage is cheese
type sienaCentreListPage struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	UserNavi         string
	Title            string
	PageTitle        string
	SienaCentreCount int
	SienaCentreList  []dm.Centre
}

//sienaCentrePage is cheese
type sienaCentrePage struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	UserNavi    string
	Title       string
	PageTitle   string
	ID          string
	Code        string
	Name        string
	Country     string
	CountryName string
	Action      string
	CountryList []dm.Country
}

const (
	Centre_Redirect = "/listSienaCentre/"
)

func Centre_Publish(mux http.ServeMux) {
	mux.HandleFunc("/listSienaCentre/", Centre_HandlerList)
	mux.HandleFunc("/viewSienaCentre/", Centre_HandlerView)
	mux.HandleFunc("/editSienaCentre/", Centre_HandlerEdit)
	mux.HandleFunc("/newSienaCentre/", Centre_HandlerNew)
	mux.HandleFunc("/saveSienaCentre/", Centre_HandlerSave)
	logs.Publish("Siena", dm.Centre_Title)
}

func Centre_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	//tmpl := "listSienaCentre"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []dm.Centre
	noItems, returnList, _ := dao.Centre_GetList()
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaCentreList := sienaCentreListPage{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        PageTitle(dm.Centre_Title, core.Action_List),
		SienaCentreCount: noItems,
		SienaCentreList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
		UserNavi:         "NOT USED",
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Centre_TemplateList, core.GetUserRole(r)))
	t.Execute(w, pageSienaCentreList)

}

func Centre_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	//tmpl := "viewSienaCentre"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)
	//fmt.Println(thisConnection.Stats().OpenConnections)
	//var returnList []dm.Centre
	searchID := core.GetURLparam(r, dm.Centre_QueryString)
	_, returnRecord, _ := dao.Centre_GetByID(searchID)
	//fmt.Println("NoSienaItems", noItems, searchID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	pageSienaCentreList := sienaCentrePage{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Centre_Title, core.Action_View),
		ID:          returnRecord.Code,
		Code:        returnRecord.Code,
		Name:        returnRecord.Name,
		Country:     returnRecord.Country,
		CountryName: returnRecord.CountryName,
		Action:      "",
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		UserNavi:    "NOT USED",
	}

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Centre_TemplateView, core.GetUserRole(r)))
	t.Execute(w, pageSienaCentreList)

}

func Centre_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	//tmpl := "editSienaCentre"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)
	searchID := core.GetURLparam(r, dm.Centre_QueryString)
	_, returnRecord, _ := dao.Centre_GetByID(searchID)
	_, countryList, _ := dao.Country_GetList()

	//fmt.Println(displayList)

	pageSienaCentreList := sienaCentrePage{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.Centre_Title, core.Action_Edit),
		ID:          returnRecord.Code,
		Code:        returnRecord.Code,
		Name:        returnRecord.Name,
		Country:     returnRecord.Country,
		CountryName: returnRecord.CountryName,
		Action:      "",
		CountryList: countryList,
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		UserNavi:    "NOT USED",
	}
	//fmt.Println(pageSienaCentreList)

	t, _ := template.ParseFiles(core.GetTemplateID(dm.Centre_TemplateEdit, core.GetUserRole(r)))
	t.Execute(w, pageSienaCentreList)

}

func Centre_HandlerSave(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessageAction(inUTL, "Save", "")

	var item dm.Centre

	item.Code = r.FormValue("Code")
	if len(item.Code) == 0 {
		item.Code = r.FormValue("ID")
	}
	item.Name = r.FormValue("Name")
	item.Country = r.FormValue("country")

	adaptor.Centre_XMLExport(item)

	http.Redirect(w, r, Centre_Redirect, http.StatusFound)

}

func Centre_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "newSienaCentre"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	//Get Country List & Populate and Array of dm.Country Items

	_, countryList, _ := dao.Country_GetList()

	pageSienaCentreList := sienaCentrePage{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   core.ApplicationProperties["appname"] + " - " + "Center - New",
		ID:          "NEW",
		Code:        "",
		Name:        "",
		Country:     "",
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		UserNavi:    "NOT USED",
		Action:      "NEW",
		CountryList: countryList,
	}

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaCentreList)

}
