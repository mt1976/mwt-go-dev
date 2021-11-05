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
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

var sienaPortfolioSQL = "Code, 	Name"
var sqlPRTCode, sqlPRTName sql.NullString

//sienaPortfolioPage is cheese
type portfolio_listpage struct {
	UserMenu            []dm.AppMenuItem
	UserRole            string
	UserNavi            string
	Title               string
	PageTitle           string
	SienaPortfolioCount int
	SienaPortfolioList  []portfolio_item
}

//portfolio_page is cheese
type portfolio_page struct {
	UserMenu  []dm.AppMenuItem
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Code      string
	Name      string
}

//portfolio_item is cheese
type portfolio_item struct {
	Code   string
	Name   string
	Action string
}

func Portfolio_MUX(mux http.ServeMux) {
	log.Println("MUX Siena Firm")
	mux.HandleFunc("/listSienaPortfolio/", Portfolio_HandlerList)
	mux.HandleFunc("/viewSienaPortfolio/", Portfolio_HandlerView)
	mux.HandleFunc("/editSienaPortfolio/", Portfolio_HandlerEdit)
	mux.HandleFunc("/saveSienaPortfolio/", Portfolio_HandlerSave)
	mux.HandleFunc("/newSienaPortfolio/", Portfolio_HandlerNew)
}

func Portfolio_HandlerList(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "listSienaPortfolio"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)
	thisConnection, _ := Connect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []portfolio_item
	noItems, returnList, _ := getSienaPortfolioList(thisConnection)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaPortfolioList := portfolio_listpage{
		UserMenu:            core.GetUserMenu(r),
		UserRole:            core.GetUserRole(r),
		UserNavi:            "NOT USED",
		Title:               core.ApplicationProperties["appname"],
		PageTitle:           core.ApplicationProperties["appname"] + " - " + "Bank Portfolios",
		SienaPortfolioCount: noItems,
		SienaPortfolioList:  returnList,
	}

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaPortfolioList)

}

func Portfolio_HandlerView(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "viewSienaPortfolio"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)
	thisConnection, _ := Connect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []portfolio_item
	searchID := core.GetURLparam(r, "SienaPortfolio")
	noItems, returnRecord, _ := getSienaPortfolio(thisConnection, searchID)
	fmt.Println("NoSienaItems", noItems, searchID)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	pageSienaPortfolioList := portfolio_page{
		UserMenu:  core.GetUserMenu(r),
		UserRole:  core.GetUserRole(r),
		UserNavi:  "NOT USED",
		Title:     core.ApplicationProperties["appname"],
		PageTitle: core.ApplicationProperties["appname"] + " - " + "Bank Portfolio - View",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
	}

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaPortfolioList)

}

func Portfolio_HandlerEdit(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "editSienaPortfolio"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)
	thisConnection, _ := Connect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []portfolio_item
	searchID := core.GetURLparam(r, "SienaPortfolio")
	noItems, returnRecord, _ := getSienaPortfolio(thisConnection, searchID)
	fmt.Println("NoSienaCountries", noItems)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	pageSienaPortfolioList := portfolio_page{
		UserMenu:  core.GetUserMenu(r),
		UserRole:  core.GetUserRole(r),
		UserNavi:  "NOT USED",
		Title:     core.ApplicationProperties["appname"],
		PageTitle: core.ApplicationProperties["appname"] + " - " + "Bank Portfolio - Edit",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
	}

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaPortfolioList)

}

func Portfolio_HandlerSave(w http.ResponseWriter, r *http.Request) {
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

	var item portfolio_item

	item.Code = r.FormValue("code")
	if len(item.Code) == 0 {
		item.Code = r.FormValue("ID")
	}
	item.Name = r.FormValue("name")
	item.Action = "UPDATE"

	fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE
	var sFldCode sienaKEYFIELD
	var sFldName sienaFIELD
	// POPULATE THE XML FIELDS
	sFldCode.Name = "Code"
	sFldCode.Text = item.Code

	sFldName.Name = "Description1"
	sFldName.Text = item.Name
	// IGNORE
	var sienaKeyFields []sienaKEYFIELD
	var sienaFields []sienaFIELD
	// ADD THE FIELDS TO THE LISTS ABOVE
	sienaKeyFields = append(sienaKeyFields, sFldCode)
	sienaFields = append(sienaFields, sFldName)
	// IGNORE
	sienaRecord := &sienaRECORD{KEYFIELD: sienaKeyFields, FIELD: sienaFields}
	var sienaRecords []sienaRECORD
	sienaRecords = append(sienaRecords, *sienaRecord)

	var sienaTable sienaTABLE
	sienaTable.Name = "Portfolio"
	sienaTable.Classname = "com.eurobase.siena.data.portfolio.Portfolio"
	sienaTable.RECORD = sienaRecords

	var sienaTransaction sienaTRANSACTION
	sienaTransaction.Type = "IMPORT"
	sienaTransaction.TABLE = sienaTable

	var sienaXMLContent sienaXML
	sienaXMLContent.TRANSACTIONS = sienaTransaction

	preparedXML, _ := xml.Marshal(sienaXMLContent)
	fmt.Println("PreparedXML", string(preparedXML))

	staticImporterPath := core.SienaProperties["static_in"]
	fileID := uuid.New()
	pwd, _ := os.Getwd()
	fileName := pwd + staticImporterPath + "/" + fileID.String() + ".xml"
	fmt.Println(fileName)

	err := ioutil.WriteFile(fileName, preparedXML, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	Portfolio_HandlerList(w, r)

}

func Portfolio_HandlerNew(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "newSienaPortfolio"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	pageSienaPortfolioList := portfolio_page{
		UserMenu:  core.GetUserMenu(r),
		UserRole:  core.GetUserRole(r),
		UserNavi:  "NOT USED",
		Title:     core.ApplicationProperties["appname"],
		PageTitle: core.ApplicationProperties["appname"] + " - " + "Bank Portfolio - New",
		ID:        "NEW",
		Code:      "",
		Name:      "",
	}

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, pageSienaPortfolioList)

}

// getSienaPortfolioList read all employees
func getSienaPortfolioList(db *sql.DB) (int, []portfolio_item, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaPortfolio;", sienaPortfolioSQL, core.SienaPropertiesDB["schema"])
	count, sienaPortfolioList, _, _ := fetchSienaPortfolioData(db, tsql)
	return count, sienaPortfolioList, nil
}

// getSienaPortfolioList read all employees
func getSienaPortfolio(db *sql.DB, id string) (int, portfolio_item, error) {

	//fmt.Println(db.Stats().OpenConnections)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaPortfolio WHERE Code='%s';", sienaPortfolioSQL, core.SienaPropertiesDB["schema"], id)
	_, _, sienaPortfolio, _ := fetchSienaPortfolioData(db, tsql)
	return 1, sienaPortfolio, nil
}

// getSienaPortfolioList read all employees
func putSienaPortfolio(db *sql.DB, updateItem portfolio_item) error {

	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(core.SienaPropertiesDB["schema"])
	fmt.Println(updateItem)
	return nil
}

// getSienaPortfolioList read all employees
func fetchSienaPortfolioData(db *sql.DB, tsql string) (int, []portfolio_item, portfolio_item, error) {
	log.Println("QUERY", tsql)
	var sienaPortfolioList []portfolio_item
	var sienaPortfolio portfolio_item

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaPortfolio, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlPRTCode, &sqlPRTName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaPortfolio, err
		}
		sienaPortfolio.Code = sqlPRTCode.String
		sienaPortfolio.Name = sqlPRTName.String
		sienaPortfolioList = append(sienaPortfolioList, sienaPortfolio)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	log.Println(count, sienaPortfolioList, sienaPortfolio)
	return count, sienaPortfolioList, sienaPortfolio, nil
}
