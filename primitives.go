//57.07 PRIMITIVES
//Boolean data has true and false

package main

import (
	"fmt"
)

func main()  {
	//var n bool = true
	//Everytime a variable is initialized, the value becomes 0
	
	//var n bool = true Everytime a variable is initialized, 
	//the value becomes 0 hence when a value is initialized and the value is not given, 
	//the value becomes 0 hence false
	/*x := 1 == 1
	y := 1 == 2
	z := "yes" == "yes" */
	//var n bool
	//fmt.Printf("%v",n)
	//fmt.Printf("%v %v %v",x,y,z)

	//fmt.Printf("%v ,%T",n,n)
	//a := 10
	//b := 3

	//var a int = 10
	//var b int8 = 3

	//mismatched types int and int8)
	//fmt.Println(a+b)
	//fmt.Println(a+int(b))


	/**fmt.Println(a+b)
	fmt.Println(a/b)
	fmt.Println(a*b)
	fmt.Println(a/b)
	fmt.Println(a%b)
	fmt.Println(a-b) **/

	//bit shifting
	a := 8 //2^3
	fmt.Println(a << 3) //2^(3+3) = 2^6 = 64
	fmt.Println(a >> 3) //2^(3-3) = 2^0 = 1

}