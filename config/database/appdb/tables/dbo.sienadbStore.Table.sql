USE [SRS]
GO
/****** Object:  Table [dbo].[sienadbStore]    Script Date: 04/12/2021 18:06:42 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[sienadbStore](
	[_id] [int] IDENTITY(1,1) NOT NULL,
	[ID] [nvarchar](max) NOT NULL,
	[server] [nvarchar](max) NULL,
	[port] [nvarchar](max) NULL,
	[user] [nvarchar](max) NULL,
	[password] [nvarchar](max) NULL,
	[database] [nvarchar](max) NULL,
	[schema] [nvarchar](max) NULL,
	[_created] [nvarchar](max) NULL,
	[_who] [nvarchar](max) NULL,
	[_host] [nvarchar](max) NULL,
	[_updated] [nvarchar](max) NULL,
	[active] [nvarchar](max) NULL,
	[_createdBy] [nvarchar](max) NULL,
	[_createdHost] [nvarchar](max) NULL,
	[_updatedBy] [nvarchar](max) NULL,
	[_updatedHost] [nvarchar](max) NULL
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
