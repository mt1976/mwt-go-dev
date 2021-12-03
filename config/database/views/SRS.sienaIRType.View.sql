USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaIRType]    Script Date: 24/11/2021 19:10:28 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaIRType]
AS
SELECT        Code, Style, AllCurrencies, CurrencyIsoCode, RateCategory
FROM            {{!SQL.SOURCE}}.IRType
WHERE        (InternalDeleted IS NULL)
GO
