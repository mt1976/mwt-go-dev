USE [SRS]
GO
/****** Object:  Table [dbo].[brandStore]    Script Date: 24/11/2021 19:42:59 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[brandStore](
	[_id] [int] IDENTITY(0,1) NOT NULL,
	[Id] [nvarchar](max) NOT NULL,
	[Name] [nvarchar](max) NOT NULL,
	[_createdBy] [nvarchar](max) NULL,
	[_createdHost] [nvarchar](max) NULL,
	[_updatedBy] [nvarchar](max) NULL,
	[_updatedHost] [nvarchar](max) NULL,
	[_updated] [nvarchar](max) NULL,
	[_created] [nvarchar](max) NULL
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
