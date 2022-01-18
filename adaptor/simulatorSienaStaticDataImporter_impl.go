package adaptor

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/mt1976/mwt-go-dev/core"
	"github.com/mt1976/mwt-go-dev/logs"
)

func StaticImportProcessResponse(filename string) error {
	//	logs.Success("StaticImport_Dispatch")
	logs.Success("StaticImport_ProcessResponse:" + filename)
	var err error
	return err
}

func Simulator_SienaStaticDataImporter_ProcessResponse(filename string) error {
	//	logs.Success("StaticImport_Dispatch")
	logs.Success("SienaDealImporter_ProcessResponse:" + filename)

	// tokenise the filename, get last element
	tokens := strings.Split(filename, "/")
	last := tokens[len(tokens)-1]

	//	uri := dm.Simulator_SienaFundsChecker_PathAction + "/?" + dm.Simulator_SienaFundsChecker_QueryString + "=" + last

	core.Notification_Normal("New Siena Deal Importer Response", last)
	var err error
	return err
}

func StaticImport_AddKeyField(skf []StaticImport_KeyField, fieldName string, fieldValue string) []StaticImport_KeyField {
	skf = append(skf, StaticImport_KeyField{
		Name: fieldName,
		Text: fieldValue,
	})
	return skf
}

func StaticImport_AddField(sf []StaticImport_Field, fieldName string, fieldValue string) []StaticImport_Field {
	sf = append(sf, StaticImport_Field{
		Name: fieldName,
		Text: fieldValue,
	})
	return sf
}

func StaticImport_NewTable(tableName string, className string, records []StaticImport_Record) StaticImport_Table {
	table := StaticImport_Table{
		Name:      tableName,
		Classname: className,
		RECORD:    records,
	}
	return table
}

func StaticImport_AddRecord(keys []StaticImport_KeyField, fields []StaticImport_Field) []StaticImport_Record {
	var record StaticImport_Record
	var records []StaticImport_Record
	record.KEYFIELD = keys
	record.FIELD = fields
	records = append(records, record)
	return records
}

func StaticImport_NewTransaction(txnType string, table StaticImport_Table) StaticImport_Transaction {
	var transaction StaticImport_Transaction
	transaction.Type = txnType
	transaction.TABLE = table
	return transaction
}

func StaticImport_AddTransaction(txn StaticImport_Transaction) StaticImport_XML {
	var StaticImport_XMLContent StaticImport_XML
	StaticImport_XMLContent.TRANSACTIONS = txn
	return StaticImport_XMLContent
}

func StaticImport_GenXML(record StaticImport_XML) []byte {
	preparedXML, _ := xml.Marshal(record)
	preparedXML = []byte(StaticImport_XMLHeader + string(preparedXML))
	//fmt.Println("PreparedXML", string(preparedXML))
	return preparedXML
}

func StaticImport_Create(action string, table string, sienaKeyFields []StaticImport_KeyField, sienaFields []StaticImport_Field) []byte {

	sienaRecords := StaticImport_AddRecord(sienaKeyFields, sienaFields)
	sienaTable := StaticImport_NewTable(table, core.GetSienaClassName(table), sienaRecords)
	sienaTransaction := StaticImport_NewTransaction(action, sienaTable)
	StaticImport_XMLContent := StaticImport_AddTransaction(sienaTransaction)

	XMLmessage := StaticImport_GenXML(StaticImport_XMLContent)
	//fmt.Printf("XMLmessage: %v\n", XMLmessage)
	return XMLmessage
}

func StaticImport_DispatchXML(XMLmessage []byte, msgClass string) error {
	//	logs.Success("StaticImport_Dispatch")
	var err error
	fileID := uuid.New()
	fileName := fileID.String() + ".xml"
	delivertopath := core.SienaProperties["static_out"]
	ok := core.WriteDataFileAbsolute(fileName, delivertopath, string(XMLmessage))
	if ok == -1 {
		err = fmt.Errorf("error writing file %s", fileName)
	}
	//logs.Information("Send Message", fileName)
	_ = ExternalMessage_Sent(fileID.String(), Message_FormatXML, msgClass, delivertopath, XMLmessage, fileName, 10, Message_TimeoutAction_Notify)
	//logs.Information("Return from Send Message", fileName)

	//logs.Information("StaticImport_DispatchXML:", err.Error())
	return err
}
