USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCounterpartyExtensions]    Script Date: 13/10/2021 18:25:58 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaCounterpartyExtensions]
AS
SELECT        NameFirm, NameCentre, BICCode, ContactIndicator, CoverTrade, CustomerCategory, FSCSInclusive, FeeFactor, InactiveStatus, Indemnity, KnowYourCustomerStatus, LERLimitCarveOut, LastAmended, LastLogin, 
                         LossGivenDefault, MIC, ProtectedDepositor, RPTCurrency, RateTimeout, RateValidation, Registered, RegulatoryCategory, SecuredSettlement, SettlementLimitCarveOut, SortCode, Training, TrainingCode, TrainingReceived, 
                         Unencumbered, LEIExpiryDate, MIFIDReviewDate, GDPRReviewDate, DelegatedReporting, BOReconcile, MIFIDReportableDealsAllowed, SignedInvestmentAgreement, AppropriatenessAssessment, FinancialCounterparty, 
                         Collateralisation, PortfolioCode, ReconciliationLetterFrequency, DirectDealing
FROM            {{!SQL.SOURCE}}.CounterpartyExtensions
WHERE        (InternalDeleted IS NULL)
GO
