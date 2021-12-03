USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBroker]    Script Date: 13/10/2021 18:25:58 ******/
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
