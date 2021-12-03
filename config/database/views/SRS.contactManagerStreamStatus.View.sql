USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[contactManagerStreamStatus]    Script Date: 24/11/2021 20:38:01 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[contactManagerStreamStatus]
AS
SELECT        statusId, description, recordState
FROM            CM.StreamStatus
GO