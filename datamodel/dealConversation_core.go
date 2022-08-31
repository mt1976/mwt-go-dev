package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/dealconversation.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : DealConversation (dealconversation)
// Endpoint 	        : DealConversation (ID)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 28/06/2022 at 16:10:50
// Who & Where		    : matttownsend (Matt Townsend) on silicon.local
// ----------------------------------------------------------------

//DealConversation defines the datamolde for the DealConversation object
type DealConversation struct {


SienaReference       string
SienaReference_props FieldProperties
Status       string
Status_props FieldProperties
MessageType       string
MessageType_props FieldProperties
ContractNumber       string
ContractNumber_props FieldProperties
AckReference       string
AckReference_props FieldProperties
NewTX       string
NewTX_props FieldProperties
LegNo       string
LegNo_props FieldProperties
Summary       string
Summary_props FieldProperties
BusinessDate       string
BusinessDate_props FieldProperties
TXNo       string
TXNo_props FieldProperties
ExternalSystem       string
ExternalSystem_props FieldProperties
MessageLogReference       string
MessageLogReference_props FieldProperties
 // Any lookups will be added below












}

const (
	DealConversation_Title       = "Deal Conversation"
	DealConversation_SQLTable    = "sienaDealConversationLog"
	DealConversation_SQLSearchID = "MessageLogReference"
	DealConversation_QueryString = "ID"
	///
	/// Handler Defintions
	///
	DealConversation_Template     = "DealConversation"
	DealConversation_TemplateList = "/DealConversation/DealConversation_List"
	DealConversation_TemplateView = "/DealConversation/DealConversation_View"
	DealConversation_TemplateEdit = "/DealConversation/DealConversation_Edit"
	DealConversation_TemplateNew  = "/DealConversation/DealConversation_New"
	///
	/// Handler Monitor Paths
	///
	DealConversation_Path       = "/API/DealConversation/"
	DealConversation_PathList   = "/DealConversationList/"
	DealConversation_PathView   = "/DealConversationView/"
	DealConversation_PathEdit   = "/DealConversationEdit/"
	DealConversation_PathNew    = "/DealConversationNew/"
	DealConversation_PathSave   = "/DealConversationSave/"
	DealConversation_PathDelete = "/DealConversationDelete/"
	///
	//DealConversation_Redirect provides a page to return to aftern an action
	DealConversation_Redirect = DealConversation_PathList
	
	///
	///
	/// SQL Field Definitions
	///
	DealConversation_SienaReference_sql   = "SienaReference" // SienaReference is a String
	DealConversation_Status_sql   = "Status" // Status is a String
	DealConversation_MessageType_sql   = "MessageType" // MessageType is a String
	DealConversation_ContractNumber_sql   = "ContractNumber" // ContractNumber is a String
	DealConversation_AckReference_sql   = "AckReference" // AckReference is a String
	DealConversation_NewTX_sql   = "NewTX" // NewTX is a Bool
	DealConversation_LegNo_sql   = "LegNo" // LegNo is a Int
	DealConversation_Summary_sql   = "Summary" // Summary is a String
	DealConversation_BusinessDate_sql   = "BusinessDate" // BusinessDate is a Time
	DealConversation_TXNo_sql   = "TXNo" // TXNo is a Int
	DealConversation_ExternalSystem_sql   = "ExternalSystem" // ExternalSystem is a String
	DealConversation_MessageLogReference_sql   = "MessageLogReference" // MessageLogReference is a String

	/// Definitions End
	///
	/// Application Field Definitions
	///
	DealConversation_SienaReference_scrn   = "SienaReference" // SienaReference is a String
	DealConversation_Status_scrn   = "Status" // Status is a String
	DealConversation_MessageType_scrn   = "MessageType" // MessageType is a String
	DealConversation_ContractNumber_scrn   = "ContractNumber" // ContractNumber is a String
	DealConversation_AckReference_scrn   = "AckReference" // AckReference is a String
	DealConversation_NewTX_scrn   = "NewTX" // NewTX is a Bool
	DealConversation_LegNo_scrn   = "LegNo" // LegNo is a Int
	DealConversation_Summary_scrn   = "Summary" // Summary is a String
	DealConversation_BusinessDate_scrn   = "BusinessDate" // BusinessDate is a Time
	DealConversation_TXNo_scrn   = "TXNo" // TXNo is a Int
	DealConversation_ExternalSystem_scrn   = "ExternalSystem" // ExternalSystem is a String
	DealConversation_MessageLogReference_scrn   = "MessageLogReference" // MessageLogReference is a String

	/// Definitions End
	///
)

//dealconversation_PageList provides the information for the template for a list of DealConversations
type DealConversation_PageList struct {
	SessionInfo      SessionInfo
	UserMenu         AppMenuItem
	UserRole         string
	Title            string
	PageTitle        string
	ItemsOnPage 	 int
	ItemList  		 []DealConversation
	Context	 appContext
}

//dealconversation_Page provides the information for the template for an individual DealConversation
type DealConversation_Page struct {
	SessionInfo      SessionInfo
	UserMenu    	 AppMenuItem
	UserRole    	 string
	Title       	 string
	PageTitle   	 string
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	//	
	SienaReference         string
	SienaReference_props     FieldProperties
	Status         string
	Status_props     FieldProperties
	MessageType         string
	MessageType_props     FieldProperties
	ContractNumber         string
	ContractNumber_props     FieldProperties
	AckReference         string
	AckReference_props     FieldProperties
	NewTX         string
	NewTX_props     FieldProperties
	LegNo         string
	LegNo_props     FieldProperties
	Summary         string
	Summary_props     FieldProperties
	BusinessDate         string
	BusinessDate_props     FieldProperties
	TXNo         string
	TXNo_props     FieldProperties
	ExternalSystem         string
	ExternalSystem_props     FieldProperties
	MessageLogReference         string
	MessageLogReference_props     FieldProperties
	// 
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local 
	// END
	Context	 appContext
}