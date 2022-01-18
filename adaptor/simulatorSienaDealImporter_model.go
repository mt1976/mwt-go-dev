package adaptor

import "encoding/xml"

type spotCollections struct {
	XMLName    xml.Name `xml:"Collections"`
	Text       string   `xml:",chardata"`
	Collection struct {
		Text             string `xml:",chardata"`
		DealType         string `xml:"DealType"`
		MsgID            string `xml:"MsgID"`
		ExecType         string `xml:"ExecType"`
		TradeDate        string `xml:"TradeDate"`
		ExtRefNo         string `xml:"ExtRefNo"`
		ExternalCustomer string `xml:"ExternalCustomer"`
		Direction        string `xml:"Direction"`
		Rate             string `xml:"Rate"`
		DealtCurrency    string `xml:"DealtCurrency"`
		AgainstCurrency  string `xml:"AgainstCurrency"`
		DealtAmount      string `xml:"DealtAmount"`
		AgainstAmount    string `xml:"AgainstAmount"`
		Book             string `xml:"Book"`
		ValueDate        string `xml:"ValueDate"`
		Notes            string `xml:"Notes"`
		MandatedUser     string `xml:"MandatedUser"`
		LEI              string `xml:"LEI"`
		Broker           string `xml:"Broker"`
	} `xml:"Collection"`
}

type outrightCollections struct {
	XMLName    xml.Name `xml:"Collections"`
	Text       string   `xml:",chardata"`
	Collection struct {
		Text             string `xml:",chardata"`
		DealType         string `xml:"DealType"`
		MsgID            string `xml:"MsgID"`
		ExecType         string `xml:"ExecType"`
		TradeDate        string `xml:"TradeDate"`
		ExtRefNo         string `xml:"ExtRefNo"`
		ExternalCustomer string `xml:"ExternalCustomer"`
		Direction        string `xml:"Direction"`
		Rate             string `xml:"Rate"`
		DealtCurrency    string `xml:"DealtCurrency"`
		AgainstCurrency  string `xml:"AgainstCurrency"`
		DealtAmount      string `xml:"DealtAmount"`
		AgainstAmount    string `xml:"AgainstAmount"`
		Book             string `xml:"Book"`
		ValueDate        string `xml:"ValueDate"`
		Notes            string `xml:"Notes"`
		MandatedUser     string `xml:"MandatedUser"`
		LEI              string `xml:"LEI"`
		Broker           string `xml:"Broker"`
	} `xml:"Collection"`
}
type statementCollections struct {
	XMLName    xml.Name `xml:"Collections"`
	Text       string   `xml:",chardata"`
	Collection []struct {
		Text            string `xml:",chardata"`
		ParentExtRefNo  string `xml:"ParentExtRefNo"`
		DealType        string `xml:"DealType"`
		ExecType        string `xml:"ExecType"`
		ValueDate       string `xml:"ValueDate"`
		Adjustment      string `xml:"Adjustment"`
		Notes           string `xml:"Notes"`
		ExtRefNo        string `xml:"ExtRefNo"`
		StatementNotes  string `xml:"StatementNotes"`
		RedefaultValues string `xml:"RedefaultValues"`
	} `xml:"Collection"`
}
type accountCollection struct {
	XMLName             xml.Name `xml:"Collection"`
	Text                string   `xml:",chardata"`
	DealType            string   `xml:"DealType"`
	MsgID               string   `xml:"MsgID"`
	ExecType            string   `xml:"ExecType"`
	TradeDate           string   `xml:"TradeDate"`
	ExtRefNo            string   `xml:"ExtRefNo"`
	ExternalCustomer    string   `xml:"ExternalCustomer"`
	MandatedUser        string   `xml:"MandatedUser"`
	PaymentSystem       string   `xml:"PaymentSystem"`
	OriginUser          string   `xml:"OriginUser"`
	Broker              string   `xml:"Broker"`
	Direction           string   `xml:"Direction"`
	Rate                string   `xml:"Rate"`
	MarketInterestRate  string   `xml:"MarketInterestRate"`
	DealtCurrency       string   `xml:"DealtCurrency"`
	DealtAmount         string   `xml:"DealtAmount"`
	Book                string   `xml:"Book"`
	ValueDate           string   `xml:"ValueDate"`
	MaturityDate        string   `xml:"MaturityDate"`
	InterestFrequency   string   `xml:"InterestFrequency"`
	InterestRateType    string   `xml:"InterestRateType"`
	InterestRateType2   []string `xml:"InterestRateType2"`
	FloatingRateMargin  string   `xml:"FloatingRateMargin"`
	FloatingRateMargin2 string   `xml:"FloatingRateMargin2"`
	InterestAction      string   `xml:"InterestAction"`
	Rate2               string   `xml:"Rate2"`
	AccountName         string   `xml:"AccountName"`
	AccountNumber       string   `xml:"AccountNumber"`
	CreditFrequency     string   `xml:"CreditFrequency"`
	DebitFrequency      string   `xml:"DebitFrequency"`
	Notes               string   `xml:"Notes"`
}
type accountTransactionCollection struct {
	XMLName        xml.Name `xml:"Collection"`
	Text           string   `xml:",chardata"`
	DealType       string   `xml:"DealType"`
	MsgID          string   `xml:"MsgID"`
	ExecType       string   `xml:"ExecType"`
	TradeDate      string   `xml:"TradeDate"`
	ParentExtRefNo string   `xml:"ParentExtRefNo"`
	PaymentSystem  string   `xml:"PaymentSystem"`
	Direction      string   `xml:"Direction"`
	Adjustment     string   `xml:"Adjustment"`
	ValueDate      string   `xml:"ValueDate"`
	MaturityDate   string   `xml:"MaturityDate"`
	Notes          string   `xml:"Notes"`
	Test           string   `xml:"test"`
}
type niCollections struct {
	XMLName                   xml.Name `xml:"Collections"`
	Text                      string   `xml:",chardata"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	Xsi                       string   `xml:"xsi,attr"`
	Collection                struct {
		Text             string `xml:",chardata"`
		DealType         string `xml:"DealType"`
		ExecType         string `xml:"ExecType"`
		TradeDate        string `xml:"TradeDate"`
		ExtRefNo         string `xml:"ExtRefNo"`
		ExternalCustomer string `xml:"ExternalCustomer"`
		Direction        string `xml:"Direction"`
		Rate             string `xml:"Rate"`
		DealtCurrency    string `xml:"DealtCurrency"`
		DealtAmount      string `xml:"DealtAmount"`
		Book             string `xml:"Book"`
		ValueDate        string `xml:"ValueDate"`
		NI               string `xml:"NI"`
	} `xml:"Collection"`
}
