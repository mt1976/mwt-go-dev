package main

import (
	"encoding/json"
	"io/ioutil"
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
func getappMenuData(inRole string) []AppMenuItem {
	_, thisMenuList := fetchappMenuData(inRole)
	return thisMenuList.MenuItem
}

// fetchappMenuData read all employees
func fetchappMenuData(inRole string) (int, AppMenuList) {

	file, _ := ioutil.ReadFile(getMenuID("menu"))

	data := AppMenuList{}

	_ = json.Unmarshal([]byte(file), &data)

	return len(data.MenuItem), data
}
