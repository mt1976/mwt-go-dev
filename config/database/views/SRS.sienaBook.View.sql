USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBook]    Script Date: 24/11/2021 19:10:28 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaBook]
AS
SELECT        BookName, FullName, PLManage, PLTransfer, DerivePL, CostOfCarry, CostOfFunding, LotAllocationMethod, InternalId
FROM            {{!SQL.SOURCE}}.Book
WHERE        (InternalDeleted IS NULL)
GO
