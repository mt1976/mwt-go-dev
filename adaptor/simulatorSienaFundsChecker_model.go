package adaptor

import "encoding/xml"

type FundsCheckRequest struct {
	XMLName        xml.Name `xml:"REQ"`
	Text           string   `xml:",chardata"`
	NS1            string   `xml:"NS1,attr"`
	Xs             string   `xml:"xs,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	HEADER         struct {
		Text        string `xml:",chardata"`
		XREF        string `xml:"XREF"`
		SOURCE      string `xml:"SOURCE"`
		OPERATION   string `xml:"OPERATION"`
		DESTINATION string `xml:"DESTINATION"`
	} `xml:"HEADER"`
	BODY struct {
		Text       string `xml:",chardata"`
		FundsCheck struct {
			Text    string `xml:",chardata"`
			BIC     string `xml:"BIC"`
			ACC     string `xml:"ACC"`
			CCY     string `xml:"CCY"`
			AMOUNT  string `xml:"AMOUNT"`
			VALDATE string `xml:"VALDATE"`
		} `xml:"FundsCheck"`
	} `xml:"BODY"`
}

type FundsCheckReponse struct {
	XMLName   xml.Name `xml:"NS1:MBT_RES"`
	Text      string   `xml:",chardata"`
	NS1       string   `xml:"xmlns:NS1,attr"`
	MBTHEADER struct {
		Text        string `xml:",chardata"`
		XREF        string `xml:"NS1:XREF"`
		SOURCE      string `xml:"NS1:SOURCE"`
		OPERATION   string `xml:"NS1:OPERATION"`
		DESTINATION string `xml:"NS1:DESTINATION"`
		MSGSTAT     string `xml:"NS1:MSGSTAT"`
	} `xml:"NS1:MBT_HEADER"`
	MBTBODY struct {
		Text               string `xml:",chardata"`
		FundsCheckResponse struct {
			Text           string `xml:",chardata"`
			BIC            string `xml:"NS1:BIC"`
			ACC            string `xml:"NS1:ACC"`
			CCY            string `xml:"NS1:CCY"`
			AMOUNT         string `xml:"NS1:AMOUNT"`
			QUERYTIMESTAMP string `xml:"NS1:QUERYTIMESTAMP"`
			RESULTCODE     string `xml:"NS1:RESULTCODE"`
		} `xml:"NS1:FundsCheckResponse"`
	} `xml:"NS1:MBT_BODY"`
}
