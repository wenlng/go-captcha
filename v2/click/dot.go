/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package click

import (
	"image"

	"github.com/golang/freetype/truetype"
)

// Dot represents a single point (character or shape) in the captcha
type Dot struct {
	Index  int    `json:"index"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Size   int    `json:"size"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Text   string `json:"text"`
	Shape  string `json:"shape"`
	Angle  int    `json:"angle"`
	Color  string `json:"color"`
	Color2 string `json:"color2"`
}

// DrawDot represents the dot data used for drawing
type DrawDot struct {
	Dot              *Dot
	X                int
	Y                int
	FontDPI          int
	Text             string
	Image            image.Image
	UseOriginalColor bool
	Size             int
	Width            int
	Height           int
	Angle            int
	Color            string
	Color2           string
	Font             *truetype.Font
	DrawType         DrawType
}
