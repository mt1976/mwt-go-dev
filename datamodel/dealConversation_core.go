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
// Date & Time		    : 12/12/2021 at 16:13:13
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type DealConversation struct {

SienaReference        string
Status        string
MessageType        string
ContractNumber        string
AckReference        string
NewTX        string
LegNo        string
Summary        string
BusinessDate        string
TXNo        string
ExternalSystem        string
MessageLogReference        string

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
	DealConversation_TemplateList = "DealConversation_List"
	DealConversation_TemplateView = "DealConversation_View"
	DealConversation_TemplateEdit = "DealConversation_Edit"
	DealConversation_TemplateNew  = "DealConversation_New"
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
	/// SQL Field Definitions
	///
	DealConversation_SienaReference   = "SienaReference" // SienaReference is a String
	DealConversation_Status   = "Status" // Status is a String
	DealConversation_MessageType   = "MessageType" // MessageType is a String
	DealConversation_ContractNumber   = "ContractNumber" // ContractNumber is a String
	DealConversation_AckReference   = "AckReference" // AckReference is a String
	DealConversation_NewTX   = "NewTX" // NewTX is a Bool
	DealConversation_LegNo   = "LegNo" // LegNo is a Int
	DealConversation_Summary   = "Summary" // Summary is a String
	DealConversation_BusinessDate   = "BusinessDate" // BusinessDate is a Time
	DealConversation_TXNo   = "TXNo" // TXNo is a Int
	DealConversation_ExternalSystem   = "ExternalSystem" // ExternalSystem is a String
	DealConversation_MessageLogReference   = "MessageLogReference" // MessageLogReference is a String

	/// Definitions End
)
