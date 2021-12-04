USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBook]    Script Date: 04/12/2021 18:16:50 ******/
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
