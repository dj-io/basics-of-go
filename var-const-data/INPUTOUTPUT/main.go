package main

// import multiple packages using parentheses block
import (
	"fmt"
	"stratumlabs.ai/go/io/data" // because this is my own project, I need to prefix of my project id (stratumlabs.ai/go/io)
)

// this is a package level variable, it is available to all functions in the package
var name string = "Stratum Labs"

// this init function will execute when the application starts
func init() {
	fmt.Println("init function called")
}

// this init function will also execute when the application starts after the first init function
func init() {
	fmt.Println("init function called 2")
}

func main() {
	fmt.Println(data.MaxSpeed) // using the last part of the import path we can access the code in the data package

	variableDeclarations() // this function is defined in the variable_declaration_examples.go file
	printData() // this function is defined in the functions.go file
}
