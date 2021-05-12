USE [SRS-2]
GO
/****** Object:  View [{{!SQL.DB}}].[{{!SQL.SCHEMA}}].[credentialsView]    Script Date: 11/05/2021 14:31:24 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[credentialsView] AS SELECT * FROM [{{!SQL.SOURCE}}].dbo.credentialsStore cs

GO
