USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBusinessDate]    Script Date: 13/10/2021 18:25:58 ******/
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
