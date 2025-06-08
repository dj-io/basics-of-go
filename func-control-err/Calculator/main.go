package main

import "fmt"


func main() {
 var operation string
 var number1, number2 int // we can use one int to declare multiple variables if they are of the same type being declared on the same line

 fmt.Println("CALCULATOR GO 1.0")
 fmt.Println("--------------------------------")
 fmt.Println("Which operation do you want to perform? add, subtract, multiply, divide")
// scanf takes in input from the user, using scanf we can format the input, scanf takes in two arguments, the format and the address of the variable (%s for string, %i for int) & (&varName)
 fmt.Scanf("%s", &operation) // we use the & operator to pass the address of the variable operation to the scanf function, without this operation will be a copy of the input and not the actual variable
 fmt.Println("Enter the first number:")
 fmt.Scanf("%d", &number1)
 fmt.Println("Enter the second number:")
 fmt.Scanf("%d", &number2)


 switch operation {
	case "add":
		fmt.Println(number1 + number2)
	case "subtract":
		fmt.Println(number1 - number2)
	case "multiply":
		fmt.Println(number1 * number2)
	case "divide":
		fmt.Println(number1 / number2)
 }


}
