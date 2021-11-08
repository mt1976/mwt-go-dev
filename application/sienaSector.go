package application

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

var sienaSectorSQL = "Code, 	Name"
var sqlSCTCode, sqlSCTName sql.NullString

//sienaSectorPage is cheese
type sienaSectorListPage struct {
	UserMenu         []dm.AppMenuItem
	UserRole         string
	UserNavi         string
	Title            string
	PageTitle        string
	SienaSectorCount int
	SienaSectorList  []sienaSectorItem
}

//sienaSectorPage is cheese
type sienaSectorPage struct {
	UserMenu  []dm.AppMenuItem
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Code      string
	Name      string
}

//sienaSectorItem is cheese
type sienaSectorItem struct {
	Code   string
	Name   string
	Action string
}

func Sector_Publish(mux http.ServeMux) {

	mux.HandleFunc("/listSienaSector/", ListSienaSectorHandler)
	mux.HandleFunc("/viewSienaSector/", ViewSienaSectorHandler)
	mux.HandleFunc("/editSienaSector/", EditSienaSectorHandler)
	mux.HandleFunc("/saveSienaSector/", SaveSienaSectorHandler)
	mux.HandleFunc("/newSienaSector/", NewSienaSectorHandler)
	core.LOG_mux("Siena", "Sector")
}

func ListSienaSectorHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "listSienaSector"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []sienaSectorItem
	noItems, returnList, _ := getSienaSectorList()

	pageSienaSectorList := sienaSectorListPage{
		UserMenu:         UserMenu_Get(r),
		UserRole:         core.GetUserRole(r),
		UserNavi:         "NOT USED",
		Title:            core.ApplicationProperties["appname"],
		PageTitle:        core.ApplicationProperties["appname"] + " - " + "Sectors",
		SienaSectorCount: noItems,
		SienaSectorList:  returnList,
	}

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaSectorList)

}

func ViewSienaSectorHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "viewSienaSector"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	searchID := core.GetURLparam(r, "sienaSector")
	noItems, returnRecord, _ := getSienaSector(searchID)
	fmt.Println("NoSienaCountries", noItems)

	pageSienaSectorList := sienaSectorPage{
		UserMenu:  UserMenu_Get(r),
		UserRole:  core.GetUserRole(r),
		UserNavi:  "NOT USED",
		Title:     core.ApplicationProperties["appname"],
		PageTitle: core.ApplicationProperties["appname"] + " - " + "Sectors - View",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
	}

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaSectorList)

}

func EditSienaSectorHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "editSienaSector"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	searchID := core.GetURLparam(r, "sienaSector")
	noItems, returnRecord, _ := getSienaSector(searchID)
	fmt.Println("NoSienaCountries", noItems)

	pageSienaSectorList := sienaSectorPage{
		UserMenu:  UserMenu_Get(r),
		UserRole:  core.GetUserRole(r),
		UserNavi:  "NOT USED",
		Title:     core.ApplicationProperties["appname"],
		PageTitle: core.ApplicationProperties["appname"] + " - " + "Sectors - Edit",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
	}

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaSectorList)

}

func SaveSienaSectorHandler(w http.ResponseWriter, r *http.Request) {
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

	var item sienaSectorItem

	item.Code = r.FormValue("code")
	if len(item.Code) == 0 {
		item.Code = r.FormValue("ID")
	}
	item.Name = r.FormValue("name")

	log.Println(item)

	item.Action = "UPDATE"

	fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE
	var sFldCode StaticImport_KeyField
	var sFldName StaticImport_Field
	// POPULATE THE XML FIELDS
	sFldCode.Name = "code"
	sFldCode.Text = item.Code

	sFldName.Name = "name"
	sFldName.Text = item.Name
	// IGNORE
	var sienaKeyFields []StaticImport_KeyField
	var sienaFields []StaticImport_Field
	// ADD THE FIELDS TO THE LISTS ABOVE
	sienaKeyFields = append(sienaKeyFields, sFldCode)
	sienaFields = append(sienaFields, sFldName)
	// IGNORE
	sienaRecord := &StaticImport_Record{KEYFIELD: sienaKeyFields, FIELD: sienaFields}
	fmt.Println(sienaRecord)
	var sienaRecords []StaticImport_Record
	sienaRecords = append(sienaRecords, *sienaRecord)

	var sienaTable StaticImport_Table
	sienaTable.Name = "Sector"
	sienaTable.Classname = "com.eurobase.siena.data.sector.Sector"
	sienaTable.RECORD = sienaRecords

	var sienaTransaction StaticImport_Transaction
	sienaTransaction.Type = "IMPORT"
	sienaTransaction.TABLE = sienaTable

	var StaticImport_XMLContent StaticImport_XML
	StaticImport_XMLContent.TRANSACTIONS = sienaTransaction

	preparedXML, _ := xml.Marshal(StaticImport_XMLContent)
	fmt.Println("PreparedXML", string(preparedXML))
	fmt.Println("header", xml.Header)

	staticImporterPath := core.SienaProperties["static_in"]
	fileID := uuid.New()
	//pwd, _ := os.Getwd()
	fileName := staticImporterPath + "/" + fileID.String() + ".xml"
	fmt.Println(fileName)

	xmlFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating XML file: ", err)
		return
	}
	xmlFile.WriteString(core.SienaProperties["static_xml_encoding"] + "\n")
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(StaticImport_XMLContent)
	if err != nil {
		fmt.Println("Error encoding XML to file: ", err)
		return
	}

	//err := ioutil.WriteFile(fileName, preparedXML, 0644)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	ListSienaSectorHandler(w, r)

}

func NewSienaSectorHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "newSienaSector"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageSienaSectorList := sienaSectorPage{
		UserMenu:  UserMenu_Get(r),
		UserRole:  core.GetUserRole(r),
		UserNavi:  "NOT USED",
		Title:     core.ApplicationProperties["appname"],
		PageTitle: core.ApplicationProperties["appname"] + " - " + "Sector - New",
		ID:        "NEW",
		Code:      "",
		Name:      "",
	}

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaSectorList)

}

// getSienaSectorList read all employees
func getSienaSectorList() (int, []sienaSectorItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaSector;", sienaSectorSQL, core.SienaPropertiesDB["schema"])
	count, sienaSectorList, _, _ := fetchSienaSectorData(tsql)
	return count, sienaSectorList, nil
}

// getSienaSectorList read all employees
func getSienaSector(id string) (int, sienaSectorItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaSector WHERE Code='%s';", sienaSectorSQL, core.SienaPropertiesDB["schema"], id)
	_, _, sienaSector, _ := fetchSienaSectorData(tsql)
	return 1, sienaSector, nil
}

// getSienaSectorList read all employees
func putSienaSector(updateItem sienaSectorItem) error {

	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(core.SienaPropertiesDB["schema"])
	fmt.Println(updateItem)
	return nil
}

// fetchSienaSectorData read all employees
func fetchSienaSectorData(tsql string) (int, []sienaSectorItem, sienaSectorItem, error) {

	var sienaSector sienaSectorItem
	var sienaSectorList []sienaSectorItem

	rows, err := core.SienaDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaSector, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlSCTCode, &sqlSCTName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaSector, err
		}

		sienaSector.Code = sqlSCTCode.String
		sienaSector.Name = sqlSCTName.String

		sienaSectorList = append(sienaSectorList, sienaSector)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaSectorList, sienaSector, nil
}
