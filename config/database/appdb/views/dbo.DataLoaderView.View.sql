USE [SRS-2]
GO
/****** Object:  View [{{!SQL.DB}}].[{{!SQL.SCHEMA}}].[ratesDataView]    Script Date: 11/05/2021 14:31:24 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[DataLoaderView] AS
SELECT        [{{!SQL.SOURCE}}].dbo.loaderDataStore.row, [{{!SQL.SOURCE}}].dbo.loaderMapStore.name AS Field, [{{!SQL.SOURCE}}].dbo.loaderDataStore.value, [{{!SQL.SOURCE}}].dbo.loaderStore.name AS DataMap
FROM            [{{!SQL.SOURCE}}].dbo.loaderDataStore INNER JOIN
                         [{{!SQL.SOURCE}}].dbo.loaderMapStore ON [{{!SQL.SOURCE}}].dbo.loaderDataStore.map = [{{!SQL.SOURCE}}].dbo.loaderMapStore.id INNER JOIN
                         [{{!SQL.SOURCE}}].dbo.loaderStore ON [{{!SQL.SOURCE}}].dbo.loaderDataStore.loader = [{{!SQL.SOURCE}}].dbo.loaderStore.id
GO
