USE [SRS-2]
GO
/****** Object:  View [{{!SQL.DB}}].[{{!SQL.SCHEMA}}].[ratesDataView]    Script Date: 11/05/2021 14:31:24 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[niDataView] AS
SELECT        [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.id, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.longName, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.isin, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.tidm, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.sedol, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.issueDate, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.maturityDate,
                         [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.couponValue, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.couponType, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.segment, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.sector, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.codeConventionCalculateAccrual,
                         [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.minimumDenomination, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.denominationCurrency, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.tradingCurrency, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.type, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.flatYield,
                         [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.paymentCouponDate, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.periodOfCoupon, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.exCouponDate, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.dateOfIndexInflation, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.unitOfQuotation,
                         [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.issuer, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.issueAmount, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.runningYield, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.LEI, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.CUSIP, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore._created, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore._who,
                         [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore._host, [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore._updated, [{{!SQL.SOURCE}}].dbo.niSelectedStore.id AS isSelected
FROM            [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore LEFT OUTER JOIN
                         [{{!SQL.SOURCE}}].dbo.niSelectedStore ON [{{!SQL.SOURCE}}].dbo.lseGiltsDataStore.id = [{{!SQL.SOURCE}}].dbo.niSelectedStore.id
GO
