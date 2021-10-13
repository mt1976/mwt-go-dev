USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaFirm]    Script Date: 13/10/2021 18:25:58 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaFirm]
AS
SELECT        {{!SQL.SOURCE}}.Firm.FirmName, {{!SQL.SOURCE}}.Firm.FullName, {{!SQL.SOURCE}}.Firm.Country, {{!SQL.SOURCE}}.Firm.Sector, {{!SQL.SOURCE}}.Sector.Name AS SectorName, {{!SQL.SOURCE}}.Country.Name AS CountryName
FROM            {{!SQL.SOURCE}}.Firm INNER JOIN
                         {{!SQL.SOURCE}}.Sector ON {{!SQL.SOURCE}}.Firm.Sector = {{!SQL.SOURCE}}.Sector.Code INNER JOIN
                         {{!SQL.SOURCE}}.Country ON {{!SQL.SOURCE}}.Firm.Country = {{!SQL.SOURCE}}.Country.Code
WHERE        ({{!SQL.SOURCE}}.Firm.InternalDeleted IS NULL)
GO
