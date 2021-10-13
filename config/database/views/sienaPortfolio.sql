USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaPortfolio]    Script Date: 13/10/2021 18:25:58 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaPortfolio]
AS
SELECT        Code, Description1 AS Name
FROM            {{!SQL.SOURCE}}.Portfolio
WHERE        (InternalDeleted IS NULL)
GO
