# **Sector** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Sector** (sector) |
|Endpoint 	    |**/Sector...** [^1]|
|Endpoint Query |**Sector**|
|REST API|**/API/Sector/**|
Glyph|**fas fa-industry** ()
Friendly Name|**Sector**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Sector/SectorList) [Exportable]
* **View** (/Sector/SectorView)
* **Edit** (/Sector/SectorEdit)
* **Save** (/Sector/SectorSave)
* **New** (/Sector/SectorNew)
* **Delete** (/Sector/SectorDelete)







##  Provides
 * Lookup (Code Name)
 * Reverse Lookup (Name)





##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaSector**
SQL Table Key | **Code**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**Code**|String|true|true|false|false|||||Y|Code||false|false|false|text||
|**Name**|String|false|true|false|false|||||Y|Name||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/sector_core.go |
| code | **adaptor** | /adaptor/sector_impl.go_template |
| code | **api** | /application/sector_api.go |
| code | **dao** | /dao/sector_core.go |
| code | **datamodel** | /datamodel/sector_core.go |
| code | **menu** | /design/menu/sector.json |
| html | **list** | /Sector_List.html |
| html | **view** | /Sector_View.html |
| html | **edit** | /Sector_Edit.html |
| html | **new** | /Sector_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **28/06/2022** at **16:10:57**
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