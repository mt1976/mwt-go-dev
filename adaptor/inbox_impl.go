package adaptor

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/mt1976/mwt-go-dev/core"
	dm "github.com/mt1976/mwt-go-dev/datamodel"
	logs "github.com/mt1976/mwt-go-dev/logs"
)

const (
	Inbox_FormatXML            = "XML"
	Inbox_TimeoutAction_Notify = "Notify"
)

func Inbox_SendMail(mailTo string, mailSubject string, mailContent string, mailFrom string, source string) error {
	//logs.Success("Message Sent:" + id)

	now := time.Now()

	var mail dm.Inbox
	mail.MailId = getNewFileID()
	mail.MailTo = mailTo
	mail.MailSubject = mailSubject
	mail.MailContent = mailContent
	mail.MailFrom = mailFrom
	mail.MailSent = now.Format(core.DATETIME)
	mail.MailSource = source

	var err error
	//fmt.Printf("msItem: %v\n", msItem)
	//spew.Dump(msItem)
	//err = dao.ExternalInbox_StoreSystem(msItem)

	json_data, err := json.Marshal(mail)
	//fmt.Println(string(json_data))
	if err != nil {
		return err
	}

	// get current ip address
	ip := "localhost"
	uri := core.ApplicationProperties["protocol"] + "://" + ip + ":" + core.ApplicationProperties["port"] + dm.Inbox_Path

	//logs.Information("uri", uri)

	resp, err := http.Post(uri, "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	//var res map[string]interface{}
	//logs.Information("response Status:", resp.Status)
	//json.NewDecoder(resp.Body).Decode(&res)
	if resp.StatusCode != 200 {
		logs.Error("Error: "+strconv.Itoa(resp.StatusCode), nil)
		return err
	}
	if resp.StatusCode == 404 {
		logs.Warning("Unable to Store Message:" + mail.MailId)
	}

	logs.Post(uri + mail.MailId)
	// logs.Accessing(uri + id)

	// params := url.Values{}
	// params.Add(dm.ExternalInbox_QueryString, id)

	// //	fmt.Println(res["json"])
	// gesp, err2 := http.Get(uri + "?" + params.Encode())

	// if err2 != nil {
	// 	log.Fatal(err2)
	// }

	// defer gesp.Body.Close()

	// gbody, err2 := ioutil.ReadAll(gesp.Body)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// //logs.Information("response Status:", gesp.Status)

	// //spew.Dump(gesp)
	// //json.NewDecoder(resp.Body).Decode(&res)
	// if gesp.StatusCode != 200 {
	// 	logs.Error("Error: "+strconv.Itoa(gesp.StatusCode), nil)
	// 	return err
	// }
	// if gesp.StatusCode == 404 {
	// 	logs.Warning("Unable to Store Message:" + id)
	// }

	// //fmt.Printf("gesp.Body: %v\n", gesp.Body)
	// //fmt.Println("RESPONSE>>" + string(gbody))

	// var data *dm.ExternalMessage
	// json.Unmarshal(gbody, &data)

	//json.NewDecoder(gesp.Body).Decode(data)

	//fmt.Printf("Results: %v\n", data)
	//spew.Dump(data)
	return nil
}
