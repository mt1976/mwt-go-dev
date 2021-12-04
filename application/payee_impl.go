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

	_, returnRecord, _ := dao.Payee_GetByFullKey(idSource, idFirm, idCentre, idCCY, idName, idNumber, idDirection, idType)

	pageDetail := Payee_Page{
		Title:     core.ApplicationProperties["appname"],
		PageTitle: PageTitle(dm.Payee_Title, core.Action_View),
		UserMenu:  UserMenu_Get(r),
		UserRole:  Session_GetUserRole(r),

		SourceTable:           returnRecord.SourceTable,
		KeyCounterpartyFirm:   returnRecord.KeyCounterpartyFirm,
		KeyCounterpartyCentre: returnRecord.KeyCounterpartyCentre,
		KeyCurrency:           returnRecord.KeyCurrency,
		KeyName:               returnRecord.KeyName,
		KeyNumber:             returnRecord.KeyNumber,
		KeyDirection:          returnRecord.KeyDirection,
		KeyType:               returnRecord.KeyType,
		FullName:              returnRecord.FullName,
		Address:               returnRecord.Address,
		PhoneNo:               returnRecord.PhoneNo,
		Country:               returnRecord.Country,
		Bic:                   returnRecord.Bic,
		Iban:                  returnRecord.Iban,
		AccountNo:             returnRecord.AccountNo,
		FedWireNo:             returnRecord.FedWireNo,
		SortCode:              returnRecord.SortCode,
		BankName:              returnRecord.BankName,
		BankPinCode:           returnRecord.BankPinCode,
		BankAddress:           returnRecord.BankAddress,
		Reason:                returnRecord.Reason,
		BankSettlementAcct:    returnRecord.BankSettlementAcct,
		UpdatedUserId:         returnRecord.UpdatedUserId,
	}

	_, Country_Lookup, _ := dao.Country_GetByID(returnRecord.Country)
	pageDetail.Country_Impl = Country_Lookup.Name
	_, KeyCounterpartyFirm_Lookup, _ := dao.Firm_GetByID(returnRecord.KeyCounterpartyFirm)
	pageDetail.Firm_Impl = KeyCounterpartyFirm_Lookup.FullName
	_, KeyCounterpartyCentre_Lookup, _ := dao.Centre_GetByID(returnRecord.KeyCounterpartyCentre)
	pageDetail.Centre_Impl = KeyCounterpartyCentre_Lookup.Name
	_, KeyCurrency_Lookup, _ := dao.Currency_GetByID(returnRecord.KeyCurrency)
	pageDetail.Currency_Impl = KeyCurrency_Lookup.Name

	ExecuteTemplate(dm.Payee_TemplateView, w, r, pageDetail)

}
