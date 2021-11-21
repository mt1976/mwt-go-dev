package application

import (
	"net/http"

	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

func mandate_PublishImpl(mux http.ServeMux) http.ServeMux          { return mux }
func mandate_HandlerEditImpl(pageDetail Mandate_Page) Mandate_Page { return pageDetail }
func mandate_HandlerViewImpl(pageDetail Mandate_Page) Mandate_Page { return pageDetail }
func mandate_HandlerSaveImpl(item dm.Mandate) dm.Mandate           { return item }
