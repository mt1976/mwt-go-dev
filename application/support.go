package application

import (
	"html/template"
	"net/http"

	core "github.com/mt1976/mwt-go-dev/core"
)

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

func ExecuteTemplate(tname string, w http.ResponseWriter, r *http.Request, data interface{}) {
	t := make(map[string]*template.Template)
	baseTemplateID := core.GetTemplateID(tname, core.GetUserRole(r))
	headerTemplateID := core.GetTemplateID("core/head", core.GetUserRole(r))
	t[tname] = template.Must(template.ParseFiles(baseTemplateID, headerTemplateID))
	t[tname].Execute(w, data)
}
