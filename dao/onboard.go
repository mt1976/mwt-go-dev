package dao

import (
	"encoding/json"
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

type Onboard_Message struct {
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
	OurSortCode       string
	Category          string
	ImportIDs         []Onboard_CounterpartyImport
	Payees            []Onboard_Payee
	TradingEntity     string
	MandatedUsers     []Onboard_User
	FirmCountryList   []dm.Lookup_Item
	FirmCentreList    []dm.Lookup_Item
	FirmSectorList    []dm.Lookup_Item
	CustomerTypeList  []dm.Lookup_Item
	BaseCCYList       []dm.Lookup_Item
	CategoryList      []dm.Lookup_Item
	TradingEntityList []dm.Lookup_Item
	SystemUserList    []dm.Lookup_Item
	OriginList        []dm.Lookup_Item
	OwnerList         []dm.Lookup_Item
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
	UserFullName     string
	UserPhoneNumber  string
	UserEmail        string
	SystemUser       string
	UserDOB          string
	UserMaidenName   string
	UserPlaceOfBirth string
	UserMiddleName   string
}

func Onboard_Test() {
	xx := Onboard_Build()
	xx.FirmSector = "FirmSector"
	spew.Dump(xx)

	file, _ := json.MarshalIndent(xx, "", " ")

	//Get a uuid
	id := uuid.New().String()
	//Get current path
	path := "/Volumes/External HD/matttownsend/Documents/GitHub/mwt-go-dev/data/onboarding/"
	filename := path + id + ".json"

	_ = ioutil.WriteFile(filename, file, 0644)

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
	onboard_message.CustomerTypeList = CounterpartyType_GetLookup()
	onboard_message.FirmAddress = "FirmAddress"
	onboard_message.PhoneNumber = "PhoneNumber"
	onboard_message.Owner = "Owner"
	onboard_message.OwnerList = Owner_GetLookup()
	onboard_message.BOI_BOLNO = "BOI_BOLNO"
	onboard_message.BOI_RDC = "BOI_RDC"
	onboard_message.BOI_GM = "BOI_GM"
	onboard_message.BaseCCY = "BaseCCY"
	ccyList := Currency_GetLookup()
	onboard_message.BaseCCYList = ccyList
	onboard_message.OurSortCode = "OurSortCode"
	onboard_message.Category = "Category"
	onboard_message.CategoryList = Counterparty_MiFIDCategory_GetLookup()
	onboard_message.ImportIDs = []Onboard_CounterpartyImport{}
	onboard_message.Payees = []Onboard_Payee{}
	onboard_message.TradingEntity = "TradingEntity"
	onboard_message.TradingEntityList = SalesDesk_GetLookup()
	onboard_message.MandatedUsers = []Onboard_User{}
	onboard_message.OriginList = Counterparty_Origin_GetLookup()
	onboard_message.SystemUserList = Counterparty_SystemUser_GetLookup()

	var cpi Onboard_CounterpartyImport
	cpi.Origin = "Origin"

	cpi.ID = "ID"
	onboard_message.ImportIDs = append(onboard_message.ImportIDs, cpi)
	var payee Onboard_Payee
	payee.PayeeID = "PayeeID"
	payee.PayeeCCY = ccyList[0].ID
	payee.PayeeAddress = "PayeeAddress"
	payee.PayeeCountry = onboard_message.FirmCountryList[0].ID

	payee.PayeeBIC = "PayeeBIC"
	payee.PayeeIBAN = "PayeeIBAN"
	payee.PayeeFullName = "PayeeFullName"
	payee.PayeePhoneNumber = "PayeePhoneNumber"
	payee.PayeeSortCode = "PayeeSortCode"
	payee.BankSettlementAccount = "BankSettlementAccount"
	payee.PayeeReason = "PayeeReason"
	payee.PayeeAccountNo = "PayeeAccountNo"
	payee.PayeeBankName = "PayeeBankName"
	onboard_message.Payees = append(onboard_message.Payees, payee)
	var user Onboard_User
	user.UserFullName = "UserFullName"
	user.UserPhoneNumber = "UserPhoneNumber"
	user.UserEmail = "UserEmail"
	user.SystemUser = onboard_message.SystemUserList[0].ID

	user.UserDOB = "UserDOB"
	user.UserMaidenName = "UserMaidenName"
	user.UserPlaceOfBirth = "UserPlaceOfBirth"
	user.UserMiddleName = "UserMiddleName"
	onboard_message.MandatedUsers = append(onboard_message.MandatedUsers, user)

	return onboard_message
}
