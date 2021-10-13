USE [SRS-2]
GO
/****** Object:  View [{{!SQL.DB}}].[{{!SQL.SCHEMA}}].[ratesDataView]    Script Date: 11/05/2021 14:31:24 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[niDataView] AS
SELECT        [{{!SQL.DB}}].dbo.lseGiltsDataStore.id, [{{!SQL.DB}}].dbo.lseGiltsDataStore.longName, [{{!SQL.DB}}].dbo.lseGiltsDataStore.isin, [{{!SQL.DB}}].dbo.lseGiltsDataStore.tidm, [{{!SQL.DB}}].dbo.lseGiltsDataStore.sedol, [{{!SQL.DB}}].dbo.lseGiltsDataStore.issueDate, [{{!SQL.DB}}].dbo.lseGiltsDataStore.maturityDate,
                         [{{!SQL.DB}}].dbo.lseGiltsDataStore.couponValue, [{{!SQL.DB}}].dbo.lseGiltsDataStore.couponType, [{{!SQL.DB}}].dbo.lseGiltsDataStore.segment, [{{!SQL.DB}}].dbo.lseGiltsDataStore.sector, [{{!SQL.DB}}].dbo.lseGiltsDataStore.codeConventionCalculateAccrual,
                         [{{!SQL.DB}}].dbo.lseGiltsDataStore.minimumDenomination, [{{!SQL.DB}}].dbo.lseGiltsDataStore.denominationCurrency, [{{!SQL.DB}}].dbo.lseGiltsDataStore.tradingCurrency, [{{!SQL.DB}}].dbo.lseGiltsDataStore.type, [{{!SQL.DB}}].dbo.lseGiltsDataStore.flatYield,
                         [{{!SQL.DB}}].dbo.lseGiltsDataStore.paymentCouponDate, [{{!SQL.DB}}].dbo.lseGiltsDataStore.periodOfCoupon, [{{!SQL.DB}}].dbo.lseGiltsDataStore.exCouponDate, [{{!SQL.DB}}].dbo.lseGiltsDataStore.dateOfIndexInflation, [{{!SQL.DB}}].dbo.lseGiltsDataStore.unitOfQuotation,
                         [{{!SQL.DB}}].dbo.lseGiltsDataStore.issuer, [{{!SQL.DB}}].dbo.lseGiltsDataStore.issueAmount, [{{!SQL.DB}}].dbo.lseGiltsDataStore.runningYield, [{{!SQL.DB}}].dbo.lseGiltsDataStore.LEI, [{{!SQL.DB}}].dbo.lseGiltsDataStore.CUSIP, [{{!SQL.DB}}].dbo.lseGiltsDataStore._created, [{{!SQL.DB}}].dbo.lseGiltsDataStore._who,
                         [{{!SQL.DB}}].dbo.lseGiltsDataStore._host, [{{!SQL.DB}}].dbo.lseGiltsDataStore._updated, [{{!SQL.DB}}].dbo.niSelectedStore.id AS isSelected
FROM            [{{!SQL.DB}}].dbo.lseGiltsDataStore LEFT OUTER JOIN
                         [{{!SQL.DB}}].dbo.niSelectedStore ON [{{!SQL.DB}}].dbo.lseGiltsDataStore.id = [{{!SQL.DB}}].dbo.niSelectedStore.id
GO
