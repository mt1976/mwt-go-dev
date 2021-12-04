USE [SRS]
GO
/****** Object:  Table [dbo].[template]    Script Date: 04/12/2021 18:06:42 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[template](
	[_id] [int] IDENTITY(1,1) NOT NULL,
	[FIELD1] [nvarchar](max) NULL,
	[FIELD2] [nvarchar](max) NULL,
	[_created] [nvarchar](max) NULL,
	[_createdBy] [varchar](100) NULL,
	[_createdHost] [varchar](100) NULL,
	[_updated] [varchar](100) NULL,
	[_updatedHost] [varchar](100) NULL,
	[_updatedBy] [varchar](100) NULL,
	[ID] [nvarchar](max) NULL,
 CONSTRAINT [PK_template] PRIMARY KEY CLUSTERED 
(
	[_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
