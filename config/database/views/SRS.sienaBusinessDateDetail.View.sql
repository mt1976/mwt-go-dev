USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBusinessDateDetail]    Script Date: 04/12/2021 18:16:51 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaBusinessDateDetail]
AS
SELECT        SRS.sienaBusinessDate.Today, {{!SQL.SOURCE}}.BusinessDate.CurrentDate, {{!SQL.SOURCE}}.BusinessDate.PreviousBusinessDate, {{!SQL.SOURCE}}.BusinessDate.CurrentBusinessDate, {{!SQL.SOURCE}}.BusinessDate.NextBusinessDate, {{!SQL.SOURCE}}.BusinessDate.DealingStatus, 
                         {{!SQL.SOURCE}}.BusinessDate.RolloverType
FROM            {{!SQL.SOURCE}}.BusinessDate INNER JOIN
                         SRS.sienaBusinessDate ON {{!SQL.SOURCE}}.BusinessDate.CurrentDate = SRS.sienaBusinessDate.Today
GO
