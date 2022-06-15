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
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**Name**|String|true|true|false|false|||||Y|Name||false|false|false|
|**AcceptReducedAmount**|Bool|false|true|false|false|||||Y|AcceptReducedAmount|True|false|false|false|
|**QuoteAsIndicative**|Bool|false|true|false|false|||||Y|QuoteAsIndicative|True|false|false|false|
|**RateTimeOut**|Int|false|true|false|false|||||Y|RateTimeOut|0|false|false|false|
|**PropagationDelay**|Int|false|true|false|false|||||Y|PropagationDelay|0|false|false|false|
|**CheckLiquidity**|Bool|false|true|false|false|||||Y|CheckLiquidity|True|false|false|false|
|**ChangeQuoteDirection**|Bool|false|true|false|false|||||Y|ChangeQuoteDirection|True|false|false|false|
|**GenerateRejectedDeals**|Bool|false|true|false|false|||||Y|GenerateRejectedDeals|True|false|false|false|
|**SpotUpdatesForForwardQuotes**|Bool|false|true|false|false|||||Y|SpotUpdatesForForwardQuotes|True|false|false|false|
|**SettlementInstructionStyle**|String|false|true|false|false|||||Y|SettlementInstructionStyle||false|false|false|
|**CanRetractQuotes**|Bool|false|true|false|false|||||Y|CanRetractQuotes|True|false|false|false|
|**CancelESPifNotPriced**|Bool|false|true|false|false|||||Y|CancelESPifNotPriced|True|false|false|false|
|**CancelRFQSifNotPriced**|Bool|false|true|false|false|||||Y|CancelRFQSifNotPriced|True|false|false|false|
|**CancelonDealingSuspended**|Bool|false|true|false|false|||||Y|CancelonDealingSuspended|True|false|false|false|
|**CreditCheckedatDI**|Bool|false|true|false|false|||||Y|CreditCheckedatDI|True|false|false|false|
|**DuplicateCheckonExternalRef**|Bool|false|true|false|false|||||Y|DuplicateCheckonExternalRef|True|false|false|false|
|**LimitCheckQuote**|Bool|false|true|false|false|||||Y|LimitCheckQuote|True|false|false|false|
|**LimitCheckonRFQDealSubmission**|Bool|false|true|false|false|||||Y|LimitCheckonRFQDealSubmission|True|false|false|false|
|**ListenonLimits**|Bool|false|true|false|false|||||Y|ListenonLimits|True|false|false|false|
|**MarginStyle**|String|false|true|false|false|||||Y|MarginStyle||false|false|false|
|**UseRerouteDefinitionOnly**|Bool|false|true|false|false|||||Y|UseRerouteDefinitionOnly|True|false|false|false|
|**BypassConfirmation**|Bool|false|true|false|false|||||Y|BypassConfirmation|True|false|false|false|
|**DIOnAcceptance**|Bool|false|true|false|false|||||Y|DIOnAcceptance|True|false|false|false|
|**IgnoreESPAmountRules**|Bool|false|true|false|false|||||Y|IgnoreESPAmountRules|True|false|false|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/dealingInterface_core.go |
| code | **adaptor** | /adaptor/dealingInterface_impl.go_template |
| code | **dao** | /dao/dealingInterface_core.go |
| code | **datamodel** | /datamodel/dealingInterface_core.go |
| code | **menu** | /design/menu/dealingInterface.json |
| html | **list** | /html/DealingInterface_List.html |
| html | **view** | /html/DealingInterface_View.html |
| html | **edit** | /html/DealingInterface_Edit.html |
| html | **new** | /html/DealingInterface_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **14/06/2022** at **21:32:03**
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