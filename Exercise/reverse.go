package main

import (
	"fmt"
)

//reverse a number
func main() {
	name := "Welcome"
	name_length := len(name)

	rune := make([]rune, name_length)
	//when created all values become 0
	//Strings are made of bytes and they can contain valid characters that can be represented using runes.
	// rune() function to convert string to an array of runes.
	n := 0
	for _, r := range name {
		rune[n] = r
		n++
		//fmt.Println(n)
	}
	//fmt.Println(n)
	//rune = rune[0:n]
	fmt.Println(rune)
	// Reverse
	for i := 0; i < n/2; i++ {
		fmt.Println(i) //prints 0,1,2
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
		//shifting
		fmt.Println(rune[i], rune[n-1-i])
	}

	//convert to utf-8
	output := string(rune)
	fmt.Println(output)

	//type byte = uint8
	//type rune = int32

}
