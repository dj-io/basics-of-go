package main

import (
	"fmt"
	"os"
	"stratumlabs.ai/go/files/fileutils"
)

func main() {
	rootPath, _ := os.Getwd() // use the os package to get the current working directory (here we ignore the error since it is not important)
	filePath := rootPath + "/data/text.txt"
	c, err := fileutils.ReadTextFile(filePath) // when reading a file, we must pass in the absolute path to the filename

	// this is the pattern seen ALOT in go, any function that can possibly error, will be declared and used with this pattern
	if err == nil {
		fmt.Println(c)
		// with complex output, go has a fmt.Sprintf func that allows us to format the output as a string\
		// %v is a placeholder for the value of the variable
		// \n is a new line character

		newContent := fmt.Sprintf("Original: %v\n Double Original: %v%v", c, c, c)
		fileutils.WriteToFile(filePath+".output.txt", newContent)
	} else {
		fmt.Printf("Error reading file: %v\n", err)
	}
}
