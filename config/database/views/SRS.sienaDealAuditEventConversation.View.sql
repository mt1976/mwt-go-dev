USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDealAuditEventConversation]    Script Date: 24/11/2021 19:10:28 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaDealAuditEventConversation]
AS
SELECT        TOP (100) PERCENT DealRefNo, EventIndex, Message
FROM            {{!SQL.SOURCE}}.DealAuditEventConversation
WHERE        (InternalDeleted IS NULL)
ORDER BY DealRefNo, EventIndex, MessageIndex
GO
