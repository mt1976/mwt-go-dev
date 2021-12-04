USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCounterparty]    Script Date: 04/12/2021 18:16:52 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaCounterparty]
AS
SELECT        {{!SQL.SOURCE}}.Counterparty.NameCentre, {{!SQL.SOURCE}}.Counterparty.NameFirm, {{!SQL.SOURCE}}.Counterparty.FullName, {{!SQL.SOURCE}}.Counterparty.TelephoneNumber, {{!SQL.SOURCE}}.Counterparty.EmailAddress, {{!SQL.SOURCE}}.Counterparty.CustomerType, {{!SQL.SOURCE}}.Counterparty.AccountOfficer, 
                         {{!SQL.SOURCE}}.Counterparty.CountryCode, {{!SQL.SOURCE}}.Counterparty.SectorCode, {{!SQL.SOURCE}}.Counterparty.CpartyGroupName, {{!SQL.SOURCE}}.Counterparty.Notes, {{!SQL.SOURCE}}.Counterparty.Owner, {{!SQL.SOURCE}}.Counterparty.Authorised, {{!SQL.SOURCE}}.Firm.FullName AS NameFirmName, 
                         {{!SQL.SOURCE}}.Centre.FullName AS NameCentreName, {{!SQL.SOURCE}}.Country.Name AS CountryCodeName, {{!SQL.SOURCE}}.Sector.Name AS SectorCodeName, { fn CONCAT({{!SQL.SOURCE}}.Counterparty.NameFirm, { fn CONCAT('|', {{!SQL.SOURCE}}.Counterparty.NameCentre) }) 
                         } AS CompID
FROM            {{!SQL.SOURCE}}.Counterparty INNER JOIN
                         {{!SQL.SOURCE}}.Firm ON {{!SQL.SOURCE}}.Counterparty.NameFirm = {{!SQL.SOURCE}}.Firm.FirmName INNER JOIN
                         {{!SQL.SOURCE}}.Centre ON {{!SQL.SOURCE}}.Counterparty.NameCentre = {{!SQL.SOURCE}}.Centre.ShortName INNER JOIN
                         {{!SQL.SOURCE}}.Country ON {{!SQL.SOURCE}}.Counterparty.CountryCode = {{!SQL.SOURCE}}.Country.Code INNER JOIN
                         {{!SQL.SOURCE}}.Sector ON {{!SQL.SOURCE}}.Counterparty.SectorCode = {{!SQL.SOURCE}}.Sector.Code
WHERE        ({{!SQL.SOURCE}}.Counterparty.InternalDeleted IS NULL)
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane1', @value=N'[0E232FF0-B466-11cf-A24F-00AA00A3EFFF, 1.00]
Begin DesignProperties = 
   Begin PaneConfigurations = 
      Begin PaneConfiguration = 0
         NumPanes = 4
         Configuration = "(H (1[40] 4[20] 2[20] 3) )"
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
         Begin Table = "Counterparty (RG)"
            Begin Extent = 
               Top = 6
               Left = 38
               Bottom = 136
               Right = 371
            End
            DisplayFlags = 280
            TopColumn = 0
         End
         Begin Table = "Firm (RG)"
            Begin Extent = 
               Top = 138
               Left = 38
               Bottom = 268
               Right = 243
            End
            DisplayFlags = 280
            TopColumn = 0
         End
         Begin Table = "Centre (RG)"
            Begin Extent = 
               Top = 270
               Left = 38
               Bottom = 400
               Right = 243
            End
            DisplayFlags = 280
            TopColumn = 0
         End
         Begin Table = "Country (RG)"
            Begin Extent = 
               Top = 402
               Left = 38
               Bottom = 532
               Right = 243
            End
            DisplayFlags = 280
            TopColumn = 0
         End
         Begin Table = "Sector (RG)"
            Begin Extent = 
               Top = 534
               Left = 38
               Bottom = 664
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
      Begin ColumnWidths = 18
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
         Widt' , @level0type=N'SCHEMA',@level0name=N'SRS', @level1type=N'VIEW',@level1name=N'sienaCounterparty'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane2', @value=N'h = 1500
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
' , @level0type=N'SCHEMA',@level0name=N'SRS', @level1type=N'VIEW',@level1name=N'sienaCounterparty'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPaneCount', @value=2 , @level0type=N'SCHEMA',@level0name=N'SRS', @level1type=N'VIEW',@level1name=N'sienaCounterparty'
GO
