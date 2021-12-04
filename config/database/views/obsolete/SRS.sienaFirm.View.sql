USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaFirm]    Script Date: 24/11/2021 19:10:28 ******/
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
