package dao

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

var sienaAccountSQL = "SienaReference, 	CustomerSienaView, 	SienaCommonRef, 	Status, 	StartDate, 	MaturityDate, 	ContractNumber, 	ExternalReference, 	CCY, 	Book, 	MandatedUser, 	BackOfficeNotes, 	CashBalance, 	AccountNumber, 	AccountName, 	LedgerBalance, 	Portfolio, 	AgreementId, 	BackOfficeRefNo, 	PaymentSystemSienaView, 	ISIN, 	UTI, 	CCYName, 	BookName, 	PortfolioName, 	Centre, 	Firm, 	CCYDp"
var sqlACCTSienaReference, sqlACCTCustomerSienaView, sqlACCTSienaCommonRef, sqlACCTStatus, sqlACCTStartDate, sqlACCTMaturityDate, sqlACCTContractNumber, sqlACCTExternalReference, sqlACCTCCY, sqlACCTBook, sqlACCTMandatedUser, sqlACCTBackOfficeNotes, sqlACCTCashBalance, sqlACCTAccountNumber, sqlACCTAccountName, sqlACCTLedgerBalance, sqlACCTPortfolio, sqlACCTAgreementId, sqlACCTBackOfficeRefNo, sqlACCTPaymentSystemSienaView, sqlACCTISIN, sqlACCTUTI, sqlACCTCCYName, sqlACCTBookName, sqlACCTPortfolioName, sqlACCTCentre, sqlACCTFirm, sqlACCTCCYDp sql.NullString

// getSienaAccountList read all employees
func Account_GetList() (int, []dm.Account, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaAccount;", sienaAccountSQL, core.SienaPropertiesDB["schema"])
	count, sienaAccountList, _, _ := fetchSienaAccountData(tsql)
	return count, sienaAccountList, nil
}

// getSienaAccountListByCounterParty read all employees
func Account_GetListByCounterparty(idFirm string, idCentre string) (int, []dm.Account, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaAccount WHERE Firm='%s' AND Centre='%s';", sienaAccountSQL, core.SienaPropertiesDB["schema"], idFirm, idCentre)
	count, sienaAccountList, _, _ := fetchSienaAccountData(tsql)
	return count, sienaAccountList, nil
}

// getSienaAccount read all employees
func Account_GetByID(id string) (int, dm.Account, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaAccount WHERE SienaReference='%s';", sienaAccountSQL, core.SienaPropertiesDB["schema"], id)
	_, _, sienaAccount, _ := fetchSienaAccountData(tsql)
	return 1, sienaAccount, nil
}

// getSienaAccountList read all employees
func Account_Store(updateItem dm.Account) error {

	//fmt.Println(db.Stats().OpenConnections)
	//fmt.Println(core.SienaPropertiesDB["schema"])
	//fmt.Println(updateItem)
	return nil
}

// fetchSienaAccountData read all employees
func fetchSienaAccountData(tsql string) (int, []dm.Account, dm.Account, error) {

	var sienaAccount dm.Account
	var sienaAccountList []dm.Account

	rows, err := core.SienaDB.Query(tsql)
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
		sienaAccount.StartDate = core.SqlDateToHTMLDate(sqlACCTStartDate.String)
		sienaAccount.MaturityDate = core.SqlDateToHTMLDate(sqlACCTMaturityDate.String)
		sienaAccount.ContractNumber = sqlACCTContractNumber.String
		sienaAccount.ExternalReference = sqlACCTExternalReference.String
		sienaAccount.CCY = sqlACCTCCY.String
		sienaAccount.Book = sqlACCTBook.String
		sienaAccount.MandatedUser = sqlACCTMandatedUser.String
		sienaAccount.BackOfficeNotes = sqlACCTBackOfficeNotes.String
		sienaAccount.CashBalance = core.FormatCurrencyDps(sqlACCTCashBalance.String, sqlACCTCCY.String, sqlACCTCCYDp.String)
		sienaAccount.AccountNumber = sqlACCTAccountNumber.String
		sienaAccount.AccountName = sqlACCTAccountName.String
		sienaAccount.LedgerBalance = core.FormatCurrencyDps(sqlACCTLedgerBalance.String, sqlACCTCCY.String, sqlACCTCCYDp.String)
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
