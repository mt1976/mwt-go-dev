package adaptor

import (
	//"encoding/xml"
	//"io/ioutil"
	//"os"

	"github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	"github.com/mt1976/mwt-go-dev/logs"
)

// Catalog_GetList() returns a list of all API records
func Catalog_GetList_Impl() (int, []dm.Catalog, error) {

	//	count, apiList, _, _ := adaptor.Catalog_GetList_Impl()
	var apiList []dm.Catalog
	count := len(core.Catalog)
	for _, v := range core.Catalog {
		apiList = append(apiList, dm.Catalog{ID: v.ID, Descr: v.Descr, Endpoint: v.Path, Query: v.Query, Source: v.Source})
	}
	return count, apiList, nil
}

// Catalog_GetByID() returns a single API record
func Catalog_GetByID_Impl(id string) (int, dm.Catalog, error) {

	//_, _, apiItem, _ := adaptor.Catalog_GetByID_Impl(id)
	var apiItem dm.Catalog

	for _, v := range core.Catalog {
		if v.ID == id {
			apiItem = dm.Catalog{ID: v.ID, Descr: v.Descr, Endpoint: v.Path, Query: v.Query, Source: v.Source}
		}
	}
	//fmt.Printf("apiItem: %v\n", apiItem)
	return 1, apiItem, nil
}

// Catalog_DeleteByID() deletes a single API record
func Catalog_Delete_Impl(id string) error {
	logs.Information("Catalog_Delete_Impl()", "NO ACTION")
	//adaptor.Catalog_Delete_Impl(id)
	return nil
}

// Catalog_DeleteByID() deletes a single API record
func Catalog_Update_Impl(r dm.Catalog, usr string) error {
	logs.Information("Catalog_Delete_Impl()", "NO ACTION")
	//adaptor.Catalog_Delete_Impl(id)
	return nil
}

// Catalog_DeleteByID() deletes a single API record
func Catalog_NewID_Impl(r dm.Catalog) string {
	logs.Information("Catalog_Delete_Impl()", "NO ACTION")
	//adaptor.Catalog_Delete_Impl(id)
	return r.ID
}
