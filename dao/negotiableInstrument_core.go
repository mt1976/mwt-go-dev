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
// Date & Time		    : 12/12/2021 at 16:13:16
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"log"
	
	"fmt"
	"net/http"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
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

	err := negotiableinstrument_Save(r,Audit_User(req))

	return err
}

// NegotiableInstrument_StoreSystem() saves/stores a NegotiableInstrument record to the database
func NegotiableInstrument_StoreSystem(r dm.NegotiableInstrument) error {
	
	err := negotiableinstrument_Save(r,Audit_Host())

	return err
}

// negotiableinstrument_Save() saves/stores a NegotiableInstrument record to the database
func negotiableinstrument_Save(r dm.NegotiableInstrument,usr string) error {

    var err error

	logs.Storing("NegotiableInstrument",fmt.Sprintf("%s", r))

	if len(r.Id) == 0 {
		r.Id = NegotiableInstrument_NewID(r)
	}


//Deal with the if its Application or null add this bit, otherwise dont.
	
	r.SYSCreated = Audit_Update(r.SYSCreated, Audit_TimeStamp())
	r.SYSCreatedBy = Audit_Update(r.SYSCreatedBy, usr)
	r.SYSCreatedHost = Audit_Update(r.SYSCreatedHost,Audit_Host())
	r.SYSUpdated = Audit_Update("", Audit_TimeStamp())
	r.SYSUpdatedBy = Audit_Update("",usr)
	r.SYSUpdatedHost = Audit_Update("",Audit_Host())

	ts := SQLData{}

ts = addData(ts, dm.NegotiableInstrument_SYSId, r.SYSId)
ts = addData(ts, dm.NegotiableInstrument_Id, r.Id)
ts = addData(ts, dm.NegotiableInstrument_LongName, r.LongName)
ts = addData(ts, dm.NegotiableInstrument_Isin, r.Isin)
ts = addData(ts, dm.NegotiableInstrument_Tidm, r.Tidm)
ts = addData(ts, dm.NegotiableInstrument_Sedol, r.Sedol)
ts = addData(ts, dm.NegotiableInstrument_IssueDate, r.IssueDate)
ts = addData(ts, dm.NegotiableInstrument_MaturityDate, r.MaturityDate)
ts = addData(ts, dm.NegotiableInstrument_CouponValue, r.CouponValue)
ts = addData(ts, dm.NegotiableInstrument_CouponType, r.CouponType)
ts = addData(ts, dm.NegotiableInstrument_Segment, r.Segment)
ts = addData(ts, dm.NegotiableInstrument_Sector, r.Sector)
ts = addData(ts, dm.NegotiableInstrument_CodeConventionCalculateAccrual, r.CodeConventionCalculateAccrual)
ts = addData(ts, dm.NegotiableInstrument_MinimumDenomination, r.MinimumDenomination)
ts = addData(ts, dm.NegotiableInstrument_DenominationCurrency, r.DenominationCurrency)
ts = addData(ts, dm.NegotiableInstrument_TradingCurrency, r.TradingCurrency)
ts = addData(ts, dm.NegotiableInstrument_Type, r.Type)
ts = addData(ts, dm.NegotiableInstrument_FlatYield, r.FlatYield)
ts = addData(ts, dm.NegotiableInstrument_PaymentCouponDate, r.PaymentCouponDate)
ts = addData(ts, dm.NegotiableInstrument_PeriodOfCoupon, r.PeriodOfCoupon)
ts = addData(ts, dm.NegotiableInstrument_ExCouponDate, r.ExCouponDate)
ts = addData(ts, dm.NegotiableInstrument_DateOfIndexInflation, r.DateOfIndexInflation)
ts = addData(ts, dm.NegotiableInstrument_UnitOfQuotation, r.UnitOfQuotation)
ts = addData(ts, dm.NegotiableInstrument_SYSCreated, r.SYSCreated)
ts = addData(ts, dm.NegotiableInstrument_SYSWho, r.SYSWho)
ts = addData(ts, dm.NegotiableInstrument_SYSHost, r.SYSHost)
ts = addData(ts, dm.NegotiableInstrument_SYSUpdated, r.SYSUpdated)
ts = addData(ts, dm.NegotiableInstrument_Issuer, r.Issuer)
ts = addData(ts, dm.NegotiableInstrument_IssueAmount, r.IssueAmount)
ts = addData(ts, dm.NegotiableInstrument_RunningYield, r.RunningYield)
ts = addData(ts, dm.NegotiableInstrument_LEI, r.LEI)
ts = addData(ts, dm.NegotiableInstrument_CUSIP, r.CUSIP)
ts = addData(ts, dm.NegotiableInstrument_SYSUpdatedHost, r.SYSUpdatedHost)
ts = addData(ts, dm.NegotiableInstrument_SYSCreatedBy, r.SYSCreatedBy)
ts = addData(ts, dm.NegotiableInstrument_SYSCreatedHost, r.SYSCreatedHost)
ts = addData(ts, dm.NegotiableInstrument_SYSUpdatedBy, r.SYSUpdatedBy)
	

	tsql := "INSERT INTO " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.NegotiableInstrument_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	NegotiableInstrument_Delete(r.Id)
	das.Execute(tsql)



	return err

}


// negotiableinstrument_Fetch read all employees
func negotiableinstrument_Fetch(tsql string) (int, []dm.NegotiableInstrument, dm.NegotiableInstrument, error) {

	var recItem dm.NegotiableInstrument
	var recList []dm.NegotiableInstrument

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.SYSId  = get_Int(rec, dm.NegotiableInstrument_SYSId, "0")
   recItem.Id  = get_String(rec, dm.NegotiableInstrument_Id, "")
   recItem.LongName  = get_String(rec, dm.NegotiableInstrument_LongName, "")
   recItem.Isin  = get_String(rec, dm.NegotiableInstrument_Isin, "")
   recItem.Tidm  = get_String(rec, dm.NegotiableInstrument_Tidm, "")
   recItem.Sedol  = get_String(rec, dm.NegotiableInstrument_Sedol, "")
   recItem.IssueDate  = get_String(rec, dm.NegotiableInstrument_IssueDate, "")
   recItem.MaturityDate  = get_String(rec, dm.NegotiableInstrument_MaturityDate, "")
   recItem.CouponValue  = get_String(rec, dm.NegotiableInstrument_CouponValue, "")
   recItem.CouponType  = get_String(rec, dm.NegotiableInstrument_CouponType, "")
   recItem.Segment  = get_String(rec, dm.NegotiableInstrument_Segment, "")
   recItem.Sector  = get_String(rec, dm.NegotiableInstrument_Sector, "")
   recItem.CodeConventionCalculateAccrual  = get_String(rec, dm.NegotiableInstrument_CodeConventionCalculateAccrual, "")
   recItem.MinimumDenomination  = get_String(rec, dm.NegotiableInstrument_MinimumDenomination, "")
   recItem.DenominationCurrency  = get_String(rec, dm.NegotiableInstrument_DenominationCurrency, "")
   recItem.TradingCurrency  = get_String(rec, dm.NegotiableInstrument_TradingCurrency, "")
   recItem.Type  = get_String(rec, dm.NegotiableInstrument_Type, "")
   recItem.FlatYield  = get_String(rec, dm.NegotiableInstrument_FlatYield, "")
   recItem.PaymentCouponDate  = get_String(rec, dm.NegotiableInstrument_PaymentCouponDate, "")
   recItem.PeriodOfCoupon  = get_String(rec, dm.NegotiableInstrument_PeriodOfCoupon, "")
   recItem.ExCouponDate  = get_String(rec, dm.NegotiableInstrument_ExCouponDate, "")
   recItem.DateOfIndexInflation  = get_String(rec, dm.NegotiableInstrument_DateOfIndexInflation, "")
   recItem.UnitOfQuotation  = get_String(rec, dm.NegotiableInstrument_UnitOfQuotation, "")
   recItem.SYSCreated  = get_String(rec, dm.NegotiableInstrument_SYSCreated, "")
   recItem.SYSWho  = get_String(rec, dm.NegotiableInstrument_SYSWho, "")
   recItem.SYSHost  = get_String(rec, dm.NegotiableInstrument_SYSHost, "")
   recItem.SYSUpdated  = get_String(rec, dm.NegotiableInstrument_SYSUpdated, "")
   recItem.Issuer  = get_String(rec, dm.NegotiableInstrument_Issuer, "")
   recItem.IssueAmount  = get_String(rec, dm.NegotiableInstrument_IssueAmount, "")
   recItem.RunningYield  = get_String(rec, dm.NegotiableInstrument_RunningYield, "")
   recItem.LEI  = get_String(rec, dm.NegotiableInstrument_LEI, "")
   recItem.CUSIP  = get_String(rec, dm.NegotiableInstrument_CUSIP, "")
   recItem.SYSUpdatedHost  = get_String(rec, dm.NegotiableInstrument_SYSUpdatedHost, "")
   recItem.SYSCreatedBy  = get_String(rec, dm.NegotiableInstrument_SYSCreatedBy, "")
   recItem.SYSCreatedHost  = get_String(rec, dm.NegotiableInstrument_SYSCreatedHost, "")
   recItem.SYSUpdatedBy  = get_String(rec, dm.NegotiableInstrument_SYSUpdatedBy, "")
// If there are fields below, create the methods in dao\NegotiableInstrument_Impl.go









































































	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func NegotiableInstrument_NewID(r dm.NegotiableInstrument) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

