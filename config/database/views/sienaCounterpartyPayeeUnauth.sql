/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCounterpartyPayeeUnauth]    Script Date: 09/05/2021 14:36:12 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaCounterpartyPayeeUnauth]
AS
SELECT        {{!SQL.SOURCE}}.UnauthorisedPayee.KeyCounterpartyFirm, {{!SQL.SOURCE}}.UnauthorisedPayee.KeyCounterpartyCentre, {{!SQL.SOURCE}}.UnauthorisedPayee.KeyName, {{!SQL.SOURCE}}.UnauthorisedPayee.Currency, {{!SQL.SOURCE}}.UnauthorisedPayee.FullName, 
                         {{!SQL.SOURCE}}.UnauthorisedPayee.Address, {{!SQL.SOURCE}}.UnauthorisedPayee.PhoneNo, {{!SQL.SOURCE}}.UnauthorisedPayee.Country, {{!SQL.SOURCE}}.UnauthorisedPayee.Bic, {{!SQL.SOURCE}}.UnauthorisedPayee.Iban, {{!SQL.SOURCE}}.UnauthorisedPayee.AccountNo, 
                         {{!SQL.SOURCE}}.UnauthorisedPayee.FedWireNo, {{!SQL.SOURCE}}.UnauthorisedPayee.SortCode, {{!SQL.SOURCE}}.UnauthorisedPayee.BankName, {{!SQL.SOURCE}}.UnauthorisedPayee.BankPinCode, {{!SQL.SOURCE}}.UnauthorisedPayee.BankAddress, {{!SQL.SOURCE}}.UnauthorisedPayee.Reason, 
                         {{!SQL.SOURCE}}.UnauthorisedPayee.BankSettlementAcct, {{!SQL.SOURCE}}.UnauthorisedPayee.IntermediaryBankName, {{!SQL.SOURCE}}.UnauthorisedPayee.IntermediaryBankAddress, {{!SQL.SOURCE}}.UnauthorisedPayee.IntermediaryBankSwiftBic, 
                         {{!SQL.SOURCE}}.UnauthorisedPayee.IntermediaryAcctNo, {{!SQL.SOURCE}}.UnauthorisedPayee.IntermediaryIban, {{!SQL.SOURCE}}.UnauthorisedPayee.NostroInternalAccountNumber, {{!SQL.SOURCE}}.UnauthorisedPayee.NostroBIC, {{!SQL.SOURCE}}.UnauthorisedPayee.Usr, 
                         {{!SQL.SOURCE}}.UnauthorisedPayee.Status, {{!SQL.SOURCE}}.UnauthorisedPayee.CreationDate, {{!SQL.SOURCE}}.UnauthorisedPayee.AmendedDate, {{!SQL.SOURCE}}.Country.Name AS CountryName, {{!SQL.SOURCE}}.Currency.Name AS CurrencyName
FROM            {{!SQL.SOURCE}}.UnauthorisedPayee INNER JOIN
                         {{!SQL.SOURCE}}.Country ON {{!SQL.SOURCE}}.UnauthorisedPayee.Country = {{!SQL.SOURCE}}.Country.Code INNER JOIN
                         {{!SQL.SOURCE}}.Currency ON {{!SQL.SOURCE}}.UnauthorisedPayee.Currency = {{!SQL.SOURCE}}.Currency.Code
WHERE        ({{!SQL.SOURCE}}.UnauthorisedPayee.InternalDeleted IS NULL)
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
         Begin Table = "UnauthorisedPayee"
            Begin Extent = 
               Top = 6
               Left = 38
               Bottom = 136
               Right = 292
            End
            DisplayFlags = 280
            TopColumn = 1
         End
         Begin Table = "Currency"
            Begin Extent = 
               Top = 6
               Left = 330
               Bottom = 136
               Right = 702
            End
            DisplayFlags = 280
            TopColumn = 0
         End
         Begin Table = "Country"
            Begin Extent = 
               Top = 166
               Left = 315
               Bottom = 296
               Right = 520
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
      Begin ColumnWidths = 31
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
         Column = 1440
         Alias = 900
         Table ' , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaCounterpartyPayeeUnauth'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane2', @value=N'= 1170
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
' , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaCounterpartyPayeeUnauth'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPaneCount', @value=2 , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaCounterpartyPayeeUnauth'
GO
