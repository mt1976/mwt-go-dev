package application

import (
	"net/http"

	core "github.com/mt1976/mwt-go-dev/core"
	dao "github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

//AccountTransaction_Publish annouces the endpoints available for this object
func AccountTransaction_Publish_Impl(mux http.ServeMux) {
	mux.HandleFunc(dm.AccountTransaction_PathViewTransaction, AccountTransaction_HandlerViewTransaction)

	logs.Publish("Siena", dm.AccountTransaction_Title+"Impl")

}

//AccountTransaction_HandlerList is the handler for the list page
func AccountTransaction_HandlerViewTransaction(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)
	searchID := core.GetURLparam(r, dm.AccountTransaction_QueryString)
	var returnList []dm.AccountTransaction
	noItems, returnList, _ := dao.AccountTransaction_GetListByID(searchID)

	pageDetail := AccountTransaction_PageList{
		Title:       core.ApplicationProperties["appname"],
		PageTitle:   PageTitle(dm.AccountTransaction_Title, core.Action_List),
		ItemsOnPage: noItems,
		ItemList:    returnList,
		UserMenu:    UserMenu_Get(r),
		UserRole:    Session_GetUserRole(r),
	}
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)

	ExecuteTemplate(dm.AccountTransaction_TemplateList, w, r, pageDetail)

}
