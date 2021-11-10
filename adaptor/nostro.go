package adaptor

import "encoding/xml"

type NostroXML struct {
	XMLName xml.Name `xml:"BulkTradeCapture"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text      string `xml:",chardata"`
		TimeStamp string `xml:"timeStamp,attr"`
		MsgID     string `xml:"msgID,attr"`
	} `xml:"Header"`
	BulkInfo struct {
		Text     string `xml:",chardata"`
		DealType string `xml:"dealType,attr"`
	} `xml:"BulkInfo"`
	DealData []struct {
		Text     string `xml:",chardata"`
		DealInfo struct {
			Text      string `xml:",chardata"`
			Dir       string `xml:"dir,attr"`
			ExtRefNo  string `xml:"extRefNo,attr"`
			ValueDate string `xml:"valueDate,attr"`
		} `xml:"DealInfo"`
		Monies struct {
			Text     string `xml:",chardata"`
			DealtAmt string `xml:"dealtAmt,attr"`
			DealtCcy string `xml:"dealtCcy,attr"`
		} `xml:"Monies"`
		SettlementInfo []struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			ID   string `xml:"ID,attr"`
		} `xml:"SettlementInfo"`
		TradingEntity struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
			ID   string `xml:"ID,attr"`
		} `xml:"TradingEntity"`
	} `xml:"DealData"`
}
