USE [SRS-2]
GO
/****** Object:  View [{{!SQL.DB}}].[{{!SQL.SCHEMA}}].[ratesDataView]    Script Date: 11/05/2021 14:31:24 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[DataLoaderView] AS
SELECT        [{{!SQL.DB}}].dbo.loaderDataStore.row, [{{!SQL.DB}}].dbo.loaderMapStore.name AS Field, [{{!SQL.DB}}].dbo.loaderDataStore.value, [{{!SQL.DB}}].dbo.loaderStore.name AS DataMap
FROM            [{{!SQL.DB}}].dbo.loaderDataStore INNER JOIN
                         [{{!SQL.DB}}].dbo.loaderMapStore ON [{{!SQL.DB}}].dbo.loaderDataStore.map = [{{!SQL.DB}}].dbo.loaderMapStore.id INNER JOIN
                         [{{!SQL.DB}}].dbo.loaderStore ON [{{!SQL.DB}}].dbo.loaderDataStore.loader = [{{!SQL.DB}}].dbo.loaderStore.id
GO
