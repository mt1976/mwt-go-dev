package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	support "github.com/mt1976/mwt-go-dev/appsupport"
)

//loginPage is cheese
type loginPage struct {
	AppName          string
	UserName         string
	UserPassword     string
	WebServerVersion string
	LicenceType      string
	LicenceLink      string
	ResponseError    string
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	tmpl := "login"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	appName := wctProperties["appname"]

	appServerVersion := wctProperties["releaseid"] + " [r" + wctProperties["releaselevel"] + "-" + wctProperties["releasenumber"] + "]"

	loginPageContent := loginPage{
		AppName:          appName,
		UserName:         "",
		UserPassword:     "",
		WebServerVersion: appServerVersion,
		LicenceType:      wctProperties["licname"],
		LicenceLink:      wctProperties["liclink"],
		ResponseError:    gSecurityViolation,
	}

	//fmt.Println("Page Data", loginPageContent)

	//thisTemplate:= support.GetTemplateID(tmpl,gUserRole)
	t, err := template.ParseFiles(support.GetTemplateID(tmpl, gUserRole))
	log.Println(t, support.GetTemplateID(tmpl, gUserRole), err)
	t.Execute(w, loginPageContent)

}

func valLoginHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	//tmpl := "login"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	uName := r.FormValue("username")
	uPassword := r.FormValue("password")
	appToken := wctProperties["applicationtoken"]
	requestID := uuid.New()
	var requestPayload []string

	requestPayload = support.QmBundleAdd(requestPayload, "username", uName)
	requestPayload = support.QmBundleAdd(requestPayload, "password", uPassword)
	requestPayload = support.QmBundleAdd(requestPayload, "apptoken", appToken)

	requestPayload = support.QmBundleAdd(requestPayload, "ipUser", support.ReadUserIP(r))
	requestPayload = support.QmBundleAdd(requestPayload, "hostUser", r.Host)

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	requestPayload = support.QmBundleAdd(requestPayload, "ipWebServer", support.GetLocalIP())
	requestPayload = support.QmBundleAdd(requestPayload, "hostWebServer", hostname)

	//fmt.Println("requestPayload", requestPayload)
	//fmt.Println("BUNDLE", support.QmBundleToString(requestPayload))

	loginRequest := buildRequestMessage(requestID.String(), "LOGIN", "", "", support.QmBundleToString(requestPayload), wctProperties)

	//fmt.Println("loginRequest", loginRequest)
	//fmt.Println("SEND MESSAGE")
	sendRequest(loginRequest, requestID.String(), wctProperties)

	loginResponse := getResponseAsync(requestID.String(), wctProperties, r)
	//fmt.Println("loginResponse", loginResponse)

	//outString := ""
	//noRows := len(loginResponse.ResponseContent.ResponseContentRow)

	responseCode := loginResponse.ResponseStatus
	//fmt.Println("Response Code", responseCode)
	//fmt.Println("SESSION", loginResponse.ResponseContent.ResponseContentRow[1])
	newToken := loginResponse.ResponseContent.ResponseContentRow[0]
	//fmt.Println("Response Content", newToken)

	log.Println(responseCode, newToken)
	//todo Encryp password etc

	if loginResponse.ResponseStatus == "200" {

		newUUID := loginResponse.ResponseContent.ResponseContentRow[1]
		gUserRole = loginResponse.ResponseContent.ResponseContentRow[2]
		gUserNavi = support.GetNavigationID(gUserRole)
		gUserKnowAs = loginResponse.ResponseContent.ResponseContentRow[3]
		log.Println("ACCESS GRANTED", responseCode, uName, gUserRole)
		gUserName = uName
		gSessionToken = newToken
		gUUID = newUUID
		gSecurityViolation = ""
		homePageHandler(w, r)
	} else {
		gSecurityViolation = loginResponse.ResponseContent.ResponseContentRow[0]
		log.Println("SECURITY INCIDENT", responseCode, gSecurityViolation)
		loginHandler(w, r)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {

	wctProperties := support.GetProperties(APPCONFIG)
	//tmpl := "login"
	inUTL := r.URL.Path
	w.Header().Set("Content-Type", "text/html")
	log.Println("Servicing :", inUTL)

	requestID := uuid.New()

	logoutRequest := buildRequestMessage(requestID.String(), "LOGOUT", "", "", "", wctProperties)

	//fmt.Println("logoutRequest", logoutRequest)
	//fmt.Println("SEND MESSAGE")
	sendRequest(logoutRequest, requestID.String(), wctProperties)

	logoutResponse := getResponseAsync(requestID.String(), wctProperties, r)
	//fmt.Println("logoutResponse", logoutResponse)

	//outString := ""

	//responseCode := logoutResponse.ResponseStatus
	log.Println("LOGOUT", logoutResponse.ResponseStatus, gSessionToken)
	gSessionToken = ""
	gUUID = ""
	gSecurityViolation = ""
	gUserRole = ""
	//fmt.Println("SESSION", loginResponse.ResponseContent.ResponseContentRow[1])
	//todo Encryp password etc

	loginHandler(w, r)
}
