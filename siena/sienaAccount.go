package siena

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	support "github.com/mt1976/mwt-go-dev/appsupport"
)

var sienaAccountSQL = "SienaReference, 	CustomerSienaView, 	SienaCommonRef, 	Status, 	StartDate, 	MaturityDate, 	ContractNumber, 	ExternalReference, 	CCY, 	Book, 	MandatedUser, 	BackOfficeNotes, 	CashBalance, 	AccountNumber, 	AccountName, 	LedgerBalance, 	Portfolio, 	AgreementId, 	BackOfficeRefNo, 	PaymentSystemSienaView, 	ISIN, 	UTI, 	CCYName, 	BookName, 	PortfolioName, 	Centre, 	Firm, 	CCYDp"
var sqlACCTSienaReference, sqlACCTCustomerSienaView, sqlACCTSienaCommonRef, sqlACCTStatus, sqlACCTStartDate, sqlACCTMaturityDate, sqlACCTContractNumber, sqlACCTExternalReference, sqlACCTCCY, sqlACCTBook, sqlACCTMandatedUser, sqlACCTBackOfficeNotes, sqlACCTCashBalance, sqlACCTAccountNumber, sqlACCTAccountName, sqlACCTLedgerBalance, sqlACCTPortfolio, sqlACCTAgreementId, sqlACCTBackOfficeRefNo, sqlACCTPaymentSystemSienaView, sqlACCTISIN, sqlACCTUTI, sqlACCTCCYName, sqlACCTBookName, sqlACCTPortfolioName, sqlACCTCentre, sqlACCTFirm, sqlACCTCCYDp sql.NullString

//sienaAccountPage is cheese
type sienaAccountListPage struct {
	UserMenu          []AppMenuItem
	Title             string
	PageTitle         string
	SienaAccountCount int
	SienaAccountList  []sienaAccountItem
	UserRole          string
	UserNavi          string
}

//sienaAccountPage is cheese
type sienaAccountPage struct {
	UserMenu               []AppMenuItem
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

func listSienaAccountHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
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
		UserMenu:          getappMenuData(gUserRole),
		UserRole:          gUserRole,
		UserNavi:          gUserNavi,
	}

	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, gUserRole))
	t.Execute(w, pageSienaAccountList)

}

func viewSienaAccountHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	tmpl := "viewSienaAccount"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	//var returnList []sienaAccountItem
	searchID := support.GetURLparam(r, "SienaAccountID")
	_, returnRecord, _ := getSienaAccount(thisConnection, searchID)
	//fmt.Println("NoSienaItems", noItems, searchID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	pageSienaAccountList := sienaAccountPage{
		UserMenu:               getappMenuData(gUserRole),
		UserRole:               gUserRole,
		UserNavi:               gUserNavi,
		Title:                  wctProperties["appname"],
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

	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, gUserRole))
	t.Execute(w, pageSienaAccountList)

}

func editSienaAccountHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	tmpl := "editSienaAccount"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	//var returnList []sienaAccountItem
	searchID := support.GetURLparam(r, "SienaAccount")
	_, returnRecord, _ := getSienaAccount(thisConnection, searchID)
	//fmt.Println("NoSienaItems", noItems, searchID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	//Get Country List & Populate and Array of sienaCountryItem Items
	_, countryList, _ := getSienaCountryList(thisConnection)
	//_, sectorList, _ := getSienaSectorList(thisConnection)

	//fmt.Println(displayList)

	pageSienaAccountList := sienaAccountPage{
		UserMenu:               getappMenuData(gUserRole),
		UserRole:               gUserRole,
		UserNavi:               gUserNavi,
		Title:                  wctProperties["appname"],
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

	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, gUserRole))
	t.Execute(w, pageSienaAccountList)

}

func saveSienaAccountHandler(w http.ResponseWriter, r *http.Request) {

	//sienaProperties := support.GetProperties(cSIENACONFIG)
	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL, " : Save")

	//	var item sienaAccountItem

	listSienaAccountHandler(w, r)

}

func newSienaAccountHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	tmpl := "newSienaAccount"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	//Get Country List & Populate and Array of sienaCountryItem Items
	thisConnection, _ := sienaConnect()

	_, countryList, _ := getSienaCountryList(thisConnection)
	//	_, sectorList, _ := getSienaSectorList(thisConnection)

	pageSienaAccountList := sienaAccountPage{
		UserMenu:    getappMenuData(gUserRole),
		UserRole:    gUserRole,
		UserNavi:    gUserNavi,
		Title:       wctProperties["appname"],
		PageTitle:   "View Siena Account",
		ID:          "NEW",
		Action:      "NEW",
		CountryList: countryList,
	}

	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, gUserRole))
	t.Execute(w, pageSienaAccountList)

}

// getSienaAccountList read all employees
func getSienaAccountList(db *sql.DB) (int, []sienaAccountItem, error) {
	mssqlConfig := support.GetProperties(SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaAccount;", sienaAccountSQL, mssqlConfig["schema"])
	count, sienaAccountList, _, _ := fetchSienaAccountData(db, tsql)
	return count, sienaAccountList, nil
}

// getSienaAccountListByCounterParty read all employees
func getSienaAccountListByCounterParty(db *sql.DB, idFirm string, idCentre string) (int, []sienaAccountItem, error) {
	mssqlConfig := support.GetProperties(SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaAccount WHERE Firm='%s' AND Centre='%s';", sienaAccountSQL, mssqlConfig["schema"], idFirm, idCentre)
	count, sienaAccountList, _, _ := fetchSienaAccountData(db, tsql)
	return count, sienaAccountList, nil
}

// getSienaAccount read all employees
func getSienaAccount(db *sql.DB, id string) (int, sienaAccountItem, error) {
	mssqlConfig := support.GetProperties(SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaAccount WHERE SienaReference='%s';", sienaAccountSQL, mssqlConfig["schema"], id)
	_, _, sienaAccount, _ := fetchSienaAccountData(db, tsql)
	return 1, sienaAccount, nil
}

// getSienaAccountList read all employees
func putSienaAccount(db *sql.DB, updateItem sienaAccountItem) error {
	mssqlConfig := support.GetProperties(SQLCONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(mssqlConfig["schema"])
	fmt.Println(updateItem)
	return nil
}

// fetchSienaAccountData read all employees
func fetchSienaAccountData(db *sql.DB, tsql string) (int, []sienaAccountItem, sienaAccountItem, error) {

	var sienaAccount sienaAccountItem
	var sienaAccountList []sienaAccountItem

	rows, err := db.Query(tsql)
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
		sienaAccount.StartDate = support.SqlDateToHTMLDate(sqlACCTStartDate.String)
		sienaAccount.MaturityDate = support.SqlDateToHTMLDate(sqlACCTMaturityDate.String)
		sienaAccount.ContractNumber = sqlACCTContractNumber.String
		sienaAccount.ExternalReference = sqlACCTExternalReference.String
		sienaAccount.CCY = sqlACCTCCY.String
		sienaAccount.Book = sqlACCTBook.String
		sienaAccount.MandatedUser = sqlACCTMandatedUser.String
		sienaAccount.BackOfficeNotes = sqlACCTBackOfficeNotes.String
		sienaAccount.CashBalance = support.FormatCurrencyDps(sqlACCTCashBalance.String, sqlACCTCCY.String, sqlACCTCCYDp.String)
		sienaAccount.AccountNumber = sqlACCTAccountNumber.String
		sienaAccount.AccountName = sqlACCTAccountName.String
		sienaAccount.LedgerBalance = support.FormatCurrencyDps(sqlACCTLedgerBalance.String, sqlACCTCCY.String, sqlACCTCCYDp.String)
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
