package dao

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	das "github.com/mt1976/mwt-go-dev/das"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

func init() {

}

// getCredentialsStoreList read all employees
func Credentials_GetList() (int, []dm.Credentials, error) {

	tsql := "SELECT * FROM " + tableName()

	count, appCredentialsStoreList, _, _ := fetchCredentialsStoreData(tsql)
	return count, appCredentialsStoreList, nil
}

// getCredentialsStoreList read all employees
func Credentials_GetByID(id string) (int, dm.Credentials, error) {

	tsql := "SELECT * FROM " + tableName()
	tsql = tsql + " WHERE " + dm.Credentials_ID + " = '" + id + "'"

	_, _, appCredentialsStoreItem, _ := fetchCredentialsStoreData(tsql)
	return 1, appCredentialsStoreItem, nil
}

// getCredentialsStoreList read all employees
func Credentials_GetByUserName(userName string) (int, dm.Credentials, error) {

	tsql := "SELECT * FROM " + tableName()
	tsql = tsql + " WHERE " + dm.Credentials_Username + " = '" + userName + "'"

	_, _, appCredentialsStoreItem, _ := fetchCredentialsStoreData(tsql)
	return 1, appCredentialsStoreItem, nil
}

func Credentials_Store(r dm.Credentials, req *http.Request) {

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

	tsql := "INSERT INTO " + tableName()
	tsql = tsql + " (" + dm.Credentials_ID + ", " + dm.Credentials_Username + ", " + dm.Credentials_Password + ", " + dm.Credentials_Firstname + ", " + dm.Credentials_Lastname + ", " + dm.Credentials_Knownas + ", " + dm.Credentials_Email + ", " + dm.Credentials_Issued + ", " + dm.Credentials_Expiry + ", " + dm.Credentials_Role + ", " + dm.Credentials_Brand + ", " + dm.Credentials_SYSCreated + ", " + dm.Credentials_SYSWho + ", " + dm.Credentials_SYSHost + ", " + dm.Credentials_SYSUpdated + ")"
	tsql = tsql + " VALUES (" + sq(r.Id) + ", " + sq(r.Username) + ", " + sq(r.Password) + ", " + sq(r.Firstname) + ", " + sq(r.Lastname) + ", " + sq(r.Knownas) + ", " + sq(r.Email) + ", " + sq(r.Issued) + ", " + sq(r.Expiry) + ", " + sq(r.Role) + ", " + sq(r.Brand) + ", " + sq(r.SYSCreated) + ", " + sq(r.SYSWho) + ", " + sq(r.SYSHost) + ", " + sq(r.SYSUpdated) + ")"

	DeleteCredentialsStore(r.Id)

	das.Execute(tsql)

}

func DeleteCredentialsStore(id string) {

	//log.Println("id:", id)
	Credentials_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.Credentials_SQLTable

	tsql := "DELETE FROM " + Credentials_Table
	tsql = tsql + " WHERE " + dm.Credentials_ID + " = '" + id + "'"

	das.Execute(tsql)

}

// fetchCredentialsStoreData read all employees
func fetchCredentialsStoreData(tsql string) (int, []dm.Credentials, dm.Credentials, error) {

	var appCredentialsStore dm.Credentials
	var appCredentialsStoreList []dm.Credentials

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
