USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCounterpartyType]    Script Date: 04/12/2021 18:17:00 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaCounterpartyType] AS
SELECT DISTINCT (CustomerType) FROM [ReportingDb_sal_prd_demo_sys-3].[RG].[Counterparty]

GO
