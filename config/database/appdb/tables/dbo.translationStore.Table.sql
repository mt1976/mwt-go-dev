USE [SRS]
GO
/****** Object:  Table [dbo].[translationStore]    Script Date: 24/11/2021 19:43:00 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[translationStore](
	[_id] [int] IDENTITY(1,1) NOT NULL,
	[Id] [nvarchar](max) NULL,
	[Class] [nvarchar](max) NULL,
	[Message] [nvarchar](max) NULL,
	[Translation] [nvarchar](max) NULL,
	[_created] [nvarchar](max) NULL,
	[_who] [nvarchar](max) NULL,
	[_host] [nvarchar](max) NULL,
	[_updated] [nvarchar](max) NULL,
	[_createdBy] [nvarchar](max) NULL,
	[_createdHost] [nvarchar](max) NULL,
	[_updatedBy] [nvarchar](max) NULL,
	[_updatedHost] [nvarchar](max) NULL,
 CONSTRAINT [PK_translationStore] PRIMARY KEY CLUSTERED 
(
	[_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
