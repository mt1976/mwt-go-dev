package dao

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	core "github.com/mt1976/mwt-go-dev/core"
)

type SQLData struct {
	fields []string
	values []string
}

func addData(d SQLData, f string, v string) SQLData {
	if f != "_id" {
		d.fields = append(d.fields, f)
		d.values = append(d.values, sq(v))
	}
	return d
}

func fields(ts SQLData) string {
	return strings.Join(ts.fields[:], ",")
}

func values(ts SQLData) string {
	return strings.Join(ts.values[:], ",")
}

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

func get_TableName(in string, in2 string) string {
	return in + "." + in2
}

//Gets a string from the interface map & default a value if not found
func get_String(t map[string]interface{}, v string, d string) string {
	i := t[v]
	fmt.Printf("%q: %q %q\n", v, i, d)
	if i == nil {
		return d
	}
	return i.(string)
}

//Gets a time.Time from the interface map & default a value if not found
func get_Time(t map[string]interface{}, v string, d string) string {
	i := t[v]

	fmt.Printf("%q: %q %q\n", v, i, d)
	if i == nil {
		return d
	}
	return core.TimeToString(i.(time.Time))
}

//Convert time.Time to string
func get_Int(t map[string]interface{}, v string, d string) string {
	i := t[v]
	fmt.Printf("%q: %q %q\n", v, i, d)

	if i != nil {
		return fmt.Sprintf("%d", i.(int64))
	}
	return d
}

//Convert float64 to string
func get_Float(t map[string]interface{}, v string, d string) string {
	i := t[v]
	fmt.Printf("%q: %q %q\n", v, i, d)
	if i != nil {
		return fmt.Sprintf("%f", i.(float64))
	}
	return d
}
