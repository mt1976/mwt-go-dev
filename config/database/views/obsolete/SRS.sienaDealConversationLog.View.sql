USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDealConversationLog]    Script Date: 24/11/2021 19:10:28 ******/
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
