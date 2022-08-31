package dao

import (
	"net/http"
	"time"

	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
)

func MarketRates_Update(r dm.MarketRates, req *http.Request) error {
	var err error
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//hist.Date = Today()
	hist := marketRatesHistory_Build(r)

	err = MarketRates_Store(r, req)
	err = MarketRatesHistory_Store(hist, req)

	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//
	return err
}

func MarketRates_UpdateSystem(r dm.MarketRates) error {
	var err error
	// START
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	//
	//hist.Date = Today()
	hist := marketRatesHistory_Build(r)

	err = MarketRates_StoreSystem(r)
	err = MarketRatesHistory_StoreSystem(hist)

	//
	// Dynamically generated 28/06/2022 by matttownsend (Matt Townsend) on silicon.local
	// END
	//
	return err
}

func marketRatesHistory_Build(r dm.MarketRates) dm.MarketRatesHistory {
	hist := dm.MarketRatesHistory{}
	hist.Id = MarketRatesHistory_NewID(hist)
	hist.Bid = r.Bid
	hist.Mid = r.Mid
	hist.Offer = r.Offer
	hist.Market = r.Market
	hist.Tenor = r.Tenor
	hist.Series = r.Series
	hist.Name = r.Name
	hist.Source = r.Source
	hist.Destination = r.Destination
	hist.Class = r.Class

	hist.Date = time.Now().Format(core.DATEFORMATSIENA)
	hist.Time = time.Now().Format(core.TIMEHMS)
	hist.Datetime = time.Now().Format(core.DATETIME)
	return hist
}
