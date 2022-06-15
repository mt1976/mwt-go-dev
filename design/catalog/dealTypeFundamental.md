# **DealTypeFundamental** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**DealTypeFundamental** (dealtypefundamental) |
|Endpoint 	    |**/DealTypeFundamental...** [^1]|
|Endpoint Query |**DealTypeKey**|
Glyph|**fas fa-toolbox** ()
Friendly Name|**Deal Type Fundamental**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/DealTypeFundamental/DealTypeFundamentalList) [Exportable]
* **View** (/DealTypeFundamental/DealTypeFundamentalView)
* **Edit** (/DealTypeFundamental/DealTypeFundamentalEdit)
* **Save** (/DealTypeFundamental/DealTypeFundamentalSave)
* **New** (/DealTypeFundamental/DealTypeFundamentalNew)
* **Delete** (/DealTypeFundamental/DealTypeFundamentalDelete)







##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaDealTypeFundamentals**
SQL Table Key | **DealTypeKey**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**DealTypeKey**|String|true|true|false|false|||||Y|DealTypeKey||false|false|false|
|**Amendment**|Bool|false|true|false|false|||||Y|Amendment|True|false|false|false|
|**DefaultFrequency**|Int|false|true|false|false|||||Y|DefaultFrequency|0|false|false|false|
|**Directions**|String|false|true|false|false|||||Y|Directions||false|false|false|
|**SettledTermDealType**|String|false|true|false|false|||||Y|SettledTermDealType||false|false|false|
|**Optn**|Bool|false|true|false|false|||||Y|Optn|True|false|false|false|
|**AllowPledge**|Bool|false|true|false|false|||||Y|AllowPledge|True|false|false|false|
|**Takeup**|Bool|false|true|false|false|||||Y|Takeup|True|false|false|false|
|**MismatchDealType**|String|false|true|false|false|||||Y|MismatchDealType||false|false|false|
|**SettledHedgeTermDealType**|String|false|true|false|false|||||Y|SettledHedgeTermDealType||false|false|false|
|**SettlementCode**|String|false|true|false|false|||||Y|SettlementCode||false|false|false|
|**TermSubType**|String|false|true|false|false|||||Y|TermSubType||false|false|false|
|**FundingDealType**|String|false|true|false|false|||||Y|FundingDealType||false|false|false|
|**TransferType**|String|false|true|false|false|||||Y|TransferType||false|false|false|
|**TermDealType**|String|false|true|false|false|||||Y|TermDealType||false|false|false|
|**NegotiableInstrumentType**|String|false|true|false|false|||||Y|NegotiableInstrumentType||false|false|false|
|**Mismatch**|Bool|false|true|false|false|||||Y|Mismatch|True|false|false|false|
|**ComplexTransferSubType**|String|false|true|false|false|||||Y|ComplexTransferSubType||false|false|false|
|**LayOffDealType**|String|false|true|false|false|||||Y|LayOffDealType||false|false|false|
|**NIAccount**|Int|false|true|false|false|||||Y|NIAccount|0|false|false|false|
|**SimpleMMsubtype**|Int|false|true|false|false|||||Y|SimpleMMsubtype|0|false|false|false|
|**SwapDealType**|String|false|true|false|false|||||Y|SwapDealType||false|false|false|
|**Positions**|String|false|true|false|false|||||Y|Positions||false|false|false|
|**OptionOutright**|String|false|true|false|false|||||Y|OptionOutright||false|false|false|
|**SettledHedgeSpotDealType**|String|false|true|false|false|||||Y|SettledHedgeSpotDealType||false|false|false|
|**StraightThroughInterestMethod**|Bool|false|true|false|false|||||Y|StraightThroughInterestMethod|True|false|false|false|
|**SubType**|String|false|true|false|false|||||Y|SubType||false|false|false|
|**Rollover**|Bool|false|true|false|false|||||Y|Rollover|True|false|false|false|
|**DefaultIssuer**|String|false|true|false|false|||||Y|DefaultIssuer||false|false|false|
|**DefaultStartDate**|Int|false|true|false|false|||||Y|DefaultStartDate|0|false|false|false|
|**Fee**|String|false|true|false|false|||||Y|Fee||false|false|false|
|**NDF**|Bool|false|true|false|false|||||Y|NDF|True|false|false|false|
|**FXFX**|Bool|false|true|false|false|||||Y|FXFX|True|false|false|false|
|**ONIA**|Bool|false|true|false|false|||||Y|ONIA|True|false|false|false|
|**MarginSubType**|Int|false|true|false|false|||||Y|MarginSubType|0|false|false|false|
|**TransferDealType**|String|false|true|false|false|||||Y|TransferDealType||false|false|false|
|**IsFX**|Bool|false|true|false|false|||||Y|IsFX|True|false|false|false|
|**Ordr**|String|false|true|false|false|||||Y|Ordr||false|false|false|
|**OptionStyle**|String|false|true|false|false|||||Y|OptionStyle||false|false|false|
|**SpotDealType**|String|false|true|false|false|||||Y|SpotDealType||false|false|false|
|**CanIssue**|Bool|false|true|false|false|||||Y|CanIssue|True|false|false|false|
|**CanShort**|Bool|false|true|false|false|||||Y|CanShort|True|false|false|false|
|**FXMarginTradingType**|Int|false|true|false|false|||||Y|FXMarginTradingType|0|false|false|false|
|**Internal**|Bool|false|true|false|false|||||Y|Internal|True|false|false|false|
|**TicketBasename**|String|false|true|false|false|||||Y|TicketBasename||false|false|false|
|**InterestRateFutureType**|String|false|true|false|false|||||Y|InterestRateFutureType||false|false|false|
|**TradingLimitProductCode**|String|false|true|false|false|||||Y|TradingLimitProductCode||false|false|false|
|**Forward**|Bool|false|true|false|false|||||Y|Forward|True|false|false|false|
|**MaturityNotificationPeriod**|String|false|true|false|false|||||Y|MaturityNotificationPeriod||false|false|false|
|**NotificationEvents**|String|false|true|false|false|||||Y|NotificationEvents||false|false|false|
|**SwapSubType**|String|false|true|false|false|||||Y|SwapSubType||false|false|false|
|**ProductCode**|String|false|true|false|false|||||Y|ProductCode||false|false|false|
|**Funding**|Bool|false|true|false|false|||||Y|Funding|True|false|false|false|
|**AllocationPricing**|String|false|true|false|false|||||Y|AllocationPricing||false|false|false|
|**CancelPeriod**|String|false|true|false|false|||||Y|CancelPeriod||false|false|false|
|**MMMarginTradingType**|Int|false|true|false|false|||||Y|MMMarginTradingType|0|false|false|false|
|**OptionSpot**|String|false|true|false|false|||||Y|OptionSpot||false|false|false|
|**Transfer**|Bool|false|true|false|false|||||Y|Transfer|True|false|false|false|
|**NotificationPeriod**|String|false|true|false|false|||||Y|NotificationPeriod||false|false|false|
|**Paymentdateshift**|Int|false|true|false|false|||||Y|Paymentdateshift|0|false|false|false|
|**CloseOut**|Bool|false|true|false|false|||||Y|CloseOut|True|false|false|false|
|**FXOptionPricing**|String|false|true|false|false|||||Y|FXOptionPricing||false|false|false|
|**SettledHedgeOutrightDealType**|String|false|true|false|false|||||Y|SettledHedgeOutrightDealType||false|false|false|
|**RepoBond**|String|false|true|false|false|||||Y|RepoBond||false|false|false|
|**RepoTerm**|String|false|true|false|false|||||Y|RepoTerm||false|false|false|
|**RepoType**|Int|false|true|false|false|||||Y|RepoType|0|false|false|false|
|**DateRule**|String|false|true|false|false|||||Y|DateRule||false|false|false|
|**CorpTransferDealType**|String|false|true|false|false|||||Y|CorpTransferDealType||false|false|false|
|**GenerateFXImage**|Bool|false|true|false|false|||||Y|GenerateFXImage|True|false|false|false|
|**Variant**|String|false|true|false|false|||||Y|Variant||false|false|false|
|**HedgeTermDealType**|String|false|true|false|false|||||Y|HedgeTermDealType||false|false|false|
|**PricingModel**|String|false|true|false|false|||||Y|PricingModel||false|false|false|
|**HedgeOutrightDealType**|String|false|true|false|false|||||Y|HedgeOutrightDealType||false|false|false|
|**Fixing**|Bool|false|true|false|false|||||Y|Fixing|True|false|false|false|
|**Payment**|Bool|false|true|false|false|||||Y|Payment|True|false|false|false|
|**MT**|Bool|false|true|false|false|||||Y|MT|True|false|false|false|
|**SettlementInstructionStyle**|String|false|true|false|false|||||Y|SettlementInstructionStyle||false|false|false|
|**QuoteHistoryRequired**|Bool|false|true|false|false|||||Y|QuoteHistoryRequired|True|false|false|false|
|**Brokerage**|Bool|false|true|false|false|||||Y|Brokerage|True|false|false|false|
|**ExposureDisabled**|Bool|false|true|false|false|||||Y|ExposureDisabled|True|false|false|false|
|**CreditLine**|String|false|true|false|false|||||Y|CreditLine||false|false|false|
|**Encumbered**|Bool|false|true|false|false|||||Y|Encumbered|True|false|false|false|
|**InternalId**|Int|true|true|false|false|||||Y|InternalId|0|false|false|false|
|**InternalDeleted**|Time|false|true|false|false|||||Y|InternalDeleted||false|false|false|
|**UpdatedTransactionId**|String|false|true|false|false|||||Y|UpdatedTransactionId||false|false|false|
|**UpdatedUserId**|String|false|true|false|false|||||Y|UpdatedUserId||false|false|false|
|**UpdatedDateTime**|Time|false|true|false|false|||||Y|UpdatedDateTime||false|false|false|
|**DeletedTransactionId**|String|false|true|false|false|||||Y|DeletedTransactionId||false|false|false|
|**DeletedUserId**|String|false|true|false|false|||||Y|DeletedUserId||false|false|false|
|**ChangeType**|String|false|true|false|false|||||Y|ChangeType||false|false|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/dealTypeFundamental_core.go |
| code | **adaptor** | /adaptor/dealTypeFundamental_impl.go_template |
| code | **dao** | /dao/dealTypeFundamental_core.go |
| code | **datamodel** | /datamodel/dealTypeFundamental_core.go |
| code | **menu** | /design/menu/dealTypeFundamental.json |
| html | **list** | /html/DealTypeFundamental_List.html |
| html | **view** | /html/DealTypeFundamental_View.html |
| html | **edit** | /html/DealTypeFundamental_Edit.html |
| html | **new** | /html/DealTypeFundamental_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **14/06/2022** at **21:32:04**
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