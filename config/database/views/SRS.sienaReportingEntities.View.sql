USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaReportingEntities]    Script Date: 04/12/2021 18:16:57 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaReportingEntities]
AS
SELECT        {{!SQL.SOURCE}}.SystemCounterparty.SystemKey, {{!SQL.SOURCE}}.SystemCounterparty.NameFirm, {{!SQL.SOURCE}}.SystemCounterparty.NameCentre, {{!SQL.SOURCE}}.SystemCounterpartyExtensions.FeeType, {{!SQL.SOURCE}}.SystemCounterpartyExtensions.NonMonetaryBenefits, 
                         {{!SQL.SOURCE}}.SystemCounterpartyExtensions.Rebate, {{!SQL.SOURCE}}.SystemCounterpartyExtensions.TradingMode, {{!SQL.SOURCE}}.SystemCounterpartyExtensions.TradingPlatform, {{!SQL.SOURCE}}.SystemCounterpartyExtensions.TradingSystem, 
                         {{!SQL.SOURCE}}.SystemCounterpartyExtensions.WebsiteURL
FROM            {{!SQL.SOURCE}}.SystemCounterparty LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.SystemCounterpartyExtensions ON {{!SQL.SOURCE}}.SystemCounterparty.SystemKey = {{!SQL.SOURCE}}.SystemCounterpartyExtensions.SystemKey
WHERE        ({{!SQL.SOURCE}}.SystemCounterparty.InternalDeleted IS NULL)
GO
