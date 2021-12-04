USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCounterpartyCreditRating]    Script Date: 24/11/2021 19:10:28 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaCounterpartyCreditRating]
AS
SELECT        NameFirm, NameCentre, CreditRatingUsage, CreditRatingAgency, CreditRatingName
FROM            {{!SQL.SOURCE}}.CounterpartyCreditRatings
WHERE        (InternalDeleted IS NULL)
GO
