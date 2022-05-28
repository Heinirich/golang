package main

import (
	"fmt"
	"strings"
	"reflect"
	"strconv"
)



type Animal struct {
	Name   string `max:"2"`
	origin string
}

type Bird struct {
	//this is called composition relation which means bird has animal characteristics
	Animal
	speedKPH float32
	color    string
	canFly   bool
}

func main() {
	b := Bird{
		Animal:   Animal{Name: "", origin: "Welcome"},
		speedKPH: 300,
	}
	fmt.Println(b.Name)

	t := reflect.TypeOf(Animal{})
	field,_ := t.FieldByName("Name")
	fmt.Println(field.Tag)
	tag  := string(field.Tag)
    fmt.Println(tag)

	Validate
	validate := make(map[string]string)
	validate = map[string]string{
		//"Name": IsMax(3, len(b.Name),tag),
		"Name": IsRequired(b.Name),
	}

	fmt.Println(validate)
fmt.Println(strconv.Itoa(3))
	

}
func Validate(u_data map) string {
 return "error detected"
}

func IsRequired(value string) string{
	if (len(value) == 0) {
       return "Name Cannot Be Null"
	}
	return "Continue"
}

func IsMax(max int, value int,tag string) string {
	//get from validation
	
	//fmt.Println(field.Tag)
	textRules := struct{ rules string }{rules: string(tag)}
	law := strings.Split(textRules.rules, ":")
	layout_rule := law[1]
	var strLen string= RemoveQuotes(layout_rule)

	return strLen
	//fmt.Println(x)
}
func RemoveQuotes(i string) string {
	quoted_data := i[1 : len(i)-1]
	res := quoted_data
	return res
}
