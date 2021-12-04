USE [SRS]
GO
/****** Object:  Table [dbo].[systemStore]    Script Date: 04/12/2021 18:06:42 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[systemStore](
	[_id] [int] IDENTITY(1,1) NOT NULL,
	[Id] [nvarchar](max) NOT NULL,
	[Name] [nvarchar](max) NULL,
	[Staticin] [nvarchar](max) NULL,
	[Staticout] [nvarchar](max) NULL,
	[Txnin] [nvarchar](max) NULL,
	[Txnout] [nvarchar](max) NULL,
	[Fundscheckin] [nvarchar](max) NULL,
	[Fundscheckout] [nvarchar](max) NULL,
	[_created] [nvarchar](max) NULL,
	[_who] [nvarchar](max) NULL,
	[_host] [nvarchar](max) NULL,
	[_updated] [nvarchar](max) NULL,
	[_createdBy] [nvarchar](max) NULL,
	[_createdHost] [nvarchar](max) NULL,
	[_updatedBy] [nvarchar](max) NULL,
	[_updatedHost] [nvarchar](max) NULL,
	[SWIFTin] [nvarchar](max) NULL,
	[SWIFTout] [nvarchar](max) NULL,
 CONSTRAINT [PK_systemStore] PRIMARY KEY CLUSTERED 
(
	[_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
