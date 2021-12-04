USE [SRS]
GO
/****** Object:  Table [dbo].[loaderStore]    Script Date: 04/12/2021 18:06:41 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[loaderStore](
	[_id] [int] IDENTITY(1,1) NOT NULL,
	[id] [nvarchar](max) NOT NULL,
	[name] [nvarchar](max) NULL,
	[description] [nvarchar](max) NULL,
	[filename] [nvarchar](max) NULL,
	[lastrun] [nvarchar](max) NULL,
	[_created] [nvarchar](max) NULL,
	[_who] [nvarchar](max) NULL,
	[_host] [nvarchar](max) NULL,
	[_updated] [nvarchar](max) NULL,
	[type] [nvarchar](max) NULL,
	[instance] [nvarchar](max) NULL,
	[extension] [nvarchar](max) NULL,
	[_createdBy] [nvarchar](max) NULL,
	[_updatedHost] [nvarchar](max) NULL,
	[_updatedBy] [nvarchar](max) NULL,
	[_createdHost] [nvarchar](max) NULL,
 CONSTRAINT [PK_loaderStore] PRIMARY KEY CLUSTERED 
(
	[_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
