USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[Broker]    Script Date: 04/12/2021 18:17:02 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [{{!SQL.SCHEMA}}].[Broker]
AS
SELECT        Code
FROM            {{!SQL.SOURCE}}.Broker
WHERE        (InternalDeleted IS NULL)

GO
