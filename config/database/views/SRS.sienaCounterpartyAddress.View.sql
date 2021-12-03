USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCounterpartyAddress]    Script Date: 24/11/2021 19:10:28 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaCounterpartyAddress]
AS
SELECT        NameFirm, NameCentre, Address1, Address2, CityTown, County, Postcode
FROM            {{!SQL.SOURCE}}.CounterpartyAddress
WHERE        (InternalDeleted IS NULL)
GO
