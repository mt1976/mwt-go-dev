package jobs

// ----------------------------------------------------------------
// Automatically generated  "/application/template.go"
// ----------------------------------------------------------------
// Package            : jobs
// Object 			  : Template
// Endpoint Root 	  : Template
// Search QueryString : TemplateID
// From   			  :
// ----------------------------------------------------------------
// Template Generator : RussetAlbatross [r0-21.11.01]
// ----------------------------------------------------------------
// Date & Time		  : 17/11/2021 at 22:28:20
// Who & Where		  : matttownsend on silicon.local
// ----------------------------------------------------------------

import (
	application "github.com/mt1976/mwt-go-dev/application"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
	cron "github.com/robfig/cron/v3"
)

func Template_Job() dm.JobDefinition {
	var j dm.JobDefinition
	j.ID = "Template"
	j.Name = "Template"
	j.Period = "10 1 * * *"
	j.Description = "Template processing"
	j.Type = core.HouseKeeping
	return j
}

func Template_Register(c *cron.Cron) {
	application.Schedule_Register(Template_Job())
	c.AddFunc(Template_Job().Period, func() { Template_Run() })
}

// RunJobRollover is a Rollover function
func Template_Run() {
	logs.StartJob(Template_Job().Name)
	message := ""
	/// CONTENT STARTS

	/// CONTENT ENDS
	application.Schedule_Update(Template_Job(), message)
	logs.EndJob(Template_Job().Name)
}
