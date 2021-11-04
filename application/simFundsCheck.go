package application

import (
	"encoding/xml"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
)

// Defines the Fields to Fetch from SQL
/*
var simFundsCheckSQL = "id, 	name, 	staticIn, 	staticout, 	txnin, 	txnout, 	fundscheckin, 	fundscheckout, 	_created, 	_who, 	_host, 	_updated"

var sqlFundsCheckId, sqlFundsCheckName, sqlFundsCheckStaticin, sqlFundsCheckStaticout, sqlFundsCheckTxnin, sqlFundsCheckTxnout, sqlFundsCheckFundscheckin, sqlFundsCheckFundscheckout, sqlFundsCheckSYSCreated, sqlFundsCheckSYSWho, sqlFundsCheckSYSHost, sqlFundsCheckSYSUpdated sql.NullString

var simFundsCheckNoParams = strings.Count(simFundsCheckSQL, ",") + 1
var simFundsCheckParams = strings.Repeat("'%s',", simFundsCheckNoParams)
var simFundsCheckSQLINSERT = "INSERT INTO %s.fundsCheck(%s) VALUES(" + strings.TrimRight(simFundsCheckParams, ",") + ");"
var simFundsCheckSQLDELETE = "DELETE FROM %s.fundsCheck WHERE id='%s';"
var simFundsCheckSQLSELECT = "SELECT %s FROM %s.fundsCheck;"
var simFundsCheckSQLGET = "SELECT %s FROM %s.fundsCheck WHERE id='%s';"
*/
//simFundsCheckPage is cheese
type simFundsCheckListPage struct {
	UserMenu        []AppMenuItem
	UserRole        string
	UserNavi        string
	Title           string
	PageTitle       string
	FundsCheckCount int
	FundsCheckList  []FundsCheckItem
}

//simFundsCheckPage is cheese
type simFundsCheckPage struct {
	UserMenu  []AppMenuItem
	UserRole  string
	UserNavi  string
	Title     string
	PageTitle string
	ID        string
	Action    string
	// Variable Bits Below
	Id      string
	Name    string
	Source  string
	Content string
	Request FundsCheckRequest
}

//FundsCheckItem is cheese
type FundsCheckItem struct {
	Id      string
	Name    string
	Source  string
	Content string
	Request FundsCheckRequest
}

type FundsCheckRequest struct {
	XMLName        xml.Name `xml:"REQ"`
	Text           string   `xml:",chardata"`
	NS1            string   `xml:"NS1,attr"`
	Xs             string   `xml:"xs,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	HEADER         struct {
		Text        string `xml:",chardata"`
		XREF        string `xml:"XREF"`
		SOURCE      string `xml:"SOURCE"`
		OPERATION   string `xml:"OPERATION"`
		DESTINATION string `xml:"DESTINATION"`
	} `xml:"HEADER"`
	BODY struct {
		Text       string `xml:",chardata"`
		FundsCheck struct {
			Text    string `xml:",chardata"`
			BIC     string `xml:"BIC"`
			ACC     string `xml:"ACC"`
			CCY     string `xml:"CCY"`
			AMOUNT  string `xml:"AMOUNT"`
			VALDATE string `xml:"VALDATE"`
		} `xml:"FundsCheck"`
	} `xml:"BODY"`
}

type FundsCheckReponse struct {
	XMLName   xml.Name `xml:"NS1:MBT_RES"`
	Text      string   `xml:",chardata"`
	NS1       string   `xml:"xmlns:NS1,attr"`
	MBTHEADER struct {
		Text        string `xml:",chardata"`
		XREF        string `xml:"NS1:XREF"`
		SOURCE      string `xml:"NS1:SOURCE"`
		OPERATION   string `xml:"NS1:OPERATION"`
		DESTINATION string `xml:"NS1:DESTINATION"`
		MSGSTAT     string `xml:"NS1:MSGSTAT"`
	} `xml:"NS1:MBT_HEADER"`
	MBTBODY struct {
		Text               string `xml:",chardata"`
		FundsCheckResponse struct {
			Text           string `xml:",chardata"`
			BIC            string `xml:"NS1:BIC"`
			ACC            string `xml:"NS1:ACC"`
			CCY            string `xml:"NS1:CCY"`
			AMOUNT         string `xml:"NS1:AMOUNT"`
			QUERYTIMESTAMP string `xml:"NS1:QUERYTIMESTAMP"`
			RESULTCODE     string `xml:"NS1:RESULTCODE"`
		} `xml:"NS1:FundsCheckResponse"`
	} `xml:"NS1:MBT_BODY"`
}

func ListFundsCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "FundsCheckList"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)
	var returnList []FundsCheckItem
	noItems, returnList, _ := GetFundsCheckList()

	fundsCheckPage := simFundsCheckListPage{
		UserMenu:        GetUserMenu(r),
		UserRole:        GetUserRole(r),
		UserNavi:        "NOT USED",
		Title:           "Outstanding Request",
		PageTitle:       core.ApplicationProperties["appname"] + " - " + "Funds Check Approvals",
		FundsCheckCount: noItems,
		FundsCheckList:  returnList,
	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, fundsCheckPage)

}

func ViewFundsCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "FundsCheckView"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	fundsCheckPage := editViewFundsCheck(w, r)
	fundsCheckPage.PageTitle = core.ApplicationProperties["appname"] + " - Funds Check - View"
	fundsCheckPage.Title = "Request Message Detail"
	//log.Println(fundsCheckPage)
	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, fundsCheckPage)

}

func ActionFundsCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below
	tmpl := "FundsCheckAction"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	fundsCheckPage := editViewFundsCheck(w, r)
	fundsCheckPage.PageTitle = core.ApplicationProperties["appname"] + " - " + "Funds Check Approval - Process"
	fundsCheckPage.Title = "Response Message Detail"

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, fundsCheckPage)

}

func editViewFundsCheck(w http.ResponseWriter, r *http.Request) simFundsCheckPage {

	searchID := GetURLparam(r, "FundsCheck")
	_, returnRecord, _ := GetFundsCheckByID(searchID)

	fundsCheckPage := simFundsCheckPage{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: core.ApplicationProperties["appname"] + " - " + "Funds Check Approval - Request",
		Action:    "",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		// Above are mandatory
		// Below are variable
		Id:      returnRecord.Id,
		Name:    returnRecord.Name,
		Source:  returnRecord.Source,
		Content: returnRecord.Content,
		Request: returnRecord.Request,
	}

	return fundsCheckPage //fmt.Println(fundsCheckPage)
}

func SubmitFundsCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	serviceMessageAction(inUTL, "Save", r.FormValue("Id"))

	thisID := r.FormValue("ID")
	balance := r.FormValue("Balance")
	resultCode := r.FormValue("ResultCode")

	searchID := thisID
	_, returnRecord, _ := GetFundsCheckByID(searchID)

	simFundsCheckItem := returnRecord.Request

	var simFundsCheckResponse FundsCheckReponse
	simFundsCheckResponse.NS1 = "http://dnb.lt/dnb-xst/MBT"
	simFundsCheckResponse.MBTHEADER.XREF = simFundsCheckItem.HEADER.XREF
	simFundsCheckResponse.MBTHEADER.SOURCE = "MB_TODO"
	simFundsCheckResponse.MBTHEADER.OPERATION = simFundsCheckItem.HEADER.OPERATION
	simFundsCheckResponse.MBTHEADER.DESTINATION = simFundsCheckItem.HEADER.SOURCE
	simFundsCheckResponse.MBTHEADER.MSGSTAT = "SUCCESS"
	simFundsCheckResponse.MBTBODY.FundsCheckResponse.BIC = simFundsCheckItem.BODY.FundsCheck.BIC
	simFundsCheckResponse.MBTBODY.FundsCheckResponse.ACC = simFundsCheckItem.BODY.FundsCheck.ACC
	simFundsCheckResponse.MBTBODY.FundsCheckResponse.CCY = simFundsCheckItem.BODY.FundsCheck.CCY
	simFundsCheckResponse.MBTBODY.FundsCheckResponse.AMOUNT = balance
	createDate := time.Now().Format(core.DFNANO)
	simFundsCheckResponse.MBTBODY.FundsCheckResponse.QUERYTIMESTAMP = createDate
	simFundsCheckResponse.MBTBODY.FundsCheckResponse.RESULTCODE = resultCode

	//log.Printf("%# v", simFundsCheckResponse)

	newMsg, err := xml.Marshal(simFundsCheckResponse)
	if err != nil {
		log.Println(string(newMsg), err)
	}

	newID := uuid.New().String()
	fileName := newID + ".xml"

	delivertopath := core.SienaProperties["funds_out"]
	deletefrompath := core.SienaProperties["funds_in"]

	log.Printf("Delivery      : %s -> %s", deletefrompath, delivertopath)
	WriteDataFile(fileName, delivertopath, string(newMsg))

	resp := DeleteDataFile(searchID, deletefrompath)
	if resp != 1 {
		//do nothing
	}
	ListFundsCheckHandler(w, r)

}

func DeleteFundsCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	inUTL := r.URL.Path
	searchID := GetURLparam(r, "FundsCheck")
	serviceMessageAction(inUTL, "Delete", searchID)
	DeleteFundsCheck(searchID)
	ListFundsCheckHandler(w, r)

}

func NewFundsCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(SessionValidate(w, r)) {
		LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "FundsCheckNew"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	serviceMessage(inUTL)

	fundsCheckPage := simFundsCheckPage{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: core.ApplicationProperties["appname"] + " - " + "Funds Check Approval - New",
		UserMenu:  GetUserMenu(r),
		UserRole:  GetUserRole(r),
		UserNavi:  "NOT USED",
		Action:    "",
		// Above are mandatory
		// Below are variable

	}

	t, _ := template.ParseFiles(GetTemplateID(tmpl, GetUserRole(r)))
	t.Execute(w, fundsCheckPage)

}

// getFundsCheckList read all employees
func GetFundsCheckList() (int, []FundsCheckItem, error) {
	//tsql := fmt.Sprintf(simFundsCheckSQLSELECT, simFundsCheckSQL, core.ApplicationPropertiesDB["schema"])
	var simFundsCheckList []FundsCheckItem
	requestPath := core.SienaProperties["funds_in"]
	pwd, _ := os.Getwd()
	files, err := ioutil.ReadDir(pwd + requestPath + "/")
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(files)

	for _, k := range files {
		//fmt.Println("key:", k)
		var fci FundsCheckItem
		fci.Id = k.Name()
		fci.Name = strings.ReplaceAll(k.Name(), ".xml", "")
		fci.Source = requestPath
		simFundsCheckList = append(simFundsCheckList, fci)
	}

	//count, simFundsCheckList, _, _ := fetchFundsCheckData("")
	return len(files), simFundsCheckList, nil
}

// getFundsCheckList read all employees
func GetFundsCheckByID(id string) (int, FundsCheckItem, error) {
	var simFundsCheckItem FundsCheckItem
	requestPath := core.SienaProperties["funds_in"]
	pwd, _ := os.Getwd()
	dat, err := ioutil.ReadFile(pwd + requestPath + "/" + id)
	if err != nil {
		log.Fatal(err)
	}
	simFundsCheckItem.Id = id
	simFundsCheckItem.Name = strings.ReplaceAll(id, ".xml", "")
	simFundsCheckItem.Source = requestPath
	simFundsCheckItem.Content = string(dat)

	var test FundsCheckRequest
	xml.Unmarshal(dat, &test)

	//log.Println(test)
	//log.Println(test.HEADER.XREF)
	//	log.Printf("%# v", test)

	simFundsCheckItem.Request = test

	return 1, simFundsCheckItem, nil
}

/*
func NewFundsCheck(r FundsCheckItem) {
	if len(r.Id) == 0 {
		newID := uuid.New().String()
		r.Id = newID
	}
	putFundsCheck(r)
}
*/
func PutFundsCheck(r FundsCheckItem) {
	putFundsCheck(r)
}

func putFundsCheck(r FundsCheckItem) {
	//fmt.Println(credentialStore)

	//	fmt.Println("RECORD", r)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	//	deletesql := fmt.Sprintf(simFundsCheckSQLDELETE, core.ApplicationPropertiesDB["schema"], r.Id)
	//	inserttsql := fmt.Sprintf(simFundsCheckSQLINSERT, core.ApplicationPropertiesDB["schema"], simFundsCheckSQL, r.Id, r.Name, r.Staticin, r.Staticout, r.Txnin, r.Txnout, r.Fundscheckin, r.Fundscheckout, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated)

	//	log.Println("DELETE:", deletesql, core.ApplicationDB)
	//	log.Println("INSERT:", inserttsql, core.ApplicationDB)
	/*
		_, err2 := core.ApplicationDB.Exec(deletesql)
		if err2 != nil {
			log.Println(err2.Error())
		}
		//	log.Println(fred2, err2)
		_, err := core.ApplicationDB.Exec(inserttsql)
		//	log.Println(fred, err)
		if err != nil {
			log.Println(err.Error())
		}
	*/
}

func DeleteFundsCheck(id string) {
	//fmt.Println(credentialStore)
	//	deletesql := fmt.Sprintf(simFundsCheckSQLDELETE, core.ApplicationPropertiesDB["schema"], id)
	//log.Println("DELETE:", deletesql, core.ApplicationDB)
	//_, err := core.ApplicationDB.Exec(deletesql)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//log.Println(fred2, err2)
}

// fetchFundsCheckData read all employees
func fetchFundsCheckData(tsql string) (int, []FundsCheckItem, FundsCheckItem, error) {
	//log.Println(tsql)
	var simFundsCheck FundsCheckItem
	var simFundsCheckList []FundsCheckItem

	count := 0

	// Populate Arrays etc.
	simFundsCheck.Id = "id"
	simFundsCheck.Name = "name"
	simFundsCheck.Source = "source"
	// no change below
	simFundsCheckList = append(simFundsCheckList, simFundsCheck)
	//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
	count++
	//log.Println(count, simFundsCheckList, simFundsCheck)
	return count, simFundsCheckList, simFundsCheck, nil
}
