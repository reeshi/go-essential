package main

import (
	"fmt"
)

func main(){
	// decalre a array(fixed size) in go
	animals := []string{"dog", "cat", "lion", "tiger"}
	// animals[4] = "leopard" 
	animals = append(animals,"leopard")
	// length 
	fmt.Println("length =", len(animals))

	// print the element at specific index
	fmt.Println(animals[3]) // tiger

	fmt.Println(animals)

	// decalre a slice (basically a dyanamic array)
	var rollNo []int
	for i := 0; i<10; i++{
		// to add element at last in slice
		rollNo = append(rollNo, i)
	}
	fmt.Println("rollNo :", rollNo)

	// find the maximum value of this slice
	nums := []int{100,2,5,89,102}

	// insert 90 at index 3 in slice nums
	nums = append(nums, 0)
	copy(nums[4:], nums[3:])
	nums[3] = 90
	fmt.Println("nums : ", nums);

	// max logic
	max := nums[0]
	for _, val := range nums {
		if val > max{
			max = val
		}
	}

	fmt.Println("Maximum of nums :", max)
}