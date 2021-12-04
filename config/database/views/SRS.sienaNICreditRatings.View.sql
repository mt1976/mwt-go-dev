USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaNICreditRatings]    Script Date: 04/12/2021 18:16:57 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaNICreditRatings]
AS
SELECT        Name, Agency, Rating
FROM            {{!SQL.SOURCE}}.NegotiableInstrumentCreditRatings
WHERE        (InternalDeleted IS NULL) AND (NOT (Agency IS NULL)) AND (NOT (Rating IS NULL))
GO
