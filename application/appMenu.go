package application

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	globals "github.com/mt1976/mwt-go-dev/globals"
)

//AppMenuItem is cheese
type AppMenuItem struct {
	MenuID        string `json:"ID"`
	MenuText      string `json:"Text"`
	MenuGlyph     string `json:"Glyph"`
	MenuHREF      string `json:"HREF"`
	MenuOnClick   string `json:"OnClick"`
	MenuTextClass string `json:"TextClass"`
	MenuBreak     string `json:"Break"`
}

//AppMenuList is cheese
type AppMenuList struct {
	MenuItem []AppMenuItem `json:"menu_Nodes"`
}

//getappMenuData
func GetUserMenu(r *http.Request) []AppMenuItem {
	session_role := globals.SessionManager.GetString(r.Context(), globals.SessionRole)
	_, thisMenuList := FetchappMenuData(session_role)
	return thisMenuList.MenuItem
}

// fetchappMenuData read all employees
func FetchappMenuData(inRole string) (int, AppMenuList) {
	file, _ := ioutil.ReadFile(GetMenuID("menu", inRole))
	data := AppMenuList{}
	_ = json.Unmarshal([]byte(file), &data)
	return len(data.MenuItem), data
}
