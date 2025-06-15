package util

import (
	"fmt"
	"hash/crc32"
	"image"
	"image/color"
	"image/draw"
)

func toBW(img image.Image) image.Image {
	bounds := img.Bounds()
	// Create a new grayscale image
	grayImg := image.NewGray(bounds)

	// Convert each pixel to grayscale
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Get the color at each pixel
			oldPixel := img.At(x, y)
			// Convert to grayscale using the standard formula
			grayPixel := color.GrayModel.Convert(oldPixel)
			grayImg.Set(x, y, grayPixel)
		}
	}

	return grayImg
}

func ConvertLabelsToImage(data [][]int) image.Image {
	width := len(data)
	height := len(data[0])

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			brightnes := data[x][y]

			if brightnes <= 0 {
				// Pixel is already 0 alpha
				continue
			}

			position := (y-img.Rect.Min.Y)*img.Stride + (x-img.Rect.Min.X)*4
			r, g, b := randomCol(int(brightnes))
			img.Pix[position+0] = r
			img.Pix[position+1] = g
			img.Pix[position+2] = b
			img.Pix[position+3] = 255
		}
	}

	return img
}

func randomCol(cluster int) (uint8, uint8, uint8) {
	checksum := crc32.ChecksumIEEE([]byte(fmt.Sprintf("%d", cluster)))

	r := uint8((checksum >> 16) & 0xFF) // Extract red
	g := uint8((checksum >> 8) & 0xFF)  // Extract green
	b := uint8(checksum & 0xFF)         // Extract blue

	return r, g, b
}

func Blend(background, overlay image.Image) image.Image {
	bounds := background.Bounds()
	result := image.NewRGBA(bounds)

	// Draw the background first
	draw.Draw(result, bounds, background, bounds.Min, draw.Src)

	// Create a semi-transparent version of the overlay
	mask := image.NewUniform(color.Alpha{128}) // 50% opacity

	// Draw the overlay with the opacity mask
	draw.DrawMask(result, bounds, overlay, bounds.Min, mask, bounds.Min, draw.Over)

	return result
}
