USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDealTypeHelper]    Script Date: 24/11/2021 19:10:28 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaDealTypeHelper]
AS
SELECT        {{!SQL.SOURCE}}.DealType.DealTypeKey, {{!SQL.SOURCE}}.DealType.DealTypeShortName, {{!SQL.SOURCE}}.FundamentalDealTypeSwitches.NegotiableInstrumentType, {{!SQL.SOURCE}}.FundamentalDealTypeSwitches.ProductCode, {{!SQL.SOURCE}}.Product.Name AS ProductCodeName
FROM            {{!SQL.SOURCE}}.Product RIGHT OUTER JOIN
                         {{!SQL.SOURCE}}.FundamentalDealTypeSwitches ON {{!SQL.SOURCE}}.Product.Code = {{!SQL.SOURCE}}.FundamentalDealTypeSwitches.ProductCode RIGHT OUTER JOIN
                         {{!SQL.SOURCE}}.DealType ON {{!SQL.SOURCE}}.FundamentalDealTypeSwitches.DealTypeKey = {{!SQL.SOURCE}}.DealType.DealTypeKey
WHERE        ({{!SQL.SOURCE}}.DealType.InternalDeleted IS NULL) AND ({{!SQL.SOURCE}}.DealType.IsActive = 1) AND ({{!SQL.SOURCE}}.FundamentalDealTypeSwitches.Internal = 0)
GO
