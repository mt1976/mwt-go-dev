USE [SRS]
GO
/****** Object:  Table [{{!SQL.DB}}].[{{!SQL.SCHEMA}}].[applicationStore]    Script Date: 10/05/2021 22:16:26 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [{{!SQL.DB}}].[{{!SQL.SCHEMA}}].[applicationStore](
	[_id] [int] IDENTITY(1,1) NOT NULL,
	[id] [nvarchar](max) NOT NULL,
	[name] [nvarchar](max) NULL,
	[level] [nvarchar](max) NULL,
	[number] [nvarchar](max) NULL,
	[signature] [nvarchar](max) NULL,
	[cacheupdate] [nvarchar](max) NULL,
	[_created] [nvarchar](max) NULL,
	[_who] [nvarchar](max) NULL,
	[_host] [nvarchar](max) NULL,
	[_updated] [nvarchar](max) NULL,
 CONSTRAINT [PK_applicationStore] PRIMARY KEY CLUSTERED
(
	[_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
