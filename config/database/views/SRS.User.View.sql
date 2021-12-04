USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[User]    Script Date: 04/12/2021 18:17:03 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [{{!SQL.SCHEMA}}].[User]
AS
SELECT        UserName
FROM            {{!SQL.SOURCE}}.Usr
WHERE        (Type = 'CORE') AND (InternalDeleted IS NULL)

GO
