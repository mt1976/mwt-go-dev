USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaTradeDesk]    Script Date: 04/12/2021 18:16:47 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [{{!SQL.SCHEMA}}].[sienaTradeDesk]
AS
SELECT        Name
FROM            {{!SQL.SOURCE}}.TraderTradingEntity
WHERE        (InternalDeleted IS NULL)

GO
