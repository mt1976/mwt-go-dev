package main

import "encoding/xml"

//SIENA_IMPORT are cheese
type sienaStaticImportXML struct {
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

type sienaFIELD struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr"`
}

type sienaKEYFIELD struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr"`
}

type sienaRECORD struct {
	Text     string `xml:",chardata"`
	KEYFIELD []sienaKEYFIELD
	FIELD    []sienaFIELD
}

type sienaTABLE struct {
	Text      string `xml:",chardata"`
	Name      string `xml:"name,attr"`
	Classname string `xml:"classname,attr"`
	RECORD    []sienaRECORD
}

type sienaTRANSACTION struct {
	Text  string     `xml:",chardata"`
	Type  string     `xml:"type,attr"`
	TABLE sienaTABLE `xml:"TABLE"`
}

type sienaXML struct {
	XMLName      xml.Name         `xml:"TRANSACTIONS"`
	TRANSACTIONS sienaTRANSACTION `xml:"TRANSACTION"`
}
