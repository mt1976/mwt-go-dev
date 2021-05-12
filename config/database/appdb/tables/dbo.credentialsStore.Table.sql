USE [SRS]
GO
/****** Object:  Table [{{!SQL.DB}}].[{{!SQL.SCHEMA}}].[credentialsStore]    Script Date: 10/05/2021 22:16:26 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [{{!SQL.DB}}].[{{!SQL.SCHEMA}}].[credentialsStore](
	[_id] [int] IDENTITY(1,1) NOT NULL,
	[id] [nvarchar](max) NOT NULL,
	[username] [nvarchar](max) NOT NULL,
	[password] [nvarchar](max) NULL,
	[firstname] [nvarchar](max) NULL,
	[lastname] [nvarchar](max) NULL,
	[knownas] [nvarchar](max) NULL,
	[email] [nvarchar](max) NULL,
	[issued] [nvarchar](max) NULL,
	[expiry] [nvarchar](max) NULL,
	[role] [nvarchar](max) NULL,
	[brand] [nvarchar](max) NULL,
	[_created] [nvarchar](max) NULL,
	[_who] [nvarchar](max) NULL,
	[_host] [nvarchar](max) NULL,
	[_updated] [nvarchar](max) NULL,
 CONSTRAINT [PK_credentialsStore] PRIMARY KEY CLUSTERED 
(
	[_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
