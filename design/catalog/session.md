# **Session** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Session** (session) |
|Endpoint 	    |**/Session...** [^1]|
|Endpoint Query |**SessionID**|
Glyph|**fas fa-history** (text-info)
Friendly Name|**Session**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Session/SessionList) 
* **View** (/Session/SessionView)

* **Save** (/Session/SessionSave)









##  Provides


* Auditing 




##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sessionStore**
SQL Table Key | **Id**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**SYSId**|Int|true|true|false|false|||||NH|_id|0|false|false|true|
|**Apptoken**|String|true|true|false|false|||||Y|Apptoken||false|false|false|
|**Createdate**|String|false|true|false|false|||||Y|Createdate||false|false|false|
|**Createtime**|String|false|true|false|false|||||Y|Createtime||false|false|false|
|**Uniqueid**|String|false|true|false|false|||||Y|Uniqueid||false|false|false|
|**Sessiontoken**|String|false|true|false|false|||||Y|Sessiontoken||false|false|false|
|**Username**|String|false|true|false|false|||||Y|Username||false|false|false|
|**Password**|String|false|true|false|false|||||Y|Password||false|false|false|
|**Userip**|String|false|true|false|false|||||Y|Userip||false|false|false|
|**Userhost**|String|false|true|false|false|||||Y|Userhost||false|false|false|
|**Appip**|String|false|true|false|false|||||Y|Appip||false|false|false|
|**Apphost**|String|false|true|false|false|||||Y|Apphost||false|false|false|
|**Issued**|String|false|true|false|false|||||Y|Issued||false|false|false|
|**Expiry**|String|false|true|false|false|||||Y|Expiry||false|false|false|
|**Expiryraw**|String|false|true|false|false|||||Y|Expiryraw||false|false|false|
|**Brand**|String|false|true|false|false|||||Y|Brand||false|false|false|
|**SYSCreated**|String|false|true|false|false|||||NH|_created||false|false|true|
|**SYSWho**|String|false|true|false|false|||||NH|_who||false|false|true|
|**SYSHost**|String|false|true|false|false|||||NH|_host||false|false|true|
|**SYSUpdated**|String|false|true|false|false|||||NH|_updated||false|false|true|
|**Id**|String|false|true|false|false|||||Y|Id||false|false|false|
|**Expires**|Time|false|true|false|false|||||Y|Expires||false|false|false|
|**SYSCreatedBy**|String|false|true|false|false|||||NH|_createdBy||false|false|true|
|**SYSCreatedHost**|String|false|true|false|false|||||NH|_createdHost||false|false|true|
|**SYSUpdatedBy**|String|false|true|false|false|||||NH|_updatedBy||false|false|true|
|**SYSUpdatedHost**|String|false|true|false|false|||||NH|_updatedHost||false|false|true|
|**SessionRole**|String|false|true|false|false|||||Y|SessionRole||false|false|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/session_core.go |
| code | **dao** | /dao/session_core.go |
| code | **datamodel** | /datamodel/session_core.go |
| code | **menu** | /design/menu/session.json |
| html | **list** | /html/Session_List.html |
| html | **view** | /html/Session_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **14/06/2022** at **21:32:09**
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