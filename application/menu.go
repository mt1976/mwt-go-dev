package application

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

//AppMenuList is cheese
type AppMenuList struct {
	MenuItem []dm.AppMenuItem `json:"menu_Nodes"`
}

//UserMenu_Get
func UserMenu_Get(r *http.Request) []dm.AppMenuItem {
	session_role := core.SessionManager.GetString(r.Context(), core.SessionRole)
	_, thisMenuList := userMenu_Fetch(session_role)
	return thisMenuList.MenuItem
}

// fetchappMenuData read all employees
func userMenu_Fetch(inRole string) (int, AppMenuList) {
	file, _ := ioutil.ReadFile(core.GetMenuID("menu", inRole))
	data := AppMenuList{}
	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.MenuItem); i++ {
		data.MenuItem[i].MenuText = Translation_Lookup("Menu", data.MenuItem[i].MenuText)
		data.MenuItem[i].MenuHeaderText = core.ApplicationProperties["appname"]
	}

	return len(data.MenuItem), data
}
