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
// Date & Time		    : 28/06/2022 at 16:10:55
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

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	return 1, productItem, nil
}

// Product_GetByReverseLookup(id string) returns a single Product record using the Name field as key.
func Product_GetByReverseLookup(id string) (int, dm.Product, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.Product_SQLTable)
	tsql = tsql + " WHERE Name = '" + id + "'"

	_, _, productItem, _ := product_Fetch(tsql)
	
	return 1, productItem, nil
}

// Product_DeleteByID() deletes a single Product record
func Product_Delete(id string) {


	adaptor.Product_Delete_impl(id)
	
}


// Product_Store() saves/stores a Product record to the database
func Product_Store(r dm.Product,req *http.Request) error {

	err, r := Product_Validate(r)
	if err == nil {
		err = product_Save(r, Audit_User(req))
	} else {
		logs.Information("Product_Store()", err.Error())
	}

	return err
}

// Product_StoreSystem() saves/stores a Product record to the database
func Product_StoreSystem(r dm.Product) error {
	
	err, r := Product_Validate(r)
	if err == nil {
		err = product_Save(r, Audit_Host())
	} else {
		logs.Information("Product_Store()", err.Error())
	}

	return err
}

// Product_Validate() validates for saves/stores a Product record to the database
func Product_Validate(r dm.Product) (error,dm.Product) {
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

// product_Save() saves/stores a Product record to the database
func product_Save(r dm.Product,usr string) error {

    var err error



	

	if len(r.Code) == 0 {
		r.Code = Product_NewID(r)
	}

// If there are fields below, create the methods in dao\product_impl.go














	
logs.Storing("Product",fmt.Sprintf("%s", r))

// Please Create Functions Below in the adaptor/Product_impl.go file
	err1 := adaptor.Product_Delete_impl(r.Code)
	err2 := adaptor.Product_Update_impl(r.Code,r,usr)
	if err1 != nil {
		err = err1
	}
	if err2 != nil {
		err = err2
	}


	return err

}



// product_Fetch read all Product's
func product_Fetch(tsql string) (int, []dm.Product, dm.Product, error) {

	var recItem dm.Product
	var recList []dm.Product

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		logs.Fatal(err.Error(),err)
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	   recItem.Code  = get_String(rec, dm.Product_Code_sql, "")
	   recItem.Name  = get_String(rec, dm.Product_Name_sql, "")
	   recItem.Factor  = get_Float(rec, dm.Product_Factor_sql, "0.00")
	   recItem.MaxTermPrecedence  = get_Bool(rec, dm.Product_MaxTermPrecedence_sql, "True")
	   recItem.InternalId  = get_Int(rec, dm.Product_InternalId_sql, "0")
	   recItem.InternalDeleted  = get_Time(rec, dm.Product_InternalDeleted_sql, "")
	   recItem.UpdatedTransactionId  = get_String(rec, dm.Product_UpdatedTransactionId_sql, "")
	   recItem.UpdatedUserId  = get_String(rec, dm.Product_UpdatedUserId_sql, "")
	   recItem.UpdatedDateTime  = get_Time(rec, dm.Product_UpdatedDateTime_sql, "")
	   recItem.DeletedTransactionId  = get_String(rec, dm.Product_DeletedTransactionId_sql, "")
	   recItem.DeletedUserId  = get_String(rec, dm.Product_DeletedUserId_sql, "")
	   recItem.ChangeType  = get_String(rec, dm.Product_ChangeType_sql, "")
	
	// If there are fields below, create the methods in adaptor\Product_impl.go
	
	
	
	
	
	
	
	
	
	
	
	
	
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
	


func Product_NewID(r dm.Product) string {
	
			id := uuid.New().String()
	
	return id
}



// product_Fetch read all Product's
func Product_New() (int, []dm.Product, dm.Product, error) {

	var r = dm.Product{}
	var rList []dm.Product
	

	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END


	rList = append(rList, r)

	return 1, rList, r, nil
}