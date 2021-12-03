USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaImportOrigins]    Script Date: 24/11/2021 19:10:28 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaImportOrigins]
AS
SELECT DISTINCT KeyOriginID
FROM            {{!SQL.SOURCE}}.CounterpartyImportID
GO
