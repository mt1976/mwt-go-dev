USE [ReportingDb_sal_prd_demo_sys-3]
GO
/****** Object:  View [{{!SQL.SCHEMA}}].[sienaContactStream]    Script Date: 04/12/2021 18:17:08 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaContactStream]
AS
SELECT        CM.StreamStatus.description AS Status, CM.StreamStatusAction.actionId AS ActionID, CM.StreamStatusAction.allowed AS Allowed, CM.StreamType.description AS Campaign, CM.Stream.streamId AS StreamID, 
                         CM.Stream.contactId AS ContactID, CM.Stream.usr AS Usr, CM.Stream.description AS Description, CM.Stream.streamTypeId AS StreamTypeID, CM.Stream.streamStatusId AS StreamStatusID, 
                         CM.Stream.openedDateTime AS Opened, CM.Stream.closedDateTime AS Closed
FROM            CM.StreamType INNER JOIN
                         CM.Stream ON CM.StreamType.typeId = CM.Stream.streamTypeId INNER JOIN
                         CM.StreamStatusAction INNER JOIN
                         CM.StreamStatus ON CM.StreamStatusAction.streamStatusId = CM.StreamStatus.statusId ON CM.Stream.streamStatusId = CM.StreamStatus.statusId
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane1', @value=N'[0E232FF0-B466-11cf-A24F-00AA00A3EFFF, 1.00]
Begin DesignProperties = 
   Begin PaneConfigurations = 
      Begin PaneConfiguration = 0
         NumPanes = 4
         Configuration = "(H (1[43] 4[24] 2[3] 3) )"
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
         Begin Table = "Stream (CM)"
            Begin Extent = 
               Top = 6
               Left = 38
               Bottom = 195
               Right = 217
            End
            DisplayFlags = 280
            TopColumn = 3
         End
         Begin Table = "StreamStatus (CM)"
            Begin Extent = 
               Top = 6
               Left = 255
               Bottom = 160
               Right = 425
            End
            DisplayFlags = 280
            TopColumn = 0
         End
         Begin Table = "StreamStatusAction (CM)"
            Begin Extent = 
               Top = 6
               Left = 463
               Bottom = 199
               Right = 633
            End
            DisplayFlags = 280
            TopColumn = 0
         End
         Begin Table = "StreamType (CM)"
            Begin Extent = 
               Top = 6
               Left = 671
               Bottom = 177
               Right = 841
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
         Width = 2865
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
         Width = 2475
         Width = 1500
      End
   End
   Begin CriteriaPane = 
      Begin ColumnWidths = 11
         Column = 1440
         Alias = 900
         Table = 1170
      ' , @level0type=N'SCHEMA',@level0name=N'SRS', @level1type=N'VIEW',@level1name=N'sienaContactStream'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane2', @value=N'   Output = 720
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
' , @level0type=N'SCHEMA',@level0name=N'SRS', @level1type=N'VIEW',@level1name=N'sienaContactStream'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPaneCount', @value=2 , @level0type=N'SCHEMA',@level0name=N'SRS', @level1type=N'VIEW',@level1name=N'sienaContactStream'
GO
