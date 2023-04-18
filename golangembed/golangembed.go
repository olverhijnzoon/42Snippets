package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	_ "image/png"
)

//go:embed example.*
var exampleText string

//go:embed 42SnippetsLogo.png
var logoBytes []byte

func main() {

	fmt.Println("# 42Snippets")
	fmt.Println("## Golang Embed")

	/*
		This Go code demonstrates the use of the embed package to include an external file's contents within the program. By leveraging the //go:embed directive, the content is stored as a string in the exampleText variable. The main function then prints the embedded file's contents, showcasing a simple yet effective use of Go's embed feature.
	*/

	fmt.Println("Contents of example.txt:")
	fmt.Println(exampleText)

	img, _, err := image.Decode(bytes.NewReader(logoBytes))
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}
	fmt.Printf("Image size: %d x %d\n", img.Bounds().Dx(), img.Bounds().Dy())

}
