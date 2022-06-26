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
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**DealTypeKey**|String|true|true|false|false|||||Y|DealTypeKey||false|false|false|text||
|**Amendment**|Bool|false|true|false|false|||||Y|Amendment|True|false|false|false|text||
|**DefaultFrequency**|Int|false|true|false|false|||||Y|DefaultFrequency|0|false|false|false|text||
|**Directions**|String|false|true|false|false|||||Y|Directions||false|false|false|text||
|**SettledTermDealType**|String|false|true|false|false|||||Y|SettledTermDealType||false|false|false|text||
|**Optn**|Bool|false|true|false|false|||||Y|Optn|True|false|false|false|text||
|**AllowPledge**|Bool|false|true|false|false|||||Y|AllowPledge|True|false|false|false|text||
|**Takeup**|Bool|false|true|false|false|||||Y|Takeup|True|false|false|false|text||
|**MismatchDealType**|String|false|true|false|false|||||Y|MismatchDealType||false|false|false|text||
|**SettledHedgeTermDealType**|String|false|true|false|false|||||Y|SettledHedgeTermDealType||false|false|false|text||
|**SettlementCode**|String|false|true|false|false|||||Y|SettlementCode||false|false|false|text||
|**TermSubType**|String|false|true|false|false|||||Y|TermSubType||false|false|false|text||
|**FundingDealType**|String|false|true|false|false|||||Y|FundingDealType||false|false|false|text||
|**TransferType**|String|false|true|false|false|||||Y|TransferType||false|false|false|text||
|**TermDealType**|String|false|true|false|false|||||Y|TermDealType||false|false|false|text||
|**NegotiableInstrumentType**|String|false|true|false|false|||||Y|NegotiableInstrumentType||false|false|false|text||
|**Mismatch**|Bool|false|true|false|false|||||Y|Mismatch|True|false|false|false|text||
|**ComplexTransferSubType**|String|false|true|false|false|||||Y|ComplexTransferSubType||false|false|false|text||
|**LayOffDealType**|String|false|true|false|false|||||Y|LayOffDealType||false|false|false|text||
|**NIAccount**|Int|false|true|false|false|||||Y|NIAccount|0|false|false|false|text||
|**SimpleMMsubtype**|Int|false|true|false|false|||||Y|SimpleMMsubtype|0|false|false|false|text||
|**SwapDealType**|String|false|true|false|false|||||Y|SwapDealType||false|false|false|text||
|**Positions**|String|false|true|false|false|||||Y|Positions||false|false|false|text||
|**OptionOutright**|String|false|true|false|false|||||Y|OptionOutright||false|false|false|text||
|**SettledHedgeSpotDealType**|String|false|true|false|false|||||Y|SettledHedgeSpotDealType||false|false|false|text||
|**StraightThroughInterestMethod**|Bool|false|true|false|false|||||Y|StraightThroughInterestMethod|True|false|false|false|text||
|**SubType**|String|false|true|false|false|||||Y|SubType||false|false|false|text||
|**Rollover**|Bool|false|true|false|false|||||Y|Rollover|True|false|false|false|text||
|**DefaultIssuer**|String|false|true|false|false|||||Y|DefaultIssuer||false|false|false|text||
|**DefaultStartDate**|Int|false|true|false|false|||||Y|DefaultStartDate|0|false|false|false|text||
|**Fee**|String|false|true|false|false|||||Y|Fee||false|false|false|text||
|**NDF**|Bool|false|true|false|false|||||Y|NDF|True|false|false|false|text||
|**FXFX**|Bool|false|true|false|false|||||Y|FXFX|True|false|false|false|text||
|**ONIA**|Bool|false|true|false|false|||||Y|ONIA|True|false|false|false|text||
|**MarginSubType**|Int|false|true|false|false|||||Y|MarginSubType|0|false|false|false|text||
|**TransferDealType**|String|false|true|false|false|||||Y|TransferDealType||false|false|false|text||
|**IsFX**|Bool|false|true|false|false|||||Y|IsFX|True|false|false|false|text||
|**Ordr**|String|false|true|false|false|||||Y|Ordr||false|false|false|text||
|**OptionStyle**|String|false|true|false|false|||||Y|OptionStyle||false|false|false|text||
|**SpotDealType**|String|false|true|false|false|||||Y|SpotDealType||false|false|false|text||
|**CanIssue**|Bool|false|true|false|false|||||Y|CanIssue|True|false|false|false|text||
|**CanShort**|Bool|false|true|false|false|||||Y|CanShort|True|false|false|false|text||
|**FXMarginTradingType**|Int|false|true|false|false|||||Y|FXMarginTradingType|0|false|false|false|text||
|**Internal**|Bool|false|true|false|false|||||Y|Internal|True|false|false|false|text||
|**TicketBasename**|String|false|true|false|false|||||Y|TicketBasename||false|false|false|text||
|**InterestRateFutureType**|String|false|true|false|false|||||Y|InterestRateFutureType||false|false|false|text||
|**TradingLimitProductCode**|String|false|true|false|false|||||Y|TradingLimitProductCode||false|false|false|text||
|**Forward**|Bool|false|true|false|false|||||Y|Forward|True|false|false|false|text||
|**MaturityNotificationPeriod**|String|false|true|false|false|||||Y|MaturityNotificationPeriod||false|false|false|text||
|**NotificationEvents**|String|false|true|false|false|||||Y|NotificationEvents||false|false|false|text||
|**SwapSubType**|String|false|true|false|false|||||Y|SwapSubType||false|false|false|text||
|**ProductCode**|String|false|true|false|false|||||Y|ProductCode||false|false|false|text||
|**Funding**|Bool|false|true|false|false|||||Y|Funding|True|false|false|false|text||
|**AllocationPricing**|String|false|true|false|false|||||Y|AllocationPricing||false|false|false|text||
|**CancelPeriod**|String|false|true|false|false|||||Y|CancelPeriod||false|false|false|text||
|**MMMarginTradingType**|Int|false|true|false|false|||||Y|MMMarginTradingType|0|false|false|false|text||
|**OptionSpot**|String|false|true|false|false|||||Y|OptionSpot||false|false|false|text||
|**Transfer**|Bool|false|true|false|false|||||Y|Transfer|True|false|false|false|text||
|**NotificationPeriod**|String|false|true|false|false|||||Y|NotificationPeriod||false|false|false|text||
|**Paymentdateshift**|Int|false|true|false|false|||||Y|Paymentdateshift|0|false|false|false|text||
|**CloseOut**|Bool|false|true|false|false|||||Y|CloseOut|True|false|false|false|text||
|**FXOptionPricing**|String|false|true|false|false|||||Y|FXOptionPricing||false|false|false|text||
|**SettledHedgeOutrightDealType**|String|false|true|false|false|||||Y|SettledHedgeOutrightDealType||false|false|false|text||
|**RepoBond**|String|false|true|false|false|||||Y|RepoBond||false|false|false|text||
|**RepoTerm**|String|false|true|false|false|||||Y|RepoTerm||false|false|false|text||
|**RepoType**|Int|false|true|false|false|||||Y|RepoType|0|false|false|false|text||
|**DateRule**|String|false|true|false|false|||||Y|DateRule||false|false|false|text||
|**CorpTransferDealType**|String|false|true|false|false|||||Y|CorpTransferDealType||false|false|false|text||
|**GenerateFXImage**|Bool|false|true|false|false|||||Y|GenerateFXImage|True|false|false|false|text||
|**Variant**|String|false|true|false|false|||||Y|Variant||false|false|false|text||
|**HedgeTermDealType**|String|false|true|false|false|||||Y|HedgeTermDealType||false|false|false|text||
|**PricingModel**|String|false|true|false|false|||||Y|PricingModel||false|false|false|text||
|**HedgeOutrightDealType**|String|false|true|false|false|||||Y|HedgeOutrightDealType||false|false|false|text||
|**Fixing**|Bool|false|true|false|false|||||Y|Fixing|True|false|false|false|text||
|**Payment**|Bool|false|true|false|false|||||Y|Payment|True|false|false|false|text||
|**MT**|Bool|false|true|false|false|||||Y|MT|True|false|false|false|text||
|**SettlementInstructionStyle**|String|false|true|false|false|||||Y|SettlementInstructionStyle||false|false|false|text||
|**QuoteHistoryRequired**|Bool|false|true|false|false|||||Y|QuoteHistoryRequired|True|false|false|false|text||
|**Brokerage**|Bool|false|true|false|false|||||Y|Brokerage|True|false|false|false|text||
|**ExposureDisabled**|Bool|false|true|false|false|||||Y|ExposureDisabled|True|false|false|false|text||
|**CreditLine**|String|false|true|false|false|||||Y|CreditLine||false|false|false|text||
|**Encumbered**|Bool|false|true|false|false|||||Y|Encumbered|True|false|false|false|text||
|**InternalId**|Int|true|true|false|false|||||Y|InternalId|0|false|false|false|text||
|**InternalDeleted**|Time|false|true|false|false|||||Y|InternalDeleted||false|false|false|text||
|**UpdatedTransactionId**|String|false|true|false|false|||||Y|UpdatedTransactionId||false|false|false|text||
|**UpdatedUserId**|String|false|true|false|false|||||Y|UpdatedUserId||false|false|false|text||
|**UpdatedDateTime**|Time|false|true|false|false|||||Y|UpdatedDateTime||false|false|false|text||
|**DeletedTransactionId**|String|false|true|false|false|||||Y|DeletedTransactionId||false|false|false|text||
|**DeletedUserId**|String|false|true|false|false|||||Y|DeletedUserId||false|false|false|text||
|**ChangeType**|String|false|true|false|false|||||Y|ChangeType||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/dealTypeFundamental_core.go |
| code | **adaptor** | /adaptor/dealTypeFundamental_impl.go_template |
| code | **dao** | /dao/dealTypeFundamental_core.go |
| code | **datamodel** | /datamodel/dealTypeFundamental_core.go |
| code | **menu** | /design/menu/dealTypeFundamental.json |
| html | **list** | /DealTypeFundamental_List.html |
| html | **view** | /DealTypeFundamental_View.html |
| html | **edit** | /DealTypeFundamental_Edit.html |
| html | **new** | /DealTypeFundamental_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **26/06/2022** at **18:48:28**
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