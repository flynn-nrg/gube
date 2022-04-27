package gube

import (
	"image"
	"image/color"
)

func (gi *GubeImpl) ProcessImage(input image.Image) (image.Image, error) {
	floatImage := toFloat(input)

	sizeX := input.Bounds().Max.X
	sizeY := input.Bounds().Max.Y

	i := 0
	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			r := floatImage[i]
			g := floatImage[i+1]
			b := floatImage[i+2]
			rgb, err := gi.LookUp(r, g, b)
			if err != nil {
				return nil, err
			}
			floatImage[i] = rgb[0]
			floatImage[i+1] = rgb[1]
			floatImage[i+2] = rgb[2]
			i += 3
		}
	}

	return toImage(floatImage, sizeX, sizeY), nil
}

func toFloat(i image.Image) []float64 {
	var data []float64
	sizeX := i.Bounds().Max.X
	sizeY := i.Bounds().Max.Y
	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			pixel := color.NRGBAModel.Convert(i.At(x, y)).(color.NRGBA)
			r := pixel.R
			g := pixel.G
			b := pixel.B
			data = append(data, float64(r)/255.0, float64(g)/255.0, float64(b)/255.0)
		}
	}

	return data
}

func toImage(data []float64, sizeX int, sizeY int) image.Image {
	var i int
	imageData := image.NewNRGBA(image.Rectangle{Max: image.Point{X: sizeX, Y: sizeY}})
	for y := 0; y < sizeY; y++ {
		for x := 0; x < sizeX; x++ {
			r := data[i]
			i++
			g := data[i]
			i++
			b := data[i]
			i++
			pixel := color.NRGBA{R: uint8(r * 255), G: uint8(g * 255), B: uint8(b * 255), A: 255}
			imageData.Set(x, y, pixel)
		}
	}

	return imageData
}
