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

func Notification_Test() {
	Notification_Normal("normal message", "normal title")
	Notification_High("high message", "high title")
	Notification_Emergency("emergancy message", "emergancy title")
	Notification_Low("low message", "low title")
}

func Notification_Emergency(messageContent string, title string) {
	cfg, _ := notification_GetConfig()
	//fmt.Printf("cfg: %v\n", cfg)
	app := pushover.New(cfg.PushoverKey)

	// Create a new recipient
	recipient := pushover.NewRecipient(cfg.PushoverToken)
	port := ApplicationProperties["port"]
	// Create the message to send
	//message := pushover.NewMessageWithTitle(messageContent, title)
	fmt.Printf("SystemHostname: %v\n", SystemHostname)
	message := &pushover.Message{
		Message:     messageContent,
		Title:       title,
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

func Notification_Normal(messageContent string, title string) {
	cfg, _ := notification_GetConfig()
	//fmt.Printf("cfg: %v\n", cfg)
	app := pushover.New(cfg.PushoverKey)

	// Create a new recipient
	recipient := pushover.NewRecipient(cfg.PushoverToken)
	port := ApplicationProperties["port"]
	// Create the message to send
	//message := pushover.NewMessageWithTitle(messageContent, title)
	//fmt.Printf("SystemHostname: %v\n", SystemHostname)
	message := &pushover.Message{
		Message:     messageContent,
		Title:       title,
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

func Notification_High(messageContent string, title string) {
	cfg, _ := notification_GetConfig()
	//fmt.Printf("cfg: %v\n", cfg)
	app := pushover.New(cfg.PushoverKey)

	// Create a new recipient
	recipient := pushover.NewRecipient(cfg.PushoverToken)
	port := ApplicationProperties["port"]
	// Create the message to send
	//message := pushover.NewMessageWithTitle(messageContent, title)
	//fmt.Printf("SystemHostname: %v\n", SystemHostname)
	message := &pushover.Message{
		Message:     messageContent,
		Title:       title,
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

func Notification_Low(messageContent string, title string) {
	cfg, _ := notification_GetConfig()
	//fmt.Printf("cfg: %v\n", cfg)
	app := pushover.New(cfg.PushoverKey)

	// Create a new recipient
	recipient := pushover.NewRecipient(cfg.PushoverToken)
	port := ApplicationProperties["port"]
	// Create the message to send
	//message := pushover.NewMessageWithTitle(messageContent, title)
	//fmt.Printf("SystemHostname: %v\n", SystemHostname)
	message := &pushover.Message{
		Message:     messageContent,
		Title:       title,
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
