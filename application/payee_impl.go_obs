package application

import (
	"net/http"

	core "github.com/mt1976/mwt-go-dev/core"
	dao "github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

func payee_HandlerViewImpl(pageDetail Payee_Page) Payee_Page { return pageDetail }
func payee_HandlerEditImpl(pageDetail Payee_Page) Payee_Page { return pageDetail }
func payee_HandlerSaveImpl(item dm.Payee) dm.Payee {
	return item
}

func Payee_PublishImpl(mux http.ServeMux) {

	// payee_PublishImpl should be specified in application/payee_Impl.go
	// to provide the implementation for the special case.
	// override should return mux - override function should be defined as
	// payee_PublishImpl(mux http.ServeMux) http.ServeMux {...}

	logs.Publish("Siena", dm.Payee_Title+" Special")
	mux.HandleFunc("/PayeeViewItem/", Payee_HandlerViewItem)

}

func Payee_HandlerViewItem(w http.ResponseWriter, r *http.Request) {
	//log.Println("poo")
	// Mandatory Security Validation
	if !(Session_Validate(w, r)) {
		core.Logout(w, r)
		return
	}

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

	idSource := core.GetURLparam(r, "csrc")
	idFirm := core.GetURLparam(r, "cfrm")
	idCentre := core.GetURLparam(r, "ccen")
	idCCY := core.GetURLparam(r, "cccy")
	idName := core.GetURLparam(r, "cnam")
	idNumber := core.GetURLparam(r, "cnum")
	idDirection := core.GetURLparam(r, "cdir")
	idType := core.GetURLparam(r, "ctyp")

	_, rD, _ := dao.Payee_GetByFullKey(idSource, idFirm, idCentre, idCCY, idName, idNumber, idDirection, idType)

	pageDetail := Payee_Page{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: PageTitle(dm.Payee_Title, core.Action_View),
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),

		SourceTable:           rD.SourceTable,
		KeyCounterpartyFirm:   rD.KeyCounterpartyFirm,
		KeyCounterpartyCentre: rD.KeyCounterpartyCentre,
		KeyCurrency:           rD.KeyCurrency,
		KeyName:               rD.KeyName,
		KeyNumber:             rD.KeyNumber,
		KeyDirection:          rD.KeyDirection,
		KeyType:               rD.KeyType,
		FullName:              rD.FullName,
		Address:               rD.Address,
		PhoneNo:               rD.PhoneNo,
		Country:               rD.Country,
		Bic:                   rD.Bic,
		Iban:                  rD.Iban,
		AccountNo:             rD.AccountNo,
		FedWireNo:             rD.FedWireNo,
		SortCode:              rD.SortCode,
		BankName:              rD.BankName,
		BankPinCode:           rD.BankPinCode,
		BankAddress:           rD.BankAddress,
		Reason:                rD.Reason,
		BankSettlementAcct:    rD.BankSettlementAcct,
		UpdatedUserId:         rD.UpdatedUserId,
	}

	_, Country_Lookup_Name, _ := dao.Country_GetByID(rD.Country)
	pageDetail.Country_Lookup = Country_Lookup_Name.Name
	_, KeyCounterpartyFirm_Lookup_FullName, _ := dao.Firm_GetByID(rD.KeyCounterpartyFirm)
	pageDetail.Firm_Lookup = KeyCounterpartyFirm_Lookup_FullName.FullName
	_, KeyCounterpartyCentre_Lookup_Name, _ := dao.Centre_GetByID(rD.KeyCounterpartyCentre)
	pageDetail.Centre_Lookup = KeyCounterpartyCentre_Lookup_Name.Name
	_, KeyCurrency_Lookup_Name, _ := dao.Currency_GetByID(rD.KeyCurrency)
	pageDetail.Currency_Lookup = KeyCurrency_Lookup_Name.Name

	pageDetail.SessionInfo, _ = Session_GetSessionInfo(r)
	ExecuteTemplate(dm.Payee_TemplateView, w, r, pageDetail)

}
