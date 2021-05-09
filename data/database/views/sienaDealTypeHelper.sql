/****** Object:  View [{{!SQL.SCHEMA}}].[sienaDealTypeHelper]    Script Date: 09/05/2021 14:36:12 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaDealTypeHelper]
AS
SELECT        {{!SQL.SOURCE}}.DealType.DealTypeKey, {{!SQL.SOURCE}}.DealType.DealTypeShortName, {{!SQL.SOURCE}}.FundamentalDealTypeSwitches.NegotiableInstrumentType, {{!SQL.SOURCE}}.FundamentalDealTypeSwitches.ProductCode, {{!SQL.SOURCE}}.Product.Name AS ProductCodeName
FROM            {{!SQL.SOURCE}}.Product RIGHT OUTER JOIN
                         {{!SQL.SOURCE}}.FundamentalDealTypeSwitches ON {{!SQL.SOURCE}}.Product.Code = {{!SQL.SOURCE}}.FundamentalDealTypeSwitches.ProductCode RIGHT OUTER JOIN
                         {{!SQL.SOURCE}}.DealType ON {{!SQL.SOURCE}}.FundamentalDealTypeSwitches.DealTypeKey = {{!SQL.SOURCE}}.DealType.DealTypeKey
WHERE        ({{!SQL.SOURCE}}.DealType.InternalDeleted IS NULL) AND ({{!SQL.SOURCE}}.DealType.IsActive = 1) AND ({{!SQL.SOURCE}}.FundamentalDealTypeSwitches.Internal = 0)
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane1', @value=N'[0E232FF0-B466-11cf-A24F-00AA00A3EFFF, 1.00]
Begin DesignProperties = 
   Begin PaneConfigurations = 
      Begin PaneConfiguration = 0
         NumPanes = 4
         Configuration = "(H (1[46] 4[16] 2[14] 3) )"
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
         Begin Table = "DealType"
            Begin Extent = 
               Top = 6
               Left = 38
               Bottom = 170
               Right = 264
            End
            DisplayFlags = 280
            TopColumn = 20
         End
         Begin Table = "FundamentalDealTypeSwitches"
            Begin Extent = 
               Top = 6
               Left = 292
               Bottom = 162
               Right = 548
            End
            DisplayFlags = 280
            TopColumn = 85
         End
         Begin Table = "Product"
            Begin Extent = 
               Top = 204
               Left = 513
               Bottom = 334
               Right = 718
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
      Begin ColumnWidths = 11
         Width = 284
         Width = 1980
         Width = 1950
         Width = 2280
         Width = 1500
         Width = 2250
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
      End
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
' , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaDealTypeHelper'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPaneCount', @value=1 , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaDealTypeHelper'
GO
