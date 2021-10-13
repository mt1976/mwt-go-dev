USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBI_DealsPerDay]    Script Date: 13/10/2021 18:25:58 ******/
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
