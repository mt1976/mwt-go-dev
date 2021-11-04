package siena

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	application "github.com/mt1976/mwt-go-dev/application"
	core "github.com/mt1976/mwt-go-dev/core"
)

var sienaCountrySQL = "Code, 	Name, 	ShortCode, 	EU_EEA"
var sqlCNTRCode, sqlCNTRName, sqlCNTRShortCode, sqlCNTREU_EEA sql.NullString

//sienaCountryPage is cheese
type sienaCountryListPage struct {
	UserMenu          []application.AppMenuItem
	UserRole          string
	UserNavi          string
	Title             string
	PageTitle         string
	SienaCountryCount int
	SienaCountryList  []sienaCountryItem
}

//sienaCountryPage is cheese
type sienaCountryPage struct {
	UserMenu  []application.AppMenuItem
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Code      string
	Name      string
	ShortCode string
	EU_EEA    string
	YNList    []sienaYNItem
}

//sienaCountryItem is cheese
type sienaCountryItem struct {
	Code      string
	Name      string
	ShortCode string
	EU_EEA    string
	Action    string
}

func Country_MUX(mux http.ServeMux) {
	core.LOG_success("Muxed Siena Country")
	mux.HandleFunc("/listSienaCountry/", ListSienaCountryHandler)
	mux.HandleFunc("/viewSienaCountry/", ViewSienaCountryHandler)
	mux.HandleFunc("/editSienaCountry/", EditSienaCountryHandler)
	mux.HandleFunc("/saveSienaCountry/", SaveSienaCountryHandler)
	mux.HandleFunc("/newSienaCountry/", NewSienaCountryHandler)
}

func ListSienaCountryHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "listSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessage(inUTL)

	var returnList []sienaCountryItem
	noItems, returnList, _ := getSienaCountryList()

	pageSienaCountryList := sienaCountryListPage{
		UserMenu:          application.GetUserMenu(r),
		UserRole:          application.GetUserRole(r),
		UserNavi:          "NOT USED",
		Title:             core.ApplicationProperties["appname"],
		PageTitle:         core.ApplicationProperties["appname"] + " - " + "Countries",
		SienaCountryCount: noItems,
		SienaCountryList:  returnList,
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaCountryList)

}

func ViewSienaCountryHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "viewSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessage(inUTL)

	searchID := application.GetURLparam(r, "sienaCountry")
	_, returnRecord, _ := getSienaCountry(searchID)
	_, ynList, _ := getSienaYNList()

	///	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaCountryList := sienaCountryPage{
		UserMenu:  application.GetUserMenu(r),
		UserRole:  application.GetUserRole(r),
		UserNavi:  "NOT USED",
		Title:     core.ApplicationProperties["appname"],
		PageTitle: core.ApplicationProperties["appname"] + " - " + "Country - View",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
		ShortCode: returnRecord.ShortCode,
		EU_EEA:    returnRecord.EU_EEA,
		YNList:    ynList,
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaCountryList)

}

func EditSienaCountryHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "editSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessage(inUTL)

	searchID := application.GetURLparam(r, "sienaCountry")
	_, returnRecord, _ := getSienaCountry(searchID)
	_, ynList, _ := getSienaYNList()

	pageSienaCountryList := sienaCountryPage{
		UserMenu:  application.GetUserMenu(r),
		UserRole:  application.GetUserRole(r),
		UserNavi:  "NOT USED",
		Title:     core.ApplicationProperties["appname"],
		PageTitle: core.ApplicationProperties["appname"] + " - " + "Country - Edit",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
		ShortCode: returnRecord.ShortCode,
		EU_EEA:    returnRecord.EU_EEA,
		YNList:    ynList,
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaCountryList)

}

func SaveSienaCountryHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessageAction(inUTL, "Save", "")

	var item sienaCountryItem

	item.Code = r.FormValue("code")
	if len(item.Code) == 0 {
		item.Code = r.FormValue("ID")
	}
	item.ShortCode = r.FormValue("shortcode")
	item.EU_EEA = r.FormValue("isEUEEA")
	item.Name = r.FormValue("name")
	item.Action = "UPDATE"

	fmt.Println("ITEM", item)

	var sFldCode sienaKEYFIELD
	var sFldName sienaFIELD
	var sFldShortCode sienaFIELD
	var sFldEUEEA sienaFIELD

	sFldCode.Name = "Code"
	sFldCode.Text = item.Code

	sFldName.Name = "name"
	sFldName.Text = item.Name

	sFldShortCode.Name = "shortCode"
	sFldShortCode.Text = item.ShortCode

	sFldEUEEA.Name = "EU_EEA"
	sFldEUEEA.Text = item.EU_EEA

	var sienaKeyFields []sienaKEYFIELD
	var sienaFields []sienaFIELD

	sienaKeyFields = append(sienaKeyFields, sFldCode)
	sienaFields = append(sienaFields, sFldName)
	sienaFields = append(sienaFields, sFldShortCode)
	sienaFields = append(sienaFields, sFldEUEEA)

	sienaRecord := &sienaRECORD{KEYFIELD: sienaKeyFields, FIELD: sienaFields}
	var sienaRecords []sienaRECORD
	sienaRecords = append(sienaRecords, *sienaRecord)

	var sienaTable sienaTABLE
	sienaTable.Name = "Country"
	sienaTable.Classname = "com.eurobase.siena.data.country.Country"
	sienaTable.RECORD = sienaRecords

	var sienaTransaction sienaTRANSACTION
	sienaTransaction.Type = "IMPORT"
	sienaTransaction.TABLE = sienaTable

	var sienaXMLContent sienaXML
	sienaXMLContent.TRANSACTIONS = sienaTransaction

	thisError := sienaDispatchStaticDataXML(sienaXMLContent)
	if thisError != nil {
		log.Println("Error in XML dispatch: ", thisError)
	}
	ListSienaCountryHandler(w, r)
}

func NewSienaCountryHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "newSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessage(inUTL)

	_, ynList, _ := getSienaYNList()

	pageSienaCountryList := sienaCountryPage{
		UserMenu:  application.GetUserMenu(r),
		UserRole:  application.GetUserRole(r),
		UserNavi:  "NOT USED",
		Title:     core.ApplicationProperties["appname"],
		PageTitle: core.ApplicationProperties["appname"] + " - " + "Country - New",
		ID:        "NEW",
		Code:      "",
		Name:      "",
		ShortCode: "",
		EU_EEA:    "",
		YNList:    ynList,
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaCountryList)

}

// getSienaCountryList read all employees
func getSienaCountryList() (int, []sienaCountryItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCountry;", sienaCountrySQL, core.SienaPropertiesDB["schema"])
	count, sienaCountryList, _, _ := fetchSienaCountryData(tsql)
	return count, sienaCountryList, nil
}

// getSienaCountryList read all employees
func getSienaCountry(id string) (int, sienaCountryItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCountry WHERE Code='%s';", sienaCountrySQL, core.SienaPropertiesDB["schema"], id)
	_, _, sienaCountry, _ := fetchSienaCountryData(tsql)
	return 1, sienaCountry, nil
}

// getSienaCountryList read all employees
func putSienaCountry(updateItem sienaCountryItem) error {

	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(core.SienaPropertiesDB["schema"])
	fmt.Println(updateItem)

	return nil
}

// fetchSienaCountryData read all employees
func fetchSienaCountryData(tsql string) (int, []sienaCountryItem, sienaCountryItem, error) {

	var sienaCountry sienaCountryItem
	var sienaCountryList []sienaCountryItem

	rows, err := core.SienaDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaCountry, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlCNTRCode, &sqlCNTRName, &sqlCNTRShortCode, &sqlCNTREU_EEA)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaCountry, err
		}

		sienaCountry.Code = sqlCNTRCode.String
		sienaCountry.Name = sqlCNTRName.String
		sienaCountry.ShortCode = sqlCNTRShortCode.String
		sienaCountry.EU_EEA = sienaYN(sqlCNTREU_EEA.String)

		sienaCountryList = append(sienaCountryList, sienaCountry)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaCountryList, sienaCountry, nil
}
