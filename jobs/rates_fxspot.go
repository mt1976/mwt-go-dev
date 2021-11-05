package jobs

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/mt1976/common"
	application "github.com/mt1976/mwt-go-dev/application"
	core "github.com/mt1976/mwt-go-dev/core"
	dao "github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	tools "github.com/mt1976/mwtgostringtools"
	cron "github.com/robfig/cron/v3"
)

const constRateLen = 8

var funcName = ""

type fxRate struct {
	ccyPair   string
	bidRate   float64
	bidString string
	askRate   float64
	askString string
}

type fxRateCard struct {
	fxRates []fxRate
}

func RatesFXSpot_Job() dm.JobDefinition {
	var j dm.JobDefinition
	j.ID = "RATES_FXSP"
	j.Name = "RATES_FXSP"
	j.Period = "*/10 7-19 * * 1-5"
	j.Description = "Update FX Spot rate from barchart.com"
	j.Type = core.Aquirer
	return j
}

func RatesFXSpot_Register(c *cron.Cron) {
	application.RegisterSchedule(RatesFXSpot_Job())
	c.AddFunc(RatesFXSpot_Job().Period, func() { RatesFXSpot_Run() })
}

// RunJobRollover is a Rollover function
func RatesFXSpot_Run() {
	logStart(RatesFXSpot_Job().Name)
	var message string
	/// CONTENT STARTS

	_, cacheList, err := dao.DataCache_GetListByObject("CurrencyPair")
	if err != nil {
		log.Println(err.Error())
	}

	var rateCard fxRateCard

	for _, cacheData := range cacheList {
		//	fmt.Println(i, ccyPair)
		rateCard.fxRates = append(rateCard.fxRates, getFXrate(cacheData.Value))
	}

	if err != nil {
		message = err.Error()
	}

	/// CONTENT ENDS
	application.UpdateSchedule(RatesFXSpot_Job(), message)
	logEnd(RatesFXSpot_Job().Name)
}

func getFXrate(inCCYpair string) fxRate {
	thisRate := fxRate{}
	thisRate.ccyPair = inCCYpair
	thisPair := "%5E" + inCCYpair
	url := fmt.Sprintf("https://www.barchart.com/forex/quotes/%s/overview", thisPair)
	//fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		log.Println(err, inCCYpair)
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err, inCCYpair)
		panic(err)
	}
	// show the HTML code as a string %s
	//fmt.Printf("%s\n", html)
	inString := string(html)

	searchString := "\"bidPrice\":\""
	searchString2 := "\",\"askPrice\":\""
	searchString3 := "\",\"bidSize\":\""

	bidPriceStart := tools.FindStartPos(inString, searchString)
	bidPriceStop := tools.FindPos(inString, searchString2)
	askPriceStart := tools.FindStartPos(inString, searchString2)
	askPriceStop := tools.FindPos(inString, searchString3)
	if askPriceStop == -1 {
		askPriceStop = askPriceStart + 7
	}

	thisRate.bidRate, _ = strconv.ParseFloat(inString[bidPriceStart:bidPriceStop], 64)
	thisRate.bidString = tools.TruncateString(inString[bidPriceStart:bidPriceStop], constRateLen)
	//fmt.Println("bidPrice=", bidPrice)

	thisRate.askRate, _ = strconv.ParseFloat(inString[askPriceStart:askPriceStop], 64)
	thisRate.askString = tools.TruncateString(inString[askPriceStart:askPriceStop], constRateLen)
	//fmt.Println("askPrice=", askPrice)
	//log.Println(thisRate.ccyPair, bidPriceStart, askPriceStart, askPriceStop, thisRate.bidRate, thisRate.askRate)
	var ratesData RatesDataStore
	ratesData.bid = fmt.Sprintf("%f", thisRate.bidRate)
	ratesData.offer = fmt.Sprintf("%f", thisRate.askRate)
	ratesData.market = "FX"
	ratesData.tenor = "SP"
	ratesData.series = inCCYpair
	ratesData.class = "Market"
	ratesData.source = "Barchart.com"
	ratesData.destination = "RVMARKET"

	RatesDataStorePut(ratesData)
	return thisRate
}

var wg sync.WaitGroup

func getASYCFXrate(inCCYpair string, rateChan chan fxRate) {
	defer wg.Done()
	common.Snooze()
	thisRate := fxRate{}
	thisRate.ccyPair = inCCYpair
	thisPair := "%5E" + inCCYpair
	url := fmt.Sprintf("https://www.barchart.com/forex/quotes/%s/overview", thisPair)
	//fmt.Printf("HTML code of %s ...\n", url)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		log.Println(err, inCCYpair)
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err, inCCYpair)
		panic(err)
	}
	// show the HTML code as a string %s
	//fmt.Printf("%s\n", html)
	inString := string(html)

	searchString := "\"bidPrice\":\""
	searchString2 := "\",\"askPrice\":\""
	searchString3 := "\",\"bidSize\":\""

	bidPriceStart := tools.FindStartPos(inString, searchString)
	bidPriceStop := tools.FindPos(inString, searchString2)
	askPriceStart := tools.FindStartPos(inString, searchString2)
	askPriceStop := tools.FindPos(inString, searchString3)
	if askPriceStop == -1 {
		askPriceStop = askPriceStart + 7
	}

	thisRate.bidRate, _ = strconv.ParseFloat(inString[bidPriceStart:bidPriceStop], 64)
	thisRate.bidString = tools.TruncateString(inString[bidPriceStart:bidPriceStop], constRateLen)
	//fmt.Println("bidPrice=", bidPrice)

	thisRate.askRate, _ = strconv.ParseFloat(inString[askPriceStart:askPriceStop], 64)
	thisRate.askString = tools.TruncateString(inString[askPriceStart:askPriceStop], constRateLen)
	//fmt.Println("askPrice=", askPrice)

	rateChan <- thisRate

	//log.Println(thisRate.ccyPair, bidPriceStart, askPriceStart, askPriceStop, thisRate.bidRate, thisRate.askRate)

}

func buildRateCard() {
	var rateCard fxRateCard
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("AUDUSD"))
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("EURGBP"))
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("EURJPY"))
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("EURUSD"))
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("GBPUSD"))
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("NZDUSD"))
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("USDCAD"))
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("USDCHF"))
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("USDHKD"))
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("USDJPY"))
	rateCard.fxRates = append(rateCard.fxRates, getFXrate("USDSGD"))
}

func buldFXRVRates(rateCard fxRateCard) string {
	outputString := ""
	for _, s := range rateCard.fxRates {
		//fmt.Println(i, s)
		//fmt.Println(i, s)
		//fmt.Sprint(s)
		abc := fmt.Sprintf("s%ssptD%11sD%11s", s.ccyPair, s.bidString, s.askString)
		//fmt.Println(abc)
		outputString = outputString + abc + "\n"
	}
	//log.Println("\n", outputString)
	return outputString
}

func deliverRVData(name string, record string) {

	f, err := os.Create(name)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(record)

	if err2 != nil {
		log.Fatal(err2)
	}

}
