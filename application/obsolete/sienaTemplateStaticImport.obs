package application

import "encoding/xml"

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
