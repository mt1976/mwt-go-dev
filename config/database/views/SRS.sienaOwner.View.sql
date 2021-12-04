USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaOwner]    Script Date: 04/12/2021 18:17:00 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaOwner]
AS
SELECT        {{!SQL.SOURCE}}.Usr.UserName, {{!SQL.SOURCE}}.Usr.FullName, {{!SQL.SOURCE}}.Usr.Type, {{!SQL.SOURCE}}.Usr.TradingEntity, {{!SQL.SOURCE}}.UserDetailsSettings.DefaultEnterBook, {{!SQL.SOURCE}}.UserDetailsSettings.EmailAddress, {{!SQL.SOURCE}}.UserDetailsSettings.Enabled, 
                         {{!SQL.SOURCE}}.UserDetailsSettings.ExternalUserIds, {{!SQL.SOURCE}}.UserDetailsSettings.Language, {{!SQL.SOURCE}}.UserDetailsSettings.LocalCurrency, {{!SQL.SOURCE}}.UserDetailsSettings.Role, {{!SQL.SOURCE}}.UserDetailsSettings.TelephoneNumber, {{!SQL.SOURCE}}.UserDetailsSettings.TokenId, 
                         {{!SQL.SOURCE}}.UserDetailsSettings.TradingEntity AS Entity, {{!SQL.SOURCE}}.UserDetailsSettings.UserCode
FROM            {{!SQL.SOURCE}}.Usr INNER JOIN
                         {{!SQL.SOURCE}}.UserDetailsSettings ON {{!SQL.SOURCE}}.Usr.UserName = {{!SQL.SOURCE}}.UserDetailsSettings.UserName
WHERE        ({{!SQL.SOURCE}}.Usr.InternalDeleted IS NULL) AND ({{!SQL.SOURCE}}.UserDetailsSettings.InternalDeleted IS NULL) AND ({{!SQL.SOURCE}}.Usr.Type = 'CORE') AND (NOT ({{!SQL.SOURCE}}.UserDetailsSettings.TradingEntity IS NULL)) AND ({{!SQL.SOURCE}}.UserDetailsSettings.Enabled = 'true') 
                         AND (NOT ({{!SQL.SOURCE}}.Usr.TradingEntity IS NULL))
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane1', @value=N'[0E232FF0-B466-11cf-A24F-00AA00A3EFFF, 1.00]
Begin DesignProperties = 
   Begin PaneConfigurations = 
      Begin PaneConfiguration = 0
         NumPanes = 4
         Configuration = "(H (1[12] 4[50] 2[14] 3) )"
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
         Begin Table = "Usr (RG)"
            Begin Extent = 
               Top = 6
               Left = 38
               Bottom = 136
               Right = 243
            End
            DisplayFlags = 280
            TopColumn = 0
         End
         Begin Table = "UserDetailsSettings (RG)"
            Begin Extent = 
               Top = 6
               Left = 281
               Bottom = 136
               Right = 514
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
      Begin ColumnWidths = 16
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
         Width = 2115
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
' , @level0type=N'SCHEMA',@level0name=N'SRS', @level1type=N'VIEW',@level1name=N'sienaOwner'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPaneCount', @value=1 , @level0type=N'SCHEMA',@level0name=N'SRS', @level1type=N'VIEW',@level1name=N'sienaOwner'
GO
