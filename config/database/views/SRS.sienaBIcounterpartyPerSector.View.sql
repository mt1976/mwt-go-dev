USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBIcounterpartyPerSector]    Script Date: 04/12/2021 18:16:50 ******/
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
