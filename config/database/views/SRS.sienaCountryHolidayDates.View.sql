USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCountryHolidayDates]    Script Date: 04/12/2021 18:17:05 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [{{!SQL.SCHEMA}}].[sienaCountryHolidayDates]
AS
SELECT        Code, InternalDeleted, DATEFROMPARTS(YEAR(GETDATE()), HMonth, HDay) AS HDate, Name, SettleOK
FROM            {{!SQL.SOURCE}}.CountryHolidaysAnnualDates
WHERE        (InternalDeleted IS NULL)

GO
