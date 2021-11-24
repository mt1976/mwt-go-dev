USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCountry]    Script Date: 24/11/2021 19:10:28 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaCountry]
AS
SELECT        Code, Name, ShortCode, EU_EEA, HolidaysWeekend
FROM            {{!SQL.SOURCE}}.Country
WHERE        (InternalDeleted IS NULL)
GO
