USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBIdealEventsPerDay]    Script Date: 24/11/2021 19:10:28 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaBIdealEventsPerDay]
AS
SELECT        TOP (100) PERCENT StartInterestDate, COUNT(InternalId) AS Count
FROM            {{!SQL.SOURCE}}.DealLegsInfo
WHERE        (InternalDeleted IS NULL)
GROUP BY StartInterestDate
HAVING        (NOT (StartInterestDate IS NULL))
ORDER BY StartInterestDate
GO
