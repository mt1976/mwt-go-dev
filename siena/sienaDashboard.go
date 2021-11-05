package siena

import (
	"html/template"
	"net/http"
	"strconv"

	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

//sienaDashboardPage is cheese
type sienaDashboardPage struct {
	UserMenu          []dm.AppMenuItem
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
	if !(core.SessionValidate(w, r)) {
		core.LogoutHandler(w, r)
		return
	}
	// Code Continues Below

	tmpl := "dashboard"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	core.ServiceMessage(inUTL)

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
		UserMenu:          core.GetUserMenu(r),
		UserRole:          core.GetUserRole(r),
		UserNavi:          "NOT USED",
		Title:             core.ApplicationProperties["appname"],
		PageTitle:         core.ApplicationProperties["appname"] + " - " + "Dashboard",
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

	t, _ := template.ParseFiles(core.GetTemplateID(tmpl, core.GetUserRole(r)))
	t.Execute(w, p)

}
