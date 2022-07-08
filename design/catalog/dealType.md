# **DealType** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**DealType** (dealtype) |
|Endpoint 	    |**/DealType...** [^1]|
|Endpoint Query |**DealTypeKey**|
Glyph|**fas fa-toolbox** ()
Friendly Name|**Deal Type**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/DealType/DealTypeList) [Exportable]
* **View** (/DealType/DealTypeView)
* **Edit** (/DealType/DealTypeEdit)
* **Save** (/DealType/DealTypeSave)
* **New** (/DealType/DealTypeNew)
* **Delete** (/DealType/DealTypeDelete)







##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaDealType**
SQL Table Key | **DealTypeKey**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**DealTypeKey**|String|true|true|false|false|||||Y|DealTypeKey||false|false|false|text||
|**DealTypeShortName**|String|false|true|false|false|||||Y|DealTypeShortName||false|false|false|text||
|**HostKey**|String|false|true|false|false|||||Y|HostKey||false|false|false|text||
|**IsActive**|Bool|false|true|false|false|||||Y|IsActive|True|false|false|false|text||
|**Interbook**|Bool|false|true|false|false|||||Y|Interbook|True|false|false|false|text||
|**BackOfficeLink**|Bool|false|true|false|false|||||Y|BackOfficeLink|True|false|false|false|text||
|**HasTicket**|Bool|false|true|false|false|||||Y|HasTicket|True|false|false|false|text||
|**CurrencyOverride**|Bool|false|true|false|false|||||Y|CurrencyOverride|True|false|false|false|text||
|**CurrencyHolderCurrency**|String|false|true|false|false|||||Y|CurrencyHolderCurrency||false|false|false|text||
|**AllBooks**|Bool|false|true|false|false|||||Y|AllBooks|True|false|false|false|text||
|**FundamentalDealTypeKey**|String|false|true|false|false|||||Y|FundamentalDealTypeKey||false|false|false|text||
|**RelatedDealType**|String|false|true|false|false|||||Y|RelatedDealType||false|false|false|text||
|**BookName**|String|false|true|false|false|||||Y|BookName||false|false|false|text||
|**ExportMethod**|String|false|true|false|false|||||Y|ExportMethod||false|false|false|text||
|**DefaultUserLayoffBooks**|Bool|false|true|false|false|||||Y|DefaultUserLayoffBooks|True|false|false|false|text||
|**RFQ**|Bool|false|true|false|false|||||Y|RFQ|True|false|false|false|text||
|**OBS**|Bool|false|true|false|false|||||Y|OBS|True|false|false|false|text||
|**KID**|Bool|false|true|false|false|||||Y|KID|True|false|false|false|text||
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
| code | **application** | /application/dealType_core.go |
| code | **adaptor** | /adaptor/dealType_impl.go_template |
| code | **dao** | /dao/dealType_core.go |
| code | **datamodel** | /datamodel/dealType_core.go |
| code | **menu** | /design/menu/dealType.json |
| html | **list** | /DealType_List.html |
| html | **view** | /DealType_View.html |
| html | **edit** | /DealType_Edit.html |
| html | **new** | /DealType_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **28/06/2022** at **16:10:51**
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