//1:48:11
//all about arrays here

package main

import "fmt"

func main()  {
	//creating an array...this method however add array thrice
	//grades := [3]int{97,96,99} //number, type and array
	//grades := [...]int{97,96,99} 
	/**var students [4] string
	fmt.Printf("Students %v\n",students)
	students[0] = "Lisa"
	students[1] = "1"

	fmt.Printf("%v\n",students)
	//dereference
	fmt.Printf("Student 1: %v\n",students[0])
	//how big array is
	fmt.Printf("Number of students: %v\n",len(students))
	//make array of arrays **/

	/**var identityMatrix [3][3]int

	identityMatrix[0] = [3] int{0,1,2}
	identityMatrix[1] = [3] int{1,2,3}

	fmt.Println(identityMatrix)**/
	//arrays are considered as values

	//a := [...]int{1,2,3}
	//use pointers easily
	//b := a
	//this changes values for both of them. There useful is in slice much in go
	//b := &a
	//b[1] = 5

	//fmt.Printf("%v",a)
	//fmt.Printf("%v",b)
	//slice example
	/**a := [] int {1,2,3}
	fmt.Printf("%v\n",a)
	fmt.Print(cap(a)) //Capacity in slice
	//a change in slice may change the whole values similar to slices **/


	/**a := []int{1,2,3,4,5,6,7,8,9}
	//adding ...adds it to arrays
	a := [...]int{1,2,3,4,5,6,7,8,9}
	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:6]
	//this changes everything
	a[5] = 42
	fmt.Println(a)//all elements
	fmt.Println(b)//all elements
	fmt.Println(c)//all elements from slice with index 3 to end
	fmt.Println(d)//all elements from slice with element 1 to 6
	fmt.Println(e)//slices elemt 3 to element 6
	//first number is exclusive last ins inclusive
	**/
	//use make function to create a slide
	//The make built-in function allocates and initializes an object of type slice, map, or chan (only).
	// Like new, the first argument is a type, not a value. 
	//Unlike new, make's return type is the same as the type of its argument, not a pointer to it. 
	//The specification of the result depends on the type:
	/**a := make([]int,3,100)
	fmt.Println(a)
	fmt.Printf("Length of a %v\n",len(a))
	fmt.Printf("Capacity of a %v\n",cap(a)) */
	//The append built-in function appends elements to the end of a slice. If
	//it has sufficient capacity, the destination is resliced to accommodate the
	//new elements. If it does not, a new underlying array will be allocated.
	/**a := []int {}
	fmt.Println(a)
	a = append(a, 120)
	fmt.Printf("%v\n",a)
	a = append(a, 10,12,11)
	fmt.Println(a)
	//use ... if you know underlaying array is big
	//this is spread operator
	a = append(a, []int{10,12,11}...) **/

	a := []int{1,2,3,4,5,6,7,8,9}
	//remove first elemment
	//b :=  a[1:]
	//remove last element
	//c := a[:len(a)-1]

	//fmt.Printf(" %v , %v",b,c)
	

}