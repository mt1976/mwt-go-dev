package application

// ----------------------------------------------------------------
// Automatically generated  "/application/marketrateshistory_api.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : MarketRatesHistory (marketrateshistory)
// Endpoint 	        : MarketRatesHistory (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : Dysprosium [r4-21.12.31]
// Date & Time		    : 31/08/2022 at 14:19:48
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

import (
	"encoding/json"

	"net/http"


	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
)

//MarketRatesHistory_Handler is the handler for the api calls
func MarketRatesHistory_Handler(w http.ResponseWriter, r *http.Request) {
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
		marketrateshistory_MethodGet(w, r)

	case http.MethodPost:
		marketrateshistory_MethodPost(w, r)

	case http.MethodPut:
		marketrateshistory_MethodPost(w, r)
	case http.MethodDelete:

		marketrateshistory_MethodDelete(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

// Handles GET requests for MarketRatesHistory
func marketrateshistory_MethodGet(w http.ResponseWriter, r *http.Request) {
	//logs.Information("PATH", r.URL.Path)
	searchID := core.GetURLparam(r, dm.MarketRatesHistory_QueryString)
	//logs.Information("GET", searchID)
	w.Header().Set("Content-Type", "application/json")

	if searchID == "" {
		//Get All Entities
		//logs.Information("GET", "All")
		noRecs, records, _ := dao.MarketRatesHistory_GetList()
		var ci core.ContentList
		ci.Count = noRecs
		ci.Key = dm.MarketRatesHistory_QueryString
		for _, v := range records {
			ciContent := core.ContentListItem{ID:v.Id,Query:"?" + ci.Key +"="+ v.Id}
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
		_, record, _ := dao.MarketRatesHistory_GetByID(searchID)
		//spew.Dump(record)
		json_data, _ := json.Marshal(record)
		w.Write(json_data)

		if record.Id == "" {
		    w.WriteHeader(int(http.StatusNotFound))
		} else {
			w.WriteHeader(int(http.StatusOK))
		}
	}


}

//Handles POST & PUT requests for MarketRatesHistory
func marketrateshistory_MethodPost(w http.ResponseWriter, r *http.Request) {
	//logs.Processing("POST")
	//fmt.Printf("r.Body: %v\n", r.Body)

	decoder := json.NewDecoder(r.Body)
	var t dm.MarketRatesHistory
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(int(http.StatusNotFound))
		panic(err)
	} else {
		w.WriteHeader(int(http.StatusOK))
	}
	//spew.Dump(t)
	
			w.WriteHeader(int(http.StatusMethodNotAllowed))
	
	//logs.Processing("POST BACK")
	//logs.Information("POST", err.Error())
	
	//logs.Success("POST")
}
//Handles DELETE requests for MarketRatesHistory
func marketrateshistory_MethodDelete(w http.ResponseWriter, r *http.Request) {
	//logs.Processing("DELETE")
	//logs.Information("DELETE", deleteID)
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(int(http.StatusMethodNotAllowed))


	//fmt.Printf("json_data: %v\n", json_data)

	//logs.Success("DELETE")
}
