package main

import "fmt"

func main()  {
	sum(1,2,3,4,5,6)
	subtract(84,7)

	func() {
		fmt.Println("Hello")
	}()
}

func subtract(value1, value2 int)(sub int) {

	sub = int(value1-value2)

	fmt.Println(sub)

	return
}
func sum(values ...int) int {
	fmt.Println(values)

	sum := 0

	for _,v := range values{
		
		sum = sum+v
	}

	fmt.Println("The sum is ",sum)

	return sum
}