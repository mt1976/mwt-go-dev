USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDealConversationLog]    Script Date: 04/12/2021 18:16:55 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaDealConversationLog]
AS
SELECT        TOP (100) PERCENT SienaReference, Status, MessageType, ContractNumber, AckReference, NewTX, LegNo, Summary, BusinessDate, TXNo, ExternalSystem, MessageLogReference
FROM            {{!SQL.SOURCE}}.DealConversationLog
WHERE        (InternalDeleted IS NULL)
ORDER BY SienaReference, LegNo
GO
