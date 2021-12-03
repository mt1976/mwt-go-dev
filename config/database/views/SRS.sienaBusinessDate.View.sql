USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBusinessDate]    Script Date: 24/11/2021 19:10:28 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaBusinessDate]
AS
SELECT        MAX(CurrentDate) AS Today
FROM            {{!SQL.SOURCE}}.BusinessDate
WHERE        (InternalDeleted IS NULL) AND (DealingStatus = 'Permitted')
GO
