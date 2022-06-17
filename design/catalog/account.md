# **Account** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**Account** (account) |
|Endpoint 	    |**/Account...** [^1]|
|Endpoint Query |**AccountNo**|
|REST API|**/API/Account/**|
Glyph|**fas fa-landmark** ()
Friendly Name|**Account**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/Account/AccountList) [Exportable]
* **View** (/Account/AccountView)











##  Provides
 * Lookup (SienaReference AccountName)
 * Reverse Lookup (CashBalance)





##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaAccount**
SQL Table Key | **SienaReference**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal | Display | Mask |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: | -- | -- |
|**SienaReference**|String|true|true|false|false|||||Y|SienaReference||false|false|false|text||
|**CustomerSienaView**|String|false|true|false|false|||||Y|CustomerSienaView||false|false|false|text||
|**SienaCommonRef**|String|true|true|false|false|||||Y|SienaCommonRef||false|false|false|text||
|**Status**|String|false|true|false|false|||||Y|Status||false|false|false|text||
|**StartDate**|Time|false|true|false|false|||||Y|StartDate||false|false|false|text||
|**MaturityDate**|Time|false|true|false|false|||||Y|MaturityDate||false|false|false|text||
|**ContractNumber**|String|false|true|false|false|||||Y|ContractNumber||false|false|false|text||
|**ExternalReference**|String|false|true|false|false|||||Y|ExternalReference||false|false|false|text||
|**CCY**|String|false|true|false|false|OL|Currency|CCY|Name|Y|CCY||true|false|false|text||
|**Book**|String|false|true|false|false|OL|Book|Book|FullName|N|Book||false|false|false|text||
|**MandatedUser**|String|false|true|false|false|||||Y|MandatedUser||false|false|false|text||
|**BackOfficeNotes**|String|false|true|false|false|||||Y|BackOfficeNotes||false|false|false|text||
|**CashBalance**|Float|false|true|false|false|||||Y|CashBalance|0.00|false|false|false|text||
|**AccountNumber**|String|false|true|false|false|||||Y|AccountNumber||false|false|false|text||
|**AccountName**|String|false|true|false|false|||||Y|AccountName||false|false|false|text||
|**LedgerBalance**|Float|false|true|false|false|||||Y|LedgerBalance|0.00|false|false|false|text||
|**Portfolio**|String|false|true|false|false|OL|Portfolio|Portfolio|Description1|N|Portfolio||false|false|false|text||
|**AgreementId**|Int|false|true|false|false|||||Y|AgreementId|0|false|false|false|text||
|**BackOfficeRefNo**|String|false|true|false|false|||||Y|BackOfficeRefNo||false|false|false|text||
|**ISIN**|String|false|true|false|false|||||Y|ISIN||false|false|false|text||
|**UTI**|String|false|true|false|false|||||Y|UTI||false|false|false|text||
|**CCYName**|String|false|true|false|false|||||Y|CCYName||false|false|false|text||
|**BookName**|String|false|true|false|false|||||Y|BookName||false|false|false|text||
|**PortfolioName**|String|false|true|false|false|||||Y|PortfolioName||false|false|false|text||
|**Centre**|String|false|true|false|false|OL|Centre|Centre|Name|N|Centre||false|false|false|text||
|**DealTypeKey**|String|false|true|false|false|||||Y|DealTypeKey||false|false|false|text||
|**DealTypeShortName**|String|false|true|false|false|||||Y|DealTypeShortName||false|false|false|text||
|**InternalId**|Int|false|true|false|false|||||Y|InternalId|0|false|false|false|text||
|**InternalDeleted**|Time|false|true|false|false|||||Y|InternalDeleted||false|false|false|text||
|**UpdatedTransactionId**|String|false|true|false|false|||||Y|UpdatedTransactionId||false|false|false|text||
|**UpdatedUserId**|String|false|true|false|false|||||Y|UpdatedUserId||false|false|false|text||
|**UpdatedDateTime**|Time|false|true|false|false|||||Y|UpdatedDateTime||false|false|false|text||
|**DeletedTransactionId**|String|false|true|false|false|||||Y|DeletedTransactionId||false|false|false|text||
|**DeletedUserId**|String|false|true|false|false|||||Y|DeletedUserId||false|false|false|text||
|**ChangeType**|String|false|true|false|false|||||Y|ChangeType||false|false|false|text||
|**CCYDp**|Int|false|true|false|false|||||Y|CCYDp|0|false|false|false|text||
|**CompID**|String|false|true|false|false|||||Y|CompID||false|false|false|text||
|**Firm**|String|false|true|false|false|OL|Firm|Firm|FullName|N|Firm||false|false|false|text||
|**DealType**|String|false|true|false|false|||||Y|DealType||false|false|false|text||
|**FullDealType**|String|false|true|false|false|||||Y|FullDealType||false|false|false|text||
|**DealingInterface**|String|false|true|false|false|||||Y|DealingInterface||false|false|false|text||
|**DealtAmount**|Float|false|true|false|false|||||Y|DealtAmount|0.00|false|false|false|text||
|**ParentContractNumber**|String|false|true|false|false|||||Y|ParentContractNumber||false|false|false|text||
|**InterestFrequency**|String|false|true|false|false|||||Y|InterestFrequency||false|false|false|text||
|**InterestAction**|String|false|true|false|false|||||Y|InterestAction||false|false|false|text||
|**InterestStrategy**|String|false|true|false|false|||||Y|InterestStrategy||false|false|false|text||
|**InterestBasis**|String|false|true|false|false|||||Y|InterestBasis||false|false|false|text||
|**SienaDealer**|String|false|true|false|false|||||Y|SienaDealer||false|false|false|text||
|**DealOwner**|String|false|true|false|false|||||Y|DealOwner||false|false|false|text||
|**OriginUser**|String|false|true|false|false|||||Y|OriginUser||false|false|false|text||
|**EditedByUser**|String|false|true|false|false|||||Y|EditedByUser||false|false|false|text||
|**DealOwnerMnemonic**|String|false|true|false|false|||||Y|DealOwnerMnemonic||false|false|false|text||
|**UTCOriginTime**|String|false|true|false|false|||||Y|UTCOriginTime||false|false|false|text||
|**UTCUpdateTime**|String|false|true|false|false|||||Y|UTCUpdateTime||false|false|false|text||
|**CustomerStatementNotes**|String|false|true|false|false|||||Y|customerStatementNotes||false|false|false|text||
|**NotesMargin**|Float|false|true|false|false|||||Y|NotesMargin|0.00|false|false|false|text||
|**RequestedBy**|String|false|true|false|false|||||Y|RequestedBy||false|false|false|text||
|**EditReason**|String|false|true|false|false|||||Y|EditReason||false|false|false|text||
|**EditOtherReason**|String|false|true|false|false|||||Y|EditOtherReason||false|false|false|text||
|**NoticeDays**|Int|false|true|false|false|||||Y|NoticeDays|0|false|false|false|text||
|**DebitFrequency**|String|false|true|false|false|||||Y|DebitFrequency||false|false|false|text||
|**CreditFrequency**|String|false|true|false|false|||||Y|CreditFrequency||false|false|false|text||
|**EURAmount**|Float|false|true|false|false|||||Y|EURAmount|0.00|false|false|false|text||
|**EUROtherAmount**|Float|false|true|false|false|||||Y|EUROtherAmount|0.00|false|false|false|text||
|**PaymentSystemSienaView**|String|false|true|false|false|||||Y|PaymentSystemSienaView||false|false|false|text||
|**PaymentSystemExternalView**|String|false|true|false|false|||||Y|PaymentSystemExternalView||false|false|false|text||
|**DealtCA**|String|false|false|true|false|||||N|||false|true|false|text||
|**AgainstCA**|String|false|false|true|false|||||N|||false|true|false|text||
|**LedgerCA**|String|false|false|true|false|||||N|||false|true|false|text||
|**CashBalanceCA**|String|false|false|true|false|||||N|||false|true|false|text||


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/account_core.go |
| code | **adaptor** | /adaptor/account_impl.go_template |
| code | **api** | /application/account_api.go |
| code | **dao** | /dao/account_core.go |
| code | **datamodel** | /datamodel/account_core.go |
| code | **menu** | /design/menu/account.json |
| html | **list** | /Account_List.html |
| html | **view** | /Account_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **17/06/2022** at **18:38:05**
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