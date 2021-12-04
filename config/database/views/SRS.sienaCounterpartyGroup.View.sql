USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCounterpartyGroup]    Script Date: 04/12/2021 18:16:53 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaCounterpartyGroup]
AS
SELECT        {{!SQL.SOURCE}}.CpartyGroup.Name, {{!SQL.SOURCE}}.CpartyGroup.CountryCode, {{!SQL.SOURCE}}.CpartyGroup.SuperGroup
FROM            {{!SQL.SOURCE}}.CpartyGroup 
WHERE        ({{!SQL.SOURCE}}.CpartyGroup.InternalDeleted IS NULL)
GO
