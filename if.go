//2:48
//if and switch statement
package main

import (
	"fmt"
	"math"
)

func main()  {
	/**if true{
		fmt.Printf("Hello World")
	}
	
	//interrogating the mao

	statePopulations := map[string] int{
		"California" : 123456,
		"Texas" : 12212,
		"Kenya" : 121212,
	}

	if pop,ok := statePopulations["California"];ok{
		fmt.Println(pop)
	}
	//Not operate returns the opposite of the other result
	//Short curcuiting means or does not evaluate the rest if the others are true
	//if ....{} else if "statetement" {} **/
	myNum := 0.1
	if myNum == math.Pow(math.Sqrt(myNum),2){
		fmt.Println("They are the same")
	}else{
		fmt.Println("They are different")
	}


}