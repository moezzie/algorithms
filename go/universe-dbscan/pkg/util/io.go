package util

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"strings"

	"golang.org/x/image/tiff"
)

func LoadImage(filePath string) image.Image {

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %s %v\n", filePath, err)
		panic(err)
	}
	defer file.Close()

	var img image.Image
	if strings.HasSuffix(filePath, "tif") {
		img = decodeTif(file)
	} else if strings.HasSuffix(filePath, "jpg") {
		img = decodeJpg(file)
	}

	img = toBW(img)

	return img
}

func decodeTif(file *os.File) image.Image {
	img, err := tiff.Decode(file)
	if err != nil {
		fmt.Printf("Error decoding TIFF: %v\n", err)
		panic(err)
	}

	return img
}

func decodeJpg(file *os.File) image.Image {
	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Printf("Error decoding JPEG: %v\n", err)
		panic(err)
	}

	return img
}
