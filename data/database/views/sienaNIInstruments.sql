/****** Object:  View [{{!SQL.SCHEMA}}].[sienaNIInstruments]    Script Date: 09/05/2021 14:36:12 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaNIInstruments]
AS
SELECT        {{!SQL.SOURCE}}.NegotiableInstrument.NIName, {{!SQL.SOURCE}}.NegotiableInstrument.Status, {{!SQL.SOURCE}}.NegotiableInstrument.Type, {{!SQL.SOURCE}}.NegotiableInstrument.CurrencyCode, {{!SQL.SOURCE}}.NegotiableInstrument.IssuerFirm, {{!SQL.SOURCE}}.NegotiableInstrument.IssuerCentre, 
                         {{!SQL.SOURCE}}.NegotiableInstrument.AcceptorFirm, {{!SQL.SOURCE}}.NegotiableInstrument.AcceptorCentre, {{!SQL.SOURCE}}.NegotiableInstrument.IssueDate, {{!SQL.SOURCE}}.NegotiableInstrument.MaturityDate, {{!SQL.SOURCE}}.NegotiableInstrument.RateType, 
                         {{!SQL.SOURCE}}.NegotiableInstrument.CouponRate, {{!SQL.SOURCE}}.NegotiableInstrument.InterestBasis, {{!SQL.SOURCE}}.NegotiableInstrument.InterestFrequency, {{!SQL.SOURCE}}.NegotiableInstrument.ClearingSystem, {{!SQL.SOURCE}}.NegotiableInstrument.ParentDealType, 
                         {{!SQL.SOURCE}}.NegotiableInstrument.SecurityID, {{!SQL.SOURCE}}.NegotiableInstrument.Encumbered, {{!SQL.SOURCE}}.NegotiableInstrument.Marketable, {{!SQL.SOURCE}}.NegotiableInstrument.Hypothecated, {{!SQL.SOURCE}}.NegotiableInstrument.RepoFactor, 
                         {{!SQL.SOURCE}}.NegotiableInstrument.LotSize, {{!SQL.SOURCE}}.Currency.Name AS CurrencyName, {{!SQL.SOURCE}}.Firm.FullName AS IssuerFirmName, {{!SQL.SOURCE}}.Centre.FullName AS IssuerCentreName
FROM            {{!SQL.SOURCE}}.NegotiableInstrument LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Firm ON {{!SQL.SOURCE}}.NegotiableInstrument.IssuerFirm = {{!SQL.SOURCE}}.Firm.FirmName LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Centre ON {{!SQL.SOURCE}}.NegotiableInstrument.IssuerCentre = {{!SQL.SOURCE}}.Centre.ShortName LEFT OUTER JOIN
                         {{!SQL.SOURCE}}.Currency ON {{!SQL.SOURCE}}.NegotiableInstrument.CurrencyCode = {{!SQL.SOURCE}}.Currency.Code
WHERE        ({{!SQL.SOURCE}}.NegotiableInstrument.InternalDeleted IS NULL)
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane1', @value=N'[0E232FF0-B466-11cf-A24F-00AA00A3EFFF, 1.00]
Begin DesignProperties = 
   Begin PaneConfigurations = 
      Begin PaneConfiguration = 0
         NumPanes = 4
         Configuration = "(H (1[41] 4[20] 2[20] 3) )"
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
         Begin Table = "NegotiableInstrument"
            Begin Extent = 
               Top = 6
               Left = 38
               Bottom = 226
               Right = 198
            End
            DisplayFlags = 280
            TopColumn = 10
         End
         Begin Table = "Currency"
            Begin Extent = 
               Top = 6
               Left = 236
               Bottom = 136
               Right = 608
            End
            DisplayFlags = 344
            TopColumn = 0
         End
         Begin Table = "Firm"
            Begin Extent = 
               Top = 41
               Left = 332
               Bottom = 171
               Right = 537
            End
            DisplayFlags = 344
            TopColumn = 0
         End
         Begin Table = "Centre"
            Begin Extent = 
               Top = 65
               Left = 278
               Bottom = 195
               Right = 483
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
      Begin ColumnWidths = 27
         Width = 284
         Width = 2700
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
         Width = 1650
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width = 1500
         Width' , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaNIInstruments'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane2', @value=N' = 1500
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
' , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaNIInstruments'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPaneCount', @value=2 , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaNIInstruments'
GO
