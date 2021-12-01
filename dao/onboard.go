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

func Onboard_Test() {
	xx := CounterpartyOnboarding_Build()
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

func CounterpartyOnboarding_Build() dm.CounterpartyOnboarding {
	var onboarding_Message dm.CounterpartyOnboarding
	onboarding_Message.FirmName = "Telia"
	onboarding_Message.FirmFullName = "Telia International"
	onboarding_Message.FirmCountry = "GBR"
	onboarding_Message.FirmCountryList = Country_GetLookup()
	onboarding_Message.FirmCentre = "WEB"
	///_, centrelist, _ := Centre_GetList()
	onboarding_Message.FirmCentreList = Centre_GetLookup()
	onboarding_Message.FirmSector = "Telephony"
	onboarding_Message.FirmSectorList = Sector_GetLookup()
	onboarding_Message.CustomerType = "Corporate"
	onboarding_Message.CustomerTypeList = CounterpartyType_GetLookup()
	onboarding_Message.FirmAddress = "22 Bedford Place, Leeds"
	onboarding_Message.PhoneNumber = "+44 7837 8272"
	onboarding_Message.Owner = "SalesDesk"
	onboarding_Message.OwnerList = Owner_GetLookup()
	onboarding_Message.BOI_BOLNO = "123456789"
	onboarding_Message.BOI_RDC = "23456789"
	onboarding_Message.BOI_GM = "34567890"
	onboarding_Message.BaseCCY = "GBP"
	ccyList := Currency_GetLookup()
	onboarding_Message.BaseCCYPair = "GBPUSD"
	onboarding_Message.BaseCCYList = ccyList
	onboarding_Message.OurSortCode = "60-09-09"
	onboarding_Message.Category = "Retail Client"
	onboarding_Message.CategoryList = Counterparty_MiFIDCategory_GetLookup()
	onboarding_Message.ImportIDs = []dm.Onboard_CounterpartyImport{}
	onboarding_Message.Payees = []dm.Onboard_Payee{}
	onboarding_Message.TradingEntity = "Sales Desk"
	onboarding_Message.TradingEntityList = SalesDesk_GetLookup()
	onboarding_Message.MandatedUsers = []dm.Onboard_User{}
	onboarding_Message.OriginList = Counterparty_Origin_GetLookup()
	onboarding_Message.SystemUserList = Counterparty_SystemUser_GetLookup()

	var cpi dm.Onboard_CounterpartyImport
	cpi.Origin = "T24"

	cpi.ID = "500300"
	onboarding_Message.ImportIDs = append(onboarding_Message.ImportIDs, cpi)
	var payee dm.Onboard_Payee
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
	onboarding_Message.Payees = append(onboarding_Message.Payees, payee)
	var user dm.Onboard_User
	user.UserName = "tim@apple.com"
	user.UserFullName = "Tim Cook"
	user.UserPhoneNumber = "+1 123456789"
	user.UserEmail = "tim@apple.com"
	user.SystemUser = onboarding_Message.SystemUserList[0].ID
	user.UserDOB = "1960-11-01"
	user.UserMaidenName = "Brown"
	user.UserPlaceOfBirth = "Mobile, Alabama"
	user.UserMiddleName = "Donald"
	onboarding_Message.MandatedUsers = append(onboarding_Message.MandatedUsers, user)

	var user2 dm.Onboard_User
	user2.UserName = "steve@apple.com"
	user2.UserFullName = "Steve Wozniak"
	user2.UserPhoneNumber = "+1 123456789"
	user2.UserEmail = "steve@apple.com"
	user2.SystemUser = onboarding_Message.SystemUserList[1].ID
	user2.UserDOB = "1950-08-11"
	user2.UserMaidenName = "Plum"
	user2.UserPlaceOfBirth = "San Jose, California"
	user2.UserMiddleName = "Gary"
	onboarding_Message.MandatedUsers = append(onboarding_Message.MandatedUsers, user2)

	id := uuid.New().String()
	onboarding_Message.TxnID = id
	return onboarding_Message
}
