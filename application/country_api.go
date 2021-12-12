package application

// ----------------------------------------------------------------
// Automatically generated  "/application/country_api.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Country (country)
// Endpoint 	        : Country (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:10
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"encoding/json"

	"net/http"


	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
)

//Country_Handler is the handler for the api calls
func Country_Handler(w http.ResponseWriter, r *http.Request) {
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
		country_MethodGet(w, r)

	case http.MethodPost:
		country_MethodPost(w, r)

	case http.MethodPut:
		country_MethodPost(w, r)
	case http.MethodDelete:

		country_MethodDelete(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

// Handles GET requests for Country
func country_MethodGet(w http.ResponseWriter, r *http.Request) {
	//logs.Information("PATH", r.URL.Path)
	searchID := core.GetURLparam(r, dm.Country_QueryString)
	//logs.Information("GET", searchID)
	w.Header().Set("Content-Type", "application/json")

	if searchID == "" {
		//Get All Entities
		//logs.Information("GET", "All")
		noRecs, records, _ := dao.Country_GetList()
		var ci core.ContentList
		ci.Count = noRecs
		ci.Key = dm.Country_QueryString
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
		_, record, _ := dao.Country_GetByID(searchID)
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

//Handles POST & PUT requests for Country
func country_MethodPost(w http.ResponseWriter, r *http.Request) {
	//logs.Processing("POST")
	//fmt.Printf("r.Body: %v\n", r.Body)

	decoder := json.NewDecoder(r.Body)
	var t dm.Country
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
			w.WriteHeader(int(http.StatusNotFound))
	} else {
		w.WriteHeader(int(http.StatusOK))
	}
	//spew.Dump(t)
	err = dao.Country_StoreSystem(t)
	//logs.Processing("POST BACK")
	//logs.Information("POST", err.Error())
	if err != nil {
		//	panic(err)
		w.WriteHeader(int(http.StatusNotFound))
	} else {
		w.WriteHeader(int(http.StatusOK))
	}
	//logs.Success("POST")
}
//Handles DELETE requests for Country
func country_MethodDelete(w http.ResponseWriter, r *http.Request) {
	//logs.Processing("DELETE")
	deleteID := core.GetURLparam(r, dm.Country_QueryString)
	//logs.Information("DELETE", deleteID)

	dao.Country_Delete(deleteID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(http.StatusOK))
	//fmt.Printf("json_data: %v\n", json_data)

	//logs.Success("DELETE")
}
