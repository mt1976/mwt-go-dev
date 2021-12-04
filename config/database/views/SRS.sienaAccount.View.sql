USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaAccount]    Script Date: 04/12/2021 18:17:08 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaAccount]
AS
SELECT        {{!SQL.SOURCE}}.Deals.SienaReference, {{!SQL.SOURCE}}.Deals.CustomerSienaView, {{!SQL.SOURCE}}.Deals.SienaCommonRef, {{!SQL.SOURCE}}.Deals.Status, {{!SQL.SOURCE}}.Deals.StartDate, {{!SQL.SOURCE}}.Deals.MaturityDate, {{!SQL.SOURCE}}.Deals.ContractNumber, {{!SQL.SOURCE}}.Deals.ExternalReference, 
                         {{!SQL.SOURCE}}.Deals.DealtCcy AS CCY, {{!SQL.SOURCE}}.Deals.Book, {{!SQL.SOURCE}}.Deals.MandatedUser, {{!SQL.SOURCE}}.Deals.BackOfficeNotes, {{!SQL.SOURCE}}.Deals.CashBalance, {{!SQL.SOURCE}}.Deals.AccountNumber, {{!SQL.SOURCE}}.Deals.AccountName, {{!SQL.SOURCE}}.Deals.LedgerBalance, {{!SQL.SOURCE}}.Deals.Portfolio, 
                         {{!SQL.SOURCE}}.Deals.AgreementId, {{!SQL.SOURCE}}.Deals.BackOfficeRefNo, {{!SQL.SOURCE}}.Deals.ISIN, {{!SQL.SOURCE}}.Deals.UTI, {{!SQL.SOURCE}}.Currency.Name AS CCYName, {{!SQL.SOURCE}}.Book.FullName AS BookName, {{!SQL.SOURCE}}.Portfolio.Description1 AS PortfolioName, 
                         RIGHT({{!SQL.SOURCE}}.Deals.CustomerSienaView, 3) AS Centre, {{!SQL.SOURCE}}.FundamentalDealType.*, {{!SQL.SOURCE}}.Currency.AmountDp AS CCYDp, { fn CONCAT(RTRIM(LTRIM(LEFT({{!SQL.SOURCE}}.Deals.CustomerSienaView, 10))), { fn CONCAT('|', 
                         RIGHT({{!SQL.SOURCE}}.Deals.CustomerSienaView, 3)) }) } AS CompID, RTRIM(LTRIM(LEFT({{!SQL.SOURCE}}.Deals.CustomerSienaView, 10))) AS Firm, {{!SQL.SOURCE}}.Deals.DealType, {{!SQL.SOURCE}}.Deals.FullDealType, {{!SQL.SOURCE}}.Deals.DealingInterface, {{!SQL.SOURCE}}.Deals.DealtAmount, 
                         {{!SQL.SOURCE}}.Deals.ParentContractNumber, {{!SQL.SOURCE}}.Deals.InterestFrequency, {{!SQL.SOURCE}}.Deals.InterestAction, {{!SQL.SOURCE}}.Deals.InterestStrategy, {{!SQL.SOURCE}}.Deals.InterestBasis, {{!SQL.SOURCE}}.Deals.SienaDealer, {{!SQL.SOURCE}}.Deals.DealOwner, {{!SQL.SOURCE}}.Deals.OriginUser, 
                         {{!SQL.SOURCE}}.Deals.EditedByUser, {{!SQL.SOURCE}}.Deals.DealOwnerMnemonic, {{!SQL.SOURCE}}.Deals.UTCOriginTime, {{!SQL.SOURCE}}.Deals.UTCUpdateTime, {{!SQL.SOURCE}}.Deals.customerStatementNotes, {{!SQL.SOURCE}}.Deals.NotesMargin, {{!SQL.SOURCE}}.Deals.RequestedBy, {{!SQL.SOURCE}}.Deals.EditReason, 
                         {{!SQL.SOURCE}}.Deals.EditOtherReason, {{!SQL.SOURCE}}.Deals.NoticeDays, {{!SQL.SOURCE}}.Deals.DebitFrequency, {{!SQL.SOURCE}}.Deals.CreditFrequency, {{!SQL.SOURCE}}.Deals.EURAmount, {{!SQL.SOURCE}}.Deals.EUROtherAmount, {{!SQL.SOURCE}}.Deals.PaymentSystemSienaView, 
                         {{!SQL.SOURCE}}.Deals.PaymentSystemExternalView
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
         Configuration = "(H (1[41] 4[25] 2[16] 3) )"
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
         Top = -192
         Left = 0
      End
      Begin Tables = 
         Begin Table = "FundamentalDealType (RG)"
            Begin Extent = 
               Top = 6
               Left = 38
               Bottom = 187
               Right = 243
            End
            DisplayFlags = 280
            TopColumn = 0
         End
         Begin Table = "DealType (RG)"
            Begin Extent = 
               Top = 22
               Left = 377
               Bottom = 152
               Right = 603
            End
            DisplayFlags = 280
            TopColumn = 0
         End
         Begin Table = "Deals (RG)"
            Begin Extent = 
               Top = 204
               Left = 38
               Bottom = 400
               Right = 292
            End
            DisplayFlags = 280
            TopColumn = 213
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
         Begin Table = "Book (RG)"
            Begin Extent = 
               Top = 534
               Left = 38
               Bottom = 664
               Right = 243
            End
            DisplayFlags = 280
            TopColumn = 0
         End
         Begin Table = "Portfolio (RG)"
            Begin Extent = 
               Top = 666
               Left = 38
               Bottom = 796
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
      Begin ColumnWidths = 66
         Width = 284
      ' , @level0type=N'SCHEMA',@level0name=N'SRS', @level1type=N'VIEW',@level1name=N'sienaAccount'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane2', @value=N'   Width = 1500
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
         Width = 1500
      End
   End
   Begin CriteriaPane = 
      Begin ColumnWidths = 11
         Column = 3840
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
' , @level0type=N'SCHEMA',@level0name=N'SRS', @level1type=N'VIEW',@level1name=N'sienaAccount'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPaneCount', @value=2 , @level0type=N'SCHEMA',@level0name=N'SRS', @level1type=N'VIEW',@level1name=N'sienaAccount'
GO
