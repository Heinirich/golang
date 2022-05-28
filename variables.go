package main

import "fmt"
import "strconv"

//var i int = 42
//var actorName string = "Johhny Depp"

//var (
	//actorName string = "Johnny Depp"
	//season int = 11
//)
var i int = 42;

func main(){
	//var i int
	//i = 42
	//i = 55
	//var i float32 = 42
	//i := 42; collon syntax

	//fmt.Println(i)
	//fmt.Printf("%i",i)
	//T = Type
	//v = Value
	//Cant declare variable again in same scope
	//varibles must always be used
	//Naming controls visibility
	//Make acronyms uppercase i. e var theHTTP not var thehttp string
	//j = int(i) or float32(1)
	//a string is stream of bytes
	//strconv converts any other type 
	//lower case is scoped topackage
	//uppercase is globaly scoped
	//a rune in int32
	var x int = 11
	fmt.Print(isEven(x))
	fmt.Println(i)
	var j string;
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T",j,j)
}



func isEven(i int) string{
	if i%2 == 0{
		return "Even"
	}else{
		return "Odd"
	}
	
}


