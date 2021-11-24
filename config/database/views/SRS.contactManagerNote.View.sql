USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[contactManagerNote]    Script Date: 24/11/2021 20:38:01 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[contactManagerNote]
AS
SELECT        noteId, streamId, summary, details, recordState, createdBy, createdDateTime
FROM            CM.Note
GO