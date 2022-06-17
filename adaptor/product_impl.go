package adaptor

// ----------------------------------------------------------------
// Automatically generated  "/adaptor/product.go"
// ----------------------------------------------------------------
// Package              : adaptor
// Object 			    : Product (product)
// Endpoint 	        : Product (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 08/12/2021 at 15:59:18
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func Product_Delete_impl(id string) error {
	var er error

	message := "Implement Product_Delete: " + id

	// Implement Product_Delete_impl in product_impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Product_Delete_impl(item)
	//

	logs.Success(message)
	return er
}

func Product_Update_impl(id string, item dm.Product, usr string) error {
	var er error

	message := "Implement Product_Update: " + item.Code

	// Implement Product_Update_impl in product_impl.go
	// Uncomment the line below to use the implementation
	//
	// er := Product_Update_impl(item)
	//

	logs.Success(message)
	return er
}
