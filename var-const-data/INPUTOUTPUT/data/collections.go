package data

import "fmt"


// ------------------------------------------------------------
// INIT FUNCTION
// ------------------------------------------------------------

// this is an init function
// It is a way to ensure a valid state is set for package-level variables
// It is automatically invoked once per package before the ⁠main function (or when the package is imported).
// init functions are similar to constructor functions in other languages
func init() {
	fmt.Println("init function called")
}


// multiple init functions can be defined in a single file
// init functions are executed in the order they are defined



// ------------------------------------------------------------
// DECLARE COLLECTIONS
// ------------------------------------------------------------

var Countries [10]string // we declare an array of 10 strings with no values
var Slice []int // we declare a slice of integers with no values
var Codes map[string]int // we declare a map of strings to integers with no values

// here we use the init function to set the values of the array
func init() {

	// ------------------------------------------------------------
	// SET THE VALUES OF THE COLLECTIONS
	// ------------------------------------------------------------

	// here we set the values of the array
	Countries = [10]string{
		"United States",
		"Canada",
		"United Kingdom",
		"Australia",
		"New Zealand",
		"India",
		"Brazil",
		"Argentina",
		"Chile",
		"Mexico",
	}

	// we can also set the values of the array index by index
	Countries[0] = "United States"
	Countries[1] = "Canada"
	Countries[2] = "United Kingdom"
	Countries[3] = "Australia"
	Countries[4] = "New Zealand"

	// It does not matter if we set the values in order
	Countries[5] = "India"
	Countries[8] = "New Zealand"

	Slice = []int{1, 2, 3, 4, 5}

	Codes = map[string]int{
		"United States": 1,
		"Canada": 2,
		"United Kingdom": 3,
		"Australia": 4,
		"New Zealand": 5,
	}

	qty := len(Countries)
	fmt.Println("The number of countries in the array is", qty)
}

func init() {
	// ------------------------------------------------------------
	// DECLARE COLLECTIONS AND SET VALUES
	// ------------------------------------------------------------

	// here we declare an ARRAY of strings and set the values
	Cities := [2]string { "New York", "Toronto" } // the curly braces look like code blocks, but they are like constructors
	fmt.Println(len(Cities), Cities)

	// here we declare a SLICE of strings and set the values
	names := []string {"John", "Jane", "Jim", "Jill"}

	// here we use the global function append to add a value to the slice
	names = append(names, "John")
	fmt.Println(len(names))

	// here we declare a MAP of strings to integers with values
	Rooms := map[string]int{
		"101": 1,
		"102": 2,
		"103": 3,
		"104": 4,
		"105": 5,
	}

	fmt.Println("The number of rooms in the map is", len(Rooms))

	// ------------------------------------------------------------

}


