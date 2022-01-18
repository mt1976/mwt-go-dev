package datamodel

import "encoding/xml"

// ----------------------------------------------------------------
// Manually generated  "/datamodel/Simulator_SienaFundsChecker.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Simulator_SienaFundsChecker (Simulator_SienaFundsChecker)
// Endpoint 	        : Simulator_SienaFundsChecker (Code)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : cryptoidCalcium [r0-21.11.01]
// Date & Time		    : 22/11/2021 at 21:11:42
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

const (
	Simulator_SienaFundsChecker_Title       = "Funds Checker"
	Simulator_SienaFundsChecker_SQLTable    = "sienaSimulator_SienaFundsChecker"
	Simulator_SienaFundsChecker_SQLSearchID = "Code"
	Simulator_SienaFundsChecker_QueryString = "FundsCheck"
	///
	/// Handler Defintions
	///
	Simulator_SienaFundsChecker_TemplateList   = "Simulator_SienaFundsChecker_List"
	Simulator_SienaFundsChecker_TemplateView   = "Simulator_SienaFundsChecker_View"
	Simulator_SienaFundsChecker_TemplateAction = "Simulator_SienaFundsChecker_Action"
	Simulator_SienaFundsChecker_TemplateSubmit = "Simulator_SienaFundsChecker_Submit"
	Simulator_SienaFundsChecker_TemplateDelete = "Simulator_SienaFundsChecker_Delete"

	///
	/// Handler Monitor Paths
	///
	Simulator_SienaFundsChecker_PathList   = "/Simulator_SienaFundsCheckerList/"
	Simulator_SienaFundsChecker_PathView   = "/Simulator_SienaFundsCheckerView/"
	Simulator_SienaFundsChecker_PathAction = "/Simulator_SienaFundsCheckerAction/"
	Simulator_SienaFundsChecker_PathSubmit = "/Simulator_SienaFundsCheckerSubmit/"
	//Simulator_SienaFundsChecker_PathSave   = "/Simulator_SienaFundsCheckerSave/"
	Simulator_SienaFundsChecker_PathDelete = "/Simulator_SienaFundsCheckerDelete/"
	///
	/// SQL Field Definitions
	///

	/// Definitions End
)

//Simulator_SienaFundsChecker_Item is cheese
type Simulator_SienaFundsChecker_Item struct {
	Id      string
	Name    string
	Source  string
	Content string
	Request Simulator_SienaFundsChecker_Request
}

type Simulator_SienaFundsChecker_Request struct {
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

type Simulator_SienaFundsChecker_Response struct {
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
