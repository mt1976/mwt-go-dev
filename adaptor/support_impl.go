package adaptor

import (
	"encoding/xml"
	"fmt"
	"os"
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

func Simulator_SienaStaticImporter_DispatchXML(StaticImport_XMLContent StaticImport_XML) error {

	//preparedXML, _ := xml.Marshal(StaticImport_XMLContent)
	//fmt.Println("PreparedXML", string(preparedXML))

	fileID := uuid.New()
	fileName := staticImportPath() + "/" + fileID.String() + ".xml"
	//fmt.Println(fileName)

	xmlFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating XML file: ", err)
		return err
	}
	xmlFile.WriteString(core.SienaProperties["static_xml_encoding"] + "\n")
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(StaticImport_XMLContent)
	if err != nil {
		fmt.Println("Error encoding XML to file: ", err)
		return err
	}

	return err
}
