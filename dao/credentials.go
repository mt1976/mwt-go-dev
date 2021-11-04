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
	FullTableName := core.ApplicationPropertiesDB["schema"] + "." + "credentialsStore"
	dsCredentials = dm.DataStoreMessages{
		Table: FullTableName,
		//		Columns:   "id, username, password, firstname, lastname, knownas, email, issued, expiry, role, brand, _created, _who, _host, _updated",
		Insert:    "INSERT INTO " + FullTableName + "(%s) VALUES('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s','%s');",
		Delete:    "DELETE FROM " + FullTableName + " WHERE id='%s';",
		Select:    "SELECT %s FROM " + FullTableName + ";",
		Get:       "SELECT %s FROM " + FullTableName + " WHERE id='%s';",
		GetAlt:    "SELECT %s FROM " + FullTableName + " WHERE username='%s';",
		DeleteAlt: "",
	}
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
func GetCredentialsStoreList() (int, []dm.AppCredentialsStoreItem, error) {
	//tsql := fmt.Sprintf(dsCredentials.Select, sqlstruct.Columns(appCredentialsStoreItem{}))

	selectsql := buildStatement()

	//log.Println("SELECT:", selectsql, core.ApplicationDB)

	tsql, args := sqlexpr.Build(selectsql)

	count, appCredentialsStoreList, _, _ := fetchCredentialsStoreData(tsql, args)
	return count, appCredentialsStoreList, nil
}

// getCredentialsStoreList read all employees
func GetCredentialsStoreByID(id string) (int, dm.AppCredentialsStoreItem, error) {
	//tsql := fmt.Sprintf(dsCredentials.Get, sqlstruct.Columns(appCredentialsStoreItem{}), id)

	selectsql := buildStatement()
	selectsql.AddWhere(sqlexpr.Eq(credId, id))

	//log.Println("SELECT:", selectsql, core.ApplicationDB)

	tsql, args := sqlexpr.Build(selectsql)

	_, _, appCredentialsStoreItem, _ := fetchCredentialsStoreData(tsql, args)
	return 1, appCredentialsStoreItem, nil
}

// getCredentialsStoreList read all employees
func GetCredentialsStoreByUserName(userName string) (int, dm.AppCredentialsStoreItem, error) {
	//tsql := fmt.Sprintf(dsCredentials.GetAlt, sqlstruct.Columns(appCredentialsStoreItem{}), userName)

	selectsql := buildStatement()
	selectsql.AddWhere(sqlexpr.Eq(credUsername, userName))

	tsql, args := sqlexpr.Build(selectsql)

	_, _, appCredentialsStoreItem, _ := fetchCredentialsStoreData(tsql, args)
	return 1, appCredentialsStoreItem, nil
}

func PutCredentialsStore(r dm.AppCredentialsStoreItem, req *http.Request) {
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

	deletesql := fmt.Sprintf(dsCredentials.Delete, r.Id)

	//inserttsql := fmt.Sprintf(dsCredentials.Insert, dsCredentials.Columns, r.Id, r.Username, r.Password, r.Firstname, r.Lastname, r.Knownas, r.Email, r.Issued, r.Expiry, r.Role, r.Brand, r.SYSCreated, r.SYSWho, r.SYSHost, r.SYSUpdated)

	//log.Println("DELETE:", deletesql, core.ApplicationDB)
	//log.Println("INSERT:", inserttsql, core.ApplicationDB)

	//tableName :=

	inserttsql := sqlexpr.Insert{Table: sqlexpr.Table(core.ApplicationPropertiesDB["schema"] + dsCredentials.Table)}
	inserttsql.Set(credId, r.Id)

	inserttsql.Set(credUsername, r.Username)
	inserttsql.Set(credPassword, r.Password)
	inserttsql.Set(credFirstname, r.Firstname)
	inserttsql.Set(credLastname, r.Lastname)
	inserttsql.Set(credKnownas, r.Knownas)
	inserttsql.Set(credEmail, r.Email)
	inserttsql.Set(credIssued, r.Issued)
	inserttsql.Set(credExpiry, r.Expiry)
	inserttsql.Set(credRole, r.Role)
	inserttsql.Set(credBrand, r.Brand)
	inserttsql.Set(credSYSCreated, r.SYSCreated)
	inserttsql.Set(credSYSWho, r.SYSWho)
	inserttsql.Set(credSYSHost, r.SYSHost)
	inserttsql.Set(credSYSUpdated, r.SYSCreated)

	log.Println("DELETE:", deletesql, core.ApplicationDB)

	sql, args := sqlexpr.Build(inserttsql)

	fmt.Printf("sql: %v\n", sql)
	fmt.Printf("args: %v\n", args)

	_, err2 := core.ApplicationDB.Exec(deletesql)
	if err2 != nil {
		log.Panicf("%e", err2)
	}
	_, err := core.ApplicationDB.Exec(sql, args...)
	if err != nil {
		log.Panicf("%e", err)
	}
}

func DeleteCredentialsStore(id string) {
	//fmt.Println(credentialStore)

	deletesql := sqlexpr.Delete{Table: sqlexpr.Table(core.ApplicationPropertiesDB["schema"] + dsCredentials.Table)}
	deletesql.AddWhere(sqlexpr.Eq(credId, id))

	tsql, args := sqlexpr.Build(deletesql)
	fmt.Printf("tsql: %v\n", tsql)
	fmt.Printf("args: %v\n", args)
	//deletesql := fmt.Sprintf(dsCredentials.Delete, id)
	_, err := core.ApplicationDB.Exec(tsql, args...)
	if err != nil {
		log.Panicf("%e", err)
	}
}

func FetchCredentialsStoreData(tsql string, args []interface{}) (int, []dm.AppCredentialsStoreItem, dm.AppCredentialsStoreItem, error) {
	c, l, i, _ := fetchCredentialsStoreData(tsql, args)
	return c, l, i, nil
}

// fetchCredentialsStoreData read all employees
func fetchCredentialsStoreData(tsql string, args []interface{}) (int, []dm.AppCredentialsStoreItem, dm.AppCredentialsStoreItem, error) {

	fmt.Printf("tsql: %v\n", tsql)
	fmt.Printf("args: %v\n", args)

	var appCredentialsStore dm.AppCredentialsStoreItem
	var appCredentialsStoreList []dm.AppCredentialsStoreItem

	rows, err := core.ApplicationDB.Query(tsql, args...)
	//fmt.Println("back from dq Q")
	if err != nil {
		log.Println("Error reading rows: " + err.Error())
		return -1, nil, appCredentialsStore, err
	}
	//fmt.Println(rows)
	defer rows.Close()
	count := 0
	for rows.Next() {
		err := rows.Scan(
			&appCredentialsStore.Id,
			&appCredentialsStore.Username,
			&appCredentialsStore.Password,
			&appCredentialsStore.Firstname,
			&appCredentialsStore.Lastname,
			&appCredentialsStore.Knownas,
			&appCredentialsStore.Email,
			&appCredentialsStore.Issued,
			&appCredentialsStore.Expiry,
			&appCredentialsStore.Role,
			&appCredentialsStore.Brand,
			&appCredentialsStore.SYSCreated,
			&appCredentialsStore.SYSWho,
			&appCredentialsStore.SYSHost,
			&appCredentialsStore.SYSUpdated,
		)
		if err != nil {
			log.Println("Error reading rows: " + err.Error())
			return -1, nil, appCredentialsStore, err
		}

		fmt.Printf("appCredentialsStore: %v\n", appCredentialsStore)
		// no change below
		appCredentialsStoreList = append(appCredentialsStoreList, appCredentialsStore)
		//log.Printf("Code: %s, Name: %s, Shortcode: %s, eu_eea: %t\n", code, name, shortcode, eu_eea)
		count++
	}
	//log.Println(count, appCredentialsStoreList, appCredentialsStore)
	return count, appCredentialsStoreList, appCredentialsStore, nil
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

func buildStatement() sqlexpr.Select {
	selectsql := sqlexpr.Select{From: sqlexpr.Table(core.ApplicationPropertiesDB["schema"] + dsCredentials.Table)}
	selectsql.AddField(credId)
	selectsql.AddField(credUsername)
	selectsql.AddField(credPassword)
	selectsql.AddField(credFirstname)
	selectsql.AddField(credLastname)
	selectsql.AddField(credKnownas)
	selectsql.AddField(credEmail)
	selectsql.AddField(credIssued)
	selectsql.AddField(credExpiry)
	selectsql.AddField(credRole)
	selectsql.AddField(credBrand)
	selectsql.AddField(credSYSCreated)
	selectsql.AddField(credSYSWho)
	selectsql.AddField(credSYSHost)
	selectsql.AddField(credSYSUpdated)
	return selectsql
}
