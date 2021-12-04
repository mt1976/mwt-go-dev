USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[Reason]    Script Date: 04/12/2021 18:17:03 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [{{!SQL.SCHEMA}}].[Reason]
AS
SELECT        Reason
FROM            {{!SQL.SOURCE}}.EditDealReason
WHERE        (InternalDeleted IS NULL)

GO
