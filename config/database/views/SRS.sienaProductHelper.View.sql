USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaProductHelper]    Script Date: 24/11/2021 19:10:28 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaProductHelper]
AS
SELECT        Code, Name
FROM            {{!SQL.SOURCE}}.Product
WHERE        (InternalDeleted IS NULL)
GO
