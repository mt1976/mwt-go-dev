package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// Defines the Fields to Fetch from SQL
var sienaAccountLadderSQL = "SienaReference, BusinessDate, ContractNumber, Balance, DealtCcy,AmountDP"
var sqlACCLSienaReference, sqlACCLBusinessDate, sqlACCLContractNumber, sqlACCLBalance, sqlACCLDealtCcy, sqlACCLAmountDP sql.NullString

//sienaAccountLadderPage is cheese
type sienaAccountLadderListPage struct {
	UserMenu                []AppMenuItem
	UserRole                string
	UserNavi                string
	Title                   string
	PageTitle               string
	SienaAccountLadderCount int
	SienaAccountLadderList  []sienaAccountLadderItem
	ID                      string
	Name                    string
}

//sienaAccountLadderPage is cheese
type sienaAccountLadderPage struct {
	UserMenu       []AppMenuItem
	UserRole       string
	UserNavi       string
	Title          string
	PageTitle      string
	ID             string
	Action         string
	SienaReference string
	BusinessDate   string
	ContractNumber string
	Balance        string
	DealtCcy       string
	AmountDP       string
}

//sienaAccountLadderItem is cheese
type sienaAccountLadderItem struct {
	Action         string
	SienaReference string
	BusinessDate   string
	ContractNumber string
	Balance        string
	DealtCcy       string
	AmountDP       string
}

func listSienaAccountLadderHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "listSienaAccountLadder"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaAccountLadderItem
	accountID := getURLparam(r, "SienaAccountID")
	accountCCY := getURLparam(r, "CCY")
	noItems, returnList, _ := getSienaAccountLadderList(thisConnection, accountID, accountCCY)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)
	_, account, _ := getSienaAccount(thisConnection, accountID)

	pageSienaAccountLadderList := sienaAccountLadderListPage{
		UserMenu:                getappMenuData(gUserRole),
		UserRole:                gUserRole,
		UserNavi:                gUserNavi,
		Title:                   wctProperties["appname"],
		PageTitle:               "List Siena AccountLadders",
		SienaAccountLadderCount: noItems,
		SienaAccountLadderList:  returnList,
		ID:                      account.AccountNumber,
		Name:                    account.AccountName,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaAccountLadderList)

}

// getSienaAccountLadderList read all employees
func getSienaAccountLadderList(db *sql.DB, idRef string, acctCCY string) (int, []sienaAccountLadderItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaAccountLadder WHERE SienaReference=%s;", sienaAccountLadderSQL, mssqlConfig["schema"], idRef)
	count, sienaAccountLadderList, _, _ := fetchSienaAccountLadderData(db, tsql)
	return count, sienaAccountLadderList, nil
}

// fetchSienaAccountLadderData read all employees
func fetchSienaAccountLadderData(db *sql.DB, tsql string) (int, []sienaAccountLadderItem, sienaAccountLadderItem, error) {
	log.Println(tsql)
	var sienaAccountLadder sienaAccountLadderItem
	var sienaAccountLadderList []sienaAccountLadderItem

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaAccountLadder, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlACCLSienaReference, &sqlACCLBusinessDate, &sqlACCLContractNumber, &sqlACCLBalance, &sqlACCLDealtCcy, &sqlACCLAmountDP)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaAccountLadder, err
		}

		sienaAccountLadder.SienaReference = sqlACCLSienaReference.String
		sienaAccountLadder.BusinessDate = sqlDateToHTMLDate(sqlACCLBusinessDate.String)
		sienaAccountLadder.ContractNumber = sqlACCLContractNumber.String
		sienaAccountLadder.DealtCcy = sqlACCLDealtCcy.String
		sienaAccountLadder.Balance = formatCurrencyDps(sqlACCLBalance.String, sqlACCLDealtCcy.String, sqlACCLAmountDP.String)
		sienaAccountLadder.AmountDP = sqlACCLAmountDP.String
		sienaAccountLadderList = append(sienaAccountLadderList, sienaAccountLadder)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	log.Println(count, sienaAccountLadderList, sienaAccountLadder)
	return count, sienaAccountLadderList, sienaAccountLadder, nil
}
