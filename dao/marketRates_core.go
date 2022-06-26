package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/marketrates.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : MarketRates (marketrates)
// Endpoint 	        : MarketRates (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:30
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

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
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
func MarketRates_Store(r dm.MarketRates,req *http.Request) error {

	err, r := MarketRates_Validate(r)
	if err == nil {
		err = marketrates_Save(r, Audit_User(req))
	} else {
		logs.Information("MarketRates_Store()", err.Error())
	}

	return err
}

// MarketRates_StoreSystem() saves/stores a MarketRates record to the database
func MarketRates_StoreSystem(r dm.MarketRates) error {
	
	err, r := MarketRates_Validate(r)
	if err == nil {
		err = marketrates_Save(r, Audit_Host())
	} else {
		logs.Information("MarketRates_Store()", err.Error())
	}

	return err
}

// MarketRates_Validate() validates for saves/stores a MarketRates record to the database
func MarketRates_Validate(r dm.MarketRates) (error,dm.MarketRates) {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return nil,r
}
//

// marketrates_Save() saves/stores a MarketRates record to the database
func marketrates_Save(r dm.MarketRates,usr string) error {

    var err error



	

	if len(r.Id) == 0 {
		r.Id = MarketRates_NewID(r)
	}

// If there are fields below, create the methods in dao\marketrates_impl.go























	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	
logs.Storing("MarketRates",fmt.Sprintf("%s", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.MarketRates_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.MarketRates_Id_sql, r.Id)
	ts = addData(ts, dm.MarketRates_Bid_sql, r.Bid)
	ts = addData(ts, dm.MarketRates_Mid_sql, r.Mid)
	ts = addData(ts, dm.MarketRates_Offer_sql, r.Offer)
	ts = addData(ts, dm.MarketRates_Market_sql, r.Market)
	ts = addData(ts, dm.MarketRates_Tenor_sql, r.Tenor)
	ts = addData(ts, dm.MarketRates_Series_sql, r.Series)
	ts = addData(ts, dm.MarketRates_Name_sql, r.Name)
	ts = addData(ts, dm.MarketRates_Source_sql, r.Source)
	ts = addData(ts, dm.MarketRates_Destination_sql, r.Destination)
	ts = addData(ts, dm.MarketRates_Class_sql, r.Class)
	ts = addData(ts, dm.MarketRates_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.MarketRates_SYSWho_sql, r.SYSWho)
	ts = addData(ts, dm.MarketRates_SYSHost_sql, r.SYSHost)
	ts = addData(ts, dm.MarketRates_Date_sql, r.Date)
	ts = addData(ts, dm.MarketRates_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.MarketRates_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.MarketRates_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.MarketRates_SYSUpdatedBy_sql, r.SYSUpdatedBy)
	ts = addData(ts, dm.MarketRates_SYSUpdatedHost_sql, r.SYSUpdatedHost)
		
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.MarketRates_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	MarketRates_Delete(r.Id)
	das.Execute(tsql)



	return err

}



// marketrates_Fetch read all MarketRates's
func marketrates_Fetch(tsql string) (int, []dm.MarketRates, dm.MarketRates, error) {

	var recItem dm.MarketRates
	var recList []dm.MarketRates

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.MarketRates_SYSId_sql, "0")
	   recItem.Id  = get_String(rec, dm.MarketRates_Id_sql, "")
	   recItem.Bid  = get_String(rec, dm.MarketRates_Bid_sql, "")
	   recItem.Mid  = get_String(rec, dm.MarketRates_Mid_sql, "")
	   recItem.Offer  = get_String(rec, dm.MarketRates_Offer_sql, "")
	   recItem.Market  = get_String(rec, dm.MarketRates_Market_sql, "")
	   recItem.Tenor  = get_String(rec, dm.MarketRates_Tenor_sql, "")
	   recItem.Series  = get_String(rec, dm.MarketRates_Series_sql, "")
	   recItem.Name  = get_String(rec, dm.MarketRates_Name_sql, "")
	   recItem.Source  = get_String(rec, dm.MarketRates_Source_sql, "")
	   recItem.Destination  = get_String(rec, dm.MarketRates_Destination_sql, "")
	   recItem.Class  = get_String(rec, dm.MarketRates_Class_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.MarketRates_SYSCreated_sql, "")
	   recItem.SYSWho  = get_String(rec, dm.MarketRates_SYSWho_sql, "")
	   recItem.SYSHost  = get_String(rec, dm.MarketRates_SYSHost_sql, "")
	   recItem.Date  = get_String(rec, dm.MarketRates_Date_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.MarketRates_SYSUpdated_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.MarketRates_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.MarketRates_SYSCreatedHost_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.MarketRates_SYSUpdatedBy_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.MarketRates_SYSUpdatedHost_sql, "")
	
	// If there are fields below, create the methods in adaptor\MarketRates_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func MarketRates_NewID(r dm.MarketRates) string {
	
			id := uuid.New().String()
	
	return id
}



// marketrates_Fetch read all MarketRates's
func MarketRates_New() (int, []dm.MarketRates, dm.MarketRates, error) {

	var r = dm.MarketRates{}
	var rList []dm.MarketRates
	

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}