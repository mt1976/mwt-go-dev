# **DealConversation** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**DealConversation** (dealconversation) |
|Endpoint 	    |**/DealConversation...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-exchange-alt** ()
Friendly Name|**Deal Conversation**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/DealConversation/DealConversationList) [Exportable]
* **View** (/DealConversation/DealConversationView)











##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaDealConversationLog**
SQL Table Key | **MessageLogReference**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SienaReference**|String|false|true|false|false|||||Y|SienaReference||false|false|false|text||
|**Status**|String|false|true|false|false|||||Y|Status||false|false|false|text||
|**MessageType**|String|false|true|false|false|||||Y|MessageType||false|false|false|text||
|**ContractNumber**|String|false|true|false|false|||||Y|ContractNumber||false|false|false|text||
|**AckReference**|String|false|true|false|false|||||Y|AckReference||false|false|false|text||
|**NewTX**|Bool|false|true|false|false|||||Y|NewTX|True|false|false|false|text||
|**LegNo**|Int|false|true|false|false|||||Y|LegNo|0|false|false|false|text||
|**Summary**|String|false|true|false|false|||||Y|Summary||false|false|false|text||
|**BusinessDate**|Time|false|true|false|false|||||Y|BusinessDate||false|false|false|text||
|**TXNo**|Int|false|true|false|false|||||Y|TXNo|0|false|false|false|text||
|**ExternalSystem**|String|false|true|false|false|||||Y|ExternalSystem||false|false|false|text||
|**MessageLogReference**|String|true|true|false|false|||||Y|MessageLogReference||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/dealConversation_core.go |
| code | **dao** | /dao/dealConversation_core.go |
| code | **datamodel** | /datamodel/dealConversation_core.go |
| code | **menu** | /design/menu/dealConversation.json |
| html | **list** | /DealConversation_List.html |
| html | **view** | /DealConversation_View.html |


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