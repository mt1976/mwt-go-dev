package jobs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	application "github.com/mt1976/mwt-go-dev/application"
	core "github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	"github.com/mt1976/mwt-go-dev/logs"
	cron "github.com/robfig/cron/v3"
)

const (
	apiKey    = "d4d50b886c05baa2795c363319ca64db"
	uri       = "https://api.stlouisfed.org/fred/series/observations?series_id=%s&api_key=%s&file_type=json"
	seriesuri = "https://api.stlouisfed.org/fred/series?series_id=%s&api_key=%s&file_type=json"
)

type FredSDW struct {
	RealtimeStart    string `json:"realtime_start"`
	RealtimeEnd      string `json:"realtime_end"`
	ObservationStart string `json:"observation_start"`
	ObservationEnd   string `json:"observation_end"`
	Units            string `json:"units"`
	OutputType       int    `json:"output_type"`
	FileType         string `json:"file_type"`
	OrderBy          string `json:"order_by"`
	SortOrder        string `json:"sort_order"`
	Count            int    `json:"count"`
	Offset           int    `json:"offset"`
	Limit            int    `json:"limit"`
	Observations     []struct {
		RealtimeStart string `json:"realtime_start"`
		RealtimeEnd   string `json:"realtime_end"`
		Date          string `json:"date"`
		Value         string `json:"value"`
	} `json:"observations"`
}

type FredSeriesInfo struct {
	RealtimeStart string `json:"realtime_start"`
	RealtimeEnd   string `json:"realtime_end"`
	Seriess       []struct {
		ID                      string `json:"id"`
		RealtimeStart           string `json:"realtime_start"`
		RealtimeEnd             string `json:"realtime_end"`
		Title                   string `json:"title"`
		ObservationStart        string `json:"observation_start"`
		ObservationEnd          string `json:"observation_end"`
		Frequency               string `json:"frequency"`
		FrequencyShort          string `json:"frequency_short"`
		Units                   string `json:"units"`
		UnitsShort              string `json:"units_short"`
		SeasonalAdjustment      string `json:"seasonal_adjustment"`
		SeasonalAdjustmentShort string `json:"seasonal_adjustment_short"`
		LastUpdated             string `json:"last_updated"`
		Popularity              int    `json:"popularity"`
		Notes                   string `json:"notes"`
	} `json:"seriess"`
}

func InstFRED_Job() dm.JobDefinition {
	var j dm.JobDefinition
	j.ID = "INST_FRED"
	j.Name = "INST_FRED"
	j.Period = "58 7-19 * * 1-5"
	j.Description = "Update Bond like Instruments from FRED"
	j.Type = core.Aquirer
	return j
}

func InstFRED_Register(c *cron.Cron) {
	application.Schedule_Register(InstFRED_Job())
	c.AddFunc(InstFRED_Job().Period, func() { InstFRED_Run() })
}

// RunJobRollover is a Rollover function
func InstFRED_Run() {
	logs.StartJob(InstFRED_Job().Name)
	var message string
	/// CONTENT STARTS

	seriesList := []string{"TB3MS", "TB1YR", "DTB6", "TB4WK", "DTB4WK", "BAMLC1A0C13YEY", "SOFR", "CBBTCUSD", "CPALTT01USM657N", "JPNCPIALLMINMEI", "GBRCPIALLMINMEI"}

	for _, s := range seriesList {
		requestURI := fmt.Sprintf(uri, s, apiKey)
		seriesRequestURI := fmt.Sprintf(seriesuri, s, apiKey)
		//logit(actionType, requestURI)
		value, valueDate := getFredAPIData(requestURI)
		//logit(actionType, value)
		//logit(actionType, valueDate)
		//logit(actionType, seriesRequestURI)
		//getFredSeriesData(seriesRequestURI)

		var ratesData RatesDataStore
		ratesData.mid = value
		ratesData.date = valueDate
		ratesData.market = "NI"
		ratesData.tenor = ""
		ratesData.series = s
		ratesData.class = "USD"
		ratesData.source = "Fred"
		ratesData.destination = "RVNI"
		ratesData.name = getFredSeriesData(seriesRequestURI)

		RatesDataStorePut(ratesData)

	}

	/// CONTENT ENDS
	application.Schedule_Update(InstFRED_Job(), message)
	logs.EndJob(InstFRED_Job().Name)
}

func getFredAPIData(inURI string) (string, string) {

	resp, err := http.Get(inURI)
	// handle the error if there is one
	if err != nil {
		log.Println(err)
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	//fmt.Println(string(html))

	var data FredSDW
	err2 := json.Unmarshal(html, &data)
	if err2 != nil {
		fmt.Println("error:", err2)
	}
	//	fmt.Printf("%+v", data)

	//fmt.Println("count", data.Count)
	//fmt.Println("value", data.Observations[data.Count-1].Value)
	//fmt.Println("date", data.Observations[data.Count-1].Date)
	return data.Observations[data.Count-1].Value, data.Observations[data.Count-1].Date
}

func getFredSeriesData(inURI string) string {

	resp, err := http.Get(inURI)
	// handle the error if there is one
	if err != nil {
		log.Println(err)
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	//fmt.Println(string(html))

	var data FredSeriesInfo
	err2 := json.Unmarshal(html, &data)
	if err2 != nil {
		fmt.Println("error:", err2)
	}
	//fmt.Printf("%+v", data)
	output := data.Seriess[0].Title
	////logit("getFredSeriesData", output)

	return output
	//fmt.Println("count", data.Count)
	//fmt.Println("value", data.Observations[data.Count-1].Value)
	//fmt.Println("date", data.Observations[data.Count-1].Date)
}
