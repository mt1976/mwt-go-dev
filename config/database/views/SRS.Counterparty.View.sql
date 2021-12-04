USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[Counterparty]    Script Date: 04/12/2021 18:17:02 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO

CREATE VIEW [{{!SQL.SCHEMA}}].[Counterparty]
AS
SELECT        KeyImportID
FROM            {{!SQL.SOURCE}}.CounterpartyImportID
WHERE        (InternalDeleted IS NULL) AND (KeyOriginID = 'ExternalDealImporter')

GO
