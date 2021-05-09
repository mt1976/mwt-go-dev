/****** Object:  View [{{!SQL.SCHEMA}}].[sienaBookedRates]    Script Date: 09/05/2021 14:36:12 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaBookedRates]
AS
SELECT        {{!SQL.SOURCE}}.Deals.SienaReference, {{!SQL.SOURCE}}.Deals.UTI, {{!SQL.SOURCE}}.Deals.FullDealType, {{!SQL.SOURCE}}.Deals.Broker, {{!SQL.SOURCE}}.Deals.Status, {{!SQL.SOURCE}}.Deals.TradeDate, {{!SQL.SOURCE}}.Deals.StartDate, {{!SQL.SOURCE}}.Deals.MaturityDate, {{!SQL.SOURCE}}.Deals.ContractNumber, 
                         {{!SQL.SOURCE}}.Deals.ExternalReference, {{!SQL.SOURCE}}.Deals.DealingInterface, {{!SQL.SOURCE}}.Deals.DealtAmount, {{!SQL.SOURCE}}.Deals.AgainstAmount, {{!SQL.SOURCE}}.Deals.DealtCcy, {{!SQL.SOURCE}}.Deals.AgainstCcy, {{!SQL.SOURCE}}.Deals.MajorCcy, {{!SQL.SOURCE}}.Deals.MinorCcy, {{!SQL.SOURCE}}.Deals.AllInRate, 
                         {{!SQL.SOURCE}}.Deals.Book, {{!SQL.SOURCE}}.Deals.SettleCcy, {{!SQL.SOURCE}}.Deals.Direction, {{!SQL.SOURCE}}.Deals.MandatedUser, {{!SQL.SOURCE}}.Deals.CustomerSienaView, {{!SQL.SOURCE}}.Deals.PayInstruction, {{!SQL.SOURCE}}.Deals.ReceiptInstruction, {{!SQL.SOURCE}}.Deals.VenueUTI, {{!SQL.SOURCE}}.Deals.Portfolio, 
                         {{!SQL.SOURCE}}.Deals.ISIN, {{!SQL.SOURCE}}.Deals.Waiver, {{!SQL.SOURCE}}.Deals.LEI, {{!SQL.SOURCE}}.Deals.AgreementId, {{!SQL.SOURCE}}.Deals.DealOwner, {{!SQL.SOURCE}}.Deals.OriginUser, {{!SQL.SOURCE}}.Deals.SienaCommonRef
FROM            {{!SQL.SOURCE}}.FundamentalDealType RIGHT OUTER JOIN
                         {{!SQL.SOURCE}}.DealType ON {{!SQL.SOURCE}}.FundamentalDealType.DealTypeKey = {{!SQL.SOURCE}}.DealType.FundamentalDealTypeKey RIGHT OUTER JOIN
                         {{!SQL.SOURCE}}.Deals ON {{!SQL.SOURCE}}.DealType.DealTypeKey = {{!SQL.SOURCE}}.Deals.FullDealType
WHERE        ({{!SQL.SOURCE}}.Deals.InternalDeleted IS NULL) AND ({{!SQL.SOURCE}}.FundamentalDealType.DealTypeShortName = 'BR')
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane1', @value=N'[0E232FF0-B466-11cf-A24F-00AA00A3EFFF, 1.00]
Begin DesignProperties = 
   Begin PaneConfigurations = 
      Begin PaneConfiguration = 0
         NumPanes = 4
         Configuration = "(H (1[41] 4[20] 2[7] 3) )"
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
               Left = 38
               Bottom = 189
               Right = 292
            End
            DisplayFlags = 280
            TopColumn = 0
         End
         Begin Table = "DealType"
            Begin Extent = 
               Top = 6
               Left = 330
               Bottom = 226
               Right = 556
            End
            DisplayFlags = 280
            TopColumn = 7
         End
         Begin Table = "FundamentalDealType"
            Begin Extent = 
               Top = 6
               Left = 594
               Bottom = 222
               Right = 799
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
      Begin ColumnWidths = 37
         Width = 284
         Width = 1500
         Width = 1995
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 2340
         Width = 1500
         Width = 1935
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
  ' , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaBookedRates'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane2', @value=N'    End
   End
   Begin CriteriaPane = 
      Begin ColumnWidths = 11
         Column = 1440
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
' , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaBookedRates'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPaneCount', @value=2 , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaBookedRates'
GO
