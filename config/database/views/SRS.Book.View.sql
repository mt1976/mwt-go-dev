USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[Book]    Script Date: 04/12/2021 18:17:02 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [{{!SQL.SCHEMA}}].[Book]
AS
SELECT        BookName
FROM            {{!SQL.SOURCE}}.Book
WHERE        (InternalDeleted IS NULL)

GO
