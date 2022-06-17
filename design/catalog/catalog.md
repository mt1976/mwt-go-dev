# **Catalog** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Catalog** (catalog) |
|Endpoint 	    |**/Catalog...** [^1]|
|Endpoint Query |**ID**|
|REST API|**/API/Catalog/**|
Glyph|**fas fa-cubes** (text-info)
Friendly Name|**API Catalog**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Catalog/CatalogList) [Exportable]
* **View** (/Catalog/CatalogView)











##  Provides







##  Data Source 
|   |   |
|---|---|

Fetch|<ul><li>**Implement in Adaptor**</li><li> func Catalog_GetList_impl() (int, []dm.Catalog, error) {return 0, nil, nil}</li><li>func Catalog_GetByID_impl(id string) (int, dm.Catalog, error) {return 0, dm.Catalog{}, nil}</li></ul>
Store|<ul><li>**Implement in Adaptor**</li><li>func Catalog_NewID_impl(rec dm.Catalog) (string) { return rec.ID } </li><li>func Catalog_Delete_impl(id string) (error) {return nil}</li><li>func Catalog_Update_impl(id string,rec dm.Catalog, usr string) (error) {return nil}</li></ul>

##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**ID**|String|true|true|false|false|||||Y|ID||false|false|false|text||
|**Endpoint**|String|true|true|false|false|||||Y|Endpoint||false|false|false|text||
|**Descr**|String|true|true|false|false|||||Y|Descr||false|false|false|text||
|**Query**|String|true|true|false|false|||||Y|Query||false|false|false|text||
|**Source**|String|true|true|false|false|||||Y|Source||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/catalog_core.go |
| code | **adaptor** | /adaptor/catalog_impl.go_template |
| code | **api** | /application/catalog_api.go |
| code | **dao** | /dao/catalog_core.go |
| code | **datamodel** | /datamodel/catalog_core.go |
| code | **menu** | /design/menu/catalog.json |
| html | **list** | /Catalog_List.html |
| html | **view** | /Catalog_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **17/06/2022** at **18:38:06**
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