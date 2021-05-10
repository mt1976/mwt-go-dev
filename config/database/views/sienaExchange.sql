/****** Object:  View [{{!SQL.SCHEMA}}].[sienaExchange]    Script Date: 09/05/2021 14:36:12 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaExchange]
AS
SELECT        {{!SQL.SOURCE}}.Exchange.FullName, {{!SQL.SOURCE}}.Exchange.Broker, {{!SQL.SOURCE}}.Exchange.Country, {{!SQL.SOURCE}}.Exchange.Centre, {{!SQL.SOURCE}}.Exchange.CounterpartyFirm, {{!SQL.SOURCE}}.Exchange.CounterpartyCentre, {{!SQL.SOURCE}}.Exchange.LEI, {{!SQL.SOURCE}}.Country.Name AS CountryName, 
                         Centre_1.FullName AS CentreName, {{!SQL.SOURCE}}.Firm.FullName AS CounterpartyFirmName, {{!SQL.SOURCE}}.Centre.FullName AS CounterpartCentreName, {{!SQL.SOURCE}}.Broker.Name AS BrokerName, {{!SQL.SOURCE}}.Broker.FullName AS BrokerFullName
FROM            {{!SQL.SOURCE}}.Exchange LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Broker ON {{!SQL.SOURCE}}.Exchange.Broker = {{!SQL.SOURCE}}.Broker.Code LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Centre ON {{!SQL.SOURCE}}.Exchange.CounterpartyCentre = {{!SQL.SOURCE}}.Centre.ShortName LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Firm ON {{!SQL.SOURCE}}.Exchange.CounterpartyFirm = {{!SQL.SOURCE}}.Firm.FirmName LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Centre AS Centre_1 ON {{!SQL.SOURCE}}.Exchange.Centre = Centre_1.ShortName LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Country ON {{!SQL.SOURCE}}.Exchange.Country = {{!SQL.SOURCE}}.Country.Code
WHERE        ({{!SQL.SOURCE}}.Exchange.InternalDeleted IS NULL)
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane1', @value=N'[0E232FF0-B466-11cf-A24F-00AA00A3EFFF, 1.00]
Begin DesignProperties = 
   Begin PaneConfigurations = 
      Begin PaneConfiguration = 0
         NumPanes = 4
         Configuration = "(H (1[40] 4[21] 2[11] 3) )"
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
         Begin Table = "Exchange"
            Begin Extent = 
               Top = 6
               Left = 38
               Bottom = 224
               Right = 243
            End
            DisplayFlags = 280
            TopColumn = 1
         End
         Begin Table = "Country"
            Begin Extent = 
               Top = 6
               Left = 281
               Bottom = 136
               Right = 486
            End
            DisplayFlags = 344
            TopColumn = 0
         End
         Begin Table = "Centre"
            Begin Extent = 
               Top = 6
               Left = 524
               Bottom = 136
               Right = 729
            End
            DisplayFlags = 344
            TopColumn = 0
         End
         Begin Table = "Centre_1"
            Begin Extent = 
               Top = 53
               Left = 286
               Bottom = 183
               Right = 491
            End
            DisplayFlags = 344
            TopColumn = 0
         End
         Begin Table = "Firm"
            Begin Extent = 
               Top = 53
               Left = 523
               Bottom = 183
               Right = 728
            End
            DisplayFlags = 344
            TopColumn = 0
         End
         Begin Table = "Broker"
            Begin Extent = 
               Top = 76
               Left = 464
               Bottom = 206
               Right = 677
            End
            DisplayFlags = 344
            TopColumn = 8
         End
      End
   End
   Begin SQLPane = 
   End
   Begin DataPane = 
      Begin ParameterDefaults = ""
      End
      Begin ColumnWidths = 14
         Width = 284
         Width = 2805
         Width = 1500
         Wi' , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaExchange'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane2', @value=N'dth = 1500
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
' , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaExchange'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPaneCount', @value=2 , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaExchange'
GO
