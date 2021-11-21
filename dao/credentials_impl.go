package dao

import (
	"github.com/google/uuid"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

func credentials_NewIDImpl(r dm.Credentials) string {
	return uuid.New().String()
}
