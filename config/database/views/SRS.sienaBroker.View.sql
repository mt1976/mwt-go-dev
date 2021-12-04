USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBroker]    Script Date: 04/12/2021 18:16:51 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaBroker]
AS
SELECT        Code, Name, FullName, Contact, Address, LEI
FROM            {{!SQL.SOURCE}}.Broker
WHERE        (InternalDeleted IS NULL)
GO
