# **SimulatorSWIFT** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**SimulatorSWIFT** (simulatorswift) |
|Endpoint 	    |**/SimulatorSWIFT...** [^1]|
|Endpoint Query |**ID**|
|REST API|**/API/SimulatorSWIFT/**|
Glyph|**fas fa-globe** (text-primary)
Friendly Name|**SWIFT Simulator**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/SimulatorSWIFT/SimulatorSWIFTList) [Exportable]
* **View** (/SimulatorSWIFT/SimulatorSWIFTView)
* **Edit** (/SimulatorSWIFT/SimulatorSWIFTEdit)
* **Save** (/SimulatorSWIFT/SimulatorSWIFTSave)
* **New** (/SimulatorSWIFT/SimulatorSWIFTNew)
* **Delete** (/SimulatorSWIFT/SimulatorSWIFTDelete)







##  Provides



* Monitoring 



##  Data Source 
|   |   |
|---|---|

Fetch|<ul><li>**Implement in Adaptor**</li><li> func SimulatorSWIFT_GetList_impl() (int, []dm.SimulatorSWIFT, error) {return 0, nil, nil}</li><li>func SimulatorSWIFT_GetByID_impl(id string) (int, dm.SimulatorSWIFT, error) {return 0, dm.SimulatorSWIFT{}, nil}</li></ul>
Store|<ul><li>**Implement in Adaptor**</li><li>func SimulatorSWIFT_NewID_impl(rec dm.SimulatorSWIFT) (string) { return rec.ID } </li><li>func SimulatorSWIFT_Delete_impl(id string) (error) {return nil}</li><li>func SimulatorSWIFT_Update_impl(id string,rec dm.SimulatorSWIFT, usr string) (error) {return nil}</li></ul>

##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**ID**|String|true|true|false|false|||||Y|ID||false|false|false|text||
|**FileName**|String|true|true|false|false|||||Y|FileName||false|false|false|text||
|**MessageRaw**|String|false|true|false|false|||||Y|MessageRaw||false|false|false|text||
|**MessageFmt**|String|false|true|false|true|||||N|MessageFmt||false|false|false|text||
|**Action**|String|false|true|false|false|||||Y|Action||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/simulatorSWIFT_core.go |
| code | **adaptor** | /adaptor/simulatorSWIFT_impl.go_template |
| code | **api** | /application/simulatorSWIFT_api.go |
| code | **dao** | /dao/simulatorSWIFT_core.go |
| code | **datamodel** | /datamodel/simulatorSWIFT_core.go |
| code | **menu** | /design/menu/simulatorSWIFT.json |
| html | **list** | /SimulatorSWIFT_List.html |
| html | **view** | /SimulatorSWIFT_View.html |
| html | **edit** | /SimulatorSWIFT_Edit.html |
| html | **new** | /SimulatorSWIFT_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **26/06/2022** at **18:48:33**
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