package application

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"

	core "github.com/mt1976/mwt-go-dev/core"
	dao "github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

var sienaCentreSQL = "Code, 	Name, 	Country, 	CountryName"
var sqlCENTCode, sqlCENTName, sqlCENTCountry, sqlCENTCountryName sql.NullString

//sienaCentrePage is cheese
type sienaCentreListPage struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	UserNavi         string
	Title            string
	PageTitle        string
	SienaCentreCount int
	SienaCentreList  []sienaCentreItem
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

//sienaCentreItem is cheese
type sienaCentreItem struct {
	Code        string
	Name        string
	Country     string
	Action      string
	CountryName string
}

func ListSienaCentreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "listSienaCentre"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCentreItem
	noItems, returnList, _ := getSienaCentreList()
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaCentreList := sienaCentreListPage{
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        core.ApplicationProperties["appname"] + " - " + "Centers",
		SienaCentreCount: noItems,
		SienaCentreList:  returnList,
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
		UserNavi:         "NOT USED",
	}

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaCentreList)

}

func ViewSienaCentreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "viewSienaCentre"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)
	//fmt.Println(thisConnection.Stats().OpenConnections)
	//var returnList []sienaCentreItem
	searchID := core.GetURLparam(r, "SienaCentre")
	_, returnRecord, _ := getSienaCentre(searchID)
	//fmt.Println("NoSienaItems", noItems, searchID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	pageSienaCentreList := sienaCentrePage{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   core.ApplicationProperties["appname"] + " - " + "Center - View",
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

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaCentreList)

}

func EditSienaCentreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "editSienaCentre"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)
	searchID := core.GetURLparam(r, "SienaCentre")
	_, returnRecord, _ := getSienaCentre(searchID)
	_, countryList, _ := dao.Country_GetList()

	//fmt.Println(displayList)

	pageSienaCentreList := sienaCentrePage{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   core.ApplicationProperties["appname"] + " - " + "Center - Edit",
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

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaCentreList)

}

func SaveSienaCentreHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessageAction(inUTL, "Save", "")

	var item sienaCentreItem

	item.Code = r.FormValue("Code")
	if len(item.Code) == 0 {
		item.Code = r.FormValue("ID")
	}
	item.Name = r.FormValue("Name")
	item.Country = r.FormValue("country")

	item.Action = "UPDATE"

	//fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE
	var sFldCode StaticImport_KeyField
	var sFldName StaticImport_Field
	var sFldCountry StaticImport_Field

	// POPULATE THE XML FIELDS
	sFldCode.Name = "shortName"
	sFldCode.Text = item.Code

	sFldName.Name = "fullName"
	sFldName.Text = item.Name

	sFldCountry.Name = "country"
	sFldCountry.Text = item.Country

	// IGNORE
	var sienaKeyFields []StaticImport_KeyField
	var sienaFields []StaticImport_Field
	// ADD THE FIELDS TO THE LISTS ABOVE
	sienaKeyFields = append(sienaKeyFields, sFldCode)
	sienaFields = append(sienaFields, sFldName)
	sienaFields = append(sienaFields, sFldCountry)
	// IGNORE
	sienaRecord := &StaticImport_Record{KEYFIELD: sienaKeyFields, FIELD: sienaFields}
	var sienaRecords []StaticImport_Record
	sienaRecords = append(sienaRecords, *sienaRecord)

	var sienaTable StaticImport_Table
	sienaTable.Name = "Centre"
	sienaTable.Classname = "com.eurobase.siena.data.centres.Centre"
	sienaTable.RECORD = sienaRecords

	var sienaTransaction StaticImport_Transaction
	sienaTransaction.Type = "IMPORT"
	sienaTransaction.TABLE = sienaTable

	var StaticImport_XMLContent StaticImport_XML
	StaticImport_XMLContent.TRANSACTIONS = sienaTransaction

	preparedXML, _ := xml.Marshal(StaticImport_XMLContent)
	//fmt.Println("PreparedXML", string(preparedXML))

	staticImporterPath := core.SienaProperties["static_in"]
	fileID := uuid.New()
	pwd, _ := os.Getwd()
	fileName := pwd + staticImporterPath + "/" + fileID.String() + ".xml"
	//fmt.Println(fileName)

	err := ioutil.WriteFile(fileName, preparedXML, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	ListSienaCentreHandler(w, r)

}

func NewSienaCentreHandler(w http.ResponseWriter, r *http.Request) {
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

// getSienaCentreList read all employees
func getSienaCentreList() (int, []sienaCentreItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCentre;", sienaCentreSQL, core.SienaPropertiesDB["schema"])
	count, sienaCentreList, _, _ := fetchSienaCentreData(tsql)
	return count, sienaCentreList, nil
}

// getSienaCentreList read all employees
func getSienaCentre(id string) (int, sienaCentreItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCentre WHERE Code='%s';", sienaCentreSQL, core.SienaPropertiesDB["schema"], id)
	_, _, sienaCentre, _ := fetchSienaCentreData(tsql)
	return 1, sienaCentre, nil
}

// getSienaCentreList read all employees
func putSienaCentre(updateItem sienaCentreItem) error {

	//fmt.Println(db.Stats().OpenConnections)
	//fmt.Println(core.SienaPropertiesDB["schema"])
	fmt.Println(updateItem)
	return nil
}

// fetchSienaCentreData read all employees
func fetchSienaCentreData(tsql string) (int, []sienaCentreItem, sienaCentreItem, error) {

	var sienaCentre sienaCentreItem
	var sienaCentreList []sienaCentreItem

	rows, err := core.SienaDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaCentre, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlCENTCode, &sqlCENTName, &sqlCENTCountry, &sqlCENTCountryName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaCentre, err
		}

		sienaCentre.Code = sqlCENTCode.String
		sienaCentre.Name = sqlCENTName.String
		sienaCentre.Country = sqlCENTCountry.String
		sienaCentre.CountryName = sqlCENTCountryName.String

		sienaCentreList = append(sienaCentreList, sienaCentre)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaCentreList, sienaCentre, nil
}
