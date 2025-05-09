package main

import "fmt"

func printData() {
	// safe print - using the fmt package
	fmt.Print("Hello ")
	fmt.Println("Safe Print")
	fmt.Println(name)

	// unsafe print - using the print function
	print("Hello ")
	println("Unsafe Print")
	print(name) // we can access the package level variable from the main.go file because main.go and functions.go are apart of the same package
}
