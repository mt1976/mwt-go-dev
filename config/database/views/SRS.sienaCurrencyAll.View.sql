USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCurrencyAll]    Script Date: 04/12/2021 18:16:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaCurrencyAll]
AS
SELECT        {{!SQL.SOURCE}}.Currency.Code, {{!SQL.SOURCE}}.Currency.Name, {{!SQL.SOURCE}}.Currency.AmountDp, {{!SQL.SOURCE}}.Currency.Country, {{!SQL.SOURCE}}.Country.Name AS CountryName, {{!SQL.SOURCE}}.Currency.IntBase, {{!SQL.SOURCE}}.Currency.KeydateBase, 
                         {{!SQL.SOURCE}}.Currency.InterestRateTolerance, {{!SQL.SOURCE}}.Currency.CheckPayTo, {{!SQL.SOURCE}}.Currency.LatinAmericanSettlement, {{!SQL.SOURCE}}.Currency.DefaultLayOffBookKey, {{!SQL.SOURCE}}.Currency.CutOffTimeCutOffTime, {{!SQL.SOURCE}}.Currency.CutOffTimeTimeZone, 
                         {{!SQL.SOURCE}}.Currency.CutOffTimeDerivedDataUTCOffset, {{!SQL.SOURCE}}.Currency.CutOffTimeDerivedDataHasDaylightSaving, {{!SQL.SOURCE}}.Currency.CutOffTimeDerivedDataDaylightStart, {{!SQL.SOURCE}}.Currency.CutOffTimeDerivedDataDaylightEnd, 
                         {{!SQL.SOURCE}}.Currency.DealerInterventionQuoteTimeout, {{!SQL.SOURCE}}.Currency.CutOffTimeCutOffPeriod, {{!SQL.SOURCE}}.Currency.StripRateFutureExchangeCode, {{!SQL.SOURCE}}.Currency.StripRateFutureCurrencyContractCurrencyIsoCode, 
                         {{!SQL.SOURCE}}.Currency.StripRateFutureCurrencyContractFutureContractCode, {{!SQL.SOURCE}}.Currency.OvernightFundingSpreadBid, {{!SQL.SOURCE}}.Currency.OvernightFundingSpreadOffer
FROM            {{!SQL.SOURCE}}.Currency INNER JOIN
                         {{!SQL.SOURCE}}.Country ON {{!SQL.SOURCE}}.Currency.Country = {{!SQL.SOURCE}}.Country.Code
WHERE        ({{!SQL.SOURCE}}.Currency.InternalDeleted IS NULL) 
GO
