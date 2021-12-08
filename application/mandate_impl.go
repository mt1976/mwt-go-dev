package application

import (
	"net/http"

	core "github.com/mt1976/mwt-go-dev/core"
	dao "github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func mandate_PublishImpl(mux http.ServeMux) http.ServeMux          { return mux }
func mandate_HandlerEditImpl(pageDetail Mandate_Page) Mandate_Page { return pageDetail }
func mandate_HandlerViewImpl(pageDetail Mandate_Page) Mandate_Page { return pageDetail }
func mandate_HandlerSaveImpl(item dm.Mandate) dm.Mandate           { return item }

//Mandate_Publish annouces the endpoints available for this object
func Mandate_PublishImpl(mux http.ServeMux) {

	mux.HandleFunc(dm.Mandate_PathViewItem, Mandate_HandlerViewItem)
	mux.HandleFunc(dm.Mandate_PathEditItem, Mandate_HandlerEditItem)
	logs.Publish("Siena", dm.Mandate_Title+"Impl")

	// mandate_PublishImpl should be specified in application/mandate_Impl.go
	// to provide the implementation for the special case.
	// override function should be defined as
	// mandate_PublishImpl(mux http.ServeMux) {...}
	// TODO - this is a temporary hack to get the special case working
	// Add to main.go >>> mandate_PublishImpl(mux)

}

//Mandate_HandlerView is the handler used to View a page
func Mandate_HandlerViewItem(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	muID := core.GetURLparam(r, "SU")
	muFirm := core.GetURLparam(r, "SF")
	muCentre := core.GetURLparam(r, "SC")

	//SU={{.MandatedUserKeyUserName}}&SF={{.MandatedUserKeyCounterpartyFirm}}&SC={{.MandatedUserKeyCounterpartyCentre}}

	_, rD, _ := dao.Mandate_GetByFullID(muID, muFirm, muCentre)

	pageDetail := Mandate_Page{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: PageTitle(dm.Mandate_Title, core.Action_View),
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
	}

	//
	// Automatically generated 02/12/2021 by matttownsend on silicon.local - START
	pageDetail.MandatedUserKeyCounterpartyFirm = rD.MandatedUserKeyCounterpartyFirm
	pageDetail.MandatedUserKeyCounterpartyCentre = rD.MandatedUserKeyCounterpartyCentre
	pageDetail.MandatedUserKeyUserName = rD.MandatedUserKeyUserName
	pageDetail.TelephoneNumber = rD.TelephoneNumber
	pageDetail.EmailAddress = rD.EmailAddress
	pageDetail.Active = rD.Active
	pageDetail.FirstName = rD.FirstName
	pageDetail.Surname = rD.Surname
	pageDetail.DateOfBirth = rD.DateOfBirth
	pageDetail.Postcode = rD.Postcode
	pageDetail.NationalIDNo = rD.NationalIDNo
	pageDetail.PassportNo = rD.PassportNo
	pageDetail.Country = rD.Country
	pageDetail.CountryName = rD.CountryName
	pageDetail.FirmName = rD.FirmName
	pageDetail.CentreName = rD.CentreName
	pageDetail.Notify = rD.Notify
	pageDetail.SystemUser = rD.SystemUser
	pageDetail.CompID = rD.CompID
	// Automatically generated 02/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
	_, Country_Lookup, _ := dao.Country_GetByID(rD.Country)
	pageDetail.Country_Lookup = Country_Lookup.Name
	_, MandatedUserKeyCounterpartyFirm_Lookup, _ := dao.Firm_GetByID(rD.MandatedUserKeyCounterpartyFirm)
	pageDetail.Firm_Lookup = MandatedUserKeyCounterpartyFirm_Lookup.FullName
	_, MandatedUserKeyCounterpartyCentre_Lookup, _ := dao.Centre_GetByID(rD.MandatedUserKeyCounterpartyCentre)
	pageDetail.Centre_Lookup = MandatedUserKeyCounterpartyCentre_Lookup.Name
	// Automatically generated 02/12/2021 by matttownsend on silicon.local - END
	//

	// mandate_HandlerViewImpl should be specified in application/mandate_Impl.go
	// to provide the implementation for the special case.
	// override should return mux - override function should be defined as
	// func mandate_HandlerViewImpl(pageDetail Mandate_Page) Mandate_Page {return pageDetail}
	pageDetail = mandate_HandlerViewImpl(pageDetail)

	// Automatically generated 02/12/2021 by matttownsend on silicon.local - END
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	ExecuteTemplate(dm.Mandate_TemplateView, w, r, pageDetail)
}

//Mandate_HandlerEdit is the handler used generate the Edit page
func Mandate_HandlerEditItem(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}
	// Code Continues Below

	w.Header().Set("Content-Type", "text/html")
	logs.Servicing(r.URL.Path)

	muID := core.GetURLparam(r, "SU")
	muFirm := core.GetURLparam(r, "SF")
	muCentre := core.GetURLparam(r, "SC")

	//SU={{.MandatedUserKeyUserName}}&SF={{.MandatedUserKeyCounterpartyFirm}}&SC={{.MandatedUserKeyCounterpartyCentre}}

	_, rD, _ := dao.Mandate_GetByFullID(muID, muFirm, muCentre)

	pageDetail := Mandate_Page{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: PageTitle(dm.Mandate_Title, core.Action_Edit),
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),
	}

	//
	// Automatically generated 02/12/2021 by matttownsend on silicon.local - START
	pageDetail.MandatedUserKeyCounterpartyFirm = rD.MandatedUserKeyCounterpartyFirm
	pageDetail.MandatedUserKeyCounterpartyCentre = rD.MandatedUserKeyCounterpartyCentre
	pageDetail.MandatedUserKeyUserName = rD.MandatedUserKeyUserName
	pageDetail.TelephoneNumber = rD.TelephoneNumber
	pageDetail.EmailAddress = rD.EmailAddress
	pageDetail.Active = rD.Active
	pageDetail.FirstName = rD.FirstName
	pageDetail.Surname = rD.Surname
	pageDetail.DateOfBirth = rD.DateOfBirth
	pageDetail.Postcode = rD.Postcode
	pageDetail.NationalIDNo = rD.NationalIDNo
	pageDetail.PassportNo = rD.PassportNo
	pageDetail.Country = rD.Country
	pageDetail.CountryName = rD.CountryName
	pageDetail.FirmName = rD.FirmName
	pageDetail.CentreName = rD.CentreName
	pageDetail.Notify = rD.Notify
	pageDetail.SystemUser = rD.SystemUser
	pageDetail.CompID = rD.CompID
	// Automatically generated 02/12/2021 by matttownsend on silicon.local - Enrichment Fields Below
	_, Country_Lookup, _ := dao.Country_GetByID(rD.Country)
	pageDetail.Country_Lookup = Country_Lookup.Name
	_, pageDetail.Country_Lookup_List, _ = dao.Country_GetList()
	_, MandatedUserKeyCounterpartyFirm_Lookup, _ := dao.Firm_GetByID(rD.MandatedUserKeyCounterpartyFirm)
	pageDetail.Firm_Lookup = MandatedUserKeyCounterpartyFirm_Lookup.FullName
	_, pageDetail.Firm_Lookup_List, _ = dao.Firm_GetList()
	_, MandatedUserKeyCounterpartyCentre_Lookup, _ := dao.Centre_GetByID(rD.MandatedUserKeyCounterpartyCentre)
	pageDetail.Centre_Lookup = MandatedUserKeyCounterpartyCentre_Lookup.Name
	_, pageDetail.Centre_Lookup_List, _ = dao.Centre_GetList()
	// Automatically generated 02/12/2021 by matttownsend on silicon.local - END
	//

	// mandate_HandlerEditImpl should be specified in application/mandate_Impl.go
	// to provide the implementation for the special case.
	// override should return mux - override function should be defined as
	// func mandate_HandlerEditImpl(pageDetail Mandate_Page) Mandate_Page {return pageDetail}
	pageDetail = mandate_HandlerEditImpl(pageDetail)

	// Automatically generated 02/12/2021 by matttownsend on silicon.local - END
	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	ExecuteTemplate(dm.Mandate_TemplateEdit, w, r, pageDetail)

}
