//3:41:36
//invocation of functions

package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)
func main()  {
	//defer is used for delay put off (an action or event) to a later time; postpone.
	//LIFO(Last in first out)
	/**defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three") **/
	res, err := http.Get("https://www.google.com/robots.txt")

	if err != nil {
		log.Fatal(err)
	}
	
	defer res.Body.Close()
	
	robots, err :=ioutil.ReadAll(res.Body)
	

	if err!=nil{
		log.Fatal(err)
	}
	fmt.Printf("%s",robots)
}