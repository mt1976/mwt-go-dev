package siena

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	application "github.com/mt1976/mwt-go-dev/application"
	support "github.com/mt1976/mwt-go-dev/appsupport"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

// Defines the Fields to Fetch from SQL
var sienaAccountTransactionsSQL = "SienaReference, 	LegNo, 	MMLegNo, 	Narrative, 	Amount, 	StartInterestDate, 	EndInterestDate, 	Amortisation, 	InterestAmount, 	InterestAction, 	FixingDate, 	InterestCalculationDate, 	AmendmentAmount, 	DealtCcy, AmountDp"
var sqlACTXSienaReference, sqlACTXLegNo, sqlACTXMMLegNo, sqlACTXNarrative, sqlACTXAmount, sqlACTXStartInterestDate, sqlACTXEndInterestDate, sqlACTXAmortisation, sqlACTXInterestAmount, sqlACTXInterestAction, sqlACTXFixingDate, sqlACTXInterestCalculationDate, sqlACTXAmendmentAmount, sqlACTXDealtCcy, sqlACTXAmountDp sql.NullString

//sienaAccountTransactionsPage is cheese
type sienaAccountTransactionListPage struct {
	UserMenu                     []application.AppMenuItem
	UserRole                     string
	UserNavi                     string
	Title                        string
	PageTitle                    string
	SienaAccountTransactionCount int
	SienaAccountTransactionList  []sienaAccountTransactionItem
	ID                           string
	Name                         string
}

//sienaAccountTransactionsPage is cheese
type sienaAccountTransactionPage struct {
	UserMenu                []application.AppMenuItem
	UserRole                string
	UserNavi                string
	Title                   string
	PageTitle               string
	ID                      string
	Action                  string
	SienaReference          string
	LegNo                   string
	MMLegNo                 string
	Narrative               string
	Amount                  string
	StartInterestDate       string
	EndInterestDate         string
	Amortisation            string
	InterestAmount          string
	InterestAction          string
	FixingDate              string
	InterestCalculationDate string
	AmendmentAmount         string
	DealtCcy                string
	AmountDp                string
}

//sienaAccountTransactionItem is cheese
type sienaAccountTransactionItem struct {
	Action                  string
	SienaReference          string
	LegNo                   string
	MMLegNo                 string
	Narrative               string
	Amount                  string
	StartInterestDate       string
	EndInterestDate         string
	Amortisation            string
	InterestAmount          string
	InterestAction          string
	FixingDate              string
	InterestCalculationDate string
	AmendmentAmount         string
	DealtCcy                string
	AmountDp                string
}

func listSienaAccountTransactionsHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(globals.APPCONFIG)
	tmpl := "listSienaAccountTransactions"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := Connect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaAccountTransactionItem
	accountID := support.GetURLparam(r, "SienaAccountID")
	accountCCY := support.GetURLparam(r, "CCY")
	noItems, returnList, _ := getSienaAccountTransactionsList(thisConnection, accountID, accountCCY)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)
	_, account, _ := getSienaAccount(thisConnection, accountID)

	pageSienaAccountTransactionsList := sienaAccountTransactionListPage{
		Title:                        wctProperties["appname"],
		PageTitle:                    "List Siena AccountTransactionss",
		SienaAccountTransactionCount: noItems,
		SienaAccountTransactionList:  returnList,
		ID:                           account.AccountNumber,
		Name:                         account.AccountName,
		UserMenu:                     application.GetAppMenuData(globals.UserRole),
		UserRole:                     globals.UserRole,
		UserNavi:                     globals.UserNavi,
	}

	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, pageSienaAccountTransactionsList)

}

// getSienaAccountTransactionsList read all employees
func getSienaAccountTransactionsList(db *sql.DB, idRef string, accountCCY string) (int, []sienaAccountTransactionItem, error) {
	mssqlConfig := support.GetProperties(globals.SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaAccountTransactions WHERE SienaReference=%s;", sienaAccountTransactionsSQL, mssqlConfig["schema"], idRef)
	count, sienaAccountTransactionsList, _, _ := fetchSienaAccountTransactionData(db, tsql)
	return count, sienaAccountTransactionsList, nil
}

// fetchSienaAccountTransactionData read all employees
func fetchSienaAccountTransactionData(db *sql.DB, tsql string) (int, []sienaAccountTransactionItem, sienaAccountTransactionItem, error) {

	var sienaAccountTransaction sienaAccountTransactionItem
	var sienaAccountTransactionList []sienaAccountTransactionItem

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaAccountTransaction, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlACTXSienaReference, &sqlACTXLegNo, &sqlACTXMMLegNo, &sqlACTXNarrative, &sqlACTXAmount, &sqlACTXStartInterestDate, &sqlACTXEndInterestDate, &sqlACTXAmortisation, &sqlACTXInterestAmount, &sqlACTXInterestAction, &sqlACTXFixingDate, &sqlACTXInterestCalculationDate, &sqlACTXAmendmentAmount, &sqlACTXDealtCcy, &sqlACTXAmountDp)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaAccountTransaction, err
		}
		sienaAccountTransaction.SienaReference = sqlACTXSienaReference.String
		sienaAccountTransaction.LegNo = sqlACTXLegNo.String
		sienaAccountTransaction.MMLegNo = sqlACTXMMLegNo.String
		sienaAccountTransaction.Narrative = sqlACTXNarrative.String
		sienaAccountTransaction.Amount = support.FormatCurrencyDps(sqlACTXAmount.String, sqlACTXDealtCcy.String, sqlACTXAmountDp.String)
		sienaAccountTransaction.StartInterestDate = support.SqlDateToHTMLDate(sqlACTXStartInterestDate.String)
		sienaAccountTransaction.EndInterestDate = support.SqlDateToHTMLDate(sqlACTXEndInterestDate.String)
		sienaAccountTransaction.Amortisation = sqlACTXAmortisation.String
		sienaAccountTransaction.InterestAmount = sqlACTXInterestAmount.String
		sienaAccountTransaction.InterestAction = sqlACTXInterestAction.String
		sienaAccountTransaction.FixingDate = support.SqlDateToHTMLDate(sqlACTXFixingDate.String)
		sienaAccountTransaction.InterestCalculationDate = support.SqlDateToHTMLDate(sqlACTXInterestCalculationDate.String)
		sienaAccountTransaction.AmendmentAmount = support.FormatCurrencyDps(sqlACTXAmendmentAmount.String, sqlACTXDealtCcy.String, sqlACTXAmountDp.String)
		sienaAccountTransaction.DealtCcy = sqlACTXDealtCcy.String
		sienaAccountTransaction.AmountDp = sqlACTXAmountDp.String
		sienaAccountTransactionList = append(sienaAccountTransactionList, sienaAccountTransaction)
		count++
	}
	log.Println(count, sienaAccountTransactionList, sienaAccountTransaction)
	return count, sienaAccountTransactionList, sienaAccountTransaction, nil
}
