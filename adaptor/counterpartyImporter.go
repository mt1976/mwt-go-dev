package adaptor

import "encoding/xml"

type counterpartyImportRECORD struct {
	XMLName       xml.Name `xml:"RECORDS"`
	Text          string   `xml:",chardata"`
	COUNTERPARTYS struct {
		Text         string `xml:",chardata"`
		COUNTERPARTY struct {
			Text                 string `xml:",chardata"`
			NAME                 string `xml:"NAME"`
			FULLNAME             string `xml:"FULLNAME"`
			COUNTRY              string `xml:"COUNTRY"`
			SECTOR               string `xml:"SECTOR"`
			ADDRESS              string `xml:"ADDRESS"`
			TELEPHONENUMBER      string `xml:"TELEPHONENUMBER"`
			FAXNUMBER            string `xml:"FAXNUMBER"`
			COUNTERPARTYIDSOURCE string `xml:"COUNTERPARTYIDSOURCE"`
			COUNTERPARTYID       string `xml:"COUNTERPARTYID"`
		} `xml:"COUNTERPARTY"`
	} `xml:"COUNTERPARTYS"`
	USERS struct {
		Text string `xml:",chardata"`
		USER struct {
			Text            string `xml:",chardata"`
			NAME            string `xml:"NAME"`
			SOURCEUSER      string `xml:"SOURCEUSER"`
			FULLNAME        string `xml:"FULLNAME"`
			EMAILADDRESS    string `xml:"EMAILADDRESS"`
			ENABLED         string `xml:"ENABLED"`
			LANGUAGE        string `xml:"LANGUAGE"`
			TELEPHONENUMBER string `xml:"TELEPHONENUMBER"`
			COUNTERPARTY    string `xml:"COUNTERPARTY"`
			CURRENCY        string `xml:"CURRENCY"`
			CURRENCYPAIR    string `xml:"CURRENCYPAIR"`
			MMESPMODE       string `xml:"MMESPMODE"`
			GROUPS          struct {
				Text  string   `xml:",chardata"`
				GROUP []string `xml:"GROUP"`
			} `xml:"GROUPS"`
			MANDATEDUSER struct {
				Text            string `xml:",chardata"`
				NAME            string `xml:"NAME"`
				TYPE            string `xml:"TYPE"`
				TELEPHONENUMBER string `xml:"TELEPHONENUMBER"`
				EMAILADDRESS    string `xml:"EMAILADDRESS"`
				NOTIFY          string `xml:"NOTIFY"`
				ACTIVE          string `xml:"ACTIVE"`
			} `xml:"MANDATEDUSER"`
		} `xml:"USER"`
	} `xml:"USERS"`
}
