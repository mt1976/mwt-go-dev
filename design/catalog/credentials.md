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
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|
|**Id**|String|true|true|false|false|||||Y|Id||false|false|false|
|**Username**|String|true|true|false|false|||||Y|Username||false|false|false|
|**Password**|String|false|true|false|false|||||Y|Password||false|false|false|
|**Firstname**|String|false|true|false|false|||||Y|Firstname||false|false|false|
|**Lastname**|String|false|true|false|false|||||Y|Lastname||false|false|false|
|**Knownas**|String|false|true|false|false|||||Y|Knownas||false|false|false|
|**Email**|String|false|true|false|false|||||Y|Email||false|false|false|
|**Issued**|String|false|true|false|false|||||Y|Issued||false|false|false|
|**Expiry**|String|false|true|false|false|||||Y|Expiry||false|false|false|
|**RoleType**|String|false|true|false|false|||||Y|RoleType||false|false|false|
|**Brand**|String|false|true|false|false|||||Y|Brand||false|false|false|
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|
|**SYSWho**|String|false|true|false|false|||||NH|_who||false|false|true|
|**SYSHost**|String|false|true|false|false|||||NH|_host||false|false|true|
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/credentials_core.go |
| code | **dao** | /dao/credentials_core.go |
| code | **datamodel** | /datamodel/credentials_core.go |
| code | **menu** | /design/menu/credentials.json |
| html | **list** | /html/Credentials_List.html |
| html | **view** | /html/Credentials_View.html |
| html | **edit** | /html/Credentials_Edit.html |
| html | **new** | /html/Credentials_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **14/06/2022** at **21:32:00**
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