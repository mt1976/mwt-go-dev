package siena

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

var sienaAccountSQL = "SienaReference, 	CustomerSienaView, 	SienaCommonRef, 	Status, 	StartDate, 	MaturityDate, 	ContractNumber, 	ExternalReference, 	CCY, 	Book, 	MandatedUser, 	BackOfficeNotes, 	CashBalance, 	AccountNumber, 	AccountName, 	LedgerBalance, 	Portfolio, 	AgreementId, 	BackOfficeRefNo, 	PaymentSystemSienaView, 	ISIN, 	UTI, 	CCYName, 	BookName, 	PortfolioName, 	Centre, 	Firm, 	CCYDp"
var sqlACCTSienaReference, sqlACCTCustomerSienaView, sqlACCTSienaCommonRef, sqlACCTStatus, sqlACCTStartDate, sqlACCTMaturityDate, sqlACCTContractNumber, sqlACCTExternalReference, sqlACCTCCY, sqlACCTBook, sqlACCTMandatedUser, sqlACCTBackOfficeNotes, sqlACCTCashBalance, sqlACCTAccountNumber, sqlACCTAccountName, sqlACCTLedgerBalance, sqlACCTPortfolio, sqlACCTAgreementId, sqlACCTBackOfficeRefNo, sqlACCTPaymentSystemSienaView, sqlACCTISIN, sqlACCTUTI, sqlACCTCCYName, sqlACCTBookName, sqlACCTPortfolioName, sqlACCTCentre, sqlACCTFirm, sqlACCTCCYDp sql.NullString

//sienaAccountPage is cheese
type sienaAccountListPage struct {
	UserMenu          []application.AppMenuItem
	Title             string
	PageTitle         string
	SienaAccountCount int
	SienaAccountList  []sienaAccountItem
	UserRole          string
	UserNavi          string
}

//sienaAccountPage is cheese
type sienaAccountPage struct {
	UserMenu               []application.AppMenuItem
	UserRole               string
	UserNavi               string
	Title                  string
	PageTitle              string
	ID                     string
	SienaReference         string
	CustomerSienaView      string
	SienaCommonRef         string
	Status                 string
	StartDate              string
	MaturityDate           string
	ContractNumber         string
	ExternalReference      string
	CCY                    string
	Book                   string
	MandatedUser           string
	BackOfficeNotes        string
	CashBalance            string
	AccountNumber          string
	AccountName            string
	LedgerBalance          string
	Portfolio              string
	AgreementId            string
	BackOfficeRefNo        string
	PaymentSystemSienaView string
	ISIN                   string
	UTI                    string
	CCYName                string
	BookName               string
	PortfolioName          string
	Centre                 string
	Firm                   string
	CCYDp                  string
	Action                 string
	CountryList            []sienaCountryItem

	//	SectorList  []sienaSectorItem
}

//sienaAccountItem is cheese
type sienaAccountItem struct {
	SienaReference         string
	CustomerSienaView      string
	SienaCommonRef         string
	Status                 string
	StartDate              string
	MaturityDate           string
	ContractNumber         string
	ExternalReference      string
	CCY                    string
	Book                   string
	MandatedUser           string
	BackOfficeNotes        string
	CashBalance            string
	AccountNumber          string
	AccountName            string
	LedgerBalance          string
	Portfolio              string
	AgreementId            string
	BackOfficeRefNo        string
	PaymentSystemSienaView string
	ISIN                   string
	UTI                    string
	CCYName                string
	BookName               string
	PortfolioName          string
	Centre                 string
	Firm                   string
	CCYDp                  string
}

func ListSienaAccountHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "listSienaAccount"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaAccountItem
	noItems, returnList, _ := getSienaAccountList()
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaAccountList := sienaAccountListPage{
		Title:             globals.ApplicationProperties["appname"],
		PageTitle:         "List Siena Accounts",
		SienaAccountCount: noItems,
		SienaAccountList:  returnList,
		UserMenu:          application.GetAppMenuData(globals.UserRole),
		UserRole:          globals.UserRole,
		UserNavi:          globals.UserNavi,
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageSienaAccountList)

}

func ViewSienaAccountHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "viewSienaAccount"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	searchID := application.GetURLparam(r, "SienaAccountID")
	_, returnRecord, _ := getSienaAccount(searchID)
	//fmt.Println("NoSienaItems", noItems, searchID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	pageSienaAccountList := sienaAccountPage{
		UserMenu:               application.GetAppMenuData(globals.UserRole),
		UserRole:               globals.UserRole,
		UserNavi:               globals.UserNavi,
		Title:                  globals.ApplicationProperties["appname"],
		PageTitle:              "View Siena Account",
		ID:                     returnRecord.SienaReference,
		SienaReference:         returnRecord.SienaReference,
		CustomerSienaView:      returnRecord.CustomerSienaView,
		SienaCommonRef:         returnRecord.SienaCommonRef,
		Status:                 returnRecord.Status,
		StartDate:              returnRecord.StartDate,
		MaturityDate:           returnRecord.MaturityDate,
		ContractNumber:         returnRecord.ContractNumber,
		ExternalReference:      returnRecord.ExternalReference,
		CCY:                    returnRecord.CCY,
		Book:                   returnRecord.Book,
		MandatedUser:           returnRecord.MandatedUser,
		BackOfficeNotes:        returnRecord.BackOfficeNotes,
		CashBalance:            returnRecord.CashBalance,
		AccountNumber:          returnRecord.AccountNumber,
		AccountName:            returnRecord.AccountName,
		LedgerBalance:          returnRecord.LedgerBalance,
		Portfolio:              returnRecord.Portfolio,
		AgreementId:            returnRecord.AgreementId,
		BackOfficeRefNo:        returnRecord.BackOfficeRefNo,
		PaymentSystemSienaView: returnRecord.PaymentSystemSienaView,
		ISIN:                   returnRecord.ISIN,
		UTI:                    returnRecord.UTI,
		CCYName:                returnRecord.CCYName,
		BookName:               returnRecord.BookName,
		PortfolioName:          returnRecord.PortfolioName,
		Centre:                 returnRecord.Centre,
		Firm:                   returnRecord.Firm,
		CCYDp:                  returnRecord.CCYDp,
		Action:                 "",
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageSienaAccountList)

}

func EditSienaAccountHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "editSienaAccount"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	searchID := application.GetURLparam(r, "SienaAccount")
	_, returnRecord, _ := getSienaAccount(searchID)
	//fmt.Println("NoSienaItems", noItems, searchID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	//Get Country List & Populate and Array of sienaCountryItem Items
	_, countryList, _ := getSienaCountryList()
	//_, sectorList, _ := getSienaSectorList(thisConnection)

	//fmt.Println(displayList)

	pageSienaAccountList := sienaAccountPage{
		UserMenu:               application.GetAppMenuData(globals.UserRole),
		UserRole:               globals.UserRole,
		UserNavi:               globals.UserNavi,
		Title:                  globals.ApplicationProperties["appname"],
		PageTitle:              "View Siena Account",
		ID:                     returnRecord.SienaReference,
		SienaReference:         returnRecord.SienaReference,
		CustomerSienaView:      returnRecord.CustomerSienaView,
		SienaCommonRef:         returnRecord.SienaCommonRef,
		Status:                 returnRecord.Status,
		StartDate:              returnRecord.StartDate,
		MaturityDate:           returnRecord.MaturityDate,
		ContractNumber:         returnRecord.ContractNumber,
		ExternalReference:      returnRecord.ExternalReference,
		CCY:                    returnRecord.CCY,
		Book:                   returnRecord.Book,
		MandatedUser:           returnRecord.MandatedUser,
		BackOfficeNotes:        returnRecord.BackOfficeNotes,
		CashBalance:            returnRecord.CashBalance,
		AccountNumber:          returnRecord.AccountNumber,
		AccountName:            returnRecord.AccountName,
		LedgerBalance:          returnRecord.LedgerBalance,
		Portfolio:              returnRecord.Portfolio,
		AgreementId:            returnRecord.AgreementId,
		BackOfficeRefNo:        returnRecord.BackOfficeRefNo,
		PaymentSystemSienaView: returnRecord.PaymentSystemSienaView,
		ISIN:                   returnRecord.ISIN,
		UTI:                    returnRecord.UTI,
		CCYName:                returnRecord.CCYName,
		BookName:               returnRecord.BookName,
		PortfolioName:          returnRecord.PortfolioName,
		Centre:                 returnRecord.Centre,
		Firm:                   returnRecord.Firm,
		CCYDp:                  returnRecord.CCYDp,
		Action:                 "",
		CountryList:            countryList,

		//	SectorList:  sectorList,
	}
	//fmt.Println(pageSienaAccountList)

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageSienaAccountList)

}

func SaveSienaAccountHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	//globals.SienaProperties := application.GetProperties(cSIENACONFIG)
	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL, " : Save")

	//	var item sienaAccountItem

	ListSienaAccountHandler(w, r)

}

func NewSienaAccountHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "newSienaAccount"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	//Get Country List & Populate and Array of sienaCountryItem Items
	_, countryList, _ := getSienaCountryList()
	//	_, sectorList, _ := getSienaSectorList(thisConnection)

	pageSienaAccountList := sienaAccountPage{
		UserMenu:    application.GetAppMenuData(globals.UserRole),
		UserRole:    globals.UserRole,
		UserNavi:    globals.UserNavi,
		Title:       globals.ApplicationProperties["appname"],
		PageTitle:   "View Siena Account",
		ID:          "NEW",
		Action:      "NEW",
		CountryList: countryList,
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageSienaAccountList)

}

// getSienaAccountList read all employees
func getSienaAccountList() (int, []sienaAccountItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaAccount;", sienaAccountSQL, globals.SienaPropertiesDB["schema"])
	count, sienaAccountList, _, _ := fetchSienaAccountData(tsql)
	return count, sienaAccountList, nil
}

// getSienaAccountListByCounterParty read all employees
func getSienaAccountListByCounterParty(idFirm string, idCentre string) (int, []sienaAccountItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaAccount WHERE Firm='%s' AND Centre='%s';", sienaAccountSQL, globals.SienaPropertiesDB["schema"], idFirm, idCentre)
	count, sienaAccountList, _, _ := fetchSienaAccountData(tsql)
	return count, sienaAccountList, nil
}

// getSienaAccount read all employees
func getSienaAccount(id string) (int, sienaAccountItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaAccount WHERE SienaReference='%s';", sienaAccountSQL, globals.SienaPropertiesDB["schema"], id)
	_, _, sienaAccount, _ := fetchSienaAccountData(tsql)
	return 1, sienaAccount, nil
}

// getSienaAccountList read all employees
func putSienaAccount(updateItem sienaAccountItem) error {

	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(globals.SienaPropertiesDB["schema"])
	fmt.Println(updateItem)
	return nil
}

// fetchSienaAccountData read all employees
func fetchSienaAccountData(tsql string) (int, []sienaAccountItem, sienaAccountItem, error) {

	var sienaAccount sienaAccountItem
	var sienaAccountList []sienaAccountItem

	rows, err := globals.SienaDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaAccount, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlACCTSienaReference, &sqlACCTCustomerSienaView, &sqlACCTSienaCommonRef, &sqlACCTStatus, &sqlACCTStartDate, &sqlACCTMaturityDate, &sqlACCTContractNumber, &sqlACCTExternalReference, &sqlACCTCCY, &sqlACCTBook, &sqlACCTMandatedUser, &sqlACCTBackOfficeNotes, &sqlACCTCashBalance, &sqlACCTAccountNumber, &sqlACCTAccountName, &sqlACCTLedgerBalance, &sqlACCTPortfolio, &sqlACCTAgreementId, &sqlACCTBackOfficeRefNo, &sqlACCTPaymentSystemSienaView, &sqlACCTISIN, &sqlACCTUTI, &sqlACCTCCYName, &sqlACCTBookName, &sqlACCTPortfolioName, &sqlACCTCentre, &sqlACCTFirm, &sqlACCTCCYDp)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaAccount, err
		}

		sienaAccount.SienaReference = sqlACCTSienaReference.String
		sienaAccount.CustomerSienaView = sqlACCTCustomerSienaView.String
		sienaAccount.SienaCommonRef = sqlACCTSienaCommonRef.String
		sienaAccount.Status = sqlACCTStatus.String
		sienaAccount.StartDate = application.SqlDateToHTMLDate(sqlACCTStartDate.String)
		sienaAccount.MaturityDate = application.SqlDateToHTMLDate(sqlACCTMaturityDate.String)
		sienaAccount.ContractNumber = sqlACCTContractNumber.String
		sienaAccount.ExternalReference = sqlACCTExternalReference.String
		sienaAccount.CCY = sqlACCTCCY.String
		sienaAccount.Book = sqlACCTBook.String
		sienaAccount.MandatedUser = sqlACCTMandatedUser.String
		sienaAccount.BackOfficeNotes = sqlACCTBackOfficeNotes.String
		sienaAccount.CashBalance = application.FormatCurrencyDps(sqlACCTCashBalance.String, sqlACCTCCY.String, sqlACCTCCYDp.String)
		sienaAccount.AccountNumber = sqlACCTAccountNumber.String
		sienaAccount.AccountName = sqlACCTAccountName.String
		sienaAccount.LedgerBalance = application.FormatCurrencyDps(sqlACCTLedgerBalance.String, sqlACCTCCY.String, sqlACCTCCYDp.String)
		sienaAccount.Portfolio = sqlACCTPortfolio.String
		sienaAccount.AgreementId = sqlACCTAgreementId.String
		sienaAccount.BackOfficeRefNo = sqlACCTBackOfficeRefNo.String
		sienaAccount.PaymentSystemSienaView = sqlACCTPaymentSystemSienaView.String
		sienaAccount.ISIN = sqlACCTISIN.String
		sienaAccount.UTI = sqlACCTUTI.String
		sienaAccount.CCYName = sqlACCTCCYName.String
		sienaAccount.BookName = sqlACCTBookName.String
		sienaAccount.PortfolioName = sqlACCTPortfolioName.String
		sienaAccount.Centre = strings.TrimSpace(sqlACCTCentre.String)
		sienaAccount.Firm = strings.TrimSpace(sqlACCTFirm.String)
		sienaAccount.CCYDp = sqlACCTCCYDp.String

		sienaAccountList = append(sienaAccountList, sienaAccount)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaAccountList, sienaAccount, nil
}
