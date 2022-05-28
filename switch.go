package main

import "fmt"
//test cases have to be unique
func main()  {

	/**we can use initializzation switch i:= 2+3;i
	switch 8 {
	case 1,5,10:
		fmt.Println("one five ten")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("Not found")
	} **/

	/*i := 10
	switch  {
	case i <90:
		fmt.Println("This is 2")
		//fallthrough says go to next case though rarely used
		fallthrough
	case i >= 10:
		fmt.Println("Greater than 10")
		fallthrough
	default:
		fmt.Println("Is default")
	}*/

	var i interface{} = 1.0
	switch i.(type) {
	case int:
		fmt.Println("This is int")
		//break can switch an from code quickly
		
	case float64:
		fmt.Println("This is float")
	default:
		fmt.Println("Undefined type")
	}
}