/**
 * @Author Awen
 * @Description Captcha Palette
 * @Date 2021/7/18
 * @Email wengaolng@gmail.com
 **/

package captcha

import (
	"image"
	"image/color"
	"math"
)

// Point is a type
/**
 * @Description: 点
 */
type Point struct {
	X, Y int
}

// Palette is a type
/**
 * @Description: 调色板
 */
type Palette struct {
	*image.Paletted
}

// NewPalette is a function
/**
 * @Description: 创建调色板
 * @param r
 * @param p
 * @return *Palette
 */
func NewPalette(r image.Rectangle, p color.Palette) *Palette {
	return &Palette{
		image.NewPaletted(r, p),
	}
}

// Rotate is a function
/**
 * @Description: 旋转任意角度
 * @receiver p
 * @param angle
 */
func (p *Palette) Rotate(angle int) {
	tarImg := p
	width := tarImg.Bounds().Max.X
	height := tarImg.Bounds().Max.Y
	r := width / 2
	retImg := image.NewPaletted(image.Rect(0, 0, width, height), tarImg.Palette)
	for x := 0; x <= retImg.Bounds().Max.X; x++ {
		for y := 0; y <= retImg.Bounds().Max.Y; y++ {
			tx, ty := p.angleSwapPoint(float64(x), float64(y), float64(r), float64(angle))
			retImg.SetColorIndex(x, y, tarImg.ColorIndexAt(int(tx), int(ty)))
		}
	}

	nW := retImg.Bounds().Max.X
	nH := retImg.Bounds().Max.Y
	for x := 0; x < nW; x++ {
		for y := 0; y < nH; y++ {
			p.SetColorIndex(x, y, retImg.ColorIndexAt(x, y))
		}
	}
}

/**
 * @Description: 画圆点
 * @receiver p
 * @param x
 * @param y
 * @param radius
 * @param colorIdx
 */
func (p *Palette) drawCircle(x, y, radius int, colorIdx uint8) {
	f := 1 - radius
	dfx := 1
	dfy := -2 * radius
	xo := 0
	yo := radius

	p.SetColorIndex(x, y + radius, colorIdx)
	p.SetColorIndex(x, y - radius, colorIdx)
	p.drawHorizLine(x - radius, x + radius, y, colorIdx)

	for xo < yo {
		if f >= 0 {
			yo--
			dfy += 2
			f += dfy
		}
		xo++
		dfx += 2
		f += dfx
		p.drawHorizLine(x - xo, x + xo, y + yo, colorIdx)
		p.drawHorizLine(x - xo, x + xo, y - yo, colorIdx)
		p.drawHorizLine(x - yo, x + yo, y + xo, colorIdx)
		p.drawHorizLine(x - yo, x + yo, y - xo, colorIdx)
	}
}

/**
 * @Description: 画线
 * @receiver p
 * @param fromX
 * @param toX
 * @param y
 * @param colorIdx
 */
func (p *Palette) drawHorizLine(fromX, toX, y int, colorIdx uint8) {
	for x := fromX; x <= toX; x++ {
		p.SetColorIndex(x, y, colorIdx)
	}
}

/**
 * @Description: 扭曲
 * @receiver p
 * @param amplude
 * @param period
 */
func (p *Palette) distort(amplude float64, period float64) {
	w := p.Bounds().Max.X
	h := p.Bounds().Max.Y
	newP := NewPalette(image.Rect(0, 0, w, h), p.Palette)
	dx := 2.0 * math.Pi / period
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			xo := amplude * math.Sin(float64(y) * dx)
			yo := amplude * math.Cos(float64(x) * dx)
			newP.SetColorIndex(x, y, p.ColorIndexAt(x + int(xo), y + int(yo)))
		}
	}

	nW := newP.Bounds().Max.X
	nH := newP.Bounds().Max.Y
	for x := 0; x < nW; x++ {
		for y := 0; y < nH; y++ {
			p.SetColorIndex(x, y, newP.ColorIndexAt(x, y))
		}
	}

	newP.Palette = nil
}

/**
 * @Description: 画点到点直线
 * @receiver p
 * @param point1
 * @param point2
 * @param lineColor
 */
func (p *Palette) drawBeeline(point1 Point, point2 Point, lineColor color.RGBA) {
	dx := math.Abs(float64(point1.X - point2.X))
	dy := math.Abs(float64(point2.Y - point1.Y))
	sx, sy := 1, 1
	if point1.X >= point2.X {
		sx = -1
	}
	if point1.Y >= point2.Y {
		sy = -1
	}
	err := dx - dy
	for {
		p.Set(point1.X, point1.Y, lineColor)
		p.Set(point1.X + 1, point1.Y, lineColor)
		p.Set(point1.X - 1, point1.Y, lineColor)
		p.Set(point1.X + 2, point1.Y, lineColor)
		p.Set(point1.X - 2, point1.Y, lineColor)
		if point1.X == point2.X && point1.Y == point2.Y {
			return
		}
		e2 := err * 2
		if e2 > -dy {
			err -= dy
			point1.X += sx
		}
		if e2 < dx {
			err += dx
			point1.Y += sy
		}
	}
}

/**
 * @Description: 角度转换点坐标
 * @receiver p
 * @param x
 * @param y
 * @param r
 * @param angle
 * @return tarX
 * @return tarY
 */
func (p *Palette) angleSwapPoint(x, y, r, angle float64) (tarX, tarY float64) {
	x -= r
	y = r - y
	sinVal := math.Sin(angle * (math.Pi / 180))
	cosVal := math.Cos(angle * (math.Pi / 180))
	tarX = x * cosVal + y * sinVal
	tarY = -x * sinVal + y * cosVal
	tarX += r
	tarY = r - tarY
	return
}
