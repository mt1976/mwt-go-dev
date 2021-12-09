package adaptor

import "encoding/xml"

type CoppClarkCalender struct {
	XMLName xml.Name `xml:"root"`
	Text    string   `xml:",chardata"`
	Row     []struct {
		Text                   string `xml:",chardata"`
		CenterID               string `xml:"CenterID"`
		ISOCurrencyCode        string `xml:"ISOCurrencyCode"`
		ISOCountryCode         string `xml:"ISOCountryCode"`
		RelatedFinancialCentre string `xml:"RelatedFinancialCentre"`
		FileType               string `xml:"FileType"`
		EventYear              []struct {
			Text  string `xml:",chardata"`
			Year  string `xml:"Year"`
			Event []struct {
				Text           string `xml:",chardata"`
				EventDate      string `xml:"EventDate"`
				EventDayOfWeek string `xml:"EventDayOfWeek"`
				EventName      string `xml:"EventName"`
			} `xml:"Event"`
		} `xml:"EventYear"`
	} `xml:"row"`
}
