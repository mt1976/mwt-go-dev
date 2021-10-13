USE [SRS-2]
GO
/****** Object:  View [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].[ratesDataView]    Script Date: 11/05/2021 14:31:24 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[niDataView] AS
SELECT        [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.id, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.longName, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.isin, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.tidm, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.sedol, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.issueDate, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.maturityDate,
                         [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.couponValue, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.couponType, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.segment, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.sector, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.codeConventionCalculateAccrual,
                         [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.minimumDenomination, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.denominationCurrency, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.tradingCurrency, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.type, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.flatYield,
                         [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.paymentCouponDate, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.periodOfCoupon, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.exCouponDate, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.dateOfIndexInflation, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.unitOfQuotation,
                         [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.issuer, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.issueAmount, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.runningYield, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.LEI, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.CUSIP, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore._created, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore._who,
                         [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore._host, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore._updated, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].niSelectedStore.id AS isSelected
FROM            [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore LEFT OUTER JOIN
                         [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].niSelectedStore ON [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].lseGiltsDataStore.id = [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].niSelectedStore.id
GO
