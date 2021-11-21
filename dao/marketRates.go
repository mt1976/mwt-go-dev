package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/marketrates.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : MarketRates (marketrates)
// Endpoint 	        : MarketRates (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 21/11/2021 at 15:44:04
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"log"
	"fmt"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	das  "github.com/mt1976/mwt-go-dev/das"
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
	
)

// MarketRates_GetList() returns a list of all MarketRates records
func MarketRates_GetList() (int, []dm.MarketRates, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.MarketRates_SQLTable)
	count, marketratesList, _, _ := marketrates_Fetch(tsql)
	return count, marketratesList, nil
}

// MarketRates_GetByID() returns a single MarketRates record
func MarketRates_GetByID(id string) (int, dm.MarketRates, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.MarketRates_SQLTable)
	tsql = tsql + " WHERE " + dm.MarketRates_SQLSearchID + "='" + id + "'"

	_, _, marketratesItem, _ := marketrates_Fetch(tsql)
	return 1, marketratesItem, nil
}



// MarketRates_DeleteByID() deletes a single MarketRates record
func MarketRates_Delete(id string) {

	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.MarketRates_SQLTable

	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.MarketRates_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
}

// MarketRates_Store() saves/stores a MarketRates record to the database
func MarketRates_Store(r dm.MarketRates) error {

	logs.Warning(fmt.Sprintf("%s", r))

	if len(r.Id) == 0 {
		r.Id = MarketRates_NewID(r)
	}




//Deal with the if its Application or null add this bit, otherwise dont.
		//fmt.Println(credentialStore)

	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, Audit_User())
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",Audit_User())
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

	ts = addData(ts, dm.MarketRates_SYSId, r.SYSId)
	ts = addData(ts, dm.MarketRates_Id, r.Id)
	ts = addData(ts, dm.MarketRates_Bid, r.Bid)
	ts = addData(ts, dm.MarketRates_Mid, r.Mid)
	ts = addData(ts, dm.MarketRates_Offer, r.Offer)
	ts = addData(ts, dm.MarketRates_Market, r.Market)
	ts = addData(ts, dm.MarketRates_Tenor, r.Tenor)
	ts = addData(ts, dm.MarketRates_Series, r.Series)
	ts = addData(ts, dm.MarketRates_Name, r.Name)
	ts = addData(ts, dm.MarketRates_Source, r.Source)
	ts = addData(ts, dm.MarketRates_Destination, r.Destination)
	ts = addData(ts, dm.MarketRates_Class, r.Class)
	ts = addData(ts, dm.MarketRates_SYSCreated, r.SYSCreated)
	ts = addData(ts, dm.MarketRates_SYSWho, r.SYSWho)
	ts = addData(ts, dm.MarketRates_SYSHost, r.SYSHost)
	ts = addData(ts, dm.MarketRates_Date, r.Date)
	ts = addData(ts, dm.MarketRates_SYSUpdated, r.SYSUpdated)
	ts = addData(ts, dm.MarketRates_SYSCreatedBy, r.SYSCreatedBy)
	ts = addData(ts, dm.MarketRates_SYSCreatedHost, r.SYSCreatedHost)
	ts = addData(ts, dm.MarketRates_SYSUpdatedBy, r.SYSUpdatedBy)
	ts = addData(ts, dm.MarketRates_SYSUpdatedHost, r.SYSUpdatedHost)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.MarketRates_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	MarketRates_Delete(r.Id)
	das.Execute(tsql)


	return nil
}

// marketrates_Fetch read all employees
func marketrates_Fetch(tsql string) (int, []dm.MarketRates, dm.MarketRates, error) {

	var recItem dm.MarketRates
	var recList []dm.MarketRates

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 21/11/2021 by matttownsend on silicon.local - START
    recItem.AppInternalID = get_String(rec, dm.MarketRates_Id,"")
   recItem.SYSId  = get_Int(rec, dm.MarketRates_SYSId, "0")
   recItem.Id  = get_String(rec, dm.MarketRates_Id, "")
   recItem.Bid  = get_String(rec, dm.MarketRates_Bid, "")
   recItem.Mid  = get_String(rec, dm.MarketRates_Mid, "")
   recItem.Offer  = get_String(rec, dm.MarketRates_Offer, "")
   recItem.Market  = get_String(rec, dm.MarketRates_Market, "")
   recItem.Tenor  = get_String(rec, dm.MarketRates_Tenor, "")
   recItem.Series  = get_String(rec, dm.MarketRates_Series, "")
   recItem.Name  = get_String(rec, dm.MarketRates_Name, "")
   recItem.Source  = get_String(rec, dm.MarketRates_Source, "")
   recItem.Destination  = get_String(rec, dm.MarketRates_Destination, "")
   recItem.Class  = get_String(rec, dm.MarketRates_Class, "")
   recItem.SYSCreated  = get_String(rec, dm.MarketRates_SYSCreated, "")
   recItem.SYSWho  = get_String(rec, dm.MarketRates_SYSWho, "")
   recItem.SYSHost  = get_String(rec, dm.MarketRates_SYSHost, "")
   recItem.Date  = get_String(rec, dm.MarketRates_Date, "")
   recItem.SYSUpdated  = get_String(rec, dm.MarketRates_SYSUpdated, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.MarketRates_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.MarketRates_SYSCreatedHost, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.MarketRates_SYSUpdatedBy, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.MarketRates_SYSUpdatedHost, "")
// Automatically generated 21/11/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func MarketRates_NewID(r dm.MarketRates) string {
	
	
			id := uuid.New().String()

	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

