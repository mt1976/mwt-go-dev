package application

import (
	"database/sql"
	"fmt"
	"log"

	core "github.com/mt1976/mwt-go-dev/core"
)

var sienaBIdealEventsPerDaySQL = "StartInterestDate, Count"

var sqlBIDEPDStartInterestDate, sqlBIDEPDCount sql.NullString

//sienaBIdealEventsPerDayItem is cheese
type sienaBIdealEventsPerDayItem struct {
	Action            string
	StartInterestDate string
	Count             string
}

// getSienaBIdealEventsPerDayList read all employees
func getSienaBIdealEventsPerDayList() (int, []sienaBIdealEventsPerDayItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaBIdealEventsPerDay", sienaBIdealEventsPerDaySQL, core.SienaPropertiesDB["schema"])
	//tsql = fmt.Sprintf("%s GROUP BY %s", tsql, "StartInterestDate")
	tsql = fmt.Sprintf("%s ORDER BY %s", tsql, "StartInterestDate")
	count, sienaBIdealEventsPerDayList, _, _ := fetchSienaBIdealEventsPerDayData(tsql)
	return count, sienaBIdealEventsPerDayList, nil
}

// fetchSienaBIdealEventsPerDayData read all employees
func fetchSienaBIdealEventsPerDayData(tsql string) (int, []sienaBIdealEventsPerDayItem, sienaBIdealEventsPerDayItem, error) {
	fmt.Printf("tsql: %v\n", tsql)
	var sienaBIdealEventsPerDay sienaBIdealEventsPerDayItem
	var sienaBIdealEventsPerDayList []sienaBIdealEventsPerDayItem

	rows, err := core.SienaDB.Query(tsql)
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

		sienaBIdealEventsPerDay.StartInterestDate = core.SqlDateToHTMLDate(sqlBIDEPDStartInterestDate.String)
		sienaBIdealEventsPerDay.Count = sqlBIDEPDCount.String

		sienaBIdealEventsPerDayList = append(sienaBIdealEventsPerDayList, sienaBIdealEventsPerDay)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaBIdealEventsPerDayList, sienaBIdealEventsPerDay, nil
}
