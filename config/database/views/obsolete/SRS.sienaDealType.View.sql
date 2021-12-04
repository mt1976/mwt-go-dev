USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDealType]    Script Date: 24/11/2021 19:10:28 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaDealType]
AS
SELECT        *
FROM            {{!SQL.SOURCE}}.DealType 
WHERE        ({{!SQL.SOURCE}}.DealType.InternalDeleted IS NULL)
GO
