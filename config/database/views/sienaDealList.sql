/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDealList]    Script Date: 09/05/2021 14:36:12 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaDealList]
AS
SELECT        {{!SQL.SOURCE}}.Deals.SienaReference, {{!SQL.SOURCE}}.Deals.CustomerSienaView, {{!SQL.SOURCE}}.Deals.Status, {{!SQL.SOURCE}}.Deals.StartDate AS ValueDate, {{!SQL.SOURCE}}.Deals.MaturityDate, {{!SQL.SOURCE}}.Deals.ContractNumber, {{!SQL.SOURCE}}.Deals.ExternalReference, {{!SQL.SOURCE}}.Deals.Book, 
                         {{!SQL.SOURCE}}.Deals.MandatedUser, {{!SQL.SOURCE}}.Deals.Portfolio, {{!SQL.SOURCE}}.Deals.AgreementId, {{!SQL.SOURCE}}.Deals.BackOfficeRefNo, {{!SQL.SOURCE}}.Deals.ISIN, {{!SQL.SOURCE}}.Deals.UTI, {{!SQL.SOURCE}}.Book.FullName AS BookName, RIGHT({{!SQL.SOURCE}}.Deals.CustomerSienaView, 3) AS Centre, 
                         RTRIM(LTRIM(LEFT({{!SQL.SOURCE}}.Deals.CustomerSienaView, 10))) AS Firm, {{!SQL.SOURCE}}.FundamentalDealType.DealTypeShortName, {{!SQL.SOURCE}}.Deals.FullDealType, {{!SQL.SOURCE}}.Deals.TradeDate, {{!SQL.SOURCE}}.Deals.DealtCcy, {{!SQL.SOURCE}}.Deals.DealtAmount, 
                         {{!SQL.SOURCE}}.Deals.AgainstAmount, {{!SQL.SOURCE}}.Deals.AgainstCcy, {{!SQL.SOURCE}}.Deals.AllInRate, {{!SQL.SOURCE}}.Deals.MktRate, {{!SQL.SOURCE}}.Deals.SettleCcy, {{!SQL.SOURCE}}.Deals.Direction, {{!SQL.SOURCE}}.Deals.NpvRate, {{!SQL.SOURCE}}.Deals.OriginUser, {{!SQL.SOURCE}}.Deals.PayInstruction, 
                         {{!SQL.SOURCE}}.Deals.ReceiptInstruction, {{!SQL.SOURCE}}.Deals.NIName, { fn CONCAT({{!SQL.SOURCE}}.Deals.DealtCcy, {{!SQL.SOURCE}}.Deals.AgainstCcy) } AS CCYPair, ISNULL(NULLIF ({{!SQL.SOURCE}}.Deals.NIName, NULL), { fn CONCAT({{!SQL.SOURCE}}.Deals.DealtCcy, {{!SQL.SOURCE}}.Deals.AgainstCcy) }) 
                         AS Instrument, {{!SQL.SOURCE}}.Portfolio.Description1 AS PortfolioName, {{!SQL.SOURCE}}.DealRevaluation.Date AS RVDate, {{!SQL.SOURCE}}.DealRevaluation.MarkToMarket AS RVMTM, {{!SQL.SOURCE}}.Deals.CounterBook, Book_1.FullName AS CounterBookName, 
                         ISNULL(NULLIF ({{!SQL.SOURCE}}.Deals.CounterBook, NULL), {{!SQL.SOURCE}}.Deals.CustomerSienaView) AS Party, ISNULL(NULLIF (Book_1.FullName, NULL), {{!SQL.SOURCE}}.Deals.CustomerSienaView) AS PartyName
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
         Configuration = "(H (1[6] 4[53] 2[16] 3) )"
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
         Left = -288
      End
      Begin Tables = 
         Begin Table = "Deals"
            Begin Extent = 
               Top = 13
               Left = 40
               Bottom = 143
               Right = 294
            End
            DisplayFlags = 280
            TopColumn = 29
         End
         Begin Table = "DealType"
            Begin Extent = 
               Top = 59
               Left = 74
               Bottom = 189
               Right = 300
            End
            DisplayFlags = 344
            TopColumn = 22
         End
         Begin Table = "FundamentalDealType"
            Begin Extent = 
               Top = 2
               Left = 600
               Bottom = 132
               Right = 805
            End
            DisplayFlags = 344
            TopColumn = 0
         End
         Begin Table = "Currency"
            Begin Extent = 
               Top = 96
               Left = 328
               Bottom = 226
               Right = 700
            End
            DisplayFlags = 344
            TopColumn = 0
         End
         Begin Table = "Book_1"
            Begin Extent = 
               Top = 27
               Left = 331
               Bottom = 157
               Right = 536
            End
            DisplayFlags = 344
            TopColumn = 0
         End
         Begin Table = "DealRevaluation"
            Begin Extent = 
               Top = 148
               Left = 79
               Bottom = 278
               Right = 284
            End
            DisplayFlags = 344
            TopColumn = 3
         End
         Begin Table = "Book"
            Begin Extent = 
               Top = 8
               Left = 142
               Bottom = 138
               Right = 347
            End
            Dis' , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaDealList'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane2', @value=N'playFlags = 344
            TopColumn = 0
         End
         Begin Table = "Portfolio"
            Begin Extent = 
               Top = 175
               Left = 340
               Bottom = 333
               Right = 545
            End
            DisplayFlags = 344
            TopColumn = 0
         End
      End
   End
   Begin SQLPane = 
   End
   Begin DataPane = 
      Begin ParameterDefaults = ""
      End
      Begin ColumnWidths = 97
         Width = 284
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1800
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1725
         Width = 1500
         Width = 2070
         Width = 2505
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1680
         Width = 1500
         Width = 2625
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1890
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 2970
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1290
         Width = 420
         Width = 555
         Width = 945
         Width = 1155
         Width = 1335
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
         Width = 2160
         Width = 2130
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
         Column = 12375
         Alias = 900
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
' , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaDealList'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPaneCount', @value=2 , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaDealList'
GO
