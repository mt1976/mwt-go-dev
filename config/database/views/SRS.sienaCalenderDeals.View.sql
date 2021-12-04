USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCalenderDeals]    Script Date: 04/12/2021 18:16:52 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaCalenderDeals]
AS
SELECT        SienaReference AS ID, 'Deal' AS Type, TradeDate AS Date, TradeTime AS Time, { fn CONCAT('Deal ', ContractNumber) } AS ShortName, { fn CONCAT({ fn CONCAT(CustomerSienaView, ' - ') }, FullDealType) } AS Description
FROM            {{!SQL.SOURCE}}.Deals
WHERE        (NOT ({ fn CONCAT('Deal ', ContractNumber) } IS NULL)) AND (NOT (CustomerSienaView IS NULL))
GO
