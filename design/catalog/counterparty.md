# **Counterparty** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Counterparty** (counterparty) |
|Endpoint 	    |**/Counterparty...** [^1]|
|Endpoint Query |**ID**|
|REST API|**/API/Counterparty/**|
Glyph|**fas fa-user** ()
Friendly Name|**Counterparty**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Counterparty/CounterpartyList) [Exportable]
* **View** (/Counterparty/CounterpartyView)
* **Edit** (/Counterparty/CounterpartyEdit)
* **Save** (/Counterparty/CounterpartySave)

* **Delete** (/Counterparty/CounterpartyDelete)







##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaCounterparty**
SQL Table Key | **CompID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**NameCentre**|String|true|true|false|false|OL|Centre|Code|Name|Y|NameCentre||true|false|false|text||
|**NameFirm**|String|true|true|false|false|OL|Firm|FirmName|FullName|Y|NameFirm||true|false|false|text||
|**FullName**|String|false|true|false|false|||||Y|FullName||false|false|false|text||
|**TelephoneNumber**|String|false|true|false|true|||||Y|TelephoneNumber||false|false|false|tel||
|**EmailAddress**|String|false|true|false|true|||||Y|EmailAddress||false|false|false|email||
|**CustomerType**|String|true|true|false|false|LL|counterpartytypes|||Y|CustomerType||false|false|false|text||
|**AccountOfficer**|String|false|true|false|false|||||Y|AccountOfficer||false|false|false|text||
|**CountryCode**|String|false|true|false|false|OL|Country|||Y|CountryCode||false|false|false|text||
|**SectorCode**|String|false|true|false|false|OL|Sector|||Y|SectorCode||false|false|false|text||
|**CpartyGroupName**|String|false|true|false|false|OL|CounterpartyGroup|||Y|CpartyGroupName||false|false|false|text||
|**Notes**|String|false|true|false|true|||||Y|Notes||false|false|false|textarea||
|**Owner**|String|false|true|false|false|OL|Owner|||Y|Owner||false|false|false|text||
|**Authorised**|Bool|false|true|false|false|LL|tf|||Y|Authorised|True|false|false|false|text||
|**NameFirmName**|String|false|true|false|true|||||H|NameFirmName||false|false|false|text||
|**NameCentreName**|String|false|true|false|true|||||H|NameCentreName||false|false|false|text||
|**CountryCodeName**|String|false|true|false|true|||||H|CountryCodeName||false|false|false|text||
|**SectorCodeName**|String|false|true|false|true|||||H|SectorCodeName||false|false|false|text||
|**CompID**|String|false|true|false|true|||||H|CompID||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/counterparty_core.go |
| code | **adaptor** | /adaptor/counterparty_impl.go_template |
| code | **api** | /application/counterparty_api.go |
| code | **dao** | /dao/counterparty_core.go |
| code | **datamodel** | /datamodel/counterparty_core.go |
| code | **menu** | /design/menu/counterparty.json |
| html | **list** | /Counterparty_List.html |
| html | **view** | /Counterparty_View.html |
| html | **edit** | /Counterparty_Edit.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **26/06/2022** at **18:48:21**
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