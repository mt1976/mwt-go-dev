# **Product** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Product** (product) |
|Endpoint 	    |**/Product...** [^1]|
|Endpoint Query |**Code**|
|REST API|**/API/Product/**|
Glyph|**fas fa-object-group** ()
Friendly Name|**Product Group**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Product/ProductList) [Exportable]
* **View** (/Product/ProductView)
* **Edit** (/Product/ProductEdit)
* **Save** (/Product/ProductSave)
* **New** (/Product/ProductNew)
* **Delete** (/Product/ProductDelete)







##  Provides

 * Reverse Lookup (Name)





##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaProduct**
SQL Table Key | **Code**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**Code**|String|true|true|false|false|||||Y|Code||false|false|false|text||
|**Name**|String|false|true|false|false|||||Y|Name||false|false|false|text||
|**Factor**|Float|false|true|false|false|||||Y|Factor|0.00|false|false|false|text||
|**MaxTermPrecedence**|Bool|false|true|false|false|||||Y|MaxTermPrecedence|True|false|false|false|text||
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
| code | **application** | /application/product_core.go |
| code | **adaptor** | /adaptor/product_impl.go_template |
| code | **api** | /application/product_api.go |
| code | **dao** | /dao/product_core.go |
| code | **datamodel** | /datamodel/product_core.go |
| code | **menu** | /design/menu/product.json |
| html | **list** | /Product_List.html |
| html | **view** | /Product_View.html |
| html | **edit** | /Product_Edit.html |
| html | **new** | /Product_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **26/06/2022** at **18:48:31**
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