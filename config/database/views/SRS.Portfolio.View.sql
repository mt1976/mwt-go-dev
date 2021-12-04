USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[Portfolio]    Script Date: 04/12/2021 18:17:02 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [{{!SQL.SCHEMA}}].[Portfolio]
AS
SELECT        Description1
FROM            {{!SQL.SOURCE}}.Portfolio
WHERE        (InternalDeleted IS NULL)

GO
