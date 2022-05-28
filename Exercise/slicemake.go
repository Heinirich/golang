package main

import "fmt"

func main()  {
	//slice
	slice := make([]float64, 5)

	//array
	var array [5] int

	//maps
	//var maps map[string]int

	maps  := make(map[string] int, 0)
	maps["key"] = 120

	array[1] = 20.0

	fmt.Println(slice)
	fmt.Println(array)
	fmt.Println(maps)
}