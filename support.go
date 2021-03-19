package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/leekchan/accounting"
)

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/favicon.ico")
}

func favicon16Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/favicon-16x16.png")
}

func favicon32Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/favicon-32x32.png")
}

func faviconManifestHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "site.webmanifest")
}

func faviconBrowserConfigHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "browserconfig.xml")
}

func shutdownHandler(w http.ResponseWriter, r *http.Request) {
	wctProperties := getProperties(CONST_CONFIG_FILE)

	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	requestID := uuid.New()

	log.Println("Servicing :", inUTL)

	requestMessage := buildRequestMessage(requestID.String(), "SHUTDOWN", "", "", "", wctProperties)

	fmt.Println("requestMessage", requestMessage)
	fmt.Println("SEND MESSAGE")
	sendRequest(requestMessage, requestID.String(), wctProperties)
	m := http.NewServeMux()

	s := http.Server{Addr: wctProperties["port"], Handler: m}
	s.Shutdown(context.Background())
	//	r.URL.Path = "/viewResponse?uuid=" + requestID.String()
	//	viewResponseHandler(w, r)

}

func getURLparam(r *http.Request, paramID string) string {
	//fmt.Println(paramID)
	//fmt.Println(r.URL)
	key := r.FormValue(paramID)
	log.Println("Url Param '" + paramID + "' is: " + string(key))
	return key
}

func doSnooze(inPollingInterval string) {
	pollingInterval, _ := strconv.Atoi(inPollingInterval)
	log.Println("Snoooze... Zzzzzz.... ", pollingInterval)
	time.Sleep(time.Duration(pollingInterval) * time.Second)
}

func arrToString(strArray []string) string {
	return strings.Join(strArray, "\n")
}

//RemoveContents has a comment
func RemoveContents(dir string) error {
	log.Println("TRASH", dir)
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	//	fmt.Println("do Clear", files)
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return err
}

func clearQueuesHandler(w http.ResponseWriter, r *http.Request) {

	//var propertiesFileName = "config/properties.cfg"
	wctProperties := getProperties(CONST_CONFIG_FILE)
	//	tmpl := "viewResponse"
	inUTL := r.URL.Path
	//requestID := uuid.New()

	log.Println("Servicing :", inUTL)

	//fmt.Println("delivPath", wctProperties["deliverpath"])
	err1 := RemoveContents(wctProperties["deliverpath"])
	if err1 != nil {
		fmt.Println(err1)
	}

	//fmt.Println("recPath", wctProperties["receivepath"])

	err2 := RemoveContents(wctProperties["receivepath"])
	if err2 != nil {
		fmt.Println(err2)
	}
	//fmt.Println("procPath", wctProperties["processedpath"])

	err3 := RemoveContents(wctProperties["processedpath"])
	if err3 != nil {
		fmt.Println(err3)
	}

	homePageHandler(w, r)

}

func getTemplateID(tmpl string) string {
	templateName := "html/" + tmpl + ".html"
	roleTemplate := "html" + gUserRole + "/" + tmpl + ".html"
	log.Println("Testing", roleTemplate, fileExists(roleTemplate))
	log.Println("Testing", templateName, fileExists(templateName))
	if fileExists(roleTemplate) {
		//	templateName = roleTemplate
	}
	log.Println("TEMPLATE", templateName)
	return templateName
}

func getNavigationID(inUserRole string) string {
	templateName := "../assets/navigation.html"
	roleTemplate := "../assets" + inUserRole + "_navigation.html"
	log.Println("Testing", templateName, fileExists(templateName))
	log.Println("Testing", roleTemplate, fileExists(roleTemplate))
	if fileExists(roleTemplate) {
		//templateName = roleTemplate
	}
	log.Println("NAVIGATION", templateName)
	return templateName
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func pickEpochToDateTimeString(pickEpoch string) string {
	//pickEpochLayout := "20060102T150405"
	t, err := time.Parse(wctEpochDateFormat, pickEpoch)
	if err != nil {
		fmt.Println(err)
	}
	tString := t.Format(time.RFC1123)
	return tString
}

func strArrayToString(inArray []string) string {

	return strArrayToStringWithSep(inArray, "\n")
}

func strArrayToStringWithSep(inArray []string, inSep string) string {

	outString := ""
	noRows := len(inArray)
	for ii := 0; ii < noRows; ii++ {
		outString += inArray[ii] + inSep
	}
	return outString
}

func qmBundleAdd(inBundle []string, name string, value string) []string {
	return append(inBundle, name+"¡"+value)
}

func qmBundleToString(inBundle []string) string {
	return strArrayToStringWithSep(inBundle, ";")
}

//ipRange - a structure that holds the start and end of a range of ip addresses
type ipRange struct {
	start net.IP
	end   net.IP
}

// inRange - check to see if a given ip address is within a range given
func inRange(r ipRange, ipAddress net.IP) bool {
	// strcmp type byte comparison
	if bytes.Compare(ipAddress, r.start) >= 0 && bytes.Compare(ipAddress, r.end) < 0 {
		return true
	}
	return false
}

var privateRanges = []ipRange{
	ipRange{
		start: net.ParseIP("10.0.0.0"),
		end:   net.ParseIP("10.255.255.255"),
	},
	ipRange{
		start: net.ParseIP("100.64.0.0"),
		end:   net.ParseIP("100.127.255.255"),
	},
	ipRange{
		start: net.ParseIP("172.16.0.0"),
		end:   net.ParseIP("172.31.255.255"),
	},
	ipRange{
		start: net.ParseIP("192.0.0.0"),
		end:   net.ParseIP("192.0.0.255"),
	},
	ipRange{
		start: net.ParseIP("192.168.0.0"),
		end:   net.ParseIP("192.168.255.255"),
	},
	ipRange{
		start: net.ParseIP("198.18.0.0"),
		end:   net.ParseIP("198.19.255.255"),
	},
}

// isPrivateSubnet - check to see if this ip is in a private subnet
func isPrivateSubnet(ipAddress net.IP) bool {
	// my use case is only concerned with ipv4 atm
	if ipCheck := ipAddress.To4(); ipCheck != nil {
		// iterate over all our ranges
		for _, r := range privateRanges {
			// check if this ip is in a private range
			if inRange(r, ipAddress) {
				return true
			}
		}
	}
	return false
}

func getIPAdress(r *http.Request) string {
	var ipAddress string
	for _, h := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		for _, ip := range strings.Split(r.Header.Get(h), ",") {
			// header can contain spaces too, strip those out.
			ip = strings.TrimSpace(ip)
			realIP := net.ParseIP(ip)
			if !realIP.IsGlobalUnicast() || isPrivateSubnet(realIP) {
				// bad address, go to next
				continue
			} else {
				ipAddress = ip
				goto Done
			}
		}
	}
Done:
	return ipAddress
}

func readUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func getLocalIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	//handle err...

	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.String()
}

func sqlDateToHTMLDate(inDate string) (outDate string) {
	var rtnDate string
	if inDate != "" {
		rtnDate = inDate[0:10]
	}
	return rtnDate
}

func formatCurrency(inAmount string, inCCY string) string {
	ac := accounting.Accounting{Symbol: inCCY, Precision: 2, Format: "%v", FormatNegative: "-%v", FormatZero: "\u2013 ;\u2013"}
	bum, _ := strconv.ParseFloat(inAmount, 64)
	return ac.FormatMoney(bum)
}

func formatCurrencyFull(inAmount string, inCCY string) string {

	thisConnection, _ := sienaConnect()
	_, ccyData, _ := getSienaCurrency(thisConnection, inCCY)
	prec, _ := strconv.Atoi(ccyData.AmountDP)
	ac := accounting.Accounting{Symbol: inCCY, Precision: prec, Format: "%v", FormatNegative: "-%v", FormatZero: "\u2013 \u2013"}
	bum, _ := strconv.ParseFloat(inAmount, 64)
	return ac.FormatMoney(bum)
}

func formatCurrencyDps(inAmount string, inCCY string, inPrec string) string {
	prec, _ := strconv.Atoi(inPrec)
	ac := accounting.Accounting{Symbol: inCCY, Precision: prec, Format: "%v", FormatNegative: "-%v", FormatZero: "\u2013 \u2013"}
	bum, _ := strconv.ParseFloat(inAmount, 64)
	return ac.FormatMoney(bum)
}
