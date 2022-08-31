package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/counterpartyextensions.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : CounterpartyExtensions (counterpartyextensions)
// Endpoint 	        : CounterpartyExtensions (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:47
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (

	"fmt"
	"net/http"
core "github.com/mt1976/mwt-go-dev/core"
"github.com/google/uuid"
das  "github.com/mt1976/mwt-go-dev/das"
	 adaptor   "github.com/mt1976/mwt-go-dev/adaptor"
	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs   "github.com/mt1976/mwt-go-dev/logs"
)

// CounterpartyExtensions_GetList() returns a list of all CounterpartyExtensions records
func CounterpartyExtensions_GetList() (int, []dm.CounterpartyExtensions, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CounterpartyExtensions_SQLTable)
	count, counterpartyextensionsList, _, _ := counterpartyextensions_Fetch(tsql)
	
	return count, counterpartyextensionsList, nil
}



// CounterpartyExtensions_GetByID() returns a single CounterpartyExtensions record
func CounterpartyExtensions_GetByID(id string) (int, dm.CounterpartyExtensions, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.CounterpartyExtensions_SQLTable)
	tsql = tsql + " WHERE " + dm.CounterpartyExtensions_SQLSearchID + "='" + id + "'"
	_, _, counterpartyextensionsItem, _ := counterpartyextensions_Fetch(tsql)

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, counterpartyextensionsItem, nil
}



// CounterpartyExtensions_DeleteByID() deletes a single CounterpartyExtensions record
func CounterpartyExtensions_Delete(id string) {


	adaptor.CounterpartyExtensions_Delete_impl(id)
	
}


// CounterpartyExtensions_Store() saves/stores a CounterpartyExtensions record to the database
func CounterpartyExtensions_Store(r dm.CounterpartyExtensions,req *http.Request) error {

	err, r := CounterpartyExtensions_Validate(r)
	if err == nil {
		err = counterpartyextensions_Save(r, Audit_User(req))
	} else {
		logs.Information("CounterpartyExtensions_Store()", err.Error())
	}

	return err
}

// CounterpartyExtensions_StoreSystem() saves/stores a CounterpartyExtensions record to the database
func CounterpartyExtensions_StoreSystem(r dm.CounterpartyExtensions) error {
	
	err, r := CounterpartyExtensions_Validate(r)
	if err == nil {
		err = counterpartyextensions_Save(r, Audit_Host())
	} else {
		logs.Information("CounterpartyExtensions_Store()", err.Error())
	}

	return err
}

// CounterpartyExtensions_Validate() validates for saves/stores a CounterpartyExtensions record to the database
func CounterpartyExtensions_Validate(r dm.CounterpartyExtensions) (error,dm.CounterpartyExtensions) {
	var err error
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	//
	

	return err,r
}
//

// counterpartyextensions_Save() saves/stores a CounterpartyExtensions record to the database
func counterpartyextensions_Save(r dm.CounterpartyExtensions,usr string) error {

    var err error



	

	if len(r.CompID) == 0 {
		r.CompID = CounterpartyExtensions_NewID(r)
	}

// If there are fields below, create the methods in dao\counterpartyextensions_impl.go













































	
logs.Storing("CounterpartyExtensions",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/CounterpartyExtensions_impl.go file
	err1 := adaptor.CounterpartyExtensions_Delete_impl(r.CompID)
	err2 := adaptor.CounterpartyExtensions_Update_impl(r.CompID,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// counterpartyextensions_Fetch read all CounterpartyExtensions's
func counterpartyextensions_Fetch(tsql string) (int, []dm.CounterpartyExtensions, dm.CounterpartyExtensions, error) {

	var recItem dm.CounterpartyExtensions
	var recList []dm.CounterpartyExtensions

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.NameFirm  = get_String(rec, dm.CounterpartyExtensions_NameFirm_sql, "")
	   recItem.NameCentre  = get_String(rec, dm.CounterpartyExtensions_NameCentre_sql, "")
	   recItem.BICCode  = get_String(rec, dm.CounterpartyExtensions_BICCode_sql, "")
	   recItem.ContactIndicator  = get_Bool(rec, dm.CounterpartyExtensions_ContactIndicator_sql, "True")
	   recItem.CoverTrade  = get_Bool(rec, dm.CounterpartyExtensions_CoverTrade_sql, "True")
	   recItem.CustomerCategory  = get_String(rec, dm.CounterpartyExtensions_CustomerCategory_sql, "")
	   recItem.FSCSInclusive  = get_Bool(rec, dm.CounterpartyExtensions_FSCSInclusive_sql, "True")
	   recItem.FeeFactor  = get_Float(rec, dm.CounterpartyExtensions_FeeFactor_sql, "0.00")
	   recItem.InactiveStatus  = get_Bool(rec, dm.CounterpartyExtensions_InactiveStatus_sql, "True")
	   recItem.Indemnity  = get_Bool(rec, dm.CounterpartyExtensions_Indemnity_sql, "True")
	   recItem.KnowYourCustomerStatus  = get_Bool(rec, dm.CounterpartyExtensions_KnowYourCustomerStatus_sql, "True")
	   recItem.LERLimitCarveOut  = get_Bool(rec, dm.CounterpartyExtensions_LERLimitCarveOut_sql, "True")
	   recItem.LastAmended  = get_Time(rec, dm.CounterpartyExtensions_LastAmended_sql, "")
	   recItem.LastLogin  = get_Time(rec, dm.CounterpartyExtensions_LastLogin_sql, "")
	   recItem.LossGivenDefault  = get_Float(rec, dm.CounterpartyExtensions_LossGivenDefault_sql, "0.00")
	   recItem.MIC  = get_String(rec, dm.CounterpartyExtensions_MIC_sql, "")
	   recItem.ProtectedDepositor  = get_Bool(rec, dm.CounterpartyExtensions_ProtectedDepositor_sql, "True")
	   recItem.RPTCurrency  = get_String(rec, dm.CounterpartyExtensions_RPTCurrency_sql, "")
	   recItem.RateTimeout  = get_Int(rec, dm.CounterpartyExtensions_RateTimeout_sql, "0")
	   recItem.RateValidation  = get_String(rec, dm.CounterpartyExtensions_RateValidation_sql, "")
	   recItem.Registered  = get_Bool(rec, dm.CounterpartyExtensions_Registered_sql, "True")
	   recItem.RegulatoryCategory  = get_String(rec, dm.CounterpartyExtensions_RegulatoryCategory_sql, "")
	   recItem.SecuredSettlement  = get_Bool(rec, dm.CounterpartyExtensions_SecuredSettlement_sql, "True")
	   recItem.SettlementLimitCarveOut  = get_Bool(rec, dm.CounterpartyExtensions_SettlementLimitCarveOut_sql, "True")
	   recItem.SortCode  = get_String(rec, dm.CounterpartyExtensions_SortCode_sql, "")
	   recItem.Training  = get_Bool(rec, dm.CounterpartyExtensions_Training_sql, "True")
	   recItem.TrainingCode  = get_String(rec, dm.CounterpartyExtensions_TrainingCode_sql, "")
	   recItem.TrainingReceived  = get_Bool(rec, dm.CounterpartyExtensions_TrainingReceived_sql, "True")
	   recItem.Unencumbered  = get_Bool(rec, dm.CounterpartyExtensions_Unencumbered_sql, "True")
	   recItem.LEIExpiryDate  = get_Time(rec, dm.CounterpartyExtensions_LEIExpiryDate_sql, "")
	   recItem.MIFIDReviewDate  = get_Time(rec, dm.CounterpartyExtensions_MIFIDReviewDate_sql, "")
	   recItem.GDPRReviewDate  = get_Time(rec, dm.CounterpartyExtensions_GDPRReviewDate_sql, "")
	   recItem.DelegatedReporting  = get_String(rec, dm.CounterpartyExtensions_DelegatedReporting_sql, "")
	   recItem.BOReconcile  = get_Bool(rec, dm.CounterpartyExtensions_BOReconcile_sql, "True")
	   recItem.MIFIDReportableDealsAllowed  = get_Bool(rec, dm.CounterpartyExtensions_MIFIDReportableDealsAllowed_sql, "True")
	   recItem.SignedInvestmentAgreement  = get_Bool(rec, dm.CounterpartyExtensions_SignedInvestmentAgreement_sql, "True")
	   recItem.AppropriatenessAssessment  = get_Bool(rec, dm.CounterpartyExtensions_AppropriatenessAssessment_sql, "True")
	   recItem.FinancialCounterparty  = get_Bool(rec, dm.CounterpartyExtensions_FinancialCounterparty_sql, "True")
	   recItem.Collateralisation  = get_String(rec, dm.CounterpartyExtensions_Collateralisation_sql, "")
	   recItem.PortfolioCode  = get_String(rec, dm.CounterpartyExtensions_PortfolioCode_sql, "")
	   recItem.ReconciliationLetterFrequency  = get_String(rec, dm.CounterpartyExtensions_ReconciliationLetterFrequency_sql, "")
	   recItem.DirectDealing  = get_Bool(rec, dm.CounterpartyExtensions_DirectDealing_sql, "True")
	   recItem.CompID  = get_String(rec, dm.CounterpartyExtensions_CompID_sql, "")
	
	// If there are fields below, create the methods in adaptor\CounterpartyExtensions_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	///
	//Add to the list
	//
		recList = append(recList, recItem)
	}

	return noitems, recList, recItem, nil
}
	


func CounterpartyExtensions_NewID(r dm.CounterpartyExtensions) string {
	
			id := uuid.New().String()
	
	return id
}



// counterpartyextensions_Fetch read all CounterpartyExtensions's
func CounterpartyExtensions_New() (int, []dm.CounterpartyExtensions, dm.CounterpartyExtensions, error) {

	var r = dm.CounterpartyExtensions{}
	var rList []dm.CounterpartyExtensions
	

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}