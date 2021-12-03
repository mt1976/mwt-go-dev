package application

import (
	"database/sql"
	"fmt"
	"log"

	core "github.com/mt1976/mwt-go-dev/core"
)

var sienaDealTypeHelperSQL = "DealTypeKey, 	DealTypeShortName, 	NegotiableInstrumentType, 	ProductCode, 	ProductCodeName"

var sqlDealTypeHelperDealTypeKey, sqlDealTypeHelperDealTypeShortName, sqlDealTypeHelperNegotiableInstrumentType, sqlDealTypeHelperProductCode, sqlDealTypeHelperProductCodeName sql.NullString

//sienaDealTypeHelperItem is cheese
type sienaDealTypeHelperItem struct {
	Action                   string
	DealTypeKey              string
	DealTypeShortName        string
	NegotiableInstrumentType string
	ProductCode              string
	ProductCodeName          string
}

// getSienaDealTypeHelperList read all employees
func getSienaDealTypeHelperList(db *sql.DB) (int, []sienaDealTypeHelperItem, error) {

	tsql := fmt.Sprintf("SELECT %s FROM %s.sienaDealTypeHelper;", sienaDealTypeHelperSQL, core.SienaPropertiesDB["schema"])
	count, sienaDealTypeHelperList, _, _ := fetchSienaDealTypeHelperData(db, tsql)
	return count, sienaDealTypeHelperList, nil
}

// fetchSienaDealTypeHelperData read all employees
func fetchSienaDealTypeHelperData(db *sql.DB, tsql string) (int, []sienaDealTypeHelperItem, sienaDealTypeHelperItem, error) {

	var sienaDealTypeHelper sienaDealTypeHelperItem
	var sienaDealTypeHelperList []sienaDealTypeHelperItem

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, sienaDealTypeHelper, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&sqlDealTypeHelperDealTypeKey, &sqlDealTypeHelperDealTypeShortName, &sqlDealTypeHelperNegotiableInstrumentType, &sqlDealTypeHelperProductCode, &sqlDealTypeHelperProductCodeName)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, sienaDealTypeHelper, err
		}

		sienaDealTypeHelper.DealTypeKey = sqlDealTypeHelperDealTypeKey.String
		sienaDealTypeHelper.DealTypeShortName = sqlDealTypeHelperDealTypeShortName.String
		sienaDealTypeHelper.NegotiableInstrumentType = sqlDealTypeHelperNegotiableInstrumentType.String
		sienaDealTypeHelper.ProductCode = sqlDealTypeHelperProductCode.String
		sienaDealTypeHelper.ProductCodeName = sqlDealTypeHelperProductCodeName.String

		sienaDealTypeHelperList = append(sienaDealTypeHelperList, sienaDealTypeHelper)
		count++
	}
	return count, sienaDealTypeHelperList, sienaDealTypeHelper, nil
}
