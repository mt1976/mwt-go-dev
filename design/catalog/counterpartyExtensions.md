# **CounterpartyExtensions** - Object Definition
---
##  Information
|   |   |
|---|---|
|Object         |**CounterpartyExtensions** (counterpartyextensions) |
|Endpoint 	    |**/CounterpartyExtensions...** [^1]|
|Endpoint Query |**ID**|
Glyph|**fas fa-passport** ()
Friendly Name|**Counterparty Extension**|
|For Project    |github.com/mt1976/mwt-go-dev/|

##  Actions {#action-id}
* **List** (/CounterpartyExtensions/CounterpartyExtensionsList) [Exportable]
* **View** (/CounterpartyExtensions/CounterpartyExtensionsView)
* **Edit** (/CounterpartyExtensions/CounterpartyExtensionsEdit)
* **Save** (/CounterpartyExtensions/CounterpartyExtensionsSave)
* **New** (/CounterpartyExtensions/CounterpartyExtensionsNew)
* **Delete** (/CounterpartyExtensions/CounterpartyExtensionsDelete)







##  Provides







##  Data Source 
|   |   |
|---|---|
SQL Table Name       | **sienaCounterpartyExtensions**
SQL Table Key | **CompID**



##  Properties / Fields
| Field Name| Type | Mandatory | Core | Virtual | Overide | Lookup [^2]| Lookup Object      | Lookup Field Source         | Lookup Return Value                | Inputable [^3]|DB Column|Default Value| No Change | Callout | Internal |
| -- | --  | :--: | :--: | :--: |:--: |:--: |:--: |-- |-- |:--: |-- | --| :--: | :--: | :--: |
|**NameFirm**|String|true|true|false|false|||||Y|NameFirm||false|false|false|
|**NameCentre**|String|true|true|false|false|||||Y|NameCentre||false|false|false|
|**BICCode**|String|false|true|false|false|||||Y|BICCode||false|false|false|
|**ContactIndicator**|Bool|false|true|false|false|||||Y|ContactIndicator|True|false|false|false|
|**CoverTrade**|Bool|false|true|false|false|||||Y|CoverTrade|True|false|false|false|
|**CustomerCategory**|String|false|true|false|false|||||Y|CustomerCategory||false|false|false|
|**FSCSInclusive**|Bool|false|true|false|false|||||Y|FSCSInclusive|True|false|false|false|
|**FeeFactor**|Float|false|true|false|false|||||Y|FeeFactor|0.00|false|false|false|
|**InactiveStatus**|Bool|false|true|false|false|||||Y|InactiveStatus|True|false|false|false|
|**Indemnity**|Bool|false|true|false|false|||||Y|Indemnity|True|false|false|false|
|**KnowYourCustomerStatus**|Bool|false|true|false|false|||||Y|KnowYourCustomerStatus|True|false|false|false|
|**LERLimitCarveOut**|Bool|false|true|false|false|||||Y|LERLimitCarveOut|True|false|false|false|
|**LastAmended**|Time|false|true|false|false|||||Y|LastAmended||false|false|false|
|**LastLogin**|Time|false|true|false|false|||||Y|LastLogin||false|false|false|
|**LossGivenDefault**|Float|false|true|false|false|||||Y|LossGivenDefault|0.00|false|false|false|
|**MIC**|String|false|true|false|false|||||Y|MIC||false|false|false|
|**ProtectedDepositor**|Bool|false|true|false|false|||||Y|ProtectedDepositor|True|false|false|false|
|**RPTCurrency**|String|false|true|false|false|||||Y|RPTCurrency||false|false|false|
|**RateTimeout**|Int|false|true|false|false|||||Y|RateTimeout|0|false|false|false|
|**RateValidation**|String|false|true|false|false|||||Y|RateValidation||false|false|false|
|**Registered**|Bool|false|true|false|false|||||Y|Registered|True|false|false|false|
|**RegulatoryCategory**|String|false|true|false|false|||||Y|RegulatoryCategory||false|false|false|
|**SecuredSettlement**|Bool|false|true|false|false|||||Y|SecuredSettlement|True|false|false|false|
|**SettlementLimitCarveOut**|Bool|false|true|false|false|||||Y|SettlementLimitCarveOut|True|false|false|false|
|**SortCode**|String|false|true|false|false|||||Y|SortCode||false|false|false|
|**Training**|Bool|false|true|false|false|||||Y|Training|True|false|false|false|
|**TrainingCode**|String|false|true|false|false|||||Y|TrainingCode||false|false|false|
|**TrainingReceived**|Bool|false|true|false|false|||||Y|TrainingReceived|True|false|false|false|
|**Unencumbered**|Bool|false|true|false|false|||||Y|Unencumbered|True|false|false|false|
|**LEIExpiryDate**|Time|false|true|false|false|||||Y|LEIExpiryDate||false|false|false|
|**MIFIDReviewDate**|Time|false|true|false|false|||||Y|MIFIDReviewDate||false|false|false|
|**GDPRReviewDate**|Time|false|true|false|false|||||Y|GDPRReviewDate||false|false|false|
|**DelegatedReporting**|String|false|true|false|false|||||Y|DelegatedReporting||false|false|false|
|**BOReconcile**|Bool|false|true|false|false|||||Y|BOReconcile|True|false|false|false|
|**MIFIDReportableDealsAllowed**|Bool|false|true|false|false|||||Y|MIFIDReportableDealsAllowed|True|false|false|false|
|**SignedInvestmentAgreement**|Bool|false|true|false|false|||||Y|SignedInvestmentAgreement|True|false|false|false|
|**AppropriatenessAssessment**|Bool|false|true|false|false|||||Y|AppropriatenessAssessment|True|false|false|false|
|**FinancialCounterparty**|Bool|false|true|false|false|||||Y|FinancialCounterparty|True|false|false|false|
|**Collateralisation**|String|false|true|false|false|||||Y|Collateralisation||false|false|false|
|**PortfolioCode**|String|false|true|false|false|||||Y|PortfolioCode||false|false|false|
|**ReconciliationLetterFrequency**|String|false|true|false|false|||||Y|ReconciliationLetterFrequency||false|false|false|
|**DirectDealing**|Bool|false|true|false|false|||||Y|DirectDealing|True|false|false|false|
|**CompID**|String|true|true|false|false|||||Y|CompID||false|false|false|


##  Artifacts Generated
| Type | Artifact | Path|
| :--: | -- | -- |
| code | **application** | /application/counterpartyExtensions_core.go |
| code | **adaptor** | /adaptor/counterpartyExtensions_impl.go_template |
| code | **dao** | /dao/counterpartyExtensions_core.go |
| code | **datamodel** | /datamodel/counterpartyExtensions_core.go |
| code | **menu** | /design/menu/counterpartyExtensions.json |
| html | **list** | /html/CounterpartyExtensions_List.html |
| html | **view** | /html/CounterpartyExtensions_View.html |
| html | **edit** | /html/CounterpartyExtensions_Edit.html |
| html | **new** | /html/CounterpartyExtensions_New.html |


## Audit Information
|   |   |
|---|---|
Template Generator Version   | **delinquentDysprosium [r4-21.12.31]**
Date & Time		     | **14/06/2022** at **21:31:55**
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