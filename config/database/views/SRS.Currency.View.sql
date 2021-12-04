USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[Currency]    Script Date: 04/12/2021 18:17:02 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [{{!SQL.SCHEMA}}].[Currency]
AS
SELECT        Code
FROM            {{!SQL.SOURCE}}.Currency
WHERE        (InternalDeleted IS NULL) AND (Active = 1)

GO
