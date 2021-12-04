USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[CurrencyPair]    Script Date: 04/12/2021 18:17:02 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [{{!SQL.SCHEMA}}].[CurrencyPair]
AS
SELECT        {{!SQL.SOURCE}}.CurrencyPair.CodeMajorCurrencyIsoCode, {{!SQL.SOURCE}}.CurrencyPair.CodeMinorCurrencyIsoCode, {{!SQL.SOURCE}}.CurrencyPair.Active, {{!SQL.SOURCE}}.CurrencyPair.SpotRerouteDataRerouteCurrency, {{!SQL.SOURCE}}.CurrencyPairRateRICCode.RICCode,
                         {{!SQL.SOURCE}}.CurrencyPairRateRICCode.Period, { fn CONCAT({{!SQL.SOURCE}}.CurrencyPair.CodeMajorCurrencyIsoCode, { fn CONCAT('\', {{!SQL.SOURCE}}.CurrencyPair.CodeMinorCurrencyIsoCode) }) } AS CODE1,
                         { fn CONCAT({{!SQL.SOURCE}}.CurrencyPair.CodeMajorCurrencyIsoCode, {{!SQL.SOURCE}}.CurrencyPair.CodeMinorCurrencyIsoCode) } AS CODE2
FROM            {{!SQL.SOURCE}}.CurrencyPair INNER JOIN
                         {{!SQL.SOURCE}}.CurrencyPairRateRICCode ON {{!SQL.SOURCE}}.CurrencyPair.CodeMajorCurrencyIsoCode = {{!SQL.SOURCE}}.CurrencyPairRateRICCode.CodeMajorCurrencyIsoCode AND
                         {{!SQL.SOURCE}}.CurrencyPair.CodeMinorCurrencyIsoCode = {{!SQL.SOURCE}}.CurrencyPairRateRICCode.CodeMinorCurrencyIsoCode
WHERE        ({{!SQL.SOURCE}}.CurrencyPairRateRICCode.RICCode <> '') AND ({{!SQL.SOURCE}}.CurrencyPair.SpotRerouteDataRerouteCurrency = '') AND ({{!SQL.SOURCE}}.CurrencyPairRateRICCode.Period = 'Spot') AND ({{!SQL.SOURCE}}.CurrencyPair.Active = 1)

GO
