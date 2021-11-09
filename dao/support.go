package dao

import (
	"log"
	"strconv"
)

func sienaYN(inValue string) string {
	log.Println("inValue", inValue)
	var outValue string
	outValue = ""
	if inValue == "true" {
		outValue = "Yes"
	}
	if inValue == "false" {
		outValue = "No"
	}
	return outValue
}

func sienaDBBoolYN(inBool bool) string {

	inValue := strconv.FormatBool(inBool)

	var outValue string
	outValue = ""
	if inValue == "true" {
		outValue = "Yes"
	}
	if inValue == "false" {
		outValue = "No"
	}
	return outValue
}

func sq(in string) string {
	return "'" + in + "'"
}
