# **ContactDeals** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**ContactDeals** (contactdeals) |
|Endpoint 	    |**/ContactDeals...** [^1]|
|Endpoint Query |**ID**|
Glyph|**far fa-comment** ("")
Friendly Name|**Conversation Deals**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/ContactDeals/ContactDealsList) [Exportable]
* **View** (/ContactDeals/ContactDealsView)











##  Provides
 * Lookup (NoteId RefNo)






##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **contactManagerDealRefNo**
SQL Table Key | **noteId**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**RefNo**|String|true|true|false|false|||||Y|refNo||false|false|false|text||
|**NoteId**|Int|false|true|false|false|||||Y|noteId|0|false|false|false|text||
|**RecordState**|Int|false|true|false|false|||||Y|recordState|0|false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/contactDeals_core.go |
| code | **adaptor** | /adaptor/contactDeals_impl.go_template |
| code | **dao** | /dao/contactDeals_core.go |
| code | **datamodel** | /datamodel/contactDeals_core.go |
| code | **menu** | /design/menu/contactDeals.json |
| html | **list** | /ContactDeals_List.html |
| html | **view** | /ContactDeals_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **24/06/2022** at **16:01:55**
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