USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDealConversationLog]    Script Date: 13/10/2021 18:25:58 ******/
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
