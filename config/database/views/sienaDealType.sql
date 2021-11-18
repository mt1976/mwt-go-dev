USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDealType]    Script Date: 13/10/2021 18:25:58 ******/
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
