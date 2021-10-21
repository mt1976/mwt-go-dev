package globals

import(
		"log"
		"fmt"
)

func LOG_header(s string){
	msg_header(s)
}

func LOG_success(s string){
	msg_done(s)
}

func LOG_info(w string,v string){
	msg_info(w,v)
}

func msg_header(s string) {
	log.Println(ColorYellow + "Information   : " + s + " " + ColorReset)
	//log.Println(strings.Repeat("-", len(s)))
}
func msg_done(s string) {
	log.Println(ColorYellow + "Success       : " + s + " " + ColorReset + Tick)
}
func msg_info(what string, value string) {
	output := fmt.Sprintf("Information   : %-25s %s", what, value)
	log.Println(ColorCyan + output + ColorReset)
}