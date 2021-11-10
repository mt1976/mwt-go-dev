package adaptor

import (
	"encoding/xml"
	"io/ioutil"
	"os"

	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

func Centre_XMLExport(item dm.Centre) {

	core.LOG_message("Centre_XMLExport", "Doing it!")

	// Create XML File
	item.Action = "UPDATE"

	//fmt.Println("ITEM", item)
	// DEFINE THE XML FIELDS/KEYFIELDS HERE
	var sFldCode StaticImport_KeyField
	var sFldName StaticImport_Field
	var sFldCountry StaticImport_Field

	// POPULATE THE XML FIELDS
	sFldCode.Name = "shortName"
	sFldCode.Text = item.Code

	sFldName.Name = "fullName"
	sFldName.Text = item.Name

	sFldCountry.Name = "country"
	sFldCountry.Text = item.Country

	// IGNORE
	var sienaKeyFields []StaticImport_KeyField
	var sienaFields []StaticImport_Field
	// ADD THE FIELDS TO THE LISTS ABOVE
	sienaKeyFields = append(sienaKeyFields, sFldCode)
	sienaFields = append(sienaFields, sFldName)
	sienaFields = append(sienaFields, sFldCountry)
	// IGNORE
	sienaRecord := &StaticImport_Record{KEYFIELD: sienaKeyFields, FIELD: sienaFields}
	var sienaRecords []StaticImport_Record
	sienaRecords = append(sienaRecords, *sienaRecord)

	var sienaTable StaticImport_Table
	sienaTable.Name = "Centre"
	sienaTable.Classname = "com.eurobase.siena.data.centres.Centre"
	sienaTable.RECORD = sienaRecords

	var sienaTransaction StaticImport_Transaction
	sienaTransaction.Type = "IMPORT"
	sienaTransaction.TABLE = sienaTable

	var StaticImport_XMLContent StaticImport_XML
	StaticImport_XMLContent.TRANSACTIONS = sienaTransaction

	preparedXML, _ := xml.Marshal(StaticImport_XMLContent)

	pwd, _ := os.Getwd()
	fileName := pwd + staticImportPath() + "/" + getNewFileID() + ".xml"
	core.LOG_info("Centre_XMLExport", "Writing XML to file: "+fileName)
	err := ioutil.WriteFile(fileName, preparedXML, 0644)
	if err != nil {
		core.LOG_error(err.Error(), err)
	}
	core.LOG_message("Centre_XMLExport", "Done it!")

}
