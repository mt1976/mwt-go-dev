USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaPortfolio]    Script Date: 04/12/2021 18:16:57 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaPortfolio]
AS
SELECT        *
FROM            {{!SQL.SOURCE}}.Portfolio
WHERE        (InternalDeleted IS NULL)
GO
