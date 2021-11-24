USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDeskSales]    Script Date: 24/11/2021 19:10:28 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaDeskSales]
AS
SELECT        Name, ReportDealsOver, ReportDealsOverCCY, AccountTransferCutOffTime, AccountTransferCutOffTimeTimeZone, AccountTransferCutOffTimeCutOffPeriod
FROM            {{!SQL.SOURCE}}.SalesTradingEntity
WHERE        (InternalDeleted IS NULL)
GO
