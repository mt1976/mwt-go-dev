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
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SienaReference**|String|true|true|false|false|||||Y|SienaReference||false|false|false|text||
|**LegNo**|Int|true|true|false|false|||||Y|LegNo|0|false|false|false|text||
|**MMLegNo**|Int|true|true|false|false|||||Y|MMLegNo|0|false|false|false|text||
|**Narrative**|String|false|true|false|false|||||Y|Narrative||false|false|false|text||
|**Amount**|Float|false|true|false|false|||||Y|Amount|0.00|false|false|false|text||
|**StartInterestDate**|Time|false|true|false|false|||||Y|StartInterestDate||false|false|false|text||
|**EndInterestDate**|Time|false|true|false|false|||||Y|EndInterestDate||false|false|false|text||
|**Amortisation**|Float|false|true|false|false|||||Y|Amortisation|0.00|false|false|false|text||
|**InterestAmount**|Float|false|true|false|false|||||Y|InterestAmount|0.00|false|false|false|text||
|**InterestAction**|String|false|true|false|false|||||Y|InterestAction||false|false|false|text||
|**FixingDate**|Time|false|true|false|false|||||Y|FixingDate||false|false|false|text||
|**InterestCalculationDate**|Time|false|true|false|false|||||Y|InterestCalculationDate||false|false|false|text||
|**AmendmentAmount**|Float|false|true|false|false|||||Y|AmendmentAmount|0.00|false|false|false|text||
|**DealtCcy**|String|false|true|false|false|||||Y|DealtCcy||false|false|false|text||
|**AmountDp**|Int|false|true|false|false|||||Y|AmountDp|0|false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/accountTransaction_core.go |
| code | **adaptor** | /adaptor/accountTransaction_impl.go_template |
| code | **dao** | /dao/accountTransaction_core.go |
| code | **datamodel** | /datamodel/accountTransaction_core.go |
| code | **menu** | /design/menu/accountTransaction.json |
| html | **list** | /AccountTransaction_List.html |
| html | **view** | /AccountTransaction_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **28/06/2022** at **16:10:42**
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