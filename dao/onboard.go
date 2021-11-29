package dao

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

type Onboard_Message struct {
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
	BaseCCYPairList   []dm.Lookup_Item
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

func Onboard_Test() {
	xx := Onboard_Build()
	xx.FirmSector = "FirmSector"
	spew.Dump(xx)

	file, _ := json.MarshalIndent(xx, "", " ")

	//Get a uuid
	//Get current path
	path := "/Volumes/External HD/matttownsend/Documents/GitHub/mwt-go-dev/data/onboarding/"
	filename := path + xx.TxnID + ".json"
	filenameXML := path + xx.TxnID + ".xml"

	_ = ioutil.WriteFile(filename, file, 0644)

	fp := path + "/templates/" + "a982234_Strawberry.xml"
	//read template from file

	t, err := template.ParseFiles(fp)
	if err != nil {
		logs.Error("Load Template :", err)
	}
	f, err := os.Create(filenameXML)
	if err != nil {
		logs.Error("Create file: ", err)

	}

	err2 := t.Execute(f, xx)
	if err2 != nil {
		logs.Error("Process Template", err2)
	}
	f.Close()
	logs.Created(f.Name())

}

func Onboard_Build() Onboard_Message {
	var onboard_message Onboard_Message
	onboard_message.FirmName = "Telia"
	onboard_message.FirmFullName = "Telia International"
	onboard_message.FirmCountry = "GBR"
	onboard_message.FirmCountryList = Country_GetLookup()
	onboard_message.FirmCentre = "WEB"
	///_, centrelist, _ := Centre_GetList()
	onboard_message.FirmCentreList = Centre_GetLookup()
	onboard_message.FirmSector = "Telephony"
	onboard_message.FirmSectorList = Sector_GetLookup()
	onboard_message.CustomerType = "Corporate"
	onboard_message.CustomerTypeList = CounterpartyType_GetLookup()
	onboard_message.FirmAddress = "22 Bedford Place, Leeds"
	onboard_message.PhoneNumber = "+44 7837 8272"
	onboard_message.Owner = "SalesDesk"
	onboard_message.OwnerList = Owner_GetLookup()
	onboard_message.BOI_BOLNO = "123456789"
	onboard_message.BOI_RDC = "23456789"
	onboard_message.BOI_GM = "34567890"
	onboard_message.BaseCCY = "GBP"
	ccyList := Currency_GetLookup()
	onboard_message.BaseCCYPair = "GBPUSD"
	onboard_message.BaseCCYList = ccyList
	onboard_message.OurSortCode = "60-09-09"
	onboard_message.Category = "Retail Client"
	onboard_message.CategoryList = Counterparty_MiFIDCategory_GetLookup()
	onboard_message.ImportIDs = []Onboard_CounterpartyImport{}
	onboard_message.Payees = []Onboard_Payee{}
	onboard_message.TradingEntity = "Sales Desk"
	onboard_message.TradingEntityList = SalesDesk_GetLookup()
	onboard_message.MandatedUsers = []Onboard_User{}
	onboard_message.OriginList = Counterparty_Origin_GetLookup()
	onboard_message.SystemUserList = Counterparty_SystemUser_GetLookup()

	var cpi Onboard_CounterpartyImport
	cpi.Origin = "Origin"

	cpi.ID = "ID"
	onboard_message.ImportIDs = append(onboard_message.ImportIDs, cpi)
	var payee Onboard_Payee
	payee.PayeeID = "APPLEUSD"
	payee.PayeeCCY = "USD"
	payee.PayeeAddress = "1 Apple Way, Cupertino"
	payee.PayeeCountry = "USA"

	payee.PayeeBIC = "CSCHUS6SXXX"
	payee.PayeeIBAN = "GB94BARC10201530093459"
	payee.PayeeFullName = "Apple USA"
	payee.PayeePhoneNumber = "+1 2345678"
	payee.PayeeSortCode = "99-99-99"
	payee.BankSettlementAccount = "false"
	payee.PayeeReason = "Goods & Services"
	payee.PayeeAccountNo = "987654321"
	payee.PayeeBankName = "Charles Schwabb"
	onboard_message.Payees = append(onboard_message.Payees, payee)
	var user Onboard_User
	user.UserName = "tim@apple.com"
	user.UserFullName = "Tim Cook"
	user.UserPhoneNumber = "+1 123456789"
	user.UserEmail = "tim@apple.com"
	user.SystemUser = onboard_message.SystemUserList[0].ID

	user.UserDOB = "1960-11-1"
	user.UserMaidenName = "Brown"
	user.UserPlaceOfBirth = "Mobile, Alabama"
	user.UserMiddleName = "Donald"
	onboard_message.MandatedUsers = append(onboard_message.MandatedUsers, user)

	var user2 Onboard_User
	user2.UserName = "steve@apple.com"
	user2.UserFullName = "Steve Wozniak"
	user2.UserPhoneNumber = "+1 123456789"
	user2.UserEmail = "steve@apple.com"
	user2.SystemUser = onboard_message.SystemUserList[1].ID

	user2.UserDOB = "1950-08-11"
	user2.UserMaidenName = "Steve Woznia"
	user2.UserPlaceOfBirth = "San Jose, California"
	user2.UserMiddleName = "Gary"
	onboard_message.MandatedUsers = append(onboard_message.MandatedUsers, user2)

	id := uuid.New().String()
	onboard_message.TxnID = id
	return onboard_message
}
