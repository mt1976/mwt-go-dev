USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[SystemDate]    Script Date: 04/12/2021 18:17:03 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [{{!SQL.SCHEMA}}].[SystemDate]
AS
SELECT        MAX(CurrentDate) AS Today
FROM            {{!SQL.SOURCE}}.BusinessDate
WHERE        (InternalDeleted IS NULL) AND (DealingStatus = 'Permitted')

GO
