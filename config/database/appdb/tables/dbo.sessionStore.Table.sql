USE [SRS]
GO
/****** Object:  Table [{{!SQL.DB}}].[{{!SQL.SCHEMA}}].[sessionStore]    Script Date: 10/05/2021 22:16:26 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [{{!SQL.DB}}].[{{!SQL.SCHEMA}}].[sessionStore](
	[_id] [int] IDENTITY(1,1) NOT NULL,
	[apptoken] [nvarchar](max) NOT NULL,
	[createdate] [nvarchar](max) NULL,
	[createtime] [nvarchar](max) NULL,
	[uniqueid] [nvarchar](max) NULL,
	[sessiontoken] [nvarchar](max) NULL,
	[username] [nvarchar](max) NULL,
	[password] [nvarchar](max) NULL,
	[userip] [nvarchar](max) NULL,
	[userhost] [nvarchar](max) NULL,
	[appip] [nvarchar](max) NULL,
	[apphost] [nvarchar](max) NULL,
	[issued] [nvarchar](max) NULL,
	[expiry] [nvarchar](max) NULL,
	[expiryraw] [nvarchar](max) NULL,
	[role] [nvarchar](max) NULL,
	[brand] [nvarchar](max) NULL,
	[_created] [nvarchar](max) NULL,
	[_who] [nvarchar](max) NULL,
	[_host] [nvarchar](max) NULL,
	[_updated] [nvarchar](max) NULL,
	[id] [nvarchar](max) NULL,
	[expires] [datetime2](7) NULL,
 CONSTRAINT [PK_sessionStore] PRIMARY KEY CLUSTERED 
(
	[_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
