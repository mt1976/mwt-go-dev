package jobs

import (
	"fmt"

	application "github.com/mt1976/mwt-go-dev/application"
	"github.com/openprovider/rates"
	"github.com/openprovider/rates/providers"
)

func RunJobECB(actionType string) {
	//logit(actionType, "*** START ***")
	service := rates.New(
		// any collection of providers which implement rates.Provider interface
		providers.NewECBProvider(new(rates.Options)),
	)
	rates, errors := service.FetchLast()
	if len(errors) != 0 {
		fmt.Println(errors)
	}
	fmt.Println(service.Name(), "exchange rates for today")
	for index, rate := range rates {
		fmt.Printf("%d. %s - %v\r\n", index+1, rate.Currency, rate.Value)

		var ratesData RatesDataStore
		ratesData.bid = fmt.Sprintf("%v\r\n", rate.Value)
		ratesData.mid = fmt.Sprintf("%v\r\n", rate.Value)
		ratesData.offer = fmt.Sprintf("%v\r\n", rate.Value)
		ratesData.market = "FX"
		ratesData.tenor = "SP"
		ratesData.series = fmt.Sprintf("%s", rate.Currency)
		ratesData.class = "ECB"
		ratesData.source = "ECB"
		ratesData.destination = "RVMARKET"
		RatesDataStorePut(ratesData)

	}
	application.UpdateSchedule(actionType, Aquirer, "")
	//logit(actionType, "*** DONE ***")
}
