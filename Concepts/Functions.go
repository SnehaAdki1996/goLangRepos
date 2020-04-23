package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Main Function")
	var a int = 10
	var b int = 20
	//multiple passing Parameter
	rtnValue := dummyFunctionAdd(a, b)
	fmt.Println("Function Returned is : ", rtnValue)
	port, errorValue := server(10)
	fmt.Println("Multiple Return Value From function Port Value:", port, " Error Value:", errorValue)

	//write only variable
	_, err_value := server(300)
	fmt.Println("Multiple Return Value From function and using only one Error Value:", err_value)
}

//2nd way of Definig Paramerts
// func dummyFunctionAdd(num1, num2 int) {
// 	fmt.Println("Addition of num1:", num1, " and num2:", num2, num1+num2)
// }

//Reuring Data from function can be bool,error,int....
//error specifies something bad has hapenned
//multiple return values
func dummyFunctionAdd(num1 int, num2 int) error {
	fmt.Println("Addition of num1:", num1, " and num2:", num2, num1+num2)
	return errors.New("Something Went Wrong")
}

func server(port int) (int, error) {
	return port, nil
}

// OUTPUT:
// C:\Go_Code\Concepts>go run Functions.go
// Main Function
// Addition of num1: 10  and num2: 20 30
// Function Returned is :  30
//Function Returned is :  <nil> return nil
//Function Returned is :  Something Went Wrong
//Multiple Return Value From function Port Value: 10  Error Value: <nil>
//Multiple Return Value From function and using only one Error Value: <nil>
//NOTE Declare Local Variable to be used must
