package main


// this variable is global-scoped (declared outside of a function, available to the package)
var url = "https://stratumlabs.ai"

func main() {

	// --------
	// ZERO VALUE DECLARATIONS (least common)
	// --------

	// We can declare variables without assigning a value to them, and they will be assigned a zero value (empty string, 0, false)
	// Declaring variables this way means we set the value later
	var message string
	var x int
	var y float64
	var z bool

	// setting the values
	message = "Hello from go"
	x = 10
	y = 34.4
	z = true

	// --------
	//  DECLARATIONS WITH ASSIGNMENT (TYPE INFERENCE)
	// --------

	// Even though go is a statically typed language, we can omit the type if we are assigning a value, and go will infer the type
	// You typically want to set the type if you want more precision or if you want to save memory i.e, an int8 is smaller than an int32

	var message2 = "Hello from go" // string type is inferred
	var price = 34.4 // float type is inferred

	// --------
	// SHORTHAND DECLARATIONS (most common)
	// --------

	// we dont have to use var to declare variables we want to assign a value to, we can use the := operator
	// we can only use the := operator within a function, not at the package level
	walrusVarMessage := "Hello from go"
	walrusVarPrice := 34.4

	// --------
	// CONSTANTS
	// --------

	// for constants we use the const keyword, and we cannot use the := operator
	const pi = 3.14 // we can omit the type if we want, float type is inferred
	const e float64 = 2.71828 // we can specify the type if we want



	// ------------------------------------------------------------
	// OUTPUT VARIABLES
	// ------------------------------------------------------------

	// we can pass in unused variables to print to quiet the linter
	print(message, price, message2, x, y, z, walrusVarMessage, walrusVarPrice, pi, e, url) // print is a function that takes a variable number of arguments
}
