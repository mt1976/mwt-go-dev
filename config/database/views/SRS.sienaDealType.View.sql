USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDealType]    Script Date: 04/12/2021 18:16:55 ******/
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
