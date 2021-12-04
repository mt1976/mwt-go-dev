USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCounterpartyPayeeAll]    Script Date: 04/12/2021 18:17:05 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [{{!SQL.SCHEMA}}].[sienaCounterpartyPayeeAll]
AS
Select '{{!SQL.SOURCE}}.Payee' AS SourceTable
        ,KeyCounterpartyFirm
        ,KeyCounterpartyCentre
        ,KeyCurrency
,KeyName
,KeyNumber
,KeyDirection
,KeyType
,FullName
,Address
,PhoneNo
,Country
,Bic
,Iban
,AccountNo
,FedWireNo
,SortCode
,BankName
,BankPinCode
,BankAddress
,Reason
,BankSettlementAcct
,UpdatedUserId

From {{!SQL.SOURCE}}.Payee

UNION ALL

Select '{{!SQL.SOURCE}}.UnauthorisedPayee' AS SourceTable
        ,KeyCounterpartyFirm
        ,KeyCounterpartyCentre
              ,Currency
        ,KeyName
,'1'
,''
,''
,FullName
,Address
,PhoneNo
,Country
,Bic
,Iban
,AccountNo
,FedWireNo
,SortCode
,BankName
,BankPinCode
,BankAddress
,Reason
,BankSettlementAcct
,Usr

From {{!SQL.SOURCE}}.UnauthorisedPayee

GO
