package dao

import (
	"database/sql"
	"log"
	"strings"

	core "github.com/mt1976/mwt-go-dev/core"
	"github.com/mt1976/mwt-go-dev/das"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

var sienaAccountSQL = "SienaReference, 	CustomerSienaView, 	SienaCommonRef, 	Status, 	StartDate, 	MaturityDate, 	ContractNumber, 	ExternalReference, 	CCY, 	Book, 	MandatedUser, 	BackOfficeNotes, 	CashBalance, 	AccountNumber, 	AccountName, 	LedgerBalance, 	Portfolio, 	AgreementId, 	BackOfficeRefNo, 	PaymentSystemSienaView, 	ISIN, 	UTI, 	CCYName, 	BookName, 	PortfolioName, 	Centre, 	Firm, 	CCYDp"
var sqlACCTSienaReference, sqlACCTCustomerSienaView, sqlACCTSienaCommonRef, sqlACCTStatus, sqlACCTStartDate, sqlACCTMaturityDate, sqlACCTContractNumber, sqlACCTExternalReference, sqlACCTCCY, sqlACCTBook, sqlACCTMandatedUser, sqlACCTBackOfficeNotes, sqlACCTCashBalance, sqlACCTAccountNumber, sqlACCTAccountName, sqlACCTLedgerBalance, sqlACCTPortfolio, sqlACCTAgreementId, sqlACCTBackOfficeRefNo, sqlACCTPaymentSystemSienaView, sqlACCTISIN, sqlACCTUTI, sqlACCTCCYName, sqlACCTBookName, sqlACCTPortfolioName, sqlACCTCentre, sqlACCTFirm, sqlACCTCCYDp sql.NullString

// getSienaAccountList read all employees
func Account_GetList() (int, []dm.Account, error) {

	//	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaAccount;", sienaAccountSQL, core.SienaPropertiesDB["schema"])

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Account_SQLTable)
	count, sienaAccountList, _, _ := fetchSienaAccountData(tsql)
	return count, sienaAccountList, nil
}

// getSienaAccountListByCounterParty read all employees
func Account_GetListByCounterparty(idFirm string, idCentre string) (int, []dm.Account, error) {

	//tsql := fmt.Sprintf("SELECT %s FROM %s.sienaAccount WHERE Firm='%s' AND Centre='%s';", sienaAccountSQL, core.SienaPropertiesDB["schema"], idFirm, idCentre)
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Account_SQLTable)
	tsql = tsql + " WHERE " + dm.Account_Firm + "='" + idFirm + "' AND " + dm.Account_Centre + "='" + idCentre + "'"

	count, sienaAccountList, _, _ := fetchSienaAccountData(tsql)
	return count, sienaAccountList, nil
}

// getSienaAccount read all employees
func Account_GetByID(id string) (int, dm.Account, error) {

	//tsql := fmt.Sprintf("SELECT %s FROM %s.sienaAccount WHERE SienaReference='%s';", sienaAccountSQL, core.SienaPropertiesDB["schema"], id)

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Account_SQLTable)
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

		sienaAccount.SienaReference = get_String(rec, dm.Account_SienaReference, "")
		sienaAccount.CustomerSienaView = get_String(rec, dm.Account_CustomerSienaView, "")
		sienaAccount.SienaCommonRef = get_String(rec, dm.Account_SienaCommonRef, "")
		sienaAccount.Status = get_String(rec, dm.Account_Status, "")
		sienaAccount.StartDate = get_Time(rec, dm.Account_StartDate, "")
		sienaAccount.MaturityDate = get_Time(rec, dm.Account_MaturityDate, "")
		sienaAccount.ContractNumber = get_String(rec, dm.Account_ContractNumber, "")
		sienaAccount.ExternalReference = get_String(rec, dm.Account_ExternalReference, "")
		sienaAccount.CCY = get_String(rec, dm.Account_CCY, "")
		sienaAccount.Book = get_String(rec, dm.Account_Book, "")
		sienaAccount.MandatedUser = get_String(rec, dm.Account_MandatedUser, "")
		sienaAccount.BackOfficeNotes = get_String(rec, dm.Account_BackOfficeNotes, "")
		sienaAccount.CashBalance = get_Float(rec, dm.Account_CashBalance, "0.0")
		sienaAccount.AccountNumber = get_String(rec, dm.Account_AccountNumber, "")
		sienaAccount.AccountName = get_String(rec, dm.Account_AccountName, "")
		sienaAccount.LedgerBalance = get_Float(rec, dm.Account_LedgerBalance, "0.0")
		sienaAccount.Portfolio = get_String(rec, dm.Account_Portfolio, "")
		sienaAccount.AgreementId = get_String(rec, dm.Account_AgreementId, "")
		sienaAccount.BackOfficeRefNo = get_String(rec, dm.Account_BackOfficeRefNo, "")
		sienaAccount.PaymentSystemSienaView = get_String(rec, dm.Account_PaymentSystemSienaView, "")
		sienaAccount.ISIN = get_String(rec, dm.Account_ISIN, "")
		sienaAccount.UTI = get_String(rec, dm.Account_UTI, "")
		sienaAccount.CCYName = get_String(rec, dm.Account_CCYName, "")
		sienaAccount.Firm = get_String(rec, dm.Account_Firm, "")
		sienaAccount.Centre = get_String(rec, dm.Account_Centre, "")
		sienaAccount.CCYDp = get_Int(rec, dm.Account_CCYDp, "2")
		sienaAccount.BookName = get_String(rec, dm.Account_BookName, "")
		sienaAccount.PortfolioName = get_String(rec, dm.Account_PortfolioName, "")
		sienaAccount.CCYName = get_String(rec, dm.Account_CCYName, "")
		//Post Import Actions
		sienaAccount.CashBalance = core.FormatCurrencyDps(sienaAccount.CashBalance, sienaAccount.CCY, sienaAccount.CCYDp)
		sienaAccount.LedgerBalance = core.FormatCurrencyDps(sienaAccount.LedgerBalance, sienaAccount.CCY, sienaAccount.CCYDp)
		sienaAccount.Centre = strings.TrimSpace(sienaAccount.Centre)
		sienaAccount.Firm = strings.TrimSpace(sienaAccount.Firm)
		//Add to the list
		sienaAccountList = append(sienaAccountList, sienaAccount)

	}

	return noitems, sienaAccountList, sienaAccount, nil
}
