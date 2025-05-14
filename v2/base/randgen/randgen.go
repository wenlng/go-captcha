/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package randgen

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/golang/freetype/truetype"
	"github.com/wenlng/go-captcha/v2/base/helper"
	"github.com/wenlng/go-captcha/v2/base/random"
)

// RandFont randomly selects a font
func RandFont(fonts []*truetype.Font) *truetype.Font {
	index := helper.RandIndex(len(fonts))
	if index < 0 {
		return nil
	}

	return fonts[index]
}

// RandHexColor randomly selects a hex color
func RandHexColor(colors []string) string {
	index := helper.RandIndex(len(colors))
	if index < 0 {
		return ""
	}

	return colors[index]
}

// RandImage randomly selects an image
func RandImage(images []image.Image) image.Image {
	index := helper.RandIndex(len(images))
	if index < 0 {
		return nil
	}

	return images[index]
}

// RandString randomly selects a string
func RandString(chars []string) string {
	k := rand.Intn(len(chars))
	return chars[k]
}

// RandColor randomly selects an RGBA color
func RandColor(co []color.Color) color.RGBA {
	colorLen := len(co)
	index := random.RandInt(0, colorLen)
	if index >= colorLen {
		index = colorLen - 1
	}

	r, g, b, a := co[index].RGBA()
	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
}

// RangCutImagePos randomly selects an image cropping position
func RangCutImagePos(width int, height int, img image.Image) image.Point {
	b := img.Bounds()
	iW := b.Max.X
	iH := b.Max.Y
	curX := 0
	curY := 0

	if iW-width > 0 {
		curX = random.RandInt(0, iW-width)
	}
	if iH-height > 0 {
		curY = random.RandInt(0, iH-height)
	}

	return image.Point{
		X: curX,
		Y: curY,
	}
}
