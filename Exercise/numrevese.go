package main

import (
	"fmt"
)

func main()  {

	number := 5645464

	fmt.Println(numReverse((number)))
	
}

func numReverse(num int) int{
	res := 0

	for num >0{
		remainder := num%10
		res = (res*10)+remainder
		num  /= 10
	}
	return res
}