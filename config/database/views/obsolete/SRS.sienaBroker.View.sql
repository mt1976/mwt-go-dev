USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBroker]    Script Date: 24/11/2021 19:10:28 ******/
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
