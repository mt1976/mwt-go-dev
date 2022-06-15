# **AccountTransaction** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**AccountTransaction** (accounttransaction) |
|Endpoint 	    |**/AccountTransaction...** [^1]|
|Endpoint Query |**AccountNo**|
Glyph|**fas fa-landmark** ()
Friendly Name|**Account Transaction**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/AccountTransaction/AccountTransactionList) [Exportable]
* **View** (/AccountTransaction/AccountTransactionView)











##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaAccountTransactions**
SQL Table Key | **SienaReference**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**SienaReference**|String|true|true|false|false|||||Y|SienaReference||false|false|false|
|**LegNo**|Int|true|true|false|false|||||Y|LegNo|0|false|false|false|
|**MMLegNo**|Int|true|true|false|false|||||Y|MMLegNo|0|false|false|false|
|**Narrative**|String|false|true|false|false|||||Y|Narrative||false|false|false|
|**Amount**|Float|false|true|false|false|||||Y|Amount|0.00|false|false|false|
|**StartInterestDate**|Time|false|true|false|false|||||Y|StartInterestDate||false|false|false|
|**EndInterestDate**|Time|false|true|false|false|||||Y|EndInterestDate||false|false|false|
|**Amortisation**|Float|false|true|false|false|||||Y|Amortisation|0.00|false|false|false|
|**InterestAmount**|Float|false|true|false|false|||||Y|InterestAmount|0.00|false|false|false|
|**InterestAction**|String|false|true|false|false|||||Y|InterestAction||false|false|false|
|**FixingDate**|Time|false|true|false|false|||||Y|FixingDate||false|false|false|
|**InterestCalculationDate**|Time|false|true|false|false|||||Y|InterestCalculationDate||false|false|false|
|**AmendmentAmount**|Float|false|true|false|false|||||Y|AmendmentAmount|0.00|false|false|false|
|**DealtCcy**|String|false|true|false|false|||||Y|DealtCcy||false|false|false|
|**AmountDp**|Int|false|true|false|false|||||Y|AmountDp|0|false|false|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/accountTransaction_core.go |
| code | **adaptor** | /adaptor/accountTransaction_impl.go_template |
| code | **dao** | /dao/accountTransaction_core.go |
| code | **datamodel** | /datamodel/accountTransaction_core.go |
| code | **menu** | /design/menu/accountTransaction.json |
| html | **list** | /html/AccountTransaction_List.html |
| html | **view** | /html/AccountTransaction_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **14/06/2022** at **21:31:48**
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