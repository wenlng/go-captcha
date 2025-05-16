/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package canvas

import (
	"image/color"

	"github.com/golang/freetype/truetype"
)

// AreaRect struct for defining a rectangular area
type AreaRect struct {
	MinX, MaxX, MinY, MaxY int
}

// MakeAreaRect creates an area rectangle
func MakeAreaRect(x0, y0, x1, y1 int) *AreaRect {
	return &AreaRect{
		MinX: x0,
		MaxX: x1,
		MinY: y0,
		MaxY: y1,
	}
}

// PositionRect struct for defining a rectangle's position and size
type PositionRect struct {
	X, Y, Width, Height int
}

// MakePositionRect creates a position rectangle
func MakePositionRect(x, y, h, w int) *PositionRect {
	return &PositionRect{
		X:      x,
		Y:      y,
		Height: h,
		Width:  w,
	}
}

// DrawStringParams struct for string drawing parameters
type DrawStringParams struct {
	Color   color.Color
	Size    int
	Width   int
	Height  int
	FontDPI int
	Text    string
	Font    *truetype.Font
}
