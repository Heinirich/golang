package main

import (
	"fmt"
)

func main()  {
	num := 6

	fmt.Println(Factorial(num))
}


func Factorial(i int) int{
	
	

	if i <= 1 {
		return 1
	}

	fmt.Printf("%v *",i)
	return i*Factorial(i-1)
	
}