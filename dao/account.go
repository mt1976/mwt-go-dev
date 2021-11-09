package dao

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	core "github.com/mt1976/mwt-go-dev/core"
	"github.com/mt1976/mwt-go-dev/das"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

var sienaAccountSQL = "SienaReference, 	CustomerSienaView, 	SienaCommonRef, 	Status, 	StartDate, 	MaturityDate, 	ContractNumber, 	ExternalReference, 	CCY, 	Book, 	MandatedUser, 	BackOfficeNotes, 	CashBalance, 	AccountNumber, 	AccountName, 	LedgerBalance, 	Portfolio, 	AgreementId, 	BackOfficeRefNo, 	PaymentSystemSienaView, 	ISIN, 	UTI, 	CCYName, 	BookName, 	PortfolioName, 	Centre, 	Firm, 	CCYDp"
var sqlACCTSienaReference, sqlACCTCustomerSienaView, sqlACCTSienaCommonRef, sqlACCTStatus, sqlACCTStartDate, sqlACCTMaturityDate, sqlACCTContractNumber, sqlACCTExternalReference, sqlACCTCCY, sqlACCTBook, sqlACCTMandatedUser, sqlACCTBackOfficeNotes, sqlACCTCashBalance, sqlACCTAccountNumber, sqlACCTAccountName, sqlACCTLedgerBalance, sqlACCTPortfolio, sqlACCTAgreementId, sqlACCTBackOfficeRefNo, sqlACCTPaymentSystemSienaView, sqlACCTISIN, sqlACCTUTI, sqlACCTCCYName, sqlACCTBookName, sqlACCTPortfolioName, sqlACCTCentre, sqlACCTFirm, sqlACCTCCYDp sql.NullString

// getSienaAccountList read all employees
func Account_GetList() (int, []dm.Account, error) {

	//	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaAccount;", sienaAccountSQL, core.SienaPropertiesDB["schema"])

	tsql := "SELECT * FROM " + queryTableName(core.SienaPropertiesDB["schema"], dm.Account_SQLTable)
	count, sienaAccountList, _, _ := fetchSienaAccountData(tsql)
	return count, sienaAccountList, nil
}

// getSienaAccountListByCounterParty read all employees
func Account_GetListByCounterparty(idFirm string, idCentre string) (int, []dm.Account, error) {

	//tsql := fmt.Sprintf("SELECT %s FROM %s.sienaAccount WHERE Firm='%s' AND Centre='%s';", sienaAccountSQL, core.SienaPropertiesDB["schema"], idFirm, idCentre)
	tsql := "SELECT * FROM " + queryTableName(core.SienaPropertiesDB["schema"], dm.Account_SQLTable)
	tsql = tsql + " WHERE " + dm.Account_Firm + "='" + idFirm + "' AND " + dm.Account_Centre + "='" + idCentre + "'"

	count, sienaAccountList, _, _ := fetchSienaAccountData(tsql)
	return count, sienaAccountList, nil
}

// getSienaAccount read all employees
func Account_GetByID(id string) (int, dm.Account, error) {

	//tsql := fmt.Sprintf("SELECT %s FROM %s.sienaAccount WHERE SienaReference='%s';", sienaAccountSQL, core.SienaPropertiesDB["schema"], id)

	tsql := "SELECT * FROM " + queryTableName(core.SienaPropertiesDB["schema"], dm.Account_SQLTable)
	tsql = tsql + " WHERE " + dm.Account_SienaReference + "='" + id + "'"
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

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	//	spew.Dump(returnList)

	for i := 0; i < noitems; i++ {

		rec := returnList[i]

		sienaAccount.SienaReference = rec[dm.Account_SienaReference].(string)
		sienaAccount.CustomerSienaView = rec[dm.Account_CustomerSienaView].(string)
		sienaAccount.SienaCommonRef = loadString(rec[dm.Account_SienaCommonRef])
		sienaAccount.Status = loadString(rec[dm.Account_Status])
		sienaAccount.StartDate = loadTime(rec[dm.Account_StartDate])
		sienaAccount.MaturityDate = loadTime(rec[dm.Account_MaturityDate])
		sienaAccount.ContractNumber = rec[dm.Account_ContractNumber].(string)
		sienaAccount.ExternalReference = loadString(rec[dm.Account_ExternalReference])
		sienaAccount.CCY = rec[dm.Account_CCY].(string)
		sienaAccount.Book = rec[dm.Account_Book].(string)
		sienaAccount.MandatedUser = loadString(rec[dm.Account_MandatedUser])
		sienaAccount.BackOfficeNotes = loadString(rec[dm.Account_BackOfficeNotes])
		sienaAccount.CashBalance = fmt.Sprintf("%f", rec[dm.Account_CashBalance].(float64))
		sienaAccount.AccountNumber = rec[dm.Account_AccountNumber].(string)
		sienaAccount.AccountName = rec[dm.Account_AccountName].(string)
		sienaAccount.LedgerBalance = fmt.Sprintf("%f", rec[dm.Account_LedgerBalance].(float64))
		sienaAccount.Portfolio = loadString(rec[dm.Account_Portfolio])
		sienaAccount.AgreementId = loadString(rec[dm.Account_AgreementId])
		sienaAccount.BackOfficeRefNo = loadString(rec[dm.Account_BackOfficeRefNo])
		sienaAccount.PaymentSystemSienaView = loadString(rec[dm.Account_PaymentSystemSienaView])
		sienaAccount.ISIN = loadString(rec[dm.Account_ISIN])
		sienaAccount.UTI = loadString(rec[dm.Account_UTI])
		sienaAccount.CCYName = loadString(rec[dm.Account_CCYName])
		sienaAccount.Firm = loadString(rec[dm.Account_Firm])
		sienaAccount.Centre = loadString(rec[dm.Account_Centre])
		sienaAccount.CCYDp = loadInt(rec[dm.Account_CCYDp])
		sienaAccount.BookName = loadString(rec[dm.Account_BookName])
		sienaAccount.PortfolioName = loadString(rec[dm.Account_PortfolioName])
		sienaAccount.CCYName = loadString(rec[dm.Account_CCYName])

		sienaAccount.CashBalance = core.FormatCurrencyDps(sienaAccount.CashBalance, sienaAccount.CCY, sienaAccount.CCYDp)
		sienaAccount.LedgerBalance = core.FormatCurrencyDps(sienaAccount.LedgerBalance, sienaAccount.CCY, sienaAccount.CCYDp)
		sienaAccount.Centre = strings.TrimSpace(sienaAccount.Centre)
		sienaAccount.Firm = strings.TrimSpace(sienaAccount.Firm)
		sienaAccountList = append(sienaAccountList, sienaAccount)

	}

	return noitems, sienaAccountList, sienaAccount, nil
}

//Convert time.Time to string
func loadTime(t interface{}) string {
	fmt.Printf("t: %v\n", t)
	if t == nil {
		return ""
	}
	return core.TimeToString(t.(time.Time))
}

//Convert time.Time to string
func loadString(t interface{}) string {
	fmt.Printf("t: %v\n", t)
	if t == nil {
		return ""
	}
	return t.(string)
}

//Convert time.Time to string
func loadInt(t interface{}) string {
	fmt.Printf("t: %v\n", t)
	if t != nil {
		return fmt.Sprintf("%d", t.(int64))
	}
	return ""
}
