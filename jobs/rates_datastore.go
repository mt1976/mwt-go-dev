package jobs

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	core "github.com/mt1976/mwt-go-dev/core"
	"github.com/mt1976/mwt-go-dev/dao"
)

type RatesDataStore struct {
	bid         string
	mid         string
	offer       string
	market      string
	tenor       string
	series      string
	name        string
	source      string
	destination string
	class       string
	date        string
}

var appRatesDataStoreSQL = "id, bid, mid, offer, market, tenor, series, name, [source], destination, class, _created, _who, _host, [date]"

// Commented out sqlRatesDataStore... until we implement a read function
//var sqlRatesDataStoreID, sqlRatesDataStorebid, sqlRatesDataStoremid, sqlRatesDataStoreoffer, sqlRatesDataStoremarket, sqlRatesDataStoretenor, sqlRatesDataStoreseries, sqlRatesDataStorename, sqlRatesDataStoresource, sqlRatesDataStoredestination, sqlRatesDataStoreclass, sqlRatesDataStorecreated, sqlRatesDataStoreby, sqlRatesDataStoreon, sqlRatesDataStoredate sql.NullString
var appRatesDataStoreSQLINSERT = "INSERT INTO %s.rateDataStore(%s) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');"
var appRatesDataStoreSQLDELETE = "DELETE FROM %s.rateDataStore WHERE id='%s';"

type DataStoreSQL struct {
	ID          string `sql:"que"`
	bid         string `sql:"nana"`
	mid         string `sql:"mid"`
	offer       string `sql:"offer"`
	market      string `sql:"market"`
	tenor       string `sql:"tenor"`
	series      string `sql:"series"`
	name        string `sql:"name"`
	source      string `sql:"source"`
	destination string `sql:"destination"`
	class       string `sql:"class"`
	created     string `sql:"created"`
	who         string `sql:"by"`
	host        string `sql:"on"`
	date        string `sql:"date"`
}

func RunJobRatesDataStore(actionType string) {
	var ratesData RatesDataStore
	ratesData.bid = "1.3847"
	ratesData.mid = ""
	ratesData.offer = "1.3849"
	ratesData.market = "FX"
	ratesData.tenor = "SP"
	ratesData.series = "GBPUSD"
	ratesData.class = "Market"
	ratesData.name = "Cable"
	ratesData.source = "BOE"
	ratesData.destination = "RVMARKET"

	RatesDataStorePut(ratesData)

	var mm RatesDataStore
	mm.bid = ""
	mm.mid = "0.01"
	mm.offer = ""
	mm.market = "MM"
	mm.tenor = "ON"
	mm.series = "USD"
	mm.class = "SOFR"
	mm.name = "SOFR"
	mm.source = "BOE"
	mm.destination = "RVMARKET"
	RatesDataStorePut(mm)

	var cpi RatesDataStore
	cpi.bid = ""
	cpi.mid = "0.7"
	cpi.offer = ""
	cpi.market = "CPi"
	cpi.tenor = "ON"
	cpi.series = "GBR"
	cpi.class = "INDEX"
	cpi.name = "CPI"
	cpi.source = "BOE"
	cpi.destination = "RVMARKET"
	RatesDataStorePut(cpi)

	var cpi2 RatesDataStore
	cpi2.bid = ""
	cpi2.mid = "264.79"
	cpi2.offer = ""
	cpi2.market = "CPi"
	cpi2.tenor = "ON"
	cpi2.series = "USA"
	cpi2.class = "INDEX"
	cpi2.name = "CPI"
	cpi2.source = "FRED"
	cpi2.destination = "RVMARKET"
	RatesDataStorePut(cpi2)

	//logit(actionType, "*** DONE ***")
}

func RatesDataStorePut(ratesData RatesDataStore) {
	//fmt.Println(ratesData)
	if len(ratesData.mid) == 0 {
		b, _ := strconv.ParseFloat(ratesData.bid, 64)
		o, _ := strconv.ParseFloat(ratesData.offer, 64)
		m := (b + o) / 2
		ratesData.mid = fmt.Sprintf("%f", m)
	}

	if len(ratesData.date) == 0 {
		ratesData.date = time.Now().Format("2006-01-02")
	}

	ratesData.bid = strings.Replace(ratesData.bid, "\r\n", "", -1)
	ratesData.mid = strings.Replace(ratesData.mid, "\r\n", "", -1)
	ratesData.offer = strings.Replace(ratesData.offer, "\r\n", "", -1)

	ratesData.series = strings.Replace(ratesData.series, "/", "", -1)
	ratesData.series = strings.Replace(ratesData.series, "*", "", -1)

	recordID := getRatesDataStoreKey(ratesData)

	_, cd, _ := dao.MarketRates_GetByID(recordID)

	cd.Id = recordID
	cd.Bid = ratesData.bid
	cd.Offer = ratesData.offer
	cd.Mid = ratesData.mid
	cd.Market = ratesData.market
	cd.Tenor = ratesData.tenor
	cd.Series = ratesData.series
	cd.Name = ratesData.name
	cd.Source = ratesData.source
	cd.Destination = ratesData.destination
	cd.Class = ratesData.class

	dao.MarketRates_StoreSystem(cd)

}

func getRatesDataStoreKey(ratesData RatesDataStore) string {
	key := strings.ToUpper(ratesData.market) + core.IDSep + strings.ToUpper(ratesData.tenor) + core.IDSep + strings.ToUpper(ratesData.class) + core.IDSep + strings.ToUpper(ratesData.series)
	return key
}
