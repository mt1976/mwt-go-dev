USE [RG3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCounterpartyCreditRating]    Script Date: 13/10/2021 18:25:58 ******/
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
