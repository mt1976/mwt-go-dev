package datamodel

//AppMenuItem holds the menu data for a specfic login instance
type AppMenuItem struct {
	MenuID         string `json:"ID"`
	MenuText       string `json:"Text"`
	MenuGlyph      string `json:"Glyph"`
	MenuHREF       string `json:"HREF"`
	MenuOnClick    string `json:"OnClick"`
	MenuTextClass  string `json:"TextClass"`
	MenuBreak      string `json:"Break"`
	MenuHeaderText string
}
