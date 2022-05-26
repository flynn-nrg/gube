package gube

import (
	"image"
	"image/color"
)

func (gi *GubeImpl) ProcessImage(input image.Image) (image.Image, error) {
	output := image.NewNRGBA(image.Rectangle{
		Min: image.Point{X: input.Bounds().Min.X, Y: input.Bounds().Min.Y},
		Max: image.Point{X: input.Bounds().Max.X, Y: input.Bounds().Max.Y}})

	for y := input.Bounds().Min.Y; y <= input.Bounds().Max.Y; y++ {
		for x := input.Bounds().Min.X; x <= input.Bounds().Max.X; x++ {
			pixel := color.NRGBAModel.Convert(input.At(x, y)).(color.NRGBA)
			dMin, dMax := gi.Domain()
			r := dMin[0] + (float64(pixel.R)/255.0)*(dMax[0]-dMin[0])
			g := dMin[1] + (float64(pixel.G)/255.0)*(dMax[1]-dMin[1])
			b := dMin[2] + (float64(pixel.B)/255.0)*(dMax[2]-dMin[2])
			rgb, err := gi.LookUp(r, g, b)
			if err != nil {
				return nil, err
			}
			outPixel := color.NRGBA{
				R: uint8(((rgb[0] - dMin[0]) / (dMax[0] - dMin[0])) * 255.0),
				G: uint8(((rgb[1] - dMin[1]) / (dMax[1] - dMin[1])) * 255.0),
				B: uint8(((rgb[2] - dMin[2]) / (dMax[2] - dMin[2])) * 255.0),
				A: pixel.A}
			output.Set(x, y, outPixel)
		}
	}

	return output, nil
}
