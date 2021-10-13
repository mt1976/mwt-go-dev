USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDealAuditEvent]    Script Date: 13/10/2021 18:25:58 ******/
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
