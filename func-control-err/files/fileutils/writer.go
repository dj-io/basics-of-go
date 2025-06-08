package fileutils

import "os"

func WriteToFile(filename string, content string) error {
	// Writefile expects a slice of bytes, since our function takes in a string as a param go uses the slice of bytes func to convert the string to a slice of bytes
	return os.WriteFile(filename, []byte(content), 0644) // WriteFile takes in 3 args, the filename, the content, and the permissions -> permissions are set to 0644 which means the file is readable by the owner and writable by the owner and readable by the group

}
