# **CMNotes** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**CMNotes** (cmnotes) |
|Endpoint 	    |**/CMNotes...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-poo** ("")
Friendly Name|**Customer Note**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/CMNotes/CMNotesList) [Exportable]
* **View** (/CMNotes/CMNotesView)











##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **contactManagerNote**
SQL Table Key | **noteId**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**NoteId**|Int|true|true|false|false|||||Y|noteId|0|false|false|false|
|**StreamId**|Int|false|true|false|false|||||Y|streamId|0|false|false|false|
|**Summary**|String|false|true|false|false|||||Y|summary||false|false|false|
|**Details**|String|false|true|false|false|||||Y|details||false|false|false|
|**RecordState**|Int|false|true|false|false|||||Y|recordState|0|false|false|false|
|**CreatedBy**|String|false|true|false|false|||||Y|createdBy||false|false|false|
|**CreatedDateTime**|Time|false|true|false|false|||||Y|createdDateTime||false|false|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/cMNotes_core.go |
| code | **adaptor** | /adaptor/cMNotes_impl.go_template |
| code | **dao** | /dao/cMNotes_core.go |
| code | **datamodel** | /datamodel/cMNotes_core.go |
| code | **menu** | /design/menu/cMNotes.json |
| html | **list** | /html/CMNotes_List.html |
| html | **view** | /html/CMNotes_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **14/06/2022** at **21:31:52**
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