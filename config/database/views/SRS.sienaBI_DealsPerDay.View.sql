USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBI_DealsPerDay]    Script Date: 04/12/2021 18:16:49 ******/
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
