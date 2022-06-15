# **Mandate** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Mandate** (mandate) |
|Endpoint 	    |**/Mandate...** [^1]|
|Endpoint Query |**Mandate**|
|REST API|**/API/Mandate/**|
Glyph|**fas fa-city** ()
Friendly Name|**Mandate**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Mandate/MandateList) [Exportable]
* **View** (/Mandate/MandateView)
* **Edit** (/Mandate/MandateEdit)
* **Save** (/Mandate/MandateSave)
* **New** (/Mandate/MandateNew)








##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaMandatedUser**
SQL Table Key | **CompID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**MandatedUserKeyCounterpartyFirm**|String|false|true|false|false|OL|Firm|MandatedUserKeyCounterpartyFirm|FullName|N|MandatedUserKeyCounterpartyFirm||true|false|false|
|**MandatedUserKeyCounterpartyCentre**|String|false|true|false|false|OL|Centre|MandatedUserKeyCounterpartyCentre|Name|N|MandatedUserKeyCounterpartyCentre||true|false|false|
|**MandatedUserKeyUserName**|String|true|true|false|false|||||Y|MandatedUserKeyUserName||false|false|false|
|**TelephoneNumber**|String|false|true|false|false|||||Y|TelephoneNumber||false|false|false|
|**EmailAddress**|String|false|true|false|false|||||Y|EmailAddress||false|false|false|
|**Active**|Bool|false|true|false|false|LL|tf|||Y|Active|True|false|false|false|
|**FirstName**|String|false|true|false|false|||||Y|FirstName||false|false|false|
|**Surname**|String|false|true|false|false|||||Y|Surname||false|false|false|
|**DateOfBirth**|Time|true|true|false|false|||||Y|DateOfBirth||false|false|false|
|**Postcode**|String|false|true|false|false|||||Y|Postcode||false|false|false|
|**NationalIDNo**|String|false|true|false|false|||||Y|NationalIDNo||false|false|false|
|**PassportNo**|String|false|true|false|false|||||Y|PassportNo||false|false|false|
|**Country**|String|false|true|false|false|OL|Country|Country|Name|N|Country||false|false|false|
|**CountryName**|String|false|true|false|false|||||Y|CountryName||false|false|false|
|**FirmName**|String|false|true|false|false|||||Y|FirmName||false|false|false|
|**CentreName**|String|false|true|false|false|||||Y|CentreName||false|false|false|
|**Notify**|Bool|false|true|false|false|LL|tf|||Y|Notify|True|false|false|false|
|**SystemUser**|String|false|true|false|false|||||Y|SystemUser||false|false|false|
|**CompID**|String|true|true|false|false|||||Y|CompID||false|false|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/mandate_core.go |
| code | **adaptor** | /adaptor/mandate_impl.go_template |
| code | **api** | /application/mandate_api.go |
| code | **dao** | /dao/mandate_core.go |
| code | **datamodel** | /datamodel/mandate_core.go |
| code | **menu** | /design/menu/mandate.json |
| html | **list** | /html/Mandate_List.html |
| html | **view** | /html/Mandate_View.html |
| html | **edit** | /html/Mandate_Edit.html |
| html | **new** | /html/Mandate_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **14/06/2022** at **21:58:56**
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