USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBookedRatePayments]    Script Date: 04/12/2021 18:16:50 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaBookedRatePayments]
AS
SELECT        {{!SQL.SOURCE}}.Deals.SienaReference, {{!SQL.SOURCE}}.Deals.UTI, {{!SQL.SOURCE}}.Deals.FullDealType, {{!SQL.SOURCE}}.Deals.Broker, {{!SQL.SOURCE}}.Deals.Status, {{!SQL.SOURCE}}.Deals.TradeDate, {{!SQL.SOURCE}}.Deals.StartDate, {{!SQL.SOURCE}}.Deals.MaturityDate, {{!SQL.SOURCE}}.Deals.ContractNumber, 
                         {{!SQL.SOURCE}}.Deals.ExternalReference, {{!SQL.SOURCE}}.Deals.DealingInterface, {{!SQL.SOURCE}}.Deals.DealtAmount, {{!SQL.SOURCE}}.Deals.AgainstAmount, {{!SQL.SOURCE}}.Deals.DealtCcy, {{!SQL.SOURCE}}.Deals.AgainstCcy, {{!SQL.SOURCE}}.Deals.MajorCcy, {{!SQL.SOURCE}}.Deals.MinorCcy, {{!SQL.SOURCE}}.Deals.AllInRate, 
                         {{!SQL.SOURCE}}.Deals.Book, {{!SQL.SOURCE}}.Deals.SettleCcy, {{!SQL.SOURCE}}.Deals.Direction, {{!SQL.SOURCE}}.Deals.MandatedUser, {{!SQL.SOURCE}}.Deals.CustomerSienaView, {{!SQL.SOURCE}}.Deals.PayInstruction, {{!SQL.SOURCE}}.Deals.ReceiptInstruction, {{!SQL.SOURCE}}.Deals.VenueUTI, {{!SQL.SOURCE}}.Deals.Portfolio, 
                         {{!SQL.SOURCE}}.Deals.ISIN, {{!SQL.SOURCE}}.Deals.LEI, {{!SQL.SOURCE}}.Deals.AgreementId, {{!SQL.SOURCE}}.Deals.DealOwner, {{!SQL.SOURCE}}.Deals.OriginUser, {{!SQL.SOURCE}}.Deals.SienaCommonRef, {{!SQL.SOURCE}}.FundamentalDealType.DealTypeShortName, {{!SQL.SOURCE}}.Deals.FastPay, {{!SQL.SOURCE}}.Deals.PaymentFee, 
                         {{!SQL.SOURCE}}.Deals.PaymentFeePolicy, {{!SQL.SOURCE}}.Deals.PaymentReason, {{!SQL.SOURCE}}.Deals.BackOfficeRefNo
FROM            {{!SQL.SOURCE}}.FundamentalDealType RIGHT OUTER JOIN
                         {{!SQL.SOURCE}}.DealType ON {{!SQL.SOURCE}}.FundamentalDealType.DealTypeKey = {{!SQL.SOURCE}}.DealType.FundamentalDealTypeKey RIGHT OUTER JOIN
                         {{!SQL.SOURCE}}.Deals ON {{!SQL.SOURCE}}.DealType.DealTypeKey = {{!SQL.SOURCE}}.Deals.FullDealType
WHERE        ({{!SQL.SOURCE}}.Deals.InternalDeleted IS NULL) AND ({{!SQL.SOURCE}}.FundamentalDealType.DealTypeShortName = 'FPay')
GO
