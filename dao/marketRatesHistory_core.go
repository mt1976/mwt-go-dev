package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/marketrateshistory.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : MarketRatesHistory (marketrateshistory)
// Endpoint 	        : MarketRatesHistory (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 31/08/2022 at 14:19:48
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (

	"fmt"
	"net/http"
core "github.com/mt1976/mwt-go-dev/core"
"github.com/google/uuid"
das  "github.com/mt1976/mwt-go-dev/das"
	
	
	
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
)

// MarketRatesHistory_GetList() returns a list of all MarketRatesHistory records
func MarketRatesHistory_GetList() (int, []dm.MarketRatesHistory, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.MarketRatesHistory_SQLTable)
	count, marketrateshistoryList, _, _ := marketrateshistory_Fetch(tsql)
	
	return count, marketrateshistoryList, nil
}



// MarketRatesHistory_GetByID() returns a single MarketRatesHistory record
func MarketRatesHistory_GetByID(id string) (int, dm.MarketRatesHistory, error) {


	tsql := "SELECT * FROM " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.MarketRatesHistory_SQLTable)
	tsql = tsql + " WHERE " + dm.MarketRatesHistory_SQLSearchID + "='" + id + "'"
	_, _, marketrateshistoryItem, _ := marketrateshistory_Fetch(tsql)

	// START
	// Dynamically generated 31/08/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 31/08/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, marketrateshistoryItem, nil
}



// MarketRatesHistory_DeleteByID() deletes a single MarketRatesHistory record
func MarketRatesHistory_Delete(id string) {


	object_Table := core.GetSQLSchema(core.ApplicationPropertiesDB) + "." + dm.MarketRatesHistory_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.MarketRatesHistory_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// MarketRatesHistory_Store() saves/stores a MarketRatesHistory record to the database
func MarketRatesHistory_Store(r dm.MarketRatesHistory,req *http.Request) error {

	err, r := MarketRatesHistory_Validate(r)
	if err == nil {
		err = marketrateshistory_Save(r, Audit_User(req))
	} else {
		logs.Information("MarketRatesHistory_Store()", err.Error())
	}

	return err
}

// MarketRatesHistory_StoreSystem() saves/stores a MarketRatesHistory record to the database
func MarketRatesHistory_StoreSystem(r dm.MarketRatesHistory) error {
	
	err, r := MarketRatesHistory_Validate(r)
	if err == nil {
		err = marketrateshistory_Save(r, Audit_Host())
	} else {
		logs.Information("MarketRatesHistory_Store()", err.Error())
	}

	return err
}

// MarketRatesHistory_Validate() validates for saves/stores a MarketRatesHistory record to the database
func MarketRatesHistory_Validate(r dm.MarketRatesHistory) (error,dm.MarketRatesHistory) {
	var err error
	// START
	// Dynamically generated 31/08/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 31/08/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	

	return err,r
}
//

// marketrateshistory_Save() saves/stores a MarketRatesHistory record to the database
func marketrateshistory_Save(r dm.MarketRatesHistory,usr string) error {

    var err error



	

	if len(r.Id) == 0 {
		r.Id = MarketRatesHistory_NewID(r)
	}

// If there are fields below, create the methods in dao\marketrateshistory_impl.go

























	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	
logs.Storing("MarketRatesHistory",fmt.Sprintf("%s", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 31/08/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.MarketRatesHistory_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.MarketRatesHistory_Id_sql, r.Id)
	ts = addData(ts, dm.MarketRatesHistory_Bid_sql, r.Bid)
	ts = addData(ts, dm.MarketRatesHistory_Mid_sql, r.Mid)
	ts = addData(ts, dm.MarketRatesHistory_Offer_sql, r.Offer)
	ts = addData(ts, dm.MarketRatesHistory_Market_sql, r.Market)
	ts = addData(ts, dm.MarketRatesHistory_Tenor_sql, r.Tenor)
	ts = addData(ts, dm.MarketRatesHistory_Series_sql, r.Series)
	ts = addData(ts, dm.MarketRatesHistory_Name_sql, r.Name)
	ts = addData(ts, dm.MarketRatesHistory_Source_sql, r.Source)
	ts = addData(ts, dm.MarketRatesHistory_Destination_sql, r.Destination)
	ts = addData(ts, dm.MarketRatesHistory_Class_sql, r.Class)
	ts = addData(ts, dm.MarketRatesHistory_Date_sql, r.Date)
	ts = addData(ts, dm.MarketRatesHistory_Time_sql, r.Time)
	ts = addData(ts, dm.MarketRatesHistory_Datetime_sql, r.Datetime)
	ts = addData(ts, dm.MarketRatesHistory_SYSWho_sql, r.SYSWho)
	ts = addData(ts, dm.MarketRatesHistory_SYSHost_sql, r.SYSHost)
	ts = addData(ts, dm.MarketRatesHistory_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.MarketRatesHistory_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.MarketRatesHistory_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.MarketRatesHistory_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.MarketRatesHistory_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.MarketRatesHistory_SYSUpdatedHost_sql, r.SYSUpdatedHost)
		
	// 
	// Dynamically generated 31/08/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := "INSERT INTO " + get_TableName(core.GetSQLSchema(core.ApplicationPropertiesDB), dm.MarketRatesHistory_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	MarketRatesHistory_Delete(r.Id)
	das.Execute(tsql)



	return err

}



// marketrateshistory_Fetch read all MarketRatesHistory's
func marketrateshistory_Fetch(tsql string) (int, []dm.MarketRatesHistory, dm.MarketRatesHistory, error) {

	var recItem dm.MarketRatesHistory
	var recList []dm.MarketRatesHistory

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 31/08/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.MarketRatesHistory_SYSId_sql, "0")
	   recItem.Id  = get_String(rec, dm.MarketRatesHistory_Id_sql, "")
	   recItem.Bid  = get_String(rec, dm.MarketRatesHistory_Bid_sql, "")
	   recItem.Mid  = get_String(rec, dm.MarketRatesHistory_Mid_sql, "")
	   recItem.Offer  = get_String(rec, dm.MarketRatesHistory_Offer_sql, "")
	   recItem.Market  = get_String(rec, dm.MarketRatesHistory_Market_sql, "")
	   recItem.Tenor  = get_String(rec, dm.MarketRatesHistory_Tenor_sql, "")
	   recItem.Series  = get_String(rec, dm.MarketRatesHistory_Series_sql, "")
	   recItem.Name  = get_String(rec, dm.MarketRatesHistory_Name_sql, "")
	   recItem.Source  = get_String(rec, dm.MarketRatesHistory_Source_sql, "")
	   recItem.Destination  = get_String(rec, dm.MarketRatesHistory_Destination_sql, "")
	   recItem.Class  = get_String(rec, dm.MarketRatesHistory_Class_sql, "")
	   recItem.Date  = get_String(rec, dm.MarketRatesHistory_Date_sql, "")
	   recItem.Time  = get_String(rec, dm.MarketRatesHistory_Time_sql, "")
	   recItem.Datetime  = get_String(rec, dm.MarketRatesHistory_Datetime_sql, "")
	   recItem.SYSWho  = get_String(rec, dm.MarketRatesHistory_SYSWho_sql, "")
	   recItem.SYSHost  = get_String(rec, dm.MarketRatesHistory_SYSHost_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.MarketRatesHistory_SYSCreated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.MarketRatesHistory_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.MarketRatesHistory_SYSCreatedHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.MarketRatesHistory_SYSUpdated_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.MarketRatesHistory_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.MarketRatesHistory_SYSUpdatedHost_sql, "")
	
	// If there are fields below, create the methods in adaptor\MarketRatesHistory_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 31/08/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func MarketRatesHistory_NewID(r dm.MarketRatesHistory) string {
	
			id := uuid.New().String()
	
	return id
}



// marketrateshistory_Fetch read all MarketRatesHistory's
func MarketRatesHistory_New() (int, []dm.MarketRatesHistory, dm.MarketRatesHistory, error) {

	var r = dm.MarketRatesHistory{}
	var rList []dm.MarketRatesHistory
	

	// START
	// Dynamically generated 31/08/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 31/08/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}