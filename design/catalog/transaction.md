# **Transaction** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Transaction** (transaction) |
|Endpoint 	    |**/Transaction...** [^1]|
|Endpoint Query |**Ref**|
Glyph|**far fa-handshake** ()
Friendly Name|**Transaction**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Transaction/TransactionList) [Exportable]
* **View** (/Transaction/TransactionView)











##  Provides
 * Lookup (SienaReference SienaReference)






##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaDealList**
SQL Table Key | **SienaReference**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SienaReference**|String|true|true|false|false|||||Y|SienaReference||false|false|false|text||
|**Status**|String|false|true|false|false|||||Y|Status||false|false|false|text||
|**ValueDate**|Time|false|true|false|false|||||Y|ValueDate||false|false|false|text||
|**MaturityDate**|Time|false|true|false|false|||||Y|MaturityDate||false|false|false|text||
|**ContractNumber**|String|false|true|false|false|||||Y|ContractNumber||false|false|false|text||
|**ExternalReference**|String|false|true|false|false|||||Y|ExternalReference||false|false|false|text||
|**Book**|String|false|true|false|false|||||Y|Book||false|false|false|text||
|**MandatedUser**|String|false|true|false|false|||||Y|MandatedUser||false|false|false|text||
|**Portfolio**|String|false|true|false|false|||||Y|Portfolio||false|false|false|text||
|**AgreementId**|Int|false|true|false|false|||||Y|AgreementId|0|false|false|false|text||
|**BackOfficeRefNo**|String|false|true|false|false|||||Y|BackOfficeRefNo||false|false|false|text||
|**ISIN**|String|false|true|false|false|||||Y|ISIN||false|false|false|text||
|**UTI**|String|false|true|false|false|||||Y|UTI||false|false|false|text||
|**BookName**|String|false|true|false|false|||||Y|BookName||false|false|false|text||
|**Centre**|String|false|true|false|false|||||Y|Centre||false|false|false|text||
|**Firm**|String|false|true|false|false|||||Y|Firm||false|false|false|text||
|**DealTypeShortName**|String|false|true|false|false|||||Y|DealTypeShortName||false|false|false|text||
|**FullDealType**|String|false|true|false|false|||||Y|FullDealType||false|false|false|text||
|**TradeDate**|Time|false|true|false|false|||||Y|TradeDate||false|false|false|text||
|**DealtCcy**|String|false|true|false|false|||||Y|DealtCcy||false|false|false|text||
|**DealtAmount**|Float|false|true|false|false|||||Y|DealtAmount|0.00|false|false|false|text||
|**AgainstAmount**|Float|false|true|false|false|||||Y|AgainstAmount|0.00|false|false|false|text||
|**AgainstCcy**|String|false|true|false|false|||||Y|AgainstCcy||false|false|false|text||
|**AllInRate**|Float|false|true|false|false|||||Y|AllInRate|0.00|false|false|false|text||
|**MktRate**|Float|false|true|false|false|||||Y|MktRate|0.00|false|false|false|text||
|**SettleCcy**|String|false|true|false|false|||||Y|SettleCcy||false|false|false|text||
|**Direction**|String|false|true|false|false|||||Y|Direction||false|false|false|text||
|**NpvRate**|Float|false|true|false|false|||||Y|NpvRate|0.00|false|false|false|text||
|**OriginUser**|String|false|true|false|false|||||Y|OriginUser||false|false|false|text||
|**PayInstruction**|String|false|true|false|false|||||Y|PayInstruction||false|false|false|text||
|**ReceiptInstruction**|String|false|true|false|false|||||Y|ReceiptInstruction||false|false|false|text||
|**NIName**|String|false|true|false|false|||||Y|NIName||false|false|false|text||
|**CCYPair**|String|false|true|false|false|||||Y|CCYPair||false|false|false|text||
|**Instrument**|String|false|true|false|false|||||Y|Instrument||false|false|false|text||
|**PortfolioName**|String|false|true|false|false|||||Y|PortfolioName||false|false|false|text||
|**RVDate**|Time|false|true|false|false|||||Y|RVDate||false|false|false|text||
|**RVMTM**|Float|false|true|false|false|||||Y|RVMTM|0.00|false|false|false|text||
|**CounterBook**|String|false|true|false|false|||||Y|CounterBook||false|false|false|text||
|**CounterBookName**|String|false|true|false|false|||||Y|CounterBookName||false|false|false|text||
|**Party**|String|false|true|false|false|||||Y|Party||false|false|false|text||
|**PartyName**|String|false|true|false|false|||||Y|PartyName||false|false|false|text||
|**NameCentre**|String|false|true|false|false|||||Y|NameCentre||false|false|false|text||
|**NameFirm**|String|false|true|false|false|||||Y|NameFirm||false|false|false|text||
|**CustomerExternalView**|String|false|true|false|false|||||Y|CustomerExternalView||false|false|false|text||
|**CustomerSienaView**|String|false|true|false|false|||||Y|CustomerSienaView||false|false|false|text||
|**CompID**|String|false|true|false|false|||||Y|CompID||false|false|false|text||
|**SienaDealer**|String|false|true|false|false|||||Y|SienaDealer||false|false|false|text||
|**DealOwner**|String|false|true|false|false|||||Y|DealOwner||false|false|false|text||
|**DealOwnerMnemonic**|String|false|true|false|false|||||Y|DealOwnerMnemonic||false|false|false|text||
|**EditedByUser**|String|false|true|false|false|||||Y|EditedByUser||false|false|false|text||
|**UTCOriginTime**|String|false|true|false|false|||||Y|UTCOriginTime||false|false|false|text||
|**UTCUpdateTime**|String|false|true|false|false|||||Y|UTCUpdateTime||false|false|false|text||
|**MarginTrading**|Bool|false|true|false|false|||||Y|MarginTrading|True|false|false|false|text||
|**SwapPoints**|Float|false|true|false|false|||||Y|SwapPoints|0.00|false|false|false|text||
|**SpotDate**|Time|false|true|false|false|||||Y|SpotDate||false|false|false|text||
|**SpotRate**|Float|false|true|false|false|||||Y|SpotRate|0.00|false|false|false|text||
|**MktSpotRate**|Float|false|true|false|false|||||Y|MktSpotRate|0.00|false|false|false|text||
|**SpotSalesMargin**|Float|false|true|false|false|||||Y|SpotSalesMargin|0.00|false|false|false|text||
|**SpotChlMargin**|Float|false|true|false|false|||||Y|SpotChlMargin|0.00|false|false|false|text||
|**RerouteCcy**|String|false|true|false|false|||||Y|RerouteCcy||false|false|false|text||
|**CustomerPayInstruction**|String|false|true|false|false|||||Y|CustomerPayInstruction||false|false|false|text||
|**CustomerReceiptInstruction**|String|false|true|false|false|||||Y|CustomerReceiptInstruction||false|false|false|text||
|**BackOfficeNotes**|String|false|true|false|false|||||Y|BackOfficeNotes||false|false|false|text||
|**CustomerStatementNotes**|String|false|true|false|false|||||Y|customerStatementNotes||false|false|false|text||
|**NotesMargin**|Float|false|true|false|false|||||Y|NotesMargin|0.00|false|false|false|text||
|**RequestedBy**|String|false|true|false|false|||||Y|RequestedBy||false|false|false|text||
|**EditReason**|String|false|true|false|false|||||Y|EditReason||false|false|false|text||
|**EditOtherReason**|String|false|true|false|false|||||Y|EditOtherReason||false|false|false|text||
|**NICleanPrice**|Float|false|true|false|false|||||Y|NICleanPrice|0.00|false|false|false|text||
|**NIDirtyPrice**|Float|false|true|false|false|||||Y|NIDirtyPrice|0.00|false|false|false|text||
|**NIYield**|Float|false|true|false|false|||||Y|NIYield|0.00|false|false|false|text||
|**NIClearingSystem**|String|false|true|false|false|||||Y|NIClearingSystem||false|false|false|text||
|**Acceptor**|String|false|true|false|false|||||Y|Acceptor||false|false|false|text||
|**NIDiscount**|Float|false|true|false|false|||||Y|NIDiscount|0.00|false|false|false|text||
|**FastPay**|Bool|false|true|false|false|||||Y|FastPay|True|false|false|false|text||
|**PaymentFee**|Float|false|true|false|false|||||Y|PaymentFee|0.00|false|false|false|text||
|**PaymentFeePolicy**|String|false|true|false|false|||||Y|PaymentFeePolicy||false|false|false|text||
|**PaymentReason**|String|false|true|false|false|||||Y|PaymentReason||false|false|false|text||
|**PaymentDate**|Time|false|true|false|false|||||Y|PaymentDate||false|false|false|text||
|**SettlementDate**|Time|false|true|false|false|||||Y|SettlementDate||false|false|false|text||
|**FixingDate**|Time|false|true|false|false|||||Y|FixingDate||false|false|false|text||
|**VenueUTI**|String|false|true|false|false|||||Y|VenueUTI||false|false|false|text||
|**EditVersion**|Int|false|true|false|false|||||Y|EditVersion|0|false|false|false|text||
|**BrokeragePercentage**|Float|false|true|false|false|||||Y|BrokeragePercentage|0.00|false|false|false|text||
|**BrokerageAmount**|Float|false|true|false|false|||||Y|BrokerageAmount|0.00|false|false|false|text||
|**BrokerageCurrency**|String|false|true|false|false|||||Y|BrokerageCurrency||false|false|false|text||
|**BrokerageDate**|Time|false|true|false|false|||||Y|BrokerageDate||false|false|false|text||
|**AccountName**|String|false|true|false|false|||||Y|AccountName||false|false|false|text||
|**AccountNumber**|String|false|true|false|false|||||Y|AccountNumber||false|false|false|text||
|**CashBalance**|Float|false|true|false|false|||||Y|CashBalance|0.00|false|false|false|text||
|**DebitFrequency**|String|false|true|false|false|||||Y|DebitFrequency||false|false|false|text||
|**CreditFrequency**|String|false|true|false|false|||||Y|CreditFrequency||false|false|false|text||
|**ManuallyQuoted**|String|false|true|false|false|||||Y|ManuallyQuoted||false|false|false|text||
|**LedgerBalance**|Float|false|true|false|false|||||Y|LedgerBalance|0.00|false|false|false|text||
|**SettAmtOutstanding**|Float|false|true|false|false|||||Y|SettAmtOutstanding|0.00|false|false|false|text||
|**FeePercentage**|Float|false|true|false|false|||||Y|FeePercentage|0.00|false|false|false|text||
|**FeeAmount**|Float|false|true|false|false|||||Y|FeeAmount|0.00|false|false|false|text||
|**Venue**|String|false|true|false|false|||||Y|Venue||false|false|false|text||
|**EURAmount**|Float|false|true|false|false|||||Y|EURAmount|0.00|false|false|false|text||
|**EUROtherAmount**|Float|false|true|false|false|||||Y|EUROtherAmount|0.00|false|false|false|text||
|**LEI**|String|false|true|false|false|||||Y|LEI||false|false|false|text||
|**Equity**|String|false|true|false|false|||||Y|Equity||false|false|false|text||
|**Shares**|Int|false|true|false|false|||||Y|Shares|0|false|false|false|text||
|**QuoteExpiryDate**|Time|false|true|false|false|||||Y|QuoteExpiryDate||false|false|false|text||
|**Commodity**|String|false|true|false|false|||||Y|Commodity||false|false|false|text||
|**PaymentSystemSienaView**|String|false|true|false|false|||||Y|PaymentSystemSienaView||false|false|false|text||
|**PaymentSystemExternalView**|String|false|true|false|false|||||Y|PaymentSystemExternalView||false|false|false|text||
|**SalesProfit**|Float|false|true|false|false|||||Y|SalesProfit|0.00|false|false|false|text||
|**RejectReason**|String|false|true|false|false|||||Y|RejectReason||false|false|false|text||
|**PaymentError**|String|false|true|false|false|||||Y|PaymentError||false|false|false|text||
|**RepoPrincipal**|Float|false|true|false|false|||||Y|RepoPrincipal|0.00|false|false|false|text||
|**FixingFrequency**|String|false|true|false|false|||||Y|FixingFrequency||false|false|false|text||
|**Dealt**|String|false|false|true|false|||||N|||false|true|false|text||
|**Against**|String|false|false|true|false|||||N|||false|true|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/transaction_core.go |
| code | **adaptor** | /adaptor/transaction_impl.go_template |
| code | **dao** | /dao/transaction_core.go |
| code | **datamodel** | /datamodel/transaction_core.go |
| code | **menu** | /design/menu/transaction.json |
| html | **list** | /Transaction_List.html |
| html | **view** | /Transaction_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **28/06/2022** at **16:10:40**
Who & Where		     | **matttownsend (Matt Townsend)** on **silicon.local**

### Footnotes
[^1]: **Endpoint**
    * The full list of endpoints can be found in the [Actions](#action-id) section
[^2]: **Lookup Fields**
    * LL = A List Lookup. Define list in lits.cfg
    * OL = An Object Lookup. Get a list of values from an Object
    * FL = Fetches 1 value from an object based on the content of the field. 
[^3]: **Inputtable**   
    * H = Hidden Field
    * N = No Input Field
    * Y = Inputtable Field