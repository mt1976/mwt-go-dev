# **AccountLadder** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**AccountLadder** (accountladder) |
|Endpoint 	    |**/AccountLadder...** [^1]|
|Endpoint Query |**AccountNo**|
Glyph|**fas fa-landmark** ()
Friendly Name|**Account Ladder**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/AccountLadder/AccountLadderList) [Exportable]
* **View** (/AccountLadder/AccountLadderView)











##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaAccountLadder**
SQL Table Key | **SienaReference**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**SienaReference**|String|true|true|false|false|||||Y|SienaReference||false|false|false|
|**BusinessDate**|Time|true|true|false|false|||||Y|BusinessDate||false|false|false|
|**ContractNumber**|String|false|true|false|false|||||Y|ContractNumber||false|false|false|
|**Balance**|Float|false|true|false|false|||||Y|Balance|0.00|false|false|false|
|**DealtCcy**|String|false|true|false|false|||||Y|DealtCcy||false|false|false|
|**AmountDp**|Int|false|true|false|false|||||Y|AmountDp|0|false|false|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/accountLadder_core.go |
| code | **adaptor** | /adaptor/accountLadder_impl.go_template |
| code | **dao** | /dao/accountLadder_core.go |
| code | **datamodel** | /datamodel/accountLadder_core.go |
| code | **menu** | /design/menu/accountLadder.json |
| html | **list** | /html/AccountLadder_List.html |
| html | **view** | /html/AccountLadder_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **14/06/2022** at **21:31:47**
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