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







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaDealList**
SQL Table Key | **SienaReference**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**SienaReference**|String|true|true|false|false|||||Y|SienaReference||false|false|false|
|**Status**|String|false|true|false|false|||||Y|Status||false|false|false|
|**ValueDate**|Time|false|true|false|false|||||Y|ValueDate||false|false|false|
|**MaturityDate**|Time|false|true|false|false|||||Y|MaturityDate||false|false|false|
|**ContractNumber**|String|false|true|false|false|||||Y|ContractNumber||false|false|false|
|**ExternalReference**|String|false|true|false|false|||||Y|ExternalReference||false|false|false|
|**Book**|String|false|true|false|false|||||Y|Book||false|false|false|
|**MandatedUser**|String|false|true|false|false|||||Y|MandatedUser||false|false|false|
|**Portfolio**|String|false|true|false|false|||||Y|Portfolio||false|false|false|
|**AgreementId**|Int|false|true|false|false|||||Y|AgreementId|0|false|false|false|
|**BackOfficeRefNo**|String|false|true|false|false|||||Y|BackOfficeRefNo||false|false|false|
|**ISIN**|String|false|true|false|false|||||Y|ISIN||false|false|false|
|**UTI**|String|false|true|false|false|||||Y|UTI||false|false|false|
|**BookName**|String|false|true|false|false|||||Y|BookName||false|false|false|
|**Centre**|String|false|true|false|false|||||Y|Centre||false|false|false|
|**Firm**|String|false|true|false|false|||||Y|Firm||false|false|false|
|**DealTypeShortName**|String|false|true|false|false|||||Y|DealTypeShortName||false|false|false|
|**FullDealType**|String|false|true|false|false|||||Y|FullDealType||false|false|false|
|**TradeDate**|Time|false|true|false|false|||||Y|TradeDate||false|false|false|
|**DealtCcy**|String|false|true|false|false|||||Y|DealtCcy||false|false|false|
|**DealtAmount**|Float|false|true|false|false|||||Y|DealtAmount|0.00|false|false|false|
|**AgainstAmount**|Float|false|true|false|false|||||Y|AgainstAmount|0.00|false|false|false|
|**AgainstCcy**|String|false|true|false|false|||||Y|AgainstCcy||false|false|false|
|**AllInRate**|Float|false|true|false|false|||||Y|AllInRate|0.00|false|false|false|
|**MktRate**|Float|false|true|false|false|||||Y|MktRate|0.00|false|false|false|
|**SettleCcy**|String|false|true|false|false|||||Y|SettleCcy||false|false|false|
|**Direction**|String|false|true|false|false|||||Y|Direction||false|false|false|
|**NpvRate**|Float|false|true|false|false|||||Y|NpvRate|0.00|false|false|false|
|**OriginUser**|String|false|true|false|false|||||Y|OriginUser||false|false|false|
|**PayInstruction**|String|false|true|false|false|||||Y|PayInstruction||false|false|false|
|**ReceiptInstruction**|String|false|true|false|false|||||Y|ReceiptInstruction||false|false|false|
|**NIName**|String|false|true|false|false|||||Y|NIName||false|false|false|
|**CCYPair**|String|false|true|false|false|||||Y|CCYPair||false|false|false|
|**Instrument**|String|false|true|false|false|||||Y|Instrument||false|false|false|
|**PortfolioName**|String|false|true|false|false|||||Y|PortfolioName||false|false|false|
|**RVDate**|Time|false|true|false|false|||||Y|RVDate||false|false|false|
|**RVMTM**|Float|false|true|false|false|||||Y|RVMTM|0.00|false|false|false|
|**CounterBook**|String|false|true|false|false|||||Y|CounterBook||false|false|false|
|**CounterBookName**|String|false|true|false|false|||||Y|CounterBookName||false|false|false|
|**Party**|String|false|true|false|false|||||Y|Party||false|false|false|
|**PartyName**|String|false|true|false|false|||||Y|PartyName||false|false|false|
|**NameCentre**|String|false|true|false|false|||||Y|NameCentre||false|false|false|
|**NameFirm**|String|false|true|false|false|||||Y|NameFirm||false|false|false|
|**CustomerExternalView**|String|false|true|false|false|||||Y|CustomerExternalView||false|false|false|
|**CustomerSienaView**|String|false|true|false|false|||||Y|CustomerSienaView||false|false|false|
|**CompID**|String|false|true|false|false|||||Y|CompID||false|false|false|
|**SienaDealer**|String|false|true|false|false|||||Y|SienaDealer||false|false|false|
|**DealOwner**|String|false|true|false|false|||||Y|DealOwner||false|false|false|
|**DealOwnerMnemonic**|String|false|true|false|false|||||Y|DealOwnerMnemonic||false|false|false|
|**EditedByUser**|String|false|true|false|false|||||Y|EditedByUser||false|false|false|
|**UTCOriginTime**|String|false|true|false|false|||||Y|UTCOriginTime||false|false|false|
|**UTCUpdateTime**|String|false|true|false|false|||||Y|UTCUpdateTime||false|false|false|
|**MarginTrading**|Bool|false|true|false|false|||||Y|MarginTrading|True|false|false|false|
|**SwapPoints**|Float|false|true|false|false|||||Y|SwapPoints|0.00|false|false|false|
|**SpotDate**|Time|false|true|false|false|||||Y|SpotDate||false|false|false|
|**SpotRate**|Float|false|true|false|false|||||Y|SpotRate|0.00|false|false|false|
|**MktSpotRate**|Float|false|true|false|false|||||Y|MktSpotRate|0.00|false|false|false|
|**SpotSalesMargin**|Float|false|true|false|false|||||Y|SpotSalesMargin|0.00|false|false|false|
|**SpotChlMargin**|Float|false|true|false|false|||||Y|SpotChlMargin|0.00|false|false|false|
|**RerouteCcy**|String|false|true|false|false|||||Y|RerouteCcy||false|false|false|
|**CustomerPayInstruction**|String|false|true|false|false|||||Y|CustomerPayInstruction||false|false|false|
|**CustomerReceiptInstruction**|String|false|true|false|false|||||Y|CustomerReceiptInstruction||false|false|false|
|**BackOfficeNotes**|String|false|true|false|false|||||Y|BackOfficeNotes||false|false|false|
|**CustomerStatementNotes**|String|false|true|false|false|||||Y|customerStatementNotes||false|false|false|
|**NotesMargin**|Float|false|true|false|false|||||Y|NotesMargin|0.00|false|false|false|
|**RequestedBy**|String|false|true|false|false|||||Y|RequestedBy||false|false|false|
|**EditReason**|String|false|true|false|false|||||Y|EditReason||false|false|false|
|**EditOtherReason**|String|false|true|false|false|||||Y|EditOtherReason||false|false|false|
|**NICleanPrice**|Float|false|true|false|false|||||Y|NICleanPrice|0.00|false|false|false|
|**NIDirtyPrice**|Float|false|true|false|false|||||Y|NIDirtyPrice|0.00|false|false|false|
|**NIYield**|Float|false|true|false|false|||||Y|NIYield|0.00|false|false|false|
|**NIClearingSystem**|String|false|true|false|false|||||Y|NIClearingSystem||false|false|false|
|**Acceptor**|String|false|true|false|false|||||Y|Acceptor||false|false|false|
|**NIDiscount**|Float|false|true|false|false|||||Y|NIDiscount|0.00|false|false|false|
|**FastPay**|Bool|false|true|false|false|||||Y|FastPay|True|false|false|false|
|**PaymentFee**|Float|false|true|false|false|||||Y|PaymentFee|0.00|false|false|false|
|**PaymentFeePolicy**|String|false|true|false|false|||||Y|PaymentFeePolicy||false|false|false|
|**PaymentReason**|String|false|true|false|false|||||Y|PaymentReason||false|false|false|
|**PaymentDate**|Time|false|true|false|false|||||Y|PaymentDate||false|false|false|
|**SettlementDate**|Time|false|true|false|false|||||Y|SettlementDate||false|false|false|
|**FixingDate**|Time|false|true|false|false|||||Y|FixingDate||false|false|false|
|**VenueUTI**|String|false|true|false|false|||||Y|VenueUTI||false|false|false|
|**EditVersion**|Int|false|true|false|false|||||Y|EditVersion|0|false|false|false|
|**BrokeragePercentage**|Float|false|true|false|false|||||Y|BrokeragePercentage|0.00|false|false|false|
|**BrokerageAmount**|Float|false|true|false|false|||||Y|BrokerageAmount|0.00|false|false|false|
|**BrokerageCurrency**|String|false|true|false|false|||||Y|BrokerageCurrency||false|false|false|
|**BrokerageDate**|Time|false|true|false|false|||||Y|BrokerageDate||false|false|false|
|**AccountName**|String|false|true|false|false|||||Y|AccountName||false|false|false|
|**AccountNumber**|String|false|true|false|false|||||Y|AccountNumber||false|false|false|
|**CashBalance**|Float|false|true|false|false|||||Y|CashBalance|0.00|false|false|false|
|**DebitFrequency**|String|false|true|false|false|||||Y|DebitFrequency||false|false|false|
|**CreditFrequency**|String|false|true|false|false|||||Y|CreditFrequency||false|false|false|
|**ManuallyQuoted**|String|false|true|false|false|||||Y|ManuallyQuoted||false|false|false|
|**LedgerBalance**|Float|false|true|false|false|||||Y|LedgerBalance|0.00|false|false|false|
|**SettAmtOutstanding**|Float|false|true|false|false|||||Y|SettAmtOutstanding|0.00|false|false|false|
|**FeePercentage**|Float|false|true|false|false|||||Y|FeePercentage|0.00|false|false|false|
|**FeeAmount**|Float|false|true|false|false|||||Y|FeeAmount|0.00|false|false|false|
|**Venue**|String|false|true|false|false|||||Y|Venue||false|false|false|
|**EURAmount**|Float|false|true|false|false|||||Y|EURAmount|0.00|false|false|false|
|**EUROtherAmount**|Float|false|true|false|false|||||Y|EUROtherAmount|0.00|false|false|false|
|**LEI**|String|false|true|false|false|||||Y|LEI||false|false|false|
|**Equity**|String|false|true|false|false|||||Y|Equity||false|false|false|
|**Shares**|Int|false|true|false|false|||||Y|Shares|0|false|false|false|
|**QuoteExpiryDate**|Time|false|true|false|false|||||Y|QuoteExpiryDate||false|false|false|
|**Commodity**|String|false|true|false|false|||||Y|Commodity||false|false|false|
|**PaymentSystemSienaView**|String|false|true|false|false|||||Y|PaymentSystemSienaView||false|false|false|
|**PaymentSystemExternalView**|String|false|true|false|false|||||Y|PaymentSystemExternalView||false|false|false|
|**SalesProfit**|Float|false|true|false|false|||||Y|SalesProfit|0.00|false|false|false|
|**RejectReason**|String|false|true|false|false|||||Y|RejectReason||false|false|false|
|**PaymentError**|String|false|true|false|false|||||Y|PaymentError||false|false|false|
|**RepoPrincipal**|Float|false|true|false|false|||||Y|RepoPrincipal|0.00|false|false|false|
|**FixingFrequency**|String|false|true|false|false|||||Y|FixingFrequency||false|false|false|
|**Dealt**|String|false|false|true|false|||||N|||false|true|false|
|**Against**|String|false|false|true|false|||||N|||false|true|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/transaction_core.go |
| code | **adaptor** | /adaptor/transaction_impl.go_template |
| code | **dao** | /dao/transaction_core.go |
| code | **datamodel** | /datamodel/transaction_core.go |
| code | **menu** | /design/menu/transaction.json |
| html | **list** | /html/Transaction_List.html |
| html | **view** | /html/Transaction_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **14/06/2022** at **21:31:46**
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