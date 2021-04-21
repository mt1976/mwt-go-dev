package siena

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	support "github.com/mt1976/mwt-go-dev/appsupport"
)

//sienaDashboardPage is cheese
type sienaDashboardPage struct {
	UserMenu          []AppMenuItem
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

func sienaDashboardHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	tmpl := "dashboard"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	db, _ := sienaConnect()
	noCps, _, _ := getSienaCounterpartyList(db)
	noDepd, dataDepd, _ := getSienaBIdealEventsPerDayList(db)
	noSecs, dataSect, _ := getSienaBIcounterpartyPerSectorList(db)

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
		UserMenu:          getappMenuData(gUserRole),
		UserRole:          gUserRole,
		UserNavi:          gUserNavi,
		Title:             wctProperties["appname"],
		PageTitle:         "List Siena Dashboards",
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

	fmt.Println(p)

	t, _ := template.ParseFiles(support.GetTemplateID(tmpl, gUserRole))
	t.Execute(w, p)

}
