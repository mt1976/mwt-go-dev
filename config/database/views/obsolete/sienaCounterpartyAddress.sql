USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCounterpartyAddress]    Script Date: 13/10/2021 18:25:58 ******/
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
