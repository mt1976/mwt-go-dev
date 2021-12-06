package application

import (
	"net/http"

	core "github.com/mt1976/mwt-go-dev/core"
	dao "github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

//AccountLadder_Publish annouces the endpoints available for this object
func AccountLadder_Publish_Impl(mux http.ServeMux) {
	mux.HandleFunc(dm.AccountLadder_PathViewLadder, AccountLadder_HandlerViewLadder)

	logs.Publish("Siena", dm.AccountLadder_Title+"impl")

}

//AccountLadder_HandlerList is the handler for the list page
func AccountLadder_HandlerViewLadder(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	searchID := core.GetURLparam(r, dm.AccountLadder_QueryString)

	var returnList []dm.AccountLadder
	noItems, returnList, _ := dao.AccountLadder_GetListByID(searchID)

	pageDetail := AccountLadder_PageList{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.AccountLadder_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:    returnList,
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	ExecuteTemplate(dm.AccountLadder_TemplateList, w, r, pageDetail)
}
