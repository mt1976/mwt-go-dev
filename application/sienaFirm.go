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
	dao "github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

var sienaFirmSQL = "FirmName, 	FullName, 	Country, 	Sector, 	SectorName, 	CountryName"
var sqlFRMFirmName, sqlFRMFullName, sqlFRMCountry, sqlFRMSector, sqlFRMSectorName, sqlFRMCountryName sql.NullString

//sienaFirmPage is cheese
type sienaFirmListPage struct {
	UserMenu       []dm.AppMenuItem
	UserRole       string
	UserNavi       string
	Title          string
	PageTitle      string
	SienaFirmCount int
	SienaFirmList  []sienaFirmItem
}

//sienaFirmPage is cheese
type sienaFirmPage struct {
	UserMenu    []dm.AppMenuItem
	UserRole    string
	UserNavi    string
	Title       string
	PageTitle   string
	ID          string
	FirmName    string
	FullName    string
	Country     string
	Sector      string
	SectorName  string
	CountryName string
	Action      string
	CountryList []dm.Country
	SectorList  []sienaSectorItem
}

//sienaFirmItem is cheese
type sienaFirmItem struct {
	FirmName    string
	FullName    string
	Country     string
	Sector      string
	SectorName  string
	CountryName string
	Action      string
}

func Firm_Publish(mux http.ServeMux) {

	mux.HandleFunc("/listSienaFirm/", ListSienaFirmHandler)
	mux.HandleFunc("/viewSienaFirm/", ViewSienaFirmHandler)
	mux.HandleFunc("/editSienaFirm/", EditSienaFirmHandler)
	mux.HandleFunc("/saveSienaFirm/", SaveSienaFirmHandler)
	mux.HandleFunc("/newSienaFirm/", NewSienaFirmHandler)
	core.LOG_mux("Siena", "Firm")
}

func ListSienaFirmHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "listSienaFirm"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	var returnList []sienaFirmItem
	noItems, returnList, _ := getSienaFirmList()

	pageSienaFirmList := sienaFirmListPage{
		UserMenu:       UserMenu_Get(r),
		UserRole:       core.GetUserRole(r),
		UserNavi:       "NOT USED",
		Title:          core.ApplicationProperties["appname"],
		PageTitle:      core.ApplicationProperties["appname"] + " - " + "Firms",
		SienaFirmCount: noItems,
		SienaFirmList:  returnList,
	}

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaFirmList)

}

func ViewSienaFirmHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "viewSienaFirm"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	searchID := core.GetURLparam(r, "SienaFirm")
	_, returnRecord, _ := getSienaFirm(searchID)

	pageSienaFirmList := sienaFirmPage{
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		UserNavi:    "NOT USED",
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   core.ApplicationProperties["appname"] + " - " + "Firm - View",
		ID:          returnRecord.FirmName,
		FirmName:    returnRecord.FirmName,
		FullName:    returnRecord.FullName,
		Country:     returnRecord.Country,
		Sector:      returnRecord.Sector,
		CountryName: returnRecord.CountryName,
		SectorName:  returnRecord.SectorName,
		Action:      "",
	}

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaFirmList)

}

func EditSienaFirmHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "editSienaFirm"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	searchID := core.GetURLparam(r, "SienaFirm")
	_, returnRecord, _ := getSienaFirm(searchID)

	//Get Country List & Populate and Array of dm.Country Items
	_, countryList, _ := dao.Country_GetList()
	_, sectorList, _ := getSienaSectorList()

	//fmt.Println(displayList)

	pageSienaFirmList := sienaFirmPage{
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		UserNavi:    "NOT USED",
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   core.ApplicationProperties["appname"] + " - " + "Firm - Edit",
		ID:          returnRecord.FirmName,
		FirmName:    returnRecord.FirmName,
		FullName:    returnRecord.FullName,
		Country:     returnRecord.Country,
		Sector:      returnRecord.Sector,
		CountryName: returnRecord.CountryName,
		SectorName:  returnRecord.SectorName,
		Action:      "",
		CountryList: countryList,
		SectorList:  sectorList,
	}
	//fmt.Println(pageSienaFirmList)

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaFirmList)

}

func SaveSienaFirmHandler(w http.ResponseWriter, r *http.Request) {
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

	var item sienaFirmItem

	item.FirmName = r.FormValue("firmName")
	if len(item.FirmName) == 0 {
		item.FirmName = r.FormValue("ID")
	}
	item.FullName = r.FormValue("fullName")
	item.Country = r.FormValue("country")

	item.Sector = r.FormValue("sector")
	item.Action = "UPDATE"

	fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE
	var sFldFirmName StaticImport_KeyField
	var sFldFullName StaticImport_Field
	var sFldCountry StaticImport_Field
	var sFldSector StaticImport_Field

	// POPULATE THE XML FIELDS
	sFldFirmName.Name = "firmName"
	sFldFirmName.Text = item.FirmName

	sFldFullName.Name = "fullName"
	sFldFullName.Text = item.FullName

	sFldCountry.Name = "country"
	sFldCountry.Text = item.Country

	sFldSector.Name = "sector"
	sFldSector.Text = item.Sector

	// IGNORE
	var sienaKeyFields []StaticImport_KeyField
	var sienaFields []StaticImport_Field
	// ADD THE FIELDS TO THE LISTS ABOVE
	sienaKeyFields = append(sienaKeyFields, sFldFirmName)
	sienaFields = append(sienaFields, sFldFullName)
	sienaFields = append(sienaFields, sFldCountry)
	sienaFields = append(sienaFields, sFldSector)
	// IGNORE
	sienaRecord := &StaticImport_Record{KEYFIELD: sienaKeyFields, FIELD: sienaFields}
	var sienaRecords []StaticImport_Record
	sienaRecords = append(sienaRecords, *sienaRecord)

	var sienaTable StaticImport_Table
	sienaTable.Name = "Firm"
	sienaTable.Classname = "com.eurobase.siena.data.firms.Firm"
	sienaTable.RECORD = sienaRecords

	var sienaTransaction StaticImport_Transaction
	sienaTransaction.Type = "IMPORT"
	sienaTransaction.TABLE = sienaTable

	var StaticImport_XMLContent StaticImport_XML
	StaticImport_XMLContent.TRANSACTIONS = sienaTransaction

	preparedXML, _ := xml.Marshal(StaticImport_XMLContent)
	fmt.Println("PreparedXML", string(preparedXML))

	staticImporterPath := core.SienaProperties["static_in"]
	fileID := uuid.New()
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

	ListSienaFirmHandler(w, r)

}

func NewSienaFirmHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "newSienaFirm"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	//Get Country List & Populate and Array of dm.Country Items
	_, countryList, _ := dao.Country_GetList()
	_, sectorList, _ := getSienaSectorList()

	pageSienaFirmList := sienaFirmPage{
		UserMenu:    UserMenu_Get(r),
		UserRole:    core.GetUserRole(r),
		UserNavi:    "NOT USED",
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   core.ApplicationProperties["appname"] + " - " + "Firm - New",
		ID:          "NEW",
		FirmName:    "",
		FullName:    "",
		Country:     "",
		Sector:      "",
		Action:      "NEW",
		CountryList: countryList,
		SectorList:  sectorList,
	}

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaFirmList)

}

// getSienaFirmList read all employees
func getSienaFirmList() (int, []sienaFirmItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaFirm;", sienaFirmSQL, core.SienaPropertiesDB["schema"])
	count, sienaFirmList, _, _ := fetchSienaFirmData(tsql)

	return count, sienaFirmList, nil
}

// getSienaFirmList read all employees
func getSienaFirm(id string) (int, sienaFirmItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaFirm WHERE FirmName='%s';", sienaFirmSQL, core.SienaPropertiesDB["schema"], id)
	_, _, sienaFirm, _ := fetchSienaFirmData(tsql)

	return 1, sienaFirm, nil
}

// getSienaFirmList read all employees
func putSienaFirm(updateItem sienaFirmItem) error {

	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(core.SienaPropertiesDB["schema"])
	fmt.Println(updateItem)
	return nil
}

// fetchSienaFirmData read all employees
func fetchSienaFirmData(tsql string) (int, []sienaFirmItem, sienaFirmItem, error) {

	var sienaFirm sienaFirmItem
	var sienaFirmList []sienaFirmItem

	rows, err := core.SienaDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaFirm, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlFRMFirmName, &sqlFRMFullName, &sqlFRMCountry, &sqlFRMSector, &sqlFRMSectorName, &sqlFRMCountryName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaFirm, err
		}

		sienaFirm.FirmName = sqlFRMFirmName.String
		sienaFirm.FullName = sqlFRMFullName.String
		sienaFirm.Country = sqlFRMCountry.String
		sienaFirm.Sector = sqlFRMSector.String
		sienaFirm.SectorName = sqlFRMSectorName.String
		sienaFirm.CountryName = sqlFRMCountryName.String

		sienaFirmList = append(sienaFirmList, sienaFirm)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaFirmList, sienaFirm, nil
}
