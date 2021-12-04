USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaUtil_SchemaCheck]    Script Date: 04/12/2021 18:17:25 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [{{!SQL.SCHEMA}}].[sienaUtil_SchemaCheck]
AS
SELECT        COUNT(*) AS [View Count], s.name AS [Schema]
FROM            sys.views AS t INNER JOIN
                         sys.schemas AS s ON s.schema_id = t.schema_id
GROUP BY s.name

GO
