# **ExternalMessage** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**ExternalMessage** (externalmessage) |
|Endpoint 	    |**/ExternalMessage...** [^1]|
|Endpoint Query |**MessageID**|
|REST API|**/API/ExternalMessage/**|
Glyph|**fas fa-mail-bulk** (text-info)
Friendly Name|**ExternalMessage**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/ExternalMessage/ExternalMessageList) [Exportable]
* **View** (/ExternalMessage/ExternalMessageView)
* **Edit** (/ExternalMessage/ExternalMessageEdit)
* **Save** (/ExternalMessage/ExternalMessageSave)

* **Delete** (/ExternalMessage/ExternalMessageDelete)







##  Provides


* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **externalMessageStore**
SQL Table Key | **MessageID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|
|**MessageID**|String|false|true|false|false|||||Y|messageID||false|false|false|
|**MessageFormat**|String|false|true|false|false|||||Y|messageFormat||false|false|false|
|**MessageDeliveredTo**|String|false|true|false|false|||||Y|messageDeliveredTo||false|false|false|
|**MessageBody**|String|false|true|false|false|||||Y|messageBody||false|false|false|
|**MessageFilename**|String|false|true|false|false|||||Y|messageFilename||false|false|false|
|**MessageLife**|String|false|true|false|false|||||Y|messageLife||false|false|false|
|**MessageDate**|String|false|true|false|false|||||Y|messageDate||false|false|false|
|**MessageTime**|String|false|true|false|false|||||Y|messageTime||false|false|false|
|**MessageTimeoutAction**|String|false|true|false|false|||||Y|messageTimeoutAction||false|false|false|
|**MessageACKNAK**|String|false|true|false|false|||||Y|messageACKNAK||false|false|false|
|**ResponseID**|String|false|true|false|false|||||Y|responseID||false|false|false|
|**ResponseFilename**|String|false|true|false|false|||||Y|responseFilename||false|false|false|
|**ResponseBody**|String|false|true|false|false|||||Y|responseBody||false|false|false|
|**ResponseDate**|String|false|true|false|false|||||Y|responseDate||false|false|false|
|**ResponseTime**|String|false|true|false|false|||||Y|responseTime||false|false|false|
|**ResponseAdditionalInfo**|String|false|true|false|false|||||Y|responseAdditionalInfo||false|false|false|
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|
|**MessageTimeout**|String|false|true|false|false|||||Y|messageTimeout||false|false|false|
|**MessageEmitted**|String|false|true|false|false|||||Y|messageEmitted||false|false|false|
|**ResponseRecieved**|String|false|true|false|false|||||Y|responseRecieved||false|false|false|
|**MessageClass**|String|false|true|false|false|||||Y|messageClass||false|false|false|
|**AppID**|String|false|true|false|false|||||Y|appID||false|false|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/externalMessage_core.go |
| code | **api** | /application/externalMessage_api.go |
| code | **dao** | /dao/externalMessage_core.go |
| code | **datamodel** | /datamodel/externalMessage_core.go |
| code | **menu** | /design/menu/externalMessage.json |
| html | **list** | /html/ExternalMessage_List.html |
| html | **view** | /html/ExternalMessage_View.html |
| html | **edit** | /html/ExternalMessage_Edit.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **14/06/2022** at **21:32:04**
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