USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDealType]    Script Date: 13/10/2021 18:25:58 ******/
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
