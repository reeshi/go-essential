package main

import "fmt"

func main(){
	// declare a int variable in go
	/*
		var x int
		var y int

		x = 2
		y = 2
	*/

	/*
	// declare float value in go
		var x float64
		var y float64
		
		x = 2.0
		y = 2.0
	*/

	// shortcut to declare a variable and assign in one line
	x := 4.0
	y := 4.0

	// if you want to declare multiple variable  of same type and assign a value than you can do this way
	a,b:=2,2 

	fmt.Println(a,b)
	fmt.Println(b)

	fmt.Printf("x=%v type of %T\n",x,x)
	fmt.Printf("y=%v type of %T\n",y,y)

	// calculate the mean of two number
	mean := (x + y) / 2.0

	fmt.Printf("result: %v type of %T\n",mean,mean)

}