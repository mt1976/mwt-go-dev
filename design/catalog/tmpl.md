# **Tmpl** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Tmpl** (tmpl) |
|Endpoint 	    |**/Tmpl...** [^1]|
|Endpoint Query |**TemplateID**|
|REST API|**/API/Tmpl/**|
Glyph|**fas fa-poo** ()
Friendly Name|**Template Contents**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Tmpl/TmplList) [Exportable]
* **View** (/Tmpl/TmplView)
* **Edit** (/Tmpl/TmplEdit)
* **Save** (/Tmpl/TmplSave)
* **New** (/Tmpl/TmplNew)
* **Delete** (/Tmpl/TmplDelete)







##  Provides


* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **template**
SQL Table Key | **ID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**FIELD1**|String|false|true|false|false|LL|YN|||Y|FIELD1|N|false|false|false|text||
|**FIELD2**|String|false|true|false|false|OL|Firm|FirmName|FullName|Y|FIELD2||false|false|false|text||
|**SYSCreated**|String|false|true|false|true|||||H|_created||false|false|true|text||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|text||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|text||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|text||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|text||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|text||
|**ID**|String|true|true|false|false|||||Y|ID||false|false|false|text||
|**ExtraField**|String|false|false|true|false|||||Y|||false|true|false|text||
|**ExtraField2**|String|false|false|true|false|||||Y||Hummous|false|false|false|text||
|**ExtraField3**|String|false|false|true|false|FL|Firm|Firm|FullName|Y|||false|true|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/tmpl_core.go |
| code | **adaptor** | /adaptor/tmpl_impl.go_template |
| code | **api** | /application/tmpl_api.go |
| code | **dao** | /dao/tmpl_core.go |
| code | **datamodel** | /datamodel/tmpl_core.go |
| code | **job** | /jobs/tmpl_core.go |
| code | **menu** | /design/menu/tmpl.json |
| html | **list** | /Tmpl_List.html |
| html | **view** | /Tmpl_View.html |
| html | **edit** | /Tmpl_Edit.html |
| html | **new** | /Tmpl_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **17/06/2022** at **18:38:14**
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