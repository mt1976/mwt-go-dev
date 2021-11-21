package jobs

import (
	"fmt"

	application "github.com/mt1976/mwt-go-dev/application"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	"github.com/mt1976/mwt-go-dev/logs"
	"github.com/openprovider/rates"
	"github.com/openprovider/rates/providers"
	cron "github.com/robfig/cron/v3"
)

func RatesFXECB_Job() dm.JobDefinition {
	var j dm.JobDefinition
	j.ID = "RATES_FXECB"
	j.Name = "RATES_FXECB"
	j.Period = "30 16 * * 1-5"
	j.Description = "Update ECB Benchmark FX Spot rate"
	j.Type = core.Aquirer
	return j
}

func RatesFXECB_Register(c *cron.Cron) {
	application.Schedule_Register(RatesFXECB_Job())
	c.AddFunc(RatesFXECB_Job().Period, func() { RatesFXECB_Run() })
}

// RunJobRollover is a Rollover function
func RatesFXECB_Run() {
	logs.StartJob(RatesFXECB_Job().Name)
	var message string
	/// CONTENT STARTS

	service := rates.New(
		// any collection of providers which implement rates.Provider interface
		providers.NewECBProvider(new(rates.Options)),
	)
	rates, err := service.FetchLast()
	if len(err) != 0 {
		fmt.Println(err)
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

	if err != nil {
		message = err[0].Error()
	}

	/// CONTENT ENDS
	application.Schedule_Update(RatesFXECB_Job(), message)
	logs.EndJob(RatesFXECB_Job().Name)
}
