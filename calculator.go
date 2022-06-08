// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var operation string
	fmt.Println("Enter 1st Number")
	fmt.Scanln(&num1)
	fmt.Println("Enter 2nd Number")
	fmt.Scanln(&num2)
	fmt.Println("operation")
	fmt.Scanln(&operation)
	var result int

	if operation == "+" {
		result = num1 + num2
	} else if operation == "-" {
		result = num1 - num2
	} else if operation == "/" {
		if num2 == 0 {
			fmt.Println("Number is not Divided by Zero")
			return
		}
		result = num1 / num2
	} else if operation == "*" {
		result = num1 * num2
	} else {
		fmt.Println("operation in not supported please enter +,-, * or / operation")
		return
	}
	fmt.Println("The result is: ", result)
}
