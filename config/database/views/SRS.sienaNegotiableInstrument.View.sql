USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaNegotiableInstrument]    Script Date: 04/12/2021 18:16:57 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaNegotiableInstrument]
AS
SELECT        {{!SQL.SOURCE}}.NegotiableInstrument.NIName, {{!SQL.SOURCE}}.NegotiableInstrument.Status, {{!SQL.SOURCE}}.NegotiableInstrument.Type, {{!SQL.SOURCE}}.NegotiableInstrument.CurrencyCode, {{!SQL.SOURCE}}.NegotiableInstrument.IssuerFirm, {{!SQL.SOURCE}}.NegotiableInstrument.IssuerCentre, 
                         {{!SQL.SOURCE}}.NegotiableInstrument.AcceptorFirm, {{!SQL.SOURCE}}.NegotiableInstrument.AcceptorCentre, {{!SQL.SOURCE}}.NegotiableInstrument.IssueDate, {{!SQL.SOURCE}}.NegotiableInstrument.MaturityDate, {{!SQL.SOURCE}}.NegotiableInstrument.RateType, 
                         {{!SQL.SOURCE}}.NegotiableInstrument.CouponRate, {{!SQL.SOURCE}}.NegotiableInstrument.InterestBasis, {{!SQL.SOURCE}}.NegotiableInstrument.InterestFrequency, {{!SQL.SOURCE}}.NegotiableInstrument.ClearingSystem, {{!SQL.SOURCE}}.NegotiableInstrument.ParentDealType, 
                         {{!SQL.SOURCE}}.NegotiableInstrument.SecurityID, {{!SQL.SOURCE}}.NegotiableInstrument.Encumbered, {{!SQL.SOURCE}}.NegotiableInstrument.Marketable, {{!SQL.SOURCE}}.NegotiableInstrument.Hypothecated, {{!SQL.SOURCE}}.NegotiableInstrument.RepoFactor, 
                         {{!SQL.SOURCE}}.NegotiableInstrument.LotSize, {{!SQL.SOURCE}}.Currency.Name AS CurrencyName, {{!SQL.SOURCE}}.Firm.FullName AS IssuerFirmName, {{!SQL.SOURCE}}.Centre.FullName AS IssuerCentreName
FROM            {{!SQL.SOURCE}}.NegotiableInstrument INNER JOIN
                         {{!SQL.SOURCE}}.Currency ON {{!SQL.SOURCE}}.NegotiableInstrument.CurrencyCode = {{!SQL.SOURCE}}.Currency.Code INNER JOIN
                         {{!SQL.SOURCE}}.Firm ON {{!SQL.SOURCE}}.NegotiableInstrument.IssuerFirm = {{!SQL.SOURCE}}.Firm.FirmName INNER JOIN
                         {{!SQL.SOURCE}}.Centre ON {{!SQL.SOURCE}}.NegotiableInstrument.IssuerCentre = {{!SQL.SOURCE}}.Centre.ShortName
WHERE        ({{!SQL.SOURCE}}.NegotiableInstrument.InternalDeleted IS NULL)
GO
