package jobs

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	application "github.com/mt1976/mwt-go-dev/application"
	globals "github.com/mt1976/mwt-go-dev/globals"
	"golang.org/x/net/html"
)

const (
	LSEBaseURI                   = "https://www.londonstockexchange.com/download/csv/tbLSEBondsDownloadCSV_en.csv"
	LSEDateFormat                = "02-Jan-2006"
	FIIDateFormat                = "2 Jan 2006"
	FIICorporateBonds            = "https://www.fixedincomeinvestor.co.uk/x/bondtable.html?groupid=4"
	FIIIndexLinkedCorporateBonds = "https://www.fixedincomeinvestor.co.uk/x/bondtable.html?groupid=3562"
	FIIUKGilts                   = "https://www.fixedincomeinvestor.co.uk/x/bondtable.html?groupid=3"
	FIIUKIndexLinkedGilts        = "https://www.fixedincomeinvestor.co.uk/x/bondtable.html?groupid=3530"
	FFIURLPrefix                 = "https://www.fixedincomeinvestor.co.uk/x/"
)

type LSEBond struct {
	LongName                       string `csv:"LONG NAME"` // .csn Headers
	Isin                           string `csv:"ISIN"`
	Tidm                           string `csv:"TIDM"`
	Sedol                          string `csv:"SEDOL"`
	IssueDate                      string `csv:"ISSUE DATE"`
	MaturityDate                   string `csv:"MATURITY DATE"`
	CouponValue                    string `csv:"COUPON VALUE"`
	CouponType                     string `csv:"COUPON TYPE"`
	Segment                        string `csv:"SEGMENT"`
	Sector                         string `csv:"SECTOR"`
	CodeConventionCalculateAccrual string `csv:"CODE CONVENTION CALCULATE ACCRUAL"`
	MinimumDenomination            string `csv:"MINIMUM DENOMINATION"`
	DenominationCurrency           string `csv:"DENOMINATION CURRENCY"`
	TradingCurrency                string `csv:"TRADING CURRENCY"`
	Type                           string `csv:"TYPE"`
	FlatYield                      string `csv:"FLAT YIELD"`
	PaymentCouponDate              string `csv:"PAYMENT COUPON DATE"`
	PeriodOfCoupon                 string `csv:"PERIOD OF COUPON"`
	ExCouponDate                   string `csv:"EX COUPON DATE"`
	DateOfIndexInflation           string `csv:"DATE OF INDEX INFLATION"`
	UnitOfQuotation                string `csv:"UNIT OF QUOTATION"`
}

func RunJobLSE(actionType string) {

	//	log.Println("SONIA=", sonia)
	var message string
	dataFile, err := readCSVFromUrl(LSEBaseURI)
	if err != nil {
		log.Println(err.Error())
		message = err.Error()
	}

	if dataFile == nil {
		message = "No Data Aquired"
	}

	application.UpdateSchedule("LSE", globals.Aquirer, message)
	//logit(actionType, "*** DONE ***")
}

func readCSVFromUrl(url string) ([]LSEBond, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	//defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var bonds []LSEBond
	//log.Println(resp.Body)
	//log.Println(data)
	for row, theseStrings := range data {
		//	log.Println(row)
		if row > 2 {
			//log.Println(theseStrings)
			for _, thisRow := range theseStrings {
				var bondRow LSEBond
				thisRow := strings.Split(thisRow, ",")
				bondRow.LongName = application.GetTranslation("NI-Name", thisRow[0])
				bondRow.Isin = thisRow[1]
				bondRow.Tidm = thisRow[2]
				bondRow.Sedol = thisRow[3]
				bondRow.IssueDate = getInternalDate(thisRow[4])
				bondRow.MaturityDate = getInternalDate(thisRow[5])
				bondRow.CouponValue = thisRow[6]
				bondRow.CouponType = thisRow[7]
				bondRow.Segment = application.GetTranslation("NI-Segment", thisRow[8])
				bondRow.Sector = application.GetTranslation("NI-Sector", thisRow[9])
				bondRow.CodeConventionCalculateAccrual = application.GetTranslation("NI-AccrualType", thisRow[10])
				bondRow.MinimumDenomination = thisRow[11]
				bondRow.DenominationCurrency = thisRow[12]
				bondRow.TradingCurrency = thisRow[13]
				bondRow.Type = application.GetTranslation("NI-Type", thisRow[14])
				bondRow.FlatYield = thisRow[15]
				bondRow.PaymentCouponDate = getInternalDate(thisRow[16])
				bondRow.PeriodOfCoupon = application.GetTranslation("NI-CouponPeriod", thisRow[17])
				bondRow.ExCouponDate = getInternalDate(thisRow[18])
				bondRow.DateOfIndexInflation = thisRow[19]
				bondRow.UnitOfQuotation = thisRow[20]

				_, bondRec, _ := application.GetLSEGiltsDataStoreByID(bondRow.Isin)
				//		var bondRec application.AppLSEGiltsDataStoreItem
				if bondRec.Isin == "" {
					// ITs a new record
					bondRec.LongName = bondRow.LongName
					bondRec.Isin = bondRow.Isin
					bondRec.Tidm = bondRow.Tidm
					bondRec.Sedol = bondRow.Sedol
					bondRec.IssueDate = bondRow.IssueDate
					bondRec.MaturityDate = bondRow.MaturityDate
					bondRec.CouponValue = bondRow.CouponValue
					bondRec.CouponType = bondRow.CouponType
					bondRec.Segment = bondRow.Segment
					bondRec.Sector = bondRow.Sector
					bondRec.CodeConventionCalculateAccrual = bondRow.CodeConventionCalculateAccrual
					bondRec.MinimumDenomination = bondRow.MinimumDenomination
					bondRec.DenominationCurrency = bondRow.DenominationCurrency
					bondRec.TradingCurrency = bondRow.TradingCurrency
					bondRec.Type = bondRow.Type
					bondRec.FlatYield = bondRow.FlatYield
					bondRec.PaymentCouponDate = bondRow.PaymentCouponDate
					bondRec.PeriodOfCoupon = bondRow.PeriodOfCoupon
					bondRec.ExCouponDate = bondRow.ExCouponDate
					bondRec.DateOfIndexInflation = bondRow.DateOfIndexInflation
					bondRec.UnitOfQuotation = bondRow.UnitOfQuotation
					if bondRec.Segment == "UKGT" {
						bondRec.Issuer = application.GetTranslation("NI-Issuer", "UK Government")
					} else {
						bondRec.Issuer = application.GetTranslation("NI-Issuer", bondRec.LongName)
					}
					if len(bondRec.Type) == 0 {
						bondRec.Type = application.GetTranslation("NI-Type", bondRec.Segment)
					}
					if len(bondRec.Type) == 0 {
						bondRec.Type = application.GetTranslation("NI-Type", "Corporate Bond")
					}
				}

				bondRec.Tidm = bondRow.Tidm
				bondRec.Sedol = bondRow.Sedol
				bondRec.IssueDate = bondRow.IssueDate
				bondRec.CouponValue = bondRow.CouponValue
				bondRec.CouponType = bondRow.CouponType
				bondRec.CodeConventionCalculateAccrual = bondRow.CodeConventionCalculateAccrual
				bondRec.MinimumDenomination = bondRow.MinimumDenomination
				bondRec.DenominationCurrency = bondRow.DenominationCurrency
				bondRec.TradingCurrency = bondRow.TradingCurrency
				bondRec.FlatYield = bondRow.FlatYield
				bondRec.PaymentCouponDate = bondRow.PaymentCouponDate
				bondRec.PeriodOfCoupon = bondRow.PeriodOfCoupon
				bondRec.ExCouponDate = bondRow.ExCouponDate
				bondRec.DateOfIndexInflation = bondRow.DateOfIndexInflation
				bondRec.UnitOfQuotation = bondRow.UnitOfQuotation

				application.PutLSEGiltsDataStoreSystem(bondRec)

				//spew.Dump(bondRow)
				bonds = append(bonds, bondRow)
			}
		}
	}
	//	log.Print(bonds)
	return bonds, nil
}

func getInternalDate(in string) string {
	return getInternalDateGen(in, LSEDateFormat)
}

func getInternalDateGen(in string, format string) string {
	var intDate string
	if len(in) > 0 {
		//log.Println(LSEDateFormat, globals.DATEFORMATSIENA)
		time, err := time.Parse(format, in)
		if err != nil {
			log.Println(err.Error())
		}

		intDate = time.Format(globals.DATEFORMATSIENA)
		//log.Println(in, intDate)
	}
	return intDate
}

func getInternalDateFII(in string) string {
	return getInternalDateGen(in, FIIDateFormat)
}

func RunJobFII(actionType string) {
	//	log.Println("SONIA=", sonia)
	var message string
	OnPage(FIICorporateBonds)
	OnPage(FIIIndexLinkedCorporateBonds)
	OnPage(FIIUKGilts)
	OnPage(FIIUKIndexLinkedGilts)

	//dataFile := "poo"
	//log.Println(len(dataFile))
	//if len(dataFile) == 0 {
	//	message = "No Data Aquired"
	//}
	application.UpdateSchedule("FII", globals.Aquirer, message)
	//logit(actionType, "*** DONE ***")
}

func OnPage(link string) {

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "application/json")

	response, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}

	tokenizeFIIhtml(response)

	defer response.Body.Close()

	//	bytes, err := ioutil.ReadAll(response.Body)
	//	if err != nil {
	//		log.Fatalln(err)

	return
}

func tokenizeFIIhtml(response *http.Response) {

	inTable := false
	inTableBody := false
	inTableRow := false
	inTableCell := false
	noRows := 0
	noCols := 0
	sourceURI := ""
	var row []string
	var table [][]string
	tokenizer := html.NewTokenizer(response.Body)
	for {
		tt := tokenizer.Next()
		t := tokenizer.Token()

		err := tokenizer.Err()
		if err == io.EOF {
			break
		}
		//spew.Dump(tt)
		switch tt {
		case html.ErrorToken:
			log.Fatal(err)
		case html.TextToken:
			if inTable {
				//		spew.Dump(t)
				if inTableBody {
					if inTableRow {
						if inTableCell {
							//node := t.Type
							data := strings.TrimSpace(t.Data)
							//fmt.Println(inTable, inTableBody, inTableRow, node, data, noRows, noCols)
							row = append(row, data)
						}
					}
				}
			}
		case html.StartTagToken:
			if t.Data == "a" {

				ok, thisURI := getFFUURI(t)
				if ok {
					sourceURI = thisURI
				} else {
					sourceURI = ""
				} //spew.Dump(t)
				//log.Println(url)
			}
			//log.Println("!!!!!START TOKEN!!!!", t.Type, t.Data)
			if t.Data == "table" {
				inTable = true
			}
			if inTable {
				if t.Data == "tbody" {
					inTableBody = true
				}
			}
			if inTableBody {
				if t.Data == "tr" {
					inTableRow = true
				}
			}
			if inTableRow {
				if t.Data == "td" {
					inTableCell = true
					noCols = noCols + 1
				}
			}

		case html.EndTagToken:
			//	log.Println("!!!!!END TOKEN!!!!!", t.Type, t.Data)
			//spew.Dump(t)
			if t.Data == "table" {
				inTable = false
			}
			if inTable {
				if t.Data == "tbody" {
					inTableBody = false
				}
			}
			if inTableBody {
				if t.Data == "tr" {
					inTableRow = false
					noRows = noRows + 1
					processRow(row, noCols, sourceURI)

					//log.Print(row)

					table = append(table, row)
					noCols = 0
					row = nil

				}
			}
			if inTableRow {
				if t.Data == "td" {
					inTableCell = false
				}
			}
		}
	}
	// Print to check the slice's content
	//fmt.Println("table", table, noRows)
	//spew.Dump(table)
	if len(sourceURI) == 0 {
	}
}

func getFFUURI(t html.Token) (bool, string) {
	//	spew.Dump(t)
	ok := false
	href := ""
	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range t.Attr {
		if a.Key == "href" {
			//	fmt.Println("Got href")
			href = FFIURLPrefix + a.Val
			ok = true
		}
	}

	// "bare" return will return the variables (ok, href) as defined in
	// the function definition

	return ok, href
}

func processRow(row []string, noCols int, inURI string) {
	processDefinition(row, noCols, inURI)
	processPrice(row, noCols)
	//	log.Println("PROCESSED", row, noCols, inURI)
}

func processDefinition(row []string, noCols int, inURI string) {
	//	var bondRec application.AppLSEGiltsDataStoreItem
	_, bondRec, _ := application.GetLSEGiltsDataStoreByID(row[3])
	bondRec.LongName = application.GetTranslation("NI-Name", row[2])
	bondRec.Isin = row[3]
	//	bondRec.IssueDate = bondRow.IssueDate
	bondRec.MaturityDate = getInternalDateFII(row[5])
	bondRec.CouponValue = strings.ReplaceAll(row[4], "%", "")

	bondRec.DenominationCurrency = row[0]
	bondRec.TradingCurrency = row[0]

	if len(inURI) != 0 {
		bondRec = getFIIEnrichment(inURI, bondRec)
	}

	if bondRec.Segment == "UKGT" {
		bondRec.Issuer = application.GetTranslation("NI-Issuer", "UK Government")
	} else {
		bondRec.Issuer = application.GetTranslation("NI-Issuer", bondRec.LongName)
	}
	if len(bondRec.Type) == 0 {
		bondRec.Type = application.GetTranslation("NI-Type", bondRec.Segment)
	}
	if len(bondRec.Type) == 0 {
		bondRec.Type = application.GetTranslation("NI-Type", "Corporate Bond")
	}
	//// TODO: Use Longname/ISIN to get ISIN info and/or lookup a siena newSienaCounterparty
	// TODO: if countarparty is not found create Firm, add to center, create counterparty (as issuer), create counterpartyimportID
	// TODO: isinlookup  (might need to go into the Dispatcher Job)

	application.PutLSEGiltsDataStoreSystem(bondRec)
}

func getFIIEnrichment(inURI string, bondRec application.AppLSEGiltsDataStoreItem) application.AppLSEGiltsDataStoreItem {

	req, err := http.NewRequest("GET", inURI, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "application/json")

	response, err := http.Get(inURI)
	if err != nil {
		log.Fatal(err)
	}

	sector, couponFreq, yield, runningYield, issueAmount := tokenizeFIIenrichment(response)
	bondRec.Sector = application.GetTranslation("NI-Sector", sector)
	bondRec.PeriodOfCoupon = application.GetTranslation("NI-CouponPeriod", couponFreq)
	bondRec.FlatYield = strings.ReplaceAll(yield, "%", "")
	bondRec.RunningYield = strings.ReplaceAll(runningYield, "%", "")
	bondRec.IssueAmount = issueAmount

	return bondRec
}

func tokenizeFIIenrichment(response *http.Response) (string, string, string, string, string) {

	inTable := false
	inTableRow := false
	inTableCell := false
	noRows := 0
	noCols := 0
	sourceURI := ""
	sector := ""
	couponFreq := ""
	yield := ""
	runningYield := ""
	issueAmount := ""
	//prevData := []string{"", ""}

	var row []string
	var table [][]string
	tokenizer := html.NewTokenizer(response.Body)
	for {
		tt := tokenizer.Next()
		t := tokenizer.Token()

		err := tokenizer.Err()
		if err == io.EOF {
			break
		}
		//spew.Dump(tt)
		switch tt {
		case html.ErrorToken:
			log.Fatal(err)
		case html.TextToken:
			if inTable {
				//		spew.Dump(t)

				if inTableRow {
					//	prevData = nil
					if inTableCell {
						//node := t.Type
						data := strings.TrimSpace(t.Data)
						//fmt.Println(inTable, inTableBody, inTableRow, node, data, noRows, noCols)
						row = append(row, data)
					}
				}

			}
		case html.StartTagToken:
			if t.Data == "a" {

				ok, thisURI := getFFUURI(t)
				if ok {
					sourceURI = thisURI
				} else {
					sourceURI = ""
				} //spew.Dump(t)
				//log.Println(url)
			}
			//log.Println("!!!!!START TOKEN!!!!", t.Type, t.Data)
			if t.Data == "table" {
				inTable = true
			}

			if t.Data == "tr" {

				inTableRow = true
			}

			if inTableRow {
				if t.Data == "td" {
					inTableCell = true
					noCols = noCols + 1
				}
			}

		case html.EndTagToken:
			//	log.Println("!!!!!END TOKEN!!!!!", t.Type, t.Data)
			//spew.Dump(t)
			if t.Data == "table" {
				inTable = false
			}

			if t.Data == "tr" {
				inTableRow = false
				noRows = noRows + 1
				//	processRow(row, noCols, sourceURI)
				//spew.Dump(row)
				//log.Print("row=", row)
				if row[0] == "Industy:" {
					sector = row[1]
				}
				if row[0] == "Coupon Freq:" {
					couponFreq = row[1]
				}
				if row[0] == "Yield:" {
					yield = row[1]
				}
				if row[0] == "Running Yield:" {
					runningYield = row[1]
				}
				if row[0] == "Amount In Issue:" {
					issueAmount = row[1]
				}

				table = append(table, row)
				noCols = 0
				row = nil

			}
			if inTableRow {
				if t.Data == "td" {
					inTableCell = false
				}
			}
		}
	}
	// Print to check the slice's content
	//fmt.Println("table", table, noRows)
	//spew.Dump(table)
	if len(sourceURI) == 0 {
	}
	//log.Println("RESULT", sector, couponFreq)
	return sector, couponFreq, yield, runningYield, issueAmount
}

func processPrice(row []string, noCols int) {
	var ratesData RatesDataStore
	//	ratesData.bid = fmt.Sprintf("%v\r\n", bondRow.CouponValue)
	ratesData.mid = row[7]
	//	ratesData.offer = fmt.Sprintf("%v\r\n", bondRow.CouponValue)
	ratesData.market = "NI"
	ratesData.tenor = ""
	ratesData.series = row[3]
	ratesData.name = application.GetTranslation("NI-Name", row[2])
	ratesData.class = "Market"
	ratesData.source = "FII"
	ratesData.destination = "RVNI"
	RatesDataStorePut(ratesData)
}
