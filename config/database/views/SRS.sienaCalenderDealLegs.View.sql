USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCalenderDealLegs]    Script Date: 04/12/2021 18:16:51 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaCalenderDealLegs]
AS
SELECT        {{!SQL.SOURCE}}.DealLegsInfo.InternalId AS ID, 'Transaction' AS Type, {{!SQL.SOURCE}}.DealLegsInfo.Narrative AS Description, {{!SQL.SOURCE}}.DealLegsInfo.StartInterestDate AS Date, { fn CONCAT({ fn CONCAT({ fn CONCAT('Deal ', {{!SQL.SOURCE}}.Deals.SienaReference) }, '-') 
                         }, {{!SQL.SOURCE}}.Deals.FullDealType) } AS ShortName, ' ' AS Time
FROM            {{!SQL.SOURCE}}.Currency LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Deals ON {{!SQL.SOURCE}}.Currency.Code = {{!SQL.SOURCE}}.Deals.DealtCcy RIGHT OUTER JOIN
                         {{!SQL.SOURCE}}.DealLegsInfo ON {{!SQL.SOURCE}}.Deals.SienaReference = {{!SQL.SOURCE}}.DealLegsInfo.SienaReference
WHERE        ({{!SQL.SOURCE}}.Deals.InternalDeleted IS NULL) AND ({{!SQL.SOURCE}}.DealLegsInfo.InternalDeleted IS NULL)
GO
