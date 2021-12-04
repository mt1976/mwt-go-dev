USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDealingInterface]    Script Date: 04/12/2021 18:16:56 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaDealingInterface]
AS
SELECT        {{!SQL.SOURCE}}.DealingInterface.Name, {{!SQL.SOURCE}}.DealingInterface.AcceptReducedAmount, {{!SQL.SOURCE}}.DealingInterface.QuoteAsIndicative, {{!SQL.SOURCE}}.DealingInterface.RateTimeOut, {{!SQL.SOURCE}}.DealingInterface.PropagationDelay, {{!SQL.SOURCE}}.DealingInterface.CheckLiquidity, 
                         {{!SQL.SOURCE}}.DealingInterface.ChangeQuoteDirection, {{!SQL.SOURCE}}.DealingInterface.GenerateRejectedDeals, {{!SQL.SOURCE}}.DealingInterface.SpotUpdatesForForwardQuotes, {{!SQL.SOURCE}}.DealingInterface.SettlementInstructionStyle, 
                         {{!SQL.SOURCE}}.DealingInterfaceExtensions.CanRetractQuotes, {{!SQL.SOURCE}}.DealingInterfaceExtensions.CancelESPifNotPriced, {{!SQL.SOURCE}}.DealingInterfaceExtensions.CancelRFQSifNotPriced, {{!SQL.SOURCE}}.DealingInterfaceExtensions.CancelonDealingSuspended, 
                         {{!SQL.SOURCE}}.DealingInterfaceExtensions.CreditCheckedatDI, {{!SQL.SOURCE}}.DealingInterfaceExtensions.DuplicateCheckonExternalRef, {{!SQL.SOURCE}}.DealingInterfaceExtensions.LimitCheckQuote, 
                         {{!SQL.SOURCE}}.DealingInterfaceExtensions.LimitCheckonRFQDealSubmission, {{!SQL.SOURCE}}.DealingInterfaceExtensions.ListenonLimits, {{!SQL.SOURCE}}.DealingInterfaceExtensions.MarginStyle, {{!SQL.SOURCE}}.DealingInterfaceExtensions.UseRerouteDefinitionOnly, 
                         {{!SQL.SOURCE}}.DealingInterfaceExtensions.BypassConfirmation, {{!SQL.SOURCE}}.DealingInterfaceExtensions.DIOnAcceptance, {{!SQL.SOURCE}}.DealingInterfaceExtensions.IgnoreESPAmountRules
FROM            {{!SQL.SOURCE}}.DealingInterface INNER JOIN
                         {{!SQL.SOURCE}}.DealingInterfaceExtensions ON {{!SQL.SOURCE}}.DealingInterface.Name = {{!SQL.SOURCE}}.DealingInterfaceExtensions.Name
WHERE        ({{!SQL.SOURCE}}.DealingInterface.InternalDeleted IS NULL)
GO
