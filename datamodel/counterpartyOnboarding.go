package datamodel

type CounterpartyOnboarding struct {
	TxnID             string
	FirmName          string
	FirmFullName      string
	FirmCountry       string
	FirmCentre        string
	FirmSector        string
	CustomerType      string
	FirmAddress       string
	PhoneNumber       string
	Owner             string
	BOI_BOLNO         string
	BOI_RDC           string
	BOI_GM            string
	BaseCCY           string
	BaseCCYPair       string
	OurSortCode       string
	Category          string
	ImportIDs         []Onboard_CounterpartyImport
	Payees            []Onboard_Payee
	TradingEntity     string
	MandatedUsers     []Onboard_User
	FirmCountryList   []Lookup_Item
	FirmCentreList    []Lookup_Item
	FirmSectorList    []Lookup_Item
	CustomerTypeList  []Lookup_Item
	BaseCCYList       []Lookup_Item
	CategoryList      []Lookup_Item
	TradingEntityList []Lookup_Item
	SystemUserList    []Lookup_Item
	OriginList        []Lookup_Item
	OwnerList         []Lookup_Item
	BaseCCYPairList   []Lookup_Item
}

type Onboard_CounterpartyImport struct {
	Origin string
	//OriginList []dm.Lookup_Item
	ID string
}

type Onboard_Payee struct {
	PayeeID               string
	PayeeCCY              string
	PayeeAddress          string
	PayeeCountry          string
	PayeeBIC              string
	PayeeIBAN             string
	PayeeFullName         string
	PayeePhoneNumber      string
	PayeeSortCode         string
	BankSettlementAccount string
	PayeeReason           string
	PayeeAccountNo        string
	PayeeBankName         string
}

type Onboard_User struct {
	UserName         string
	UserFullName     string
	UserPhoneNumber  string
	UserEmail        string
	SystemUser       string
	UserDOB          string
	UserMaidenName   string
	UserPlaceOfBirth string
	UserMiddleName   string
}

const (
	CounterpartyOnboarding_Title = "CounterpartyOnboarding Contents"
	//CounterpartyOnboarding_SQLTable    = "template"
	//CounterpartyOnboarding_SQLSearchID = "ID"
	CounterpartyOnboarding_QueryString = "CounterpartyOnboardingID"
	///
	/// Handler Defintions
	///
	CounterpartyOnboarding_CounterpartyOnboardingList = "CounterpartyOnboarding_List"
	CounterpartyOnboarding_CounterpartyOnboardingView = "CounterpartyOnboarding_View"
	CounterpartyOnboarding_CounterpartyOnboardingEdit = "CounterpartyOnboarding_Edit"
	CounterpartyOnboarding_CounterpartyOnboardingNew  = "CounterpartyOnboarding_New"
	///
	/// Handler Monitor Paths
	///
	CounterpartyOnboarding_PathList   = "/CounterpartyOnboardingList/"
	CounterpartyOnboarding_PathView   = "/CounterpartyOnboardingView/"
	CounterpartyOnboarding_PathEdit   = "/CounterpartyOnboardingEdit/"
	CounterpartyOnboarding_PathNew    = "/CounterpartyOnboardingNew/"
	CounterpartyOnboarding_PathSave   = "/CounterpartyOnboardingSave/"
	CounterpartyOnboarding_PathDelete = "/CounterpartyOnboardingDelete/"
	///
	/// SQL Field Definitions
	///

	/// Definitions End
)
