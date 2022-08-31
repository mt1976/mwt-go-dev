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
// Date & Time		    : 28/06/2022 at 16:10:53
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//NegotiableInstrument defines the datamolde for the NegotiableInstrument object
type NegotiableInstrument struct {


SYSId       string
SYSId_props FieldProperties
Id       string
Id_props FieldProperties
LongName       string
LongName_props FieldProperties
Isin       string
Isin_props FieldProperties
Tidm       string
Tidm_props FieldProperties
Sedol       string
Sedol_props FieldProperties
IssueDate       string
IssueDate_props FieldProperties
MaturityDate       string
MaturityDate_props FieldProperties
CouponValue       string
CouponValue_props FieldProperties
CouponType       string
CouponType_props FieldProperties
Segment       string
Segment_props FieldProperties
Sector       string
Sector_props FieldProperties
CodeConventionCalculateAccrual       string
CodeConventionCalculateAccrual_props FieldProperties
MinimumDenomination       string
MinimumDenomination_props FieldProperties
DenominationCurrency       string
DenominationCurrency_props FieldProperties
TradingCurrency       string
TradingCurrency_props FieldProperties
Type       string
Type_props FieldProperties
FlatYield       string
FlatYield_props FieldProperties
PaymentCouponDate       string
PaymentCouponDate_props FieldProperties
PeriodOfCoupon       string
PeriodOfCoupon_props FieldProperties
ExCouponDate       string
ExCouponDate_props FieldProperties
DateOfIndexInflation       string
DateOfIndexInflation_props FieldProperties
UnitOfQuotation       string
UnitOfQuotation_props FieldProperties
SYSCreated       string
SYSCreated_props FieldProperties
SYSWho       string
SYSWho_props FieldProperties
SYSHost       string
SYSHost_props FieldProperties
SYSUpdated       string
SYSUpdated_props FieldProperties
Issuer       string
Issuer_props FieldProperties
IssueAmount       string
IssueAmount_props FieldProperties
RunningYield       string
RunningYield_props FieldProperties
LEI       string
LEI_props FieldProperties
CUSIP       string
CUSIP_props FieldProperties
SYSUpdatedHost       string
SYSUpdatedHost_props FieldProperties
SYSCreatedBy       string
SYSCreatedBy_props FieldProperties
SYSCreatedHost       string
SYSCreatedHost_props FieldProperties
SYSUpdatedBy       string
SYSUpdatedBy_props FieldProperties
 // Any lookups will be added below




































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
	NegotiableInstrument_TemplateList = "/NegotiableInstrument/NegotiableInstrument_List"
	NegotiableInstrument_TemplateView = "/NegotiableInstrument/NegotiableInstrument_View"
	NegotiableInstrument_TemplateEdit = "/NegotiableInstrument/NegotiableInstrument_Edit"
	NegotiableInstrument_TemplateNew  = "/NegotiableInstrument/NegotiableInstrument_New"
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
	//NegotiableInstrument_Redirect provides a page to return to aftern an action
	NegotiableInstrument_Redirect = NegotiableInstrument_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	NegotiableInstrument_SYSId_sql   = "_id" // SYSId is a Int
	NegotiableInstrument_Id_sql   = "id" // Id is a String
	NegotiableInstrument_LongName_sql   = "longName" // LongName is a String
	NegotiableInstrument_Isin_sql   = "isin" // Isin is a String
	NegotiableInstrument_Tidm_sql   = "tidm" // Tidm is a String
	NegotiableInstrument_Sedol_sql   = "sedol" // Sedol is a String
	NegotiableInstrument_IssueDate_sql   = "issueDate" // IssueDate is a String
	NegotiableInstrument_MaturityDate_sql   = "maturityDate" // MaturityDate is a String
	NegotiableInstrument_CouponValue_sql   = "couponValue" // CouponValue is a String
	NegotiableInstrument_CouponType_sql   = "couponType" // CouponType is a String
	NegotiableInstrument_Segment_sql   = "segment" // Segment is a String
	NegotiableInstrument_Sector_sql   = "sector" // Sector is a String
	NegotiableInstrument_CodeConventionCalculateAccrual_sql   = "codeConventionCalculateAccrual" // CodeConventionCalculateAccrual is a String
	NegotiableInstrument_MinimumDenomination_sql   = "minimumDenomination" // MinimumDenomination is a String
	NegotiableInstrument_DenominationCurrency_sql   = "denominationCurrency" // DenominationCurrency is a String
	NegotiableInstrument_TradingCurrency_sql   = "tradingCurrency" // TradingCurrency is a String
	NegotiableInstrument_Type_sql   = "type" // Type is a String
	NegotiableInstrument_FlatYield_sql   = "flatYield" // FlatYield is a String
	NegotiableInstrument_PaymentCouponDate_sql   = "paymentCouponDate" // PaymentCouponDate is a String
	NegotiableInstrument_PeriodOfCoupon_sql   = "periodOfCoupon" // PeriodOfCoupon is a String
	NegotiableInstrument_ExCouponDate_sql   = "exCouponDate" // ExCouponDate is a String
	NegotiableInstrument_DateOfIndexInflation_sql   = "dateOfIndexInflation" // DateOfIndexInflation is a String
	NegotiableInstrument_UnitOfQuotation_sql   = "unitOfQuotation" // UnitOfQuotation is a String
	NegotiableInstrument_SYSCreated_sql   = "_created" // SYSCreated is a String
	NegotiableInstrument_SYSWho_sql   = "_who" // SYSWho is a String
	NegotiableInstrument_SYSHost_sql   = "_host" // SYSHost is a String
	NegotiableInstrument_SYSUpdated_sql   = "_updated" // SYSUpdated is a String
	NegotiableInstrument_Issuer_sql   = "issuer" // Issuer is a String
	NegotiableInstrument_IssueAmount_sql   = "issueAmount" // IssueAmount is a String
	NegotiableInstrument_RunningYield_sql   = "runningYield" // RunningYield is a String
	NegotiableInstrument_LEI_sql   = "LEI" // LEI is a String
	NegotiableInstrument_CUSIP_sql   = "CUSIP" // CUSIP is a String
	NegotiableInstrument_SYSUpdatedHost_sql   = "_updatedHost" // SYSUpdatedHost is a String
	NegotiableInstrument_SYSCreatedBy_sql   = "_createdBy" // SYSCreatedBy is a String
	NegotiableInstrument_SYSCreatedHost_sql   = "_createdHost" // SYSCreatedHost is a String
	NegotiableInstrument_SYSUpdatedBy_sql   = "_updatedBy" // SYSUpdatedBy is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	NegotiableInstrument_SYSId_scrn   = "SYSId" // SYSId is a Int
	NegotiableInstrument_Id_scrn   = "Id" // Id is a String
	NegotiableInstrument_LongName_scrn   = "LongName" // LongName is a String
	NegotiableInstrument_Isin_scrn   = "Isin" // Isin is a String
	NegotiableInstrument_Tidm_scrn   = "Tidm" // Tidm is a String
	NegotiableInstrument_Sedol_scrn   = "Sedol" // Sedol is a String
	NegotiableInstrument_IssueDate_scrn   = "IssueDate" // IssueDate is a String
	NegotiableInstrument_MaturityDate_scrn   = "MaturityDate" // MaturityDate is a String
	NegotiableInstrument_CouponValue_scrn   = "CouponValue" // CouponValue is a String
	NegotiableInstrument_CouponType_scrn   = "CouponType" // CouponType is a String
	NegotiableInstrument_Segment_scrn   = "Segment" // Segment is a String
	NegotiableInstrument_Sector_scrn   = "Sector" // Sector is a String
	NegotiableInstrument_CodeConventionCalculateAccrual_scrn   = "CodeConventionCalculateAccrual" // CodeConventionCalculateAccrual is a String
	NegotiableInstrument_MinimumDenomination_scrn   = "MinimumDenomination" // MinimumDenomination is a String
	NegotiableInstrument_DenominationCurrency_scrn   = "DenominationCurrency" // DenominationCurrency is a String
	NegotiableInstrument_TradingCurrency_scrn   = "TradingCurrency" // TradingCurrency is a String
	NegotiableInstrument_Type_scrn   = "Type" // Type is a String
	NegotiableInstrument_FlatYield_scrn   = "FlatYield" // FlatYield is a String
	NegotiableInstrument_PaymentCouponDate_scrn   = "PaymentCouponDate" // PaymentCouponDate is a String
	NegotiableInstrument_PeriodOfCoupon_scrn   = "PeriodOfCoupon" // PeriodOfCoupon is a String
	NegotiableInstrument_ExCouponDate_scrn   = "ExCouponDate" // ExCouponDate is a String
	NegotiableInstrument_DateOfIndexInflation_scrn   = "DateOfIndexInflation" // DateOfIndexInflation is a String
	NegotiableInstrument_UnitOfQuotation_scrn   = "UnitOfQuotation" // UnitOfQuotation is a String
	NegotiableInstrument_SYSCreated_scrn   = "SYSCreated" // SYSCreated is a String
	NegotiableInstrument_SYSWho_scrn   = "SYSWho" // SYSWho is a String
	NegotiableInstrument_SYSHost_scrn   = "SYSHost" // SYSHost is a String
	NegotiableInstrument_SYSUpdated_scrn   = "SYSUpdated" // SYSUpdated is a String
	NegotiableInstrument_Issuer_scrn   = "Issuer" // Issuer is a String
	NegotiableInstrument_IssueAmount_scrn   = "IssueAmount" // IssueAmount is a String
	NegotiableInstrument_RunningYield_scrn   = "RunningYield" // RunningYield is a String
	NegotiableInstrument_LEI_scrn   = "LEI" // LEI is a String
	NegotiableInstrument_CUSIP_scrn   = "CUSIP" // CUSIP is a String
	NegotiableInstrument_SYSUpdatedHost_scrn   = "SYSUpdatedHost" // SYSUpdatedHost is a String
	NegotiableInstrument_SYSCreatedBy_scrn   = "SYSCreatedBy" // SYSCreatedBy is a String
	NegotiableInstrument_SYSCreatedHost_scrn   = "SYSCreatedHost" // SYSCreatedHost is a String
	NegotiableInstrument_SYSUpdatedBy_scrn   = "SYSUpdatedBy" // SYSUpdatedBy is a String

	/// Definitions End
	///
)

//negotiableinstrument_PageList provides the information for the template for a list of NegotiableInstruments
type NegotiableInstrument_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []NegotiableInstrument
	Context	 appContext
}

//negotiableinstrument_Page provides the information for the template for an individual NegotiableInstrument
type NegotiableInstrument_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SYSId         string
	SYSId_props     FieldProperties
	Id         string
	Id_props     FieldProperties
	LongName         string
	LongName_props     FieldProperties
	Isin         string
	Isin_props     FieldProperties
	Tidm         string
	Tidm_props     FieldProperties
	Sedol         string
	Sedol_props     FieldProperties
	IssueDate         string
	IssueDate_props     FieldProperties
	MaturityDate         string
	MaturityDate_props     FieldProperties
	CouponValue         string
	CouponValue_props     FieldProperties
	CouponType         string
	CouponType_props     FieldProperties
	Segment         string
	Segment_props     FieldProperties
	Sector         string
	Sector_props     FieldProperties
	CodeConventionCalculateAccrual         string
	CodeConventionCalculateAccrual_props     FieldProperties
	MinimumDenomination         string
	MinimumDenomination_props     FieldProperties
	DenominationCurrency         string
	DenominationCurrency_props     FieldProperties
	TradingCurrency         string
	TradingCurrency_props     FieldProperties
	Type         string
	Type_props     FieldProperties
	FlatYield         string
	FlatYield_props     FieldProperties
	PaymentCouponDate         string
	PaymentCouponDate_props     FieldProperties
	PeriodOfCoupon         string
	PeriodOfCoupon_props     FieldProperties
	ExCouponDate         string
	ExCouponDate_props     FieldProperties
	DateOfIndexInflation         string
	DateOfIndexInflation_props     FieldProperties
	UnitOfQuotation         string
	UnitOfQuotation_props     FieldProperties
	SYSCreated         string
	SYSCreated_props     FieldProperties
	SYSWho         string
	SYSWho_props     FieldProperties
	SYSHost         string
	SYSHost_props     FieldProperties
	SYSUpdated         string
	SYSUpdated_props     FieldProperties
	Issuer         string
	Issuer_props     FieldProperties
	IssueAmount         string
	IssueAmount_props     FieldProperties
	RunningYield         string
	RunningYield_props     FieldProperties
	LEI         string
	LEI_props     FieldProperties
	CUSIP         string
	CUSIP_props     FieldProperties
	SYSUpdatedHost         string
	SYSUpdatedHost_props     FieldProperties
	SYSCreatedBy         string
	SYSCreatedBy_props     FieldProperties
	SYSCreatedHost         string
	SYSCreatedHost_props     FieldProperties
	SYSUpdatedBy         string
	SYSUpdatedBy_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}