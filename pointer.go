package main

import "fmt"


func main()  {
	var a int = 42
	var b*int = &a

	fmt.Println(a,*b)

	c := [3]int{1,2,3}

	d := &c[0]

	fmt.Println(*d)

	var ms *myStruct

	ms = &myStruct{foo:42}

	//ms = new(myStruct)

	//(*ms).foo =42

	//ms.foo = 42
	
	fmt.Println(ms)

	
}

type myStruct struct{
	foo int
}