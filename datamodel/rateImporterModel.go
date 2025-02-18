<<<<<<< Updated upstream:datamodel/rateImporterModel.go
package datamodel

import (
	"encoding/xml"
	"log"
)

type RateImporterMsg struct {
	XMLName                   xml.Name `xml:"msg"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xmlns:xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	Coll                      struct {
		Text string                `xml:",chardata"`
		Rt   []RateImporterMsgItem `xml:"rt"`
	} `xml:"coll"`
}

type RateImporterMsgItem struct {
	Text   string `xml:",chardata"`
	Cat    string `xml:"cat,attr"`
	ID     string `xml:"id,attr"`
	Src    string `xml:"src,attr"`
	Bid    string `xml:"bid"`
	Ask    string `xml:"ask"`
	Own    string `xml:"own"`
	Rsk    string `xml:"rsk"`
	DTme   string `xml:"dTme"`
	Amt    string `xml:"amt"`
	AmtCcy string `xml:"amtCcy"`
	Sts    string `xml:"sts"`
}

func RateImporterTest() {
	var knickers RateImporterMsg
	knickers.NoNamespaceSchemaLocation = "eurobase-rate.xsd"
	knickers.Xsi = "http://www.w3.org/2001/XMLSchema-instance"

	var list []RateImporterMsgItem
	var item RateImporterMsgItem

	item.Cat = "SourceA"
	item.ID = "GBP="
	item.Src = "SRS"
	item.Bid = "1.5"
	item.Ask = "7.6"
	item.Own = "DealerA"
	item.Rsk = "LON"
	item.DTme = "2012-11-28T10:10:10"
	item.Amt = "10000"
	item.AmtCcy = "USD"
	item.Sts = "OK"

	list = append(list, item)

	item.Cat = "SourceA"
	item.ID = "EUR="
	item.Src = "SRS"
	item.Bid = "1.5"
	item.Ask = "7.6"
	item.Own = "DealerA"
	item.Rsk = "LON"
	item.DTme = "2012-11-28T10:10:10"
	item.Amt = "10000"
	item.AmtCcy = "USD"
	item.Sts = "OK"

	list = append(list, item)

	knickers.Coll.Rt = list
	//spew.Dump(knickers)

	newMsg, err := xml.Marshal(knickers)
	newMsg = []byte(xml.Header + string(newMsg))
	if err != nil {
		log.Println(string(newMsg), err)
	}
	log.Println(string(newMsg))

}
=======
package siena

import (
	"encoding/xml"
	"log"

	"github.com/davecgh/go-spew/spew"
)

type RateImporterMsg struct {
	XMLName                   xml.Name `xml:"msg"`
	Text                      string   `xml:",chardata"`
	Xsi                       string   `xml:"xmlns:xsi,attr"`
	NoNamespaceSchemaLocation string   `xml:"noNamespaceSchemaLocation,attr"`
	Coll                      struct {
		Text string                `xml:",chardata"`
		Rt   []RateImporterMsgItem `xml:"rt"`
	} `xml:"coll"`
}

type RateImporterMsgItem struct {
	Text string `xml:",chardata"`
	Cat  string `xml:"cat,attr"`
	ID   string `xml:"id,attr"`
	Src  string `xml:"src,attr"`
	Uid  string `xml:"uid,attr"`
	Bid  string `xml:"bid"`
	Ask  string `xml:"ask"`
	Own  string `xml:"own"`
	Rsk  string `xml:"rsk"`
	DTme string `xml:"dTme"`
	Amt  string `xml:"amt"`
	Vols struct {
		Text string `xml:",chardata"`
		Px   struct {
			Text string `xml:",chardata"`
			Stk  string `xml:"stk"`
			Vol  string `xml:"vol"`
		} `xml:"px"`
	} `xml:"vols"`
	AmtCcy string `xml:"amtCcy"`
	Sts    string `xml:"sts"`
}

func RateImporterTest() {
	var knickers RateImporterMsg
	knickers.NoNamespaceSchemaLocation = "eurobase-rate.xsd"
	knickers.Xsi = "http://www.w3.org/2001/XMLSchema-instance"

	var list []RateImporterMsgItem
	var item RateImporterMsgItem

	item.Cat = "SourceA"
	item.ID = "GBP="
	item.Src = "SRS"
	item.Bid = "1.5"
	item.Ask = "7.6"
	item.Own = "DealerA"
	item.Rsk = "LON"
	item.DTme = "2012-11-28T10:10:10"
	item.Amt = "10000"
	item.AmtCcy = "USD"
	item.Sts = "OK"

	list = append(list, item)

	item.Cat = "SourceA"
	item.ID = "EUR="
	item.Src = "SRS"
	item.Bid = "1.5"
	item.Ask = "7.6"
	item.Own = "DealerA"
	item.Rsk = "LON"
	item.DTme = "2012-11-28T10:10:10"
	item.Amt = "10000"
	item.Vols.Px.Stk = "123.45"
	item.Vols.Px.Vol = "456.78"
	item.AmtCcy = "USD"
	item.Sts = "OK"

	list = append(list, item)

	knickers.Coll.Rt = list
	spew.Dump(knickers)

	newMsg, err := xml.Marshal(knickers)
	newMsg = []byte(xml.Header + string(newMsg))
	if err != nil {
		log.Println(string(newMsg), err)
	}
	log.Println(string(newMsg))

}
>>>>>>> Stashed changes:siena/rateImporterModel.go
