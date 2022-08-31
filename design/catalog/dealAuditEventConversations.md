# **DealAuditEventConversations** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**DealAuditEventConversations** (dealauditeventconversations) |
|Endpoint 	    |**/DealAuditEventConversations...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-toolbox** ()
Friendly Name|**Deal Audit Conversation**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/DealAuditEventConversations/DealAuditEventConversationsList) [Exportable]
* **View** (/DealAuditEventConversations/DealAuditEventConversationsView)











##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaDealAuditEventConversation**
SQL Table Key | **InternalId**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**DealRefNo**|String|true|true|false|false|||||Y|DealRefNo||false|false|false|text||
|**EventIndex**|Int|true|true|false|false|||||Y|EventIndex|0|false|false|false|text||
|**MessageIndex**|Int|true|true|false|false|||||Y|MessageIndex|0|false|false|false|text||
|**Message**|String|false|true|false|false|||||Y|Message||false|false|false|text||
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
| code | **application** | /application/dealAuditEventConversations_core.go |
| code | **adaptor** | /adaptor/dealAuditEventConversations_impl.go_template |
| code | **dao** | /dao/dealAuditEventConversations_core.go |
| code | **datamodel** | /datamodel/dealAuditEventConversations_core.go |
| code | **menu** | /design/menu/dealAuditEventConversations.json |
| html | **list** | /DealAuditEventConversations_List.html |
| html | **view** | /DealAuditEventConversations_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **28/06/2022** at **16:10:50**
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