package main

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"fmt"
	"image"
	_ "image/png"
)

//go:embed example.*
var exampleText string

//go:embed 42SnippetsLogo.png
var logoBytes []byte

//go:embed 42SnippetsLogo.png.gz
var compressedLogo []byte

func main() {

	fmt.Println("# 42Snippets")
	fmt.Println("## Golang Embed")

	/*
		This Go code demonstrates the use of the embed package to include an external file's contents within the program. By leveraging the //go:embed directive, the content is stored as a string in the exampleText variable. The main function then prints the embedded file's contents, showcasing a simple yet effective use of Go's embed feature.

		Furthermore, a compressed image can be embedded into the binary and subsequently decompressed during runtime.
	*/

	fmt.Println("Contents of example.txt:")
	fmt.Println(exampleText)

	img, _, err := image.Decode(bytes.NewReader(logoBytes))
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}
	fmt.Printf("Image size: %d x %d\n", img.Bounds().Dx(), img.Bounds().Dy())

	// Decompress the embedded image
	gzipReader, err := gzip.NewReader(bytes.NewReader(compressedLogo))
	if err != nil {
		fmt.Println("Error creating gzip reader:", err)
		return
	}
	defer gzipReader.Close()

	// Decode the image
	imgDecoded, _, err := image.Decode(gzipReader)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	// Use the image
	fmt.Printf("Image size of the decompressed/decoded: %d x %d\n", imgDecoded.Bounds().Dx(), imgDecoded.Bounds().Dy())

}
