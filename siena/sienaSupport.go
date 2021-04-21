package siena

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	support "github.com/mt1976/mwt-go-dev/appsupport"
)

var sienaBusinessDateSQL = "Today"
var sqlSYSDToday sql.NullString

const SIENAXMLIMPORT = "IMPORT"
const SIENAXMLDELETE = "DELETE"
const SIENAXMLUPDATE = "UPDATE"
const SIENAXMLNEW = "INSERT"

//sienaBusinessDateItem is cheese
type sienaBusinessDateItem struct {
	Today    string
	Internal time.Time
	Siena    string
	YYYYMMDD string
	qmEpoch  string
}

type sienaYNItem struct {
	Code string
	Name string
}

func getSienaYNList() (int, []sienaYNItem, error) {
	var YNList []sienaYNItem
	var YNItem sienaYNItem
	YNItem.Code = "true"
	YNItem.Name = "Yes"
	YNList = append(YNList, YNItem)
	YNItem.Code = "false"
	YNItem.Name = "No"
	YNList = append(YNList, YNItem)
	return 2, YNList, nil
}

func getSienaYNTickList() (int, []sienaYNItem, error) {
	var YNList []sienaYNItem
	var YNItem sienaYNItem
	YNItem.Code = "true"
	YNItem.Name = "checked"
	YNList = append(YNList, YNItem)
	YNItem.Code = "false"
	YNItem.Name = ""
	YNList = append(YNList, YNItem)
	return 2, YNList, nil
}

func sienaYN(inValue string) string {
	var outValue string
	outValue = ""
	if inValue == "true" {
		outValue = "Yes"
	}
	if inValue == "false" {
		outValue = "No"
	}
	return outValue
}

func setChecked(inValue string) string {
	var outValue string
	outValue = ""
	if inValue == "true" {
		outValue = "checked"
	}
	if inValue == "false" {
		outValue = ""
	}
	return outValue
}

func isChecked(inValue string) string {
	var outValue string
	fmt.Println(inValue)
	outValue = ""
	if inValue == "true" {
		outValue = "checked"
	}
	if inValue == "false" {
		outValue = ""
	}
	return outValue
}

func sienaConnect() (*sql.DB, error) {
	// Connect to SQL Server DB
	mssqlConfig := support.GetProperties(SQLCONFIG)

	server := mssqlConfig["server"]
	user := mssqlConfig["user"]
	password := mssqlConfig["password"]
	port := mssqlConfig["port"]
	database := mssqlConfig["database"]
	/*
		fmt.Println("SQL SERVER - CONNECTING...")
		fmt.Println("Server     :", server)
		fmt.Println("User       :", user)
		fmt.Println("Password   :", strings.Repeat("*", len(password)))
		fmt.Println("Port       :", port)
		fmt.Println("Database   :", database)
		fmt.Println("")
	*/
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		server, user, password, port, database)
	dbInstance, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	//fmt.Printf("Connected!\n")
	//fmt.Println("")

	//defer dbInstance.Close()

	stmt, err := dbInstance.Prepare("select @@version")
	row := stmt.QueryRow()
	var result string

	err = row.Scan(&result)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	//fmt.Printf("%s\n", result)
	return dbInstance, err
}

// getSienaBusinessDate read all employees
func getSienaBusinessDate(db *sql.DB) (int, sienaBusinessDateItem, error) {
	mssqlConfig := support.GetProperties(SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.SienaBusinessDate;", sienaBusinessDateSQL, mssqlConfig["schema"])
	_, _, SienaBusinessDate, _ := fetchSienaBusinessDateData(db, tsql)
	return 1, SienaBusinessDate, nil
}

// fetchSienaBusinessDateData read all employees
func fetchSienaBusinessDateData(db *sql.DB, tsql string) (int, []sienaBusinessDateItem, sienaBusinessDateItem, error) {

	var sienaBusinessDate sienaBusinessDateItem
	var sienaBusinessDateList []sienaBusinessDateItem

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaBusinessDate, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlSYSDToday)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaBusinessDate, err
		}

		sienaBusinessDate.Today = sqlSYSDToday.String

		sienaBusinessDate.Internal, _ = time.Parse(time.RFC3339, sqlSYSDToday.String)
		sienaBusinessDate.Siena = sienaBusinessDate.Internal.Format(DATEFORMATSIENA)
		sienaBusinessDate.YYYYMMDD = sienaBusinessDate.Internal.Format(DATEFORMATYMD)
		sienaBusinessDate.qmEpoch = sienaBusinessDate.Internal.Format(DATEFORMATPICK)
		sienaBusinessDateList = append(sienaBusinessDateList, sienaBusinessDate)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	//log.Println("Dates", sienaBusinessDate)
	return count, sienaBusinessDateList, sienaBusinessDate, nil
}

func sienaDispatchStaticDataXML(sienaXMLContent sienaXML) error {

	sienaProperties := support.GetProperties(SIENACONFIG)

	preparedXML, _ := xml.Marshal(sienaXMLContent)
	fmt.Println("PreparedXML", string(preparedXML))

	staticImporterPath := sienaProperties["static_in"]
	fileID := uuid.New()
	fileName := staticImporterPath + "/" + fileID.String() + ".xml"
	fmt.Println(fileName)

	xmlFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating XML file: ", err)
		return err
	}
	xmlFile.WriteString(sienaProperties["static_xml_encoding"] + "\n")
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(sienaXMLContent)
	if err != nil {
		fmt.Println("Error encoding XML to file: ", err)
		return err
	}
	return err
}
