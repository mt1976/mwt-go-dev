package dao

import (
	"github.com/davecgh/go-spew/spew"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

type Onboard_Message struct {
	FirmName          string
	FirmFullName      string
	FirmCountry       string
	FirmCountryList   []dm.Lookup_Item
	FirmCentre        string
	FirmCentreList    []dm.Lookup_Item
	FirmSector        string
	FirmSectorList    []dm.Lookup_Item
	CustomerType      string
	CustomerTypeList  []dm.Lookup_Item
	FirmAddress       string
	PhoneNumber       string
	Owner             string
	OwnerList         []dm.Lookup_Item
	BOI_BOLNO         string
	BOI_RDC           string
	BOI_GM            string
	BaseCCY           string
	BaseCCYList       []dm.Lookup_Item
	OurSortCode       string
	Category          string
	CategoryList      []dm.Lookup_Item
	ImportIDs         []Onboard_CounterpartyImport
	Payees            []Onboard_Payee
	TradingEntity     string
	TradingEntityList []dm.Lookup_Item
	MandatedUsers     []Onboard_User
}

type Onboard_CounterpartyImport struct {
	Origin     string
	OriginList []dm.Lookup_Item
	ID         string
}

type Onboard_Payee struct {
	PayeeID               string
	PayeeCCY              string
	PayeeList             []dm.Lookup_Item
	PayeeAddress          string
	PayeeCountry          string
	PayeeCountryList      []dm.Lookup_Item
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
	UserFullName     string
	UserPhoneNumber  string
	UserEmail        string
	SystemUser       string
	SystemUserList   []dm.Lookup_Item
	UserDOB          string
	UserMaidenName   string
	UserPlaceOfBirth string
	UserMiddleName   string
}

func Onboard_Test() {
	xx := Onboard_Build()
	spew.Dump(xx)
}

func Onboard_Build() Onboard_Message {
	var onboard_message Onboard_Message
	onboard_message.FirmName = "FirmName"
	onboard_message.FirmFullName = "FirmFullName"
	onboard_message.FirmCountry = "FirmCountry"

	onboard_message.FirmCountryList = Country_GetLookup()
	onboard_message.FirmCentre = "FirmCentre"
	///_, centrelist, _ := Centre_GetList()
	onboard_message.FirmCentreList = Centre_GetLookup()
	onboard_message.FirmSector = "FirmSector"
	onboard_message.FirmSectorList = Sector_GetLookup()
	onboard_message.CustomerType = "CustomerType"
	onboard_message.CustomerTypeList = []dm.Lookup_Item{}
	onboard_message.FirmAddress = "FirmAddress"
	onboard_message.PhoneNumber = "PhoneNumber"
	onboard_message.Owner = "Owner"
	onboard_message.OwnerList = []dm.Lookup_Item{}
	onboard_message.BOI_BOLNO = "BOI_BOLNO"
	onboard_message.BOI_RDC = "BOI_RDC"
	onboard_message.BOI_GM = "BOI_GM"
	onboard_message.BaseCCY = "BaseCCY"

	onboard_message.BaseCCYList = []dm.Lookup_Item{}
	onboard_message.OurSortCode = "OurSortCode"
	onboard_message.Category = "Category"
	onboard_message.CategoryList = []dm.Lookup_Item{}
	onboard_message.ImportIDs = []Onboard_CounterpartyImport{}
	onboard_message.Payees = []Onboard_Payee{}
	onboard_message.TradingEntity = "TradingEntity"
	onboard_message.TradingEntityList = []dm.Lookup_Item{}
	onboard_message.MandatedUsers = []Onboard_User{}

	return onboard_message
}
