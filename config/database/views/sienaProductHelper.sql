USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaProductHelper]    Script Date: 13/10/2021 18:25:58 ******/
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
