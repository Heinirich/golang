package main

import "fmt"

func main(){

	/**i := 0
	//for ;i<5;i++{
		for ;i<5;i++{
	//for i :=0; i<5; i=i+2{
		fmt.Println(i)
		//i++
		//continue statements are important in showing when or when not to do an action
	}

	//A Loop:  label break Loop breaks in all loops
	**/

	s := [] int{1,2,3}

	for k,v := range s{
		fmt.Println(k,v)
		//fmt.Println(v)
	}
}