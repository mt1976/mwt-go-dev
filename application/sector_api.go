package application

// ----------------------------------------------------------------
// Automatically generated  "/application/sector_api.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Sector (sector)
// Endpoint 	        : Sector (Sector)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:57
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"encoding/json"

	"net/http"


	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
)

//Sector_Handler is the handler for the api calls
func Sector_Handler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	//TODO: Add your security validation here
	//        new => POST
	//  	 read => GET
	// 	   update => PUT
	//     delete => DELETE

	httpMethod := r.Method
	
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	//responseStatus := http.StatusOK

	switch httpMethod {
	case http.MethodGet:
		sector_MethodGet(w, r)

	case http.MethodPost:
		sector_MethodPost(w, r)

	case http.MethodPut:
		sector_MethodPost(w, r)
	case http.MethodDelete:

		sector_MethodDelete(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

// Handles GET requests for Sector
func sector_MethodGet(w http.ResponseWriter, r *http.Request) {
	//logs.Information("PATH", r.URL.Path)
	searchID := core.GetURLparam(r, dm.Sector_QueryString)
	//logs.Information("GET", searchID)
	w.Header().Set("Content-Type", "application/json")

	if searchID == "" {
		//Get All Entities
		//logs.Information("GET", "All")
		noRecs, records, _ := dao.Sector_GetList()
		var ci core.ContentList
		ci.Count = noRecs
		ci.Key = dm.Sector_QueryString
		for _, v := range records {
			ciContent := core.ContentListItem{ID:v.Code,Query:"?" + ci.Key +"="+ v.Code}
			ci.Items= append(ci.Items, ciContent)
		}
		json_data, _ := json.Marshal(ci)
		w.Write(json_data)

		if noRecs == 0 {
			w.WriteHeader(int(http.StatusNotFound))
		} else {
			w.WriteHeader(int(http.StatusOK))
		}


	} else {
		//Get a specific entity
		_, record, _ := dao.Sector_GetByID(searchID)
		//spew.Dump(record)
		json_data, _ := json.Marshal(record)
		w.Write(json_data)

		if record.Code == "" {
		    w.WriteHeader(int(http.StatusNotFound))
		} else {
			w.WriteHeader(int(http.StatusOK))
		}
	}


}

//Handles POST & PUT requests for Sector
func sector_MethodPost(w http.ResponseWriter, r *http.Request) {
	//logs.Processing("POST")
	//fmt.Printf("r.Body: %v\n", r.Body)

	decoder := json.NewDecoder(r.Body)
	var t dm.Sector
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(int(http.StatusNotFound))
		panic(err)
	} else {
		w.WriteHeader(int(http.StatusOK))
	}
	//spew.Dump(t)
	
	err = dao.Sector_StoreSystem(t)
	if err != nil {

		//	panic(err)

		w.WriteHeader(int(http.StatusNotFound))

	} else {

		w.WriteHeader(int(http.StatusOK))

	}
	
	//logs.Processing("POST BACK")
	//logs.Information("POST", err.Error())
	
	//logs.Success("POST")
}
//Handles DELETE requests for Sector
func sector_MethodDelete(w http.ResponseWriter, r *http.Request) {
	//logs.Processing("DELETE")
	//logs.Information("DELETE", deleteID)
		w.Header().Set("Content-Type", "application/json")

	deleteID := core.GetURLparam(r, dm.Sector_QueryString)

	dao.Sector_Delete(deleteID)
		w.WriteHeader(int(http.StatusOK))


	//fmt.Printf("json_data: %v\n", json_data)

	//logs.Success("DELETE")
}
