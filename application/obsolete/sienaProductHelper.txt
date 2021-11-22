package application

import (
	"database/sql"
	"fmt"
	"log"

	core "github.com/mt1976/mwt-go-dev/core"
)

var sienaProductHelperSQL = "Code, 	Name"

var sqlProductHelperCode, sqlProductHelperName sql.NullString

//sienaProductHelperItem is cheese
type sienaProductHelperItem struct {
	Action string
	Code   string
	Name   string
}

// getSienaProductHelperList read all employees
func getSienaProductHelperList(db *sql.DB) (int, []sienaProductHelperItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaProductHelper;", sienaProductHelperSQL, core.SienaPropertiesDB["schema"])
	count, sienaProductHelperList, _, _ := fetchSienaProductHelperData(db, tsql)
	return count, sienaProductHelperList, nil
}

// fetchSienaProductHelperData read all employees
func fetchSienaProductHelperData(db *sql.DB, tsql string) (int, []sienaProductHelperItem, sienaProductHelperItem, error) {

	var sienaProductHelper sienaProductHelperItem
	var sienaProductHelperList []sienaProductHelperItem

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaProductHelper, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlProductHelperCode, &sqlProductHelperName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaProductHelper, err
		}

		sienaProductHelper.Code = sqlProductHelperCode.String
		sienaProductHelper.Name = sqlProductHelperName.String

		sienaProductHelperList = append(sienaProductHelperList, sienaProductHelper)
		count++
	}
	return count, sienaProductHelperList, sienaProductHelper, nil
}
