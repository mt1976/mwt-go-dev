# **Broker** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Broker** (broker) |
|Endpoint 	    |**/Broker...** [^1]|
|Endpoint Query |**Code**|
|REST API|**/API/Broker/**|
Glyph|**fas fa-plug** ()
Friendly Name|**Broker**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Broker/BrokerList) [Exportable]
* **View** (/Broker/BrokerView)
* **Edit** (/Broker/BrokerEdit)
* **Save** (/Broker/BrokerSave)
* **New** (/Broker/BrokerNew)
* **Delete** (/Broker/BrokerDelete)







##  Provides

 * Reverse Lookup (Name)





##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaBroker**
SQL Table Key | **Code**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**Code**|String|true|true|false|false|||||Y|Code||false|false|false|
|**Name**|String|false|true|false|false|||||Y|Name||false|false|false|
|**FullName**|String|false|true|false|false|||||Y|FullName||false|false|false|
|**Contact**|String|false|true|false|false|||||Y|Contact||false|false|false|
|**Address**|String|false|true|false|false|||||Y|Address||false|false|false|
|**LEI**|String|false|true|false|false|||||Y|LEI||false|false|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/broker_core.go |
| code | **adaptor** | /adaptor/broker_impl.go_template |
| code | **api** | /application/broker_api.go |
| code | **dao** | /dao/broker_core.go |
| code | **datamodel** | /datamodel/broker_core.go |
| code | **menu** | /design/menu/broker.json |
| html | **list** | /html/Broker_List.html |
| html | **view** | /html/Broker_View.html |
| html | **edit** | /html/Broker_Edit.html |
| html | **new** | /html/Broker_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **14/06/2022** at **21:31:49**
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