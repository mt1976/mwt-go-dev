USE [SRS]
GO
/****** Object:  Table [dbo].[lseGiltsDataStore]    Script Date: 04/12/2021 18:06:41 ******/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[lseGiltsDataStore](
	[_id] [int] IDENTITY(1,1) NOT NULL,
	[id] [nvarchar](max) NULL,
	[longName] [nvarchar](max) NULL,
	[isin] [nvarchar](max) NULL,
	[tidm] [nvarchar](max) NULL,
	[sedol] [nvarchar](max) NULL,
	[issueDate] [nvarchar](max) NULL,
	[maturityDate] [nvarchar](max) NULL,
	[couponValue] [nvarchar](max) NULL,
	[couponType] [nvarchar](max) NULL,
	[segment] [nvarchar](max) NULL,
	[sector] [nvarchar](max) NULL,
	[codeConventionCalculateAccrual] [nvarchar](max) NULL,
	[minimumDenomination] [nvarchar](max) NULL,
	[denominationCurrency] [nvarchar](max) NULL,
	[tradingCurrency] [nvarchar](max) NULL,
	[type] [nvarchar](max) NULL,
	[flatYield] [nvarchar](max) NULL,
	[paymentCouponDate] [nvarchar](max) NULL,
	[periodOfCoupon] [nvarchar](max) NULL,
	[exCouponDate] [nvarchar](max) NULL,
	[dateOfIndexInflation] [nvarchar](max) NULL,
	[unitOfQuotation] [nvarchar](max) NULL,
	[_created] [nvarchar](max) NULL,
	[_who] [nvarchar](max) NULL,
	[_host] [nvarchar](max) NULL,
	[_updated] [nvarchar](max) NULL,
	[issuer] [nvarchar](max) NULL,
	[issueAmount] [nvarchar](max) NULL,
	[runningYield] [nvarchar](max) NULL,
	[LEI] [nvarchar](max) NULL,
	[CUSIP] [nvarchar](max) NULL,
	[_updatedHost] [nvarchar](max) NULL,
	[_createdBy] [nvarchar](max) NULL,
	[_createdHost] [nvarchar](max) NULL,
	[_updatedBy] [nvarchar](max) NULL,
 CONSTRAINT [PK_lseGiltsDataStore] PRIMARY KEY CLUSTERED 
(
	[_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY] TEXTIMAGE_ON [PRIMARY]
GO
