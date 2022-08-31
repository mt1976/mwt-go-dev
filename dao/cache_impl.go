package dao

import (
	"database/sql"
	"fmt"
	"log"

	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// getCacheStoreList read all employees
func Cache_GetListByObject(id string) (int, []dm.Cache, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationSQLSchema(), dm.Cache_SQLTable)
	tsql = tsql + " WHERE " + dm.Cache_Object_sql + "='" + id + "'"

	//log.Println(tsql)
	count, returnItem, _, _ := cache_Fetch(tsql)
	//log.Println(count, returnItem)
	return count, returnItem, nil
}

// getCacheStoreList read all employees
//func HousekeepCacheStore() (int, error) {
//	expiry := time.Now().Format(core.DATETIMEFORMATSQLSERVER)
//	deletesql := fmt.Sprintf(appCacheStoreSQLDELETEEXPIRED, expiry)
//	log.Println("DELETE:", deletesql, core.ApplicationDB)
//	_, err := core.ApplicationDB.Exec(deletesql)
//	return 0, err
//}

func Cache_GetSample(o string, f string) string {
	// returns a random cache item
	return ""
}

func Cache_GetSampleAll(o string, f string) string {
	// returns a random cache item
	return ""
}

func Cache_StoreItem(o string, f string, v string, sourceDB string) {
	// returns a random cache item
	var r dm.Cache
	r.Object = o
	r.Field = f
	r.Value = v
	r.Source = sourceDB
	r.Id = newCacheStoreID(r)

	Cache_StoreSystem(r)

}

//func Cache_Store_Impl(r dm.Cache) {
//	Cache_StoreSystem(r)
//}

//func Cache_DeleteItemByID(id string) {
//	Cache_Delete(id)
//}

func newCacheStoreID(r dm.Cache) string {
	id := r.Object + core.IDSep + r.Field + core.IDSep + r.Value + core.IDSep + r.Source
	return id
}

//DataCache_Build builds a cache of data from a particular siena table & field.
func DataCache_Build(table string, field string, fieldName string, objectName string, tsql string, sourceProperties map[string]string) {
	//log.Println(tsql)
	//var appCredentialsStore appCredentialsStoreItem
	//var appCredentialsStoreList []appCredentialsStoreItem
	var sqlDataItem sql.NullString
	sourceDB, _ := sourceConnect(sourceProperties)
	rows, err := sourceDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows1: " + err.Error())
		return
	}
	//fmt.Println(rows)
	defer rows.Close()
	//count := 0
	for rows.Next() {
		err := rows.Scan(&sqlDataItem)
		if err != nil {
			log.Println("Error reading rows2: " + err.Error())
			return
		}
		storeField := fieldName
		if len(fieldName) == 0 {
			storeField = field
		}
		storeObjectName := objectName
		if len(objectName) == 0 {
			storeObjectName = table
		}

		Cache_StoreItem(storeObjectName, storeField, sqlDataItem.String, sourceProperties["database"])
		//appCredentialsStore.Id =
		// no change below
		//appCredentialsStoreList = append(appCredentialsStoreList, appCredentialsStore)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		//count++
	}
	//log.Println(count, appCredentialsStoreList, appCredentialsStore)
}

func sourceConnect(sourceProperties map[string]string) (*sql.DB, error) {
	// Connect to SQL Server DB

	server := sourceProperties["server"]
	user := sourceProperties["user"]
	password := sourceProperties["password"]
	port := sourceProperties["port"]
	database := sourceProperties["database"]
	/*
		fmt.Println("SQL SERVER - CONNECTING...")
		fmt.Println("Server     :", server)
		fmt.Println("User       :", user)
		fmt.Println("Password   :", strings.Repeat("*", len(password)))
		fmt.Println("Port       :", port)
		fmt.Println("Database   :", database)
		fmt.Println("")
	*/
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		server, user, password, port, database)
	dbInstance, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	//fmt.Printf("Connected!\n")
	//fmt.Println("")

	//defer dbInstance.Close()

	stmt, _ := dbInstance.Prepare("select @@version")
	row := stmt.QueryRow()
	var result string

	err = row.Scan(&result)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	//fmt.Printf("%s\n", result)
	return dbInstance, err
}
