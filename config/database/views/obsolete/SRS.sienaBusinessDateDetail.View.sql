USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBusinessDateDetail]    Script Date: 24/11/2021 19:10:28 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaBusinessDateDetail]
AS
SELECT        {{!SQL.SCHEMA}}.sienaBusinessDate.Today, {{!SQL.SOURCE}}.BusinessDate.CurrentDate, {{!SQL.SOURCE}}.BusinessDate.PreviousBusinessDate, {{!SQL.SOURCE}}.BusinessDate.CurrentBusinessDate, {{!SQL.SOURCE}}.BusinessDate.NextBusinessDate, {{!SQL.SOURCE}}.BusinessDate.DealingStatus, 
                         {{!SQL.SOURCE}}.BusinessDate.RolloverType
FROM            {{!SQL.SOURCE}}.BusinessDate INNER JOIN
                         {{!SQL.SCHEMA}}.sienaBusinessDate ON {{!SQL.SOURCE}}.BusinessDate.CurrentDate = {{!SQL.SCHEMA}}.sienaBusinessDate.Today
GO
