# **CounterpartyNotes** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**CounterpartyNotes** (counterpartynotes) |
|Endpoint 	    |**/CounterpartyNotes...** [^1]|
|Endpoint Query |**ID**|
Glyph|**far fa-sticky-note** ("")
Friendly Name|**Counterparty Notes**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/CounterpartyNotes/CounterpartyNotesList) [Exportable]
* **View** (/CounterpartyNotes/CounterpartyNotesView)











##  Provides
 * Lookup (NoteId Summary)






##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **contactManagerNote**
SQL Table Key | **noteId**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**NoteId**|Int|true|true|false|false|||||Y|noteId|0|false|false|false|text||
|**StreamId**|Int|false|true|false|false|OL|ContactStream|||N|streamId|0|false|false|false|text||
|**Summary**|String|false|true|false|false|||||Y|summary||false|false|false|text||
|**Details**|String|false|true|false|false|||||Y|details||false|false|false|text||
|**RecordState**|Int|false|true|false|false|LL|cmrecordstate|||N|recordState|0|false|false|false|text||
|**CreatedBy**|String|false|true|false|false|||||Y|createdBy||false|false|false|text||
|**CreatedDateTime**|Time|false|true|false|true|||||Y|createdDateTime||false|false|false|datetime||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/counterpartyNotes_core.go |
| code | **adaptor** | /adaptor/counterpartyNotes_impl.go_template |
| code | **dao** | /dao/counterpartyNotes_core.go |
| code | **datamodel** | /datamodel/counterpartyNotes_core.go |
| code | **menu** | /design/menu/counterpartyNotes.json |
| html | **list** | /CounterpartyNotes_List.html |
| html | **view** | /CounterpartyNotes_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **28/06/2022** at **16:10:48**
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