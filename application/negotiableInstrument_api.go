package application

// ----------------------------------------------------------------
// Automatically generated  "/application/negotiableinstrument_api.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : NegotiableInstrument (negotiableinstrument)
// Endpoint 	        : NegotiableInstrument (Id)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:16
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	"encoding/json"

	"net/http"


	core    "github.com/mt1976/mwt-go-dev/core"
	dao     "github.com/mt1976/mwt-go-dev/dao"
	dm      "github.com/mt1976/mwt-go-dev/datamodel"
)

//NegotiableInstrument_Handler is the handler for the api calls
func NegotiableInstrument_Handler(w http.ResponseWriter, r *http.Request) {
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
		negotiableinstrument_MethodGet(w, r)

	case http.MethodPost:
		negotiableinstrument_MethodPost(w, r)

	case http.MethodPut:
		negotiableinstrument_MethodPost(w, r)
	case http.MethodDelete:

		negotiableinstrument_MethodDelete(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

// Handles GET requests for NegotiableInstrument
func negotiableinstrument_MethodGet(w http.ResponseWriter, r *http.Request) {
	//logs.Information("PATH", r.URL.Path)
	searchID := core.GetURLparam(r, dm.NegotiableInstrument_QueryString)
	//logs.Information("GET", searchID)
	w.Header().Set("Content-Type", "application/json")

	if searchID == "" {
		//Get All Entities
		//logs.Information("GET", "All")
		noRecs, records, _ := dao.NegotiableInstrument_GetList()
		var ci core.ContentList
		ci.Count = noRecs
		ci.Key = dm.NegotiableInstrument_QueryString
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
		_, record, _ := dao.NegotiableInstrument_GetByID(searchID)
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

//Handles POST & PUT requests for NegotiableInstrument
func negotiableinstrument_MethodPost(w http.ResponseWriter, r *http.Request) {
	//logs.Processing("POST")
	//fmt.Printf("r.Body: %v\n", r.Body)

	decoder := json.NewDecoder(r.Body)
	var t dm.NegotiableInstrument
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
			w.WriteHeader(int(http.StatusNotFound))
	} else {
		w.WriteHeader(int(http.StatusOK))
	}
	//spew.Dump(t)
	err = dao.NegotiableInstrument_StoreSystem(t)
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
//Handles DELETE requests for NegotiableInstrument
func negotiableinstrument_MethodDelete(w http.ResponseWriter, r *http.Request) {
	//logs.Processing("DELETE")
	deleteID := core.GetURLparam(r, dm.NegotiableInstrument_QueryString)
	//logs.Information("DELETE", deleteID)

	dao.NegotiableInstrument_Delete(deleteID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(http.StatusOK))
	//fmt.Printf("json_data: %v\n", json_data)

	//logs.Success("DELETE")
}
