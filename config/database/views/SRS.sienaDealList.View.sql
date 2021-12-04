USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDealList]    Script Date: 04/12/2021 18:16:55 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaDealList]
AS
SELECT        {{!SQL.SOURCE}}.Deals.SienaReference, {{!SQL.SOURCE}}.Deals.Status, {{!SQL.SOURCE}}.Deals.StartDate AS ValueDate, {{!SQL.SOURCE}}.Deals.MaturityDate, {{!SQL.SOURCE}}.Deals.ContractNumber, {{!SQL.SOURCE}}.Deals.ExternalReference, {{!SQL.SOURCE}}.Deals.Book, {{!SQL.SOURCE}}.Deals.MandatedUser, {{!SQL.SOURCE}}.Deals.Portfolio, 
                         {{!SQL.SOURCE}}.Deals.AgreementId, {{!SQL.SOURCE}}.Deals.BackOfficeRefNo, {{!SQL.SOURCE}}.Deals.ISIN, {{!SQL.SOURCE}}.Deals.UTI, {{!SQL.SOURCE}}.Book.FullName AS BookName, RIGHT({{!SQL.SOURCE}}.Deals.CustomerSienaView, 3) AS Centre, RTRIM(LTRIM(LEFT({{!SQL.SOURCE}}.Deals.CustomerSienaView, 10))) 
                         AS Firm, {{!SQL.SOURCE}}.FundamentalDealType.DealTypeShortName, {{!SQL.SOURCE}}.Deals.FullDealType, {{!SQL.SOURCE}}.Deals.TradeDate, {{!SQL.SOURCE}}.Deals.DealtCcy, {{!SQL.SOURCE}}.Deals.DealtAmount, {{!SQL.SOURCE}}.Deals.AgainstAmount, {{!SQL.SOURCE}}.Deals.AgainstCcy, {{!SQL.SOURCE}}.Deals.AllInRate, 
                         {{!SQL.SOURCE}}.Deals.MktRate, {{!SQL.SOURCE}}.Deals.SettleCcy, {{!SQL.SOURCE}}.Deals.Direction, {{!SQL.SOURCE}}.Deals.NpvRate, {{!SQL.SOURCE}}.Deals.OriginUser, {{!SQL.SOURCE}}.Deals.PayInstruction, {{!SQL.SOURCE}}.Deals.ReceiptInstruction, {{!SQL.SOURCE}}.Deals.NIName, { fn CONCAT({{!SQL.SOURCE}}.Deals.DealtCcy, 
                         {{!SQL.SOURCE}}.Deals.AgainstCcy) } AS CCYPair, ISNULL(NULLIF ({{!SQL.SOURCE}}.Deals.NIName, NULL), { fn CONCAT({{!SQL.SOURCE}}.Deals.DealtCcy, {{!SQL.SOURCE}}.Deals.AgainstCcy) }) AS Instrument, {{!SQL.SOURCE}}.Portfolio.Description1 AS PortfolioName, 
                         {{!SQL.SOURCE}}.DealRevaluation.Date AS RVDate, {{!SQL.SOURCE}}.DealRevaluation.MarkToMarket AS RVMTM, {{!SQL.SOURCE}}.Deals.CounterBook, Book_1.FullName AS CounterBookName, ISNULL(NULLIF ({{!SQL.SOURCE}}.Deals.CounterBook, NULL), 
                         {{!SQL.SOURCE}}.Deals.CustomerSienaView) AS Party, ISNULL(NULLIF (Book_1.FullName, NULL), {{!SQL.SOURCE}}.Deals.CustomerSienaView) AS PartyName, RIGHT({{!SQL.SOURCE}}.Deals.CustomerSienaView, 3) AS NameCentre, 
                         RTRIM(LTRIM(LEFT({{!SQL.SOURCE}}.Deals.CustomerSienaView, 10))) AS NameFirm, {{!SQL.SOURCE}}.Deals.CustomerExternalView, {{!SQL.SOURCE}}.Deals.CustomerSienaView, { fn CONCAT(RTRIM(LTRIM(LEFT({{!SQL.SOURCE}}.Deals.CustomerSienaView, 10))), { fn CONCAT('|', 
                         RIGHT({{!SQL.SOURCE}}.Deals.CustomerSienaView, 3)) }) } AS CompID, {{!SQL.SOURCE}}.Deals.SienaDealer, {{!SQL.SOURCE}}.Deals.DealOwner, {{!SQL.SOURCE}}.Deals.DealOwnerMnemonic, {{!SQL.SOURCE}}.Deals.EditedByUser, {{!SQL.SOURCE}}.Deals.UTCOriginTime, {{!SQL.SOURCE}}.Deals.UTCUpdateTime, 
                         {{!SQL.SOURCE}}.Deals.MarginTrading, {{!SQL.SOURCE}}.Deals.SwapPoints, {{!SQL.SOURCE}}.Deals.SpotDate, {{!SQL.SOURCE}}.Deals.SpotRate, {{!SQL.SOURCE}}.Deals.MktSpotRate, {{!SQL.SOURCE}}.Deals.SpotSalesMargin, {{!SQL.SOURCE}}.Deals.SpotChlMargin, {{!SQL.SOURCE}}.Deals.RerouteCcy, {{!SQL.SOURCE}}.Deals.CustomerPayInstruction, 
                         {{!SQL.SOURCE}}.Deals.CustomerReceiptInstruction, {{!SQL.SOURCE}}.Deals.BackOfficeNotes, {{!SQL.SOURCE}}.Deals.customerStatementNotes, {{!SQL.SOURCE}}.Deals.NotesMargin, {{!SQL.SOURCE}}.Deals.RequestedBy, {{!SQL.SOURCE}}.Deals.EditReason, {{!SQL.SOURCE}}.Deals.EditOtherReason, 
                         {{!SQL.SOURCE}}.Deals.NICleanPrice, {{!SQL.SOURCE}}.Deals.NIDirtyPrice, {{!SQL.SOURCE}}.Deals.NIYield, {{!SQL.SOURCE}}.Deals.NIClearingSystem, {{!SQL.SOURCE}}.Deals.Acceptor, {{!SQL.SOURCE}}.Deals.NIDiscount, {{!SQL.SOURCE}}.Deals.FastPay, {{!SQL.SOURCE}}.Deals.PaymentFee, {{!SQL.SOURCE}}.Deals.PaymentFeePolicy, 
                         {{!SQL.SOURCE}}.Deals.PaymentReason, {{!SQL.SOURCE}}.Deals.PaymentDate, {{!SQL.SOURCE}}.Deals.SettlementDate, {{!SQL.SOURCE}}.Deals.FixingDate, {{!SQL.SOURCE}}.Deals.VenueUTI, {{!SQL.SOURCE}}.Deals.EditVersion, {{!SQL.SOURCE}}.Deals.BrokeragePercentage, {{!SQL.SOURCE}}.Deals.BrokerageAmount, 
                         {{!SQL.SOURCE}}.Deals.BrokerageCurrency, {{!SQL.SOURCE}}.Deals.BrokerageDate, {{!SQL.SOURCE}}.Deals.AccountName, {{!SQL.SOURCE}}.Deals.AccountNumber, {{!SQL.SOURCE}}.Deals.CashBalance, {{!SQL.SOURCE}}.Deals.DebitFrequency, {{!SQL.SOURCE}}.Deals.CreditFrequency, {{!SQL.SOURCE}}.Deals.ManuallyQuoted, 
                         {{!SQL.SOURCE}}.Deals.LedgerBalance, {{!SQL.SOURCE}}.Deals.SettAmtOutstanding, {{!SQL.SOURCE}}.Deals.FeePercentage, {{!SQL.SOURCE}}.Deals.FeeAmount, {{!SQL.SOURCE}}.Deals.Venue, {{!SQL.SOURCE}}.Deals.EURAmount, {{!SQL.SOURCE}}.Deals.EUROtherAmount, {{!SQL.SOURCE}}.Deals.LEI, {{!SQL.SOURCE}}.Deals.Equity, {{!SQL.SOURCE}}.Deals.Shares, 
                         {{!SQL.SOURCE}}.Deals.QuoteExpiryDate, {{!SQL.SOURCE}}.Deals.Commodity, {{!SQL.SOURCE}}.Deals.PaymentSystemSienaView, {{!SQL.SOURCE}}.Deals.PaymentSystemExternalView, {{!SQL.SOURCE}}.Deals.SalesProfit, {{!SQL.SOURCE}}.Deals.RejectReason, {{!SQL.SOURCE}}.Deals.PaymentError, 
                         {{!SQL.SOURCE}}.Deals.RepoPrincipal, {{!SQL.SOURCE}}.Deals.FixingFrequency
FROM            {{!SQL.SOURCE}}.Deals INNER JOIN
                         {{!SQL.SOURCE}}.DealType ON {{!SQL.SOURCE}}.Deals.FullDealType = {{!SQL.SOURCE}}.DealType.DealTypeKey INNER JOIN
                         {{!SQL.SOURCE}}.FundamentalDealType ON {{!SQL.SOURCE}}.DealType.FundamentalDealTypeKey = {{!SQL.SOURCE}}.FundamentalDealType.DealTypeKey INNER JOIN
                         {{!SQL.SOURCE}}.Currency ON {{!SQL.SOURCE}}.Deals.DealtCcy = {{!SQL.SOURCE}}.Currency.Code LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Book AS Book_1 ON {{!SQL.SOURCE}}.Deals.CounterBook = Book_1.BookName LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.DealRevaluation ON {{!SQL.SOURCE}}.Deals.SienaReference = {{!SQL.SOURCE}}.DealRevaluation.DealRefNo LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Book ON {{!SQL.SOURCE}}.Deals.Book = {{!SQL.SOURCE}}.Book.BookName LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Portfolio ON {{!SQL.SOURCE}}.Deals.Portfolio = {{!SQL.SOURCE}}.Portfolio.Code
WHERE        ({{!SQL.SOURCE}}.Deals.InternalDeleted IS NULL) AND ({{!SQL.SOURCE}}.FundamentalDealType.DealTypeShortName <> 'Acct') AND ({{!SQL.SOURCE}}.FundamentalDealType.InternalDeleted IS NULL) AND ({{!SQL.SOURCE}}.DealType.InternalDeleted IS NULL) AND 
                         ({{!SQL.SOURCE}}.Deals.LimitOrderType IS NULL)
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane1', @value=N'[0E232FF0-B466-11cf-A24F-00AA00A3EFFF, 1.00]
Begin DesignProperties = 
   Begin PaneConfigurations = 
      Begin PaneConfiguration = 0
         NumPanes = 4
         Configuration = "(H (1[30] 4[32] 2[20] 3) )"
      End
      Begin PaneConfiguration = 1
         NumPanes = 3
         Configuration = "(H (1 [50] 4 [25] 3))"
      End
      Begin PaneConfiguration = 2
         NumPanes = 3
         Configuration = "(H (1 [50] 2 [25] 3))"
      End
      Begin PaneConfiguration = 3
         NumPanes = 3
         Configuration = "(H (4 [30] 2 [40] 3))"
      End
      Begin PaneConfiguration = 4
         NumPanes = 2
         Configuration = "(H (1 [56] 3))"
      End
      Begin PaneConfiguration = 5
         NumPanes = 2
         Configuration = "(H (2 [66] 3))"
      End
      Begin PaneConfiguration = 6
         NumPanes = 2
         Configuration = "(H (4 [50] 3))"
      End
      Begin PaneConfiguration = 7
         NumPanes = 1
         Configuration = "(V (3))"
      End
      Begin PaneConfiguration = 8
         NumPanes = 3
         Configuration = "(H (1[56] 4[18] 2) )"
      End
      Begin PaneConfiguration = 9
         NumPanes = 2
         Configuration = "(H (1 [75] 4))"
      End
      Begin PaneConfiguration = 10
         NumPanes = 2
         Configuration = "(H (1[66] 2) )"
      End
      Begin PaneConfiguration = 11
         NumPanes = 2
         Configuration = "(H (4 [60] 2))"
      End
      Begin PaneConfiguration = 12
         NumPanes = 1
         Configuration = "(H (1) )"
      End
      Begin PaneConfiguration = 13
         NumPanes = 1
         Configuration = "(V (4))"
      End
      Begin PaneConfiguration = 14
         NumPanes = 1
         Configuration = "(V (2))"
      End
      ActivePaneConfig = 0
   End
   Begin DiagramPane = 
      Begin Origin = 
         Top = 0
         Left = 0
      End
      Begin Tables = 
         Begin Table = "Deals (RG)"
            Begin Extent = 
               Top = 6
               Left = 38
               Bottom = 211
               Right = 292
            End
            DisplayFlags = 280
            TopColumn = 213
         End
         Begin Table = "DealType (RG)"
            Begin Extent = 
               Top = 25
               Left = 550
               Bottom = 155
               Right = 776
            End
            DisplayFlags = 280
            TopColumn = 0
         End
         Begin Table = "FundamentalDealType (RG)"
            Begin Extent = 
               Top = 270
               Left = 38
               Bottom = 400
               Right = 243
            End
            DisplayFlags = 280
            TopColumn = 0
         End
         Begin Table = "Currency (RG)"
            Begin Extent = 
               Top = 402
               Left = 38
               Bottom = 532
               Right = 410
            End
            DisplayFlags = 280
            TopColumn = 0
         End
         Begin Table = "Book_1"
            Begin Extent = 
               Top = 534
               Left = 38
               Bottom = 664
               Right = 243
            End
            DisplayFlags = 280
            TopColumn = 0
         End
         Begin Table = "DealRevaluation (RG)"
            Begin Extent = 
               Top = 666
               Left = 38
               Bottom = 796
               Right = 243
            End
            DisplayFlags = 280
            TopColumn = 0
         End
         Begin Table = "Book (RG)"
            Begin Extent = 
               Top = 798
               Left = 38
               Bottom = 928
               Right = 243
  ' , @level0type=N'SCHEMA',@level0name=N'SRS', @level1type=N'VIEW',@level1name=N'sienaDealList'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane2', @value=N'          End
            DisplayFlags = 280
            TopColumn = 0
         End
         Begin Table = "Portfolio (RG)"
            Begin Extent = 
               Top = 930
               Left = 38
               Bottom = 1060
               Right = 243
            End
            DisplayFlags = 280
            TopColumn = 0
         End
      End
   End
   Begin SQLPane = 
   End
   Begin DataPane = 
      Begin ParameterDefaults = ""
      End
      Begin ColumnWidths = 108
         Width = 284
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1305
         Width = 990
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
      End
   End
   Begin CriteriaPane = 
      Begin ColumnWidths = 11
         Column = 6225
         Alias = 1440
         Table = 1170
         Output = 720
         Append = 1400
         NewValue = 1170
         SortType = 1350
         SortOrder = 1410
         GroupBy = 1350
         Filter = 1350
         Or = 1350
         Or = 1350
         Or = 1350
      End
   End
End
' , @level0type=N'SCHEMA',@level0name=N'SRS', @level1type=N'VIEW',@level1name=N'sienaDealList'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPaneCount', @value=2 , @level0type=N'SCHEMA',@level0name=N'SRS', @level1type=N'VIEW',@level1name=N'sienaDealList'
GO
