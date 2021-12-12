package adaptor

import (
	"encoding/xml"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func Simulator_FundsChecker_Store(thisID string, balance string, resultCode string) {

	_, returnRecord, _ := Simulator_FundsChecker_GetByID(thisID)

	simFundsCheckItem := returnRecord.Request

	var simFundsCheckResponse dm.Simulator_FundsChecker_Response
	simFundsCheckResponse.NS1 = "http://dnb.lt/dnb-xst/MBT"
	simFundsCheckResponse.MBTHEADER.XREF = simFundsCheckItem.HEADER.XREF
	simFundsCheckResponse.MBTHEADER.SOURCE = "MB_TODO"
	simFundsCheckResponse.MBTHEADER.OPERATION = simFundsCheckItem.HEADER.OPERATION
	simFundsCheckResponse.MBTHEADER.DESTINATION = simFundsCheckItem.HEADER.SOURCE
	simFundsCheckResponse.MBTHEADER.MSGSTAT = "SUCCESS"
	simFundsCheckResponse.MBTBODY.FundsCheckResponse.BIC = simFundsCheckItem.BODY.FundsCheck.BIC
	simFundsCheckResponse.MBTBODY.FundsCheckResponse.ACC = simFundsCheckItem.BODY.FundsCheck.ACC
	simFundsCheckResponse.MBTBODY.FundsCheckResponse.CCY = simFundsCheckItem.BODY.FundsCheck.CCY
	simFundsCheckResponse.MBTBODY.FundsCheckResponse.AMOUNT = balance
	createDate := time.Now().Format(core.DATETIME)
	simFundsCheckResponse.MBTBODY.FundsCheckResponse.QUERYTIMESTAMP = createDate
	simFundsCheckResponse.MBTBODY.FundsCheckResponse.RESULTCODE = resultCode

	//log.Printf("%# v", simFundsCheckResponse)

	newMsg, err := xml.Marshal(simFundsCheckResponse)
	if err != nil {
		log.Println(string(newMsg), err)
	}

	newID := uuid.New().String()
	fileName := newID + ".xml"

	delivertopath := core.SienaProperties["funds_out"]
	deletefrompath := core.SienaProperties["funds_in"]

	log.Printf("Delivery      : %s -> %s", deletefrompath, delivertopath)
	core.WriteDataFileAbsolute(fileName, delivertopath, string(newMsg))

	resp := core.DeleteDataFileAbsolute(thisID, deletefrompath)
	if resp != 1 {
		//do nothing
	}
}

// getFundsCheckList read all employees
func Simulator_FundsChecker_DeleteByID(id string) error {
	requestPath := core.SienaProperties["funds_in"]
	//pwd, _ := os.Getwd()
	rtn := core.DeleteDataFileAbsolute(id, requestPath)
	if rtn != 1 {
		logs.Warning("Unable to Delete " + id)
	}
	return nil
}

func Simulator_FundsChecker_GetByID(id string) (int, dm.Simulator_FundsChecker_Item, error) {
	var simFundsCheckItem dm.Simulator_FundsChecker_Item
	requestPath := core.SienaProperties["funds_in"]

	dat, _, _ := core.ReadDataFileAbsolute(id, requestPath)
	simFundsCheckItem.Id = id
	simFundsCheckItem.Name = strings.ReplaceAll(id, ".xml", "")
	simFundsCheckItem.Source = requestPath
	simFundsCheckItem.Content = string(dat)

	var test dm.Simulator_FundsChecker_Request
	xml.Unmarshal(dat, &test)

	//log.Println(test)
	//log.Println(test.HEADER.XREF)
	//	log.Printf("%# v", test)

	simFundsCheckItem.Request = test

	return 1, simFundsCheckItem, nil
}

// getFundsCheckList read all employees
func Simulator_FundsChecker_GetList() (int, []dm.Simulator_FundsChecker_Item, error) {
	//tsql := fmt.Sprintf(simFundsCheckSQLSELECT, simFundsCheckSQL, core.ApplicationPropertiesDB["schema"])
	var simFundsCheckList []dm.Simulator_FundsChecker_Item
	requestPath := core.SienaProperties["funds_in"]

	files, _ := core.GetDirectoryContentAbsolute(requestPath)

	for _, k := range files {
		//fmt.Println("key:", k)
		var fci dm.Simulator_FundsChecker_Item
		fci.Id = k.Name()
		fci.Name = strings.ReplaceAll(k.Name(), ".xml", "")
		fci.Source = requestPath
		simFundsCheckList = append(simFundsCheckList, fci)
	}

	//count, simFundsCheckList, _, _ := fetchFundsCheckData("")
	return len(files), simFundsCheckList, nil
}
