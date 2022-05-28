package main

import "fmt"

func main()  {

	fmt.Print("Enter a number: ")
	var input string

	fmt.Scanf("%f", &input)

	name := input

	//name := "Hello"

	name_rune := make([] rune, len(name))

	n := 0
	for _,v := range name{
		name_rune[n] = v
		n++
	}

	for i := 0; i < n/2; i++ {
		name_rune[i],name_rune[n-1-i] = name_rune[n-1-i],name_rune[i]
	}

	fmt.Println(string(name_rune))
}