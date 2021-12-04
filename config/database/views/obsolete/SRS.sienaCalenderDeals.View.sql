USE [bnk6rgsys4]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCalenderDeals]    Script Date: 24/11/2021 19:10:28 ******/
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
