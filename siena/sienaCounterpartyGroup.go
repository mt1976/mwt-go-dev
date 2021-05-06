package siena

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
	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

var sienaCounterpartyGroupSQL = "Name, 	CountryCode, 	SuperGroup, 	CountryName"
var sqlCGRPName, sqlCGRPCountryCode, sqlCGRPSuperGroup, sqlCGRPCountryName sql.NullString

//sienaCounterpartyGroupPage is cheese
type sienaCounterpartyGroupListPage struct {
	UserMenu                    []application.AppMenuItem
	UserRole                    string
	UserNavi                    string
	Title                       string
	PageTitle                   string
	SienaCounterpartyGroupCount int
	SienaCounterpartyGroupList  []sienaCounterpartyGroupItem
}

//sienaCounterpartyGroupPage is cheese
type sienaCounterpartyGroupPage struct {
	UserMenu    []application.AppMenuItem
	UserRole    string
	UserNavi    string
	Title       string
	PageTitle   string
	Action      string
	Name        string
	CountryCode string
	SuperGroup  string
	CountryName string
	GroupList   []sienaCounterpartyGroupItem
	CountryList []sienaCountryItem
}

//sienaCounterpartyGroupItem is cheese
type sienaCounterpartyGroupItem struct {
	Name        string
	CountryCode string
	SuperGroup  string
	CountryName string
	Action      string
}

func ListSienaCounterpartyGroupHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "listSienaCounterpartyGroup"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessage(inUTL)

	var returnList []sienaCounterpartyGroupItem
	noItems, returnList, _ := getSienaCounterpartyGroupList()

	pageSienaCounterpartyGroupList := sienaCounterpartyGroupListPage{
		Title:                       globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - “ +                   "List Siena CounterpartyGroups",
		SienaCounterpartyGroupCount: noItems,
		SienaCounterpartyGroupList:  returnList,
		UserMenu:                    application.GetUserMenu(r),
		UserRole:                    application.GetUserRole(r),
		UserNavi:                    "NOT USED",
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaCounterpartyGroupList)

}

func ViewSienaCounterpartyGroupHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "viewSienaCounterpartyGroup"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessage(inUTL)

	searchID := application.GetURLparam(r, "SienaCounterpartyGroup")
	_, returnRecord, _ := getSienaCounterpartyGroup(searchID)

	pageSienaCounterpartyGroupList := sienaCounterpartyGroupPage{
		Title:       globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - “ +   "View Siena CounterpartyGroup",
		Name:        returnRecord.Name,
		CountryCode: returnRecord.CountryCode,
		SuperGroup:  returnRecord.SuperGroup,
		CountryName: returnRecord.CountryName,
		Action:      "",
		UserMenu:    application.GetUserMenu(r),
		UserRole:    application.GetUserRole(r),
		UserNavi:    "NOT USED",
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaCounterpartyGroupList)

}

func EditSienaCounterpartyGroupHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "editSienaCounterpartyGroup"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessage(inUTL)

	searchID := application.GetURLparam(r, "SienaCounterpartyGroup")
	_, returnRecord, _ := getSienaCounterpartyGroup(searchID)
	_, countryList, _ := getSienaCountryList()
	_, groupList, _ := getSienaCounterpartyGroupList()

	pageSienaCounterpartyGroupList := sienaCounterpartyGroupPage{
		Title:       globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - “ +   "View Siena CounterpartyGroup",
		Name:        returnRecord.Name,
		CountryCode: returnRecord.CountryCode,
		SuperGroup:  returnRecord.SuperGroup,
		CountryName: returnRecord.CountryName,
		Action:      "",
		CountryList: countryList,
		UserMenu:    application.GetUserMenu(r),
		UserRole:    application.GetUserRole(r),
		UserNavi:    "NOT USED",
		GroupList:   groupList,
	}
	//fmt.Println(pageSienaCounterpartyGroupList)

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaCounterpartyGroupList)

}

func SaveSienaCounterpartyGroupHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessageAction(inUTL,"Save","")

	var item sienaCounterpartyGroupItem

	item.Name = r.FormValue("Code")
	if len(item.Name) == 0 {
		item.Name = r.FormValue("ID")
	}
	item.Name = r.FormValue("Name")
	item.CountryCode = r.FormValue("country")

	item.Action = "UPDATE"

	fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE
	var sFldCode sienaKEYFIELD
	var sFldName sienaFIELD
	var sFldCountry sienaFIELD

	// POPULATE THE XML FIELDS
	sFldCode.Name = "shortName"
	sFldCode.Text = item.Name

	sFldName.Name = "fullName"
	sFldName.Text = item.Name

	sFldCountry.Name = "country"
	sFldCountry.Text = item.CountryCode

	// IGNORE
	var sienaKeyFields []sienaKEYFIELD
	var sienaFields []sienaFIELD
	// ADD THE FIELDS TO THE LISTS ABOVE
	sienaKeyFields = append(sienaKeyFields, sFldCode)
	sienaFields = append(sienaFields, sFldName)
	sienaFields = append(sienaFields, sFldCountry)
	// IGNORE
	sienaRecord := &sienaRECORD{KEYFIELD: sienaKeyFields, FIELD: sienaFields}
	var sienaRecords []sienaRECORD
	sienaRecords = append(sienaRecords, *sienaRecord)

	var sienaTable sienaTABLE
	sienaTable.Name = "CounterpartyGroup"
	sienaTable.Classname = "com.eurobase.siena.data.CounterpartyGroups.CounterpartyGroup"
	sienaTable.RECORD = sienaRecords

	var sienaTransaction sienaTRANSACTION
	sienaTransaction.Type = "IMPORT"
	sienaTransaction.TABLE = sienaTable

	var sienaXMLContent sienaXML
	sienaXMLContent.TRANSACTIONS = sienaTransaction

	preparedXML, _ := xml.Marshal(sienaXMLContent)
	fmt.Println("PreparedXML", string(preparedXML))

	staticImporterPath := globals.SienaProperties["static_in"]
	fileID := uuid.New()
	pwd, _ := os.Getwd()
	fileName := pwd + staticImporterPath + "/" + fileID.String() + ".xml"
	fmt.Println(fileName)

	err := ioutil.WriteFile(fileName, preparedXML, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	ListSienaCounterpartyGroupHandler(w, r)

}

func NewSienaCounterpartyGroupHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "newSienaCounterpartyGroup"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessage(inUTL)

	//Get Country List & Populate and Array of sienaCountryItem Items
	_, countryList, _ := getSienaCountryList()
	_, groupList, _ := getSienaCounterpartyGroupList()

	pageSienaCounterpartyGroupList := sienaCounterpartyGroupPage{
		Title:     globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - “ + "View Siena CounterpartyGroup",

		Name:        "",
		CountryCode: "",
		SuperGroup:  "",
		CountryName: "",
		UserMenu:    application.GetUserMenu(r),
		UserRole:    application.GetUserRole(r),
		UserNavi:    "NOT USED",
		Action:      "NEW",
		CountryList: countryList,
		GroupList:   groupList,
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaCounterpartyGroupList)

}

// getSienaCounterpartyGroupList read all employees
func getSienaCounterpartyGroupList() (int, []sienaCounterpartyGroupItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCounterpartyGroup;", sienaCounterpartyGroupSQL, globals.SienaPropertiesDB["schema"])
	count, sienaCounterpartyGroupList, _, _ := fetchSienaCounterpartyGroupData(tsql)
	return count, sienaCounterpartyGroupList, nil
}

// getSienaCounterpartyGroupList read all employees
func getSienaCounterpartyGroup(id string) (int, sienaCounterpartyGroupItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaCounterpartyGroup WHERE Name='%s';", sienaCounterpartyGroupSQL, globals.SienaPropertiesDB["schema"], id)
	_, _, sienaCounterpartyGroup, _ := fetchSienaCounterpartyGroupData(tsql)
	return 1, sienaCounterpartyGroup, nil
}

// getSienaCounterpartyGroupList read all employees
func putSienaCounterpartyGroup(updateItem sienaCounterpartyGroupItem) error {

	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(globals.SienaPropertiesDB["schema"])
	fmt.Println(updateItem)
	return nil
}

// fetchSienaCounterpartyGroupData read all employees
func fetchSienaCounterpartyGroupData(tsql string) (int, []sienaCounterpartyGroupItem, sienaCounterpartyGroupItem, error) {

	var sienaCounterpartyGroup sienaCounterpartyGroupItem
	var sienaCounterpartyGroupList []sienaCounterpartyGroupItem

	rows, err := globals.SienaDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaCounterpartyGroup, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlCGRPName, &sqlCGRPCountryCode, &sqlCGRPSuperGroup, &sqlCGRPCountryName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaCounterpartyGroup, err
		}

		sienaCounterpartyGroup.Name = sqlCGRPName.String
		sienaCounterpartyGroup.CountryCode = sqlCGRPCountryCode.String
		sienaCounterpartyGroup.SuperGroup = sqlCGRPSuperGroup.String
		sienaCounterpartyGroup.CountryName = sqlCGRPCountryName.String

		sienaCounterpartyGroupList = append(sienaCounterpartyGroupList, sienaCounterpartyGroup)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaCounterpartyGroupList, sienaCounterpartyGroup, nil
}
