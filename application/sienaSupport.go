package application

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
)

var sienaBusinessDateSQL = "Today"
var sqlSYSDToday sql.NullString

const SIENAXMLIMPORT = "IMPORT"
const SIENAXMLDELETE = "DELETE"
const SIENAXMLUPDATE = "UPDATE"
const SIENAXMLNEW = "INSERT"

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

func Connect() (*sql.DB, error) {
	// Connect to SQL Server DB

	server := core.SienaPropertiesDB["server"]
	user := core.SienaPropertiesDB["user"]
	password := core.SienaPropertiesDB["password"]
	port := core.SienaPropertiesDB["port"]
	database := core.SienaPropertiesDB["database"]
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
func GetBusinessDate(db *sql.DB) (int, core.DateItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaBusinessDate;", sienaBusinessDateSQL, core.SienaPropertiesDB["schema"])
	_, _, SienaBusinessDate, _ := fetchSienaBusinessDateData(db, tsql)
	return 1, SienaBusinessDate, nil
}

// fetchSienaBusinessDateData read all employees
func fetchSienaBusinessDateData(db *sql.DB, tsql string) (int, []core.DateItem, core.DateItem, error) {

	var sienaBusinessDate core.DateItem
	var sienaBusinessDateList []core.DateItem

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
		sienaBusinessDate.Siena = sienaBusinessDate.Internal.Format(core.DATEFORMATSIENA)
		sienaBusinessDate.YYYYMMDD = sienaBusinessDate.Internal.Format(core.DATEFORMATYMD)
		sienaBusinessDate.PICKEpoch = sienaBusinessDate.Internal.Format(core.DATEFORMATPICK)
		sienaBusinessDateList = append(sienaBusinessDateList, sienaBusinessDate)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	//log.Println("Dates", sienaBusinessDate)
	return count, sienaBusinessDateList, sienaBusinessDate, nil
}

func sienaDispatchStaticDataXML(StaticImport_XMLContent StaticImport_XML) error {

	preparedXML, _ := xml.Marshal(StaticImport_XMLContent)
	fmt.Println("PreparedXML", string(preparedXML))

	staticImporterPath := core.SienaProperties["static_in"]
	fileID := uuid.New()
	fileName := staticImporterPath + "/" + fileID.String() + ".xml"
	fmt.Println(fileName)

	xmlFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating XML file: ", err)
		return err
	}
	xmlFile.WriteString(core.SienaProperties["static_xml_encoding"] + "\n")
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(StaticImport_XMLContent)
	if err != nil {
		fmt.Println("Error encoding XML to file: ", err)
		return err
	}
	return err
}
