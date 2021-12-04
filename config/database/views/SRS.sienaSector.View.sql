USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaSector]    Script Date: 04/12/2021 18:16:57 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaSector]
AS
SELECT        Code, Name
FROM            {{!SQL.SOURCE}}.Sector
WHERE        (InternalDeleted IS NULL)
GO
