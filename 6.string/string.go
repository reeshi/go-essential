package main

import "fmt"

func main(){
	// strings are immutable in go
	str := "I am rishikesh Yadav"

	fmt.Println(str)

	// length of string
	fmt.Println("length of str :",len(str))

	// get a char at specific index
	fmt.Println(string(str[0]))

	// substring from string , end is not inclusive
	fmt.Println(str[5:14])
	// it will get substring from index 5 to till last 
	fmt.Println(str[5:])
	// it will get substring from start to till index 3
	fmt.Println(str[:4])
}