# **CounterpartyAddress** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**CounterpartyAddress** (counterpartyaddress) |
|Endpoint 	    |**/CounterpartyAddress...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-passport** ()
Friendly Name|**Counterparty Address**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/CounterpartyAddress/CounterpartyAddressList) [Exportable]
* **View** (/CounterpartyAddress/CounterpartyAddressView)
* **Edit** (/CounterpartyAddress/CounterpartyAddressEdit)
* **Save** (/CounterpartyAddress/CounterpartyAddressSave)
* **New** (/CounterpartyAddress/CounterpartyAddressNew)
* **Delete** (/CounterpartyAddress/CounterpartyAddressDelete)







##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaCounterpartyAddress**
SQL Table Key | **CompID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**NameFirm**|String|true|true|false|false|||||Y|NameFirm||false|false|false|text||
|**NameCentre**|String|true|true|false|false|||||Y|NameCentre||false|false|false|text||
|**Address1**|String|false|true|false|false|||||Y|Address1||false|false|false|text||
|**Address2**|String|false|true|false|false|||||Y|Address2||false|false|false|text||
|**CityTown**|String|false|true|false|false|||||Y|CityTown||false|false|false|text||
|**County**|String|false|true|false|false|||||Y|County||false|false|false|text||
|**Postcode**|String|false|true|false|false|||||Y|Postcode||false|false|false|text||
|**CompID**|String|true|true|false|false|||||Y|CompID||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/counterpartyAddress_core.go |
| code | **adaptor** | /adaptor/counterpartyAddress_impl.go_template |
| code | **dao** | /dao/counterpartyAddress_core.go |
| code | **datamodel** | /datamodel/counterpartyAddress_core.go |
| code | **menu** | /design/menu/counterpartyAddress.json |
| html | **list** | /CounterpartyAddress_List.html |
| html | **view** | /CounterpartyAddress_View.html |
| html | **edit** | /CounterpartyAddress_Edit.html |
| html | **new** | /CounterpartyAddress_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **28/06/2022** at **16:10:47**
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