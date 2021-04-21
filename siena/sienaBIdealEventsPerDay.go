package siena

import (
	"database/sql"
	"fmt"
	"log"

	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

var sienaBIdealEventsPerDaySQL = "StartInterestDate, 	Count"

var sqlBIDEPDStartInterestDate, sqlBIDEPDCount sql.NullString

//sienaBIdealEventsPerDayItem is cheese
type sienaBIdealEventsPerDayItem struct {
	Action            string
	StartInterestDate string
	Count             string
}

// getSienaBIdealEventsPerDayList read all employees
func getSienaBIdealEventsPerDayList(db *sql.DB) (int, []sienaBIdealEventsPerDayItem, error) {
	mssqlConfig := application.GetProperties(globals.SQLCONFIG)
	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaBIdealEventsPerDay;", sienaBIdealEventsPerDaySQL, mssqlConfig["schema"])
	count, sienaBIdealEventsPerDayList, _, _ := fetchSienaBIdealEventsPerDayData(db, tsql)
	return count, sienaBIdealEventsPerDayList, nil
}

// fetchSienaBIdealEventsPerDayData read all employees
func fetchSienaBIdealEventsPerDayData(db *sql.DB, tsql string) (int, []sienaBIdealEventsPerDayItem, sienaBIdealEventsPerDayItem, error) {

	var sienaBIdealEventsPerDay sienaBIdealEventsPerDayItem
	var sienaBIdealEventsPerDayList []sienaBIdealEventsPerDayItem

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaBIdealEventsPerDay, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlBIDEPDStartInterestDate, &sqlBIDEPDCount)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaBIdealEventsPerDay, err
		}

		sienaBIdealEventsPerDay.StartInterestDate = application.SqlDateToHTMLDate(sqlBIDEPDStartInterestDate.String)
		sienaBIdealEventsPerDay.Count = sqlBIDEPDCount.String

		sienaBIdealEventsPerDayList = append(sienaBIdealEventsPerDayList, sienaBIdealEventsPerDay)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaBIdealEventsPerDayList, sienaBIdealEventsPerDay, nil
}
