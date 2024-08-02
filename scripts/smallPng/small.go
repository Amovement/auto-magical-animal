package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	// 打开源图像文件
	inputFile, err := os.Open("input.png")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// 解码源图像
	srcImg, _, err := image.Decode(inputFile)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	// 设置缩放比例
	newWidth := uint(25)   // 新的宽度
	newHeight := uint(600) // 新的高度
	// 这里可以根据需要设置缩放比例，例如保持宽高比
	newHeight = uint(float64(srcImg.Bounds().Dy()) * float64(newWidth) / float64(srcImg.Bounds().Dx()))

	// 等比缩放图像
	resizedImg := resize.Resize(newWidth, newHeight, srcImg, resize.Lanczos3)

	// 创建输出图像文件
	outputFile, err := os.Create("output.png")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// 将缩放后的图像编码并保存为PNG文件
	err = png.Encode(outputFile, resizedImg)
	if err != nil {
		fmt.Println("Error encoding image:", err)
		return
	}

	fmt.Println("Image has been resized and saved as output.png")
}
