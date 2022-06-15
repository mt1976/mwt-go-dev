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
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**DealTypeKey**|String|true|true|false|false|||||Y|DealTypeKey||false|false|false|
|**DealTypeShortName**|String|false|true|false|false|||||Y|DealTypeShortName||false|false|false|
|**HostKey**|String|false|true|false|false|||||Y|HostKey||false|false|false|
|**IsActive**|Bool|false|true|false|false|||||Y|IsActive|True|false|false|false|
|**Interbook**|Bool|false|true|false|false|||||Y|Interbook|True|false|false|false|
|**BackOfficeLink**|Bool|false|true|false|false|||||Y|BackOfficeLink|True|false|false|false|
|**HasTicket**|Bool|false|true|false|false|||||Y|HasTicket|True|false|false|false|
|**CurrencyOverride**|Bool|false|true|false|false|||||Y|CurrencyOverride|True|false|false|false|
|**CurrencyHolderCurrency**|String|false|true|false|false|||||Y|CurrencyHolderCurrency||false|false|false|
|**AllBooks**|Bool|false|true|false|false|||||Y|AllBooks|True|false|false|false|
|**FundamentalDealTypeKey**|String|false|true|false|false|||||Y|FundamentalDealTypeKey||false|false|false|
|**RelatedDealType**|String|false|true|false|false|||||Y|RelatedDealType||false|false|false|
|**BookName**|String|false|true|false|false|||||Y|BookName||false|false|false|
|**ExportMethod**|String|false|true|false|false|||||Y|ExportMethod||false|false|false|
|**DefaultUserLayoffBooks**|Bool|false|true|false|false|||||Y|DefaultUserLayoffBooks|True|false|false|false|
|**RFQ**|Bool|false|true|false|false|||||Y|RFQ|True|false|false|false|
|**OBS**|Bool|false|true|false|false|||||Y|OBS|True|false|false|false|
|**KID**|Bool|false|true|false|false|||||Y|KID|True|false|false|false|
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
| code | **application** | /application/dealType_core.go |
| code | **adaptor** | /adaptor/dealType_impl.go_template |
| code | **dao** | /dao/dealType_core.go |
| code | **datamodel** | /datamodel/dealType_core.go |
| code | **menu** | /design/menu/dealType.json |
| html | **list** | /html/DealType_List.html |
| html | **view** | /html/DealType_View.html |
| html | **edit** | /html/DealType_Edit.html |
| html | **new** | /html/DealType_New.html |


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