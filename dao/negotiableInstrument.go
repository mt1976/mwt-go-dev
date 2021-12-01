package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/negotiableinstrument.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : NegotiableInstrument (negotiableinstrument)
// Endpoint 	        : NegotiableInstrument (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 01/12/2021 at 20:36:40
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
func NegotiableInstrument_Store(r dm.NegotiableInstrument) error {

	logs.Storing("NegotiableInstrument",fmt.Sprintf("%s", r))

	return nil

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
	// Automatically generated 01/12/2021 by matttownsend on silicon.local - START
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
// Automatically generated 01/12/2021 by matttownsend on silicon.local - END
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

