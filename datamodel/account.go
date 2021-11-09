package datamodel

//sienaAccountItem is cheese
type Account struct {
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

const (
	//Action     string
	Account_Title        = "Account"
	Account_QueryString  = "SienaAccountID"
	Account_TemplateList = "listSienaAccount"
	Account_TemplateView = "viewSienaAccount"
	Account_TemplateEdit = "editSienaAccount"
	//SQL TableName (No schema)
	Account_SQLTable = "sienaAccount"
	//SQL Columns Returned from Query
	Account_SienaReference         = "SienaReference"
	Account_CustomerSienaView      = "CustomerSienaView"
	Account_SienaCommonRef         = "SienaCommonRef"
	Account_Status                 = "Status"
	Account_StartDate              = "StartDate"
	Account_MaturityDate           = "MaturityDate"
	Account_ContractNumber         = "ContractNumber"
	Account_ExternalReference      = "ExternalReference"
	Account_CCY                    = "CCY"
	Account_Book                   = "Book"
	Account_MandatedUser           = "MandatedUser"
	Account_BackOfficeNotes        = "BackOfficeNotes"
	Account_CashBalance            = "CashBalance"
	Account_AccountNumber          = "AccountNumber"
	Account_AccountName            = "AccountName"
	Account_LedgerBalance          = "LedgerBalance"
	Account_Portfolio              = "Portfolio"
	Account_AgreementId            = "AgreementId"
	Account_BackOfficeRefNo        = "BackOfficeRefNo"
	Account_PaymentSystemSienaView = "PaymentSystemSienaView"
	Account_ISIN                   = "ISIN"
	Account_UTI                    = "UTI"
	Account_CCYName                = "CCYName"
	Account_BookName               = "BookName"
	Account_PortfolioName          = "PortfolioName"
	Account_Centre                 = "Centre"
	Account_Firm                   = "Firm"
	Account_CCYDp                  = "CCYDp"
)
