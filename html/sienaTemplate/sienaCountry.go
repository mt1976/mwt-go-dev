package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//sienaCountryPage is cheese
type sienaCountryListPage struct {
	Title             string
	PageTitle         string
	SienaCountryCount int
	SienaCountryList  []sienaCountryItem
}

//sienaCountryPage is cheese
type sienaCountryPage struct {
	Title     string
	PageTitle string
	ID        string
	Code      string
	Name      string
	ShortCode string
	EU_EEA    string
}

//sienaCountryItem is cheese
type sienaCountryItem struct {
	Code      string
	Name      string
	ShortCode string
	EU_EEA    string
	Action    string
}

func listSienaCountryHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "listSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCountryItem
	noItems, returnList, _ := getSienaCountryList(thisConnection)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaCountryList := sienaCountryListPage{
		Title:             wctProperties["appname"],
		PageTitle:         "List Siena Countrys",
		SienaCountryCount: noItems,
		SienaCountryList:  returnList,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCountryList)

}

func viewSienaCountryHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "viewSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCountryItem
	searchID := getURLparam(r, "sienaCountry")
	noItems, returnRecord, _ := getSienaCountry(thisConnection, searchID)
	fmt.Println("NoSienaCountries", noItems)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	pageSienaCountryList := sienaCountryPage{
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Country",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
		ShortCode: returnRecord.ShortCode,
		EU_EEA:    returnRecord.EU_EEA,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCountryList)

}

func editSienaCountryHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "editSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaCountryItem
	searchID := getURLparam(r, "sienaCountry")
	noItems, returnRecord, _ := getSienaCountry(thisConnection, searchID)
	fmt.Println("NoSienaCountries", noItems)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	pageSienaCountryList := sienaCountryPage{
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Country",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
		ShortCode: returnRecord.ShortCode,
		EU_EEA:    returnRecord.EU_EEA,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCountryList)

}

func saveSienaCountryHandler(w http.ResponseWriter, r *http.Request) {

	//wctProperties := getProperties(CONST_CONFIG_FILE)
	//tmpl := "saveSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	//thisConnection, _ := sienaConnect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	//searchID := getURLparam(r, "sienaCountry")
	//noItems, returnRecord, _ := getSienaCountry(thisConnection, searchID)
	//fmt.Println("NoSienaCountries", noItems)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	listSienaCountryHandler(w, r)

}

func newSienaCountryHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "newSienaCountry"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	pageSienaCountryList := sienaCountryPage{
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Country",
		ID:        "NEW",
		Code:      "",
		Name:      "",
		ShortCode: "",
		EU_EEA:    "",
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaCountryList)

}

// getSienaCountryList read all employees
func getSienaCountryList(db *sql.DB) (int, []sienaCountryItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	var sienaCountryList []sienaCountryItem
	var sienaCountry sienaCountryItem
	tsql := fmt.Sprintf("SELECT Code, Name, ShortCode, EU_EEA FROM %s.sienaCountry;", mssqlConfig["schema"])
	//	fmt.Println("MS SQL:", tsql)

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		var code, name, shortcode string
		var eu_eea bool
		err := rows.Scan(&code, &name, &shortcode, &eu_eea)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, err
		}
		sienaCountry.Code = code
		sienaCountry.Name = name
		sienaCountry.ShortCode = shortcode
		sienaCountry.EU_EEA = ""
		if eu_eea {
			sienaCountry.EU_EEA = "checked"
		}
		sienaCountryList = append(sienaCountryList, sienaCountry)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaCountryList, nil
}

// getSienaCountryList read all employees
func getSienaCountry(db *sql.DB, id string) (int, sienaCountryItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	var sienaCountry sienaCountryItem
	tsql := fmt.Sprintf("SELECT Code, Name, ShortCode, EU_EEA FROM %s.sienaCountry WHERE Code='%s';", mssqlConfig["schema"], id)
	fmt.Println("MS SQL:", tsql)

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, sienaCountry, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		var code, name, shortcode string
		var eu_eea bool
		err := rows.Scan(&code, &name, &shortcode, &eu_eea)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, sienaCountry, err
		}
		sienaCountry.Code = code
		sienaCountry.Name = name
		sienaCountry.ShortCode = shortcode
		sienaCountry.EU_EEA = ""
		if eu_eea {
			sienaCountry.EU_EEA = "checked"
		}
		count++
	}
	return 1, sienaCountry, nil
}

// getSienaCountryList read all employees
func putSienaCountry(db *sql.DB, updateItem sienaCountryItem) error {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(mssqlConfig["schema"])
	fmt.Println(updateItem)
	return nil
}
