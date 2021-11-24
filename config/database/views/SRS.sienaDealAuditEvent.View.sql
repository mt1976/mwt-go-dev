USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDealAuditEvent]    Script Date: 24/11/2021 19:10:28 ******/
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
