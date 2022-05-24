package jobs

import (
	"fmt"
	"log"

	application "github.com/mt1976/mwt-go-dev/application"
	core "github.com/mt1976/mwt-go-dev/core"
	dao "github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
	cron "github.com/robfig/cron/v3"
)

func RefreshCache_Job() dm.JobDefinition {
	var j dm.JobDefinition
	j.ID = "CACHE_REFRESH"
	j.Name = "CACHE_REFRESH"
	j.Period = "*/10 7-19 * * 1-5"
	j.Description = "Update Cache rate from Source System"
	j.Type = core.Aquirer
	return j
}

func RefreshCache_Register(c *cron.Cron) {
	application.Schedule_Register(RefreshCache_Job())
	c.AddFunc(RefreshCache_Job().Period, func() { RefreshCache_Run() })
}

// RunJobRollover is a Rollover function
func RefreshCache_Run() {
	//logs.StartJob(RefreshCache_Job().Name)

	logs.StartJob(RefreshCache_Job().Name)
	var message string
	/// CONTENT STARTS
	initialiseCache()
	/// CONTENT ENDS
	application.Schedule_Update(RefreshCache_Job(), message)
	logs.EndJob(RefreshCache_Job().Name)
}

//initaliseCache rebuilds each cashed set of date table by table
func initialiseCache() {
	//	initialiseCacheData("CounterpartyImportID", "KeyImportID", "Counterparty", "ID", core.SienaPropertiesDB)
	//	initialiseCacheData("Book", "BookName", "", "ID", core.SienaPropertiesDB)
	initialiseCacheData("Currency", "Code", "", "ID", core.SienaPropertiesDB)
	initialiseCacheData("CurrencyPair", "Code", "", "ID", core.SienaPropertiesDB)
	initialiseCacheData("CryptoCurrency", "CODE2", "", "ID", core.SienaPropertiesDB)

	//	InitialiseCacheData("Reason", "Reason")
	//	initialiseCacheData("Portfolio", "Code", "", "ID", core.SienaPropertiesDB)
	//	initialiseCacheData("User", "UserName", "", "Name", core.SienaPropertiesDB)
	//	initialiseCacheData("MandatedUser", "MandatedUserKeyUserName", "", "Name", core.SienaPropertiesDB)
	//	initialiseCacheData("BusinessDate", "Today", "", "", core.SienaPropertiesDB)
	//	initialiseCacheData("Broker", "Code", "", "ID", core.SienaPropertiesDB)
}

//initialiseCacheData rebuilds a set of data from a table
func initialiseCacheData(table string, field string, objectName string, fieldName string, sourceProperties map[string]string) {

	basesql := fmt.Sprintf("SELECT %s FROM %s.siena%s;", field, core.SienaPropertiesDB["schema"], table)

	storeObjectName := objectName
	if len(objectName) == 0 {
		storeObjectName = table
	}

	log.Printf("Caching       : %-20s data -> %-20s from %-15q -> %s on %-15q", table, storeObjectName, sourceProperties["database"], sourceProperties["schema"], sourceProperties["server"])

	//log.Println(basesql)
	dao.DataCache_Build(table, field, fieldName, objectName, basesql, sourceProperties)
}
