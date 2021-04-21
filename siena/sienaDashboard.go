package siena

import (
	"fmt"
	"html/template"
	"log"
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

	wctProperties := application.GetProperties(globals.APPCONFIG)
	tmpl := "dashboard"

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)
	db, _ := Connect()
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
		UserMenu:          application.GetAppMenuData(globals.UserRole),
		UserRole:          globals.UserRole,
		UserNavi:          globals.UserNavi,
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

	t, _ := template.ParseFiles(application.GetTemplateID(tmpl, globals.UserRole))
	t.Execute(w, p)

}
