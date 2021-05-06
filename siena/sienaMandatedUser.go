package siena

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

var sienaMandatedUserSQL = "MandatedUserKeyCounterpartyFirm, 	MandatedUserKeyCounterpartyCentre, 	MandatedUserKeyUserName, 	TelephoneNumber, 	EmailAddress, 	Active, 	FirstName, 	Surname, 	DateOfBirth, 	Postcode, 	NationalIDNo, 	PassportNo, 	Country, 	CountryName, 	FirmName, 	CentreName, 	Notify, 	SystemUser"

var sqlMDUMandatedUserKeyCounterpartyFirm, sqlMDUMandatedUserKeyCounterpartyCentre, sqlMDUMandatedUserKeyUserName, sqlMDUTelephoneNumber, sqlMDUEmailAddress, sqlMDUActive, sqlMDUFirstName, sqlMDUSurname, sqlMDUDateOfBirth, sqlMDUPostcode, sqlMDUNationalIDNo, sqlMDUPassportNo, sqlMDUCountry, sqlMDUCountryName, sqlMDUFirmName, sqlMDUCentreName, sqlMDUNotify, sqlMDUSystemUser sql.NullString

//sienaMandatedUserPage is cheese
type sienaMandatedUserListPage struct {
	UserMenu               []application.AppMenuItem
	UserRole               string
	UserNavi               string
	Title                  string
	PageTitle              string
	SienaMandatedUserCount int
	SienaMandatedUserList  []sienaMandatedUserItem
}

//sienaMandatedUserPage is cheese
type sienaMandatedUserPage struct {
	UserMenu                          []application.AppMenuItem
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
	YNList                            []sienaYNItem
	CBActive                          string
	CBNotify                          string
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
	CBActive                          string
	CBNotify                          string
}

func ListSienaMandatedUserHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "listSienaMandatedUser"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessage(inUTL)
	var returnList []sienaMandatedUserItem
	noItems, returnList, _ := getSienaMandatedUserList()

	pageSienaMandatedUserList := sienaMandatedUserListPage{
		UserMenu:               application.GetUserMenu(r),
		UserRole:               application.GetUserRole(r),
		UserNavi:               "NOT USED",
		Title:                  globals.ApplicationProperties["appname"],
		PageTitle:              globals.ApplicationProperties["appname"] + " - " + "Counterparty Mandates",
		SienaMandatedUserCount: noItems,
		SienaMandatedUserList:  returnList,
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaMandatedUserList)

}

func ViewSienaMandatedUserHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "viewSienaMandatedUser"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessage(inUTL)

	suID := application.GetURLparam(r, "SU")
	sfID := application.GetURLparam(r, "SF")
	scID := application.GetURLparam(r, "SC")
	noItems, returnRecord, _ := getSienaMandatedUser(suID, sfID, scID)
	fmt.Println("NoSienaItems", noItems, suID, sfID, scID)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	pageSienaMandatedUserList := sienaMandatedUserPage{
		UserMenu:                          application.GetUserMenu(r),
		UserRole:                          application.GetUserRole(r),
		UserNavi:                          "NOT USED",
		Title:                             globals.ApplicationProperties["appname"],
		PageTitle:                         globals.ApplicationProperties["appname"] + " - " + "Counterparty Mandate - View",
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
		CBActive:                          returnRecord.CBActive,
		CBNotify:                          returnRecord.CBNotify,
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaMandatedUserList)

}

func EditSienaMandatedUserHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "editSienaMandatedUser"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessage(inUTL)

	suID := application.GetURLparam(r, "SU")
	sfID := application.GetURLparam(r, "SF")
	scID := application.GetURLparam(r, "SC")
	noItems, returnRecord, _ := getSienaMandatedUser(suID, sfID, scID)
	fmt.Println("NoSienaItems", noItems, suID, sfID, scID)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	//Get Country List & Populate and Array of sienaCountryItem Items
	_, countryList, _ := getSienaCountryList()
	_, ynList, _ := getSienaYNList()

	//fmt.Println(displayList)

	pageSienaMandatedUserList := sienaMandatedUserPage{
		UserMenu:                          application.GetUserMenu(r),
		UserRole:                          application.GetUserRole(r),
		UserNavi:                          "NOT USED",
		Title:                             globals.ApplicationProperties["appname"],
		PageTitle:                         globals.ApplicationProperties["appname"] + " - " + "Counterparty Mandate - Edit",
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
		YNList:                            ynList,
		CBActive:                          returnRecord.CBActive,
		CBNotify:                          returnRecord.CBNotify,
	}

	fmt.Println(pageSienaMandatedUserList)

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaMandatedUserList)

}

func SaveSienaMandatedUserHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessageAction(inUTL, "Save", "")

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
	item.Notify = r.FormValue("Notify")
	item.SystemUser = r.FormValue("SystemUser")

	item.Action = "UPDATE"

	fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE
	var sienaFieldMandatedUserKeyCounterpartyFirm sienaKEYFIELD
	var sienaFieldMandatedUserKeyCounterpartyCentre sienaKEYFIELD
	var sienaFieldMandatedUserKeyUserName sienaKEYFIELD
	var sienaFieldTelephoneNumber sienaFIELD
	var sienaFieldEmailAddress sienaFIELD
	var sienaFieldActive sienaFIELD
	var sienaFieldFirstName sienaFIELD
	var sienaFieldSurname sienaFIELD
	var sienaFieldDateOfBirth sienaFIELD
	var sienaFieldPostcode sienaFIELD
	var sienaFieldNationalIDNo sienaFIELD
	var sienaFieldPassportNo sienaFIELD
	var sienaFieldCountry sienaFIELD
	var sienaFieldNotify sienaFIELD
	var sienaFieldSystemUser sienaFIELD

	// POitemPULATE THE XML FIELD NAMEs
	sienaFieldMandatedUserKeyCounterpartyFirm.Name = "MandatedUserKeyCounterpartyFirm"
	sienaFieldMandatedUserKeyCounterpartyCentre.Name = "MandatedUserKeyCounterpartyCentre"
	sienaFieldMandatedUserKeyUserName.Name = "MandatedUserKeyUserName"
	sienaFieldTelephoneNumber.Name = "TelephoneNumber"
	sienaFieldEmailAddress.Name = "EmailAddress"
	sienaFieldActive.Name = "Active"
	sienaFieldFirstName.Name = "FirstName"
	sienaFieldSurname.Name = "Surname"
	sienaFieldDateOfBirth.Name = "DateOfBirth"
	sienaFieldPostcode.Name = "Postcode"
	sienaFieldNationalIDNo.Name = "NationalIDNo"
	sienaFieldPassportNo.Name = "PassportNo"
	sienaFieldCountry.Name = "Country"
	sienaFieldNotify.Name = "Notify"
	sienaFieldSystemUser.Name = "SystemUser"

	// POitemPULATE THE XML FIELD values

	sienaFieldMandatedUserKeyCounterpartyFirm.Text = item.MandatedUserKeyCounterpartyFirm
	sienaFieldMandatedUserKeyCounterpartyCentre.Text = item.MandatedUserKeyCounterpartyCentre
	sienaFieldMandatedUserKeyUserName.Text = item.MandatedUserKeyUserName
	sienaFieldTelephoneNumber.Text = item.TelephoneNumber
	sienaFieldEmailAddress.Text = item.EmailAddress
	sienaFieldActive.Text = item.Active
	sienaFieldFirstName.Text = item.FirstName
	sienaFieldSurname.Text = item.Surname
	sienaFieldDateOfBirth.Text = item.DateOfBirth
	sienaFieldPostcode.Text = item.Postcode
	sienaFieldNationalIDNo.Text = item.NationalIDNo
	sienaFieldPassportNo.Text = item.PassportNo
	sienaFieldCountry.Text = item.Country
	sienaFieldNotify.Text = item.Notify
	sienaFieldSystemUser.Text = item.SystemUser

	// IGNORE
	var sienaKeyFields []sienaKEYFIELD
	var sienaFields []sienaFIELD
	// ADD THE FIELDS TO THE LISTS ABOVE
	sienaKeyFields = append(sienaKeyFields, sienaFieldMandatedUserKeyCounterpartyFirm)
	sienaKeyFields = append(sienaKeyFields, sienaFieldMandatedUserKeyCounterpartyCentre)
	sienaKeyFields = append(sienaKeyFields, sienaFieldMandatedUserKeyUserName)
	sienaFields = append(sienaFields, sienaFieldTelephoneNumber)
	sienaFields = append(sienaFields, sienaFieldEmailAddress)
	sienaFields = append(sienaFields, sienaFieldActive)
	sienaFields = append(sienaFields, sienaFieldFirstName)
	sienaFields = append(sienaFields, sienaFieldSurname)
	sienaFields = append(sienaFields, sienaFieldDateOfBirth)
	sienaFields = append(sienaFields, sienaFieldPostcode)
	sienaFields = append(sienaFields, sienaFieldNationalIDNo)
	sienaFields = append(sienaFields, sienaFieldPassportNo)
	sienaFields = append(sienaFields, sienaFieldCountry)
	sienaFields = append(sienaFields, sienaFieldNotify)
	sienaFields = append(sienaFields, sienaFieldSystemUser)

	// IGNORE
	sienaRecord := &sienaRECORD{KEYFIELD: sienaKeyFields, FIELD: sienaFields}
	var sienaRecords []sienaRECORD
	sienaRecords = append(sienaRecords, *sienaRecord)

	var sienaTable sienaTABLE
	sienaTable.Name = "MandatedUser"
	sienaTable.Classname = "com.eurobase.siena.data.MandatedUsers.MandatedUser"
	sienaTable.RECORD = sienaRecords

	var sienaTransaction sienaTRANSACTION
	sienaTransaction.Type = SIENAXMLIMPORT
	sienaTransaction.TABLE = sienaTable

	var sienaXMLContent sienaXML
	sienaXMLContent.TRANSACTIONS = sienaTransaction

	thisError := sienaDispatchStaticDataXML(sienaXMLContent)
	if thisError != nil {
		log.Println("Error in XML dispatch: ", thisError)
	}

	ListSienaMandatedUserHandler(w, r)

}

func NewSienaMandatedUserHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "newSienaMandatedUser"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessage(inUTL)

	//Get Country List & Populate and Array of sienaCountryItem Items

	_, countryList, _ := getSienaCountryList()
	_, firmList, _ := getSienaFirmList()
	_, centreList, _ := getSienaCentreList()
	_, ynList, _ := getSienaYNList()

	pageSienaMandatedUserList := sienaMandatedUserPage{
		UserMenu:                          application.GetUserMenu(r),
		UserRole:                          application.GetUserRole(r),
		UserNavi:                          "NOT USED",
		Title:                             globals.ApplicationProperties["appname"],
		PageTitle:                         globals.ApplicationProperties["appname"] + " - " + "Counterparty Mandate - New",
		ID:                                "NEW",
		MandatedUserKeyCounterpartyFirm:   "",
		MandatedUserKeyCounterpartyCentre: "",
		MandatedUserKeyUserName:           "",
		TelephoneNumber:                   "",
		EmailAddress:                      "",
		Active:                            "No",
		FirstName:                         "",
		Surname:                           "",
		Postcode:                          "",
		NationalIDNo:                      "",
		PassportNo:                        "",
		Country:                           "",
		CountryName:                       "",
		FirmName:                          "",
		CentreName:                        "",
		Notify:                            "Yes",
		SystemUser:                        "",
		Action:                            "NEW",
		CountryList:                       countryList,
		FirmList:                          firmList,
		CentreList:                        centreList,
		YNList:                            ynList,
		CBActive:                          "",
		CBNotify:                          "",
	}

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, pageSienaMandatedUserList)

}

// getSienaMandatedUserList read all employees
func getSienaMandatedUserList() (int, []sienaMandatedUserItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaMandatedUser;", sienaMandatedUserSQL, globals.SienaPropertiesDB["schema"])
	count, sienaMandatedUserList, _, _ := fetchSienaMandatedUserData(tsql)
	return count, sienaMandatedUserList, nil
}

// getSienaMandatedUserList read all employees
func getSienaMandatedUserListByCounterparty(idFirm string, idCentre string) (int, []sienaMandatedUserItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaMandatedUser WHERE MandatedUserKeyCounterpartyFirm='%s' AND MandatedUserKeyCounterpartyCentre='%s';", sienaMandatedUserSQL, globals.SienaPropertiesDB["schema"], idFirm, idCentre)
	//	fmt.Println("MS SQL:", tsql)
	count, sienaMandatedUserList, _, _ := fetchSienaMandatedUserData(tsql)
	return count, sienaMandatedUserList, nil
}

// getSienaMandatedUserList read all employees
func getSienaMandatedUser(suid string, sfid string, scid string) (int, sienaMandatedUserItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaMandatedUser WHERE MandatedUserKeyUserName='%s' AND MandatedUserKeyCounterpartyFirm='%s' AND MandatedUserKeyCounterpartyCentre='%s';", sienaMandatedUserSQL, globals.SienaPropertiesDB["schema"], suid, sfid, scid)
	_, _, sienaMandatedUser, _ := fetchSienaMandatedUserData(tsql)
	return 1, sienaMandatedUser, nil
}

// getSienaMandatedUserList read all employees
func putSienaMandatedUser(updateItem sienaMandatedUserItem) error {

	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(globals.SienaPropertiesDB["schema"])
	fmt.Println(updateItem)
	return nil
}

// fetchSienaMandatedUserData read all employees
func fetchSienaMandatedUserData(tsql string) (int, []sienaMandatedUserItem, sienaMandatedUserItem, error) {

	var sienaMandatedUser sienaMandatedUserItem
	var sienaMandatedUserList []sienaMandatedUserItem

	rows, err := globals.SienaDB.Query(tsql)
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
		sienaMandatedUser.Active = sienaYN(sqlMDUActive.String)
		sienaMandatedUser.FirstName = sqlMDUFirstName.String
		sienaMandatedUser.Surname = sqlMDUSurname.String
		sienaMandatedUser.DateOfBirth = application.SqlDateToHTMLDate(sqlMDUDateOfBirth.String)
		sienaMandatedUser.Postcode = sqlMDUPostcode.String
		sienaMandatedUser.NationalIDNo = sqlMDUNationalIDNo.String
		sienaMandatedUser.PassportNo = sqlMDUPassportNo.String
		sienaMandatedUser.Country = sqlMDUCountry.String
		sienaMandatedUser.CountryName = sqlMDUCountryName.String
		sienaMandatedUser.FirmName = sqlMDUFirmName.String
		sienaMandatedUser.CentreName = sqlMDUCentreName.String
		sienaMandatedUser.Notify = sienaYN(sqlMDUNotify.String)
		sienaMandatedUser.SystemUser = sqlMDUSystemUser.String
		sienaMandatedUser.CBActive = setChecked(sqlMDUActive.String)
		sienaMandatedUser.CBNotify = setChecked(sqlMDUNotify.String)

		sienaMandatedUserList = append(sienaMandatedUserList, sienaMandatedUser)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaMandatedUserList, sienaMandatedUser, nil
}
