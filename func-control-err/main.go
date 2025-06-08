package main

import (
	"fmt"
)

// ------------------------------------------------------------
// FUNCTIONS: RETURNING MULTIPLE VALUES
// ------------------------------------------------------------

// declaring a function that takes in a parameter and returns two values
func calculateTax(price float32) (float32, float32) {

	// returning the values
	// the values are seperated by a comma, they will be returned in the order they are declared
	return price*0.09, price*0.02

}

// return types can be named - the names are internal to the function (this is less common)
// adding names does not change how the function is consumed, simply changing how we define the returned values
func calculateTaxNamed(price float32) (stateTax float32, cityTax float32) {

	stateTax = price * 0.09
	cityTax = price * 0.02

	return
}

func printCalculations() {
	// declaring two variables to store the returned values
	// stateTax, cityTax := calculateTax(100)
	// fmt.Println(stateTax, cityTax)

	// what if we only want to store one of the returned values?
	// we can use the blank identifier `_` to store the value we don't want to use

	_, cityTax := calculateTax(100) // here the stateTax is not used
	fmt.Println(cityTax)

}

// -----------------------------------------------------------
// FUNCTIONS: POINTERS & REFERENCES
// ----------------------------------------------------------

// ------------------------------------------------------------
// declaring a function that takes in a pointer to an int
// ------------------------------------------------------------

// prefixing the type with the star * makes the parameter a pointer
// now the param is a pointer to the age object we will pass in
// age *int reads as "here is age, a pointer to an int" now we know that age is a pointer to an int somewhere
func birthday(age *int) {
	*age++
	// by prefixing the value with the * operator we are dereferencing the pointer,
	// this means "perform some action on the value stored at the address this value points to."
	// in this case the operation we want to perform is incrementing the value
}

// ------------------------------------------------------------
// declaring a function that does not take in a pointer
// ------------------------------------------------------------

// Since the param passed in is not a pointer, this param creates a new address in memory for the value (a new variable) instead of using the value from the original address (original variable)
func yearsOfExperience(years int) {
	years++ // this will increment a COPY of the years passed in not the original years passed in
}

func actions() {
	// -------------------------
	// perform actions using pointers
	// -------------------------

	age := 25
	birthday(&age) // prefixing the argument with the & operator passes the ADDRESS of the age variable to the function so that it can reference and modify the original value
	fmt.Println(age) // 26 will be printed because the birthday func accepts a pointer and we have passed the address of the age variable in the function call denoted by the & operator

	// -------------------------
	// perform actions without using pointers
	// -------------------------
	yoe := 10
	yearsOfExperience(yoe)

	// 10 will still be printed because yearsOfExperience passed a copy of the value
	// (i.e the param passed into yearsOfExperience is not a pointer/reference to the yoe variable here but a copy of the value)
	fmt.Println(yoe)

}

// ------------------------------------------------------------
// FUNCTIONS: PANIC & DEFER
// ------------------------------------------------------------

func yearlySalary(salary *int) {
	if *salary > 140 {
		// panic is used to abort some function, panic is similar to throwing an error if some condition
		panic("Salary is to high")
	}
	*salary++
}
func panic_defer() {
	salary := 100000
	yearlySalary(&salary)
	fmt.Println(salary)
}

func main() {
	// defer is used to push the function call to the end of the function call regardless of its place in the code
	// You can have more than one defer
	// Defer is commonly used for cleanup operations like:
	// - Closing files after reading/writing
	// - Releasing database connections
	// - Unlocking mutexes
	// - Cleaning up resources
	// This ensures cleanup happens even if errors occur, since deferred
	// functions run during panic unwinding
	defer panic_defer()
	printCalculations()
	actions()
}

