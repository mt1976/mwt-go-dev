package dao
// ----------------------------------------------------------------
// Automatically generated  "/dao/product.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			    : Product (product)
// Endpoint 	        : Product (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:18
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

// Product_GetList() returns a list of all Product records
func Product_GetList() (int, []dm.Product, error) {
	
	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Product_SQLTable)
	count, productList, _, _ := product_Fetch(tsql)
	
	return count, productList, nil
}



// Product_GetByID() returns a single Product record
func Product_GetByID(id string) (int, dm.Product, error) {


	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Product_SQLTable)
	tsql = tsql + " WHERE " + dm.Product_SQLSearchID + "='" + id + "'"
	_, _, productItem, _ := product_Fetch(tsql)

	return 1, productItem, nil
}

// Product_GetByReverseLookup(id string) returns a single Product record
func Product_GetByReverseLookup(id string) (int, dm.Product, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Product_SQLTable)
	tsql = tsql + " WHERE Name = '" + id + "'"

	_, _, productItem, _ := product_Fetch(tsql)
	
	return 1, productItem, nil
}

// Product_DeleteByID() deletes a single Product record
func Product_Delete(id string) {


	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Product_SQLTable
	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.Product_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)

}


// Product_Store() saves/stores a Product record to the database
func Product_Store(r dm.Product,req *http.Request) error {

	err := product_Save(r,Audit_User(req))

	return err
}

// Product_StoreSystem() saves/stores a Product record to the database
func Product_StoreSystem(r dm.Product) error {
	
	err := product_Save(r,Audit_Host())

	return err
}

// product_Save() saves/stores a Product record to the database
func product_Save(r dm.Product,usr string) error {

    var err error

	logs.Storing("Product",fmt.Sprintf("%s", r))

	if len(r.Code) == 0 {
		r.Code = Product_NewID(r)
	}


// Please Create Functions Below in the adaptor/Product_impl.go file
	err1 := adaptor.Product_Delete_Impl(r.Code)
	err2 := adaptor.Product_Update_Impl(r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}


// product_Fetch read all employees
func product_Fetch(tsql string) (int, []dm.Product, dm.Product, error) {

	var recItem dm.Product
	var recList []dm.Product

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// Automatically generated 12/12/2021 by matttownsend on silicon.local - START
   recItem.Code  = get_String(rec, dm.Product_Code, "")
   recItem.Name  = get_String(rec, dm.Product_Name, "")
   recItem.Factor  = get_Float(rec, dm.Product_Factor, "0.00")
   recItem.MaxTermPrecedence  = get_Bool(rec, dm.Product_MaxTermPrecedence, "True")
   recItem.InternalId  = get_Int(rec, dm.Product_InternalId, "0")
   recItem.InternalDeleted  = get_Time(rec, dm.Product_InternalDeleted, "")
   recItem.UpdatedTransactionId  = get_String(rec, dm.Product_UpdatedTransactionId, "")
   recItem.UpdatedUserId  = get_String(rec, dm.Product_UpdatedUserId, "")
   recItem.UpdatedDateTime  = get_Time(rec, dm.Product_UpdatedDateTime, "")
   recItem.DeletedTransactionId  = get_String(rec, dm.Product_DeletedTransactionId, "")
   recItem.DeletedUserId  = get_String(rec, dm.Product_DeletedUserId, "")
   recItem.ChangeType  = get_String(rec, dm.Product_ChangeType, "")
// If there are fields below, create the methods in dao\Product_Impl.go

























	// Automatically generated 12/12/2021 by matttownsend on silicon.local - END
		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func Product_NewID(r dm.Product) string {
	
			id := uuid.New().String()
	
	return id
}
// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------

