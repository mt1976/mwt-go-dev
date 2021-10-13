USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCountry]    Script Date: 13/10/2021 18:25:58 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaCountry]
AS
SELECT        Code, Name, ShortCode, EU_EEA
FROM            {{!SQL.SOURCE}}.Country
WHERE        (InternalDeleted IS NULL)
GO
