package dao

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/andreyvit/sqlexpr"
	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	das "github.com/mt1976/mwt-go-dev/das"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// Defines the Fields to Fetch from SQL
//CredentialsStoreSQL = "id, username, password, firstname, lastname, knownas, email, issued, expiry, role, brand, _created, _who, _host, _updated"

//var sqlCredentialsStoreId, sqlCredentialsStoreUsername, sqlCredentialsStorePassword, sqlCredentialsStoreFirstname, sqlCredentialsStoreLastname, sqlCredentialsStoreKnownas, sqlCredentialsStoreEmail, sqlCredentialsStoreIssued, sqlCredentialsStoreExpiry, sqlCredentialsStoreRole, sqlCredentialsStoreBrand, sqlCredentialsStoreSYSCreated, sqlCredentialsStoreSYSWho, sqlCredentialsStoreSYSHost, sqlCredentialsStoreSYSUpdated sql.NullString
var dsCredentials dm.DataStoreMessages

//var appCredentialsStoreSQLINSERT = "INSERT INTO %s.credentialsStore(%s) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s','%s');"
//var appCredentialsStoreSQLDELETE = "DELETE FROM %s.credentialsStore WHERE id='%s';"
//var appCredentialsStoreSQLSELECT = "SELECT %s FROM %s.credentialsView;"
//var appCredentialsStoreSQLGET = "SELECT %s FROM %s.credentialsView WHERE id='%s';"
//var appCredentialsStoreSQLGETUSER = "SELECT %s FROM %s.credentialsView WHERE username='%s';"

func init() {
	//FullTableName := core.ApplicationPropertiesDB["schema"] + "." + "credentialsStore"
	// dsCredentials = dm.DataStoreMessages{
	// 	Table: FullTableName,
	// 	//		Columns:   "id, username, password, firstname, lastname, knownas, email, issued, expiry, role, brand, _created, _who, _host, _updated",
	// 	Insert:    "INSERT INTO " + FullTableName + "(%s) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s','%s');",
	// 	Delete:    "DELETE FROM " + FullTableName + " WHERE id='%s';",
	// 	Select:    "SELECT %s FROM " + FullTableName + ";",
	// 	Get:       "SELECT %s FROM " + FullTableName + " WHERE id='%s';",
	// 	GetAlt:    "SELECT %s FROM " + FullTableName + " WHERE username='%s';",
	// 	DeleteAlt: "",
	// }
}

//appCacheStoreItem is cheese
const (
	//Action     string
	credTable      = sqlexpr.Table("credentialsStore")
	credId         = sqlexpr.Column("id")
	credUsername   = sqlexpr.Column("username")
	credPassword   = sqlexpr.Column("password")
	credFirstname  = sqlexpr.Column("firstname")
	credLastname   = sqlexpr.Column("lastname")
	credKnownas    = sqlexpr.Column("knownas")
	credEmail      = sqlexpr.Column("email")
	credIssued     = sqlexpr.Column("issued")
	credExpiry     = sqlexpr.Column("expiry")
	credRole       = sqlexpr.Column("role")
	credBrand      = sqlexpr.Column("brand")
	credSYSCreated = sqlexpr.Column("_created")
	credSYSWho     = sqlexpr.Column("_who")
	credSYSHost    = sqlexpr.Column("_host")
	credSYSUpdated = sqlexpr.Column("_updated")
)

// getCredentialsStoreList read all employees
func Credentials_GetList() (int, []dm.AppCredentialsStoreItem, error) {
	//tsql := fmt.Sprintf(dsCredentials.Select, sqlstruct.Columns(appCredentialsStoreItem{}))

	//selectsql := buildStatement()
	//log.Printf("credBrand: %q", credBrand)
	//log.Printf("credSYSWho: %q", credSYSWho)
	//log.Println("SELECT:", selectsql, core.ApplicationDB)

	//	tsql, args := sqlexpr.Build(selectsql)

	//Credentials_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Credentials_SQLTable

	tsql := "SELECT * FROM " + tableName()
	// + " WHERE " + dm.Country_Code + " = '" + id + "'"

	count, appCredentialsStoreList, _, _ := fetchCredentialsStoreData(tsql)
	return count, appCredentialsStoreList, nil
}

// getCredentialsStoreList read all employees
func Credentials_GetByID(id string) (int, dm.AppCredentialsStoreItem, error) {
	//tsql := fmt.Sprintf(dsCredentials.Get, sqlstruct.Columns(appCredentialsStoreItem{}), id)

	//selectsql := buildStatement()
	//selectsql.AddWhere(sqlexpr.Eq(credId, id))
	//fmt.Printf("credBrand: %v\n", credBrand)

	//Credentials_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Credentials_SQLTable

	tsql := "SELECT * FROM " + tableName()
	tsql = tsql + " WHERE " + dm.Credentials_ID + " = '" + id + "'"

	// + " WHERE " + dm.Country_Code + " = '" + id + "'"
	//log.Println("SELECT:", selectsql, core.ApplicationDB)

	_, _, appCredentialsStoreItem, _ := fetchCredentialsStoreData(tsql)
	return 1, appCredentialsStoreItem, nil
}

// getCredentialsStoreList read all employees
func Credentials_GetByUserName(userName string) (int, dm.AppCredentialsStoreItem, error) {
	//tsql := fmt.Sprintf(dsCredentials.GetAlt, sqlstruct.Columns(appCredentialsStoreItem{}), userName)

	//Credentials_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Credentials_SQLTable

	tsql := "SELECT * FROM " + tableName()
	tsql = tsql + " WHERE " + dm.Credentials_Username + " = '" + userName + "'"

	_, _, appCredentialsStoreItem, _ := fetchCredentialsStoreData(tsql)
	return 1, appCredentialsStoreItem, nil
}

func Credentials_Store(r dm.AppCredentialsStoreItem, req *http.Request) {
	//fmt.Println(credentialStore)

	createDate := time.Now().Format(core.DATETIMEFORMATUSER)
	if len(r.SYSCreated) == 0 {
		r.SYSCreated = createDate
	}

	//currentUserID, _ := user.Current()
	//userID := currentUserID.Name
	host, _ := os.Hostname()

	if len(r.Id) == 0 {
		r.Id = newCredentialsStoreID()
		r.SYSCreated = createDate
		r.SYSWho = core.GetUserName(req)
		r.SYSHost = host
		r.Issued = createDate
		//expiryDate := time.Now()
		//life, _ := strconv.Atoi(core.ApplicationProperties["credentialslife"])
		//expiryDate = expiryDate.AddDate(0, 0, life)
		r.Expiry = ""
		r.Password = core.ApplicationProperties["defaultpassword"]
	}

	r.SYSUpdated = createDate

	fmt.Println("RECORD", r)
	//fmt.Printf("%s\n", sqlstruct.Columns(DataStoreSQL{}))

	//deletesql := fmt.Sprintf(dsCredentials.Delete, r.Id)

	//inserttsql := fmt.Sprintf(dsCredentials.Insert, dsCredentials.Columns, r.Id, r.Username, r.Password, r.Firstname, r.Lastname, r.Knownas, r.Email, r.Issued, r.Expiry, r.Role, r.Brand, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated)

	//log.Println("DELETE:", deletesql, core.ApplicationDB)
	//log.Println("INSERT:", inserttsql, core.ApplicationDB)

	//tableName :=

	// inserttsql := sqlexpr.Insert{Table: sqlexpr.Table(core.ApplicationPropertiesDB["schema"] + dsCredentials.Table)}
	// inserttsql.Set(credId, r.Id)

	// inserttsql.Set(credUsername, r.Username)
	// inserttsql.Set(credPassword, r.Password)
	// inserttsql.Set(credFirstname, r.Firstname)
	// inserttsql.Set(credLastname, r.Lastname)
	// inserttsql.Set(credKnownas, r.Knownas)
	// inserttsql.Set(credEmail, r.Email)
	// inserttsql.Set(credIssued, r.Issued)
	// inserttsql.Set(credExpiry, r.Expiry)
	// inserttsql.Set(credRole, r.Role)
	// inserttsql.Set(credBrand, r.Brand)
	// inserttsql.Set(credSYSCreated, r.SYSCreated)
	// inserttsql.Set(credSYSWho, r.SYSWho)
	// inserttsql.Set(credSYSHost, r.SYSHost)
	// inserttsql.Set(credSYSUpdated, r.SYSCreated)

	//log.Println("DELETE:", deletesql, core.ApplicationDB)

	// sql, args := sqlexpr.Build(inserttsql)

	// fmt.Printf("sql: %v\n", sql)
	// fmt.Printf("args: %v\n", args)

	tsql := "INSERT INTO " + tableName()
	tsql = tsql + " (" + dm.Credentials_ID + ", " + dm.Credentials_Username + ", " + dm.Credentials_Password + ", " + dm.Credentials_Firstname + ", " + dm.Credentials_Lastname + ", " + dm.Credentials_Knownas + ", " + dm.Credentials_Email + ", " + dm.Credentials_Issued + ", " + dm.Credentials_Expiry + ", " + dm.Credentials_Role + ", " + dm.Credentials_Brand + ", " + dm.Credentials_SYSCreated + ", " + dm.Credentials_SYSWho + ", " + dm.Credentials_SYSHost + ", " + dm.Credentials_SYSUpdated + ")"
	tsql = tsql + " VALUES (" + sq(r.Id) + ", " + sq(r.Username) + ", " + sq(r.Password) + ", " + sq(r.Firstname) + ", " + sq(r.Lastname) + ", " + sq(r.Knownas) + ", " + sq(r.Email) + ", " + sq(r.Issued) + ", " + sq(r.Expiry) + ", " + sq(r.Role) + ", " + sq(r.Brand) + ", " + sq(r.SYSCreated) + ", " + sq(r.SYSWho) + ", " + sq(r.SYSHost) + ", " + sq(r.SYSUpdated) + ")"

	DeleteCredentialsStore(r.Id)

	// _, err2 := core.ApplicationDB.Exec(deletesql)
	// if err2 != nil {
	// 	log.Panicf("%e", err2)
	// }

	das.Execute(tsql)

	// _, err := core.ApplicationDB.Exec(sql, args...)
	// if err != nil {
	// 	log.Panicf("%e", err)
	// }
}

func DeleteCredentialsStore(id string) {

	//log.Println("id:", id)
	Credentials_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Credentials_SQLTable

	tsql := "DELETE FROM " + Credentials_Table
	tsql = tsql + " WHERE " + dm.Credentials_ID + " = '" + id + "'"

	das.Execute(tsql)

}

// fetchCredentialsStoreData read all employees
func fetchCredentialsStoreData(tsql string) (int, []dm.AppCredentialsStoreItem, dm.AppCredentialsStoreItem, error) {

	var appCredentialsStore dm.AppCredentialsStoreItem
	var appCredentialsStoreList []dm.AppCredentialsStoreItem

	returnList, noitems, err := das.Query(core.ApplicationDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	//	spew.Dump(returnList)

	for i := 0; i < noitems; i++ {

		tc := returnList[i]

		appCredentialsStore.Id = tc[dm.Credentials_ID].(string)
		appCredentialsStore.Username = tc[dm.Credentials_Username].(string)
		appCredentialsStore.Password = tc[dm.Credentials_Password].(string)
		appCredentialsStore.Firstname = tc[dm.Credentials_Firstname].(string)
		appCredentialsStore.Lastname = tc[dm.Credentials_Lastname].(string)
		appCredentialsStore.Knownas = tc[dm.Credentials_Knownas].(string)
		appCredentialsStore.Email = tc[dm.Credentials_Email].(string)
		appCredentialsStore.Issued = tc[dm.Credentials_Issued].(string)
		appCredentialsStore.Expiry = tc[dm.Credentials_Expiry].(string)
		appCredentialsStore.Role = tc[dm.Credentials_Role].(string)
		appCredentialsStore.Brand = tc[dm.Credentials_Brand].(string)
		appCredentialsStore.SYSCreated = tc[dm.Credentials_SYSCreated].(string)
		appCredentialsStore.SYSWho = tc[dm.Credentials_SYSWho].(string)
		appCredentialsStore.SYSHost = tc[dm.Credentials_SYSHost].(string)
		appCredentialsStore.SYSUpdated = tc[dm.Credentials_SYSUpdated].(string)

		appCredentialsStoreList = append(appCredentialsStoreList, appCredentialsStore)

	}

	return noitems, appCredentialsStoreList, appCredentialsStore, nil
}

func newCredentialsStoreID() string {
	id := uuid.New().String()
	return id
}

func getExpiryDate() string {
	expiryDate := time.Now()
	life, _ := strconv.Atoi(core.ApplicationProperties["credentialslife"])
	expiryDate = expiryDate.AddDate(0, 0, life)
	return expiryDate.Format(core.DATETIMEFORMATUSER)
}

func tableName() string {
	return core.ApplicationPropertiesDB["schema"] + "." + dm.Credentials_SQLTable
}
