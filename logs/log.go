package logs

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/TwiN/go-color"
	colour "github.com/TwiN/go-color"
	viper "github.com/spf13/viper"
	"golang.org/x/term"
)

const (
	log_Warning       = "Warning "
	log_Info          = "Information"
	log_MUX           = "Publish"
	log_Error         = "Error"
	log_Success       = "Success"
	log_Panic         = "Panic"
	log_Header        = ""
	log_Poke          = "Fingering"
	log_Schedule      = "Scheduler"
	log_ScheduleStart = "Job Start"
	log_ScheduleEnd   = "Job Ends"
	log_System        = "System"
	log_Accessing     = "Accessing"
	log_Servicing     = "Servicing"
	log_Menu          = "Menu"
	log_Template      = "Template"
	log_Database      = "Database"
	log_URI           = "URI"
	log_Skipping      = "Skipping"
	log_Created       = "Created"
	log_Default       = "Default"
	log_Processing    = "Processing"
	log_Query         = "Query"
	log_Result        = "Result"
	log_Storing       = "Storing"
	log_Generating    = "Generating"
	log_Post          = "http POST"
	log_Save          = "Writing"
	log_Event         = "Event"
	log_Callout       = "Callout"

	ColorReset        = "\033[0m"
	ColorRed          = "\033[31m"
	ColorGreen        = "\033[32m"
	ColorYellow       = "\033[33m"
	ColorBlue         = "\033[34m"
	ColourMagenta     = "\033[35m"
	ColorCyan         = "\033[36m"
	ColorWhite        = "\033[37m"
	ColorBoldYellow   = "\033[1m\033[33m"
	Character_MapTo   = "⇄"
	Character_Job     = "⚙️"
	Character_Heart   = "🫀"
	Character_Poke    = "👉"
	Character_Time    = "🕒"
	Character_Break   = "≫"
	Character_Tick    = "☑️"
	Character_Warning = "⚠️"
	Character_Bike    = "🚴‍♂️"
	Character_Skip    = "⏭️"
	Character_Created = "💾"
	Character_Query   = "🔎"
	Character_Result  = "?"
	Character_Storing = "📀"
	Character_Event   = "🗂"
	Character_Callout = "📞"
)

type Config struct {
	Verbose     string `mapstructure:"VERBOSE"`
	VerboseMode bool
}

var CFG Config

func init() {
	CFG, _ = getConfig()
	//spew.Dump(CFG)
}

func Poke(w string, v string) {
	msg_raw(log_Poke, w, v, colour.Yellow)
}

func Success(s string) {
	//msg_done(s)
	msg_raw(log_Success, s, Character_Tick, colour.Green)
}

func Event(s string) {
	//msg_done(s)
	msg_raw(log_Event, s, Character_Event, colour.Green)
}

func Skipping(s string) {
	//msg_done(s)
	msg_raw(log_Skipping, s, Character_Skip, colour.Yellow)
}

func Created(s string) {
	//msg_done(s)
	msg_raw(log_Created, s, Character_Created, colour.Blue)
}

func Generate(s string) {
	//msg_done(s)
	msg_raw(log_Generating, s, Character_Created, colour.Blue)
}

func System(s string) {
	//msg_done(s)
	msg_raw(log_System, s, "", colour.Gray)
}

func Default(s string, w string) {
	//msg_done(s)
	msg_raw(log_Default, s, w, colour.Purple)
}

func Information(w string, v string) {
	//msg_info(w, v)
	if len(v) == 0 {
		msg_raw(log_Info, w, "", colour.Reset)
	} else {
		msg_raw(log_Info, w+" =", color.Bold+v, colour.Reset)
	}
	//msg_raw(log_Info, w, v, colour.Reset)

}

func Processing(s string) {
	//msg_done(s)
	msg_raw(log_Processing, s, "", colour.White)
}

func Schedule(w string) {
	//msg_info(w, v)
	msg_raw(log_Schedule, w, "", colour.Cyan)

}

func URI(w string) {
	//msg_info(w, v)
	msg_raw(log_URI, colour.Bold+colour.Purple+w, "", colour.Green)
}

func Servicing(w string) {
	//msg_info(w, v)
	msg_raw(log_Servicing, w, "", colour.Yellow)
}

func Menu(w string) {
	//msg_info(w, v)
	msg_raw(log_Menu, w, "", colour.Cyan)

}

func Template(w string) {
	//msg_info(w, v)
	msg_raw(log_Template, w, "", colour.Cyan)

}

func Accessing(w string) {
	//msg_info(w, v)
	msg_raw(log_Accessing, w, "", colour.Green)

}

func Post(w string) {
	//msg_info(w, v)
	msg_raw(log_Post, w, "", colour.Green)

}

func Query(w string) {
	//msg_info(w, v)
	if CFG.VerboseMode {
		msg_raw(log_Query, w, Character_Query, colour.White)
	}
}

func Storing(t string, w string) {
	//msg_info(w, v)
	if CFG.VerboseMode {

		msg_raw(log_Storing, t+" "+w, Character_Storing, colour.Yellow)
	}
}

func Callout(t string, w string, a string, id string) {
	//msg_info(w, v)
	if CFG.VerboseMode {
		msg_raw(log_Callout, t+" "+ColorGreen+w+ColorBoldYellow+" ["+ColorGreen+a+ColorBoldYellow+"] "+ColorWhite+id+ColorBoldYellow, Character_Callout, ColorBoldYellow)
	}
}

func Saving(t string, w string) {
	//msg_info(w, v)

	msg_raw(log_Save, t+" "+w, Character_Storing, colour.Yellow)

}

func Result(w string, r string) {
	//msg_info(w, v)
	if CFG.VerboseMode {
		msg_raw(log_Result, w, Character_Query+" = "+colour.Bold+r, colour.White)
	}
}

func Database(w string, v string) {
	//msg_info(w, v)
	if CFG.VerboseMode {
		msg_raw(log_Database, w, v, colour.Green)
	}

}

func StartJob(w string) {
	//msg_info(w, v)
	msg_raw(log_ScheduleStart, w, Character_Bike, colour.Green)

}

func EndJob(w string) {
	//msg_info(w, v)
	msg_raw(log_ScheduleEnd, w, Character_Tick, colour.Green)

}

func Warning(s string) {
	msg_raw(log_Warning, s, Character_Warning, colour.Bold+colour.Yellow)

	//log.Println(ColorYellow + "Warning       : " + s + " " + ColorReset)
}

func Message(w string, v string) {
	//log.Println(ColorReset + "Warning       : " + s + " " + ColorReset)
	//output := fmt.Sprintf("%s : %s", w, v)
	//log.Println(ColorReset + output + ColorReset)
	msg_raw(w, v, "", colour.White)

}

func Error(s string, e error) {
	msg_raw(log_Error, s, e.Error(), colour.Bold+colour.Red)
	log.Fatalln(colour.Bold + colour.Red + e.Error())
}

func Fatal(s string, e error) {

	Error(s, e)
}

func Panic(s string, e error) {
	msg_raw(log_Panic, s, e.Error(), colour.Red)
	log.Panicln(colour.Red + s + e.Error())
}

func Publish(w string, v string) {
	op := v + " " + Character_MapTo + "  " + w + " " + Character_Tick
	//	msg_mux(v + " " + Character_MapTo + "  " + w)
	msg_raw(log_MUX, op, "", colour.Bold+colour.White)
}

func msg_raw(pref string, what string, value string, clr string) {

	op := fmt.Sprintf("%-13s : %s %s", pref, what, value)
	log.Println(clr + op + colour.Reset)
}

func Break() {

	width, _, _ := term.GetSize(0)
	if isRunningInDockerContainer() {
		width = 80
	}
	if width <= 0 {
		width = 80
	}

	log.Println(colour.Bold + strings.Repeat("-", width-20) + colour.Reset)
	//log.Println("width: ", width)
}

func Header(s string) {
	log.Println(colour.Green + color.Bold + s + colour.Reset)
}

func Clear() {
	fmt.Print("\033[H\033[2J")
}

func getConfig() (config Config, err error) {
	// get current os directory path
	pwd, _ := os.Getwd()

	//fmt.Println(pwd)
	viper.AddConfigPath(pwd)
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	//spew.Dump(viper.AllKeys())
	//fmt.Printf(viper.AllSettings().(string))

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = viper.Unmarshal(&config)
	//spew.Dump(config)
	if config.Verbose == "true" {
		config.VerboseMode = true
	} else {
		config.VerboseMode = false
	}
	return
}

func isRunningInDockerContainer() bool {
	// docker creates a .dockerenv file at the root
	// of the directory tree inside the container.
	// if this file exists then the viewer is running
	// from inside a container so return true

	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}

	return false
}
