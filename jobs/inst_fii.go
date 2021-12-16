package jobs

import (
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	application "github.com/mt1976/mwt-go-dev/application"
	core "github.com/mt1976/mwt-go-dev/core"
	"github.com/mt1976/mwt-go-dev/dao"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	"github.com/mt1976/mwt-go-dev/logs"
	cron "github.com/robfig/cron/v3"
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
	MIBaseURI                    = "https://markets.businessinsider.com/bonds/"
	MIDate                       = "1/2/2006"
	FIIBenchmarks                = "https://www.fixedincomeinvestor.co.uk/x/bondtable.html?groupid=8"
	FIIBondIndex                 = "https://www.fixedincomeinvestor.co.uk/x/bondtable.html?groupid=7"
	FIIPIBS                      = "https://www.fixedincomeinvestor.co.uk/x/bondtable.html?groupid=11"
	FIIORB                       = "https://www.fixedincomeinvestor.co.uk/x/bondtable.html?groupid=3620"
	FIIORBBONDS                  = "https://www.fixedincomeinvestor.co.uk/x/bondtable.html?groupid=3653"
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

type MIBondData struct {
	ISIN              string
	Name              string
	Country           string
	Issuence          string
	Issuer            string
	IssueVolume       string
	Currency          string
	IssuePrice        string
	IssueDate         string
	Coupon            string
	Coupon2           string
	Denomination      string
	QuotationType     string
	PaymentType       string
	SpecialCouponType string
	MaturityDate      string
	CouponPaymentDate string
	PaymentFrequency  string
	NoPaymentsPerYear string
	CouponStartDate   string
	FinalCouponDate   string
	Floater           string
	About             string
	AboutIssuer       string
	URI               string
}

func InstFII_Job() dm.JobDefinition {
	var j dm.JobDefinition
	j.ID = "INST_FII"
	j.Name = "INST_FII"
	j.Period = "15 7-19 * * 1-5"
	j.Description = "Update Bond like Instruments from Fixed Income Investor"
	j.Type = core.Aquirer
	return j
}

func InstFII_Register(c *cron.Cron) {
	application.Schedule_Register(InstFII_Job())
	c.AddFunc(InstFII_Job().Period, func() { InstFII_Run() })
}

// RunJobRollover is a Rollover function
func InstFII_Run() {
	logs.StartJob(InstFII_Job().Name)
	var message string
	/// CONTENT STARTS

	OnPage(FIICorporateBonds, "Corporate Bonds")
	OnPage(FIIIndexLinkedCorporateBonds, "Index Linked Corporate Bonds")
	OnPage(FIIUKGilts, "Gilts")
	OnPage(FIIUKIndexLinkedGilts, "Index Linked Gilts")
	OnPage(FIIPIBS, "Interest Baring Share")
	OnPage(FIIORB, "Retail Bond")
	OnPage(FIIORBBONDS, "Retail Bond")

	/// CONTENT ENDS
	application.Schedule_Update(InstFII_Job(), message)
	logs.EndJob(InstFII_Job().Name)
}

func getInternalDate(in string) string {
	return getInternalDateGen(in, LSEDateFormat)
}

func getInternalDateGen(in string, format string) string {
	var intDate string
	if len(in) > 0 {
		//log.Println(LSEDateFormat, core.DATEFORMATSIENA
		if in != "Perpetual" {
			time, err := time.Parse(format, in)
			if err != nil {
				log.Println(err.Error())
			}

			intDate = time.Format(core.DATEFORMATSIENA)
		} else {
			intDate = ""
		}
		//log.Println(in, intDate, format)
	}
	return intDate
}

func getInternalDateFII(in string) string {
	return getInternalDateGen(in, FIIDateFormat)
}

func OnPage(link string, nitype string) {
	//log.Printf("link: %v\n", link)
	logs.Accessing(link)
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "application/json")

	response, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}

	tokenizeFIIhtml(response, nitype)

	defer response.Body.Close()

	//	bytes, err := ioutil.ReadAll(response.Body)
	//	if err != nil {
	//		log.Fatalln(err)

	return
}

func tokenizeFIIhtml(response *http.Response, nitype string) {

	inTable := false
	inTableBody := false
	inTableRow := false
	inTableCell := false
	noRows := 0
	noCols := 0
	sourceURI := ""
	//var row []string
	row := []string{"", "", "", "", "", "", "", "", "", ""}
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
						//fmt.Printf("consume t.Data: %q", t.Data)
						//fmt.Printf("inTableCell: %v\n", inTableCell)
						if inTableCell {
							//node := t.Type
							data := strings.TrimSpace(t.Data)
							//fmt.Println(inTable, inTableBody, inTableRow, data, noRows, noCols)
							if len(data) == 0 {
								data = "-"
							}
							//log.Println(data)
							row[noCols-1] = data
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
				//fmt.Printf("b4 inTableCell: %v\n", inTableCell)
				if t.Data == "td" {
					inTableCell = true
					noCols = noCols + 1
				}
				//fmt.Printf("t.Data: %q", t.Data)
				//fmt.Printf("noCols: %v\n", noCols)
				//fmt.Printf("a4 inTableCell: %v\n", inTableCell)
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
					processRow(row, noCols, sourceURI, nitype)

					//log.Print(row)

					table = append(table, row)
					noCols = 0
					row = []string{"", "", "", "", "", "", "", "", "", ""}

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

func processRow(row []string, noCols int, inURI string, nitype string) {
	processDefinition(row, noCols, inURI, nitype)
	processPrice(row, noCols)
	//	log.Println("PROCESSED", row, noCols, inURI)
}

func processDefinition(row []string, noCols int, inURI string, nitype string) {
	//	var bondRec application.AppLSEGiltsDataStoreItem
	_, bondRec, _ := dao.NegotiableInstrument_GetByID(row[3])
	bondRec.LongName = application.Translation_Lookup("NI-Name", row[2])
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
		bondRec.Issuer = application.Translation_Lookup("NI-Issuer", "UK Government")
	} else {
		bondRec.Type = nitype
	}

	if len(bondRec.Type) == 0 {
		bondRec.Type = application.Translation_Lookup("NI-Type", bondRec.Segment)
	}
	if len(bondRec.Type) == 0 {
		bondRec.Type = application.Translation_Lookup("NI-Type", "Corporate Bond")
	}

	//log.Println(bondRec.Issuer)
	if bondRec.Issuer == "" {
		bondRec.Issuer = application.Translation_Lookup("NI-Issuer", bondRec.LongName)
	}

	//fmt.Printf("bondRec: %v\n", bondRec)
	//fmt.Println(bondRec.LongName, bondRec.Sector)

	bondRec.LongName = strings.ToUpper(bondRec.LongName)

	//// TODO: Use Longname/ISIN to get ISIN info and/or lookup a siena newSienaCounterparty
	// TODO: if countarparty is not found create Firm, add to center, create counterparty (as issuer), create counterpartyimportID
	// TODO: isinlookup  (might need to go into the Dispatcher Job)
	//spew.Dump(bondRec)
	//fmt.Printf("bondRec: %v\n", bondRec)
	dao.NegotiableInstrument_StoreSystem(bondRec)
}

func getFIIEnrichment(inURI string, bondRec dm.NegotiableInstrument) dm.NegotiableInstrument {
	//log.Println("URI=" + inURI)
	req, err := http.NewRequest("GET", inURI, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "application/json")

	response, err := http.Get(inURI)
	if err != nil {
		log.Fatalln(err)
	}

	sector, couponFreq, yield, runningYield, issueAmount := tokenizeFIIenrichment(response)
	miData := getMIenrichment(bondRec.Isin, bondRec)

	bondRec.Sector = sector
	//fmt.Printf("bondRec.Sector: %v\n", bondRec.Sector)

	if bondRec.Sector == "" {
		bondRec.Sector = miData.Issuer
	}

	bondRec.Sector = application.Translation_Lookup("NI-Sector", bondRec.Sector)

	//fmt.Printf("bondRec.Sector: %v\n", bondRec.Sector)

	bondRec.PeriodOfCoupon = application.Translation_Lookup("NI-CouponPeriod", couponFreq)
	bondRec.FlatYield = strings.ReplaceAll(yield, "%", "")
	bondRec.RunningYield = strings.ReplaceAll(runningYield, "%", "")
	bondRec.IssueAmount = issueAmount

	bondRec.MinimumDenomination = miData.Denomination

	//TODO Translate into Internal SRS format
	bondRec.IssueDate = getInternalDateGen(miData.IssueDate, MIDate)
	bondRec.PaymentCouponDate = getInternalDateGen(miData.CouponPaymentDate, MIDate)

	if miData.Name != "" {
		bondRec.LongName = strings.ToUpper(application.Translation_Lookup("NI-Name", miData.Name))
	}
	if miData.Issuer != "" {
		bondRec.Issuer = application.Translation_Lookup("NI-Issuer", miData.Issuer)
	}
	if bondRec.Segment == "" {
		bondRec.Segment = application.Translation_Lookup("NI-Segment", miData.Issuer)
	}
	//log.Printf("ISIN=%q LEI=%q", bondRec.Isin, bondRec.Lei)
	bondRec.LEI, err = application.GLIEF_leiLookup(bondRec.Isin)
	if err != nil {
		log.Println(err.Error())
	}
	//log.Printf("ISIN=%q LEI=%q", bondRec.Isin, bondRec.Lei)

	//spew.Dump(bondRec)
	//spew.Dump(miData)
	return bondRec
}

func getMIenrichment(inISIN string, bondRec dm.NegotiableInstrument) MIBondData {

	req, err := http.NewRequest("GET", MIBaseURI+inISIN, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "application/json")

	response, err := http.Get(MIBaseURI + inISIN)
	if err != nil {
		log.Fatalln(err)
	}

	miData := processMIResponse(response)

	defer response.Body.Close()

	return miData
}

func processMIResponse(response *http.Response) MIBondData {
	var miData MIBondData

	miContent := make(map[string][]string)

	inTable := false
	inTableRow := false
	inKeyCell := false
	var dataKey string
	var dataValue string
	keyValueSearch := false

	noCols := 0
	var instCount = 0

	tokenizer := html.NewTokenizer(response.Body)
	for {
		tt := tokenizer.Next()
		t := tokenizer.Token()

		err := tokenizer.Err()
		if err == io.EOF {
			break
		}

		//log.Printf("tt={%v} attr={%v} type={%v} data={%v} da={%v} ks={%v}", tt.String(), t.Attr, t.Type, strings.TrimSpace(t.Data), t.DataAtom, keyValueSearch)
		switch tt {
		case html.ErrorToken:
			log.Fatalln(err.Error())
		case html.TextToken:
			if inTable {
				//		spew.Dump(t)

				if inTableRow {
					//	prevData = nil
					if inKeyCell && !keyValueSearch {
						//node := t.Type
						if strings.TrimSpace(t.Data) != "" {
							dataKey = strings.TrimSpace(t.Data)
							if dataValue == "" {
								keyValueSearch = true
							}
						}
					} else if inKeyCell && keyValueSearch {
						dataValue = strings.TrimSpace(t.Data)
						if dataValue != "" {
							miContent[dataKey] = append(miContent[dataKey], dataValue)
						}
						instCount = instCount + 1
					}

					//			log.Println("DK=", dataKey, "DV=", dataValue, "IC=", instCount, "NC=", noCols)

				}

			}
		case html.StartTagToken:

			if t.Data == "table" {
				inTable = true
			}

			if t.Data == "tr" {
				inTableRow = true
			}

			if inTableRow {
				if t.Data == "td" {
					inKeyCell = true
					noCols = noCols + 1
				}
			}

		case html.EndTagToken:

			if t.Data == "table" {
				inTable = false
			}

			if t.Data == "tr" {
				inTableRow = false
				noCols = 0
				//			row = nil
				inKeyCell = false
				keyValueSearch = false
			}

		}
	}
	// Print to check the slice's content
	//spew.Dump(table)

	miData.Coupon = getMIContent(miContent, "Coupon")
	miData.Denomination = getMIContent(miContent, "Denomination")
	miData.CouponPaymentDate = getMIContent(miContent, "Coupon Payment Date")
	miData.Name = getMIContent(miContent, "Name")
	miData.Issuer = getMIContent(miContent, "Issuer")
	miData.IssuePrice = getMIContent(miContent, "Issue Price")
	miData.MaturityDate = getMIContent(miContent, "Maturity Date")
	miData.CouponStartDate = getMIContent(miContent, "Coupon Start Date")
	miData.PaymentType = getMIContent(miContent, "Payment Type")
	miData.FinalCouponDate = getMIContent(miContent, "Final Coupon Date")
	miData.Country = getMIContent(miContent, "Country")
	miData.IssueVolume = getMIContent(miContent, "Issue Volume")
	miData.Currency = getMIContent(miContent, "Currency")
	miData.Floater = getMIContent(miContent, "Floater")
	miData.ISIN = getMIContent(miContent, "ISIN")
	miData.IssueDate = getMIContent(miContent, "Issue Date")
	miData.NoPaymentsPerYear = getMIContent(miContent, "No. of Payments per Year")

	return miData
}

func getMIContent(in map[string][]string, fetch string) string {

	var rv = ""
	if len(in[fetch]) != 0 {
		rv = in[fetch][0]
	}
	return rv
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
	//fmt.Printf("row: %v\n", row)
	//fmt.Printf("noCols: %v %v\n", noCols, len(row))
	//spew.Dump(row)
	var ratesData RatesDataStore
	//	ratesData.bid = fmt.Sprintf("%v\r\n", bondRow.CouponValue)
	ratesData.mid = row[7]
	//	ratesData.offer = fmt.Sprintf("%v\r\n", bondRow.CouponValue)
	ratesData.market = "NI"
	ratesData.tenor = ""
	ratesData.series = row[3]
	ratesData.name = application.Translation_Lookup("NI-Name", row[2])
	ratesData.class = "Market"
	ratesData.source = "FII"
	ratesData.destination = "RVNI"
	RatesDataStorePut(ratesData)
}
