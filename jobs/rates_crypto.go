package jobs

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/mt1976/common"
	application "github.com/mt1976/mwt-go-dev/application"
	core "github.com/mt1976/mwt-go-dev/core"
	dao "github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
	tools "github.com/mt1976/mwtgostringtools"
	cron "github.com/robfig/cron/v3"
)

func RatesCrypto_Job() dm.JobDefinition {
	var j dm.JobDefinition
	j.ID = "RATES_CRYPTO"
	j.Name = "RATES_CRYPTO"
	j.Period = "*/10 7-19 * * 1-5"
	j.Description = "Update Crypto rate from barchart.com"
	j.Type = core.Aquirer
	return j
}

func RatesCrypto_Register(c *cron.Cron) {
	application.Schedule_Register(RatesCrypto_Job())
	c.AddFunc(RatesCrypto_Job().Period, func() { RatesCrypto_Run() })
}

// RunJobRollover is a Rollover function
func RatesCrypto_Run() {
	//logs.StartJob(RatesCrypto_Job().Name)

	logs.StartJob(RatesCrypto_Job().Name)
	var message string
	/// CONTENT STARTS

	testGet()

	noItems, cacheList, err := dao.Cache_GetListByObject("CryptoCurrency")
	if err != nil {
		log.Println(err.Error())
	}

	var rateCard fxRateCard

	for i := 0; i < noItems; i++ {
		//fmt.Printf("%s\n", cacheList[i])

		cacheData := cacheList[i]

		rateCard.fxRates = append(rateCard.fxRates, getCryptoRate(cacheData.Value))
	}

	if err != nil {
		message = err.Error()
	}

	/// CONTENT ENDS
	application.Schedule_Update(RatesCrypto_Job(), message)
	logs.EndJob(RatesCrypto_Job().Name)
}

func getCryptoRate(inCCYpair string) fxRate {
	thisRate := fxRate{}
	thisRate.ccyPair = inCCYpair
	thisPair := "%5E" + inCCYpair
	//https://www.barchart.com/crypto/quotes/%5EBTCEUR/overview
	//url := fmt.Sprintf("https://www.barchart.com/forex/quotes/%s/overview", thisPair)
	url := fmt.Sprintf("https://www.barchart.com/crypto/quotes/%s/overview", thisPair)
	logs.Accessing(url)
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
	//fmt.Println(inString)

	//"lastPrice":"1,638.565","priceChange":"-5.876","
	//,"lastPrice":"
	searchString := ",\"lastPrice\":\""

	//	logs.Accessing(searchString)

	searchString2 := "\",\"priceChange\":\""
	//	logs.Accessing(searchString2)

	//	fmt.Printf("%s\n", inString)
	//	searchString3 := "\",\"bidSize\":\""
	// <span class="last-change ng-binding" data-ng-class="highlightValue('lastPrice')">1,643.311</span>
	bidPriceStart := tools.FindStartPos(inString, searchString)

	bidPriceStop := tools.FindPos(inString, searchString2)

	//	fmt.Println("bidPriceStart: ", bidPriceStart)
	//	fmt.Println("bidPriceStop: ", bidPriceStop)

	//askPriceStart := tools.FindStartPos(inString, searchString2)
	//askPriceStop := tools.FindPos(inString, searchString3)
	//if askPriceStop == -1 {
	//	askPriceStop = askPriceStart + 7
	//}

	thisRate.bidString = tools.TruncateString(inString[bidPriceStart:bidPriceStop], 18)
	thisRate.bidString = strings.ReplaceAll(thisRate.bidString, ",", "")
	thisRate.bidRate, err = strconv.ParseFloat(thisRate.bidString, 64)

	//	fmt.Printf("thisRate: %v\n", thisRate)
	//panic(err)

	//	fmt.Println("bidPrice=", thisRate.bidString)

	//thisRate.askRate, _ = strconv.ParseFloat(inString[askPriceStart:askPriceStop], 64)
	//thisRate.askString = tools.TruncateString(inString[askPriceStart:askPriceStop], constRateLen)
	//fmt.Println("askPrice=", askPrice)
	//log.Println(thisRate.ccyPair, bidPriceStart, askPriceStart, askPriceStop, thisRate.bidRate, thisRate.askRate)
	var ratesData RatesDataStore
	ratesData.bid = fmt.Sprintf("%f", thisRate.bidRate)
	//ratesData.mid = fmt.Sprintf("%f", thisRate.bidRate)
	ratesData.offer = fmt.Sprintf("%f", thisRate.bidRate)
	ratesData.market = dm.Market_Crypto
	ratesData.tenor = dm.Tenor_SP
	ratesData.series = inCCYpair
	ratesData.class = dm.Market_Crypto
	ratesData.name = application.Translation_Lookup("CryptoPair", inCCYpair)
	ratesData.source = url
	ratesData.destination = "RV" + dm.RateCategory_Market

	RatesDataStorePut(ratesData)
	return thisRate
}

//var wg sync.WaitGroup

func getASYCCrypto(inCCYpair string, rateChan chan fxRate) {
	defer wg.Done()
	common.Snooze()
	thisRate := fxRate{}
	thisRate.ccyPair = inCCYpair
	thisPair := "%5E" + inCCYpair
	url := fmt.Sprintf("https://www.barchart.com/forex/quotes/%s/overview", thisPair)
	fmt.Printf("HTML code of %s ...\n", url)
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

	fmt.Printf("thisRate: %v\n", thisRate)

	rateChan <- thisRate

	//log.Println(thisRate.ccyPair, bidPriceStart, askPriceStart, askPriceStop, thisRate.bidRate, thisRate.askRate)

}

func testGet() {

}
