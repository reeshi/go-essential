package main

import (
	"fmt"
	"strings"
)

func main(){
	// map is data structure which store data in key-value pair
	// declare in map
	capitals := map[string]string{
		"maharashtra" : "mumbai",
		"uttar pradesh" : "lucknow",
		"bihar" : "patna",
	}
	
	// length 
	fmt.Println("Length :", len(capitals))

	
	// get value of key
	fmt.Println(capitals["maharashtra"])

	// know is key present or not ?
	value, present := capitals["delhi"]

	if !present{
		fmt.Println("delhi key not present")
	}else{
		fmt.Println("delhi key present and value is :", value)
	}

	// set a new key value
	capitals["punjab"] = "chandigarh"
	fmt.Println(capitals)

	// delete a key
	delete(capitals, "punjab")
	fmt.Println(capitals)

	// iterating over map
	// ------ first method
	for key := range capitals{
		fmt.Println(capitals[key])
	}

	// --------- second method
	for key, value := range capitals{
		fmt.Println(key, "---->", value)
	}

	// ------------------  CHALLANGE QUESTION -------------------------
	// count how many times each word present in a text
	str := "GeeksforGeeks is a computer is a !"
	strArr := strings.Fields(str)

	count := map[string]int{}

	for _,word := range strArr{
		count[word]++
	}

	fmt.Println(count)
	
}