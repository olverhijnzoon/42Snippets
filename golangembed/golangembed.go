package main

import (
	_ "embed"
	"fmt"
)

//go:embed example.*
var exampleText string

func main() {

	fmt.Println("# 42Snippets")
	fmt.Println("## Golag Embed")

	/*
		This Go code demonstrates the use of the embed package to include an external file's contents within the program. By leveraging the //go:embed directive, the content is stored as a string in the exampleText variable. The main function then prints the embedded file's contents, showcasing a simple yet effective use of Go's embed feature.
	*/

	fmt.Println("Contents of example.txt:")
	fmt.Println(exampleText)
}
