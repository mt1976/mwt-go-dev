package siena

import (
	"html/template"
	"net/http"
	"strconv"

	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
)

//sienaDashboardPage is cheese
type sienaDashboardPage struct {
	UserMenu          []application.AppMenuItem
	UserRole          string
	UserNavi          string
	Title             string
	PageTitle         string
	ID                string
	Action            string
	TotCounterparties string
	NoGDPRExp         string
	NoLEIExp          string
	NoMiFIDRev        string
	DealsPerDay       []sienaBIdealEventsPerDayItem
	DEPDDataLabels    []string
	DEPDDataValues    []string
	SECTDataLabels    []string
	SECTDataValues    []string
}

func SienaDashboardHandler(w http.ResponseWriter, r *http.Request) {
	// Mandatory Security Validation
	if !(application.SessionValidate(w, r)) {
		application.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "dashboard"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	application.ServiceMessage(inUTL)

	noCps, _, _ := getSienaCounterpartyList()
	noDepd, dataDepd, _ := getSienaBIdealEventsPerDayList()
	noSecs, dataSect, _ := getSienaBIcounterpartyPerSectorList()

	var DLlist []string
	var DVlist []string
	for _, DVitem := range dataDepd {
		DLlist = append(DLlist, DVitem.StartInterestDate)
		DVlist = append(DVlist, DVitem.Count)
	}

	var SLlist []string
	var SVlist []string
	for _, SVitem := range dataSect {
		SLlist = append(SLlist, SVitem.SectorCode)
		SVlist = append(SVlist, SVitem.Count)
	}

	p := sienaDashboardPage{
		UserMenu:          application.GetUserMenu(r),
		UserRole:          application.GetUserRole(r),
		UserNavi:          "NOT USED",
		Title:             globals.ApplicationProperties["appname"],
		PageTitle: globals.ApplicationProperties["appname"] + " - “ +         "List Siena Dashboards",
		TotCounterparties: strconv.Itoa(noCps),
		NoGDPRExp:         strconv.Itoa(noDepd),
		NoLEIExp:          strconv.Itoa(noSecs),
		NoMiFIDRev:        "200",
		DealsPerDay:       dataDepd,
		DEPDDataValues:    DVlist,
		DEPDDataLabels:    DLlist,
		SECTDataValues:    SVlist,
		SECTDataLabels:    SLlist,
	}

	//fmt.Println(p)

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, application.GetUserRole(r)))
	t.Execute(w, p)

}
