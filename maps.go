package main

import "fmt"

func main()  {
	//A map making
	statePopulations := make(map[string]int)
	statePopulations = map[string] int{
		"California" : 123456,
		"Texas" : 12212,
		"Kenya" : 121212,
	}
	//pull a single value
	//statePopulations["Tanzania"] = 111111
	//fmt.Printf("%v",statePopulations["Texas"])
	//fmt.Print(statePopulations)
	//delete entries
	delete(statePopulations,"Texas")
	fmt.Printf("%v",statePopulations)
	_,ok:= statePopulations["Texas"]

	//ok is false if value is not found
	fmt.Println(ok)
	//find elements
	fmt.Printf("%v",len(statePopulations))
	//iterate through the maps
	for key,value := range statePopulations{
		fmt.Println("Key is : ",key)
		fmt.Println("Value is : ",value)
	}

}