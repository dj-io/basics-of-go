package fileutils

import "os"


// capital R sets this function as public, it 'exports' the function
func ReadTextFile(filename string) (string, error) { // we can return a string or an error, following the go error design pattern
	content, err := os.ReadFile(filename)

	if err != nil {
		// we couldnt read the file
		return "", err // 1 of 3 ways to handle errors in go return empty string as the value and the error
		// panic("HAHHHAHAHAAA") // 2 of 3 ways to handle errors in go (panic, the program will crash)
       // return "" // 3 out of 3 ways to handle errors in go (lazy way, user will not know the error)
	} else {
		// read operation was successful
		return string(content), nil // using the string func to convert the bytes to a string (UTF-8) and we return nil as the error
	}
}
