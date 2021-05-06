package siena

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	application "github.com/mt1976/mwt-go-dev/application"
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

func ListSienaAccountTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "listSienaAccountTransactions"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessage(inUTL)
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaAccountTransactionItem
	accountID := application.GetURLparam(r, "SienaAccountID")
	accountCCY := application.GetURLparam(r, "CCY")
	noItems, returnList, _ := getSienaAccountTransactionsList(accountID, accountCCY)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)
	_, account, _ := getSienaAccount(accountID)

	pageSienaAccountTransactionsList := sienaAccountTransactionListPage{
		Title:                        globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - “ +                    "List Siena AccountTransactionss",
		SienaAccountTransactionCount: noItems,
		SienaAccountTransactionList:  returnList,
		ID:                           account.AccountNumber,
		Name:                         account.AccountName,
		UserMenu:                     application.GetUserMenu(r),
		UserRole:                     application.GetUserRole(r),
		UserNavi:                     "NOT USED",
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaAccountTransactionsList)

}

// getSienaAccountTransactionsList read all employees
func getSienaAccountTransactionsList(idRef string, accountCCY string) (int, []sienaAccountTransactionItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaAccountTransactions WHERE SienaReference=%s;", sienaAccountTransactionsSQL, globals.SienaPropertiesDB["schema"], idRef)
	count, sienaAccountTransactionsList, _, _ := fetchSienaAccountTransactionData(tsql)
	return count, sienaAccountTransactionsList, nil
}

// fetchSienaAccountTransactionData read all employees
func fetchSienaAccountTransactionData(tsql string) (int, []sienaAccountTransactionItem, sienaAccountTransactionItem, error) {

	var sienaAccountTransaction sienaAccountTransactionItem
	var sienaAccountTransactionList []sienaAccountTransactionItem

	rows, err := globals.SienaDB.Query(tsql)
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
		sienaAccountTransaction.Amount = application.FormatCurrencyDps(sqlACTXAmount.String, sqlACTXDealtCcy.String, sqlACTXAmountDp.String)
		sienaAccountTransaction.StartInterestDate = application.SqlDateToHTMLDate(sqlACTXStartInterestDate.String)
		sienaAccountTransaction.EndInterestDate = application.SqlDateToHTMLDate(sqlACTXEndInterestDate.String)
		sienaAccountTransaction.Amortisation = sqlACTXAmortisation.String
		sienaAccountTransaction.InterestAmount = sqlACTXInterestAmount.String
		sienaAccountTransaction.InterestAction = sqlACTXInterestAction.String
		sienaAccountTransaction.FixingDate = application.SqlDateToHTMLDate(sqlACTXFixingDate.String)
		sienaAccountTransaction.InterestCalculationDate = application.SqlDateToHTMLDate(sqlACTXInterestCalculationDate.String)
		sienaAccountTransaction.AmendmentAmount = application.FormatCurrencyDps(sqlACTXAmendmentAmount.String, sqlACTXDealtCcy.String, sqlACTXAmountDp.String)
		sienaAccountTransaction.DealtCcy = sqlACTXDealtCcy.String
		sienaAccountTransaction.AmountDp = sqlACTXAmountDp.String
		sienaAccountTransactionList = append(sienaAccountTransactionList, sienaAccountTransaction)
		count++
	}
	log.Println(count, sienaAccountTransactionList, sienaAccountTransaction)
	return count, sienaAccountTransactionList, sienaAccountTransaction, nil
}
