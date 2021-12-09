package adaptor

import (
	"encoding/xml"
	"fmt"

	"github.com/google/uuid"
	core "github.com/mt1976/mwt-go-dev/core"
)

const (
	//	StaticImport_CreateAction = "create"
	StaticImport_UpdateAction = "IMPORT"
	StaticImport_DeleteAction = "DELETE"
	StaticImport_XMLHeader    = "<?xml version=\"1.0\" encoding=\"iso-8859-15\"?>"
)

//StaticImport defines objects that are cheese
type StaticImport struct {
	XMLName     xml.Name `xml:"TRANSACTIONS"`
	Text        string   `xml:",chardata"`
	TRANSACTION struct {
		Text  string `xml:",chardata"`
		Type  string `xml:"type,attr"`
		TABLE struct {
			Text      string `xml:",chardata"`
			Name      string `xml:"name,attr"`
			Classname string `xml:"classname,attr"`
			RECORD    []struct {
				Text     string `xml:",chardata"`
				KEYFIELD []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"KEYFIELD"`
				FIELD []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"FIELD"`
			} `xml:"RECORD"`
		} `xml:"TABLE"`
	} `xml:"TRANSACTION"`
}

//StaticImport_Field defined a Siena XML Static Importer Field Tag
type StaticImport_Field struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr"`
}

//StaticImport_KeyField defined a Siena XML Static Importer KeyField Tag
type StaticImport_KeyField struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr"`
}

//StaticImport_Record defined a Siena XML Static Importer Record
type StaticImport_Record struct {
	Text     string `xml:",chardata"`
	KEYFIELD []StaticImport_KeyField
	FIELD    []StaticImport_Field
}

//StaticImport_Table defined a Siena XML Static Importer Table Container
type StaticImport_Table struct {
	Text      string `xml:",chardata"`
	Name      string `xml:"name,attr"`
	Classname string `xml:"classname,attr"`
	RECORD    []StaticImport_Record
}

//StaticImport_Transaction defined a Siena XML Static Importer Transaction Container
type StaticImport_Transaction struct {
	Text  string             `xml:",chardata"`
	Type  string             `xml:"type,attr"`
	TABLE StaticImport_Table `xml:"TABLE"`
}

//StaticImport_XML defined a Siena XML Static Importer Content
type StaticImport_XML struct {
	XMLName      xml.Name                 `xml:"TRANSACTIONS"`
	TRANSACTIONS StaticImport_Transaction `xml:"TRANSACTION"`
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

func StaticImport_Create(action string, table string, objName string, sienaKeyFields []StaticImport_KeyField, sienaFields []StaticImport_Field) []byte {

	sienaRecords := StaticImport_AddRecord(sienaKeyFields, sienaFields)
	sienaTable := StaticImport_NewTable(table, objName, sienaRecords)
	sienaTransaction := StaticImport_NewTransaction(action, sienaTable)
	StaticImport_XMLContent := StaticImport_AddTransaction(sienaTransaction)

	XMLmessage := StaticImport_GenXML(StaticImport_XMLContent)
	fmt.Printf("XMLmessage: %v\n", XMLmessage)
	return XMLmessage
}

func StaticImport_Dispatch(XMLmessage []byte) error {
	var err error
	fileName := uuid.New().String() + ".xml"
	delivertopath := core.SienaProperties["static_out"]
	ok := core.WriteDataFileAbsolute(fileName, delivertopath, string(XMLmessage))
	if ok == -1 {
		err = fmt.Errorf("Error writing file %s", fileName)
	}
	return err
}
