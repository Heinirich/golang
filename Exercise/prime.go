//1) Write a program in Go to check whether a number is prime or not?


package main

import (
	"fmt"
)
func main()  {
	
	num := 9

	fmt.Println(primeCheck(num))


}

func primeCheck(i int) bool{

	if i%2==0 || i ==2 {
		return false
	}
	
	return true
}
