/**
 * @Author Awen
 * @Date 2024/05/01
 * @Email wengaolng@gmail.com
 **/

package canvas

import (
	"image"
	"image/color"
)

// CreatePaletteCanvas is the canvas that creates the palette
func CreatePaletteCanvas(width, height int, colour []color.RGBA) Palette {
	p := []color.Color{
		color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00},
	}

	for _, co := range colour {
		p = append(p, co)
	}

	return NewPalette(image.Rect(0, 0, width, height), p)
}

// CreateNRGBACanvas is the canvas that creates the NRGBA
func CreateNRGBACanvas(width, height int, isAlpha bool) NRGBA {
	return NewNRGBA(image.Rect(0, 0, width, height), isAlpha)
}
