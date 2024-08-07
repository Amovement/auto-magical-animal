package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	// Open the source image file
	inputFile, err := os.Open("input.png")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Decoded source image
	srcImg, _, err := image.Decode(inputFile)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	// Set scale
	newWidth := uint(15)   // New width
	newHeight := uint(600) // New height
	// Here you can set the scale as needed, for example to maintain the aspect ratio
	newHeight = uint(float64(srcImg.Bounds().Dy()) * float64(newWidth) / float64(srcImg.Bounds().Dx()))

	// Scale the image in equal proportions
	resizedImg := resize.Resize(newWidth, newHeight, srcImg, resize.Lanczos3)

	// Create the output image file
	outputFile, err := os.Create("output.png")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Encode and save the scaled image as a PNG file
	err = png.Encode(outputFile, resizedImg)
	if err != nil {
		fmt.Println("Error encoding image:", err)
		return
	}

	fmt.Println("Image has been resized and saved as output.png")
}
