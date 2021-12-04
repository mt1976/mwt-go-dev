USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDealTypeFundamentals]    Script Date: 24/11/2021 19:10:28 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaDealTypeFundamentals]
AS
SELECT        *
FROM            {{!SQL.SOURCE}}.FundamentalDealTypeSwitches 
WHERE        ({{!SQL.SOURCE}}.FundamentalDealTypeSwitches.InternalDeleted IS NULL)
GO
