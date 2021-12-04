USE [SRS]
GO
/****** Object:  Table [dbo].[sessionStore]    Script Date: 04/12/2021 18:06:42 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[sessionStore](
	[_id] [int] IDENTITY(1,1) NOT NULL,
	[Apptoken] [nvarchar](max) NOT NULL,
	[Createdate] [nvarchar](max) NULL,
	[Createtime] [nvarchar](max) NULL,
	[Uniqueid] [nvarchar](max) NULL,
	[Sessiontoken] [nvarchar](max) NULL,
	[Username] [nvarchar](max) NULL,
	[Password] [nvarchar](max) NULL,
	[Userip] [nvarchar](max) NULL,
	[Userhost] [nvarchar](max) NULL,
	[Appip] [nvarchar](max) NULL,
	[Apphost] [nvarchar](max) NULL,
	[Issued] [nvarchar](max) NULL,
	[Expiry] [nvarchar](max) NULL,
	[Expiryraw] [nvarchar](max) NULL,
	[Brand] [nvarchar](max) NULL,
	[_created] [nvarchar](max) NULL,
	[_who] [nvarchar](max) NULL,
	[_host] [nvarchar](max) NULL,
	[_updated] [nvarchar](max) NULL,
	[Id] [nvarchar](max) NULL,
	[Expires] [datetime2](7) NULL,
	[_createdBy] [nvarchar](max) NULL,
	[_createdHost] [nvarchar](max) NULL,
	[_updatedBy] [nvarchar](max) NULL,
	[_updatedHost] [nvarchar](max) NULL,
	[SessionRole] [nvarchar](max) NULL,
 CONSTRAINT [PK_sessionStore] PRIMARY KEY CLUSTERED 
(
	[_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
