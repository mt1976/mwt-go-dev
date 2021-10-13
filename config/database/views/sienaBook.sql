USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBook]    Script Date: 13/10/2021 18:25:58 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaBook]
AS
SELECT        BookName, FullName
FROM            {{!SQL.SOURCE}}.Book
WHERE        (InternalDeleted IS NULL)
GO
