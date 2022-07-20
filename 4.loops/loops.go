package main

import "fmt"

func main(){
	// there are only one loop in go i.e for loop
	for i:= 1; i<3; i++{
		fmt.Println(i)
	}

	fmt.Println("---------------")

	// use break and continue keyword
	for i:=1; i<5; i++{
		if i > 3{
			break;
		}
		fmt.Println(i)
	}

	fmt.Println("---------------")
	for i:=1; i<5; i++{
		if i < 3{
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("---------------")
	// we can use this for loop as while loop also
	i := 1
	for i <= 3{
		fmt.Println(i)
		i++
	}
}