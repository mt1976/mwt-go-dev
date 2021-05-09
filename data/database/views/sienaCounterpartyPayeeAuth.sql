/****** Object:  View [{{!SQL.SCHEMA}}].[sienaCounterpartyPayeeAuth]    Script Date: 09/05/2021 14:36:12 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[sienaCounterpartyPayeeAuth]
AS
SELECT        {{!SQL.SOURCE}}.Payee.KeyCounterpartyFirm, {{!SQL.SOURCE}}.Payee.KeyCounterpartyCentre, {{!SQL.SOURCE}}.Payee.KeyCurrency, {{!SQL.SOURCE}}.Payee.KeyName, {{!SQL.SOURCE}}.Payee.KeyNumber, {{!SQL.SOURCE}}.Payee.KeyDirection, {{!SQL.SOURCE}}.Payee.KeyType, {{!SQL.SOURCE}}.Payee.FullName, 
                         {{!SQL.SOURCE}}.Payee.Address, {{!SQL.SOURCE}}.Payee.PhoneNo, {{!SQL.SOURCE}}.Payee.Country, {{!SQL.SOURCE}}.Payee.Bic, {{!SQL.SOURCE}}.Payee.Iban, {{!SQL.SOURCE}}.Payee.AccountNo, {{!SQL.SOURCE}}.Payee.FedWireNo, {{!SQL.SOURCE}}.Payee.SortCode, {{!SQL.SOURCE}}.Payee.BankName, {{!SQL.SOURCE}}.Payee.BankPinCode, 
                         {{!SQL.SOURCE}}.Payee.BankAddress, {{!SQL.SOURCE}}.Payee.Reason, {{!SQL.SOURCE}}.Payee.BankSettlementAcct, {{!SQL.SOURCE}}.Payee.IntermediaryBankName, {{!SQL.SOURCE}}.Payee.IntermediaryBankAddress, {{!SQL.SOURCE}}.Payee.IntermediaryBankSwiftBic, {{!SQL.SOURCE}}.Payee.IntermediaryAcctNo, 
                         {{!SQL.SOURCE}}.Payee.IntermediaryIban, {{!SQL.SOURCE}}.Payee.NostroInternalAccountNumber, {{!SQL.SOURCE}}.Payee.NostroBIC, {{!SQL.SOURCE}}.Country.Name AS CountryName, {{!SQL.SOURCE}}.Currency.Name AS CurrencyName
FROM            {{!SQL.SOURCE}}.Payee INNER JOIN
                         {{!SQL.SOURCE}}.Country ON {{!SQL.SOURCE}}.Payee.Country = {{!SQL.SOURCE}}.Country.Code INNER JOIN
                         {{!SQL.SOURCE}}.Currency ON {{!SQL.SOURCE}}.Payee.KeyCurrency = {{!SQL.SOURCE}}.Currency.Code
WHERE        ({{!SQL.SOURCE}}.Payee.InternalDeleted IS NULL)
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane1', @value=N'[0E232FF0-B466-11cf-A24F-00AA00A3EFFF, 1.00]
Begin DesignProperties = 
   Begin PaneConfigurations = 
      Begin PaneConfiguration = 0
         NumPanes = 4
         Configuration = "(H (1[38] 4[19] 2[5] 3) )"
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
         Begin Table = "Payee"
            Begin Extent = 
               Top = 6
               Left = 38
               Bottom = 217
               Right = 292
            End
            DisplayFlags = 280
            TopColumn = 0
         End
         Begin Table = "Country"
            Begin Extent = 
               Top = 6
               Left = 330
               Bottom = 136
               Right = 535
            End
            DisplayFlags = 344
            TopColumn = 0
         End
         Begin Table = "Currency"
            Begin Extent = 
               Top = 6
               Left = 573
               Bottom = 136
               Right = 945
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
         Table = 1170
       ' , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaCounterpartyPayeeAuth'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPane2', @value=N'  Output = 720
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
' , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaCounterpartyPayeeAuth'
GO
EXEC sys.sp_addextendedproperty @name=N'MS_DiagramPaneCount', @value=2 , @level0type=N'SCHEMA',@level0name=N'{{!SQL.SCHEMA}}', @level1type=N'VIEW',@level1name=N'sienaCounterpartyPayeeAuth'
GO
