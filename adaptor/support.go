package adaptor

import (
	"strings"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
)

func staticImportPath() string {
	return core.SienaProperties["static_in"]
}

func getNewFileID() string {
	nID := uuid.New().String()
	nID = core.RemoveSpecialChars(nID)
	nID = strings.ReplaceAll(nID, "-", "")
	return nID
}
