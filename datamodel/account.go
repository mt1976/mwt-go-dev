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
	Account_Title        = "Account"
	Account_URI          = "SienaAccountID"
	Account_TemplateList = "listSienaAccount"
	Account_TemplateView = "viewSienaAccount"
	//Account_TemplateEdit   = "AccountStoreEdit"
	//Account_TemplateSave   = "AccountStoreSave"
	//Account_TemplateDelete = "AccountStoreDelete"
	//Account_TemplateNew    = "AccountStoreNew"
)
