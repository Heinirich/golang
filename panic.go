//panic js

package main
import (
	"fmt"
)

func main()  {
	a, b := 1,0
	ans := a/b
	panic("An error Occurred")
	fmt.Println(ans)
	
}