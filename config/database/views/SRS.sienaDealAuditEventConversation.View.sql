USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDealAuditEventConversation]    Script Date: 04/12/2021 18:16:55 ******/
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
