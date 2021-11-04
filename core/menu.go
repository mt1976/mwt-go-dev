package core

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

//AppMenuList is cheese
type AppMenuList struct {
	MenuItem []dm.AppMenuItem `json:"menu_Nodes"`
}

//getappMenuData
func GetUserMenu(r *http.Request) []dm.AppMenuItem {
	session_role := SessionManager.GetString(r.Context(), SessionRole)
	_, thisMenuList := FetchappMenuData(session_role)
	return thisMenuList.MenuItem
}

// fetchappMenuData read all employees
func FetchappMenuData(inRole string) (int, AppMenuList) {
	file, _ := ioutil.ReadFile(GetMenuID("menu", inRole))
	data := AppMenuList{}
	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.MenuItem); i++ {
		data.MenuItem[i].MenuHeaderText = ApplicationProperties["appname"]
	}

	return len(data.MenuItem), data
}
