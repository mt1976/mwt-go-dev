package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

var sienaMandatedUserSQL = "MandatedUserKeyCounterpartyFirm, 	MandatedUserKeyCounterpartyCentre, 	MandatedUserKeyUserName, 	TelephoneNumber, 	EmailAddress, 	Active, 	FirstName, 	Surname, 	DateOfBirth, 	Postcode, 	NationalIDNo, 	PassportNo, 	Country, 	CountryName, 	FirmName, 	CentreName, 	Notify, 	SystemUser"

var sqlMDUMandatedUserKeyCounterpartyFirm, sqlMDUMandatedUserKeyCounterpartyCentre, sqlMDUMandatedUserKeyUserName, sqlMDUTelephoneNumber, sqlMDUEmailAddress, sqlMDUActive, sqlMDUFirstName, sqlMDUSurname, sqlMDUDateOfBirth, sqlMDUPostcode, sqlMDUNationalIDNo, sqlMDUPassportNo, sqlMDUCountry, sqlMDUCountryName, sqlMDUFirmName, sqlMDUCentreName, sqlMDUNotify, sqlMDUSystemUser sql.NullString

//sienaMandatedUserPage is cheese
type sienaMandatedUserListPage struct {
	UserRole               string
	UserNavi               string
	Title                  string
	PageTitle              string
	SienaMandatedUserCount int
	SienaMandatedUserList  []sienaMandatedUserItem
}

//sienaMandatedUserPage is cheese
type sienaMandatedUserPage struct {
	UserRole                          string
	UserNavi                          string
	Title                             string
	PageTitle                         string
	ID                                string
	MandatedUserKeyCounterpartyFirm   string
	MandatedUserKeyCounterpartyCentre string
	MandatedUserKeyUserName           string
	TelephoneNumber                   string
	EmailAddress                      string
	Active                            string
	FirstName                         string
	Surname                           string
	DateOfBirth                       string
	Postcode                          string
	NationalIDNo                      string
	PassportNo                        string
	Country                           string
	CountryName                       string
	FirmName                          string
	CentreName                        string
	Notify                            string
	SystemUser                        string
	Action                            string
	CountryList                       []sienaCountryItem
	FirmList                          []sienaFirmItem
	CentreList                        []sienaCentreItem
	CounterpartyName                  string
}

//sienaMandatedUserItem is cheese
type sienaMandatedUserItem struct {
	MandatedUserKeyCounterpartyFirm   string
	MandatedUserKeyCounterpartyCentre string
	MandatedUserKeyUserName           string
	TelephoneNumber                   string
	EmailAddress                      string
	Active                            string
	FirstName                         string
	Surname                           string
	DateOfBirth                       string
	Postcode                          string
	NationalIDNo                      string
	PassportNo                        string
	Country                           string
	CountryName                       string
	FirmName                          string
	CentreName                        string
	Notify                            string
	SystemUser                        string
	Action                            string
	CounterpartyName                  string
}

func listSienaMandatedUserHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "listSienaMandatedUser"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaMandatedUserItem
	noItems, returnList, _ := getSienaMandatedUserList(thisConnection)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaMandatedUserList := sienaMandatedUserListPage{
		UserRole:               gUserRole,
		UserNavi:               gUserNavi,
		Title:                  wctProperties["appname"],
		PageTitle:              "List Siena MandatedUsers",
		SienaMandatedUserCount: noItems,
		SienaMandatedUserList:  returnList,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaMandatedUserList)

}

func viewSienaMandatedUserHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "viewSienaMandatedUser"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaMandatedUserItem
	suID := getURLparam(r, "SU")
	sfID := getURLparam(r, "SF")
	scID := getURLparam(r, "SC")
	noItems, returnRecord, _ := getSienaMandatedUser(thisConnection, suID, sfID, scID)
	fmt.Println("NoSienaItems", noItems, suID, sfID, scID)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	pageSienaMandatedUserList := sienaMandatedUserPage{
		UserRole:                          gUserRole,
		UserNavi:                          gUserNavi,
		Title:                             wctProperties["appname"],
		PageTitle:                         "View Siena MandatedUser",
		MandatedUserKeyCounterpartyFirm:   returnRecord.MandatedUserKeyCounterpartyFirm,
		MandatedUserKeyCounterpartyCentre: returnRecord.MandatedUserKeyCounterpartyCentre,
		MandatedUserKeyUserName:           returnRecord.MandatedUserKeyUserName,
		TelephoneNumber:                   returnRecord.TelephoneNumber,
		EmailAddress:                      returnRecord.EmailAddress,
		Active:                            returnRecord.Active,
		FirstName:                         returnRecord.FirstName,
		Surname:                           returnRecord.Surname,
		DateOfBirth:                       returnRecord.DateOfBirth,
		Postcode:                          returnRecord.Postcode,
		NationalIDNo:                      returnRecord.NationalIDNo,
		PassportNo:                        returnRecord.PassportNo,
		Country:                           returnRecord.Country,
		CountryName:                       returnRecord.CountryName,
		FirmName:                          returnRecord.FirmName,
		CentreName:                        returnRecord.CentreName,
		Notify:                            returnRecord.Notify,
		SystemUser:                        returnRecord.SystemUser,
		Action:                            "",
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaMandatedUserList)

}

func editSienaMandatedUserHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "editSienaMandatedUser"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	suID := getURLparam(r, "SU")
	sfID := getURLparam(r, "SF")
	scID := getURLparam(r, "SC")
	noItems, returnRecord, _ := getSienaMandatedUser(thisConnection, suID, sfID, scID)
	fmt.Println("NoSienaItems", noItems, suID, sfID, scID)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	//Get Country List & Populate and Array of sienaCountryItem Items
	_, countryList, _ := getSienaCountryList(thisConnection)

	//fmt.Println(displayList)

	pageSienaMandatedUserList := sienaMandatedUserPage{
		UserRole:                          gUserRole,
		UserNavi:                          gUserNavi,
		Title:                             wctProperties["appname"],
		PageTitle:                         "View Siena MandatedUser",
		MandatedUserKeyCounterpartyFirm:   returnRecord.MandatedUserKeyCounterpartyFirm,
		MandatedUserKeyCounterpartyCentre: returnRecord.MandatedUserKeyCounterpartyCentre,
		MandatedUserKeyUserName:           returnRecord.MandatedUserKeyUserName,
		TelephoneNumber:                   returnRecord.TelephoneNumber,
		EmailAddress:                      returnRecord.EmailAddress,
		Active:                            returnRecord.Active,
		FirstName:                         returnRecord.FirstName,
		Surname:                           returnRecord.Surname,
		DateOfBirth:                       returnRecord.DateOfBirth,
		Postcode:                          returnRecord.Postcode,
		NationalIDNo:                      returnRecord.NationalIDNo,
		PassportNo:                        returnRecord.PassportNo,
		Country:                           returnRecord.Country,
		CountryName:                       returnRecord.CountryName,
		FirmName:                          returnRecord.FirmName,
		CentreName:                        returnRecord.CentreName,
		Notify:                            returnRecord.Notify,
		SystemUser:                        returnRecord.SystemUser,
		Action:                            "",
		CountryList:                       countryList,
	}

	fmt.Println(pageSienaMandatedUserList)

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaMandatedUserList)

}

func saveSienaMandatedUserHandler(w http.ResponseWriter, r *http.Request) {

	sienaProperties := getProperties(cSIENACONFIG)
	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL, " : Save")

	var item sienaMandatedUserItem

	item.MandatedUserKeyCounterpartyFirm = r.FormValue("MandatedUserKeyCounterpartyFirm")
	item.MandatedUserKeyCounterpartyCentre = r.FormValue("MandatedUserKeyCounterpartyCentre")
	item.MandatedUserKeyUserName = r.FormValue("MandatedUserKeyUserName")
	item.TelephoneNumber = r.FormValue("TelephoneNumber")
	item.EmailAddress = r.FormValue("EmailAddress")
	item.Active = r.FormValue("Active")
	item.FirstName = r.FormValue("FirstName")
	item.Surname = r.FormValue("Surname")
	item.DateOfBirth = r.FormValue("DateOfBirth")
	item.Postcode = r.FormValue("Postcode")
	item.NationalIDNo = r.FormValue("NationalIDNo")
	item.PassportNo = r.FormValue("PassportNo")
	item.Country = r.FormValue("Country")
	item.CountryName = r.FormValue("CountryName")
	item.FirmName = r.FormValue("FirmName")
	item.CentreName = r.FormValue("CentreName")
	item.Notify = r.FormValue("Notify")
	item.SystemUser = r.FormValue("SystemUser")

	item.Action = "UPDATE"

	fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE
	var sFldCode sienaKEYFIELD
	var sFldName sienaFIELD
	var sFldCountry sienaFIELD

	// POPULATE THE XML FIELDS
	sFldCode.Name = "MandatedUserKeyUserName"
	sFldCode.Text = item.MandatedUserKeyUserName

	sFldName.Name = "MandatedUserKeyCounterpartyFirm"
	sFldName.Text = item.MandatedUserKeyCounterpartyFirm

	sFldCountry.Name = "MandatedUserKeyCounterpartyCentre"
	sFldCountry.Text = item.MandatedUserKeyCounterpartyCentre

	// IGNORE
	var sienaKeyFields []sienaKEYFIELD
	var sienaFields []sienaFIELD
	// ADD THE FIELDS TO THE LISTS ABOVE
	sienaKeyFields = append(sienaKeyFields, sFldCode)
	sienaFields = append(sienaFields, sFldName)
	sienaFields = append(sienaFields, sFldCountry)
	// IGNORE
	sienaRecord := &sienaRECORD{KEYFIELD: sienaKeyFields, FIELD: sienaFields}
	var sienaRecords []sienaRECORD
	sienaRecords = append(sienaRecords, *sienaRecord)

	var sienaTable sienaTABLE
	sienaTable.Name = "MandatedUser"
	sienaTable.Classname = "com.eurobase.siena.data.MandatedUsers.MandatedUser"
	sienaTable.RECORD = sienaRecords

	var sienaTransaction sienaTRANSACTION
	sienaTransaction.Type = "IMPORT"
	sienaTransaction.TABLE = sienaTable

	var sienaXMLContent sienaXML
	sienaXMLContent.TRANSACTIONS = sienaTransaction

	preparedXML, _ := xml.Marshal(sienaXMLContent)
	fmt.Println("PreparedXML", string(preparedXML))

	staticImporterPath := sienaProperties["static_in"]
	fileID := uuid.New()
	pwd, _ := os.Getwd()
	fileName := pwd + staticImporterPath + "/" + fileID.String() + ".xml"
	fmt.Println(fileName)

	err := ioutil.WriteFile(fileName, preparedXML, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	listSienaMandatedUserHandler(w, r)

}

func newSienaMandatedUserHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "newSienaMandatedUser"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	//Get Country List & Populate and Array of sienaCountryItem Items
	thisConnection, _ := sienaConnect()

	_, countryList, _ := getSienaCountryList(thisConnection)
	_, firmList, _ := getSienaFirmList(thisConnection)
	_, centreList, _ := getSienaCentreList(thisConnection)

	pageSienaMandatedUserList := sienaMandatedUserPage{
		UserRole:                          gUserRole,
		UserNavi:                          gUserNavi,
		Title:                             wctProperties["appname"],
		PageTitle:                         "View Siena MandatedUser",
		ID:                                "NEW",
		MandatedUserKeyCounterpartyFirm:   "",
		MandatedUserKeyCounterpartyCentre: "",
		MandatedUserKeyUserName:           "",
		TelephoneNumber:                   "",
		EmailAddress:                      "",
		Active:                            "",
		FirstName:                         "",
		Surname:                           "",
		Postcode:                          "",
		NationalIDNo:                      "",
		PassportNo:                        "",
		Country:                           "",
		CountryName:                       "",
		FirmName:                          "",
		CentreName:                        "",
		Notify:                            "",
		SystemUser:                        "",
		Action:                            "NEW",
		CountryList:                       countryList,
		FirmList:                          firmList,
		CentreList:                        centreList,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaMandatedUserList)

}

// getSienaMandatedUserList read all employees
func getSienaMandatedUserList(db *sql.DB) (int, []sienaMandatedUserItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaMandatedUser;", sienaMandatedUserSQL, mssqlConfig["schema"])
	count, sienaMandatedUserList, _, _ := fetchSienaMandatedUserData(db, tsql)
	return count, sienaMandatedUserList, nil
}

// getSienaMandatedUserList read all employees
func getSienaMandatedUserListByCounterparty(db *sql.DB, idFirm string, idCentre string) (int, []sienaMandatedUserItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaMandatedUser WHERE MandatedUserKeyCounterpartyFirm='%s' AND MandatedUserKeyCounterpartyCentre='%s';", sienaMandatedUserSQL, mssqlConfig["schema"], idFirm, idCentre)
	//	fmt.Println("MS SQL:", tsql)
	count, sienaMandatedUserList, _, _ := fetchSienaMandatedUserData(db, tsql)
	return count, sienaMandatedUserList, nil
}

// getSienaMandatedUserList read all employees
func getSienaMandatedUser(db *sql.DB, suid string, sfid string, scid string) (int, sienaMandatedUserItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaMandatedUser WHERE MandatedUserKeyUserName='%s' AND MandatedUserKeyCounterpartyFirm='%s' AND MandatedUserKeyCounterpartyCentre='%s';", sienaMandatedUserSQL, mssqlConfig["schema"], suid, sfid, scid)
	_, _, sienaMandatedUser, _ := fetchSienaMandatedUserData(db, tsql)
	return 1, sienaMandatedUser, nil
}

// getSienaMandatedUserList read all employees
func putSienaMandatedUser(db *sql.DB, updateItem sienaMandatedUserItem) error {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(mssqlConfig["schema"])
	fmt.Println(updateItem)
	return nil
}

// fetchSienaMandatedUserData read all employees
func fetchSienaMandatedUserData(db *sql.DB, tsql string) (int, []sienaMandatedUserItem, sienaMandatedUserItem, error) {

	var sienaMandatedUser sienaMandatedUserItem
	var sienaMandatedUserList []sienaMandatedUserItem

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaMandatedUser, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlMDUMandatedUserKeyCounterpartyFirm, &sqlMDUMandatedUserKeyCounterpartyCentre, &sqlMDUMandatedUserKeyUserName, &sqlMDUTelephoneNumber, &sqlMDUEmailAddress, &sqlMDUActive, &sqlMDUFirstName, &sqlMDUSurname, &sqlMDUDateOfBirth, &sqlMDUPostcode, &sqlMDUNationalIDNo, &sqlMDUPassportNo, &sqlMDUCountry, &sqlMDUCountryName, &sqlMDUFirmName, &sqlMDUCentreName, &sqlMDUNotify, &sqlMDUSystemUser)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaMandatedUser, err
		}

		sienaMandatedUser.MandatedUserKeyCounterpartyFirm = sqlMDUMandatedUserKeyCounterpartyFirm.String
		sienaMandatedUser.MandatedUserKeyCounterpartyCentre = sqlMDUMandatedUserKeyCounterpartyCentre.String
		sienaMandatedUser.MandatedUserKeyUserName = sqlMDUMandatedUserKeyUserName.String
		sienaMandatedUser.TelephoneNumber = sqlMDUTelephoneNumber.String
		sienaMandatedUser.EmailAddress = sqlMDUEmailAddress.String
		sienaMandatedUser.Active = sqlMDUActive.String
		sienaMandatedUser.FirstName = sqlMDUFirstName.String
		sienaMandatedUser.Surname = sqlMDUSurname.String
		sienaMandatedUser.DateOfBirth = sqlDateToHTMLDate(sqlMDUDateOfBirth.String)
		sienaMandatedUser.Postcode = sqlMDUPostcode.String
		sienaMandatedUser.NationalIDNo = sqlMDUNationalIDNo.String
		sienaMandatedUser.PassportNo = sqlMDUPassportNo.String
		sienaMandatedUser.Country = sqlMDUCountry.String
		sienaMandatedUser.CountryName = sqlMDUCountryName.String
		sienaMandatedUser.FirmName = sqlMDUFirmName.String
		sienaMandatedUser.CentreName = sqlMDUCentreName.String
		sienaMandatedUser.Notify = sqlMDUNotify.String
		sienaMandatedUser.SystemUser = sqlMDUSystemUser.String

		sienaMandatedUserList = append(sienaMandatedUserList, sienaMandatedUser)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaMandatedUserList, sienaMandatedUser, nil
}
