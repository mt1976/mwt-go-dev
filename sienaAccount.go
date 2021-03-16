package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

//sienaAccountPage is cheese
type sienaAccountListPage struct {
	Title             string
	PageTitle         string
	SienaAccountCount int
	SienaAccountList  []sienaAccountItem
}

//sienaAccountPage is cheese
type sienaAccountPage struct {
	Title             string
	PageTitle         string
	ID                string
	SienaReference    string
	Counterparty      string
	Status            string
	Opened            time.Time
	Closed            time.Time
	ExternalReference string
	CCY               string
	Book              string
	MandatedUser      string
	BackOfficeNotes   string
	CashBalance       float64
	AccountNumber     string
	AccountName       string
	LedgerBalance     float64
	Portfolio         string
	UTI               string
	CCYName           string
	BookName          string
	Action            string
	CountryList       []sienaCountryItem
	NameFirm          string
	NameCentre        string
	//	SectorList  []sienaSectorItem
}

//sienaAccountItem is cheese
type sienaAccountItem struct {
	SienaReference    string
	Counterparty      string
	Status            string
	Opened            time.Time
	Closed            time.Time
	ExternalReference string
	CCY               string
	Book              string
	MandatedUser      string
	BackOfficeNotes   string
	CashBalance       float64
	CBNeg             string
	AccountNumber     string
	AccountName       string
	LedgerBalance     float64
	LBNeg             string
	Portfolio         string
	UTI               string
	CCYName           string
	BookName          string
	Action            string
	NameFirm          string
	NameCentre        string
}

func listSienaAccountHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "listSienaAccount"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaAccountItem
	noItems, returnList, _ := getSienaAccountList(thisConnection)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaAccountList := sienaAccountListPage{
		Title:             wctProperties["appname"],
		PageTitle:         "List Siena Accounts",
		SienaAccountCount: noItems,
		SienaAccountList:  returnList,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaAccountList)

}

func viewSienaAccountHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "viewSienaAccount"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaAccountItem
	searchID := getURLparam(r, "SienaAccountID")
	noItems, returnRecord, _ := getSienaAccount(thisConnection, searchID)
	//fmt.Println("NoSienaItems", noItems, searchID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	pageSienaAccountList := sienaAccountPage{
		Title:             wctProperties["appname"],
		PageTitle:         "View Siena Account",
		ID:                returnRecord.SienaReference,
		SienaReference:    returnRecord.SienaReference,
		Counterparty:      returnRecord.Counterparty,
		Status:            returnRecord.Status,
		Opened:            returnRecord.Opened,
		Closed:            returnRecord.Closed,
		ExternalReference: returnRecord.ExternalReference,
		CCY:               returnRecord.CCY,
		Book:              returnRecord.Book,
		MandatedUser:      returnRecord.MandatedUser,
		BackOfficeNotes:   returnRecord.BackOfficeNotes,
		CashBalance:       returnRecord.CashBalance,
		AccountNumber:     returnRecord.AccountNumber,
		AccountName:       returnRecord.AccountName,
		LedgerBalance:     returnRecord.LedgerBalance,
		Portfolio:         returnRecord.Portfolio,
		UTI:               returnRecord.UTI,
		CCYName:           returnRecord.CCYName,
		BookName:          returnRecord.BookName,
		NameFirm:          returnRecord.NameFirm,
		NameCentre:        returnRecord.NameCentre,
		Action:            "",
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaAccountList)

}

func editSienaAccountHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "editSienaAccount"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaAccountItem
	searchID := getURLparam(r, "SienaAccount")
	noItems, returnRecord, _ := getSienaAccount(thisConnection, searchID)
	//fmt.Println("NoSienaItems", noItems, searchID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	//Get Country List & Populate and Array of sienaCountryItem Items
	_, countryList, _ := getSienaCountryList(thisConnection)
	//_, sectorList, _ := getSienaSectorList(thisConnection)

	//fmt.Println(displayList)

	pageSienaAccountList := sienaAccountPage{
		Title:          wctProperties["appname"],
		PageTitle:      "View Siena Account",
		ID:             returnRecord.SienaReference,
		SienaReference: returnRecord.SienaReference,
		Counterparty:   returnRecord.Counterparty,
		Status:         returnRecord.Status,
		Opened:         returnRecord.Opened,
		//	Closed:            returnRecord.Closed,
		ExternalReference: returnRecord.ExternalReference,
		CCY:               returnRecord.CCY,
		Book:              returnRecord.Book,
		MandatedUser:      returnRecord.MandatedUser,
		BackOfficeNotes:   returnRecord.BackOfficeNotes,
		CashBalance:       returnRecord.CashBalance,
		AccountNumber:     returnRecord.AccountNumber,
		AccountName:       returnRecord.AccountName,
		LedgerBalance:     returnRecord.LedgerBalance,
		Portfolio:         returnRecord.Portfolio,
		UTI:               returnRecord.UTI,
		CCYName:           returnRecord.CCYName,
		BookName:          returnRecord.BookName,
		Action:            "",
		CountryList:       countryList,
		NameFirm:          returnRecord.NameFirm,
		NameCentre:        returnRecord.NameCentre,
		//	SectorList:  sectorList,
	}
	//fmt.Println(pageSienaAccountList)

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaAccountList)

}

func saveSienaAccountHandler(w http.ResponseWriter, r *http.Request) {

	//sienaProperties := getProperties(cSIENACONFIG)
	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL, " : Save")

	//	var item sienaAccountItem

	listSienaAccountHandler(w, r)

}

func newSienaAccountHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "newSienaAccount"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	//Get Country List & Populate and Array of sienaCountryItem Items
	thisConnection, _ := sienaConnect()

	_, countryList, _ := getSienaCountryList(thisConnection)
	//	_, sectorList, _ := getSienaSectorList(thisConnection)

	pageSienaAccountList := sienaAccountPage{
		Title:       wctProperties["appname"],
		PageTitle:   "View Siena Account",
		ID:          "NEW",
		Action:      "NEW",
		CountryList: countryList,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaAccountList)

}

// getSienaAccountList read all employees
func getSienaAccountList(db *sql.DB) (int, []sienaAccountItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	var sienaAccount sienaAccountItem
	var sienaAccountList []sienaAccountItem

	tsql := fmt.Sprintf("SELECT SienaReference, CustomerSienaView, Status,StartDate,ExternalReference,CCY,Book,MandatedUser,BackOfficeNotes,CashBalance,AccountNumber,AccountName,LedgerBalance,Portfolio,UTI,CCYName,BookName, Firm, Centre FROM %s.sienaAccount;", mssqlConfig["schema"])
	//fmt.Println("MS SQL:", tsql)

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		var SienaReference, Counterparty, Status, CCY, Book, AccountNumber, AccountName, CCYName, BookName, Firm, Centre string
		//var AccountName, fullName, country, sector, CountryName, SectorName string
		var CashBalance, LedgerBalance float64
		var MandatedUser, Portfolio, UTI, ExternalReference, BackOfficeNotes sql.NullString
		var Opened time.Time
		err := rows.Scan(&SienaReference, &Counterparty, &Status, &Opened, &ExternalReference, &CCY, &Book, &MandatedUser, &BackOfficeNotes, &CashBalance, &AccountNumber, &AccountName, &LedgerBalance, &Portfolio, &UTI, &CCYName, &BookName, &Firm, &Centre)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, err
		}

		sienaAccount.SienaReference = SienaReference
		sienaAccount.Counterparty = Counterparty
		sienaAccount.Status = Status
		sienaAccount.Opened = Opened
		//sienaAccount.Closed = Closed
		sienaAccount.ExternalReference = ExternalReference.String
		sienaAccount.CCY = CCY
		sienaAccount.CCYName = CCYName
		sienaAccount.Book = Book
		sienaAccount.BookName = BookName
		sienaAccount.MandatedUser = MandatedUser.String
		sienaAccount.BackOfficeNotes = BackOfficeNotes.String
		sienaAccount.CashBalance = CashBalance
		sienaAccount.AccountNumber = AccountNumber
		sienaAccount.AccountName = AccountName
		sienaAccount.LedgerBalance = LedgerBalance
		sienaAccount.Portfolio = Portfolio.String
		sienaAccount.UTI = UTI.String
		sienaAccount.LBNeg = ""
		sienaAccount.CBNeg = ""
		sienaAccount.NameFirm = strings.TrimSpace(Firm)
		sienaAccount.NameCentre = strings.TrimSpace(Centre)

		if LedgerBalance < 0.0 {
			sienaAccount.LBNeg = "color:red"
		}
		if CashBalance < 0.0 {
			sienaAccount.CBNeg = "color:red"
		}
		sienaAccountList = append(sienaAccountList, sienaAccount)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaAccountList, nil
}

// getSienaAccountList read all employees
func getSienaAccountListByCounterParty(db *sql.DB, idFirm string, idCentre string) (int, []sienaAccountItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	var sienaAccount sienaAccountItem
	var sienaAccountList []sienaAccountItem

	tsql := fmt.Sprintf("SELECT SienaReference, CustomerSienaView, Status,StartDate,ExternalReference,CCY,Book,MandatedUser,BackOfficeNotes,CashBalance,AccountNumber,AccountName,LedgerBalance,Portfolio,UTI,CCYName,BookName, Firm, Centre FROM %s.sienaAccount WHERE Firm='%s' AND Centre='%s';", mssqlConfig["schema"], idFirm, idCentre)
	//fmt.Println("MS SQL:", tsql)

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		var SienaReference, Counterparty, Status, CCY, Book, AccountNumber, AccountName, CCYName, BookName, Firm, Centre string
		//var AccountName, fullName, country, sector, CountryName, SectorName string
		var CashBalance, LedgerBalance float64
		var MandatedUser, Portfolio, UTI, ExternalReference, BackOfficeNotes sql.NullString
		var Opened time.Time
		err := rows.Scan(&SienaReference, &Counterparty, &Status, &Opened, &ExternalReference, &CCY, &Book, &MandatedUser, &BackOfficeNotes, &CashBalance, &AccountNumber, &AccountName, &LedgerBalance, &Portfolio, &UTI, &CCYName, &BookName, &Firm, &Centre)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, err
		}

		sienaAccount.SienaReference = SienaReference
		sienaAccount.Counterparty = Counterparty
		sienaAccount.Status = Status
		sienaAccount.Opened = Opened
		//sienaAccount.Closed = Closed
		sienaAccount.ExternalReference = ExternalReference.String
		sienaAccount.CCY = CCY
		sienaAccount.CCYName = CCYName
		sienaAccount.Book = Book
		sienaAccount.BookName = BookName
		sienaAccount.MandatedUser = MandatedUser.String
		sienaAccount.BackOfficeNotes = BackOfficeNotes.String
		sienaAccount.CashBalance = CashBalance
		sienaAccount.AccountNumber = AccountNumber
		sienaAccount.AccountName = AccountName
		sienaAccount.LedgerBalance = LedgerBalance
		sienaAccount.Portfolio = Portfolio.String
		sienaAccount.UTI = UTI.String
		sienaAccount.LBNeg = ""
		sienaAccount.CBNeg = ""
		sienaAccount.NameFirm = strings.TrimSpace(Firm)
		sienaAccount.NameCentre = strings.TrimSpace(Centre)

		if LedgerBalance < 0.0 {
			sienaAccount.LBNeg = "color:red"
		}
		if CashBalance < 0.0 {
			sienaAccount.CBNeg = "color:red"
		}
		sienaAccountList = append(sienaAccountList, sienaAccount)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaAccountList, nil
}

// getSienaAccountList read all employees
func getSienaAccount(db *sql.DB, id string) (int, sienaAccountItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	var sienaAccount sienaAccountItem

	tsql := fmt.Sprintf("SELECT SienaReference, CustomerSienaView, Status,StartDate,ExternalReference,CCY,Book,MandatedUser,BackOfficeNotes,CashBalance,AccountNumber,AccountName,LedgerBalance,Portfolio,UTI,CCYName,BookName,Firm,Centre FROM %s.sienaAccount WHERE SienaReference='%s';", mssqlConfig["schema"], id)
	//fmt.Println("MS SQL:", tsql)

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, sienaAccount, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0

	for rows.Next() {
		var SienaReference, Counterparty, Status, CCY, Book, AccountNumber, AccountName, CCYName, BookName, Firm, Centre string
		//var AccountName, fullName, country, sector, CountryName, SectorName string
		var CashBalance, LedgerBalance float64
		var Opened time.Time
		var MandatedUser, Portfolio, UTI, ExternalReference, BackOfficeNotes sql.NullString
		err := rows.Scan(&SienaReference, &Counterparty, &Status, &Opened, &ExternalReference, &CCY, &Book, &MandatedUser, &BackOfficeNotes, &CashBalance, &AccountNumber, &AccountName, &LedgerBalance, &Portfolio, &UTI, &CCYName, &BookName, &Firm, &Centre)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, sienaAccount, err
		}

		sienaAccount.SienaReference = SienaReference
		sienaAccount.Counterparty = Counterparty
		sienaAccount.Status = Status
		sienaAccount.Opened = Opened
		//sienaAccount.Closed = Closed
		sienaAccount.ExternalReference = ExternalReference.String
		sienaAccount.CCY = CCY
		sienaAccount.CCYName = CCYName
		sienaAccount.Book = Book
		sienaAccount.BookName = BookName
		sienaAccount.MandatedUser = MandatedUser.String
		sienaAccount.BackOfficeNotes = BackOfficeNotes.String
		sienaAccount.CashBalance = CashBalance
		sienaAccount.AccountNumber = AccountNumber
		sienaAccount.AccountName = AccountName
		sienaAccount.LedgerBalance = LedgerBalance
		sienaAccount.Portfolio = Portfolio.String
		sienaAccount.UTI = UTI.String
		sienaAccount.NameFirm = strings.TrimSpace(Firm)
		sienaAccount.NameCentre = strings.TrimSpace(Centre)

		//sienaAccountList = append(sienaAccountList, sienaAccount)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return 1, sienaAccount, nil
}

// getSienaAccountList read all employees
func putSienaAccount(db *sql.DB, updateItem sienaAccountItem) error {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(mssqlConfig["schema"])
	fmt.Println(updateItem)
	return nil
}
