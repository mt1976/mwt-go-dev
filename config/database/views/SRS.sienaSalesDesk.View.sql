USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaSalesDesk]    Script Date: 04/12/2021 18:17:08 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [{{!SQL.SCHEMA}}].[sienaSalesDesk]
AS
SELECT        Name, ReportDealsOver, ReportDealsOverCCY, AccountTransferCutOffTime, AccountTransferCutOffTimeTimeZone, AccountTransferCutOffTimeCutOffPeriod
FROM            {{!SQL.SOURCE}}.SalesTradingEntity
WHERE        (InternalDeleted IS NULL)

GO
