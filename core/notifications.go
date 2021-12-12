package core

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gregdel/pushover"
	"github.com/spf13/viper"
)

type Notification_Config struct {
	PushoverKey   string `mapstructure:"pushoverkey"`
	PushoverToken string `mapstructure:"pushovertoken"`
}

func notification_GetConfig() (config Notification_Config, err error) {
	// get current os directory path
	pwd, _ := os.Getwd()

	//fmt.Println(pwd)
	viper.AddConfigPath(pwd + "/config/")
	viper.SetConfigName("notifications")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = viper.Unmarshal(&config)
	//spew.Dump(config)

	return config, err
}

func Notification_Emergency(messageTitle string, messageBody string) {
	cfg, _ := notification_GetConfig()
	//fmt.Printf("cfg: %v\n", cfg)
	app := pushover.New(cfg.PushoverKey)

	// Create a new recipient
	recipient := pushover.NewRecipient(cfg.PushoverToken)
	port := ApplicationProperties["port"]
	// Create the message to send
	//message := pushover.NewMessageWithTitle(messageContent, title)

	// NOTE Notification Message & Title fields are reversed (known bug in Pushover)

	fmt.Printf("SystemHostname: %v\n", SystemHostname)
	message := &pushover.Message{
		Message:     messageTitle,
		Title:       messageBody,
		Priority:    pushover.PriorityEmergency,
		URL:         "http://" + SystemHostname + ":" + port + "/",
		URLTitle:    SystemHostname,
		Timestamp:   time.Now().Unix(),
		Retry:       60 * time.Second,
		Expire:      time.Hour,
		DeviceName:  strings.ReplaceAll(SystemHostname, ".", "_"),
		CallbackURL: "http://" + SystemHostname + ":" + port + "/ACKNotification",
		Sound:       pushover.SoundCosmic,
	}

	// Send the message to the recipient
	_, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
	}

	// Print the response if you want
	//log.Println(response)
}

func Notification_Normal(messageTitle string, messageBody string) {
	cfg, _ := notification_GetConfig()
	//fmt.Printf("cfg: %v\n", cfg)
	app := pushover.New(cfg.PushoverKey)

	// Create a new recipient
	recipient := pushover.NewRecipient(cfg.PushoverToken)
	port := ApplicationProperties["port"]

	messageBody = messageBody + " - " + SystemHostname

	// NOTE Notification Message & Title fields are reversed (known bug in Pushover)
	message := &pushover.Message{
		Message:     messageTitle,
		Title:       messageBody,
		Priority:    pushover.PriorityNormal,
		URL:         "http://" + SystemHostname + ":" + port + "/",
		URLTitle:    SystemHostname,
		Timestamp:   time.Now().Unix(),
		Retry:       60 * time.Second,
		Expire:      time.Hour,
		DeviceName:  strings.ReplaceAll(SystemHostname, ".", "_"),
		CallbackURL: "http://" + SystemHostname + ":" + port + "/ACKNotification",
		Sound:       pushover.SoundCosmic,
	}

	// Send the message to the recipient
	_, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
	}

	// Print the response if you want
	//log.Println(response)
}

func Notification_High(messageTitle string, messageBody string) {
	cfg, _ := notification_GetConfig()
	//fmt.Printf("cfg: %v\n", cfg)
	app := pushover.New(cfg.PushoverKey)

	// Create a new recipient
	recipient := pushover.NewRecipient(cfg.PushoverToken)
	port := ApplicationProperties["port"]
	// Create the message to send
	//message := pushover.NewMessageWithTitle(messageContent, title)
	//fmt.Printf("SystemHostname: %v\n", SystemHostname)
	messageBody = messageBody + " - " + SystemHostname

	// NOTE Notification Message & Title fields are reversed (known bug in Pushover)

	message := &pushover.Message{
		Message:     messageTitle,
		Title:       messageBody,
		Priority:    pushover.PriorityHigh,
		URL:         "http://" + SystemHostname + ":" + port + "/",
		URLTitle:    SystemHostname,
		Timestamp:   time.Now().Unix(),
		Retry:       60 * time.Second,
		Expire:      time.Hour,
		DeviceName:  strings.ReplaceAll(SystemHostname, ".", "_"),
		CallbackURL: "http://" + SystemHostname + ":" + port + "/ACKNotification",
		Sound:       pushover.SoundCosmic,
	}

	// Send the message to the recipient
	_, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
	}

	// Print the response if you want
	//log.Println(response)
}

func Notification_Low(messageTitle string, messageBody string) {
	cfg, _ := notification_GetConfig()
	//fmt.Printf("cfg: %v\n", cfg)
	app := pushover.New(cfg.PushoverKey)

	// Create a new recipient
	recipient := pushover.NewRecipient(cfg.PushoverToken)
	port := ApplicationProperties["port"]
	// Create the message to send
	//message := pushover.NewMessageWithTitle(messageContent, title)
	//fmt.Printf("SystemHostname: %v\n", SystemHostname)
	messageBody = messageBody + " - " + SystemHostname

	// NOTE Notification Message & Title fields are reversed (known bug in Pushover)

	message := &pushover.Message{
		Message:     messageTitle,
		Title:       messageBody,
		Priority:    pushover.PriorityLow,
		URL:         "http://" + SystemHostname + ":" + port + "/",
		URLTitle:    SystemHostname,
		Timestamp:   time.Now().Unix(),
		Retry:       60 * time.Second,
		Expire:      time.Hour,
		DeviceName:  strings.ReplaceAll(SystemHostname, ".", "_"),
		CallbackURL: "http://" + SystemHostname + ":" + port + "/ACKNotification",
		Sound:       pushover.SoundCosmic,
	}

	// Send the message to the recipient
	_, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Panic(err)
	}

	// Print the response if you want
	//log.Println(response)
}
