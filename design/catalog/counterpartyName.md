# **CounterpartyName** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**CounterpartyName** (counterpartyname) |
|Endpoint 	    |**/CounterpartyName...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-passport** ()
Friendly Name|**Counterparty Name**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/CounterpartyName/CounterpartyNameList) [Exportable]
* **View** (/CounterpartyName/CounterpartyNameView)











##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaCounterpartyNameLookup**
SQL Table Key | **CompID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**NameFirm**|String|true|true|false|false|||||Y|NameFirm||false|false|false|
|**NameCentre**|String|true|true|false|false|||||Y|NameCentre||false|false|false|
|**FullName**|String|false|true|false|false|||||Y|FullName||false|false|false|
|**CompID**|String|true|true|false|false|||||Y|CompID||false|false|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/counterpartyName_core.go |
| code | **dao** | /dao/counterpartyName_core.go |
| code | **datamodel** | /datamodel/counterpartyName_core.go |
| code | **menu** | /design/menu/counterpartyName.json |
| html | **list** | /html/CounterpartyName_List.html |
| html | **view** | /html/CounterpartyName_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **14/06/2022** at **21:31:58**
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