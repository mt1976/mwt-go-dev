package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/negotiableinstrument.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : NegotiableInstrument (negotiableinstrument)
// Endpoint 	        : NegotiableInstrument (Id)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 12/12/2021 at 16:13:16
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type NegotiableInstrument struct {

SYSId        string
Id        string
LongName        string
Isin        string
Tidm        string
Sedol        string
IssueDate        string
MaturityDate        string
CouponValue        string
CouponType        string
Segment        string
Sector        string
CodeConventionCalculateAccrual        string
MinimumDenomination        string
DenominationCurrency        string
TradingCurrency        string
Type        string
FlatYield        string
PaymentCouponDate        string
PeriodOfCoupon        string
ExCouponDate        string
DateOfIndexInflation        string
UnitOfQuotation        string
SYSCreated        string
SYSWho        string
SYSHost        string
SYSUpdated        string
Issuer        string
IssueAmount        string
RunningYield        string
LEI        string
CUSIP        string
SYSUpdatedHost        string
SYSCreatedBy        string
SYSCreatedHost        string
SYSUpdatedBy        string

}

const (
	NegotiableInstrument_Title       = "Negotiable Instrument"
	NegotiableInstrument_SQLTable    = "lseGiltsDataStore"
	NegotiableInstrument_SQLSearchID = "Id"
	NegotiableInstrument_QueryString = "Id"
	///
	/// Handler Defintions
	///
	NegotiableInstrument_Template     = "NegotiableInstrument"
	NegotiableInstrument_TemplateList = "NegotiableInstrument_List"
	NegotiableInstrument_TemplateView = "NegotiableInstrument_View"
	NegotiableInstrument_TemplateEdit = "NegotiableInstrument_Edit"
	NegotiableInstrument_TemplateNew  = "NegotiableInstrument_New"
	///
	/// Handler Monitor Paths
	///
	NegotiableInstrument_Path       = "/API/NegotiableInstrument/"
	NegotiableInstrument_PathList   = "/NegotiableInstrumentList/"
	NegotiableInstrument_PathView   = "/NegotiableInstrumentView/"
	NegotiableInstrument_PathEdit   = "/NegotiableInstrumentEdit/"
	NegotiableInstrument_PathNew    = "/NegotiableInstrumentNew/"
	NegotiableInstrument_PathSave   = "/NegotiableInstrumentSave/"
	NegotiableInstrument_PathDelete = "/NegotiableInstrumentDelete/"
	///
	/// SQL Field Definitions
	///
	NegotiableInstrument_SYSId   = "_id" // SYSId is a Int
	NegotiableInstrument_Id   = "id" // Id is a String
	NegotiableInstrument_LongName   = "longName" // LongName is a String
	NegotiableInstrument_Isin   = "isin" // Isin is a String
	NegotiableInstrument_Tidm   = "tidm" // Tidm is a String
	NegotiableInstrument_Sedol   = "sedol" // Sedol is a String
	NegotiableInstrument_IssueDate   = "issueDate" // IssueDate is a String
	NegotiableInstrument_MaturityDate   = "maturityDate" // MaturityDate is a String
	NegotiableInstrument_CouponValue   = "couponValue" // CouponValue is a String
	NegotiableInstrument_CouponType   = "couponType" // CouponType is a String
	NegotiableInstrument_Segment   = "segment" // Segment is a String
	NegotiableInstrument_Sector   = "sector" // Sector is a String
	NegotiableInstrument_CodeConventionCalculateAccrual   = "codeConventionCalculateAccrual" // CodeConventionCalculateAccrual is a String
	NegotiableInstrument_MinimumDenomination   = "minimumDenomination" // MinimumDenomination is a String
	NegotiableInstrument_DenominationCurrency   = "denominationCurrency" // DenominationCurrency is a String
	NegotiableInstrument_TradingCurrency   = "tradingCurrency" // TradingCurrency is a String
	NegotiableInstrument_Type   = "type" // Type is a String
	NegotiableInstrument_FlatYield   = "flatYield" // FlatYield is a String
	NegotiableInstrument_PaymentCouponDate   = "paymentCouponDate" // PaymentCouponDate is a String
	NegotiableInstrument_PeriodOfCoupon   = "periodOfCoupon" // PeriodOfCoupon is a String
	NegotiableInstrument_ExCouponDate   = "exCouponDate" // ExCouponDate is a String
	NegotiableInstrument_DateOfIndexInflation   = "dateOfIndexInflation" // DateOfIndexInflation is a String
	NegotiableInstrument_UnitOfQuotation   = "unitOfQuotation" // UnitOfQuotation is a String
	NegotiableInstrument_SYSCreated   = "_created" // SYSCreated is a String
	NegotiableInstrument_SYSWho   = "_who" // SYSWho is a String
	NegotiableInstrument_SYSHost   = "_host" // SYSHost is a String
	NegotiableInstrument_SYSUpdated   = "_updated" // SYSUpdated is a String
	NegotiableInstrument_Issuer   = "issuer" // Issuer is a String
	NegotiableInstrument_IssueAmount   = "issueAmount" // IssueAmount is a String
	NegotiableInstrument_RunningYield   = "runningYield" // RunningYield is a String
	NegotiableInstrument_LEI   = "LEI" // LEI is a String
	NegotiableInstrument_CUSIP   = "CUSIP" // CUSIP is a String
	NegotiableInstrument_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String
	NegotiableInstrument_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	NegotiableInstrument_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	NegotiableInstrument_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String

	/// Definitions End
)
