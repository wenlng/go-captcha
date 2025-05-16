/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package canvas

import (
	"image"
	"image/color"
)

// CreatePaletteCanvas creates a canvas with a palette
func CreatePaletteCanvas(width, height int, colour []color.RGBA) Palette {
	p := make([]color.Color, 0, len(colour)+1)
	p = append(p, color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0x00})

	for _, co := range colour {
		p = append(p, co)
	}

	return NewPalette(image.Rect(0, 0, width, height), p)
}

// CreateNRGBACanvas creates an NRGBA canvas
func CreateNRGBACanvas(width, height int, isAlpha bool) NRGBA {
	return NewNRGBA(image.Rect(0, 0, width, height), isAlpha)
}
