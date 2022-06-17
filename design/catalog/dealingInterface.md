# **DealingInterface** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**DealingInterface** (dealinginterface) |
|Endpoint 	    |**/DealingInterface...** [^1]|
|Endpoint Query |**Name**|
Glyph|**fas fa-share-alt** ()
Friendly Name|**Dealing Interface**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/DealingInterface/DealingInterfaceList) [Exportable]
* **View** (/DealingInterface/DealingInterfaceView)
* **Edit** (/DealingInterface/DealingInterfaceEdit)
* **Save** (/DealingInterface/DealingInterfaceSave)
* **New** (/DealingInterface/DealingInterfaceNew)
* **Delete** (/DealingInterface/DealingInterfaceDelete)







##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaDealingInterface**
SQL Table Key | **Name**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**Name**|String|true|true|false|false|||||Y|Name||false|false|false|text||
|**AcceptReducedAmount**|Bool|false|true|false|false|||||Y|AcceptReducedAmount|True|false|false|false|text||
|**QuoteAsIndicative**|Bool|false|true|false|false|||||Y|QuoteAsIndicative|True|false|false|false|text||
|**RateTimeOut**|Int|false|true|false|false|||||Y|RateTimeOut|0|false|false|false|text||
|**PropagationDelay**|Int|false|true|false|false|||||Y|PropagationDelay|0|false|false|false|text||
|**CheckLiquidity**|Bool|false|true|false|false|||||Y|CheckLiquidity|True|false|false|false|text||
|**ChangeQuoteDirection**|Bool|false|true|false|false|||||Y|ChangeQuoteDirection|True|false|false|false|text||
|**GenerateRejectedDeals**|Bool|false|true|false|false|||||Y|GenerateRejectedDeals|True|false|false|false|text||
|**SpotUpdatesForForwardQuotes**|Bool|false|true|false|false|||||Y|SpotUpdatesForForwardQuotes|True|false|false|false|text||
|**SettlementInstructionStyle**|String|false|true|false|false|||||Y|SettlementInstructionStyle||false|false|false|text||
|**CanRetractQuotes**|Bool|false|true|false|false|||||Y|CanRetractQuotes|True|false|false|false|text||
|**CancelESPifNotPriced**|Bool|false|true|false|false|||||Y|CancelESPifNotPriced|True|false|false|false|text||
|**CancelRFQSifNotPriced**|Bool|false|true|false|false|||||Y|CancelRFQSifNotPriced|True|false|false|false|text||
|**CancelonDealingSuspended**|Bool|false|true|false|false|||||Y|CancelonDealingSuspended|True|false|false|false|text||
|**CreditCheckedatDI**|Bool|false|true|false|false|||||Y|CreditCheckedatDI|True|false|false|false|text||
|**DuplicateCheckonExternalRef**|Bool|false|true|false|false|||||Y|DuplicateCheckonExternalRef|True|false|false|false|text||
|**LimitCheckQuote**|Bool|false|true|false|false|||||Y|LimitCheckQuote|True|false|false|false|text||
|**LimitCheckonRFQDealSubmission**|Bool|false|true|false|false|||||Y|LimitCheckonRFQDealSubmission|True|false|false|false|text||
|**ListenonLimits**|Bool|false|true|false|false|||||Y|ListenonLimits|True|false|false|false|text||
|**MarginStyle**|String|false|true|false|false|||||Y|MarginStyle||false|false|false|text||
|**UseRerouteDefinitionOnly**|Bool|false|true|false|false|||||Y|UseRerouteDefinitionOnly|True|false|false|false|text||
|**BypassConfirmation**|Bool|false|true|false|false|||||Y|BypassConfirmation|True|false|false|false|text||
|**DIOnAcceptance**|Bool|false|true|false|false|||||Y|DIOnAcceptance|True|false|false|false|text||
|**IgnoreESPAmountRules**|Bool|false|true|false|false|||||Y|IgnoreESPAmountRules|True|false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/dealingInterface_core.go |
| code | **adaptor** | /adaptor/dealingInterface_impl.go_template |
| code | **dao** | /dao/dealingInterface_core.go |
| code | **datamodel** | /datamodel/dealingInterface_core.go |
| code | **menu** | /design/menu/dealingInterface.json |
| html | **list** | /DealingInterface_List.html |
| html | **view** | /DealingInterface_View.html |
| html | **edit** | /DealingInterface_Edit.html |
| html | **new** | /DealingInterface_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **17/06/2022** at **18:38:10**
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