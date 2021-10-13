USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDeskSales]    Script Date: 13/10/2021 18:25:58 ******/
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
