package dao

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/andreyvit/sqlexpr"
	"github.com/kisielk/sqlstruct"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// Defines the Fields to Fetch from SQL

//var sqlCacheStoreId, sqlCacheStoreObject, sqlCacheStoreField, sqlCacheStoreValue, sqlCacheStoreExpiry, sqlCacheStoreSYSCreated, sqlCacheStoreSYSWho, sqlCacheStoreSYSHost, sqlCacheStoreSYSUpdated, sqlCacheStoreSource sql.NullString
var dsCache dm.DataStoreMessages

func init() {
	fullTableName := core.ApplicationPropertiesDB["scheme"] + "." + "cacheStore"

	dsCache = dm.DataStoreMessages{
		//Table: fullTableName,
		//Columns: "id, 	object, 	field, 	value, 	expiry, 	_created, 	_who, 	_host, 	_updated, source",
		Insert:    "INSERT INTO " + fullTableName + "(%s) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s');",
		Delete:    "DELETE FROM " + fullTableName + " WHERE id='%s';",
		Select:    "SELECT %s FROM " + fullTableName + ";",
		Get:       "SELECT %s FROM " + fullTableName + " WHERE id='%s';",
		GetAlt:    "SELECT %s FROM " + fullTableName + " WHERE object='%s';",
		DeleteAlt: "DELETE FROM " + fullTableName + " WHERE expires < '%s';",
	}
}

//dm.AppCacheStoreItem is cheese
const (
	//Action     string
	csTable      = sqlexpr.Table("cacheStore")
	csId         = sqlexpr.Column("id")
	csObject     = sqlexpr.Column("object")
	csField      = sqlexpr.Column("field")
	csValue      = sqlexpr.Column("value")
	csExpiry     = sqlexpr.Column("expiry")
	csSYSCreated = sqlexpr.Column("_created")
	csSYSWho     = sqlexpr.Column("_who")
	csSYSHost    = sqlexpr.Column("_host")
	csSYSUpdated = sqlexpr.Column("_updated")
	csSource     = sqlexpr.Column("source")
)

// getCacheStoreList read all employees
func DataCache_GetList() (int, []dm.DataCache, error) {
	tsql := fmt.Sprintf(dsCache.Select, sqlstruct.Columns(dm.DataCache{}))
	count, appCacheStoreList, _, _ := fetchCacheStoreData(tsql)
	return count, appCacheStoreList, nil
}

// getCacheStoreList read all employees
func DataCache_GetItemByID(id string) (int, dm.DataCache, error) {
	tsql := fmt.Sprintf(dsCache.Get, sqlstruct.Columns(dm.DataCache{}), id)
	_, _, returnItem, _ := fetchCacheStoreData(tsql)
	return 1, returnItem, nil
}

// getCacheStoreList read all employees
func DataCache_GetListByObject(object string) (int, []dm.DataCache, error) {
	tsql := fmt.Sprintf(dsCache.GetAlt, sqlstruct.Columns(dm.DataCache{}), object)
	count, returnItem, _, _ := fetchCacheStoreData(tsql)
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

func GetCacheDataSample(o string, f string) string {
	// returns a random cache item
	return ""
}

func GetCacheDataSampleAll(o string, f string) string {
	// returns a random cache item
	return ""
}

func PutCacheDataSampleItem(o string, f string, v string, sourceDB string, x *http.Request) {
	// returns a random cache item
	var r dm.DataCache
	r.Object = o
	r.Field = f
	r.Value = v
	r.Source = sourceDB
	putCacheStore(r, x)

}

func DataCache_Store(r dm.DataCache, req *http.Request) {
	putCacheStore(r, req)
}

func putCacheStore(r dm.DataCache, req *http.Request) {
	//fmt.Println(credentialStore)
	createDate := time.Now().Format(core.DATETIMEFORMATUSER)
	if len(r.SYSCreated) == 0 {
		r.SYSCreated = createDate
	}

	//currentUserID, _ := user.Current()
	//userID := currentUserID.Name
	host, _ := os.Hostname()

	if len(r.Id) == 0 {
		r.Id = newCacheStoreID(r)
		r.SYSCreated = createDate
		r.SYSWho = core.GetUserName(req)
		r.SYSHost = host
	}

	r.SYSUpdated = createDate

	//fmt.Println("RECORD", r)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))
	//spew.Dump(dsCache)

	deletesql := fmt.Sprintf(dsCache.Delete, r.Id)
	//inserttsql := fmt.Sprintf(dsCache.Insert, sqlstruct.Columns(dm.AppCacheStoreItem{}), r.Id, r.Object, r.Field, r.Value, r.Expiry, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated, r.Source)
	//log.Println(sqlstruct.TagName)

	inserttsql := sqlexpr.Insert{Table: csTable}
	inserttsql.Set(csId, r.Id)
	inserttsql.Set(csObject, r.Object)
	inserttsql.Set(csField, r.Field)
	inserttsql.Set(csValue, r.Value)
	inserttsql.Set(csExpiry, r.Expiry)
	inserttsql.Set(csSYSCreated, r.SYSCreated)
	inserttsql.Set(csSYSWho, r.SYSWho)
	inserttsql.Set(csSYSHost, r.SYSHost)
	inserttsql.Set(csSYSUpdated, r.SYSUpdated)
	inserttsql.Set(csSource, r.Source)

	log.Println("DELETE:", deletesql, core.ApplicationDB)

	sql, args := sqlexpr.Build(inserttsql)

	fmt.Printf("sql: %v\n", sql)
	fmt.Printf("args: %v\n", args)

	_, err2 := core.ApplicationDB.Exec(deletesql)
	if err2 != nil {
		log.Panic(err2)
	}
	//log.Println(fred2, err2)
	_, err := core.ApplicationDB.Exec(sql, args...)
	if err != nil {
		log.Panic(err)
	}
	//log.Println(fred, err)
}

func DataCache_DeleteItemByUD(id string) {
	deleteCacheStore(id)
}

func deleteCacheStore(id string) {
	//fmt.Println(credentialStore)
	deletesql := fmt.Sprintf(dsCache.Delete, id)
	//log.Println("DELETE:", deletesql, core.ApplicationDB)
	_, err2 := core.ApplicationDB.Exec(deletesql)
	if err2 != nil {
		log.Panicln(err2.Error())
	}
	//log.Println(fred2, err2)
}

// fetchCacheStoreData read all employees
func fetchCacheStoreData(tsql string) (int, []dm.DataCache, dm.DataCache, error) {
	//log.Println(tsql)
	var appCacheStore dm.DataCache
	var appCacheStoreList []dm.DataCache

	rows, err := core.ApplicationDB.Query(tsql)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, appCacheStore, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(&appCacheStore.Id, &appCacheStore.Object, &appCacheStore.Field, &appCacheStore.Value, &appCacheStore.Expiry, &appCacheStore.SYSCreated, &appCacheStore.SYSWho, &appCacheStore.SYSHost, &appCacheStore.SYSUpdated, &appCacheStore.Source)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, appCacheStore, err
		}
		// no change below
		appCacheStoreList = append(appCacheStoreList, appCacheStore)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	//log.Println(count, appCacheStoreList, appCacheStore)
	return count, appCacheStoreList, appCacheStore, nil
}

func newCacheStoreID(r dm.DataCache) string {
	id := r.Object + core.IDSep + r.Field + core.IDSep + r.Value + core.IDSep + r.Source
	return id
}

//DataCache_Build builds a cache of data from a particular siena table & field.
func DataCache_Build(table string, field string, fieldName string, objectName string, tsql string, sourceProperties map[string]string, r *http.Request) {
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

		PutCacheDataSampleItem(storeObjectName, storeField, sqlDataItem.String, sourceProperties["database"], r)
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
