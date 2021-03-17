package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Defines the Fields to Fetch from SQL
var sienaAccountTransactionsSQL = "SienaReference, 	LegNo, 	MMLegNo, 	Narrative, 	Amount, 	StartInterestDate, 	EndInterestDate, 	Amortisation, 	InterestAmount, 	InterestAction, 	FixingDate, 	InterestCalculationDate, 	AmendmentAmount, 	DealtCcy"
var sqlACTXSienaReference, sqlACTXLegNo, sqlACTXMMLegNo, sqlACTXNarrative, sqlACTXAmount, sqlACTXStartInterestDate, sqlACTXEndInterestDate, sqlACTXAmortisation, sqlACTXInterestAmount, sqlACTXInterestAction, sqlACTXFixingDate, sqlACTXInterestCalculationDate, sqlACTXAmendmentAmount, sqlACTXDealtCcy sql.NullString

//sienaAccountTransactionsPage is cheese
type sienaAccountTransactionListPage struct {
	Title                        string
	PageTitle                    string
	SienaAccountTransactionCount int
	SienaAccountTransactionList  []sienaAccountTransactionItem
	ID                           string
	Name                         string
}

//sienaAccountTransactionsPage is cheese
type sienaAccountTransactionPage struct {
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
}

func listSienaAccountTransactionsHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "listSienaAccountTransactions"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaAccountTransactionItem
	accountID := getURLparam(r, "SienaAccountID")
	accountCCY := getURLparam(r, "CCY")
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
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaAccountTransactionsList)

}

// getSienaAccountTransactionsList read all employees
func getSienaAccountTransactionsList(db *sql.DB, idRef string, accountCCY string) (int, []sienaAccountTransactionItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
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
		err := rows.Scan(&sqlACTXSienaReference, &sqlACTXLegNo, &sqlACTXMMLegNo, &sqlACTXNarrative, &sqlACTXAmount, &sqlACTXStartInterestDate, &sqlACTXEndInterestDate, &sqlACTXAmortisation, &sqlACTXInterestAmount, &sqlACTXInterestAction, &sqlACTXFixingDate, &sqlACTXInterestCalculationDate, &sqlACTXAmendmentAmount, &sqlACTXDealtCcy)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaAccountTransaction, err
		}
		sienaAccountTransaction.SienaReference = sqlACTXSienaReference.String
		sienaAccountTransaction.LegNo = sqlACTXLegNo.String
		sienaAccountTransaction.MMLegNo = sqlACTXMMLegNo.String
		sienaAccountTransaction.Narrative = sqlACTXNarrative.String
		sienaAccountTransaction.Amount = formatCurrency(sqlACTXAmount.String, sqlACTXDealtCcy.String)
		sienaAccountTransaction.StartInterestDate = sqlDateToHTMLDate(sqlACTXStartInterestDate.String)
		sienaAccountTransaction.EndInterestDate = sqlDateToHTMLDate(sqlACTXEndInterestDate.String)
		sienaAccountTransaction.Amortisation = sqlACTXAmortisation.String
		sienaAccountTransaction.InterestAmount = sqlACTXInterestAmount.String
		sienaAccountTransaction.InterestAction = sqlACTXInterestAction.String
		sienaAccountTransaction.FixingDate = sqlDateToHTMLDate(sqlACTXFixingDate.String)
		sienaAccountTransaction.InterestCalculationDate = sqlDateToHTMLDate(sqlACTXInterestCalculationDate.String)
		sienaAccountTransaction.AmendmentAmount = formatCurrency(sqlACTXAmendmentAmount.String, sqlACTXDealtCcy.String)
		sienaAccountTransaction.DealtCcy = sqlACTXDealtCcy.String
		sienaAccountTransactionList = append(sienaAccountTransactionList, sienaAccountTransaction)
		count++
	}
	log.Println(count, sienaAccountTransactionList, sienaAccountTransaction)
	return count, sienaAccountTransactionList, sienaAccountTransaction, nil
}
