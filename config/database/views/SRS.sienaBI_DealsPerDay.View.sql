USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBI_DealsPerDay]    Script Date: 24/11/2021 19:10:28 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaBI_DealsPerDay]
AS
SELECT        TOP (100) PERCENT TradeDate, COUNT(SienaReference) AS Count
FROM            {{!SQL.SOURCE}}.Deals
GROUP BY TradeDate, InternalDeleted
HAVING        (InternalDeleted IS NULL)
ORDER BY TradeDate
GO
