# **Credentials** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Credentials** (credentials) |
|Endpoint 	    |**/Credentials...** [^1]|
|Endpoint Query |**Id**|
Glyph|**fas fa-key** (text-danger)
Friendly Name|**Credential**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Credentials/CredentialsList) [Exportable]
* **View** (/Credentials/CredentialsView)
* **Edit** (/Credentials/CredentialsEdit)
* **Save** (/Credentials/CredentialsSave)
* **New** (/Credentials/CredentialsNew)
* **Delete** (/Credentials/CredentialsDelete)







##  Provides


* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **credentialsStore**
SQL Table Key | **id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|text||
|**Id**|String|true|true|false|false|||||Y|Id||false|false|false|text||
|**Username**|String|true|true|false|false|||||Y|Username||false|false|false|text||
|**Password**|String|false|true|false|false|||||Y|Password||false|false|false|text||
|**Firstname**|String|false|true|false|false|||||Y|Firstname||false|false|false|text||
|**Lastname**|String|false|true|false|false|||||Y|Lastname||false|false|false|text||
|**Knownas**|String|false|true|false|false|||||Y|Knownas||false|false|false|text||
|**Email**|String|false|true|false|false|||||Y|Email||false|false|false|text||
|**Issued**|String|false|true|false|false|||||Y|Issued||false|false|false|text||
|**Expiry**|String|false|true|false|false|||||Y|Expiry||false|false|false|text||
|**RoleType**|String|false|true|false|false|||||Y|RoleType||false|false|false|text||
|**Brand**|String|false|true|false|false|||||Y|Brand||false|false|false|text||
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|text||
|**SYSWho**|String|false|true|false|false|||||NH|_who||false|false|true|text||
|**SYSHost**|String|false|true|false|false|||||NH|_host||false|false|true|text||
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|text||
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|text||
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|text||
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|text||
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/credentials_core.go |
| code | **dao** | /dao/credentials_core.go |
| code | **datamodel** | /datamodel/credentials_core.go |
| code | **menu** | /design/menu/credentials.json |
| html | **list** | /Credentials_List.html |
| html | **view** | /Credentials_View.html |
| html | **edit** | /Credentials_Edit.html |
| html | **new** | /Credentials_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **17/06/2022** at **18:38:08**
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