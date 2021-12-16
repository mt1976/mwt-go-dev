package datamodel

// ----------------------------------------------------------------
// Automatically generated  "/datamodel/schedule.go"
// ----------------------------------------------------------------
// Package              : application
// Object 			    : Schedule (schedule)
// Endpoint 	        : Schedule (Schedule)
// For Project          : github.com/mt1976/mwt-go-dev/
// ----------------------------------------------------------------
// Template Generator   : delinquentDysprosium [r4-21.12.31]
// Date & Time		    : 14/12/2021 at 09:54:35
// Who & Where		    : matttownsend on silicon.local
// ----------------------------------------------------------------

//Centre is cheese
type Schedule struct {

SYSId        string
Id        string
Name        string
Description        string
Schedule        string
Started        string
Lastrun        string
Message        string
SYSCreated        string
SYSWho        string
SYSHost        string
SYSUpdated        string
Type        string
SYSCreatedBy        string
SYSCreatedHost        string
SYSUpdatedBy        string
SYSUpdatedHost        string
Human        string

}

const (
	Schedule_Title       = "Scheduler"
	Schedule_SQLTable    = "scheduleStore"
	Schedule_SQLSearchID = "id"
	Schedule_QueryString = "Schedule"
	///
	/// Handler Defintions
	///
	Schedule_Template     = "Schedule"
	Schedule_TemplateList = "Schedule_List"
	Schedule_TemplateView = "Schedule_View"
	Schedule_TemplateEdit = "Schedule_Edit"
	Schedule_TemplateNew  = "Schedule_New"
	///
	/// Handler Monitor Paths
	///
	Schedule_Path       = "/API/Schedule/"
	Schedule_PathList   = "/ScheduleList/"
	Schedule_PathView   = "/ScheduleView/"
	Schedule_PathEdit   = "/ScheduleEdit/"
	Schedule_PathNew    = "/ScheduleNew/"
	Schedule_PathSave   = "/ScheduleSave/"
	Schedule_PathDelete = "/ScheduleDelete/"
	///
	/// SQL Field Definitions
	///
	Schedule_SYSId   = "_id" // SYSId is a Int
	Schedule_Id   = "id" // Id is a String
	Schedule_Name   = "name" // Name is a String
	Schedule_Description   = "description" // Description is a String
	Schedule_Schedule   = "schedule" // Schedule is a String
	Schedule_Started   = "started" // Started is a String
	Schedule_Lastrun   = "lastrun" // Lastrun is a String
	Schedule_Message   = "message" // Message is a String
	Schedule_SYSCreated   = "_created" // SYSCreated is a String
	Schedule_SYSWho   = "_who" // SYSWho is a String
	Schedule_SYSHost   = "_host" // SYSHost is a String
	Schedule_SYSUpdated   = "_updated" // SYSUpdated is a String
	Schedule_Type   = "type" // Type is a String
	Schedule_SYSCreatedBy   = "_createdBy" // SYSCreatedBy is a String
	Schedule_SYSCreatedHost   = "_createdHost" // SYSCreatedHost is a String
	Schedule_SYSUpdatedBy   = "_updatedBy" // SYSUpdatedBy is a String
	Schedule_SYSUpdatedHost   = "_updatedHost" // SYSUpdatedHost is a String
	Schedule_Human   = "human" // Human is a String

	/// Definitions End
)
