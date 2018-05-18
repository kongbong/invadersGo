package globals

import (
	"fmt"
	"image"
	"os"
)

func PrintImage(img image.Image) {
	texture := GDriver.CreateTexture(img, 1.0)
	GDriver.Call(func() {
		GWindowImg.SetTexture(texture)
	})
}

func GetImage(filePath string) image.Image {
	imgFile, err := os.Open(filePath)
	defer imgFile.Close()
	if err != nil {
		fmt.Println("Cannot read file:", err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println("Cannot decode file:", err)
	}
	return img
}
