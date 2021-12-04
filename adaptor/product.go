package adaptor
// ----------------------------------------------------------------
// Automatically generated  "/adaptor/product.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Product (product)
// Endpoint 	        : Product (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r3-21.12.04]
// Date & Time		    : 04/12/2021 at 17:36:48
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm   "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)


func Product_Delete(id string) error {
	var er error

	message:= "Implement Product_Delete: " + id

	// Implement Product_Delete_Impl in product_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Product_Delete_Impl(item)
	//

	logs.Success(message)
	return er
}

func Product_Update(item dm.Product) error {
	var er error

	message:= "Implement Product_Update: " + item.Code

	// Implement Product_Update_Impl in product_Impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Product_Update_Impl(item)
	//

	logs.Success(message)
	return er
}