# **Template** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Template** (template) |
|Endpoint 	    |**/Template...** [^1]|
|Endpoint Query |**TemplateID**|
|REST API|**/API/Template/**|
Glyph|**fas fa-poo** ()
Friendly Name|**Template Contents**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Template/TemplateList) [Exportable]
* **View** (/Template/TemplateView)
* **Edit** (/Template/TemplateEdit)
* **Save** (/Template/TemplateSave)
* **New** (/Template/TemplateNew)
* **Delete** (/Template/TemplateDelete)







##  Provides


* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **template**
SQL Table Key | **ID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value|
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --|
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|
|**FIELD1**|String|false|true|false|false|LL|YN|||Y|FIELD1|N|
|**FIELD2**|String|false|true|false|false|OL|Firm|Firm|FullName|Y|FIELD2||
|**SYSCreated**|String|false|true|false|true|||||NH|_created||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||
|**ID**|String|true|true|false|false|||||Y|ID||
|**ExtraField**|String|false|false|true|false|||||Y|||
|**ExtraField2**|String|false|false|true|false|||||N||Hummous|
|**ExtraField3**|String|false|false|true|false|FL|Firm|Firm|FullName|Y|||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/template_core.go |
| code | **adaptor** | /adaptor/template_impl.go_template |
| code | **api** | /application/template_api.go |
| code | **dao** | /dao/template_core.go |
| code | **datamodel** | /datamodel/template_core.go |
| code | **job** | /jobs/template_core.go |
| code | **menu** | /design/menu/template.json |
| html | **list** | /html/Template_List.html |
| html | **view** | /html/Template_View.html |
| html | **edit** | /html/Template_Edit.html |
| html | **new** | /html/Template_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **12/06/2022** at **17:04:23**
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