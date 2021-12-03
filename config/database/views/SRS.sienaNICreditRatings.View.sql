USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaNICreditRatings]    Script Date: 24/11/2021 19:10:28 ******/
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
