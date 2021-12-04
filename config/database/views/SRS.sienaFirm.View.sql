USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaFirm]    Script Date: 04/12/2021 18:16:56 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaFirm]
AS
SELECT        {{!SQL.SOURCE}}.Firm.FirmName, {{!SQL.SOURCE}}.Firm.FullName, {{!SQL.SOURCE}}.Firm.Country, {{!SQL.SOURCE}}.Firm.Sector
FROM            {{!SQL.SOURCE}}.Firm
WHERE        ({{!SQL.SOURCE}}.Firm.InternalDeleted IS NULL)
GO
