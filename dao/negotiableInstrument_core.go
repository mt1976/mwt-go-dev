package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/negotiableinstrument.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : NegotiableInstrument (negotiableinstrument)
// Endpoint 	        : NegotiableInstrument (Id)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 26/06/2022 at 18:48:29
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

// NegotiableInstrument_GetList() returns a list of all NegotiableInstrument records
func NegotiableInstrument_GetList() (int, []dm.NegotiableInstrument, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.NegotiableInstrument_SQLTable)
	count, negotiableinstrumentList, _, _ := negotiableinstrument_Fetch(tsql)
	
	return count, negotiableinstrumentList, nil
}



// NegotiableInstrument_GetByID() returns a single NegotiableInstrument record
func NegotiableInstrument_GetByID(id string) (int, dm.NegotiableInstrument, error) {


	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.NegotiableInstrument_SQLTable)
	tsql = tsql + " WHERE " + dm.NegotiableInstrument_SQLSearchID + "='" + id + "'"
	_, _, negotiableinstrumentItem, _ := negotiableinstrument_Fetch(tsql)

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, negotiableinstrumentItem, nil
}



// NegotiableInstrument_DeleteByID() deletes a single NegotiableInstrument record
func NegotiableInstrument_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.NegotiableInstrument_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.NegotiableInstrument_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// NegotiableInstrument_Store() saves/stores a NegotiableInstrument record to the database
func NegotiableInstrument_Store(r dm.NegotiableInstrument,req *http.Request) error {

	err, r := NegotiableInstrument_Validate(r)
	if err == nil {
		err = negotiableinstrument_Save(r, Audit_User(req))
	} else {
		logs.Information("NegotiableInstrument_Store()", err.Error())
	}

	return err
}

// NegotiableInstrument_StoreSystem() saves/stores a NegotiableInstrument record to the database
func NegotiableInstrument_StoreSystem(r dm.NegotiableInstrument) error {
	
	err, r := NegotiableInstrument_Validate(r)
	if err == nil {
		err = negotiableinstrument_Save(r, Audit_Host())
	} else {
		logs.Information("NegotiableInstrument_Store()", err.Error())
	}

	return err
}

// NegotiableInstrument_Validate() validates for saves/stores a NegotiableInstrument record to the database
func NegotiableInstrument_Validate(r dm.NegotiableInstrument) (error,dm.NegotiableInstrument) {
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return nil,r
}
//

// negotiableinstrument_Save() saves/stores a NegotiableInstrument record to the database
func negotiableinstrument_Save(r dm.NegotiableInstrument,usr string) error {

    var err error



	

	if len(r.Id) == 0 {
		r.Id = NegotiableInstrument_NewID(r)
	}

// If there are fields below, create the methods in dao\negotiableinstrument_impl.go






































	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())
	
logs.Storing("NegotiableInstrument",fmt.Sprintf("%s", r))

//Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	ts = addData(ts, dm.NegotiableInstrument_SYSId_sql, r.SYSId)
	ts = addData(ts, dm.NegotiableInstrument_Id_sql, r.Id)
	ts = addData(ts, dm.NegotiableInstrument_LongName_sql, r.LongName)
	ts = addData(ts, dm.NegotiableInstrument_Isin_sql, r.Isin)
	ts = addData(ts, dm.NegotiableInstrument_Tidm_sql, r.Tidm)
	ts = addData(ts, dm.NegotiableInstrument_Sedol_sql, r.Sedol)
	ts = addData(ts, dm.NegotiableInstrument_IssueDate_sql, r.IssueDate)
	ts = addData(ts, dm.NegotiableInstrument_MaturityDate_sql, r.MaturityDate)
	ts = addData(ts, dm.NegotiableInstrument_CouponValue_sql, r.CouponValue)
	ts = addData(ts, dm.NegotiableInstrument_CouponType_sql, r.CouponType)
	ts = addData(ts, dm.NegotiableInstrument_Segment_sql, r.Segment)
	ts = addData(ts, dm.NegotiableInstrument_Sector_sql, r.Sector)
	ts = addData(ts, dm.NegotiableInstrument_CodeConventionCalculateAccrual_sql, r.CodeConventionCalculateAccrual)
	ts = addData(ts, dm.NegotiableInstrument_MinimumDenomination_sql, r.MinimumDenomination)
	ts = addData(ts, dm.NegotiableInstrument_DenominationCurrency_sql, r.DenominationCurrency)
	ts = addData(ts, dm.NegotiableInstrument_TradingCurrency_sql, r.TradingCurrency)
	ts = addData(ts, dm.NegotiableInstrument_Type_sql, r.Type)
	ts = addData(ts, dm.NegotiableInstrument_FlatYield_sql, r.FlatYield)
	ts = addData(ts, dm.NegotiableInstrument_PaymentCouponDate_sql, r.PaymentCouponDate)
	ts = addData(ts, dm.NegotiableInstrument_PeriodOfCoupon_sql, r.PeriodOfCoupon)
	ts = addData(ts, dm.NegotiableInstrument_ExCouponDate_sql, r.ExCouponDate)
	ts = addData(ts, dm.NegotiableInstrument_DateOfIndexInflation_sql, r.DateOfIndexInflation)
	ts = addData(ts, dm.NegotiableInstrument_UnitOfQuotation_sql, r.UnitOfQuotation)
	ts = addData(ts, dm.NegotiableInstrument_SYSCreated_sql, r.SYSCreated)
	ts = addData(ts, dm.NegotiableInstrument_SYSWho_sql, r.SYSWho)
	ts = addData(ts, dm.NegotiableInstrument_SYSHost_sql, r.SYSHost)
	ts = addData(ts, dm.NegotiableInstrument_SYSUpdated_sql, r.SYSUpdated)
	ts = addData(ts, dm.NegotiableInstrument_Issuer_sql, r.Issuer)
	ts = addData(ts, dm.NegotiableInstrument_IssueAmount_sql, r.IssueAmount)
	ts = addData(ts, dm.NegotiableInstrument_RunningYield_sql, r.RunningYield)
	ts = addData(ts, dm.NegotiableInstrument_LEI_sql, r.LEI)
	ts = addData(ts, dm.NegotiableInstrument_CUSIP_sql, r.CUSIP)
	ts = addData(ts, dm.NegotiableInstrument_SYSUpdatedHost_sql, r.SYSUpdatedHost)
	ts = addData(ts, dm.NegotiableInstrument_SYSCreatedBy_sql, r.SYSCreatedBy)
	ts = addData(ts, dm.NegotiableInstrument_SYSCreatedHost_sql, r.SYSCreatedHost)
	ts = addData(ts, dm.NegotiableInstrument_SYSUpdatedBy_sql, r.SYSUpdatedBy)
		
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.NegotiableInstrument_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	NegotiableInstrument_Delete(r.Id)
	das.Execute(tsql)



	return err

}



// negotiableinstrument_Fetch read all NegotiableInstrument's
func negotiableinstrument_Fetch(tsql string) (int, []dm.NegotiableInstrument, dm.NegotiableInstrument, error) {

	var recItem dm.NegotiableInstrument
	var recList []dm.NegotiableInstrument

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.SYSId  = get_Int(rec, dm.NegotiableInstrument_SYSId_sql, "0")
	   recItem.Id  = get_String(rec, dm.NegotiableInstrument_Id_sql, "")
	   recItem.LongName  = get_String(rec, dm.NegotiableInstrument_LongName_sql, "")
	   recItem.Isin  = get_String(rec, dm.NegotiableInstrument_Isin_sql, "")
	   recItem.Tidm  = get_String(rec, dm.NegotiableInstrument_Tidm_sql, "")
	   recItem.Sedol  = get_String(rec, dm.NegotiableInstrument_Sedol_sql, "")
	   recItem.IssueDate  = get_String(rec, dm.NegotiableInstrument_IssueDate_sql, "")
	   recItem.MaturityDate  = get_String(rec, dm.NegotiableInstrument_MaturityDate_sql, "")
	   recItem.CouponValue  = get_String(rec, dm.NegotiableInstrument_CouponValue_sql, "")
	   recItem.CouponType  = get_String(rec, dm.NegotiableInstrument_CouponType_sql, "")
	   recItem.Segment  = get_String(rec, dm.NegotiableInstrument_Segment_sql, "")
	   recItem.Sector  = get_String(rec, dm.NegotiableInstrument_Sector_sql, "")
	   recItem.CodeConventionCalculateAccrual  = get_String(rec, dm.NegotiableInstrument_CodeConventionCalculateAccrual_sql, "")
	   recItem.MinimumDenomination  = get_String(rec, dm.NegotiableInstrument_MinimumDenomination_sql, "")
	   recItem.DenominationCurrency  = get_String(rec, dm.NegotiableInstrument_DenominationCurrency_sql, "")
	   recItem.TradingCurrency  = get_String(rec, dm.NegotiableInstrument_TradingCurrency_sql, "")
	   recItem.Type  = get_String(rec, dm.NegotiableInstrument_Type_sql, "")
	   recItem.FlatYield  = get_String(rec, dm.NegotiableInstrument_FlatYield_sql, "")
	   recItem.PaymentCouponDate  = get_String(rec, dm.NegotiableInstrument_PaymentCouponDate_sql, "")
	   recItem.PeriodOfCoupon  = get_String(rec, dm.NegotiableInstrument_PeriodOfCoupon_sql, "")
	   recItem.ExCouponDate  = get_String(rec, dm.NegotiableInstrument_ExCouponDate_sql, "")
	   recItem.DateOfIndexInflation  = get_String(rec, dm.NegotiableInstrument_DateOfIndexInflation_sql, "")
	   recItem.UnitOfQuotation  = get_String(rec, dm.NegotiableInstrument_UnitOfQuotation_sql, "")
	   recItem.SYSCreated  = get_String(rec, dm.NegotiableInstrument_SYSCreated_sql, "")
	   recItem.SYSWho  = get_String(rec, dm.NegotiableInstrument_SYSWho_sql, "")
	   recItem.SYSHost  = get_String(rec, dm.NegotiableInstrument_SYSHost_sql, "")
	   recItem.SYSUpdated  = get_String(rec, dm.NegotiableInstrument_SYSUpdated_sql, "")
	   recItem.Issuer  = get_String(rec, dm.NegotiableInstrument_Issuer_sql, "")
	   recItem.IssueAmount  = get_String(rec, dm.NegotiableInstrument_IssueAmount_sql, "")
	   recItem.RunningYield  = get_String(rec, dm.NegotiableInstrument_RunningYield_sql, "")
	   recItem.LEI  = get_String(rec, dm.NegotiableInstrument_LEI_sql, "")
	   recItem.CUSIP  = get_String(rec, dm.NegotiableInstrument_CUSIP_sql, "")
	   recItem.SYSUpdatedHost  = get_String(rec, dm.NegotiableInstrument_SYSUpdatedHost_sql, "")
	   recItem.SYSCreatedBy  = get_String(rec, dm.NegotiableInstrument_SYSCreatedBy_sql, "")
	   recItem.SYSCreatedHost  = get_String(rec, dm.NegotiableInstrument_SYSCreatedHost_sql, "")
	   recItem.SYSUpdatedBy  = get_String(rec, dm.NegotiableInstrument_SYSUpdatedBy_sql, "")
	
	// If there are fields below, create the methods in adaptor\NegotiableInstrument_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
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
	


func NegotiableInstrument_NewID(r dm.NegotiableInstrument) string {
	
			id := uuid.New().String()
	
	return id
}



// negotiableinstrument_Fetch read all NegotiableInstrument's
func NegotiableInstrument_New() (int, []dm.NegotiableInstrument, dm.NegotiableInstrument, error) {

	var r = dm.NegotiableInstrument{}
	var rList []dm.NegotiableInstrument
	

	// START
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 26/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}