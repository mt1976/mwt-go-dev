package core

import (
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/leekchan/accounting"
	"github.com/mt1976/mwt-go-dev/logs"
)

// Converts a siena style date, as a user readable date
func sienaDateToUserDate(in string) string {

	t, err := time.Parse(DATEFORMATSIENA, in)
	if err != nil {
		log.Println(err.Error())
	}

	ext := t.Format(DATEFORMATUSER)

	return ext
}

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

// GetURLparam returns a selected parmeter value from the a given URI
func GetURLparam(r *http.Request, paramID string) string {
	//fmt.Println(paramID)
	//fmt.Println(r.URL)
	key := r.FormValue(paramID)
	//log.Printf("URL Parameter : Key=%q Value=%q", paramID, string(key))
	logs.Information("URL Parameter :", fmt.Sprintf("Key=%q Value=%q", paramID, string(key)))
	return key
}

// ArrToString converts an array of strings to a printable string
func ArrToString(strArray []string) string {
	return strings.Join(strArray, "\n")
}

// RemoveContents clears the contents of a specified directory
func RemoveContents(dir string) error {
	log.Println("TRASH", dir)
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		log.Println(err)
		return err
	}
	//	fmt.Println("do Clear", files)
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return err
}

// GetTemplateID returns the template to use when providing HTML page templates
func GetTemplateID(tmpl string, userRole string) string {
	returnTemplate := ""
	templateName := "html/" + tmpl + ".html"
	roleTemplate := "html/versions/roles/" + userRole + "/" + tmpl + ".html"
	versionTemplate := "html/versions/" + tmpl + ".html"

	//logs.Information("Template Checks:", tmpl+" "+templateName+" "+roleTemplate+" "+versionTemplate)
	//log.Println("Testing", roleTemplate, FileExists(roleTemplate))
	//log.Println("Testing", templateName, FileExists(templateName))
	if FileExists(templateName) {
		returnTemplate = templateName
	}

	if FileExists(versionTemplate) {
		returnTemplate = versionTemplate
	}
	if userRole != "" {
		if FileExists(roleTemplate) {
			returnTemplate = roleTemplate
		}
	}
	//log.Printf("Using Template: Source %q", templateName)
	logs.Template(returnTemplate)
	return returnTemplate
}

// GetMenuID returns the menu template to use when providing HTML meny templates
func GetMenuID(tmpl string, userRole string) string {
	templateName := "config/menu/" + tmpl + ".json"
	roleTemplate := templateName
	if len(userRole) != 0 {
		roleTemplate = "config/menu/" + userRole + "/" + tmpl + ".json"
	}
	//log.Println("Testing", roleTemplate, FileExists(roleTemplate))
	//log.Println("Testing", templateName, FileExists(templateName))
	if FileExists(roleTemplate) {
		templateName = roleTemplate
	}
	//log.Printf("Using Menu    : Source %q", templateName)
	logs.Menu(templateName)
	return templateName
}

// GetNavigationID returns the navigation pane template to use when providing navigation templates
func GetNavigationID(inUserRole string) string {
	templateName := "../assets/navigation.html"
	roleTemplate := "../assets/" + inUserRole + "_navigation.html"
	//log.Println("Testing", templateName, FileExists(templateName))
	//log.Println("Testing", roleTemplate, FileExists(roleTemplate))
	if FileExists(roleTemplate) {
		//templateName = roleTemplate
	}
	//log.Println("NAVIGATION", templateName)
	return templateName
}

// FileExists returns true if the specified file existing on the filesystem
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// PickEpochToDateTimeString converts a PICK format date into a readable string.
func PickEpochToDateTimeString(pickEpoch string) string {
	//pickEpochLayout := "20060102T150405"
	t, err := time.Parse(DATEFORMATPICK, pickEpoch)
	if err != nil {
		log.Println(err)
	}
	tString := t.Format(time.RFC1123)
	return tString
}

// StrArrayToString converst a string array into a string
func StrArrayToString(inArray []string) string {
	return StrArrayToStringWithSep(inArray, "\n")
}

//StrArrayToStringWithSep converts a string array to a string using a given separator
func StrArrayToStringWithSep(inArray []string, inSep string) string {

	outString := ""
	noRows := len(inArray)
	for ii := 0; ii < noRows; ii++ {
		outString += inArray[ii] + inSep
	}
	return outString
}

// QmBundleAdd adds data to a "bundle" of data to be passed in a request message
func QmBundleAdd(inBundle []string, name string, value string) []string {
	return append(inBundle, name+"¡"+value)
}

// QmBundleToString convers a "bundle" from a request message to a readable string
func QmBundleToString(inBundle []string) string {
	return StrArrayToStringWithSep(inBundle, ";")
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

// GetIPAdress returns the current IP address
func GetIPAdress(r *http.Request) string {
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

// ReadUserIP returns the end-users IP address
func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

// GetLocalIP returns the local IP address
func GetLocalIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	//handle err...

	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.String()
}

// SqlDateToHTMLDate returns a string version of a SQL Date
func SqlDateToHTMLDate(inDate string) (outDate string) {
	var rtnDate string
	if inDate != "" {
		rtnDate = inDate[0:10]
	}
	return rtnDate
}

// FormatCurrency returns a formated string version of a CCY amount
func FormatCurrency(inAmount string, inCCY string) string {
	ac := accounting.Accounting{Symbol: inCCY, Precision: 2, Format: "%v", FormatNegative: "-%v", FormatZero: "\u2013 ;\u2013"}
	bum, _ := strconv.ParseFloat(inAmount, 64)
	return ac.FormatMoney(bum)
}

// FormatCurrencyFull returns a formated string version of a CCY amount to 7dps
func FormatCurrencyFull(inAmount string, inCCY string) string {
	prec, _ := strconv.Atoi("7")
	ac := accounting.Accounting{Symbol: inCCY, Precision: prec, Format: "%v", FormatNegative: "-%v", FormatZero: "\u2013 \u2013"}
	bum, _ := strconv.ParseFloat(inAmount, 64)
	return ac.FormatMoney(bum)
}

// FormatCurrencyDps returns a formated string version of a CCY amount to a given DPS
func FormatCurrencyDps(inAmount string, inCCY string, inPrec string) string {
	prec, _ := strconv.Atoi(inPrec)
	ac := accounting.Accounting{Symbol: inCCY, Precision: prec, Format: "%v", FormatNegative: "-%v", FormatZero: "\u2013 \u2013"}
	bum, _ := strconv.ParseFloat(inAmount, 64)
	return ac.FormatMoney(bum)
}

func GetHostIP() string {
	var ip string
	ip = ""
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrs {
		//fmt.Printf("%d %v\n", i, addr)
		ip = addr.String()
		return ip
	}
	return ""
}

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func GetIncomingRequestIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func serviceMessage(i string) {
	//msg := "Servicing     : %q"
	//log.Printf(msg, i)
	logs.Servicing(i)
}
func serviceMessageAction(i string, act string, id string) {
	msgSuffix := fmt.Sprintf("%s %s %q", i, act, id)
	serviceMessage(msgSuffix)
}
func ServiceMessageAction(i string, act string, id string) {
	serviceMessageAction(i, act, id)
}

func ServiceMessage(i string) {
	serviceMessage(i)
}

func ReadDataFile(fileName string, path string) (string, error) {
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName
	if len(path) != 0 {
		filePath = pwd + path + "/" + fileName
	}

	// Check it exists - If not create it
	if !(FileExists(filePath)) {
		WriteDataFile(fileName, path, "")
	}

	//log.Println("Read          :", filePath)
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// Convert []byte to string and print to screen
	return string(content), err
}

func GetDirectoryContentAbsolute(path string) ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(path + "/")
	if err != nil {
		log.Fatal(err)
	}
	return files, err
}

func ReadDataFileAbsolute(fileName string, path string) ([]byte, string, error) {

	pwd, _ := os.Getwd()
	//first char is a . replace with path
	path = strings.Replace(path, "..", pwd, 1)
	path = strings.Replace(path, ".", pwd, 1)

	absFileName := path + "/" + fileName
	logs.Information("AbsPath       :", absFileName)
	// Check it exists - If not create it
	if !(FileExists(absFileName)) {
		WriteDataFileAbsolute(fileName, path, "")
	}

	//log.Println("Read          :", filePath)
	// Read entire file content, giving us little control but
	// making it very simple. No need to close the file.
	content, err := ioutil.ReadFile(absFileName)
	if err != nil {
		log.Fatal(err)
	}
	// Convert []byte to string and print to screen
	return content, string(content), err
}

func WriteDataFile(fileName string, path string, content string) int {
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName
	if len(path) != 0 {
		filePath = pwd + path + "/" + fileName
	}
	//log.Println("Write         :", filePath)

	message := []byte(content)
	err := ioutil.WriteFile(filePath, message, 0644)
	if err != nil {
		log.Fatal(err)
		return -1
	}

	//	log.Println("File Write : " + fileName + " in " + path + "[" + filePath + "]")
	logs.Saving(fileName, filePath)
	return 1
}

func WriteDataFileAbsolute(fileName string, path string, content string) int {

	absFileName := path + "/" + fileName

	message := []byte(content)
	err := ioutil.WriteFile(absFileName, message, 0644)
	if err != nil {
		log.Fatal(err)
		return -1
	}

	//log.Println("File Write : " + fileName + " in " + path + "[" + absFileName + "]")
	logs.Saving(fileName, absFileName)
	return 1
}

func DeleteDataFileAbsolute(fileName string, path string) int {

	absFileName := path + "/" + fileName
	if FileExists(absFileName) {
		var err = os.Remove(absFileName)
		if err != nil {
			log.Fatal(err.Error())
			return -1
		}
	}
	log.Println("File Deletion : " + fileName + " in " + path + "[" + absFileName + "]")
	return 1

}

func DeleteDataFile(fileName string, path string) int {
	pwd, _ := os.Getwd()
	filePath := pwd + "/" + fileName
	if len(path) != 0 {
		filePath = pwd + path + "/" + fileName
	}
	//log.Println("Delete        :", filePath)

	// delete file

	if FileExists(filePath) {
		var err = os.Remove(filePath)
		if err != nil {
			log.Fatal(err.Error())
			return -1
		}
	}
	log.Println("File Deletion : " + fileName + " in " + path + "[" + filePath + "]")
	return 1
}

// CalculateSpotDate(inTime invalid type)
func CalculateSpotDate(inTime time.Time) time.Time {
	spot := inTime.AddDate(0, 0, 2)
	return wibbleDate(spot)
}

func wibbleDate(inDate time.Time) time.Time {
	if inDate.Weekday() == time.Saturday {
		inDate = inDate.AddDate(0, 0, 2)
	}
	if inDate.Weekday() == time.Sunday {
		inDate = inDate.AddDate(0, 0, 1)
	}
	return inDate
}

// CalculateSpotDate(inTime invalid type)
func CalculateTenorDate(inTime time.Time, inMonth string) time.Time {
	month, _ := strconv.Atoi(inMonth)
	spot := inTime.AddDate(0, month, 0)
	return wibbleDate(spot)
}

func CalculateFirstDateOfYear(inTime time.Time) time.Time {
	// Assuking 1st Jan is a holiday therefore first day is 2, then wibble the date.
	tempDate := time.Date(inTime.Year(), 1, 2, 0, 0, 0, inTime.Nanosecond(), inTime.Location())
	return wibbleDate(tempDate)
}

func Logit(actionType string, data string) {
	//_, caller, _, _ := runtime.Caller(1)
	//outcall := strings.Split(caller, "/")
	///depth := len(outcall) - 1
	//depth2 := depth - 1
	//log.Println(len(outcall), depth, depth2)
	//callerName := outcall[depth2] + "/" + outcall[depth]
	log.Println("Information   :", data)
	logs.Information(data, "")
	//	log.Println(callerName, actionType, data)
}

func RemoveSpecialChars(in string) string {
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	newStr := reg.ReplaceAllString(in, "-")
	return newStr
}

func finAbbrToNumeric(str string) int {
	str = strings.ToUpper(str)
	number := ""
	number = strings.ReplaceAll(str, "$", "")
	number = strings.ReplaceAll(str, "€", "")
	number = strings.ReplaceAll(str, "£", "")
	fact := strings.ToUpper(number[len(number)-1:])
	number = strings.ReplaceAll(str, "M", "")
	number = strings.ReplaceAll(str, "K", "")
	number = strings.ReplaceAll(str, "T", "")
	number = strings.ReplaceAll(str, "B", "")

	intNum, err := strconv.Atoi(number)
	if err != nil {
		log.Println(err.Error())
	}

	var retNum int
	switch fact {
	case "T":
		retNum = intNum * 1000
	case "K":
		retNum = intNum * 1000
	case "M":
		retNum = intNum * 1000000
	case "B":
		retNum = intNum * 1000000000
	default:
		retNum = intNum
	}

	return retNum
}

//Convert time.Time to string
func TimeToString(t time.Time) string {
	//fmt.Printf("t: %v\n", t)
	return t.Format(DATEFORMATSIENA)
}
