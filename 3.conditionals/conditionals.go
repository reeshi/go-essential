package main

import "fmt"

func main(){
	x := 10
	// if-else 
	if x > 5 {
		fmt.Println("x is big")
	}else{
		fmt.Println("x is small")
	}

	// using logical AND in if condition
	if x > 5 && x <= 10 {
		fmt.Println("X is in range")
	}else {
		fmt.Println("x is out of range")
	}

	// using logical OR in if condition
	if x == 5 || x == 10{
		fmt.Println("value of x is either 5 or 10")
	}

	// we can also do this
	a := 4
	b := 4
	if frac := a/b; float64(frac) > 0.5 {
		fmt.Println("frac value greater than half")
	}

	// switch cases
	n := 2
	switch n {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("many...")
	}

	// naked switch
	switch {
	case n == 1:
		fmt.Println("value of n is one")
	case n == 2:
		fmt.Println("value of n is two")
	case n == 3:
		fmt.Println("value of n is three")
	default:
		fmt.Println("value of n is many")
	}

}