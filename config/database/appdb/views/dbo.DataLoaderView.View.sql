USE [SRS-2]
GO
/****** Object:  View [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].[ratesDataView]    Script Date: 11/05/2021 14:31:24 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE VIEW [{{!SQL.SCHEMA}}].[DataLoaderView] AS
SELECT        [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].loaderDataStore.row, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].loaderMapStore.name AS Field, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].loaderDataStore.value, [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].loaderStore.name AS DataMap
FROM            [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].loaderDataStore INNER JOIN
                         [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].loaderMapStore ON [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].loaderDataStore.map = [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].loaderMapStore.id INNER JOIN
                         [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].loaderStore ON [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].loaderDataStore.loader = [{{!SQL.SOURCE}}].[{{!SQL.SCHEMA}}].loaderStore.id
GO
