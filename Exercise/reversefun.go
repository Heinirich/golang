package main

import "fmt"

func main()  {
	text := "liVERPOOL"
	
	reverse_string := NumReverse(text)
	
	fmt.Println(reverse_string)
}

func NumReverse(text string) string{
	text_rune := make([] rune,len(text))
	n := 0
	for _,v := range text{
		text_rune[n] = v
		n++ 
	}

	for i := 0; i < n/2; i++ {
		
		text_rune[i],text_rune[n-1-i] = text_rune[n-1-i],text_rune[i]
	}
	return string(text_rune);
}