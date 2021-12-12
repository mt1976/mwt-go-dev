USE [SRS]
GO

/****** Object:  Table [dbo].[message]    Script Date: 10/12/2021 09:45:14 ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [dbo].[message](
	[Id] [bigint] IDENTITY(1,1) NOT NULL,
	[messageID] [nvarchar](max) NULL,
	[messageFormat] [nvarchar](max) NULL,
	[messageDeliveredTo] [nvarchar](max) NULL,
	[messageBody] [nvarchar](max) NULL,
	[messageFilename] [nvarchar](max) NULL,
	[messageLife] [nvarchar](max) NULL,
	[messageDate] [nvarchar](max) NULL,
	[messageTime] [nvarchar](max) NULL,
	[messageTimeout] [datetime2](7) NULL,
	[messageTimeoutAction] [nvarchar](max) NULL,
	[messageACKNAK] [nvarchar](max) NULL,
	[responseID] [nvarchar](max) NULL,
	[responseFilename] [nvarchar](max) NULL,
	[responseBody] [nvarchar](max) NULL,
	[responseDate] [nvarchar](max) NULL,
	[responseTime] [nvarchar](max) NULL,
	[responseAdditionalInfo] [nvarchar](max) NULL,
	[_created] [nvarchar](max) NULL,
	[_createdBy] [nvarchar](max) NULL,
	[_createdHost] [nvarchar](max) NULL,
	[_updated] [nvarchar](max) NULL,
	[_updatedBy] [nvarchar](max) NULL,
	[_updatedHost] [nvarchar](max) NULL
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO

EXEC sys.sp_addextendedproperty @name=N'MS_Description', @value=N'Holds incoming/outcoming message info' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'TABLE',@level1name=N'message'
GO