# **Firm** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Firm** (firm) |
|Endpoint 	    |**/Firm...** [^1]|
|Endpoint Query |**FirmName**|
|REST API|**/API/Firm/**|
Glyph|**fas fa-city** ()
Friendly Name|**Firm**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Firm/FirmList) [Exportable]
* **View** (/Firm/FirmView)
* **Edit** (/Firm/FirmEdit)
* **Save** (/Firm/FirmSave)
* **New** (/Firm/FirmNew)
* **Delete** (/Firm/FirmDelete)







##  Provides
 * Lookup (FirmName FullName)
 * Reverse Lookup (FullName)





##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaFirm**
SQL Table Key | **FirmName**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**FirmName**|String|false|true|false|true|||||Y|FirmName||true|false|false|
|**FullName**|String|false|true|false|false|||||Y|FullName||false|false|false|
|**Country**|String|false|true|false|false|OL|Country|Country|Name|N|Country||false|false|false|
|**Sector**|String|false|true|false|false|OL|Sector|Sector|Name|Y|Sector||false|false|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/firm_core.go |
| code | **adaptor** | /adaptor/firm_impl.go_template |
| code | **api** | /application/firm_api.go |
| code | **dao** | /dao/firm_core.go |
| code | **datamodel** | /datamodel/firm_core.go |
| code | **menu** | /design/menu/firm.json |
| html | **list** | /html/Firm_List.html |
| html | **view** | /html/Firm_View.html |
| html | **edit** | /html/Firm_Edit.html |
| html | **new** | /html/Firm_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **14/06/2022** at **21:52:05**
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