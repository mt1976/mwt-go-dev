USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCounterpartyImportID]    Script Date: 13/10/2021 18:25:58 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaCounterpartyImportID]
AS
SELECT        TOP (100) PERCENT {{!SQL.SOURCE}}.CounterpartyImportID.KeyImportID, {{!SQL.SOURCE}}.CounterpartyImportID.CounterpartyFirm AS Firm, {{!SQL.SOURCE}}.CounterpartyImportID.CounterpartyCentre AS Centre, {{!SQL.SOURCE}}.Firm.FirmName, 
                         {{!SQL.SOURCE}}.Centre.FullName AS CentreName, {{!SQL.SOURCE}}.CounterpartyImportID.KeyOriginID
FROM            {{!SQL.SOURCE}}.CounterpartyImportID INNER JOIN
                         {{!SQL.SOURCE}}.Firm ON {{!SQL.SOURCE}}.CounterpartyImportID.CounterpartyFirm = {{!SQL.SOURCE}}.Firm.FirmName INNER JOIN
                         {{!SQL.SOURCE}}.Centre ON {{!SQL.SOURCE}}.CounterpartyImportID.CounterpartyCentre = {{!SQL.SOURCE}}.Centre.ShortName
WHERE        ({{!SQL.SOURCE}}.CounterpartyImportID.InternalDeleted IS NULL)
ORDER BY {{!SQL.SOURCE}}.CounterpartyImportID.KeyOriginID DESC
GO
