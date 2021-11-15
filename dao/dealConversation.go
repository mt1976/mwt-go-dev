package dao

// ----------------------------------------------------------------
// Automatically generated  "/dao/dealconversation.go"
// ----------------------------------------------------------------
// Package            : dao
// Object 			  : DealConversation
// Endpoint Root 	  : DealConversation
// Search QueryString : ID
// ----------------------------------------------------------------
// Template Generator : RussetAlbatross [r0-21.11.01]
// ----------------------------------------------------------------
// Date & Time		  : 15/11/2021 at 19:03:10
// Who & Where		  : matttownsend on silicon.local
// ----------------------------------------------------------------
import (
	"fmt"
	"log"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
	das "github.com/mt1976/mwt-go-dev/das"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

// DealConversation_GetList() returns a list of all DealConversation records
func DealConversation_GetList() (int, []dm.DealConversation, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.DealConversation_SQLTable)
	count, dealconversationList, _, _ := dealconversation_Fetch(tsql)
	return count, dealconversationList, nil
}

// DealConversation_GetByID() returns a single DealConversation record
func DealConversation_GetByID(id string) (int, dm.DealConversation, error) {

	tsql := "SELECT * FROM " + get_TableName(core.SienaPropertiesDB["schema"], dm.DealConversation_SQLTable)
	tsql = tsql + " WHERE " + dm.DealConversation_SQLSearchID + "='" + id + "'"

	_, _, dealconversationItem, _ := dealconversation_Fetch(tsql)
	return 1, dealconversationItem, nil
}

// DealConversation_DeleteByID() deletes a single DealConversation record
func DealConversation_Delete(id string) {

	object_Table := core.ApplicationPropertiesDB["schema"] + "." + dm.DealConversation_SQLTable

	tsql := "DELETE FROM " + object_Table
	tsql = tsql + " WHERE " + dm.DealConversation_SQLSearchID + " = '" + id + "'"

	das.Execute(tsql)
}

// DealConversation_Store() saves/stores a DealConversation record to the database
func DealConversation_Store(r dm.DealConversation) error {

	// TODO Implement Store Function for DealConversation
	fmt.Println(r)

	//TODO Deal with the if its Application or null add this bit, otherwise dont.
	//fmt.Println(credentialStore)
	// 	createDate := time.Now().Format(core.DATETIMEFORMATUSER)
	// 	if len(r.SYSCreated) == 0 {
	// 		r.SYSCreated = createDate
	// 	}

	// 	currentUserID, _ := user.Current()
	// 	userID := currentUserID.Name
	// 	host, _ := os.Hostname()

	// 	if len(r.AppInternalID) == 0 {
	// 		r.AppInternalID = newdealconversationStoreID()
	// 		r.SYSCreated = createDate
	// 		r.SYSWho = userID
	// 		r.SYSHost = host
	// 	}

	// 	r.SYSUpdated = createDate
	// //TODO Deal with the if its Application or null add this bit, otherwise dont.

	ts := SQLData{}

	ts = addData(ts, dm.DealConversation_SienaReference, r.SienaReference)
	ts = addData(ts, dm.DealConversation_Status, r.Status)
	ts = addData(ts, dm.DealConversation_MessageType, r.MessageType)
	ts = addData(ts, dm.DealConversation_ContractNumber, r.ContractNumber)
	ts = addData(ts, dm.DealConversation_AckReference, r.AckReference)
	ts = addData(ts, dm.DealConversation_NewTX, r.NewTX)
	ts = addData(ts, dm.DealConversation_LegNo, r.LegNo)
	ts = addData(ts, dm.DealConversation_Summary, r.Summary)
	ts = addData(ts, dm.DealConversation_BusinessDate, r.BusinessDate)
	ts = addData(ts, dm.DealConversation_TXNo, r.TXNo)
	ts = addData(ts, dm.DealConversation_ExternalSystem, r.ExternalSystem)
	ts = addData(ts, dm.DealConversation_MessageLogReference, r.MessageLogReference)

	tsql := "INSERT INTO " + get_TableName(core.SienaPropertiesDB["schema"], dm.DealConversation_SQLTable)
	tsql = tsql + " (" + fields(ts) + ")"
	tsql = tsql + " VALUES (" + values(ts) + ")"

	DealConversation_Delete(r.MessageLogReference)

	das.Execute(tsql)

	return nil
}

// dealconversation_Fetch read all employees
func dealconversation_Fetch(tsql string) (int, []dm.DealConversation, dm.DealConversation, error) {

	var recItem dm.DealConversation
	var recList []dm.DealConversation

	returnList, noitems, err := das.Query(core.SienaDB, tsql)
	if err != nil {
		log.Fatal(err.Error())
	}

	for i := 0; i < noitems; i++ {

		rec := returnList[i]
		// Automatically generated 15/11/2021 by matttownsend on silicon.local - START
		recItem.AppInternalID = get_String(rec, dm.DealConversation_MessageLogReference, "")
		recItem.SienaReference = get_String(rec, dm.DealConversation_SienaReference, "")
		recItem.Status = get_String(rec, dm.DealConversation_Status, "")
		recItem.MessageType = get_String(rec, dm.DealConversation_MessageType, "")
		recItem.ContractNumber = get_String(rec, dm.DealConversation_ContractNumber, "")
		recItem.AckReference = get_String(rec, dm.DealConversation_AckReference, "")
		recItem.NewTX = get_Bool(rec, dm.DealConversation_NewTX, "True")
		recItem.LegNo = get_Int(rec, dm.DealConversation_LegNo, "0")
		recItem.Summary = get_String(rec, dm.DealConversation_Summary, "")
		recItem.BusinessDate = get_Time(rec, dm.DealConversation_BusinessDate, "")
		recItem.TXNo = get_Int(rec, dm.DealConversation_TXNo, "0")
		recItem.ExternalSystem = get_String(rec, dm.DealConversation_ExternalSystem, "")
		recItem.MessageLogReference = get_String(rec, dm.DealConversation_MessageLogReference, "")
		// Automatically generated 15/11/2021 by matttownsend on silicon.local - END
		//Post Import Actions

		//Add to the list
		recList = append(recList, recItem)
	}
	return noitems, recList, recItem, nil
}

func newdealconversationStoreID() string {
	id := uuid.New().String()
	return id
}

// ----------------------------------------------------------------
// ADD Aditional Functions below this line
// ----------------------------------------------------------------
