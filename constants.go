///NAming convention
//Typed constant
//untyped constant
package main

import (
	"fmt"
)
//can be shadowed but inner declaration wins always
/**const a int16  = 27

func main()  {
	//cant set constant to something that has to be set at runtime
	//math.Sin(1.57) (value of type float64) is not constan
	//const myConst float64 = math.Sin(1.57)
	// a constant can be any of the other discussed data type
	//we can add constant to variable as long as they have same type
	// a constant  like const a = 42 means it is an int. 
	//we can do implicit coversions
	//iota means o i.e const a = iota
	// iota makes you able to create related constants

	const myConst int = 300
	fmt.Printf("%v, %T",myConst,myConst)
}**/

//storing vetenarian in vetenarian clinic
//enumerated constant
//iota always start with o
/**const (
	_ = iota
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func main() {
	//var specialistType int = catSpecialist
	var specialistType int //this is evaluated as 0
	fmt.Printf("%v\n",specialistType == catSpecialist)
}**/

//bit shifting allows us to increase values by powers of 2

const (
	_ = iota
	KB = 1 << (10 * iota) //bit shifting translates to 1024^1
	MB //translates to 1024^2
	GB // translates to 10246^3
	TB // translates to 1024^4
	PB //translates to 1024^5
)

func main()  {
	fmt.Printf("%v KB\n",KB)
	var fileSize float32 =  1024^2
	fmt.Printf("%.4fKB",fileSize/KB)
}


/**const (
	isAdmin = 1 << iota //value is 1 = 00000001
	isHeadquarters //const isHeadquarters untyped int = 2 = 00000010
	canSeeFinancials//const canSeeFinancials untyped int = 4 = 00000100

	canSeeAfrica// 00001000
	canSeeAsia// 00010000
	canSeeEurope //const canSeeEurope untyped int = 32 = 00100000
)



func main()  {
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n",roles) //translates to 00100101
} **/