USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCentre]    Script Date: 24/11/2021 19:10:28 ******/
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
