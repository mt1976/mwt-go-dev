/****** Object:  View [{{!SQL.SCHEMA}}].[sienaSector]    Script Date: 09/05/2021 14:36:12 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO


CREATE VIEW [{{!SQL.SCHEMA}}].[sienaSector]
AS
SELECT        Code, Name
FROM            {{!SQL.SOURCE}}.Sector
WHERE        (InternalDeleted IS NULL)
GO
