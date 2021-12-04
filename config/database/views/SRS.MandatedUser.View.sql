USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[MandatedUser]    Script Date: 04/12/2021 18:17:02 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [{{!SQL.SCHEMA}}].[MandatedUser]
AS
SELECT        MandatedUserKeyUserName
FROM            {{!SQL.SOURCE}}.MandatedUser
WHERE        (InternalDeleted IS NULL) AND (Active = 1)

GO
