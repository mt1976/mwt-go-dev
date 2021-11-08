package application

import (
	"database/sql"
	"fmt"
	"log"

	core "github.com/mt1976/mwt-go-dev/core"
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
func getSienaBIcounterpartyPerSectorList() (int, []sienaBIcounterpartyPerSectorItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaBIcounterpartyPerSector;", sienaBIcounterpartyPerSectorSQL, core.SienaPropertiesDB["schema"])
	count, sienaBIcounterpartyPerSectorList, _, _ := fetchSienaBIcounterpartyPerSectorData(tsql)
	return count, sienaBIcounterpartyPerSectorList, nil
}

// fetchSienaBIcounterpartyPerSectorData read all employees
func fetchSienaBIcounterpartyPerSectorData(tsql string) (int, []sienaBIcounterpartyPerSectorItem, sienaBIcounterpartyPerSectorItem, error) {

	var sienaBIcounterpartyPerSector sienaBIcounterpartyPerSectorItem
	var sienaBIcounterpartyPerSectorList []sienaBIcounterpartyPerSectorItem

	rows, err := core.SienaDB.Query(tsql)
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
