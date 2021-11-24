USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCentre]    Script Date: 13/10/2021 18:25:58 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaCentre]
AS
SELECT        {{!SQL.SOURCE}}.Centre.ShortName AS Code, {{!SQL.SOURCE}}.Centre.FullName AS Name, {{!SQL.SOURCE}}.Centre.Country
FROM            {{!SQL.SOURCE}}.Centre 
WHERE        ({{!SQL.SOURCE}}.Centre.InternalDeleted IS NULL)
GO
