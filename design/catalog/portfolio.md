# **Portfolio** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Portfolio** (portfolio) |
|Endpoint 	    |**/Portfolio...** [^1]|
|Endpoint Query |**Code**|
|REST API|**/API/Portfolio/**|
Glyph|**fas fa-plug** ()
Friendly Name|**Bank Porfolio**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Portfolio/PortfolioList) [Exportable]
* **View** (/Portfolio/PortfolioView)
* **Edit** (/Portfolio/PortfolioEdit)
* **Save** (/Portfolio/PortfolioSave)
* **New** (/Portfolio/PortfolioNew)
* **Delete** (/Portfolio/PortfolioDelete)







##  Provides
 * Lookup (Code Description1)
 * Reverse Lookup (Description1)





##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaPortfolio**
SQL Table Key | **Code**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**Code**|String|true|true|false|false|||||Y|Code||false|false|false|text||
|**Description1**|String|false|true|false|false|||||Y|Description1||false|false|false|text||
|**Description2**|String|false|true|false|false|||||Y|Description2||false|false|false|text||
|**IsDefault**|Bool|false|true|false|false|||||Y|isDefault|True|false|false|false|text||
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
| code | **application** | /application/portfolio_core.go |
| code | **adaptor** | /adaptor/portfolio_impl.go_template |
| code | **api** | /application/portfolio_api.go |
| code | **dao** | /dao/portfolio_core.go |
| code | **datamodel** | /datamodel/portfolio_core.go |
| code | **menu** | /design/menu/portfolio.json |
| html | **list** | /Portfolio_List.html |
| html | **view** | /Portfolio_View.html |
| html | **edit** | /Portfolio_Edit.html |
| html | **new** | /Portfolio_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **28/06/2022** at **16:10:55**
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