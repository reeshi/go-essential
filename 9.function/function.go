package main

import "fmt"

func main(){
	// fmt.Println("ready!!")

	// calling a add function
	val := add(5,3);
	fmt.Println(val);

	// calling quotientRemainder function
	quotient, remainder := quotientRemainder(10,2)
	fmt.Println("quotient :",quotient,"and reaminder :",remainder)

	arr := []int{1,2,3,4}
	// doubles the value at index 2
	double(arr,2)
	fmt.Println(arr)

	// calling a doubleN function to double the value of n . but it will not work because go is pass by value.
	n := 2
	doubleN(n);
	fmt.Println("normal doubleN",n);

	// if we want refelect a change in variable n then we have to pass it by pointer
	doublePtr(&n);
	fmt.Println("pointer doubleN",n);

}

/******** Create a function in go 

func functionName(parameter1Name type,  parameter2Name type) returnType {
	// function body
}

// in go a function can return mulptile things at same time
func functionName(parameter1Name type,  parameter2Name type) (returnType1,returnType2,) {
	// function body
}

**/


//  go is pass by value(default)
func double(arr []int, idx int){
	arr[idx] *= 2
}

func doubleN(n int){
	n = n * 2
}

func doublePtr(n *int){
	*n *= 2;
}

// add two numbers 
func add(a int, b int) int{
	ans := a + b;
	return ans;
}

// function which give quotient and remainder of two number
func quotientRemainder(a int, b int) (int, int){
	quotient := a / b
	remainder := a % b
	return quotient,remainder
}

