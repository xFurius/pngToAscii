package main

import (
	"fmt"
	"image/color"
	"image/png"
	"os"
)

func main() {
	file, err := os.Open("img.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	dec, err := png.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	save, err := os.Create("art.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer save.Close()

	lvls := []string{"#", "@", "%", "$", "&", "/", "=", "-", " "}

	for i := dec.Bounds().Min.Y; i < dec.Bounds().Max.Y; i++ {
		for j := dec.Bounds().Min.X; j < dec.Bounds().Max.X; j++ {
			pixel := color.GrayModel.Convert(dec.At(j, i)).(color.Gray)
			lvl := pixel.Y / 25
			if lvl >= 9 {
				lvl = 8
			}
			save.WriteString(lvls[lvl])
		}
		save.WriteString("\n")
	}
}
