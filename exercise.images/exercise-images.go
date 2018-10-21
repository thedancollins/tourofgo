package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	startAtX, startAtY int
	height, width      int
	blue, alpha        uint8
}

func (xi Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (xi Image) At(x, y int) color.Color {
	c := uint8(x ^ y)
	return color.RGBA{c, c, xi.blue, xi.alpha}
}

func (xi Image) Bounds() image.Rectangle {
	return image.Rect(xi.startAtX, xi.startAtY, xi.height, xi.width)
}

/*func ShowImages(m []image.Image) {
	var buf bytes.Buffer
	var err error
	var enc string
	for i, v := range m {
		err = png.Encode(&buf, v)
		if err != nil {
			panic(err)
		}
		enc = base64.StdEncoding.EncodeToString(buf.Bytes())
		fmt.Println("IMAGE:" + enc)
	}
}*/

func main() {
	m1 := Image{0, 0, 256, 256, 255, 255}
	//m2 := Image{0, 0, 256, 256, 155, 255}
	//m3 := Image{0, 0, 256, 256, 255, 155}
	//pic.ShowImage(m1)
	//pic.ShowImage(m2)
	//images := [3]image.Image{m1, m2, m3}
	pic.ShowImage(m1)
}
