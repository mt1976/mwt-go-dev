package datamodel

import "encoding/xml"

// ----------------------------------------------------------------
// Manually generated  "/datamodel/Simulator_FundsChecker.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Simulator_FundsChecker (Simulator_FundsChecker)
// Endpoint 	        : Simulator_FundsChecker (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 22/11/2021 at 21:11:42
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

const (
	Simulator_FundsChecker_Title       = "Funds Checker"
	Simulator_FundsChecker_SQLTable    = "sienaSimulator_FundsChecker"
	Simulator_FundsChecker_SQLSearchID = "Code"
	Simulator_FundsChecker_QueryString = "Code"
	///
	/// Handler Defintions
	///
	Simulator_FundsChecker_TemplateList   = "Simulator_FundsChecker_List"
	Simulator_FundsChecker_TemplateView   = "Simulator_FundsChecker_View"
	Simulator_FundsChecker_TemplateAction = "Simulator_FundsChecker_Action"
	Simulator_FundsChecker_TemplateSubmit = "Simulator_FundsChecker_Submit"
	Simulator_FundsChecker_TemplateDelete = "Simulator_FundsChecker_Delete"

	///
	/// Handler Monitor Paths
	///
	Simulator_FundsChecker_PathList   = "/Simulator_FundsCheckerList/"
	Simulator_FundsChecker_PathView   = "/Simulator_FundsCheckerView/"
	Simulator_FundsChecker_PathAction = "/Simulator_FundsCheckerAction/"
	Simulator_FundsChecker_PathSubmit = "/Simulator_FundsCheckerSubmit/"
	//Simulator_FundsChecker_PathSave   = "/Simulator_FundsCheckerSave/"
	Simulator_FundsChecker_PathDelete = "/Simulator_FundsCheckerDelete/"
	///
	/// SQL Field Definitions
	///

	/// Definitions End
)

//Simulator_FundsChecker_Item is cheese
type Simulator_FundsChecker_Item struct {
	Id      string
	Name    string
	Source  string
	Content string
	Request Simulator_FundsChecker_Request
}

type Simulator_FundsChecker_Request struct {
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

type Simulator_FundsChecker_Response struct {
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
