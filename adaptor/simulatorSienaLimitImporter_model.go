package adaptor

import "encoding/xml"

type SienaLimitImportCounterparty struct {
	XMLName                  xml.Name `xml:"Record"`
	Text                     string   `xml:",chardata"`
	DetailedLimitUtilisation struct {
		Text          string `xml:",chardata"`
		SetAttribute1 string `xml:"setAttribute1,attr"`
		Counterparty  struct {
			Text         string `xml:",chardata"`
			SetAttribute string `xml:"setAttribute,attr"`
			CPTY         struct {
				Text        string `xml:",chardata"`
				KeepIfTrue  string `xml:"keepIfTrue,attr"`
				SetContents string `xml:"setContents,attr"`
			} `xml:"CPTY"`
			CCY struct {
				Text        string `xml:",chardata"`
				KeepIfTrue  string `xml:"keepIfTrue,attr"`
				SetContents string `xml:"setContents,attr"`
			} `xml:"CCY"`
			LimitType struct {
				Text        string `xml:",chardata"`
				KeepIfTrue  string `xml:"keepIfTrue,attr"`
				SetContents string `xml:"setContents,attr"`
			} `xml:"LimitType"`
			SettDate struct {
				Text        string `xml:",chardata"`
				KeepIfTrue  string `xml:"keepIfTrue,attr"`
				SetContents string `xml:"setContents,attr"`
			} `xml:"SettDate"`
			Limit struct {
				Text        string `xml:",chardata"`
				KeepIfTrue  string `xml:"keepIfTrue,attr"`
				SetContents string `xml:"setContents,attr"`
			} `xml:"Limit"`
			Extension struct {
				Text        string `xml:",chardata"`
				KeepIfTrue  string `xml:"keepIfTrue,attr"`
				SetContents string `xml:"setContents,attr"`
			} `xml:"Extension"`
			CPTYUtilisation struct {
				Text        string `xml:",chardata"`
				KeepIfTrue  string `xml:"keepIfTrue,attr"`
				SetContents string `xml:"setContents,attr"`
			} `xml:"CPTYUtilisation"`
			Availability struct {
				Text        string `xml:",chardata"`
				KeepIfTrue  string `xml:"keepIfTrue,attr"`
				SetContents string `xml:"setContents,attr"`
			} `xml:"Availability"`
			MaxTerm struct {
				Text        string `xml:",chardata"`
				KeepIfTrue  string `xml:"keepIfTrue,attr"`
				SetContents string `xml:"setContents,attr"`
			} `xml:"MaxTerm"`
			ExtExpiry struct {
				Text        string `xml:",chardata"`
				KeepIfTrue  string `xml:"keepIfTrue,attr"`
				SetContents string `xml:"setContents,attr"`
			} `xml:"ExtExpiry"`
			Deal struct {
				Text          string `xml:",chardata"`
				DeleteIfTrue  string `xml:"deleteIfTrue,attr"`
				Iterations    string `xml:"iterations,attr"`
				IterationName string `xml:"iterationName,attr"`
				RefNo         struct {
					Text        string `xml:",chardata"`
					KeepIfTrue  string `xml:"keepIfTrue,attr"`
					SetContents string `xml:"setContents,attr"`
				} `xml:"RefNo"`
				DealtCCY struct {
					Text        string `xml:",chardata"`
					KeepIfTrue  string `xml:"keepIfTrue,attr"`
					SetContents string `xml:"setContents,attr"`
				} `xml:"DealtCCY"`
				DealLimitType struct {
					Text        string `xml:",chardata"`
					KeepIfTrue  string `xml:"keepIfTrue,attr"`
					SetContents string `xml:"setContents,attr"`
				} `xml:"DealLimitType"`
				UtilisationDate struct {
					Text        string `xml:",chardata"`
					KeepIfTrue  string `xml:"keepIfTrue,attr"`
					SetContents string `xml:"setContents,attr"`
				} `xml:"UtilisationDate"`
				DealType struct {
					Text        string `xml:",chardata"`
					KeepIfTrue  string `xml:"keepIfTrue,attr"`
					SetContents string `xml:"setContents,attr"`
				} `xml:"DealType"`
				DealtAmount struct {
					Text        string `xml:",chardata"`
					KeepIfTrue  string `xml:"keepIfTrue,attr"`
					SetContents string `xml:"setContents,attr"`
				} `xml:"DealtAmount"`
				DealUtilisation struct {
					Text        string `xml:",chardata"`
					KeepIfTrue  string `xml:"keepIfTrue,attr"`
					SetContents string `xml:"setContents,attr"`
				} `xml:"DealUtilisation"`
			} `xml:"Deal"`
		} `xml:"Counterparty"`
	} `xml:"DetailedLimitUtilisation"`
}

// <?xml version="1.0" encoding="iso-8859-15"?>

// <Record>
//   <DetailedLimitUtilisation setAttribute1="key=counterpartyGetter()">
//     <Counterparty setAttribute="key=counterpartyGetter()">
//       <CPTY keepIfTrue="true" setContents="counterpartyGetter()"/>
//       <CCY keepIfTrue="true" setContents="limitCurrencyGetter()"/>
//       <LimitType keepIfTrue="true" setContents="limitTypeGetter()"/>
//       <SettDate keepIfTrue="true" setContents="settlementDateGetter()"/>
//       <Limit keepIfTrue="true" setContents="limitAmountGetter()"/>
//       <Extension keepIfTrue="true" setContents="extensionGetter()"/>
//       <CPTYUtilisation keepIfTrue="true" setContents="utilisationAmountGetter()"/>
//       <Availability keepIfTrue="true" setContents="availabilityGetter()"/>
//       <MaxTerm keepIfTrue="true" setContents="maxTermGetter()"/>
//       <ExtExpiry keepIfTrue="true" setContents="extensionExpiryGetter()"/>
//       <Deal deleteIfTrue="numDealsGetter()=0" iterations="numDealsGetter()" iterationName="dealIteration">
//         <RefNo keepIfTrue="true" setContents="Deal + dealRefNoGetter()"/>
//         <DealtCCY keepIfTrue="true" setContents="dealtCurrencyGetter()"/>
//         <DealLimitType keepIfTrue="true" setContents="CPTY + limitTypeGetter()"/>
//         <UtilisationDate keepIfTrue="true" setContents="dealUtilisationDateGetter()"/>
//         <DealType keepIfTrue="true" setContents="dealTypeGetter()"/>
//         <DealtAmount keepIfTrue="true" setContents="dealtAmountGetter()"/>
//         <DealUtilisation keepIfTrue="true" setContents="dealUtilisationGetter()"/>
//       </Deal>
//     </Counterparty>
//   </DetailedLimitUtilisation>
// </Record>

type SienaLimitImportUser struct {
	XMLName          xml.Name `xml:"Record"`
	Text             string   `xml:",chardata"`
	LimitUtilisation struct {
		Text          string `xml:",chardata"`
		SetAttribute1 string `xml:"setAttribute1,attr"`
		Counterparty  struct {
			Text        string `xml:",chardata"`
			KeepIfTrue  string `xml:"keepIfTrue,attr"`
			SetContents string `xml:"setContents,attr"`
		} `xml:"Counterparty"`
		User struct {
			Text        string `xml:",chardata"`
			KeepIfTrue  string `xml:"keepIfTrue,attr"`
			SetContents string `xml:"setContents,attr"`
		} `xml:"User"`
		LimitType struct {
			Text        string `xml:",chardata"`
			KeepIfTrue  string `xml:"keepIfTrue,attr"`
			SetContents string `xml:"setContents,attr"`
		} `xml:"LimitType"`
		ProductType struct {
			Text        string `xml:",chardata"`
			KeepIfTrue  string `xml:"keepIfTrue,attr"`
			SetContents string `xml:"setContents,attr"`
		} `xml:"ProductType"`
		LimitCurrency struct {
			Text        string `xml:",chardata"`
			KeepIfTrue  string `xml:"keepIfTrue,attr"`
			SetContents string `xml:"setContents,attr"`
		} `xml:"LimitCurrency"`
		LimitAmount struct {
			Text        string `xml:",chardata"`
			KeepIfTrue  string `xml:"keepIfTrue,attr"`
			SetContents string `xml:"setContents,attr"`
			AmountView  string `xml:"amountView,attr"`
		} `xml:"LimitAmount"`
		UtilisationAmount struct {
			Text        string `xml:",chardata"`
			KeepIfTrue  string `xml:"keepIfTrue,attr"`
			SetContents string `xml:"setContents,attr"`
			AmountView  string `xml:"amountView,attr"`
		} `xml:"UtilisationAmount"`
		Availability struct {
			Text        string `xml:",chardata"`
			KeepIfTrue  string `xml:"keepIfTrue,attr"`
			SetContents string `xml:"setContents,attr"`
			AmountView  string `xml:"amountView,attr"`
		} `xml:"Availability"`
	} `xml:"LimitUtilisation"`
}

// <?xml version="1.0" encoding="windows-1252"?>

// <!--
//     Document   : UserLimitUtilisation.xml
//     Created on : 02 February 2011, 04:57
//     Author     : tpethiyagoda
//     Description: The report shows limit utilisation information for users.
// -->

// <Record>
//   <LimitUtilisation setAttribute1="key=counterpartyGetter()" >
//     <Counterparty keepIfTrue="true" setContents="counterpartyGetter()"/>
//     <User keepIfTrue="true" setContents="entityNameGetter()"/>
//     <LimitType keepIfTrue="true" setContents="limitTypeGetter()"/>
//     <ProductType keepIfTrue="true" setContents="productTypeGetter()"/>
//     <LimitCurrency keepIfTrue="true" setContents="limitCurrencyGetter()"/>
//     <LimitAmount keepIfTrue="true" setContents="limitAmountGetter()" amountView="thousand" />
//     <UtilisationAmount keepIfTrue="true" setContents="utilisationAmountGetter()" amountView="thousand"/>
//     <Availability keepIfTrue="true" setContents="availabilityGetter()" amountView="thousand"/>
//   </LimitUtilisation>
// </Record>
