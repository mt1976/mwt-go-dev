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
// Date & Time		    : 12/12/2021 at 16:13:09
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	
	"log"
	
	"fmt"
	"net/http"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
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

	return 1, counterpartyextensionsItem, nil
}



// CounterpartyExtensions_DeleteByID() deletes a single CounterpartyExtensions record
func CounterpartyExtensions_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.CounterpartyExtensions_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.CounterpartyExtensions_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// CounterpartyExtensions_Store() saves/stores a CounterpartyExtensions record to the database
func CounterpartyExtensions_Store(r dm.CounterpartyExtensions,req *http.Request) error {

	err := counterpartyextensions_Save(r,Audit_User(req))

	return err
}

// CounterpartyExtensions_StoreSystem() saves/stores a CounterpartyExtensions record to the database
func CounterpartyExtensions_StoreSystem(r dm.CounterpartyExtensions) error {
	
	err := counterpartyextensions_Save(r,Audit_Host())

	return err
}

// counterpartyextensions_Save() saves/stores a CounterpartyExtensions record to the database
func counterpartyextensions_Save(r dm.CounterpartyExtensions,usr string) error {

    var err error

	logs.Storing("CounterpartyExtensions",fmt.Sprintf("%s", r))

	if len(r.CompID) == 0 {
		r.CompID = CounterpartyExtensions_NewID(r)
	}


// Please Create Functions Below in the adaptor/CounterpartyExtensions_impl.go file
	err1 := adaptor.CounterpartyExtensions_Delete_Impl(r.CompID)
	err2 := adaptor.CounterpartyExtensions_Update_Impl(r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}


// counterpartyextensions_Fetch read all employees
func counterpartyextensions_Fetch(tsql string) (int, []dm.CounterpartyExtensions, dm.CounterpartyExtensions, error) {

	var recItem dm.CounterpartyExtensions
	var recList []dm.CounterpartyExtensions

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.NameFirm  = get_String(rec, dm.CounterpartyExtensions_NameFirm, "")
   recItem.NameCentre  = get_String(rec, dm.CounterpartyExtensions_NameCentre, "")
   recItem.BICCode  = get_String(rec, dm.CounterpartyExtensions_BICCode, "")
   recItem.ContactIndicator  = get_Bool(rec, dm.CounterpartyExtensions_ContactIndicator, "True")
   recItem.CoverTrade  = get_Bool(rec, dm.CounterpartyExtensions_CoverTrade, "True")
   recItem.CustomerCategory  = get_String(rec, dm.CounterpartyExtensions_CustomerCategory, "")
   recItem.FSCSInclusive  = get_Bool(rec, dm.CounterpartyExtensions_FSCSInclusive, "True")
   recItem.FeeFactor  = get_Float(rec, dm.CounterpartyExtensions_FeeFactor, "0.00")
   recItem.InactiveStatus  = get_Bool(rec, dm.CounterpartyExtensions_InactiveStatus, "True")
   recItem.Indemnity  = get_Bool(rec, dm.CounterpartyExtensions_Indemnity, "True")
   recItem.KnowYourCustomerStatus  = get_Bool(rec, dm.CounterpartyExtensions_KnowYourCustomerStatus, "True")
   recItem.LERLimitCarveOut  = get_Bool(rec, dm.CounterpartyExtensions_LERLimitCarveOut, "True")
   recItem.LastAmended  = get_Time(rec, dm.CounterpartyExtensions_LastAmended, "")
   recItem.LastLogin  = get_Time(rec, dm.CounterpartyExtensions_LastLogin, "")
   recItem.LossGivenDefault  = get_Float(rec, dm.CounterpartyExtensions_LossGivenDefault, "0.00")
   recItem.MIC  = get_String(rec, dm.CounterpartyExtensions_MIC, "")
   recItem.ProtectedDepositor  = get_Bool(rec, dm.CounterpartyExtensions_ProtectedDepositor, "True")
   recItem.RPTCurrency  = get_String(rec, dm.CounterpartyExtensions_RPTCurrency, "")
   recItem.RateTimeout  = get_Int(rec, dm.CounterpartyExtensions_RateTimeout, "0")
   recItem.RateValidation  = get_String(rec, dm.CounterpartyExtensions_RateValidation, "")
   recItem.Registered  = get_Bool(rec, dm.CounterpartyExtensions_Registered, "True")
   recItem.RegulatoryCategory  = get_String(rec, dm.CounterpartyExtensions_RegulatoryCategory, "")
   recItem.SecuredSettlement  = get_Bool(rec, dm.CounterpartyExtensions_SecuredSettlement, "True")
   recItem.SettlementLimitCarveOut  = get_Bool(rec, dm.CounterpartyExtensions_SettlementLimitCarveOut, "True")
   recItem.SortCode  = get_String(rec, dm.CounterpartyExtensions_SortCode, "")
   recItem.Training  = get_Bool(rec, dm.CounterpartyExtensions_Training, "True")
   recItem.TrainingCode  = get_String(rec, dm.CounterpartyExtensions_TrainingCode, "")
   recItem.TrainingReceived  = get_Bool(rec, dm.CounterpartyExtensions_TrainingReceived, "True")
   recItem.Unencumbered  = get_Bool(rec, dm.CounterpartyExtensions_Unencumbered, "True")
   recItem.LEIExpiryDate  = get_Time(rec, dm.CounterpartyExtensions_LEIExpiryDate, "")
   recItem.MIFIDReviewDate  = get_Time(rec, dm.CounterpartyExtensions_MIFIDReviewDate, "")
   recItem.GDPRReviewDate  = get_Time(rec, dm.CounterpartyExtensions_GDPRReviewDate, "")
   recItem.DelegatedReporting  = get_String(rec, dm.CounterpartyExtensions_DelegatedReporting, "")
   recItem.BOReconcile  = get_Bool(rec, dm.CounterpartyExtensions_BOReconcile, "True")
   recItem.MIFIDReportableDealsAllowed  = get_Bool(rec, dm.CounterpartyExtensions_MIFIDReportableDealsAllowed, "True")
   recItem.SignedInvestmentAgreement  = get_Bool(rec, dm.CounterpartyExtensions_SignedInvestmentAgreement, "True")
   recItem.AppropriatenessAssessment  = get_Bool(rec, dm.CounterpartyExtensions_AppropriatenessAssessment, "True")
   recItem.FinancialCounterparty  = get_Bool(rec, dm.CounterpartyExtensions_FinancialCounterparty, "True")
   recItem.Collateralisation  = get_String(rec, dm.CounterpartyExtensions_Collateralisation, "")
   recItem.PortfolioCode  = get_String(rec, dm.CounterpartyExtensions_PortfolioCode, "")
   recItem.ReconciliationLetterFrequency  = get_String(rec, dm.CounterpartyExtensions_ReconciliationLetterFrequency, "")
   recItem.DirectDealing  = get_Bool(rec, dm.CounterpartyExtensions_DirectDealing, "True")
   recItem.CompID  = get_String(rec, dm.CounterpartyExtensions_CompID, "")
// If there are fields below, create the methods in dao\CounterpartyExtensions_Impl.go























































































	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func CounterpartyExtensions_NewID(r dm.CounterpartyExtensions) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

