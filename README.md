# Gube

A golang library to use .cube LUT files.

## Usage

A convenience `ProcessImage` method is provided that will apply a LUT to an image and return a copy of the image with the changes.

Example:

```go
package main

import (
	"image/png"
	"log"
	"os"

	"github.com/flynn-nrg/gube/gube"
)

func main() {
	imageIn, err := os.Open("input.png")
	if err != nil {
		log.Fatal(err)
	}
	input, err := png.Decode(imageIn)

	cube, err := os.Open("Rec709_Kodak_2393_D65.cube")
	if err != nil {
		log.Fatal(err)
	}

	gb, err := gube.NewFromReader(cube)
	if err != nil {
		log.Fatal(err)
	}

	outImage, err := gb.ProcessImage(input)
	if err != nil {
		log.Fatal(err)
	}

	outFile, err := os.Create("output.png")
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(outFile, outImage)
	if err != nil {
		log.Fatal(err)
	}
}
```

Which will yield this result:

![Comparison before and after applying the Rec709_Kodak_2393_D65 3D LUT](./images/lut_before_after.png "Cornell box with LUT applied")
