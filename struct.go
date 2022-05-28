package main

import (
	"fmt"
	//import reflec 
	"reflect"
)
/**
//naming convention must start with capital
type Doctor struct{
	//small letters for field names
	//no underscore in field na,es
	number int
	name string
	companions [] string
	p_number int
} 

func main()  {
	aDoctor := Doctor{
		number: 3,
		name: "Johnny Depp",
		companions: [] string{
			"Amber Heard",
			"Jennifer Lopez",
		},
		p_number : 0746273465,
	}
	
	fmt.Println(aDoctor);
	//accessing is easier using dot syntax
	fmt.Println(aDoctor.p_number)
	fmt.Println(aDoctor.companions[1])

	//anonymous struct
	bDoctor := struct{name string}{name : "Johnny Depp"}
	anotherDoctor := bDoctor
	anotherDoctor.name = "Smith Heinrich"

	fmt.Println(bDoctor.name)
	fmt.Println(anotherDoctor.name)
	
	
}
**/

type Animal struct {
	Name string `max:"2"`
	origin string
}

type Bird struct{
	//this is called composition relation which means bird has animal characteristics
	Animal
	speedKPH float32
	color string
	canFly bool
}
func main()  {
	//b := Bird{}
	//b.name = "Spartan"

	b := Bird {
		Animal: Animal{Name: "Kiwi",origin:"Welcome"},
		speedKPH: 300,
	}
	fmt.Println(b.Name)
	
	//get from validation
	t := reflect.TypeOf(Animal{})
	field,_ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}

