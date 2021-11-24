USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBIcounterpartyPerSector]    Script Date: 24/11/2021 19:10:28 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaBIcounterpartyPerSector]
AS
SELECT        COUNT({{!SQL.SOURCE}}.Counterparty.InternalId) AS Count, {{!SQL.SOURCE}}.Sector.Name AS SectorCode
FROM            {{!SQL.SOURCE}}.Counterparty INNER JOIN
                         {{!SQL.SOURCE}}.Sector ON {{!SQL.SOURCE}}.Counterparty.SectorCode = {{!SQL.SOURCE}}.Sector.Code
GROUP BY {{!SQL.SOURCE}}.Counterparty.InternalDeleted, {{!SQL.SOURCE}}.Counterparty.SectorCode, {{!SQL.SOURCE}}.Sector.Name
HAVING        ({{!SQL.SOURCE}}.Counterparty.InternalDeleted IS NULL) AND (NOT ({{!SQL.SOURCE}}.Counterparty.SectorCode IS NULL))
GO
