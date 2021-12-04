package dao

import (
	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

func credentials_NewIDImpl(r dm.Credentials) string {
	return uuid.New().String()
}

// Credentials_GetByUserName() returns a single Credentials record
func Credentials_GetByUserName(id string) (int, dm.Credentials, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Credentials_SQLTable)
	tsql = tsql + " WHERE " + dm.Credentials_Username + "='" + id + "'"

	_, _, credentialsItem, _ := credentials_Fetch(tsql)
	return 1, credentialsItem, nil
}

// Credentials_GetByUUID() returns a single Credentials record
func Credentials_GetByUUID(id string) (int, dm.Credentials, error) {

	tsql := "SELECT * FROM " + get_TableName(core.ApplicationPropertiesDB["schema"], dm.Credentials_SQLTable)
	tsql = tsql + " WHERE " + dm.Credentials_Id + "='" + id + "'"

	_, _, credentialsItem, _ := credentials_Fetch(tsql)
	return 1, credentialsItem, nil
}
