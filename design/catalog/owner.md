# **Owner** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Owner** (owner) |
|Endpoint 	    |**/Owner...** [^1]|
|Endpoint Query |**Owner**|
|REST API|**/API/Owner/**|
Glyph|**fas fa-industry** ()
Friendly Name|**Owner**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Owner/OwnerList) [Exportable]
* **View** (/Owner/OwnerView)











##  Provides
 * Lookup (UserName FullName)
 * Reverse Lookup (FullName)





##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaOwner**
SQL Table Key | **UserName**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**UserName**|String|true|true|false|false|||||Y|UserName||false|false|false|text||
|**FullName**|String|false|true|false|false|||||Y|FullName||false|false|false|text||
|**Type**|String|false|true|false|false|||||Y|Type||false|false|false|text||
|**TradingEntity**|String|false|true|false|false|||||Y|TradingEntity||false|false|false|text||
|**DefaultEnterBook**|String|false|true|false|false|||||Y|DefaultEnterBook||false|false|false|text||
|**EmailAddress**|String|false|true|false|false|||||Y|EmailAddress||false|false|false|text||
|**Enabled**|String|false|true|false|false|||||Y|Enabled||false|false|false|text||
|**ExternalUserIds**|String|false|true|false|false|||||Y|ExternalUserIds||false|false|false|text||
|**Language**|String|false|true|false|false|||||Y|Language||false|false|false|text||
|**LocalCurrency**|String|false|true|false|false|||||Y|LocalCurrency||false|false|false|text||
|**Role**|String|false|true|false|false|||||Y|Role||false|false|false|text||
|**TelephoneNumber**|String|false|true|false|false|||||Y|TelephoneNumber||false|false|false|text||
|**TokenId**|String|false|true|false|false|||||Y|TokenId||false|false|false|text||
|**Entity**|String|false|true|false|false|||||Y|Entity||false|false|false|text||
|**UserCode**|String|false|true|false|false|||||Y|UserCode||false|false|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/owner_core.go |
| code | **api** | /application/owner_api.go |
| code | **dao** | /dao/owner_core.go |
| code | **datamodel** | /datamodel/owner_core.go |
| code | **menu** | /design/menu/owner.json |
| html | **list** | /Owner_List.html |
| html | **view** | /Owner_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **26/06/2022** at **18:48:30**
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