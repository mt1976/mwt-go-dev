/****** Object:  View [{{!SQL.SCHEMA}}].[sienaAccount]    Script Date: 09/05/2021 14:36:12 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaAccount]
AS
SELECT        {{!SQL.SOURCE}}.Deals.SienaReference, {{!SQL.SOURCE}}.Deals.CustomerSienaView, {{!SQL.SOURCE}}.Deals.SienaCommonRef, {{!SQL.SOURCE}}.Deals.Status, {{!SQL.SOURCE}}.Deals.StartDate, {{!SQL.SOURCE}}.Deals.MaturityDate, {{!SQL.SOURCE}}.Deals.ContractNumber, {{!SQL.SOURCE}}.Deals.ExternalReference,
                         {{!SQL.SOURCE}}.Deals.DealtCcy AS CCY, {{!SQL.SOURCE}}.Deals.Book, {{!SQL.SOURCE}}.Deals.MandatedUser, {{!SQL.SOURCE}}.Deals.BackOfficeNotes, {{!SQL.SOURCE}}.Deals.CashBalance, {{!SQL.SOURCE}}.Deals.AccountNumber, {{!SQL.SOURCE}}.Deals.AccountName, {{!SQL.SOURCE}}.Deals.LedgerBalance, {{!SQL.SOURCE}}.Deals.Portfolio,
                         {{!SQL.SOURCE}}.Deals.AgreementId, {{!SQL.SOURCE}}.Deals.BackOfficeRefNo, {{!SQL.SOURCE}}.Deals.PaymentSystemSienaView, {{!SQL.SOURCE}}.Deals.ISIN, {{!SQL.SOURCE}}.Deals.UTI, {{!SQL.SOURCE}}.Currency.Name AS CCYName, {{!SQL.SOURCE}}.Book.FullName AS BookName,
                         {{!SQL.SOURCE}}.Portfolio.Description1 AS PortfolioName, RIGHT({{!SQL.SOURCE}}.Deals.CustomerSienaView, 3) AS Centre, LEFT({{!SQL.SOURCE}}.Deals.CustomerSienaView, LEN({{!SQL.SOURCE}}.Deals.CustomerSienaView) - 3) AS Firm, {{!SQL.SOURCE}}.Currency.AmountDp AS CCYDp
FROM            {{!SQL.SOURCE}}.FundamentalDealType LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.DealType ON {{!SQL.SOURCE}}.FundamentalDealType.DealTypeKey = {{!SQL.SOURCE}}.DealType.FundamentalDealTypeKey RIGHT OUTER JOIN
                         {{!SQL.SOURCE}}.Deals ON {{!SQL.SOURCE}}.DealType.DealTypeKey = {{!SQL.SOURCE}}.Deals.FullDealType LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Currency ON {{!SQL.SOURCE}}.Deals.DealtCcy = {{!SQL.SOURCE}}.Currency.Code LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Book ON {{!SQL.SOURCE}}.Deals.Book = {{!SQL.SOURCE}}.Book.BookName LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Portfolio ON {{!SQL.SOURCE}}.Deals.Portfolio = {{!SQL.SOURCE}}.Portfolio.Code
WHERE        ({{!SQL.SOURCE}}.FundamentalDealType.DealTypeShortName = 'Acct') AND ({{!SQL.SOURCE}}.FundamentalDealType.InternalDeleted IS NULL) AND ({{!SQL.SOURCE}}.DealType.InternalDeleted IS NULL) AND ({{!SQL.SOURCE}}.Deals.InternalDeleted IS NULL)
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane1', @value=N'[0E232FF0-B466-11cf-A24F-00AA00A3EFFF, 1.00]
Begin DesignProperties =
   Begin PaneConfigurations =
      Begin PaneConfiguration = 0
         NumPanes = 4
         Configuration = "(H (1[25] 4[45] 2[12] 3) )"
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
         Begin Table = "Deals"
            Begin Extent =
               Top = 6
               Left = 39
               Bottom = 117
               Right = 292
            End
            DisplayFlags = 280
            TopColumn = 218
         End
         Begin Table = "DealType"
            Begin Extent =
               Top = 73
               Left = 459
               Bottom = 185
               Right = 685
            End
            DisplayFlags = 344
            TopColumn = 0
         End
         Begin Table = "FundamentalDealType"
            Begin Extent =
               Top = 153
               Left = 715
               Bottom = 283
               Right = 920
            End
            DisplayFlags = 344
            TopColumn = 0
         End
         Begin Table = "Currency"
            Begin Extent =
               Top = 227
               Left = 322
               Bottom = 357
               Right = 694
            End
            DisplayFlags = 280
            TopColumn = 35
         End
         Begin Table = "Book"
            Begin Extent =
               Top = 149
               Left = 477
               Bottom = 279
               Right = 682
            End
            DisplayFlags = 344
            TopColumn = 0
         End
         Begin Table = "Portfolio"
            Begin Extent =
               Top = 214
               Left = 410
               Bottom = 344
               Right = 615
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
      Begin ColumnWidths = 29
         Width = 284
         Width = 1500
         Wi' , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaAccount'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane2', @value=N'dth = 1965
         Width = 2610
         Width = 1845
         Width = 1920
         Width = 2565
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1290
         Width = 1500
         Width = 1500
         Width = 1935
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1965
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
      End
   End
   Begin CriteriaPane =
      Begin ColumnWidths = 11
         Column = 5730
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
' , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaAccount'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPaneCount', @value=2 , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaAccount'
GO
