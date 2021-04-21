package siena

import (
	"database/sql"
	"fmt"
	"log"

	support "github.com/mt1976/mwt-go-dev/appsupport"
)

var sienaBIcounterpartyPerSectorSQL = "SectorCode, 	Count"

var sqlBISECTSectorCode, sqlBISECTCount sql.NullString

//sienaBIcounterpartyPerSectorItem is cheese
type sienaBIcounterpartyPerSectorItem struct {
	Action     string
	SectorCode string
	Count      string
}

// getSienaBIcounterpartyPerSectorList read all employees
func getSienaBIcounterpartyPerSectorList(db *sql.DB) (int, []sienaBIcounterpartyPerSectorItem, error) {
	mssqlConfig := support.GetProperties(SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaBIcounterpartyPerSector;", sienaBIcounterpartyPerSectorSQL, mssqlConfig["schema"])
	count, sienaBIcounterpartyPerSectorList, _, _ := fetchSienaBIcounterpartyPerSectorData(db, tsql)
	return count, sienaBIcounterpartyPerSectorList, nil
}

// fetchSienaBIcounterpartyPerSectorData read all employees
func fetchSienaBIcounterpartyPerSectorData(db *sql.DB, tsql string) (int, []sienaBIcounterpartyPerSectorItem, sienaBIcounterpartyPerSectorItem, error) {

	var sienaBIcounterpartyPerSector sienaBIcounterpartyPerSectorItem
	var sienaBIcounterpartyPerSectorList []sienaBIcounterpartyPerSectorItem

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaBIcounterpartyPerSector, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlBISECTSectorCode, &sqlBISECTCount)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaBIcounterpartyPerSector, err
		}

		sienaBIcounterpartyPerSector.SectorCode = sqlBISECTSectorCode.String
		sienaBIcounterpartyPerSector.Count = sqlBISECTCount.String

		sienaBIcounterpartyPerSectorList = append(sienaBIcounterpartyPerSectorList, sienaBIcounterpartyPerSector)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaBIcounterpartyPerSectorList, sienaBIcounterpartyPerSector, nil
}
