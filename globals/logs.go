package globals

import (
	"fmt"
	"log"
)

func LOG_header(s string) {
	msg_header(s)
}

func LOG_success(s string) {
	msg_done(s)
}

func LOG_info(w string, v string) {
	msg_info(w, v)
}

func LOG_warning(s string) {
	log.Println(ColorYellow + "Warning       : " + s + " " + ColorReset)

}

func LOG_msg(w string, v string) {
	//log.Println(ColorReset + "Warning       : " + s + " " + ColorReset)
	output := fmt.Sprintf("%s : %s", rightPad2Len(w, " ", 14), v)
	log.Println(ColorReset + output + ColorReset)
}

func LOG_error(s string, e error) {
	log.Fatalln(ColorRed + "ERROR         : " + s + " " + e.Error() + ColorReset)
}

func LOG_it(w string) {
	LOG_msg("Information", w)
}

func msg_header(s string) {
	log.Println(ColorWhite + "Information    : " + s + " " + ColorReset)
	//log.Println(strings.Repeat("-", len(s)))
}
func msg_done(s string) {
	log.Println(ColorYellow + "Success        : " + s + " " + ColorReset + Tick)
}
func msg_info(what string, value string) {
	output := fmt.Sprintf("Information    : %s %s", rightPad2Len(what, " ", 25), value)
	log.Println(ColorCyan + output + ColorReset)
}
