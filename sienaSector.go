package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//sienaSectorPage is cheese
type sienaSectorListPage struct {
	Title            string
	PageTitle        string
	SienaSectorCount int
	SienaSectorList  []sienaSectorItem
}

//sienaSectorPage is cheese
type sienaSectorPage struct {
	Title     string
	PageTitle string
	ID        string
	Code      string
	Name      string
}

//sienaSectorItem is cheese
type sienaSectorItem struct {
	Code   string
	Name   string
	Action string
}

func listSienaSectorHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "listSienaSector"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	//	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaSectorItem
	noItems, returnList, _ := getSienaSectorList(thisConnection)
	//	fmt.Println("NoSienaCountries", noItems)
	//	fmt.Println(returnList)
	//	fmt.Println(tmpl)

	pageSienaSectorList := sienaSectorListPage{
		Title:            wctProperties["appname"],
		PageTitle:        "List Siena Sectors",
		SienaSectorCount: noItems,
		SienaSectorList:  returnList,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaSectorList)

}

func viewSienaSectorHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "viewSienaSector"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaSectorItem
	searchID := getURLparam(r, "sienaSector")
	noItems, returnRecord, _ := getSienaSector(thisConnection, searchID)
	fmt.Println("NoSienaCountries", noItems)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	pageSienaSectorList := sienaSectorPage{
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Sector",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaSectorList)

}

func editSienaSectorHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "editSienaSector"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	thisConnection, _ := sienaConnect()
	fmt.Println(thisConnection.Stats().OpenConnections)
	var returnList []sienaSectorItem
	searchID := getURLparam(r, "sienaSector")
	noItems, returnRecord, _ := getSienaSector(thisConnection, searchID)
	fmt.Println("NoSienaCountries", noItems)
	fmt.Println(returnList)
	fmt.Println(tmpl)

	pageSienaSectorList := sienaSectorPage{
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Sector",
		ID:        returnRecord.Code,
		Code:      returnRecord.Code,
		Name:      returnRecord.Name,
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaSectorList)

}

func saveSienaSectorHandler(w http.ResponseWriter, r *http.Request) {

	//wctProperties := getProperties(CONST_CONFIG_FILE)
	//tmpl := "saveSienaSector"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	//thisConnection, _ := sienaConnect()
	//fmt.Println(thisConnection.Stats().OpenConnections)
	//searchID := getURLparam(r, "sienaSector")
	//noItems, returnRecord, _ := getSienaSector(thisConnection, searchID)
	//fmt.Println("NoSienaCountries", noItems)
	//fmt.Println(returnList)
	//fmt.Println(tmpl)

	listSienaSectorHandler(w, r)

}

func newSienaSectorHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := getProperties(CONST_CONFIG_FILE)
	tmpl := "newSienaSector"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	pageSienaSectorList := sienaSectorPage{
		Title:     wctProperties["appname"],
		PageTitle: "View Siena Sector",
		ID:        "NEW",
		Code:      "",
		Name:      "",
	}

	t, _ := template.ParseFiles(getTemplateID(tmpl))
	t.Execute(w, pageSienaSectorList)

}

// getSienaSectorList read all employees
func getSienaSectorList(db *sql.DB) (int, []sienaSectorItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	var sienaSectorList []sienaSectorItem
	var sienaSector sienaSectorItem
	tsql := fmt.Sprintf("SELECT Code, Name FROM %s.sienaSector;", mssqlConfig["schema"])
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
		var code, name string
		err := rows.Scan(&code, &name)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, err
		}
		sienaSector.Code = code
		sienaSector.Name = name
		sienaSectorList = append(sienaSectorList, sienaSector)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	return count, sienaSectorList, nil
}

// getSienaSectorList read all employees
func getSienaSector(db *sql.DB, id string) (int, sienaSectorItem, error) {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	var sienaSector sienaSectorItem
	tsql := fmt.Sprintf("SELECT Code, Name FROM %s.sienaSector WHERE Code='%s';", mssqlConfig["schema"], id)
	fmt.Println("MS SQL:", tsql)

	rows, err := db.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, sienaSector, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		var code, name string

		err := rows.Scan(&code, &name)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, sienaSector, err
		}
		sienaSector.Code = code
		sienaSector.Name = name
		count++
	}
	return 1, sienaSector, nil
}

// getSienaSectorList read all employees
func putSienaSector(db *sql.DB, updateItem sienaSectorItem) error {
	mssqlConfig := getProperties(cSQL_CONFIG)
	//fmt.Println(db.Stats().OpenConnections)
	fmt.Println(mssqlConfig["schema"])
	fmt.Println(updateItem)
	return nil
}
