package dao

import (
	"github.com/google/uuid"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

func mandate_NewIDImpl(r dm.Mandate) string {
	return uuid.New().String()
}
