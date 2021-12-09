package adaptor

import "encoding/xml"

type LimitsWrapper struct {
	XMLName          xml.Name `xml:"LimitsWrapper"`
	Text             string   `xml:",chardata"`
	LimitDataWrapper struct {
		Text      string `xml:",chardata"`
		LimitData struct {
			Text         string `xml:",chardata"`
			Counterparty struct {
				Text             string `xml:",chardata"`
				Name             string `xml:"name"`
				LimitCurrency    string `xml:"limitCurrency"`
				LimitDetailsList struct {
					Text         string `xml:",chardata"`
					LimitDetails struct {
						Text                   string `xml:",chardata"`
						LimitType              string `xml:"limitType"`
						Limit                  string `xml:"limit"`
						LimitExtension         string `xml:"limitExtension"`
						ExtensionExpiry        string `xml:"extensionExpiry"`
						ValueDate              string `xml:"valueDate"`
						Product                string `xml:"product"`
						UtilisationDetailsList struct {
							Text               string `xml:",chardata"`
							UtilisationDetails struct {
								Text                string `xml:",chardata"`
								UtilisationCurrency string `xml:"utilisationCurrency"`
								Utilisation         string `xml:"utilisation"`
								ValueDate           string `xml:"valueDate"`
							} `xml:"utilisationDetails"`
						} `xml:"utilisationDetailsList"`
					} `xml:"limitDetails"`
				} `xml:"limitDetailsList"`
			} `xml:"counterparty"`
		} `xml:"limitData"`
	} `xml:"LimitDataWrapper"`
}

type LimitData struct {
	XMLName        xml.Name `xml:"LimitData"`
	Text           string   `xml:",chardata"`
	Xsi            string   `xml:"xsi,attr"`
	Xmlns          string   `xml:"xmlns,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Header         struct {
		Text      string `xml:",chardata"`
		TimeStamp string `xml:"timeStamp,attr"`
		MsgID     string `xml:"msgID,attr"`
	} `xml:"Header"`
	LimitsWrapper struct {
		Text             string `xml:",chardata"`
		LimitDataWrapper []struct {
			Text      string `xml:",chardata"`
			LimitData struct {
				Text         string `xml:",chardata"`
				Counterparty struct {
					Text             string `xml:",chardata"`
					Name             string `xml:"name"`
					LimitCurrency    string `xml:"limitCurrency"`
					LimitDetailsList struct {
						Text         string `xml:",chardata"`
						LimitDetails struct {
							Text                   string `xml:",chardata"`
							LimitType              string `xml:"limitType"`
							Limit                  string `xml:"limit"`
							LimitExtension         string `xml:"limitExtension"`
							UtilisationDetailsList struct {
								Text               string `xml:",chardata"`
								UtilisationDetails struct {
									Text                string `xml:",chardata"`
									UtilisationCurrency string `xml:"utilisationCurrency"`
									Utilisation         string `xml:"utilisation"`
									ValueDate           string `xml:"valueDate"`
								} `xml:"utilisationDetails"`
							} `xml:"utilisationDetailsList"`
						} `xml:"limitDetails"`
					} `xml:"limitDetailsList"`
				} `xml:"counterparty"`
			} `xml:"limitData"`
		} `xml:"LimitDataWrapper"`
	} `xml:"LimitsWrapper"`
}
