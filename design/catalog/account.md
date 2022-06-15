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
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**SienaReference**|String|true|true|false|false|||||Y|SienaReference||false|false|false|
|**CustomerSienaView**|String|false|true|false|false|||||Y|CustomerSienaView||false|false|false|
|**SienaCommonRef**|String|true|true|false|false|||||Y|SienaCommonRef||false|false|false|
|**Status**|String|false|true|false|false|||||Y|Status||false|false|false|
|**StartDate**|Time|false|true|false|false|||||Y|StartDate||false|false|false|
|**MaturityDate**|Time|false|true|false|false|||||Y|MaturityDate||false|false|false|
|**ContractNumber**|String|false|true|false|false|||||Y|ContractNumber||false|false|false|
|**ExternalReference**|String|false|true|false|false|||||Y|ExternalReference||false|false|false|
|**CCY**|String|false|true|false|false|OL|Currency|CCY|Name|Y|CCY||true|false|false|
|**Book**|String|false|true|false|false|OL|Book|Book|FullName|N|Book||false|false|false|
|**MandatedUser**|String|false|true|false|false|||||Y|MandatedUser||false|false|false|
|**BackOfficeNotes**|String|false|true|false|false|||||Y|BackOfficeNotes||false|false|false|
|**CashBalance**|Float|false|true|false|false|||||Y|CashBalance|0.00|false|false|false|
|**AccountNumber**|String|false|true|false|false|||||Y|AccountNumber||false|false|false|
|**AccountName**|String|false|true|false|false|||||Y|AccountName||false|false|false|
|**LedgerBalance**|Float|false|true|false|false|||||Y|LedgerBalance|0.00|false|false|false|
|**Portfolio**|String|false|true|false|false|OL|Portfolio|Portfolio|Description1|N|Portfolio||false|false|false|
|**AgreementId**|Int|false|true|false|false|||||Y|AgreementId|0|false|false|false|
|**BackOfficeRefNo**|String|false|true|false|false|||||Y|BackOfficeRefNo||false|false|false|
|**ISIN**|String|false|true|false|false|||||Y|ISIN||false|false|false|
|**UTI**|String|false|true|false|false|||||Y|UTI||false|false|false|
|**CCYName**|String|false|true|false|false|||||Y|CCYName||false|false|false|
|**BookName**|String|false|true|false|false|||||Y|BookName||false|false|false|
|**PortfolioName**|String|false|true|false|false|||||Y|PortfolioName||false|false|false|
|**Centre**|String|false|true|false|false|OL|Centre|Centre|Name|N|Centre||false|false|false|
|**DealTypeKey**|String|false|true|false|false|||||Y|DealTypeKey||false|false|false|
|**DealTypeShortName**|String|false|true|false|false|||||Y|DealTypeShortName||false|false|false|
|**InternalId**|Int|false|true|false|false|||||Y|InternalId|0|false|false|false|
|**InternalDeleted**|Time|false|true|false|false|||||Y|InternalDeleted||false|false|false|
|**UpdatedTransactionId**|String|false|true|false|false|||||Y|UpdatedTransactionId||false|false|false|
|**UpdatedUserId**|String|false|true|false|false|||||Y|UpdatedUserId||false|false|false|
|**UpdatedDateTime**|Time|false|true|false|false|||||Y|UpdatedDateTime||false|false|false|
|**DeletedTransactionId**|String|false|true|false|false|||||Y|DeletedTransactionId||false|false|false|
|**DeletedUserId**|String|false|true|false|false|||||Y|DeletedUserId||false|false|false|
|**ChangeType**|String|false|true|false|false|||||Y|ChangeType||false|false|false|
|**CCYDp**|Int|false|true|false|false|||||Y|CCYDp|0|false|false|false|
|**CompID**|String|false|true|false|false|||||Y|CompID||false|false|false|
|**Firm**|String|false|true|false|false|OL|Firm|Firm|FullName|N|Firm||false|false|false|
|**DealType**|String|false|true|false|false|||||Y|DealType||false|false|false|
|**FullDealType**|String|false|true|false|false|||||Y|FullDealType||false|false|false|
|**DealingInterface**|String|false|true|false|false|||||Y|DealingInterface||false|false|false|
|**DealtAmount**|Float|false|true|false|false|||||Y|DealtAmount|0.00|false|false|false|
|**ParentContractNumber**|String|false|true|false|false|||||Y|ParentContractNumber||false|false|false|
|**InterestFrequency**|String|false|true|false|false|||||Y|InterestFrequency||false|false|false|
|**InterestAction**|String|false|true|false|false|||||Y|InterestAction||false|false|false|
|**InterestStrategy**|String|false|true|false|false|||||Y|InterestStrategy||false|false|false|
|**InterestBasis**|String|false|true|false|false|||||Y|InterestBasis||false|false|false|
|**SienaDealer**|String|false|true|false|false|||||Y|SienaDealer||false|false|false|
|**DealOwner**|String|false|true|false|false|||||Y|DealOwner||false|false|false|
|**OriginUser**|String|false|true|false|false|||||Y|OriginUser||false|false|false|
|**EditedByUser**|String|false|true|false|false|||||Y|EditedByUser||false|false|false|
|**DealOwnerMnemonic**|String|false|true|false|false|||||Y|DealOwnerMnemonic||false|false|false|
|**UTCOriginTime**|String|false|true|false|false|||||Y|UTCOriginTime||false|false|false|
|**UTCUpdateTime**|String|false|true|false|false|||||Y|UTCUpdateTime||false|false|false|
|**CustomerStatementNotes**|String|false|true|false|false|||||Y|customerStatementNotes||false|false|false|
|**NotesMargin**|Float|false|true|false|false|||||Y|NotesMargin|0.00|false|false|false|
|**RequestedBy**|String|false|true|false|false|||||Y|RequestedBy||false|false|false|
|**EditReason**|String|false|true|false|false|||||Y|EditReason||false|false|false|
|**EditOtherReason**|String|false|true|false|false|||||Y|EditOtherReason||false|false|false|
|**NoticeDays**|Int|false|true|false|false|||||Y|NoticeDays|0|false|false|false|
|**DebitFrequency**|String|false|true|false|false|||||Y|DebitFrequency||false|false|false|
|**CreditFrequency**|String|false|true|false|false|||||Y|CreditFrequency||false|false|false|
|**EURAmount**|Float|false|true|false|false|||||Y|EURAmount|0.00|false|false|false|
|**EUROtherAmount**|Float|false|true|false|false|||||Y|EUROtherAmount|0.00|false|false|false|
|**PaymentSystemSienaView**|String|false|true|false|false|||||Y|PaymentSystemSienaView||false|false|false|
|**PaymentSystemExternalView**|String|false|true|false|false|||||Y|PaymentSystemExternalView||false|false|false|
|**DealtCA**|String|false|false|true|false|||||N|||false|true|false|
|**AgainstCA**|String|false|false|true|false|||||N|||false|true|false|
|**LedgerCA**|String|false|false|true|false|||||N|||false|true|false|
|**CashBalanceCA**|String|false|false|true|false|||||N|||false|true|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/account_core.go |
| code | **adaptor** | /adaptor/account_impl.go_template |
| code | **api** | /application/account_api.go |
| code | **dao** | /dao/account_core.go |
| code | **datamodel** | /datamodel/account_core.go |
| code | **menu** | /design/menu/account.json |
| html | **list** | /html/Account_List.html |
| html | **view** | /html/Account_View.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **14/06/2022** at **21:31:47**
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