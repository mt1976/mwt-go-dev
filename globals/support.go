package globals

func LOG_header(s string){
	msg_header(s)
}

func LOG_success(s string){
	msg_done(s)
}

func LOG_info(s string){
	msg_info(s)
}
func msg_header(s string) {
	log.Println(globals.ColorYellow + "Information   : " + s + " " + globals.ColorReset)
	//log.Println(strings.Repeat("-", len(s)))
}
func msg_done(s string) {
	log.Println(globals.ColorYellow + "Success       : " + s + " " + globals.ColorReset + globals.Tick)
}
func msg_info(what string, value string) {
	output := fmt.Sprintf("Information   : %-25s %s", what, value)
	log.Println(globals.ColorCyan + output + globals.ColorReset)
}