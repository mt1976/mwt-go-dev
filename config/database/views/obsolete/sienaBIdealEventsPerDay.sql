USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBIdealEventsPerDay]    Script Date: 13/10/2021 18:25:58 ******/
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
