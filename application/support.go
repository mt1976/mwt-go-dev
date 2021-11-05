package application

import core "github.com/mt1976/mwt-go-dev/core"

//		PageTitle: core.ApplicationProperties["appname"] + core.Character_Break + Translation_Lookup("Page", dm.Translation_Title) + core.Character_Break + Translation_Lookup("Action", "New"),
func PageTitle(
	pageTitle string,
	pageSubTitle string) string {

	pt := Translation_Lookup("Page", pageTitle)
	pst := Translation_Lookup("Action", pageSubTitle)
	appName := core.ApplicationProperties["appname"]
	if len(appName) == 0 {
		appName = "Application Name"
	}
	pageTitle = appName + core.Character_Break + pt
	if len(pst) > 0 {
		pageTitle = appName + core.Character_Break + pt + core.Character_Break + pst
	}

	return pageTitle
}
