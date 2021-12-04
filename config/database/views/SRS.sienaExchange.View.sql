USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaExchange]    Script Date: 04/12/2021 18:16:56 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaExchange]
AS
SELECT        {{!SQL.SOURCE}}.Exchange.FullName, {{!SQL.SOURCE}}.Exchange.Broker, {{!SQL.SOURCE}}.Exchange.Country, {{!SQL.SOURCE}}.Exchange.Centre, {{!SQL.SOURCE}}.Exchange.CounterpartyFirm, {{!SQL.SOURCE}}.Exchange.CounterpartyCentre, {{!SQL.SOURCE}}.Exchange.LEI, {{!SQL.SOURCE}}.Country.Name AS CountryName, 
                         Centre_1.FullName AS CentreName, {{!SQL.SOURCE}}.Firm.FullName AS CounterpartyFirmName, {{!SQL.SOURCE}}.Centre.FullName AS CounterpartCentreName, {{!SQL.SOURCE}}.Broker.Name AS BrokerName, {{!SQL.SOURCE}}.Broker.FullName AS BrokerFullName
FROM            {{!SQL.SOURCE}}.Exchange LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Broker ON {{!SQL.SOURCE}}.Exchange.Broker = {{!SQL.SOURCE}}.Broker.Code LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Centre ON {{!SQL.SOURCE}}.Exchange.CounterpartyCentre = {{!SQL.SOURCE}}.Centre.ShortName LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Firm ON {{!SQL.SOURCE}}.Exchange.CounterpartyFirm = {{!SQL.SOURCE}}.Firm.FirmName LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Centre AS Centre_1 ON {{!SQL.SOURCE}}.Exchange.Centre = Centre_1.ShortName LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Country ON {{!SQL.SOURCE}}.Exchange.Country = {{!SQL.SOURCE}}.Country.Code
WHERE        ({{!SQL.SOURCE}}.Exchange.InternalDeleted IS NULL)
GO
