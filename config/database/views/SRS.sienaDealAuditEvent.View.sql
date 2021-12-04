USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDealAuditEvent]    Script Date: 04/12/2021 18:16:54 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaDealAuditEvent]
AS
SELECT        TOP (100) PERCENT DealRefNo, EventIndex, Timestamp, EventType, Status, Usr, MessageID, Details
FROM            {{!SQL.SOURCE}}.DealAuditEvent
WHERE        (InternalDeleted IS NULL)
ORDER BY DealRefNo, EventIndex
GO
