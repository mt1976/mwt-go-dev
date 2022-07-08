# **Payee** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Payee** (payee) |
|Endpoint 	    |**/Payee...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-wallet** ("")
Friendly Name|**Payee**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Payee/PayeeList) [Exportable]
* **View** (/Payee/PayeeView)











##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaCounterpartyPayee**
SQL Table Key | **KeyCounterpartyFirm**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SourceTable**|String|true|true|false|false|||||Y|SourceTable||false|false|false|text||
|**KeyCounterpartyFirm**|String|true|true|false|false|||||Y|KeyCounterpartyFirm||false|false|false|text||
|**KeyCounterpartyCentre**|String|true|true|false|false|||||Y|KeyCounterpartyCentre||false|false|false|text||
|**KeyCurrency**|String|false|true|false|false|||||Y|KeyCurrency||false|false|false|text||
|**KeyName**|String|true|true|false|false|||||Y|KeyName||false|false|false|text||
|**KeyNumber**|String|true|true|false|false|||||Y|KeyNumber||false|false|false|text||
|**KeyDirection**|String|true|true|false|false|||||Y|KeyDirection||false|false|false|text||
|**KeyType**|String|true|true|false|false|||||Y|KeyType||false|false|false|text||
|**FullName**|String|false|true|false|false|||||Y|FullName||false|false|false|text||
|**Address**|String|false|true|false|false|||||Y|Address||false|false|false|text||
|**PhoneNo**|String|false|true|false|false|||||Y|PhoneNo||false|false|false|text||
|**Country**|String|false|true|false|false|OL|Country|Country|Name|N|Country||false|false|false|||
|**Bic**|String|false|true|false|false|||||Y|Bic||false|false|false|text||
|**Iban**|String|false|true|false|false|||||Y|Iban||false|false|false|text||
|**AccountNo**|String|false|true|false|false|||||Y|AccountNo||false|false|false|text||
|**FedWireNo**|String|false|true|false|false|||||Y|FedWireNo||false|false|false|text||
|**SortCode**|String|false|true|false|false|||||Y|SortCode||false|false|false|text||
|**BankName**|String|false|true|false|false|||||Y|BankName||false|false|false|text||
|**BankPinCode**|String|false|true|false|false|||||Y|BankPinCode||false|false|false|text||
|**BankAddress**|String|false|true|false|false|||||Y|BankAddress||false|false|false|text||
|**Reason**|String|false|true|false|true|||||N|Reason||false|false|false|text||
|**BankSettlementAcct**|Bool|false|true|false|false|||||Y|BankSettlementAcct|True|false|false|false|text||
|**UpdatedUserId**|String|false|true|false|false|||||Y|UpdatedUserId||false|false|false|text||
|**Status**|String|false|false|true|false|||||N|||false|true|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/payee_core.go |
| code | **adaptor** | /adaptor/payee_impl.go_template |
| code | **dao** | /dao/payee_core.go |
| code | **datamodel** | /datamodel/payee_core.go |
| code | **menu** | /design/menu/payee.json |
| html | **list** | /Payee_List.html |
| html | **view** | /Payee_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **28/06/2022** at **16:10:55**
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