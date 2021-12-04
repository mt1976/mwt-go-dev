USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCountry]    Script Date: 04/12/2021 18:16:54 ******/
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
