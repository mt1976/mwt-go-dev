USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBI_CounterpartyPerSector]    Script Date: 04/12/2021 18:17:03 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [{{!SQL.SCHEMA}}].[sienaBI_CounterpartyPerSector]
AS
SELECT        COUNT(InternalId) AS Count, SectorCode
FROM            {{!SQL.SOURCE}}.Counterparty
GROUP BY InternalDeleted, SectorCode
HAVING        (InternalDeleted IS NULL)

GO
