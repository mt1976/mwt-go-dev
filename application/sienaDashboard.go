package application

import (
	"net/http"
	"strconv"

	core "github.com/mt1976/mwt-go-dev/core"
	"github.com/mt1976/mwt-go-dev/dao"
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

	noCps, _, _ := dao.Counterparty_GetList()
	_, dataDepd, _ := getSienaBIdealEventsPerDayList()
	_, dataSect, _ := getSienaBIcounterpartyPerSectorList()

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
		UserMenu:          UserMenu_Get(r),
		UserRole:          core.GetUserRole(r),
		UserNavi:          "NOT USED",
		Title:             core.ApplicationProperties["appname"],
		PageTitle:         PageTitle("Dashboard", ""),
		TotCounterparties: strconv.Itoa(noCps),
		NoGDPRExp:         strconv.Itoa(50),
		NoLEIExp:          strconv.Itoa(10),
		NoMiFIDRev:        "200",
		DealsPerDay:       dataDepd,
		DEPDDataValues:    DVlist,
		DEPDDataLabels:    DLlist,
		SECTDataValues:    SVlist,
		SECTDataLabels:    SLlist,
	}

	//fmt.Println(p)

	ExecuteTemplate(tmpl, w, r, p)

}
