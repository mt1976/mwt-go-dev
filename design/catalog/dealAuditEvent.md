# **DealAuditEvent** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**DealAuditEvent** (dealauditevent) |
|Endpoint 	    |**/DealAuditEvent...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-toolbox** ()
Friendly Name|**Deal Audit History**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/DealAuditEvent/DealAuditEventList) [Exportable]
* **View** (/DealAuditEvent/DealAuditEventView)











##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaDealAuditEvent**
SQL Table Key | **InternalId**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**DealRefNo**|String|true|true|false|false|||||Y|DealRefNo||false|false|false|text||
|**EventIndex**|Int|true|true|false|false|||||Y|EventIndex|0|false|false|false|text||
|**CommonRefNo**|String|false|true|false|false|||||Y|CommonRefNo||false|false|false|text||
|**Timestamp**|Time|false|true|false|false|||||Y|Timestamp||false|false|false|text||
|**UTCTimestamp**|String|false|true|false|false|||||Y|UTCTimestamp||false|false|false|text||
|**EventType**|String|false|true|false|false|||||Y|EventType||false|false|false|text||
|**Status**|String|false|true|false|false|||||Y|Status||false|false|false|text||
|**LimitOrderStatus**|String|false|true|false|false|||||Y|LimitOrderStatus||false|false|false|text||
|**Usr**|String|false|true|false|false|||||Y|Usr||false|false|false|text||
|**DealingInterface**|String|false|true|false|false|||||Y|DealingInterface||false|false|false|text||
|**SourceIP**|String|false|true|false|false|||||Y|SourceIP||false|false|false|text||
|**MessageID**|String|false|true|false|false|||||Y|MessageID||false|false|false|text||
|**Details**|String|false|true|false|false|||||Y|Details||false|false|false|text||
|**InternalId**|Int|true|true|false|false|||||Y|InternalId|0|false|false|false|text||
|**InternalDeleted**|Time|false|true|false|false|||||Y|InternalDeleted||false|false|false|text||
|**UpdatedTransactionId**|String|false|true|false|false|||||Y|UpdatedTransactionId||false|false|false|text||
|**UpdatedUserId**|String|false|true|false|false|||||Y|UpdatedUserId||false|false|false|text||
|**UpdatedDateTime**|Time|false|true|false|false|||||Y|UpdatedDateTime||false|false|false|text||
|**DeletedTransactionId**|String|false|true|false|false|||||Y|DeletedTransactionId||false|false|false|text||
|**DeletedUserId**|String|false|true|false|false|||||Y|DeletedUserId||false|false|false|text||
|**ChangeType**|String|false|true|false|false|||||Y|ChangeType||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/dealAuditEvent_core.go |
| code | **adaptor** | /adaptor/dealAuditEvent_impl.go_template |
| code | **dao** | /dao/dealAuditEvent_core.go |
| code | **datamodel** | /datamodel/dealAuditEvent_core.go |
| code | **menu** | /design/menu/dealAuditEvent.json |
| html | **list** | /DealAuditEvent_List.html |
| html | **view** | /DealAuditEvent_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **26/06/2022** at **18:48:27**
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